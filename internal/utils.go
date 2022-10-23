package internal

import (
	"fmt"
	"io"
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
		panic("something wrong with opening file")
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

func seekFile(file *os.File, byteHead int, byteLen int) (bytes []byte) {
	_, err := file.Seek(int64(byteHead), io.SeekStart)
	if err != nil {
		panic("Something went wrong while seeking file")
	}

	readByte := make([]byte, byteLen)

	_, err = file.Read(readByte)
	if err != nil {
		panic("Something went wrong while seeking file")
	}
	return readByte
}
