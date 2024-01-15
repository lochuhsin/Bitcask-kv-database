package cache

func CacheSelector(ctype CacheType) ICache {
	var cBase ICache = nil
	switch ctype {
	case CBF:
		cBase = InitCBF()

	default:
		panic("Unsupported cache type")
	}
	return cBase
}
