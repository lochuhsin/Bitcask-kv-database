package scheduler

import (
	"rebitcask/internal/memory"
	"rebitcask/internal/segment"
	"sync"
)

var (
	scheduler *Scheduler
	sOnce     sync.Once
)

func InitScheduler() {
	sOnce.Do(
		func() {
			if scheduler == nil {
				scheduler = NewScheduler(
					memory.GetMemoryManager(),
					segment.GetSegmentManager(),
				)
				go scheduler.MemoryJobPool()
			}
		},
	)
}

func GetScheduler() *Scheduler {
	return scheduler
}
