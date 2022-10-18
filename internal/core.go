package internal

import (
	"fmt"
	"io"
	"os"
)

func toDisk(m *memoryMap, d *CurrentSegmentMap) error {
	filepath := fmt.Sprintf("%v%v/%v.log", LOGFOLDER, SEGMENTFOLDER, d.CurrentSegmentNo)
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}

	defer file.Close()

	byteHeadPosition := d.byteFileLength
	for k, v := range m.keyvalue {

		byteValue := []byte(v)
		bytes, err := file.Write(byteValue)

		if err != nil {
			panic("Something went wrong while writing to disk")
		}

		d.byteLengthMap[k] = bytes
		d.bytePositionMap[k] = byteHeadPosition

		byteHeadPosition += bytes

		// TODO: Decouple this function
		if byteHeadPosition >= FILEBYTELIMIT {
			// close file
			file.Close()

			// write diskmap to map
			splitSegment(*d, &s)

			// create new obj
			currentSegmentNo := d.CurrentSegmentNo
			*d = CurrentSegmentMap{
				bytePositionMap:  make(map[string]int),
				byteLengthMap:    make(map[string]int),
				byteFileLength:   0,
				CurrentSegmentNo: currentSegmentNo + 1,
			}
			// open new file
			filepath := fmt.Sprintf("%v%v/%v.log", LOGFOLDER, SEGMENTFOLDER, d.CurrentSegmentNo)
			file, err = os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)

			// new bytehead
			byteHeadPosition = d.byteFileLength
		}
	}

	d.byteFileLength = byteHeadPosition
	m.keyvalue = make(map[string]string)
	return nil
}

// copy CurrentSegmentMap
func splitSegment(d CurrentSegmentMap, s *DiskSegmentMap) {
	s.memo = append(s.memo, d)
}

func isKeyInSegment(k string, d *CurrentSegmentMap) (v string, status bool) {

	if _, ok := d.bytePositionMap[k]; !ok {
		return "", false
	}

	filepath := fmt.Sprintf("%v%v/%v.log", LOGFOLDER, SEGMENTFOLDER, d.CurrentSegmentNo)
	bytePos, _ := d.bytePositionMap[k]
	byteLen, _ := d.byteLengthMap[k]

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

	if string(readByte) == TOMBSTONE {
		return "", false
	}
	return string(readByte), true
}

func isKeyInSegments(k string, s *DiskSegmentMap) (v string, status bool) {

	// read backwards since SegNo. bigger means later
	for i := len(s.memo) - 1; i >= 0; i-- {
		val, ok := isKeyInSegment(k, &s.memo[i])
		if ok {
			//fmt.Println("number of segments", i)
			//fmt.Println(s.memo[i].byteLengthMap)
			//fmt.Println(s.memo[i].bytePositionMap)
			//fmt.Println(s.memo[i].CurrentSegmentNo)
			fmt.Println(s.memo)
			return val, true
		}
	}
	return "", false
}

// TODO: Convert to using total byte
func isExceedMemoLimit(m *memoryMap) bool {
	return len(m.keyvalue) >= MEMORYLIMIT
}
