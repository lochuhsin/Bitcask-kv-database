package segment

import (
	"fmt"
	"rebitcask/internal/storage/dao"
	"rebitcask/internal/storage/memory"
)

type SegmentManager struct {
	collection      SegmentCollection
	indexCollection SegmentIndexCollection
}

func NewSegmentManager() *SegmentManager {
	return &SegmentManager{collection: NewSegmentCollection(), indexCollection: NewSegmentIndexCollection()}
}

func (s *SegmentManager) Get(k dao.NilString) (val dao.Base, status bool) {
	iter := NewSegmentCollectionIterator()
	for iter.hasNext(&s.collection) {
		segment, err := iter.getNext(&s.collection)
		if err != nil {
			panic("something went wrong with looping segments")
		}
		if k.GetVal().(string) < segment.smallestKey {
			continue
		}
		sid := segment.id
		segIndex, status := s.indexCollection.Get(sid)
		if !status {
			fmt.Printf("Notice!!!! segment index of %v not found \n", sid)
		}

		offsetLen, ok := segIndex.Get(k)

		if ok {
			val, status := segment.GetbyOffset(k, offsetLen.Offset, offsetLen.Len)
			return val, status
		} else {
			// since segIndex is primary index, we assume that
			// this index will always be consistent with segment
			continue
		}
	}

	return nil, false
}

func (s *SegmentManager) ConvertToSegment(m memory.IMemory) {
	/**
	 * First we generate a new segment
	 */
	pairs := m.GetAll()
	Seg := NewSegment(int64(s.collection.GetSegmentCount()))
	SegIndex := InitSegmentIndex(Seg.id)

	// Write to segment file and generate segment index, metadata in the same time
	writeSegmentToFile(&Seg, &SegIndex, pairs)
	writeSegmentMetadata(&Seg)
	writeSegmentIndexToFile(&SegIndex)

	s.collection.Add(Seg)
	s.indexCollection.Add(SegIndex.id, &SegIndex)

	// Check compaction condition, if meets trigger it, however this should be implemented in scheduler
	if s.collection.CompactionCondition() {
		s.collection.Compaction()
	}
}
