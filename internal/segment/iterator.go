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

func NewSegmentCollectionIterator() *SegmentCollectionIterator {
	return &SegmentCollectionIterator{0, 0, nil}
}

func (sc *SegmentCollectionIterator) hasNext(segCollection *Collection) bool {

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
func (sc *SegmentCollectionIterator) getNext(segCollection *Collection) (Segment, error) {
	/**
	 * This iterator is really slow, since it first get all the segments in the same level
	 * from segCollection. Second sort the segments by timestamp, which is super slow.
	 * Most importantly, this is a getNext function, therefore this function is under a for loop
	 * The total operation of iterating segments costs O(n) * (O(n) + O(nlogn)), super slow.
	 *
	 * One solution is implement level zero with tree like structure and compare with timestamp
	 * during insertion in level 0, the only level that is sensitive to time ordering. Then remove this iterator mother fucker.
	 */

	var (
		segments []Segment
		status   bool
	)
	if sc.level == 0 { // sensitive to timestamp, should order by timestamp
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
