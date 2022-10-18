package internal

import (
	"fmt"
	"io"
	"os"
)

// TODO: Convert this to env file
var m memoryMap
var d CurrentSegmentMap
var s DiskSegmentMap
var LOGFOLDER = "./log/"
var SEGMENTFOLDER = "seg/"
var MEMORYLIMIT = 2
var FILEBYTELIMIT = 2

func init() {
	// create log storage folder

	// TODO: Convert this to env file
	_ = os.RemoveAll(LOGFOLDER)
	_ = os.MkdirAll(fmt.Sprintf("%v%v", LOGFOLDER, SEGMENTFOLDER), 0700)
	initMaps()
}

func initMaps() {
	m.keyvalue = make(map[string]string)
	d.bytePositionMap = make(map[string]int)
	d.byteLengthMap = make(map[string]int)
	d.byteFileLength = 0
	d.CurrentSegmentNo = 0
	s.memo = []CurrentSegmentMap{}

}

func Get(k string) (v string, status bool) {

	if val, ok := m.keyvalue[k]; ok {
		return val, true
	}
	if val, ok := isKeyInSegment(k, &d); ok {
		return val, true
	}

	if val, ok := isKeyInSegments(k, &s); ok {
		return val, true
	}
	return "", false
}

func Set(k string, v string) error {
	m.keyvalue[k] = v
	if isExceedMemoLimit(&m) {
		err := toDisk(&m, &d)
		if err != nil {
			fmt.Println(err)
		}
	}
	return nil
}

func GetLength() int {
	return len(m.keyvalue)
}

func GetAllInMemory() map[string]string {
	return m.keyvalue
}

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
		if byteHeadPosition > FILEBYTELIMIT {
			// close file
			file.Close()

			// write diskmap to map
			splitSegment(*d, &s)

			// create new obj
			*d = CurrentSegmentMap{
				bytePositionMap:  make(map[string]int),
				byteLengthMap:    make(map[string]int),
				byteFileLength:   0,
				CurrentSegmentNo: d.CurrentSegmentNo + 1,
			}

			// open new file
			filepath := fmt.Sprintf("%v%v/%v.log", LOGFOLDER, SEGMENTFOLDER, d.CurrentSegmentNo)
			fmt.Println(d.CurrentSegmentNo)
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

	return string(readByte), true
}

func isKeyInSegments(k string, s *DiskSegmentMap) (v string, status bool) {
	for i := len(s.memo) - 1; i >= 0; i-- {
		val, ok := isKeyInSegment(k, &s.memo[i])
		if ok {
			return val, true
		}
	}
	return "", false
}

func isExceedMemoLimit(m *memoryMap) bool {
	return len(m.keyvalue) >= MEMORYLIMIT
}
