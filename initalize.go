package rebitcask

import (
	"rebitcask/internal"
	"rebitcask/internal/setting"
)

func Setup(envPaths ...string) {
	/**
	 * Should call this, whenever the server is up
	 * Note: the order of initialization is sensitive
	 */
	setting.SetupConfig(envPaths...)
	setting.SetUpDirectory()
	internal.InitMemoryManager()
	internal.InitSegmentManager()
	internal.InitScheduler()
}
