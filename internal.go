package rebitcask

import (
	"fmt"
	"io"
	"os"
)

// TODO: Convert this to env file
var m memoryMap
var d diskMap
var LOGFOLDER = "./log/"
var NEXTLOGNo = 0

func init() {
	// create log storage folder

	// TODO: Convert this to env file
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

func toDisk(m *memoryMap, d *diskMap) error {
	filepath := fmt.Sprintf("%v/%v.log", LOGFOLDER, NEXTLOGNo)
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
	}

	d.byteFileLength = byteHeadPosition
	m.keyvalue = make(map[string]string)
	return nil
}

func splitSegment() {

}

func isKeyInDisk(k string, d *diskMap) (v string, status bool) {

	if _, ok := d.bytePositionMap[k]; !ok {
		return "", false
	}

	filepath := fmt.Sprintf("%v/%v.log", LOGFOLDER, NEXTLOGNo)
	bytePos, _ := d.bytePositionMap[k]
	byteLen, _ := d.byteLengthMap[k]

	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		return "Something went wrong while reading file", false
	}

	_, err = file.Seek(int64(bytePos), io.SeekStart)
	if err != nil {
		panic("Something went wrong while reading file")
	}

	readByte := make([]byte, byteLen)

	_, err = file.Read(readByte)
	if err != nil {
		panic("Something went wrong while reading file")
	}

	return string(readByte), true
}

func isExceedMemoLimit(m *memoryMap) bool {
	return len(m.keyvalue) >= m.memoLimit
}
