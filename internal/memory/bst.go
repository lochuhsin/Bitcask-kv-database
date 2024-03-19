package memory

import (
	"rebitcask/internal/dao"
	"rebitcask/internal/util"
	"sync"
)

type bstnode struct {
	key   []byte
	val   dao.Entry
	left  *bstnode
	right *bstnode
}

type BinarySearchTree struct {
	root   *bstnode
	size   int
	setMu  *sync.Mutex // lock for holding the set operation
	frozen bool
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{nil, 0, &sync.Mutex{}, false}
}

func (bst *BinarySearchTree) GetSize() int {
	return bst.size
}

func (bst *BinarySearchTree) GetAll() []dao.Entry {
	entry := make([]dao.Entry, 0, bst.size)
	bst.inorder(bst.root, &entry)
	return entry
}

func (bst *BinarySearchTree) inorder(node *bstnode, entry *[]dao.Entry) {
	if node == nil {
		return
	}

	if node.left != nil {
		bst.inorder(node.left, entry)
	}

	*entry = append(*entry, node.val)

	if node.right != nil {
		bst.inorder(node.right, entry)
	}
}

func (bst *BinarySearchTree) Set(entry dao.Entry) {
	bst.setMu.Lock()
	k := entry.Key
	bst.root = bst.set(bst.root, k, entry)
	bst.setMu.Unlock()
}

func (bst *BinarySearchTree) Get(k []byte) (val dao.Entry, status bool) {
	return bst.get(bst.root, k)
}

func (bst *BinarySearchTree) set(node *bstnode, key []byte, val dao.Entry) *bstnode {
	if node == nil {
		bst.size++
		return &bstnode{
			key: key,
			val: val,
		}
	}

	nKeyString := util.BytesToString(node.key)
	keyString := util.BytesToString(key)

	if nKeyString == keyString {
		node.val = val
		return node
	}

	if nKeyString > keyString {
		node.left = bst.set(node.left, key, val)
	} else {
		node.right = bst.set(node.right, key, val)
	}
	return node
}

func (bst *BinarySearchTree) get(node *bstnode, k []byte) (dao.Entry, bool) {
	if node == nil {
		return dao.Entry{}, false
	}

	nKeyString := util.BytesToString(node.key)
	keyString := util.BytesToString(k)

	if nKeyString == keyString {
		return node.val, true
	}

	if keyString < nKeyString {
		return bst.get(node.left, k)
	} else {
		return bst.get(node.right, k)
	}
}
