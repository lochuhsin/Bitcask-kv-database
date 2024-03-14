package memory

import (
	"rebitcask/internal/memory/models"
	"rebitcask/internal/setting"
	"sync"
)

var mManager *MemoryManager
var mOnce sync.Once

func InitMemoryManager() {
	mOnce.Do(func() {
		mStorage := NewBlockStorage()
		mManager = NewMemoryManager(
			mStorage,
			setting.Config.MEMORY_COUNT_LIMIT,
			setting.MEMORY_BLOCK_BUFFER_COUNT,
			models.ModelType(setting.Config.MEMORY_MODEL),
		)
		go mManager.WriteOpListener()
	})
}

func GetMemoryManager() *MemoryManager {
	return mManager
}
