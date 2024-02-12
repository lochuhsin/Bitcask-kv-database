package rebitcask

import (
	"rebitcask/internal/memory"
	"rebitcask/internal/scheduler"
	"rebitcask/internal/segment"
	"rebitcask/internal/settings"
)

func Setup(envPaths ...string) {
	/**
	 * Should call this, whenever the server is up
	 */
	settings.SetupConfig(envPaths...)
	settings.SetUpDirectory()
	memory.InitMemoryManager()
	segment.InitSegmentManager()
	scheduler.InitScheduler()
}
