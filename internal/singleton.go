package internal

import (
	"rebitcask/internal/memory"
	"rebitcask/internal/scheduler"
	"rebitcask/internal/segment"
	"rebitcask/internal/setting"
	"sync"
)

var (
	sManager *segment.Manager
	mManager *memory.Manager
	sched    *scheduler.Scheduler

	sManagerOnce  sync.Once = sync.Once{}
	mManagerOnce  sync.Once = sync.Once{}
	schedulerOnce sync.Once = sync.Once{}
)

func InitMemoryManager() {
	mManagerOnce.Do(func() {
		mStorage := memory.NewBlockStorage()
		mManager = memory.NewMemoryManager(
			mStorage,
			setting.Config.MEMORY_COUNT_LIMIT,
			setting.MEMORY_BLOCK_BUFFER_COUNT,
			memory.ModelType(setting.Config.MEMORY_MODEL),
		)
		go mManager.WriteOpListener()
	})
}

func GetMemoryManager() *memory.Manager {
	return mManager
}

func InitSegmentManager() {
	sManagerOnce.Do(func() {
		if sManager == nil {
			manager := segment.NewSegmentManager()
			sManager = &manager
		}
	})
}

func GetSegmentManager() *segment.Manager {
	return sManager
}

func InitScheduler() {
	schedulerOnce.Do(func() {
		if sched == nil {
			sched = scheduler.NewScheduler(
				GetMemoryManager(),
				GetSegmentManager(),
			)
			go sched.MemoryJobPool()
		}
	})
}
func GetScheduler() *scheduler.Scheduler {
	return sched
}
