package internal

import (
	"errors"
	"fmt"
	"os"
)

// TODO: Convert this to env file
var memory memoryMap
var currentSeg SegmentMap
var segContainer SegmentContainer
var envVar envVariables
var ENVPATH = "./rebitcask.env"

func init() {
	initGlobalEnvVar(ENVPATH)
	_ = os.RemoveAll(envVar.logFolder)
	_ = os.MkdirAll(fmt.Sprintf("%v%v", envVar.logFolder, envVar.segmentFolder), 0700)
	initMaps()
}

func initMaps() {
	memory.keyvalue = make(map[string][]byte)
	currentSeg.bytePositionMap = make(map[string]int)
	currentSeg.byteLengthMap = make(map[string]int)
	currentSeg.byteFileLength = 0
	currentSeg.CurrentSegmentNo = 0
	segContainer.memo = []SegmentMap{}

}

func Get(k string) (v string, status bool) {

	if val, ok := memory.keyvalue[k]; ok {
		str := string(val)
		return filterTombStone(str)
	}

	// check in current segment
	if val, ok := isKeyInSegment(k, &currentSeg); ok {
		str := string(val)
		return filterTombStone(str)
	}

	// check previous segments read backwards since SegNo. bigger means later
	for i := len(segContainer.memo) - 1; i >= 0; i-- {
		val, ok := isKeyInSegment(k, &segContainer.memo[i])
		str := string(val)
		if ok {
			return filterTombStone(str)
		}
	}
	return "", false
}

func Set(k string, v string) error {
	if k == envVar.tombstone {
		return errors.New("invalid input")
	}

	memory.keyvalue[k] = []byte(v)
	if isExceedMemoLimit(len(memory.keyvalue)) {
		err := toDisk(&memory, &currentSeg, &segContainer)
		if err != nil {
			fmt.Println(err)
			return err
		}
		memory.keyvalue = make(map[string][]byte)
	}
	if isSegFileMultiple(len(segContainer.memo)) {
		newSegments := compressSegments(segContainer.memo)
		segContainer.memo = newSegments
	}
	return nil
}

func Delete(k string) error {
	return Set(k, envVar.tombstone)
}

func GetLength() int {
	return len(memory.keyvalue)
}

func GetAllInMemory() map[string][]byte {
	return memory.keyvalue
}
