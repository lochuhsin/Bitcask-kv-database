package memory

import (
	"rebitcask/internal/setting"
	"sync"
)

var mManager *Manager
var mOnce sync.Once

func InitMemoryManager() {
	mOnce.Do(func() {
		mStorage := NewBlockStorage()
		mManager = NewMemoryManager(
			mStorage,
			setting.Config.MEMORY_COUNT_LIMIT,
			setting.MEMORY_BLOCK_BUFFER_COUNT,
			ModelType(setting.Config.MEMORY_MODEL),
		)
		go mManager.WriteOpListener()
	})
}

func GetMemoryManager() *Manager {
	return mManager
}
