package service

import (
	"rebitcask/internal/settings"
	"rebitcask/internal/storage/dao"
	"rebitcask/internal/storage/memory"
	"sync"
)

var memModel memory.MemoryBase

// Guards the memory model initalization
var mlock = &sync.Mutex{}

func MemoryInit(mType memory.ModelType) {
	/**
	 * Using env variable to initialize memory base model type
	 */
	if memModel == nil {
		mlock.Lock()
		defer mlock.Unlock()
		if memModel == nil {
			memModel = memoryTypeSelector(mType)
			// Implement reload from log file
		}
	}
}

func memoryTypeSelector(mType memory.ModelType) memory.MemoryBase {
	var m memory.MemoryBase = nil
	switch mType {
	case memory.HASH:
		m = memory.InitHash()
	case memory.BST:
		m = memory.InitBinarySearchTree()

	// TODO: implement these
	// case memory.AVLT:
	// 	m = memory.InitAvlTree()
	// case memory.RBT:
	// 	m = memory.InitRedBlackTree()

	default:
		panic("memory model not implemented errir")
	}
	return m
}

func MGet(k dao.NilString) (dao.Base, bool) {
	return memModel.Get(k)
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
	memModel.Set(pair)

	// TODO: write to memory log
	if memModel.GetSize() > settings.ENV.MemoryCountLimit {
		memoryToSegment(memModel)
	}

}

func MDelete(k dao.NilString) {
	pair := dao.InitTombPair(k)

	err := mLog(pair)
	if err != nil {
		panic(err)
	}

	memModel.Set(pair)

	if memModel.GetSize() > settings.ENV.MemoryCountLimit {
		memoryToSegment(memModel)
	}
}

func mLog(pair dao.Pair) error {
	return nil
}
