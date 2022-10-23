package internal

import (
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"os"
)

func createNewSegment(newSegmentNo int) (file *os.File, segmentMap SegmentMap) {
	segmentMap = SegmentMap{
		bytePositionMap:  make(map[string]int),
		byteLengthMap:    make(map[string]int),
		byteFileLength:   0,
		CurrentSegmentNo: newSegmentNo,
	}
	filepath := fmt.Sprintf("%v%v/%v.log", envVar.logFolder, envVar.segmentFolder, newSegmentNo)
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)

	if err != nil {
		panic("something wrong with opening file")
	}
	return file, segmentMap
}

func isExceedMemoLimit(memoSize int) bool {
	return memoSize >= envVar.memoryLimit
}

// TODO: find a better condition
func isSegFileMultiple(fileCount int) bool {
	return (fileCount % envVar.segFileCountLimit) == 0
}

func filterTombStone(val string) (value string, status bool) {
	if val == envVar.tombstone {
		return "", false
	}
	return val, true
}

func seekFile(file *os.File, byteHead int, byteLen int) (bytes []byte) {
	_, err := file.Seek(int64(byteHead), io.SeekStart)
	if err != nil {
		panic("Something went wrong while seeking file")
	}
	readByte := make([]byte, byteLen)
	_, err = file.Read(readByte)
	if err != nil {
		panic("Something went wrong while seeking file")
	}
	return readByte
}

func initGlobalVar(envPath string) {
	err := godotenv.Load(envPath)
	if err != nil {
		fmt.Println("env file doesn't exists")
		fmt.Println("create with default")
		envVar = envVariables{
			logFolder:         "./log/",
			segmentFolder:     "seg/",
			tombstone:         "!@#$%^&*()_+",
			memoryLimit:       20000,
			fileByteLimit:     20000,
			segFileCountLimit: 20,
		}
	} else {

	}
}
