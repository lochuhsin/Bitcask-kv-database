package memory

import (
	"rebitcask/internal/settings"
	"sync"
)

var mManager *memoryManager
var mOnce sync.Once

func InitMemoryManager(mType ModelType) {
	mOnce.Do(func() {
		mStorage := NewMemoryStorage()
		mManager = NewMemoryManager(
			mStorage,
			settings.Config.MEMORY_COUNT_LIMIT,
			settings.WORKER_COUNT,
			ModelType(settings.Config.MEMORY_MODEL),
		)
	})
}

func GetMemoryManager() *memoryManager {
	return mManager
}
