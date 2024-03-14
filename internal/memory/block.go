package memory

import (
	"errors"
	"rebitcask/internal/dao"
	"rebitcask/internal/memory/models"
	"sync"
)

type BlockId string

type Block struct {
	Timestamp int64
	Id        BlockId
	Memory    models.IMemory
	mu        sync.RWMutex
}

func NewBlock(timestamp int64, id BlockId, t models.ModelType) Block {
	return Block{
		Id:        id,
		Memory:    models.MemoryTypeSelector(t),
		Timestamp: timestamp,
		mu:        sync.RWMutex{},
	}
}

func (b *Block) Get(k []byte) (dao.Entry, bool) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.Memory.Get(k)
}
func (b *Block) Set(entry dao.Entry) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.Memory.Set(entry)
}
func (b *Block) GetSize() int {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.Memory.GetSize()
}
func (b *Block) GetAll() []dao.Entry {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.Memory.GetAll()
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
}

func NewBlockStorage() *blockStorage {
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
	}
	return pool
}

func (m *blockStorage) get(id BlockId) (*Block, bool) {
	node, ok := m.blockMap[id]
	return node.block, ok
}

func (m *blockStorage) delete(id BlockId) error {
	node, ok := m.blockMap[id]
	if !ok {
		return errors.New("task id not found in ordered map, data is missing")
	}

	node.prev.next = node.next
	node.next.prev = node.prev
	delete(m.blockMap, id)
	return nil
}

func (m *blockStorage) set(id BlockId, block *Block) {
	newNode := node{
		block: block,
		next:  nil,
		prev:  nil,
	}
	newNode.prev, newNode.next = m.bottom.prev, m.bottom
	m.bottom.prev = &newNode
	m.blockMap[id] = &newNode
}

func (m *blockStorage) getAll() []*Block {
	// iterate backwards from latest to oldest
	node := m.bottom.prev
	blocks := make([]*Block, 0, len(m.blockMap))
	for node != nil && node.block != nil {
		blocks = append(blocks, node.block)
		node = node.prev
	}
	return blocks
}
