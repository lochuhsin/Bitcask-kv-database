package scheduler

import (
	"bufio"
	"fmt"
	"os"
	"rebitcask/internal/dao"
	"rebitcask/internal/memory"
	"rebitcask/internal/segment"
	"rebitcask/internal/setting"
	"rebitcask/internal/util"
	"strings"
)

func memBlockToFile(memBlock *memory.Block) segment.Segment {
	/**
	 * Note, assuming that key in entries are sorted in ascending order
	 */
	blockId := string(memBlock.Id)
	entryList := memBlock.GetAll()

	filePath := util.GetSegmentFilePath(blockId)
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777) //TODO: optimize the mode
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	curroffset := 0
	pIndex := segment.NewSegmentIndex(blockId)
	for _, p := range entryList {
		data, err := dao.Serialize(p)
		if err != nil {
			panic("Error while serializing data")
		}
		offset, err := writer.WriteString(data + setting.DATA_SEPARATOR)
		if err != nil {
			panic("something went wrong while writing to segment")
		}
		// offset minus data saparater = the length of the data
		pIndex.Set(p.Key, curroffset, offset-len([]byte(setting.DATA_SEPARATOR)))
		curroffset += offset
	}
	writer.Flush()
	file.Sync()
	segment := segment.NewSegment(blockId, &pIndex, entryList[0].Key, len(entryList))
	return segment
}

func createSegMetaFile(sId string, level int) {
	filePath := util.GetSegmentMetaDataFilePath(sId)
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777) //TODO: optimize the mode
	if err != nil {
		panic(err)
	}
	defer file.Close()

	/**
	 * Currently only store level information for segment manager to backup
	 */
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(fmt.Sprintf("level::%v", level))
	if err != nil {
		panic("something went wrong while writing segment metadata")
	}
	writer.Flush()
	// We don't need to fd.Sync() metadata, since the read is not necessarily to do
	// immediately read, like Get operation
}

func createSegIndexFile(sId string, pIndex *segment.PrimaryIndex) {
	filePath := util.GetSegmentIndexFilePath(sId)
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777) //TODO: optimize the mode
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for key, val := range pIndex.OffsetMap {
		data := segmentIndexSerialize(key, val.Format())
		_, err := writer.WriteString(data + setting.DATA_SEPARATOR)
		if err != nil {
			panic("something went wrong while writing to segment")
		}
	}

	writer.Flush()
	file.Sync()
}

// TODO: refactor this
func segmentIndexSerialize(key string, val string) string {
	// format -> Key::offset::length
	var builder strings.Builder
	builder.WriteString(key)
	builder.WriteString("::")
	builder.WriteString(val)
	return builder.String()
}
