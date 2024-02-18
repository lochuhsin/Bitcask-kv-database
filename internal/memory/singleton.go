package memory

import (
	"rebitcask/internal/memory/models"
	"rebitcask/internal/setting"
	"sync"
)

var mManager *memoryManager
var mOnce sync.Once

func InitMemoryManager() {
	mOnce.Do(func() {
		mStorage := NewMemoryStorage()
		mManager = NewMemoryManager(
			mStorage,
			setting.Config.MEMORY_COUNT_LIMIT,
			setting.MEMORY_BLOCK_BUFFER_COUNT,
			models.ModelType(setting.Config.MEMORY_MODEL),
		)
	})
}

func GetMemoryManager() *memoryManager {
	return mManager
}
