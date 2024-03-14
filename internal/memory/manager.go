package memory

import (
	"fmt"
	"rebitcask/internal/dao"
	"rebitcask/internal/memory/models"
	"sync"
)

type MemoryManager struct {
	bStorage        *blockStorage
	blockIdCh       chan BlockId
	entryCountLimit int
	mu              sync.Mutex
	modelType       models.ModelType
}

func NewMemoryManager(bStorage *blockStorage, entryCountLimit, blockIdChanSize int, modelType models.ModelType) *MemoryManager {
	bStorage.createNewBlock(modelType)
	return &MemoryManager{
		bStorage:        bStorage,
		blockIdCh:       make(chan BlockId, blockIdChanSize),
		mu:              sync.Mutex{},
		entryCountLimit: entryCountLimit,
		modelType:       modelType,
	}
}

func (m *MemoryManager) Get(key []byte) (dao.Entry, bool) {
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

func (m *MemoryManager) Set(entry dao.Entry) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// set entry to memory
	memory := m.bStorage.getCurrentBlock().Memory
	memory.Set(entry)
	// check memory block condition if meets
	if memory.GetSize() >= m.entryCountLimit {
		bid := m.bStorage.getCurrentBlockId()
		// add new memory block
		m.bStorage.createNewBlock(m.modelType)
		// add old block id to block queue for running tests
		m.blockIdCh <- bid
	}
	return nil
}

func (m *MemoryManager) RemoveMemoryBlock(id BlockId) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.bStorage.removeMemoryBlock(id) != nil {
		panic("Invalid operation on remove memory block")
	}
}

func (m *MemoryManager) BulkRemoveMemoryBlock(ids []BlockId) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, id := range ids {
		err := m.bStorage.removeMemoryBlock(id)
		if err != nil {
			fmt.Println("Invalid operation on remove memory block", id)
			panic("err")
		}
	}
}

func (m *MemoryManager) GetMemoryBlock(id BlockId) *Block {
	m.mu.Lock()
	defer m.mu.Unlock()
	block, status := m.bStorage.getMemoryBlock(id)
	if !status {
		panic("Invalid operation while getting memory block: " + id)
	}
	return &block
}

func (m *MemoryManager) GetBlockIdQueue() <-chan BlockId {
	return m.blockIdCh
}

func (m *MemoryManager) GetTotalBlockCount() int {
	return m.bStorage.getBlockCount()
}

/**
 * User ->
 */
// Another kind of design
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
