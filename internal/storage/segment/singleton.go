package segment

import (
	"fmt"
	"sync"
)

var SegManager *SegmentManager

// Guards the initialization of segment manager
var slock = &sync.Mutex{}

func SegmentInit() {
	/**
	 * Implement segment specific data structure
	 */
	if SegManager == nil {
		slock.Lock()
		defer slock.Unlock()
		if SegManager == nil {
			SegManager = InitSegmentManager()
			fmt.Println("seg manager initialized")
			/**
			 * TODO: implement reload from segment log files
			 * 1. Segment Collection
			 * 2. Segment Index Collection from log files
			 * */
		} else {
			fmt.Println("seg manager exists")
		}
	} else {
		fmt.Println("seg manager exists")
	}
}
