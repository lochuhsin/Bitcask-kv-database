package memory

import (
	"rebitcask/internal/dao"
	"rebitcask/internal/memory/models"
	"sync"
	"time"

	"github.com/google/uuid"
)

type MemoryManager struct {
	bStorage                 *blockStorage
	activeBlock              *Block
	scheduleBlock            chan BlockId
	setResponseQ             chan dao.Entry
	setRequestQ              chan dao.Entry
	bulkRemoveBlockResponseQ chan []BlockId
	bulkRemoveBlockRequestQ  chan []BlockId
	entryCountLimit          int
	mu                       sync.RWMutex
	modelType                models.ModelType
}

func NewMemoryManager(bStorage *blockStorage, entryCountLimit, blockIdChanSize int, modelType models.ModelType) *MemoryManager {
	activeBlock := NewBlock(
		time.Now().Unix(),
		BlockId(uuid.NewString()),
		modelType,
	)
	return &MemoryManager{
		bStorage:                 bStorage,
		activeBlock:              &activeBlock,
		mu:                       sync.RWMutex{},
		entryCountLimit:          entryCountLimit,
		modelType:                modelType,
		scheduleBlock:            make(chan BlockId, blockIdChanSize),
		setResponseQ:             make(chan dao.Entry),
		setRequestQ:              make(chan dao.Entry),
		bulkRemoveBlockResponseQ: make(chan []BlockId),
		bulkRemoveBlockRequestQ:  make(chan []BlockId),
	}
}

func (m *MemoryManager) Get(key []byte) (dao.Entry, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	// first search active block, second search history blocks
	entry, status := m.activeBlock.Get(key)
	if status {
		return entry, status
	}
	blocks := m.bStorage.getAll()
	for _, block := range blocks {
		entry, status := block.Get(key)
		if status {
			return entry, status
		}
	}
	return dao.Entry{}, false
}

func (m *MemoryManager) GetBlock(id BlockId) *Block {
	m.mu.RLock()
	defer m.mu.RUnlock()
	block, status := m.bStorage.get(id)
	if status {
		return block
	}
	return nil
}

func (m *MemoryManager) GetScheduleBlockQueue() chan BlockId {
	return m.scheduleBlock
}

func (m *MemoryManager) SetRequestQ() chan dao.Entry {
	return m.setRequestQ
}

func (m *MemoryManager) SetResponseQ() chan dao.Entry {
	return m.setResponseQ
}

func (m *MemoryManager) BulkRemoveBlockRequestQ() chan []BlockId {
	return m.bulkRemoveBlockRequestQ
}

func (m *MemoryManager) BulkRemoveBlockResponseQ() chan []BlockId {
	return m.bulkRemoveBlockResponseQ
}

// Long running goroutine for listening write operations
func (m *MemoryManager) WriteOpListener() {
	for {
		select {
		case entry := <-m.setRequestQ:
			m.activeBlock.Set(entry)
			if m.activeBlock.GetSize() > m.entryCountLimit {
				m.mu.Lock()
				processBlockId := m.activeBlock.Id
				m.bStorage.set(processBlockId, m.activeBlock)
				newBlock := NewBlock(
					time.Now().Unix(),
					BlockId(uuid.NewString()),
					m.modelType,
				)
				m.activeBlock = &newBlock
				m.mu.Unlock()
				m.scheduleBlock <- processBlockId
			}
			m.setResponseQ <- entry
		case blockIds := <-m.bulkRemoveBlockRequestQ:
			m.mu.Lock()
			for _, id := range blockIds {
				m.bStorage.delete(id)
			}
			m.mu.Unlock()
		}
	}
}
