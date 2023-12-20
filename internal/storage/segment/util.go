package segment

import (
	"bufio"
	"fmt"
	"os"
	"rebitcask/internal/settings"
	"rebitcask/internal/storage/dao"
)

func getSegmentFilePath(segId string) string {
	return fmt.Sprintf("%v%v%v%v", settings.ENV.LogPath, settings.ENV.SegmentFolder, segId, settings.ENV.SegmentFileExt)
}

func writeSegmentToFile(s *Segment, sIndex *SegmentIndex, pairs []dao.Pair) {
	/**
	 * Note, assuming that key in pairs are sorted in ascending order
	 */
	filePath := getSegmentFilePath(s.id)
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	curroffset := 0
	s.smallestKey = pairs[0].Key.Val // the first key is the smallest value
	for _, p := range pairs {
		data, err := dao.Serialize(p)
		if err != nil {
			panic("Error while serializing data")
		}
		offset, err := writer.WriteString(data + "\n") // Figure out a better way to split between keys
		if err != nil {
			panic("something went wrong while writing to segment")
		}
		sIndex.Set(p.Key, curroffset)
		curroffset += offset
	}
	writer.Flush()

	s.smallestKey = pairs[0].Key.GetVal().(string)
	s.keyCount = len(pairs)
}
