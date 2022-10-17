package rebitcask

import (
	"fmt"
)

func Get(k string) (v string, status bool) {

	if val, ok := m.keyvalue[k]; ok {
		return val, true
	}
	if val, ok := isKeyInSegment(k, &d); ok {
		return val, true
	}

	if val, ok := isKeyInSegments(k, &s); ok {
		return val, true
	}
	return "", false
}
func Set(k string, v string) error {
	m.keyvalue[k] = v
	if isExceedMemoLimit(&m) {
		err := toDisk(&m, &d)
		if err != nil {
			fmt.Println(err)
		}
	}
	return nil
}

func GetLength() int {
	return len(m.keyvalue)
}

func GetAllInMemory() map[string]string {
	return m.keyvalue
}
