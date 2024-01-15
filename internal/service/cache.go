package service

import "rebitcask/internal/cache"

func CGet(k string) bool {
	return cache.Cache.Get(k)
}

func CSet(k string) {
	cache.Cache.Set(k)
}

func CDelete(k string) bool {
	return cache.Cache.Delete(k)
}
