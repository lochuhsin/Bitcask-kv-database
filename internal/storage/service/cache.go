package service

import "rebitcask/internal/storage/cache"

// TODO: Convert this to singleton
var mCache cache.Base

func CacheInit(cType cache.CacheType) {
	/**
	 * Initializes the cache overhere
	 */

	// TODO: implemented reload from log
	mCache = cacheSelector(cType)
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
