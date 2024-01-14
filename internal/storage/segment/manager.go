package segment

import (
	"fmt"
	"rebitcask/internal/storage/dao"
	"rebitcask/internal/storage/memory"
	"sync"
)

type Manager struct {
	collection      Collection
	primaryIndexMap sync.Map // type of [segmentId, *PrimaryIndex]
}

func NewSegmentManager() *Manager {
	return &Manager{collection: NewSegmentCollection(), primaryIndexMap: sync.Map{}}
}

func (s *Manager) Get(k dao.NilString) (val dao.Base, status bool) {
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
		segIndex, status := s.primaryIndexMap.Load(sid)
		sIndex := segIndex.(*PrimaryIndex)

		if !status {
			fmt.Printf("Notice!!!! segment index of %v not found \n", sid)
		}

		offsetLen, ok := sIndex.Get(k)

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

func (s *Manager) ConvertToSegment(m memory.IMemory) {
	/**
	 * First we generate a new segment
	 */
	pairs := m.GetAll()
	Seg := NewSegment(int64(s.collection.GetSegmentCount()))
	SegIndex := InitSegmentIndex(Seg.id)

	// Write to segment file and generate segment index, metadata in the same time
	writeSegmentToFile(&Seg, &SegIndex, pairs)

	// TODO: use goroutine to write concurrently, i.e wait group
	writeSegmentMetadata(&Seg)
	writeSegmentIndexToFile(&SegIndex)

	s.collection.Add(Seg)
	s.primaryIndexMap.Store(SegIndex.id, &SegIndex)

	// TODO: Check compaction condition, if meets trigger it, however this should be implemented in scheduler
	if s.collection.CompactionCondition() {
		s.collection.Compaction()
	}
}
