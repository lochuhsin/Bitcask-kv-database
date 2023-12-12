package service

import (
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
		}
	}
}

func SGet(k dao.NilString) (dao.Base, bool) {
	return segManager.Get(k)
}

func memoryToSegment(m memory.MemoryBase) error {
	segManager.ConvertToSegment(m)
	return nil
}
