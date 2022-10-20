package internal

import (
	"fmt"
	"io"
	"os"
)

func toDisk(memory *memoryMap, currSeg *CurrentSegmentMap, segContainer *SegmentContainer) error {
	filepath := fmt.Sprintf("%v%v/%v.log", LOGFOLDER, SEGMENTFOLDER, currSeg.CurrentSegmentNo)
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}

	defer file.Close()

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

		// TODO: Decouple this function
		if byteHeadPosition >= FILEBYTELIMIT {
			// close file
			file.Close()

			// write diskmap to map
			segContainer.memo = append(segContainer.memo, *currSeg)

			// create new obj
			currentSegmentNo := currSeg.CurrentSegmentNo
			*currSeg = CurrentSegmentMap{
				bytePositionMap:  make(map[string]int),
				byteLengthMap:    make(map[string]int),
				byteFileLength:   0,
				CurrentSegmentNo: currentSegmentNo + 1,
			}
			// open new file
			filepath := fmt.Sprintf("%v%v/%v.log", LOGFOLDER, SEGMENTFOLDER, currSeg.CurrentSegmentNo)
			file, err = os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)

			// new bytehead
			byteHeadPosition = 0
		}
	}

	currSeg.byteFileLength = byteHeadPosition
	memory.keyvalue = make(map[string]string)
	return nil
}

func isKeyInSegment(k string, currSeg *CurrentSegmentMap) (v string, status bool) {

	if _, ok := currSeg.bytePositionMap[k]; !ok {
		return "", false
	}

	filepath := fmt.Sprintf("%v%v/%v.log", LOGFOLDER, SEGMENTFOLDER, currSeg.CurrentSegmentNo)
	bytePos, _ := currSeg.bytePositionMap[k]
	byteLen, _ := currSeg.byteLengthMap[k]

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

func isKeySegmentContainer(k string, segContainer *SegmentContainer) (v string, status bool) {

	// read backwards since SegNo. bigger means later
	for i := len(segContainer.memo) - 1; i >= 0; i-- {
		val, ok := isKeyInSegment(k, &segContainer.memo[i])
		if ok {
			return val, true
		}
	}
	return "", false
}

// TODO: Convert to using total byte
func isExceedMemoLimit(m *memoryMap) bool {
	return len(m.keyvalue) >= MEMORYLIMIT
}
