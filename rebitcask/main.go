package rebitcask

import "rebitcask/internal"

func Get(k string) (v string, status bool) {
	return internal.Get(k)
}

func Set(k string, v string) error {
	return internal.Set(k, v)
}
