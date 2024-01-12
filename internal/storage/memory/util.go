package memory

func MemoryTypeSelector(mType ModelType) IMemory {
	var m IMemory = nil
	switch mType {
	case HASH:
		m = NewHash()
	case BST:
		m = NewBinarySearchTree()

	// TODO: implement these
	// case memory.AVLT:
	// 	m = memory.InitAvlTree()
	// case memory.RBT:
	// 	m = memory.InitRedBlackTree()

	default:
		panic("memory model not implemented errir")
	}
	return m
}
