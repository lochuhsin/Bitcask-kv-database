package internal

import "rebitcask/internal/models"

type SegmentContainer struct {
	memo models.MemoryModel // key is head, values is Segment Map
}

// Init Change this to arguments, to pass difference memory model type
func (segContainer *SegmentContainer) Init() {
	segContainer.memo = new(models.BinarySearchTree)
}

type envVariables struct {
	logFolder           string
	segmentFolder       string
	tombstone           string
	memoryKeyCountLimit int
	fileLineLimit       int
	segFileCountLimit   int
}
