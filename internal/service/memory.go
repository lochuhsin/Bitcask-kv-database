package service

import (
	"rebitcask/internal/dao"
	"rebitcask/internal/memory"
	"rebitcask/internal/settings"
	"rebitcask/internal/task"
	"sync"
	"time"

	"github.com/google/uuid"
)

var muSegConverter = &sync.Mutex{}

func MGet(k dao.NilString) (val dao.Base, status bool) {
	/**
	 * The Get function always returns value, and status
	 * status indicates whether the key exists or not
	 */
	/**
	 * 1. Get from memory model
	 * 2. Get from task pool (waiting tasks)
	 * 3. Get from schedumere
	 */

	val, status = memory.GetMemoryStorage().Get(k)
	if status {
		return val, status
	}

	val, status = getFromTaskPool(k)
	if status {
		return val, status
	}

	return nil, false
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
	mStorage := memory.GetMemoryStorage()
	// write to memory
	for mStorage.Isfrozen() {
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
	mStorage.Set(pair)
	dumpMemory(mStorage)
}

func MDelete(k dao.NilString) {
	// Optimize this
	pair := dao.InitTombPair(k)

	err := mLog(pair)
	if err != nil {
		panic(err)
	}
	mStorage := memory.GetMemoryStorage()
	for mStorage.Isfrozen() {
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
	mStorage.Set(pair)
	dumpMemory(mStorage)
}

func mLog(pair dao.Pair) error {
	// TODO: Implement this,
	return nil
}

func getFromTaskPool(k dao.NilString) (val dao.Base, status bool) {
	pool := task.GetTaskPool()
	for _, t := range pool.GetByOrder() {
		val, status := t.M.Get(k)
		if status {
			return val, status
		}
	}
	return nil, false
}

func dumpMemory(m memory.IMemory) (bool, error) {
	/**
	 * TODO: not sure how to speed this up
	 * Since double checking locking pattern is not working in
	 * Go memory model
	 */
	muSegConverter.Lock()
	if m.GetSize() > settings.ENV.MemoryCountLimit {
		m.Setfrozen(true)
		createTask(m)
		m.Reset()
	}
	muSegConverter.Unlock()
	return true, nil
}

func createTask(m memory.IMemory) {
	/**
	 * Note: Adding task to task pool and adding task id to tChan
	 * should be considered atomic.
	 *
	 * Still trying to figure out the mutex mechanism here.
	 */

	tPool := task.GetTaskPool()
	tChan := task.GetTaskChan()
	newM := m.Clone() // doing snapshot

	t := task.Task{
		Timestamp: time.Now().UnixNano(),
		Id:        task.TaskId(uuid.New().String()), // avoiding this in the future
		M:         newM,
	}
	tPool.Set(t.Id, t)
	tChan <- t.Id
}
