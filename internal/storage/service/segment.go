package service

import (
	"fmt"
	"rebitcask/internal/settings"
	"rebitcask/internal/storage/dao"
	"rebitcask/internal/storage/memory"
	"rebitcask/internal/storage/segment"
	"sync"
)

var segManager *segment.SegmentManager

// Guards the initialization of segment manager
var slock = &sync.Mutex{}

func SegmentInit() {
	/**
	 * Implement segment specific data structure
	 */
	if segManager == nil {
		slock.Lock()
		defer slock.Unlock()
		if segManager == nil {
			segManager = segment.InitSegmentManager()
			/**
			 *  TODO: implement reload from segment log files
			 * 1. Segment Collection
			 * 2. Segment Index
			 * 3. Implement transaction log files
			 * */
		} else {
			fmt.Println("seg manager exists")
		}
	} else {
		fmt.Println("seg manager exists")
	}
}

func SGet(k dao.NilString) (val dao.Base, status bool) {
	/**
	 * Get function always return two values
	 * 1. data
	 * 2. status which indicates whether the key exists or not
	 */
	return segManager.Get(k)
}

func memoryToSegment(m memory.IMemory) (bool, error) {
	/**
	 * TODO:
	 * 1. We need a lock to make sure that when the memModel is under
	 * the convertion to segment, the memory model is not allowed to
	 * perform write operations, which means frozen.
	 *
	 * In this scenario, we need a mechanism, that is able to create
	 * a new memory model to store the new write operation
	 *
	 * 2. We need another lock to ensure that there are always only one
	 * segment per memory should be created.
	 * To prevent the scenario when two concurrent writes operation reaches
	 * the condition of converting memory model to segment.
	 */
	if memory.MemModel.GetSize() > settings.ENV.MemoryCountLimit {
		segManager.ConvertToSegment(m)
	}
	return true, nil
}
