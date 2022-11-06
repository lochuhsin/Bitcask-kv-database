package internal

import (
	"errors"
	"fmt"
	"os"
	"rebitcask/internal/models"
	"sync"
)

// TODO Convert this to singleton
var memory models.AVLTree
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
	memory.Init()
	segContainer.Init()
}

func Get(k string) (v string, status bool) {

	if item, ok := memory.Get(k); ok {
		str := string(item.Val)
		return filterTombStone(str)
	}

	// check in current segment
	if val, ok := isKeyInSegments(&k, &segContainer); ok {
		str := string(val)
		return filterTombStone(str)
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

	b := models.Item{
		Val: []byte(v),
	}
	memory.Set(k, b)
	if isExceedMemoLimit(memory.GetSize()) {
		err := toSegment(&memory, &segContainer)
		if err != nil {
			fmt.Println(err)
			return err
		}
		memory.Init()
	}
	if isSegFileMultiple(segContainer.memo.GetSize()) {
		segContainer = compressSegments(&segContainer)
	}
	return nil
}

// Delete : This doesn't need lock, since Set function already contains lock
func Delete(k string) error {
	err := Set(k, ENVVAR.tombstone)
	return err
}
