package segment

import (
	"fmt"
	"rebitcask/internal/storage/dao"
	"rebitcask/internal/storage/memory"
)

type Manager struct {
	collection Collection
}

func NewSegmentManager() *Manager {
	return &Manager{collection: NewSegmentCollection()}
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
		sIndex := segment.pIndex
		if sIndex == nil {
			fmt.Println(sid, status, sIndex)
			panic("index not found")
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
	seg := NewSegment(int64(s.collection.GetSegmentCount()))
	segIndex := InitSegmentIndex(seg.id)
	seg.pIndex = &segIndex

	// Write to segment file and generate segment index, metadata in the same time
	writeSegmentToFile(&seg, pairs)
	writeSegmentMetadata(&seg)
	writeSegmentIndexToFile(&seg)

	s.collection.Add(seg)
}
