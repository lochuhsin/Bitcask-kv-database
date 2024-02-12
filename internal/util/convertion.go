package util

import "unsafe"

func StringToBytes(s string) []byte {
	if s == "" {
		panic("empty string is not allowed")
	}
	ptr := unsafe.StringData(s)
	b := unsafe.Slice(ptr, len(s))
	return b
}

func BytesToString(b []byte) string {
	if len(b) == 0 {
		panic("empty bytes is not allowed")
	}
	p := unsafe.SliceData(b)
	return unsafe.String(p, len(b))
}
