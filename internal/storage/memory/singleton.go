package memory

import (
	"fmt"
	"sync"
)

var MemModel IMemory

// Guards the memory model initalization
var mInitLock = &sync.Mutex{}
var mResetLock = &sync.Mutex{}

// TODO: Using sync.Once to implement this
func MemoryInit(mType ModelType) {
	/**
	 * Using env variable to initialize memory base model type
	 */
	if MemModel == nil {
		mInitLock.Lock()
		defer mInitLock.Unlock()
		if MemModel == nil {
			MemModel = MemoryTypeSelector(mType)
			fmt.Println("memory model initialized")
			// Implement reload from log file
		}
	}
}

func MemoryReset(mType ModelType) {
	if MemModel != nil {
		mResetLock.Lock()
		defer mResetLock.Unlock()
		if MemModel != nil {
			MemModel = MemoryTypeSelector(mType)
		}
	}
}
