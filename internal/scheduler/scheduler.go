package scheduler

import (
	"rebitcask/internal/memory"
	"rebitcask/internal/segment"
	"rebitcask/internal/settings"
	"sync"
	"time"
)

type Scheduler struct {
}

func NewScheduler() *Scheduler {
	return &Scheduler{}
}

// Long running listener for tasks
func (s *Scheduler) TaskChanListener() {
	manager := memory.GetMemoryManager()
	blockCh := manager.GetBlockIdQueue()
	runningWorker := 0
	workerCount := settings.MEMORY_CONVERT_WORKER_COUNT
	var wg sync.WaitGroup
	idList := []memory.BlockId{}
	for {
		select {
		case blockId, ok := <-blockCh: // memory block
			if !ok {
				goto END_FOR
			}
			wg.Add(1)
			go s.convertMemoryWorker(blockId, &wg) // worker
			idList = append(idList, blockId)       // [4, 5, 6]
			runningWorker++

		case <-time.After(time.Millisecond):
			wg.Wait()
			runningWorker = 0
			manager.BulkRemoveMemoryBlock(idList) //write
			idList = []memory.BlockId{}
		}

		if runningWorker >= workerCount {
			wg.Wait()
			runningWorker = 0
			manager.BulkRemoveMemoryBlock(idList) //  deadlock
			idList = []memory.BlockId{}
		}
	}

END_FOR:
	wg.Wait()
	manager.BulkRemoveMemoryBlock(idList)
}

// worker
func (s *Scheduler) convertMemoryWorker(id memory.BlockId, wg *sync.WaitGroup) {
	defer wg.Done()
	manager := segment.GetSegmentManager() //Read
	mStorage := memory.GetMemoryManager()  //Read
	block := mStorage.GetMemoryBlock(id)   //Read
	seg := memBlockToFile(*block)
	genSegmentMetadataFile(seg.Id, seg.Level)
	genSegmentIndexFile(seg.Id, seg.GetPrimaryIndex())

	manager.Add(seg)
}

func (s *Scheduler) compressSegmentWorker() {
	panic("Not implemented error")
}
