package rebitcask

import (
	"rebitcask/internal/memory"
	"rebitcask/internal/scheduler"
	"rebitcask/internal/segment"
	"rebitcask/internal/settings"
)

func Init(envPaths ...string) {
	/**
	 * Should call this, whenever the server is up
	 */
	settings.SetupConfig(envPaths...)
	settings.SetUpDirectory()
	memory.InitMemoryManager(memory.ModelType(settings.Config.MEMORY_MODEL))
	segment.InitSegmentManager()
	scheduler.InitScheduler()
}
