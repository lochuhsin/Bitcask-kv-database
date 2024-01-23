package scheduler

import (
	"rebitcask/internal/memory"
	"rebitcask/internal/segment"
)

type status struct {
	id     memory.BlockId
	status tStatus
}

type Scheduler struct {
	statusChan chan status
}

func NewScheduler() *Scheduler {
	return &Scheduler{statusChan: make(chan status, 1000)}
}

// Long running listener for tasks
func (s *Scheduler) TaskChanListener() {
	BlockIdChan := memory.GetMemoryStorage().GetBlockIdChan()
	for blockId := range BlockIdChan {
		go s.taskWorker(blockId)
	}
}

// Long running listener for finshed task signals
func (s *Scheduler) TaskSignalListner() {
	/**
	 * When the channel recieves a task finised signal,
	 * Remove the task from task pool
	 */
	mStorage := memory.GetMemoryStorage()
	for ts := range s.statusChan {
		if ts.status != FINISHED {
			panic("Some thing went wrong")
		}
		mStorage.RemoveMemoryBlock(ts.id)
	}
}

// worker
func (s *Scheduler) taskWorker(id memory.BlockId) {
	mStorage := memory.GetMemoryStorage()
	block, st := mStorage.GetMemoryBlock(id)
	if !st {
		panic("Got empty tasks, this shouldn't happen")
	}

	manager := segment.GetSegmentManager()
	manager.ConvertToSegment(block.Memory)
	s.statusChan <- status{
		id:     id,
		status: FINISHED,
	}
}
