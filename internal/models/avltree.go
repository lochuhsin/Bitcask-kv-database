package models

type avlTree struct {
	Root *avlnode
	size int
}

type avlnode struct {
	key      string
	val      string
	parent   *avlnode
	children [2]*avlnode
	b        int8 // balance factor
}

func (avl *avlTree) Init() {}

func (avl *avlTree) Set(key string, val []byte) {
}

func (avl *avlTree) GetSize() (s int) {
	return avl.size
}

func (avl *avlTree) SetMemory() {}

// using inorder traversal to make sure
// from small to large
func (avl *avlTree) GetAll() {}
