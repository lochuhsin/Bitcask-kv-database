package memory

import (
	"sync"
)

var MemModel IMemory
var mOnce sync.Once

func MemoryInit(mType ModelType) {
	mOnce.Do(func() {
		if MemModel == nil {
			MemModel = MemoryTypeSelector(mType)
		}
	})
}
