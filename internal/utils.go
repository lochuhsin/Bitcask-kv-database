package internal

import (
	"fmt"
	"os"
)

func createNewSegment(newSegmentNo int) (file *os.File, segmentMap SegmentMap) {
	segmentMap = SegmentMap{
		bytePositionMap:  make(map[string]int),
		byteLengthMap:    make(map[string]int),
		byteFileLength:   0,
		CurrentSegmentNo: newSegmentNo,
	}
	filepath := fmt.Sprintf("%v%v/%v.log", LOGFOLDER, SEGMENTFOLDER, newSegmentNo)
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)

	if err != nil {

	}
	return file, segmentMap
}

func isExceedMemoLimit(memoSize int) bool {
	return memoSize >= MEMORYLIMIT
}

// TODO: find a better condition
func isSegFileMultiple(fileCount int) bool {
	return (fileCount % SEGFILECOUNTLIMIT) == 0
}

func filterTombStone(val string) (value string, status bool) {
	if val == TOMBSTONE {
		return "", false
	}
	return val, true
}
