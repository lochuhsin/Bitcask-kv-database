package internal

import "rebitcask/internal/models"

type SegmentMap struct {
	segID   int    // for file naming
	segHead string // since the Segment will be sorted, save segment head to speed up query
	segEnd  string
}

type SegmentContainer struct {
	memo     models.BinarySearchTree // key is head, values is Segment Map
	segCount int
}

type envVariables struct {
	logFolder           string
	segmentFolder       string
	tombstone           string
	memoryKeyCountLimit int
	fileByteLimit       int
	segFileCountLimit   int
}
