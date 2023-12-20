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
	/**
	 * TODO: Implement check segment index
	 */
	//zero level segments should be ordered by timestamp
	for _, segment := range *(s.collection.noneLevelSeg.list()) {

		if k.GetVal().(string) < segment.smallestKey {
			continue
		}

		sid := segment.id
		segIndex, status := s.indexCollection.Get(sid)
		if !status {
			fmt.Println(fmt.Sprintf("Notice!!!! segment index of %v not found", sid))
		}

		offset, ok := segIndex.Get(k)

		if ok {
			fmt.Println(offset)
			// open this when get by offset is implemented
			// val, status := segment.GetbyOffset(offset)
			// return val, status
		} else {
			// since segIndex is primary index, we assume that
			// this index will always be consistent with segment
			continue
		}

		val, status := segment.Get(k)
		if !status {
			continue
		}
		return val, true
	}

	// todo: search by each segment level
	for l := 0; l < s.collection.maxLevel; l++ {
		if segments, ok := s.collection.level[l]; ok {
			for _, segment := range segments {

				sid := segment.id
				segIndex, status := s.indexCollection.Get(sid)
				if !status {
					fmt.Println(fmt.Sprintf("Notice!!!! segment index of %v not found", sid))
				}

				offset, ok := segIndex.Get(k)

				if ok {
					fmt.Println(offset)
					// open this when get by offset is implemented
					// val, status := segment.GetbyOffset(offset)
					// return val, status
				} else {
					// since segIndex is primary index, we assume that
					// this index will always be consistent with segment
					continue
				}

				val, status := segment.Get(k)
				if !status {
					continue
				}
				return val, true
			}
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

	// Write to segment file and generate segment index in the same time
	writeSegmentToFile(&newSeg, &newSegIndex, pairs)

	s.collection.Add(newSeg)
	s.indexCollection.Add(newSeg.id, &newSegIndex)

	// Check compaction condition, if meets
	if s.collection.CompactionCondition() {
		s.collection.Compaction()
	}
	/**
	 * TODO: Then we generate the Segment Index for further search
	 */
}
