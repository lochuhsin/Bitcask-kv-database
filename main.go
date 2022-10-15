package rebitcask

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var m memoryMap
var d diskMap
var LOGFOLDER = "./log/"
var NEXTLOGNo = 0

func init() {
	// create log folder
	os.MkdirAll(LOGFOLDER, 0700)
	m.keyvalue = make(map[string]string)
	m.memoLimit = 2
	d.bytePositionMap = make(map[string]int)
	d.byteLengthMap = make(map[string]int)
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

	// byteV := []byte(v)

	m.keyvalue[k] = v
	if isExceedMemoLimit() {
		fmt.Println("Saving to dict")
		err := toDisk()
		fmt.Println(err)
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
	for k, v := range m.keyvalue {
		line := fmt.Sprintf("%v,%v\n", k, v)
		_, err := file.WriteString(line)
		if err != nil {
			return err
		}
	}

	m.keyvalue = make(map[string]string)
	return nil
}

func isKeyInDisk(k string) (v string, status bool) {
	filepath := fmt.Sprintf("%v/%v.log", LOGFOLDER, NEXTLOGNo)
	file, err := os.Open(filepath)
	if err != nil {
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res := strings.Split(scanner.Text(), ",")
		fmt.Println(res[0], k)
		if res[0] == k {
			return res[1], true
		}
	}
	return "", false
}
