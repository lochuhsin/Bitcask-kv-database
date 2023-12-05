package internal

import (
	"errors"
	"fmt"
	"os"
	"rebitcask/internal/models"
	"sync"
)

// TODO Convert this to singleton
var countingbf CountingBloomFilter
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
	countingbf.Init()
}

func Get(k string) (v string, status bool) {
	if !countingbf.Get(k) {
		return "", false
	}
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
	// TODO: setting BloomFilter should be more carefully
	// Even though this might work as usual, but when large amount of duplicate keys
	// coming in, the array count will keep adding. This will increase the differences
	// between bloom filter and actual database value.
	// Resulting the fact that bloom filter becomes unreliable.
	//
	// However, for the second thought, bitcask is known for appending new key / values
	// therefor it is not reasonable for bitcask to check if the key exists before
	// setting value. Therefore, the way of adding keys to bloom filter should be
	// more carefully
	countingbf.Set(k)

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
	if countingbf.Get(k) {
		countingbf.Delete(k)
	}
	err := Set(k, ENVVAR.tombstone)
	return err
}

func Exist() (bool, error) {
	panic("Not implemented error")
}

func BulkCreate(k string) error {
	panic("Not implemented error")
}

func BulkUpdate(k string) error {
	panic("Not implemented error")
}

func BulkUpsert(k string) error {
	panic("Not implemented error")
}

func BulkDelete(k string) error {
	panic("Not implemented error")
}

func BulkGet(k ...string) ([]string, error) {
	panic("Not implemented error")
}
