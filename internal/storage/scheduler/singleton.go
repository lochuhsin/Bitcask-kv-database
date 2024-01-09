package scheduler

import (
	"fmt"
	"rebitcask/internal/settings"
	"sync"
)

/**
 * TODO: Fuck !!, I need to manage this simultaneously .... still thinking the better way
 */

var TaskPool *taskPool
var tpInitLock = sync.Mutex{} // guards the initialization of the taskPool

var TaskChan chan task
var tcInitLock = sync.Mutex{} // guards the initialization of the taskChannel

var Sched *Scheduler
var schInitLock = sync.Mutex{} // guards the initialization of the scheduler

func TaskPoolInit() {
	if TaskPool == nil {
		tpInitLock.Lock()
		defer tpInitLock.Unlock()
		if TaskPool == nil {
			TaskPool = &taskPool{
				queue: []task{},
				mu:    sync.Mutex{},
			}
			fmt.Println("Task Pool Initialized")
		} else {
			fmt.Println("Task Pool exists")
		}
	} else {
		fmt.Println("Task Pool exists")
	}
}

func TaskChannelInit() {
	if TaskChan == nil {
		tcInitLock.Lock()
		defer tcInitLock.Unlock()
		if TaskChan == nil {
			TaskChan = make(chan task, settings.TASK_POOL_SIZE)
			fmt.Println("Task Chan Initialized")
		} else {
			fmt.Println("Task Chan exists")
		}
	} else {
		fmt.Println("Task Chan exists")
	}
}

func SchedulerInit() {
	if Sched == nil {
		schInitLock.Lock()
		defer schInitLock.Unlock()
		if Sched == nil {
			Sched = InitScheduler()
			fmt.Println("Scheduler Initialized")
			// Start running daemon goroutines
			go Sched.StartTaskScheduling()
			fmt.Println("Scheduling Daemon Initialized")
			go Sched.StartTaskSignalHandler()
			fmt.Println("Signaling Daemon Initialized")
		} else {
			fmt.Println("scheduler exists")
		}
	} else {
		fmt.Println("scheduler exists")
	}

}
