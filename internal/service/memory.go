package service

import (
	"rebitcask/internal/dao"
	"rebitcask/internal/memory"
)

func MGet(k dao.NilString) (val dao.Base, status bool) {
	val, status = memory.GetMemoryManager().Get(k)
	if status {
		return val, status
	}
	return nil, false
}

func MSet(k dao.NilString, v dao.Base) {
	entry := dao.InitEntry(k, v)
	memory.GetMemoryManager().Set(entry)
}

func MDelete(k dao.NilString) {
	entry := dao.InitTombEntry(k)
	memory.GetMemoryManager().Set(entry)
}
