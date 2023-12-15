package service

import (
	"fmt"
	"rebitcask/internal/storage/dao"
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
