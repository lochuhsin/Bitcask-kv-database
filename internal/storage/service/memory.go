package service

import (
	"rebitcask/internal/storage/dao"
	"rebitcask/internal/storage/memory"
)

func MGet(k dao.NilString) (val dao.Base, status bool) {
	/**
	 * The Get function always returns value, and status
	 * status indicates whether the key exists or not
	 */
	return memory.MemModel.Get(k)
}

func MSet(k dao.NilString, v dao.Base) {
	/**
	 * Not only written to memory
	 * write to a memory log file to perform crash reload
	 */
	pair := dao.InitPair(k, v)

	// first log then memory
	err := mLog(pair)
	if err != nil {
		panic(err)
	}
	// write to memory
	memory.MemModel.Set(pair)

	memoryToSegment(memory.MemModel)
}

func MDelete(k dao.NilString) {
	pair := dao.InitTombPair(k)

	err := mLog(pair)
	if err != nil {
		panic(err)
	}

	memory.MemModel.Set(pair)
	memoryToSegment(memory.MemModel)
}

func mLog(pair dao.Pair) error {
	// Implement this,
	return nil
}
