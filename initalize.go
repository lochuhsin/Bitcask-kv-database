package rebitcask

import (
	"fmt"
	"os"
	"rebitcask/internal/memory"
	"rebitcask/internal/scheduler"
	"rebitcask/internal/segment"
	"rebitcask/internal/settings"
)

func Init() {
	/**
	 * Should call this, whenever the server is up
	 */

	settings.InitENV()
	env := settings.ENV
	segDir := fmt.Sprintf("%s%s", env.DataPath, settings.SEGMENT_FILE_FOLDER)
	indexDir := fmt.Sprintf("%s%s", env.DataPath, settings.INDEX_FILE_FOLDER)
	os.MkdirAll(segDir, os.ModePerm)
	os.MkdirAll(indexDir, os.ModePerm)

	memory.InitMemoryManager(memory.ModelType(settings.ENV.MemoryModel))
	segment.InitSegmentManager()
	scheduler.InitScheduler()
}
