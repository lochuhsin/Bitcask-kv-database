package models

type avlnode struct {
	key    string
	val    Item
	height int
	left   *avlnode
	right  *avlnode
}
type AVLTree struct {
	root *avlnode
	size int
}

func (avl *AVLTree) Init() {
	avl.size = 0
}

func (avl *AVLTree) GetSize() int {
	return avl.size
}

func (avl *AVLTree) GetAll() []KVPair {
	kvPair := make([]KVPair, 0, avl.size)
	avl.inorder(avl.root, &kvPair)
	return kvPair
}

func (avl *AVLTree) inorder(root *avlnode, kvPair *[]KVPair) {
	if root == nil {
		return
	}

	if root.left != nil {
		avl.inorder(root.left, kvPair)
	}

	*kvPair = append(*kvPair, KVPair{
		Key: root.key,
		Val: root.val,
	})

	if root.right != nil {
		avl.inorder(root.right, kvPair)
	}
}

func (avl *AVLTree) Set(key string, val Item) {
	avl.root = avl.set(avl.root, &key, &val)
}

func (avl *AVLTree) Get(key string) (val Item, status bool) {
	if res := avl.get(avl.root, &key); res != nil {
		return *res, true
	}
	return *new(Item), false
}

func (avl *AVLTree) set(root *avlnode, key *string, val *Item) *avlnode {
	if root == nil {
		avl.size++
		return &avlnode{
			key: *key,
			val: *val,
		}
	}

	if root.key == *key {
		root.val = *val
		return root
	}

	if node := *root; node.key != *key {
		if *key < node.key {
			root.left = avl.set(node.left, key, val)
		} else {
			root.right = avl.set(node.right, key, val)
		}
		return root
	}
	return nil
}

// GetAllValueUnder TODO: Optimize this
func (avl *AVLTree) GetAllValueUnder(key *string) []KVPair {
	valueList := make([]KVPair, 0)
	avl.inorder(avl.root, &valueList)

	if len(valueList) == 0 {
		return make([]KVPair, 0)
	}

	if valueList[0].Key > *key {
		return make([]KVPair, 0)
	}

	if valueList[len(valueList)-1].Key <= *key {
		return valueList
	}

	stopIndex := -1
	for i := 0; i < len(valueList); i++ {
		if valueList[i].Key > *key {
			stopIndex = i
			break
		}
	}
	return valueList[0:stopIndex]
}

func (avl *AVLTree) get(root *avlnode, key *string) (val *Item) {
	if root == nil {
		return nil
	}

	if root.key == *key {
		return &root.val
	}

	if *key < root.key {
		return avl.get(root.left, key)
	} else {
		return avl.get(root.right, key)
	}
}

func (avl *AVLTree) balance(node *avlnode) int {
	if node == nil {
		return 0
	} else {
		return avl.height(node.left) - avl.height(node.right)
	}
}
func (avl *AVLTree) rotation(root *avlnode) *avlnode {
	balanceFactor := avl.balance(root)

	// left heavy
	if balanceFactor > 1 {
		if avl.balance(root.left) < 0 {
			root.left = avl.rotateLeft(root.left)
		}
		return avl.rotateRight(root)

	} else if balanceFactor < 1 { // right heavy
		if avl.balance(root.right) > 0 {
			root.right = avl.rotateRight(root.right)
		}
		return avl.rotateLeft(root)
	}
	return root
}

func (avl *AVLTree) rotateLeft(node *avlnode) *avlnode {
	rightNode := node.right
	centerNode := rightNode.left
	rightNode.left = node
	node.right = centerNode
	avl.updateHeight(node)
	avl.updateHeight(rightNode)
	return rightNode
}

func (avl *AVLTree) rotateRight(node *avlnode) *avlnode {
	leftNode := node.left
	centerNode := leftNode.right
	leftNode.right = node
	node.left = centerNode
	avl.updateHeight(node)
	avl.updateHeight(leftNode)
	return leftNode
}

func (avl *AVLTree) updateHeight(node *avlnode) {

	leftHeight := avl.height(node.left)
	rightHeight := avl.height(node.right)

	if leftHeight > rightHeight {
		node.height = leftHeight + 1
	} else {
		node.height = rightHeight + 1
	}
}

// nil node is -1, single
func (avl *AVLTree) height(node *avlnode) int {

	if node == nil {
		return -1
	}
	return node.height
}
