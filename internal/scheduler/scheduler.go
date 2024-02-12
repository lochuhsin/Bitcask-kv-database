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
	BlockIdChan := memory.GetMemoryManager().GetBlockIdQueue()
	for blockId := range BlockIdChan {
		go s.convertMemoryWorker(blockId)
	}
}

// Long running listener for finished task signals
func (s *Scheduler) TaskSignalListner() {
	/**
	 * When the channel receives a task finished signal,
	 * Remove the task from task pool
	 */
	mStorage := memory.GetMemoryManager()
	for ts := range s.statusChan {
		if ts.status != FINISHED {
			panic("Some thing went wrong")
		}
		mStorage.RemoveMemoryBlock(ts.id)
	}
}

// worker
func (s *Scheduler) convertMemoryWorker(id memory.BlockId) {
	manager := segment.GetSegmentManager()
	mStorage := memory.GetMemoryManager()
	block := mStorage.GetMemoryBlock(id)
	seg := memBlockToFile(*block)
	genSegmentMetadataFile(seg.Id, seg.Level)
	genSegmentIndexFile(seg.Id, seg.GetPrimaryIndex())

	manager.Add(seg)
	s.statusChan <- status{
		id:     id,
		status: FINISHED,
	}
}

func (s *Scheduler) compressSegmentWorker() {
	panic("Not implemented error")
}
