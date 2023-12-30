package storage

import (
	"fmt"
	"os"
	"rebitcask/internal/settings"
	"rebitcask/internal/storage/cache"
	"rebitcask/internal/storage/memory"
	"rebitcask/internal/storage/service"
)

// move these to env
const (
	cacheType cache.CacheType = cache.CBF
)

func Init() {
	/**
	 * Should call this, whenever the server is up
	 */
	settings.InitENV()
	env := settings.ENV
	cache.CacheInit(cacheType)
	memory.MemoryInit(memory.ModelType(settings.ENV.MemoryModel))
	service.SegmentInit()
	segDir := fmt.Sprintf("%s%s", env.DataPath, settings.SEGMENT_FILE_FOLDER)
	os.MkdirAll(segDir, os.ModePerm)
	indexDir := fmt.Sprintf("%s%s", env.DataPath, settings.INDEX_FILE_FOLDER)
	os.MkdirAll(indexDir, os.ModePerm)

}