package internal

import (
	"errors"
	"fmt"
	"os"
)

// TODO: Convert this to env file
var m memoryMap
var d CurrentSegmentMap
var s DiskSegmentMap
var LOGFOLDER = "./log/"
var SEGMENTFOLDER = "seg/"
var MEMORYLIMIT = 10000
var FILEBYTELIMIT = 10000
var TOMBSTONE = "!@#$%^&*()_+"

func init() {
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

	// check if is value in memory
	if val, ok := m.keyvalue[k]; ok {
		if val == TOMBSTONE {
			return "", false
		}
		return val, true
	}

	// check in current segment
	if val, ok := isKeyInSegment(k, &d); ok {
		if val == TOMBSTONE {
			return "", false
		}
		return val, true
	}

	// check in history segment
	if val, ok := isKeyInSegments(k, &s); ok {
		if val == TOMBSTONE {
			return "", false
		}
		return val, true
	}

	return "", false
}

func Set(k string, v string) error {
	if k == TOMBSTONE {
		return errors.New("invalid input")
	}

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
