package task

import (
	"rebitcask/internal/settings"
	"sync"
)

var (
	TaskPool *taskPool
	TaskChan chan TaskId
	tOnce    sync.Once
)

func InitTaskRelated() {
	tOnce.Do(func() {
		if TaskPool == nil && TaskChan == nil {
			TaskPool = NewTaskPool()
			TaskChan = make(chan TaskId, settings.TASK_POOL_SIZE)
		}
	})
}

func GetTaskChan() chan TaskId {
	return TaskChan
}

func GetTaskPool() *taskPool {
	return TaskPool
}
