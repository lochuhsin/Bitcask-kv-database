package scheduler

import (
	"rebitcask/internal/settings"
	"sync"
)

/**
 * TODO: Fuck !!, I need to manage this simultaneously .... still thinking the better way
 */
var (
	TaskPool *taskPool
	tpOnce   sync.Once

	TaskChan chan task
	tcOnce   sync.Once

	Sched   *Scheduler
	schOnce sync.Once
)

func TaskPoolInit() {
	tpOnce.Do(func() {
		if TaskPool == nil {
			TaskPool = &taskPool{
				queue: []task{},
				mu:    sync.Mutex{},
			}
		}
	})
}

func TaskChannelInit() {
	tcOnce.Do(func() {
		if TaskChan == nil {
			TaskChan = make(chan task, settings.TASK_POOL_SIZE)
		}
	})
}

func SchedulerInit() {
	schOnce.Do(func() {
		if Sched == nil {
			Sched = NewScheduler()
		}
	})
}
