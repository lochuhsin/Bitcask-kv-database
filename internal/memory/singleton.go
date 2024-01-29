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
			settings.ENV.MemoryCountLimit,
			settings.WORKER_COUNT,
			ModelType(settings.ENV.MemoryModel),
		)
	})
}

func GetMemoryManager() *memoryManager {
	return mManager
}
