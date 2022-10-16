package rebitcask

import (
	"fmt"
	"os"
	"io"
)

var m memoryMap
var d diskMap
var LOGFOLDER = "./log/"
var NEXTLOGNo = 0

func init() {
	// create log folder
	os.RemoveAll(LOGFOLDER)
	os.MkdirAll(LOGFOLDER, 0700)
	initMaps()
}

func initMaps(){
	m.keyvalue = make(map[string]string)
	m.memoLimit = 10000
	d.bytePositionMap = make(map[string]int)
	d.byteLengthMap = make(map[string]int)
	d.byteFileLength = 0
}

func Get(k string) (v string, status bool) {

	if val, ok := m.keyvalue[k]; ok {
		return val, true
	}
	if val, ok := isKeyInDisk(k); ok {
		return val, true
	}

	return "", false
}
func Set(k string, v string) error {
	m.keyvalue[k] = v
	if isExceedMemoLimit() {
		err := toDisk()
		if err != nil{
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

func isExceedMemoLimit() bool {
	return len(m.keyvalue) >= m.memoLimit
}

func toDisk() error {
	filepath := fmt.Sprintf("%v/%v.log", LOGFOLDER, NEXTLOGNo)
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}

	defer file.Close()

	byteHeadPosition := d.byteFileLength
	for k, v := range m.keyvalue {

		byteValue := []byte(v)
		file.Write(byteValue)

		d.byteLengthMap[k] = len(byteValue)
		d.bytePositionMap[k] = byteHeadPosition

		byteHeadPosition += len(byteValue)
	}

	d.byteFileLength = byteHeadPosition
	m.keyvalue = make(map[string]string)
	return nil
}

func isKeyInDisk(k string) (v string, status bool) {

	if _, ok := d.bytePositionMap[k]; !ok{
		return "", false
	}

	filepath := fmt.Sprintf("%v/%v.log", LOGFOLDER, NEXTLOGNo)
	bytePos, _ := d.bytePositionMap[k]
	byteLen, _ := d.byteLengthMap[k]

	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil{
		return "Something went wrong while reading file", false
	}

	file.Seek(int64(bytePos), io.SeekStart)
	readByte := make([]byte, byteLen)

	file.Read(readByte)
	return string(readByte), true
}
