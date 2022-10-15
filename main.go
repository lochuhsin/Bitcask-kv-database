package rebitcask

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

type Memory struct {
	memo      map[string]string
	memoLimit int
}

var m Memory
var LOGFOLDER string = "./log/"
var NEXTLOGNo int = 0

func init() {
	// create log folder
	os.MkdirAll(LOGFOLDER, 0700)
	m.memo = make(map[string]string)
	m.memoLimit = 2
}

func Get(key string) (val string, status bool) {

	if val, ok := m.memo[key]; ok {
		return val, true
	}

	if val, ok := isKeyInDisk(key); ok{
		return val, true
	}

	return "", false
}
func Set(key string, val string) error {
	m.memo[key] = val
	if isExceedMemoLimit() {
		fmt.Println("Saving to dict")
		err := toDisk()
		fmt.Println(err)
	}
	return nil
}

func GetLength() int {
	return len(m.memo)
}

func GetAllInMemory() map[string]string {
	memoryMap := m.memo
	return memoryMap
}

func isExceedMemoLimit() bool {
	return len(m.memo) >= m.memoLimit
}

func toDisk() error {
	filepath := fmt.Sprintf("%v/%v.log", LOGFOLDER, NEXTLOGNo)
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}

	defer file.Close()
	for k, v := range m.memo {
		line := fmt.Sprintf("%v,%v\n", k, v)
		_, err := file.WriteString(line)
		if err != nil {
			return err
		}
	}

	m.memo = make(map[string]string)
	return nil
}

func isKeyInDisk(key string) (value string, status bool){
	filepath := fmt.Sprintf("%v/%v.log", LOGFOLDER, NEXTLOGNo)
	file, err := os.Open(filepath)
	if err != nil {
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res := strings.Split(scanner.Text(), ",")
		fmt.Println(res[0], key)
		if res[0] == key {
			return res[1], true
		}
	}
	return "", false
}
