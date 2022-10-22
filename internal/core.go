package internal

import (
	"fmt"
	"io"
	"os"
)

func toDisk(memory *memoryMap, currSeg *SegmentMap, segContainer *SegmentContainer) error {
	filepath := fmt.Sprintf("%v%v/%v.log", LOGFOLDER, SEGMENTFOLDER, currSeg.CurrentSegmentNo)
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}

	byteHeadPosition := currSeg.byteFileLength
	for k, v := range memory.keyvalue {

		byteValue := []byte(v)
		bytes, err := file.Write(byteValue)

		if err != nil {
			panic("Something went wrong while writing to disk")
		}

		currSeg.byteLengthMap[k] = bytes
		currSeg.bytePositionMap[k] = byteHeadPosition

		byteHeadPosition += bytes

		if byteHeadPosition >= FILEBYTELIMIT {
			// close file
			file.Close()

			// store segment
			segContainer.memo = append(segContainer.memo, *currSeg)

			// create new segment
			newSegmentNo := currSeg.CurrentSegmentNo + 1
			file, *currSeg = createNewSegment(newSegmentNo)

			// new bytehead
			byteHeadPosition = 0
		}
	}

	currSeg.byteFileLength = byteHeadPosition
	memory.keyvalue = make(map[string]string)
	file.Close()
	return nil
}

func isKeyInSegment(k string, segment *SegmentMap) (v string, status bool) {

	if _, ok := segment.bytePositionMap[k]; !ok {
		return "", false
	}

	filepath := fmt.Sprintf("%v%v/%v.log", LOGFOLDER, SEGMENTFOLDER, segment.CurrentSegmentNo)
	bytePos, _ := segment.bytePositionMap[k]
	byteLen, _ := segment.byteLengthMap[k]

	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		return "Something went wrong while opening file", false
	}

	_, err = file.Seek(int64(bytePos), io.SeekStart)
	if err != nil {
		panic("Something went wrong while seeking file")
	}

	readByte := make([]byte, byteLen)

	_, err = file.Read(readByte)
	if err != nil {
		panic("Something went wrong while reading file")
	}

	return string(readByte), true
}
