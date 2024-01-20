package memory

import (
	"sync"
)

var MemModel IMemory
var mOnce sync.Once

func InitMemory(mType ModelType) {
	mOnce.Do(func() {
		if MemModel == nil {
			MemModel = MemoryTypeSelector(mType)
		}
	})
}

func GetMemoryStorage() IMemory {
	return MemModel
}
