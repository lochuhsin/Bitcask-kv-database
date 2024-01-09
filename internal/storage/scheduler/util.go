package scheduler

import (
	"rebitcask/internal/storage/memory"
	"time"

	"github.com/google/uuid"
)

func ConvertMemoToTask(m memory.IMemory) task {
	return task{
		timestamp: time.Now().UnixNano(),
		id:        taskId(uuid.New().String()), // uuid
		m:         m,                           // should be cloned memory
	}
}

func AddTask(t task) {
	// Add pool first, preventing that other threads check err
	// referes to scheduler
	TaskPool.Add(t)
	TaskChan <- t
}
