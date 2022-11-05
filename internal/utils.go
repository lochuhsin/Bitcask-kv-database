package internal

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

func isExceedMemoLimit(memoSize int) bool {
	return memoSize >= ENVVAR.memoryKeyCountLimit
}

// TODO: find a better condition
func isSegFileMultiple(fileCount int) bool {
	return (fileCount%ENVVAR.segFileCountLimit) == 0 && fileCount != 0
}

func filterTombStone(val string) (value string, status bool) {
	if val == ENVVAR.tombstone {
		return "", false
	}
	return val, true
}

func initGlobalEnvVar(envPath string) {
	ENVVAR = envVariables{
		logFolder:           "./log",
		segmentFolder:       "/seg/",
		tombstone:           "!@#$%^&*()_+",
		memoryKeyCountLimit: 40000,
		fileLineLimit:       400000,
		segFileCountLimit:   100,
	}
	err := godotenv.Load(envPath)
	if err != nil {
		fmt.Println("env file doesn't exist")
		fmt.Println("using default setting")
		fmt.Println(ENVVAR)
	} else {

		if logFolder := os.Getenv("LOG_FOLDER_PATH"); logFolder != "" {
			ENVVAR.logFolder = logFolder
		}
		if tombstone := os.Getenv("TOMBSTONE"); tombstone != "" {
			ENVVAR.tombstone = tombstone
		}
		if memoryKeyCountLimit := os.Getenv("MEMORY_KEY_COUNT_LIMIT"); memoryKeyCountLimit != "" {
			limit, err := strconv.Atoi(memoryKeyCountLimit)
			if err != nil {
				panic("something went wrong with getting MEMORY_LIMIT")
			}
			ENVVAR.memoryKeyCountLimit = limit
		}
		if fileLineLimit := os.Getenv("FILE_LINE_LIMIT"); fileLineLimit != "" {
			limit, err := strconv.Atoi(fileLineLimit)
			if err != nil {
				panic("something went wrong with getting FILE_BYTE_LIMIT")
			}
			ENVVAR.fileLineLimit = limit
		}
		if segFileCountLimit := os.Getenv("SEGMENT_FILE_COUNT_LIMIT"); segFileCountLimit != "" {
			limit, err := strconv.Atoi(segFileCountLimit)
			if err != nil {
				panic("something went wrong with getting SEGMENT_FILE_COUNT_LIMIT")
			}
			ENVVAR.segFileCountLimit = limit
		}
	}

	fmt.Println("env setting done")
	fmt.Println(ENVVAR)
}
