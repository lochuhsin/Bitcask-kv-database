package internal

import (
	"errors"
	"fmt"
	"os"
)

// TODO: Convert this to env file
var memory memoryMap
var currentSeg CurrentSegmentMap
var segContainer SegmentContainer
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
	memory.keyvalue = make(map[string]string)
	currentSeg.bytePositionMap = make(map[string]int)
	currentSeg.byteLengthMap = make(map[string]int)
	currentSeg.byteFileLength = 0
	currentSeg.CurrentSegmentNo = 0
	segContainer.memo = []CurrentSegmentMap{}

}

func Get(k string) (v string, status bool) {

	// check if is value in memory
	if val, ok := memory.keyvalue[k]; ok {
		if val == TOMBSTONE {
			return "", false
		}
		return val, true
	}

	// check in current segment
	if val, ok := isKeyInSegment(k, &currentSeg); ok {
		if val == TOMBSTONE {
			return "", false
		}
		return val, true
	}

	// check in history segment
	if val, ok := isKeySegmentContainer(k, &segContainer); ok {
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

	memory.keyvalue[k] = v
	if isExceedMemoLimit(&memory) {
		err := toDisk(&memory, &currentSeg, &segContainer)
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
	return len(memory.keyvalue)
}

func GetAllInMemory() map[string]string {
	return memory.keyvalue
}
