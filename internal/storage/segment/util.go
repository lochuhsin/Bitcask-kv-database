package segment

import (
	"bufio"
	"fmt"
	"os"
	"rebitcask/internal/settings"
	"rebitcask/internal/storage/dao"
)

func getSegmentFilePath(segId string) string {
	return fmt.Sprintf("%v%v%v%v", settings.ENV.DataPath, settings.SEGMENT_FILE_FOLDER, segId, settings.SEGMENT_FILE_EXT)
}

func getSegmentIndexFilePath(segId string) string {
	return fmt.Sprintf("%v%v%v%v", settings.ENV.DataPath, settings.INDEX_FILE_FOLDER, segId, settings.SEGMENT_KEY_OFFSET_FILE_EXT)
}

func getSegmentIndexMetaDataFilePath(segId string) string {
	return fmt.Sprintf("%v%v%v%v", settings.ENV.DataPath, settings.SEGMENT_FILE_FOLDER, segId, settings.SEGMENT_FILE_METADATA_EXT)
}

func writeSegmentToFile(s *Segment, sIndex *PrimaryIndex, pairs []dao.Pair) {
	/**
	 * Note, assuming that key in pairs are sorted in ascending order
	 */
	filePath := getSegmentFilePath(s.id)
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777) //TODO: optimize the mode
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	curroffset := 0
	s.smallestKey = pairs[0].Key.Val // the first key is the smallest value
	for _, p := range pairs {        // TODO: convert this pair to generator pattern, hide inside segment, we don't need to know if the data needs to be serialized
		data, err := dao.Serialize(p)
		if err != nil {
			panic("Error while serializing data")
		}
		offset, err := writer.WriteString(data + settings.DATASAPARATER)
		if err != nil {
			panic("something went wrong while writing to segment")
		}
		// offset minus data saparater = the length of the data
		sIndex.Set(p.Key, curroffset, offset-len([]byte(settings.DATASAPARATER)))
		curroffset += offset
	}
	writer.Flush()

	s.smallestKey = pairs[0].Key.GetVal().(string)
	s.keyCount = len(pairs)
}

func writeSegmentMetadata(s *Segment) {
	filePath := getSegmentIndexMetaDataFilePath(s.id)
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777) //TODO: optimize the mode
	if err != nil {
		panic(err)
	}
	defer file.Close()
	/**
	 * Currently only store level information for segment manager to backup
	 */
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(fmt.Sprintf("level::%v", s.level))
	if err != nil {
		panic("something went wrong while writing segment metadata")
	}
	writer.Flush()
}

func writeSegmentIndexToFile(sIndex *PrimaryIndex) {
	filePath := getSegmentIndexFilePath(sIndex.id)
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777) //TODO: optimize the mode
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	offsetMap := sIndex.offsetMap

	for key, val := range offsetMap {
		data := segmentIndexSerialize(key.Format(), val.Format())
		_, err := writer.WriteString(data + settings.DATASAPARATER)
		if err != nil {
			panic("something went wrong while writing to segment")
		}
	}
	writer.Flush()
}

// TODO: refactor this
func segmentIndexSerialize(key string, val string) string {
	// format -> KeyDataType::KeyLen::Key::offset::length
	return fmt.Sprintf("%v::%v", key, val)
}
