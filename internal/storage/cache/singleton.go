package cache

import (
	"fmt"
	"sync"
)

var Cache ICache

// Guards the memory model initalization
var clock = &sync.Mutex{}

func CacheInit(cType CacheType) {
	/**
	 * Initializes caches overhere
	 */
	if Cache == nil {
		clock.Lock()
		defer clock.Unlock()
		if Cache == nil {
			Cache = CacheSelector(cType)
			// TODO: implemented reload from log data
		} else {
			fmt.Println("cache exists")
		}
	} else {
		fmt.Println("cache exists")
	}
}
