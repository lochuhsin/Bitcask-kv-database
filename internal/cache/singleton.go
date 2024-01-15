package cache

import (
	"sync"
)

var (
	Cache ICache
	cOnce sync.Once
)

func CacheInit(cType CacheType) {
	/**
	 * Initializes caches overhere
	 */
	cOnce.Do(func() {
		if Cache == nil {
			Cache = CacheSelector(cType)
		}
	})
}
