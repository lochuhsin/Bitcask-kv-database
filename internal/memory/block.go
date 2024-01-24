package memory

import (
	"errors"
	"rebitcask/internal/dao"
	"rebitcask/internal/settings"
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

type memoryStorage struct {
	/**
	 * TODO: ----------------------------------------------------------------
	 * This is an implementation of Bucket Hashmap, use ring queue (circular queue)
	 * to optimize this.
	 */
	blockMap       map[BlockId]*node
	top            *node
	bottom         *node
	currNode       *node
	blockTaskQueue chan BlockId
	sync.Mutex
}

func NewMemoryStorage() *memoryStorage {
	/**
	 * I'm using sentinel node to implement ordered map
	 * as it is simpler to handle edge case (i.e empty)
	 */
	top, bottom := &node{}, &node{}
	top.next, bottom.prev = bottom, top
	pool := &memoryStorage{
		blockMap:       make(map[BlockId]*node, 100),
		top:            top,
		bottom:         bottom,
		currNode:       nil,
		blockTaskQueue: make(chan BlockId, settings.WORKER_COUNT),
	}
	pool.currNode = pool.genNewNode()
	return pool
}

func (m *memoryStorage) GetMemoryBlock(id BlockId) (Block, bool) {
	m.Lock()
	defer m.Unlock()
	node, ok := m.blockMap[id]
	return *node.block, ok
}

func (m *memoryStorage) RemoveMemoryBlock(id BlockId) error {
	m.Lock()
	defer m.Unlock()
	if len(m.blockMap) == 0 {
		return errors.New("deleting empty taskOrderedMap is not allowed")
	}
	if _, ok := m.blockMap[id]; !ok {
		panic("task id not found in ordered map, data is missing")
	}
	node := m.blockMap[id]
	node.prev.next = node.next
	node.next.prev = node.prev
	delete(m.blockMap, id)
	return nil
}

func (m *memoryStorage) Get(key dao.NilString) (dao.Base, bool) {
	m.Lock()
	defer m.Unlock()

	if len(m.blockMap) == 0 {
		return nil, false
	}

	// loop backwards, from latest to oldest task
	node := m.currNode
	// the second condition stops when it reaches the
	// top node, which is also a sentinel node
	for node != nil && node.block != nil {
		val, status := node.block.Memory.Get(key)
		if status {
			return val, status
		}
		node = node.prev
	}
	return nil, false
}

func (m *memoryStorage) Set(pair dao.Pair) {
	m.Lock()
	defer m.Unlock()
	m.currNode.block.Memory.Set(pair)
	if m.currNode.block.Memory.GetSize() >= settings.ENV.MemoryCountLimit {
		// add current block to task chan and replace currentblock id to new one
		m.blockTaskQueue <- m.currNode.block.Id
		m.currNode = m.genNewNode()
	}
}

func (m *memoryStorage) genNewNode() *node {
	newBlockId := BlockId(uuid.New().String())
	newBlock := Block{
		Id:        newBlockId,
		Memory:    MemoryTypeSelector(ModelType(settings.ENV.MemoryModel)),
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
	return &newNode
}

func (m *memoryStorage) GetBlockIdChan() chan BlockId {
	return m.blockTaskQueue
}
