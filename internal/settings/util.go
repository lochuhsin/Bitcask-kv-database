package settings

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func InitENV() {
	ENV = envVar{
		LogPath:           "./log/",
		SegmentFolder:     "seg/",
		Tombstone:         "!@#$%^&*()_+",
		NilData:           ")(*&^)!@!@#$%^&*()",
		MemoryCountLimit:  100000,
		SegLineLimit:      100000 * 10,
		SegFileCountLimit: 100,
		MemoryLogFolder:   "mlog/",
		MemoryLogFile:     "m.log",
		SegmentLogFolder:  "slog/",
		SegmentLogFile:    "s.log",
		SegmentFileExt:    ".sst",
	}
	err := godotenv.Load(ENVPATH)

	if err != nil {
		fmt.Println("env file doesn't exist")
		fmt.Println("using default setting")
		fmt.Println(ENV)
	} else {

		// System settings

		if logPath := os.Getenv("LOG_FOLDER_PATH"); logPath != "" {
			ENV.LogPath = logPath
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
		if segLineLimit := os.Getenv("SEG_LINE_LIMIT"); segLineLimit != "" {
			limit, err := strconv.Atoi(segLineLimit)
			if err != nil {
				panic("something went wrong with getting FILE_BYTE_LIMIT")
			}
			ENV.SegLineLimit = limit
		}
		if segFileCountLimit := os.Getenv("SEGMENT_FILE_COUNT_LIMIT"); segFileCountLimit != "" {
			limit, err := strconv.Atoi(segFileCountLimit)
			if err != nil {
				panic("something went wrong with getting SEGMENT_FILE_COUNT_LIMIT")
			}
			ENV.SegFileCountLimit = limit
		}
	}
}
