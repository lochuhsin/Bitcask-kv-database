package service

import (
	"rebitcask/internal/settings"
	"rebitcask/internal/storage/dao"
	"rebitcask/internal/storage/memory"
	"sync"
)

var muSegConverter = &sync.Mutex{}

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

	err := mLog(pair)
	if err != nil {
		panic(err)
	}
	// write to memory
	for memory.MemModel.Isfrozen() {
		/**
		 * This for loop is a workaround when the memory is under the process of
		 * converting to segment. In this case, the memory model is frozen, in which
		 * it closes the write operation. So we had two choices:
		 *
		 * 1. Wait until the memory model, is unfrozen
		 *
		 * 2. we use a queue to hold all the frozen memory, especially when the write
		 * operation is really huge. Run a background goroutine to process these frozen memory.
		 * In the meantime, we create a new memory model for keep writing.
		 *
		 * For current implementation, we choose the first one, since it's simpler to implement.
		 * However the performance impact of write heavy cases are really huge.
		 * Therefore, we should be able to reimplement to the second method, in near future.
		 *
		 */

	}
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
	// TODO: Implement this,
	return nil
}

func memoryToSegment(m memory.IMemory) (bool, error) {
	/**
	 * TODO:
	 * 1. We need a lock to make sure that when the memModel is under
	 * the convertion to segment, the memory model is not allowed to
	 * perform write operations, which means frozen.
	 *
	 * In this scenario, we need a mechanism, that is able to create
	 * a new memory model to store the new write operation
	 *
	 */

	if memory.MemModel.GetSize() > settings.ENV.MemoryCountLimit {
		muSegConverter.Lock()
		if memory.MemModel.GetSize() > settings.ENV.MemoryCountLimit {
			memory.MemModel.Setfrozen(true)
			segManager.ConvertToSegment(m)
			memory.MemModel.Reset()
			/**
			 * TODO: Generate primary index a key offset mappings for this current segment
			 */
		}
		muSegConverter.Unlock()
	}
	return true, nil
}
