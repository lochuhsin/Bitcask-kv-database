package scheduler

import (
	"rebitcask/internal/memory"
	"rebitcask/internal/segment"
	"rebitcask/internal/setting"
	"sync"
	"time"
)

type Scheduler struct {
	mManager *memory.MemoryManager
	sManager *segment.Manager
}

func NewScheduler(mManager *memory.MemoryManager, sManager *segment.Manager) *Scheduler {
	return &Scheduler{
		mManager: mManager,
		sManager: sManager,
	}
}

func (s *Scheduler) MemoryJobPool() {
	maxWorkerCount := setting.MEMORY_CONVERT_WORKER_COUNT
	jobQ := s.mManager.GetBlockIdQueue()
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
		s.mManager.BulkRemoveMemoryBlock(batchedBlockId)
	}
}

// worker
func (s *Scheduler) memoryWorker(id memory.BlockId, wg *sync.WaitGroup) {
	defer wg.Done()
	block := s.mManager.GetMemoryBlock(id) //Read
	seg := memBlockToFile(*block)
	genSegmentMetadataFile(seg.Id, seg.Level)
	genSegmentIndexFile(seg.Id, seg.GetPrimaryIndex())
	s.sManager.Add(seg)
}

func (s *Scheduler) compressSegmentWorker() {
	panic("Not implemented error")
}
