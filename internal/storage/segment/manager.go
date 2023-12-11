package segment

import (
	"rebitcask/internal/storage/dao"
	"rebitcask/internal/storage/memory"
)

type SegmentManager struct {
	collection SegmentCollection
}

func InitSegmentManager() SegmentManager {
	return SegmentManager{collection: InitSegmentCollection()}
}

func (s *SegmentManager) Get(k dao.NilString) (dao.Base, bool) {
	/**
	 * TODO: Implement check segment index
	 */
	// Assuming zero level seg is ordered by timestamp
	for _, segment := range *(s.collection.zeroLevelSeg.list()) {
		if k.GetVal().(string) < segment.smallestKey {
			continue
		}
		val, status := segment.Get(k)
		if !status {
			continue
		}
		return val, true
	}

	// todo: search by each segment level
	for l := 1; l < s.collection.maxLevel; l++ {
		if segments, ok := s.collection.level[l]; ok {
			for _, segment := range segments {

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

func (s *SegmentManager) ConvertToSegment(m memory.MemoryBase) {
	/**
	 * First we generate a new segment
	 */
	pairs := m.GetAll()

	newSeg := InitSegment()
	newSeg.WriteFile(pairs)

	s.collection.AddSegment(newSeg)

	// Check compaction condition, if meets
	if s.collection.CompactionCondition() {
		s.collection.Compaction()
	}

	defer m.Reset()
	/**
	 * Then we generate the Segment Index for further search
	 */
	// TODO:
}
