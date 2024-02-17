package memory

import (
	"fmt"
	"rebitcask/internal/dao"
	"rebitcask/internal/memory/models"
	"sync"
)

type memoryManager struct {
	bStorage        *blockStorage
	blockIdCh       chan BlockId
	entryCountLimit int
	mu              sync.Mutex
	modelType       models.ModelType
}

func NewMemoryManager(bStorage *blockStorage, entryCountLimit, blockIdChanSize int, modelType models.ModelType) *memoryManager {
	bStorage.createNewBlock(modelType)
	return &memoryManager{
		bStorage:        bStorage,
		blockIdCh:       make(chan BlockId, blockIdChanSize),
		mu:              sync.Mutex{},
		entryCountLimit: entryCountLimit,
		modelType:       modelType,
	}
}

func (m *memoryManager) Get(key []byte) (dao.Entry, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	blocks := m.bStorage.iterateExistingBlocks()
	for _, block := range blocks {
		val, status := block.Memory.Get(key)
		if status {
			return val, status
		}
	}
	return dao.Entry{}, false
}

/**
 * User ->
 */

//  func Set(entry) {
// 	resp := chan struct{}{}
//      tasks <- (entry, resp)
//     <-resp
//  }
//  //RLock

//  // scheduler main write loop
//  go func () {
// 	for task <-tasks {//Writes
// 		//Lock()
//        ...
// 		resp <- struct{}{}
// 	}
//  }

// called by user (endpoints)

func (m *memoryManager) Set(entry dao.Entry) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// 1. set entry to memory
	memory := m.bStorage.getCurrentBlock().Memory
	memory.Set(entry)
	// 2. check memory block condition if meets
	if memory.GetSize() >= m.entryCountLimit {
		bid := m.bStorage.getCurrentBlockId()
		// 3. add new memory block
		m.bStorage.createNewBlock(m.modelType)
		// 4. add old block id to block queue for running tests
		m.blockIdCh <- bid
	}
	return nil
}

func (m *memoryManager) RemoveMemoryBlock(id BlockId) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.bStorage.removeMemoryBlock(id) != nil {
		panic("Invalid operation on remove memory block")
	}
}

func (m *memoryManager) BulkRemoveMemoryBlock(ids []BlockId) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, id := range ids {
		if m.bStorage.removeMemoryBlock(id) != nil {
			fmt.Println("Invalid operation on remove memory block", id)
			panic("error removing memory block")
		}
	}
}

func (m *memoryManager) GetMemoryBlock(id BlockId) *Block {
	m.mu.Lock()
	defer m.mu.Unlock()
	// NOTE: only allow read access, no write access
	// Implement block froze feature to avoid modification
	// of block information
	block, status := m.bStorage.getMemoryBlock(id)
	if !status {
		panic("Invalid operation while getting memory block: " + id)
	}
	return &block
}

func (m *memoryManager) GetBlockIdQueue() <-chan BlockId {
	return m.blockIdCh
}
