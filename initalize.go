package rebitcask

import (
	"rebitcask/internal/memory"
	"rebitcask/internal/scheduler"
	"rebitcask/internal/segment"
	"rebitcask/internal/setting"
)

func Setup(envPaths ...string) {
	/**
	 * Should call this, whenever the server is up
	 * Note: the order of initialization is sensitive
	 */
	setting.SetupConfig(envPaths...)
	setting.SetUpDirectory()
	memory.InitMemoryManager()
	segment.InitSegmentManager()
	scheduler.InitScheduler()
}
