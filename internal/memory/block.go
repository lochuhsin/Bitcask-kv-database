package memory

import (
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
	if m.currNode.block.Id == id {
		m.currNode = nil
	}
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
	if m.currNode == nil {
		return nil
	}
	return m.currNode.block
}

func (m *blockStorage) iterateExistingBlocks() []*Block {
	// iterate backwards
	node := m.currNode
	blocks := []*Block{}
	for node != nil && node.block != nil {
		blocks = append(blocks, node.block)
		node = node.prev
	}
	return blocks
}

func (m *blockStorage) getBlockCount() int {
	return len(m.blockMap)
}
