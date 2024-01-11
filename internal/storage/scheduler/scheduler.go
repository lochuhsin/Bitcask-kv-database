package scheduler

import (
	"fmt"
	"rebitcask/internal/settings"
	"rebitcask/internal/storage/segment"
)

type taskSignal struct {
	id     taskId
	status taskStatus
}

type Scheduler struct {
	/**
	 * fuck, we need a datastructure that is O(1) for create / delete and O(n) in loop
	 * objects in order, therefore i choose ordered map
	 */
	processPool TaskOrderedMap
	taskSignal  chan taskSignal
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		processPool: *InitTaskOrderedMap(),
		taskSignal:  make(chan taskSignal, settings.WORKER_COUNT),
	}
}

// Daemon goroutine
func (s *Scheduler) StartTaskScheduling() {

	for t := range TaskChan {
		tPool, err := TaskPool.Pop()
		if !err {
			panic("something went wrong, task pool should sync with task chen")
		}
		if tPool.id != t.id {
			panic(fmt.Sprintf("In consistent task id: %v, %v", tPool.id, t.id))
		}

		s.processPool.Set(t.id, t)
		go s.WriteToSegment(t)
	}
}

func (s *Scheduler) WriteToSegment(t task) {
	segment.SegManager.ConvertToSegment(t.m)
	s.taskSignal <- taskSignal{
		id:     t.id,
		status: FINISHED,
	}
}

// Daemon goroutine
func (s *Scheduler) StartTaskSignalHandler() {
	for signal := range s.taskSignal {
		if signal.status == FINISHED {
			s.processPool.Delete(signal.id)
		}
	}
}

func (s *Scheduler) GetByOrder() []task {
	return s.processPool.GetByOrder()
}
