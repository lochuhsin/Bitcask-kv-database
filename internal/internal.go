package internal

import (
	"errors"
	"fmt"
	"os"
	"sync"
)

// TODO Convert this to singleton
var memory memoryMap
var currentSeg SegmentMap
var segContainer SegmentContainer
var ENVVAR envVariables

const ENVPATH = "./rebitcask.env"

var mu = sync.Mutex{}

func init() {
	initGlobalEnvVar(ENVPATH)
	_ = os.RemoveAll(ENVVAR.logFolder)
	_ = os.MkdirAll(fmt.Sprintf("%v%v", ENVVAR.logFolder, ENVVAR.segmentFolder), 0700)
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

// Set TODO: Optimize this lock mechanism, this dramatically lower down the write performance
func Set(k string, v string) error {
	mu.Lock()
	defer mu.Unlock()
	if k == ENVVAR.tombstone {
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

// Delete : This doesn't need lock, since Set function already contains lock
func Delete(k string) error {
	err := Set(k, ENVVAR.tombstone)
	return err
}

func GetLength() int {
	return len(memory.keyvalue)
}

func GetAllInMemory() map[string][]byte {
	return memory.keyvalue
}
