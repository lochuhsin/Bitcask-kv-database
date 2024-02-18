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
	 */
	setting.SetupConfig(envPaths...)
	setting.SetUpDirectory()
	memory.InitMemoryManager()
	segment.InitSegmentManager()
	scheduler.InitScheduler()
}
