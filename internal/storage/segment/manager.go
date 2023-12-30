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

func InitSegmentManager() *SegmentManager {
	return &SegmentManager{collection: InitSegmentCollection(), indexCollection: InitSegmentIndexCollection()}
}

func (s *SegmentManager) Get(k dao.NilString) (val dao.Base, status bool) {
	//zero level segments should be ordered by timestamp

	iter := InitSegmentCollectionIterator()
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
	newSeg := InitSegment(int64(s.collection.noneLevelSeg.Size()))
	newSegIndex := InitSegmentIndex(newSeg.id)

	// Write to segment file and generate segment index, metadata in the same time
	writeSegmentToFile(&newSeg, &newSegIndex, pairs)
	writeSegmentMetadata(&newSeg)
	writeSegmentIndexToFile(&newSegIndex)

	s.collection.Add(newSeg)
	s.indexCollection.Add(newSeg.id, &newSegIndex)

	// Check compaction condition, if meets
	if s.collection.CompactionCondition() {
		s.collection.Compaction()
	}
}
