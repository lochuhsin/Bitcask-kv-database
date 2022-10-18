package rebitcask

import "rebitcask/internal"

func Get(k string) (v string, status bool) {
	return internal.Get(k)
}

func Set(k string, v string) error {
	return internal.Set(k, v)
}

func Delete(k string) error {
	return internal.Delete(k)
}

func GetLength() int {
	return internal.GetLength()
}

func GetAllInMemory() map[string]string {
	return internal.GetAllInMemory()
}
