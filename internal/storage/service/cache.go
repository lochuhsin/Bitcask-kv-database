package service

import (
	"rebitcask/internal/storage/cache"
	"sync"
)

var mCache cache.Base

// Guards the memory model initalization
var clock = &sync.Mutex{}

func CacheInit(cType cache.CacheType) {
	/**
	 * Initializes caches overhere
	 */
	if mCache == nil {
		clock.Lock()
		defer clock.Unlock()
		if mCache == nil {
			mCache = cacheSelector(cType)
			// TODO: implemented reload from log data
		}
	}
}

func cacheSelector(ctype cache.CacheType) cache.Base {
	var cBase cache.Base = nil
	switch ctype {
	case cache.CBF:
		cBase = cache.InitCBF()

	default:
		panic("Unsupported cache type")
	}
	return cBase
}

func CGet(k string) bool {
	return mCache.Get(k)
}

func CSet(k string) {
	mCache.Set(k)
}

func CDelete(k string) bool {
	return mCache.Delete(k)
}
