package memory

import (
	"rebitcask/internal/dao"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Manager struct {
	bStorage                 *blockStorage
	activeBlock              *Block
	scheduleBlock            chan BlockId
	setResponseQ             chan dao.Entry
	setRequestQ              chan dao.Entry
	bulkRemoveBlockResponseQ chan []BlockId
	bulkRemoveBlockRequestQ  chan []BlockId
	entryCountLimit          int
	mu                       sync.RWMutex
	modelType                ModelType
}

func NewMemoryManager(bStorage *blockStorage, entryCountLimit, blockIdChanSize int, modelType ModelType) *Manager {
	activeBlock := NewBlock(
		time.Now().Unix(),
		BlockId(uuid.NewString()),
		modelType,
	)
	return &Manager{
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

func (m *Manager) Get(key []byte) (dao.Entry, bool) {
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

func (m *Manager) GetBlock(id BlockId) *Block {
	m.mu.RLock()
	defer m.mu.RUnlock()
	block, status := m.bStorage.get(id)
	if status {
		return block
	}
	return nil
}

func (m *Manager) GetScheduleBlockQueue() chan BlockId {
	return m.scheduleBlock
}

func (m *Manager) SetRequestQ() chan dao.Entry {
	return m.setRequestQ
}

func (m *Manager) SetResponseQ() chan dao.Entry {
	return m.setResponseQ
}

func (m *Manager) BulkRemoveBlockRequestQ() chan []BlockId {
	return m.bulkRemoveBlockRequestQ
}

func (m *Manager) BulkRemoveBlockResponseQ() chan []BlockId {
	return m.bulkRemoveBlockResponseQ
}

// Long running goroutine for listening write operations
func (m *Manager) WriteOpListener() {
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
			m.bulkRemoveBlockResponseQ <- blockIds
		}
	}
}
