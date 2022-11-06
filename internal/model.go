package internal

import "rebitcask/internal/models"

type SegmentContainer struct {
	memo models.MemoryModel // key is head, values is Segment Map
}

func (segContainer *SegmentContainer) Init() {
	segContainer.memo = new(models.AVLTree)
}

type envVariables struct {
	logFolder           string
	segmentFolder       string
	tombstone           string
	memoryKeyCountLimit int
	fileLineLimit       int
	segFileCountLimit   int
}
