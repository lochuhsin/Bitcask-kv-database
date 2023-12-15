package segment

import (
	"rebitcask/internal/settings"
	"rebitcask/internal/storage/dao"
	"rebitcask/internal/storage/memory"
)

type SegmentManager struct {
	collection SegmentCollection
}

func InitSegmentManager() *SegmentManager {
	return &SegmentManager{collection: InitSegmentCollection()}
}

func (s *SegmentManager) Get(k dao.NilString) (val dao.Base, status bool) {
	/**
	 * TODO: Implement check segment index
	 */
	//zero level segments should be ordered by timestamp
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

func (s *SegmentManager) ConvertToSegment(m memory.IMemory) {
	/**
	 * First we generate a new segment
	 */
	pairs := m.GetAll()
	newSeg := InitSegment(int64(s.collection.zeroLevelSeg.Size()))
	newSeg.WriteFile(pairs)

	s.collection.AddSegment(newSeg)

	// Check compaction condition, if meets
	if s.collection.CompactionCondition() {
		s.collection.Compaction()
	}
	defer memory.MemoryReset(memory.ModelType(settings.ENV.MemoryModel))
	/**
	 * TODO: Then we generate the Segment Index for further search
	 */
}
