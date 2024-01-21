package scheduler

import (
	"rebitcask/internal/segment"
	"rebitcask/internal/settings"
	"rebitcask/internal/task"
)

type status struct {
	id     task.TaskId
	status tStatus
}

type Scheduler struct {
	statusChan      chan status
	workerSemaphore chan struct{}
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		statusChan:      make(chan status, 1000),
		workerSemaphore: make(chan struct{}, settings.WORKER_COUNT)}
}

// Long running listener for tasks
func (s *Scheduler) TaskChanListener() {

	tChan := task.GetTaskChan()
	for taskId := range tChan {
		s.workerSemaphore <- struct{}{}
		go s.taskWorker(taskId)
	}
}

// Long running listener for finshed task signals
func (s *Scheduler) TaskSignalListner() {
	/**
	 * When the channel recieves a task finised signal,
	 * Remove the task from task pool
	 */
	tPool := task.GetTaskPool()
	for ts := range s.statusChan {
		if ts.status != FINISHED {
			panic("Some thing went wrong")
		}
		tPool.Delete(ts.id)
		<-s.workerSemaphore // releasing the position in semaphore
	}
}

// worker
func (s *Scheduler) taskWorker(tid task.TaskId) {
	tPool := task.GetTaskPool()
	task, st := tPool.Get(tid)
	if !st {
		panic("Got empty tasks, this shouldn't happen")
	}

	manager := segment.GetSegmentManager()
	manager.ConvertToSegment(task.M)
	s.statusChan <- status{
		id:     tid,
		status: FINISHED,
	}
}
