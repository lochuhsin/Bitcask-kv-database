package memory

import (
	"rebitcask/internal/dao"
	"sync"
)

type bstnode struct {
	key   dao.NilString
	val   dao.Base
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

func (bst *BinarySearchTree) inorder(root *bstnode, entry *[]dao.Entry) {
	if root == nil {
		return
	}

	if root.left != nil {
		bst.inorder(root.left, entry)
	}

	*entry = append(*entry, dao.Entry{
		Key: root.key,
		Val: root.val,
	})

	if root.right != nil {
		bst.inorder(root.right, entry)
	}
}

func (bst *BinarySearchTree) Set(entry dao.Entry) {
	bst.setMu.Lock()
	k, v := entry.Key, entry.Val
	bst.root = bst.set(bst.root, k, v)
	bst.setMu.Unlock()
}

func (bst *BinarySearchTree) Get(k dao.NilString) (val dao.Base, status bool) {
	if res := bst.get(bst.root, k); res != nil {
		return res, true
	}
	return nil, false
}

func (bst *BinarySearchTree) set(root *bstnode, key dao.NilString, val dao.Base) *bstnode {
	if root == nil {
		bst.size++
		return &bstnode{
			key: key,
			val: val,
		}
	}

	if root.key == key {
		root.val = val
		return root
	}

	if node := *root; node.key != key {
		if key.IsSmaller(node.key) {
			root.left = bst.set(node.left, key, val)
		} else {
			root.right = bst.set(node.right, key, val)
		}
		return root
	}
	return nil
}

func (bst *BinarySearchTree) get(root *bstnode, k dao.NilString) (val dao.Base) {
	if root == nil {
		return nil
	}

	if root.key == k {
		return root.val
	}

	if k.IsSmaller(root.key) {
		return bst.get(root.left, k)
	} else {
		return bst.get(root.right, k)
	}
}
