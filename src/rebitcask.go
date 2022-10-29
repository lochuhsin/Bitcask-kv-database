package src

import (
	"rebitcask/internal"
)

// Currently, the main point of env setting,
// memory should be larger than 50000, since
// to make sure the system (os.system) has enough time
// to write in th actually disk

func Get(k string) (v string, status bool) {
	return internal.Get(k)
}

func Set(k string, v string) error {
	return internal.Set(k, v)
}

func Delete(k string) error {
	return internal.Delete(k)
}
