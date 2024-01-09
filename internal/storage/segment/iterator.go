package segment

import (
	"errors"
	"fmt"
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

	if sc.level < segCollection.GetLevel()-1 {
		return true
	}

	segCount, status := segCollection.GetSegmentCountByLevel(sc.level)
	if sc.level == segCollection.GetLevel()-1 && status && sc.index < segCount {
		return true
	}
	return false
}

// This is a huge performance drop, optimize this ... n * nlogln
func (sc *SegmentCollectionIterator) getNext(segCollection *SegmentCollection) (Segment, error) {
	var (
		segments []Segment
		status   bool
	)
	if sc.level == 0 { // if level is zero, we should return by timestamp, since there might be duplcate keys
		segments, status = segCollection.GetSegmentByLevel(0)
		if !status {
			fmt.Println("Something went wrong while reading level 0 segments")
			return *new(Segment), errors.New("something went wrong while reading level 0 segments")
		}
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
