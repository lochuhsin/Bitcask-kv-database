package memory

import (
	"rebitcask/internal/dao"
	"sync"
	"time"

	"github.com/google/uuid"
)

type BlockId string

type Block struct {
	Timestamp int64
	Id        BlockId
	Memory    IMemory
}

type node struct {
	block *Block
	next  *node
	prev  *node
}

type blockStorage struct {
	/**
	 * TODO: ----------------------------------------------------------------
	 * This is an implementation of Bucket Hashmap, use ring queue (circular queue)
	 * to optimize this.
	 */
	blockMap map[BlockId]*node
	top      *node
	bottom   *node
	currNode *node
}

func NewMemoryStorage() *blockStorage {
	/**
	 * I'm using sentinel node to implement ordered map
	 * as it is simpler to handle edge case (i.e empty)
	 */
	top, bottom := &node{}, &node{}
	top.next, bottom.prev = bottom, top
	pool := &blockStorage{
		blockMap: make(map[BlockId]*node, 100),
		top:      top,
		bottom:   bottom,
		currNode: nil,
	}
	return pool
}

func (m *blockStorage) getMemoryBlock(id BlockId) (Block, bool) {
	node, ok := m.blockMap[id]
	return *node.block, ok
}

func (m *blockStorage) removeMemoryBlock(id BlockId) error {
	node, ok := m.blockMap[id]
	if !ok {
		panic("task id not found in ordered map, data is missing")
	}

	node.prev.next = node.next
	node.next.prev = node.prev
	delete(m.blockMap, id)
	return nil
}

func (m *blockStorage) getCurrentBlockId() BlockId {
	return m.currNode.block.Id
}

func (m *blockStorage) createNewBlock(modelType ModelType) {
	newBlockId := BlockId(uuid.New().String())
	newBlock := Block{
		Id:        newBlockId,
		Memory:    MemoryTypeSelector(modelType),
		Timestamp: time.Now().UnixNano(),
	}
	newNode := node{
		block: &newBlock,
		next:  nil,
		prev:  nil,
	}
	newNode.prev, newNode.next = m.bottom.prev, m.bottom
	m.bottom.prev = &newNode
	m.blockMap[newBlockId] = &newNode
	m.currNode = &newNode
}

func (m *blockStorage) getCurrentBlock() *Block {
	return m.currNode.block
}

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
	curretBlock := m.bStorage.getCurrentBlock()
	return curretBlock.Memory.Get(key)
}

func (m *memoryManager) Set(entry dao.Entry) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// 1. set entry to memory
	memory := m.bStorage.getCurrentBlock().Memory
	memory.Set(entry)
	// 2. check memory block condition if meets
	if memory.GetSize() > m.entryCountLimit {
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
