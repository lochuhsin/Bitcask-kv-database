package memory

import (
	"sync"
)

var mStorage *memoryStorage
var mOnce sync.Once

func InitMemoryStorage(mType ModelType) {
	mOnce.Do(func() {
		if mStorage == nil {
			mStorage = NewMemoryStorage()
		}
	})
}

func GetMemoryStorage() *memoryStorage {
	return mStorage
}
