package service

import (
	"rebitcask/internal/dao"
	"rebitcask/internal/memory"
)

func MGet(k dao.NilString) (val dao.Base, status bool) {
	val, status = memory.GetMemoryStorage().Get(k)
	if status {
		return val, status
	}
	return nil, false
}

func MSet(k dao.NilString, v dao.Base) {
	pair := dao.InitPair(k, v)
	mStorage := memory.GetMemoryStorage()
	mStorage.Set(pair)
}

func MDelete(k dao.NilString) {
	pair := dao.InitTombPair(k)
	memory.GetMemoryStorage().Set(pair)
}
