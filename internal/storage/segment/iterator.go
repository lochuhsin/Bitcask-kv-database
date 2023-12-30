package segment

import (
	"sort"
)

type SegmentCollectionIterator struct {
	level                int
	index                int
	tempStoreSegmentList []Segment
}

func InitSegmentCollectionIterator() *SegmentCollectionIterator {
	return &SegmentCollectionIterator{0, 0, nil}
}

func (sc *SegmentCollectionIterator) hasNext(segCollection *SegmentCollection) bool {

	if sc.level < len(segCollection.levelMap)-1 {
		return true
	}
	if sc.level == len(segCollection.levelMap)-1 && sc.index < len(segCollection.levelMap[sc.level]) {
		return true
	}
	return false
}

// This is a huge performance drop, optimize this ... n * nlogln
func (sc *SegmentCollectionIterator) getNext(segCollection *SegmentCollection) (Segment, error) {
	var segments []Segment
	if sc.level == 0 {
		segments = segCollection.levelMap[0]
		sort.Slice(segments, func(i, j int) bool {
			return segments[i].timestamp > segments[j].timestamp
		})
	}
	seg := segments[sc.index]
	sc.index++
	if sc.index >= len(segments) {
		sc.index = 0
		sc.level++
	}
	return seg, nil
}
