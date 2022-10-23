package internal

import (
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"os"
	"strconv"
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
	return memoSize >= envVar.memoryKeyCountLimit
}

// TODO: find a better condition
func isSegFileMultiple(fileCount int) bool {
	return (fileCount%envVar.segFileCountLimit) == 0 && fileCount != 0
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

func initGlobalEnvVar(envPath string) {
	envVar = envVariables{
		logFolder:           "./log",
		segmentFolder:       "/seg/",
		tombstone:           "!@#$%^&*()_+",
		memoryKeyCountLimit: 20000,
		fileByteLimit:       20000,
		segFileCountLimit:   20,
	}
	err := godotenv.Load(envPath)
	if err != nil {
		fmt.Println("env file doesn't exist")
		fmt.Println("using default setting")
		fmt.Println(envVar)
	} else {

		if logFolder := os.Getenv("LOG_FOLDER_PATH"); logFolder != "" {
			envVar.logFolder = logFolder
		}
		if tombstone := os.Getenv("TOMBSTONE"); tombstone != "" {
			envVar.tombstone = tombstone
		}
		if memoryKeyCountLimit := os.Getenv("MEMORY_KEY_COUNT_LIMIT"); memoryKeyCountLimit != "" {
			limit, err := strconv.Atoi(memoryKeyCountLimit)
			if err != nil {
				panic("something went wrong with getting MEMORY_LIMIT")
			}
			envVar.memoryKeyCountLimit = limit
		}
		if fileByteLimit := os.Getenv("FILE_BYTE_LIMIT"); fileByteLimit != "" {
			limit, err := strconv.Atoi(fileByteLimit)
			if err != nil {
				panic("something went wrong with getting FILE_BYTE_LIMIT")
			}
			envVar.fileByteLimit = limit
		}
		if segFileCountLimit := os.Getenv("SEGMENT_FILE_COUNT_LIMIT"); segFileCountLimit != "" {
			limit, err := strconv.Atoi(segFileCountLimit)
			if err != nil {
				panic("something went wrong with getting SEGMENT_FILE_COUNT_LIMIT")
			}
			envVar.segFileCountLimit = limit
		}
	}

	fmt.Println("env setting done")
	fmt.Println(envVar)
}
