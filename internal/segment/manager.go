package segment

import (
	"rebitcask/internal/dao"
	"rebitcask/internal/memory"
	"sort"
)

type Manager struct {
	collection Collection
}

func NewSegmentManager() *Manager {
	// TODO: Convert to dependecy injection
	return &Manager{collection: NewSegmentCollection()}
}

func (s *Manager) Get(k dao.NilString) (val dao.Base, status bool) {
	/**
	 * We search the segements in each level.
	 * The problem with first level is that their might be duplicate keys.
	 * Therefore we need to sort the segments by timestamp (the larger timestamp
	 * comes to the front). This part could be optimized by adding a specific field
	 * that stores un-compressed segements.
	 *
	 * segments in other levels are simpler. No duplicate keys.
	 */
	for level := 0; level <= s.collection.GetLevel(); level++ {
		segments, status := s.collection.GetSegmentByLevel(level)
		if !status {
			continue
		}
		if level == 0 {
			sort.Slice(segments, func(i, j int) bool {
				return segments[i].timestamp > segments[j].timestamp
			})
		}

		for _, seg := range segments {
			if k.GetVal().(string) < seg.smallestKey {
				continue
			}
			offsetLen, ok := seg.pIndex.Get(k)
			if ok {
				val, status := seg.GetbyOffset(k, offsetLen.Offset, offsetLen.Len)
				return val, status
			} else {
				// since segIndex is primary index, we assume that
				// this index will always be consistent with segment
				continue
			}
		}
	}

	return nil, false
}

func (s *Manager) ConvertToSegment(m memory.IMemory) {
	pairs := m.GetAll()
	seg := NewSegment(int64(s.collection.GetSegmentCount()))
	segIndex := NewSegmentIndex(seg.id)
	seg.pIndex = &segIndex

	// Write to segment file and generate segment index, metadata in the same time
	segmentToFile(&seg, pairs)
	segmentToMetadata(&seg)
	segmentIndexToFile(&seg)

	s.collection.Add(seg)
}
