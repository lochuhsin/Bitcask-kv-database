package memory

import (
	"rebitcask/internal/dao"
	"sync"
)

type memoryManager struct {
	bStorage        *blockStorage
	blockIdCh       chan BlockId
	entryCountLimit int
	mu              sync.Mutex
	modelType       ModelType
}

func NewMemoryManager(bStorage *blockStorage, entryCountLimit, blockIdChanSize int, modelType ModelType) *memoryManager {
	bStorage.createNewBlock(modelType)
	return &memoryManager{
		bStorage:        bStorage,
		blockIdCh:       make(chan BlockId, blockIdChanSize),
		mu:              sync.Mutex{},
		entryCountLimit: entryCountLimit,
		modelType:       modelType,
	}
}

func (m *memoryManager) Get(key dao.NilString) (dao.Base, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	blocks := m.bStorage.iterateExistingBlocks()
	for _, block := range blocks {
		val, status := block.Memory.Get(key)
		if status {
			return val, status
		}
	}
	return nil, false
}

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
