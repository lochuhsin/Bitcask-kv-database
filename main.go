package rebitcask

import (
	"fmt"
	"os"
)

var m memoryMap
var d diskMap
var LOGFOLDER = "./log/"
var NEXTLOGNo = 0

func init() {
	// create log folder
	_ = os.RemoveAll(LOGFOLDER)
	_ = os.MkdirAll(LOGFOLDER, 0700)
	initMaps()
}

func initMaps() {
	m.keyvalue = make(map[string]string)
	m.memoLimit = 2
	d.bytePositionMap = make(map[string]int)
	d.byteLengthMap = make(map[string]int)
	d.byteFileLength = 0
}

func Get(k string) (v string, status bool) {

	if val, ok := m.keyvalue[k]; ok {
		return val, true
	}
	if val, ok := isKeyInDisk(k, &d); ok {
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
	memoryMap := m.keyvalue
	return memoryMap
}
