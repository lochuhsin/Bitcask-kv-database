package scheduler

import (
	"rebitcask/internal/memory"
	"rebitcask/internal/segment"
	"rebitcask/internal/setting"
	"sync"
	"time"
)

type Scheduler struct {
	mManager *memory.Manager
	sManager *segment.Manager
}

func NewScheduler(mManager *memory.Manager, sManager *segment.Manager) *Scheduler {
	return &Scheduler{
		mManager: mManager,
		sManager: sManager,
	}
}

func (s *Scheduler) MemoryJobPool() {
	maxWorkerCount := setting.MEMORY_CONVERT_WORKER_COUNT
	jobQ := s.mManager.GetScheduleBlockQueue()
	wg := sync.WaitGroup{}
	for {
		// optimize this without recreating list all the time
		// event though the length is small
		batchedBlockId := make([]memory.BlockId, 0, maxWorkerCount)
		for i := 0; i < maxWorkerCount; i++ {
			select {
			case blockId := <-jobQ:
				wg.Add(1)
				go s.memoryWorker(blockId, &wg)
				batchedBlockId = append(batchedBlockId, blockId)
			case <-time.After(time.Millisecond):
				goto END_INNER_FOR
			}
		}
	END_INNER_FOR:
		wg.Wait()
		s.mManager.BulkRemoveBlockRequestQ() <- batchedBlockId
		<-s.mManager.BulkRemoveBlockResponseQ()
	}
}

// worker
func (s *Scheduler) memoryWorker(id memory.BlockId, wg *sync.WaitGroup) {
	defer wg.Done()
	block := s.mManager.GetBlock(id) //Read
	seg := memBlockToFile(block)
	createSegMetaFile(seg.Id, seg.Level)
	createSegIndexFile(seg.Id, seg.GetPrimaryIndex())
	s.sManager.Add(seg)
}

func (s *Scheduler) compressSegmentWorker() {
	panic("Not implemented error")
}
