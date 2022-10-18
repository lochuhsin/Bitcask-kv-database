package internal

import (
	"fmt"
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
var TOMBSTONE = "!@#$% "

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
			return err
		}
	}
	return nil
}

func Delete(k string) error {
	return Set(k, TOMBSTONE)
}

func GetLength() int {
	return len(m.keyvalue)
}

func GetAllInMemory() map[string]string {
	return m.keyvalue
}
