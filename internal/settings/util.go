package settings

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func InitENV() {
	ENV = envVar{
		DataPath:          "./rbData/",
		SegmentFolder:     "seg/",
		Tombstone:         "!@#$%^&*()_+",
		NilData:           ")(*&^)!@!@#$%^&*()",
		MemoryModel:       "hash",
		MemoryCountLimit:  100000,
		SegFileCountLimit: 100,
	}
	err := godotenv.Load(ENVPATH)

	if err != nil {
		fmt.Println("env file doesn't exist")
		fmt.Println("using default setting")
		fmt.Println(ENV)
	} else {

		// System settings
		if DataPath := os.Getenv("DATA_FOLDER_PATH"); DataPath != "" {
			ENV.DataPath = DataPath
		}
		if tombstone := os.Getenv("TOMBSTONE"); tombstone != "" {
			ENV.Tombstone = tombstone
		}
		if memoryCountLimit := os.Getenv("MEMORY_COUNT_LIMIT"); memoryCountLimit != "" {
			limit, err := strconv.Atoi(memoryCountLimit)
			if err != nil {
				panic("something went wrong with getting MEMORY_LIMIT")
			}
			ENV.MemoryCountLimit = limit
		}
		if segFileCountLimit := os.Getenv("SEGMENT_FILE_COUNT_LIMIT"); segFileCountLimit != "" {
			limit, err := strconv.Atoi(segFileCountLimit)
			if err != nil {
				panic("something went wrong with getting SEGMENT_FILE_COUNT_LIMIT")
			}
			ENV.SegFileCountLimit = limit
		}
		if memModel := os.Getenv("MEMORY_MODEL"); memModel != "" {
			ENV.MemoryModel = memModel
		}
	}
}
