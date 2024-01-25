package segment

import (
	"bufio"
	"os"
	"rebitcask/internal/dao"
	"sort"
	"sync"
	"time"
)

/**
 * I'm going to use SSTable as segment implementation
 * The key point of SSTable is that the keys were sorted
 * in ascending order. Therefore the head of the file (usually it's the end of the file)
 * is the smallest key. This is helpful, since we store the smallest key of the segment
 * in memory. When we are looking up to see if key exists,
 * we only need to start looking at files that Segkeies who were smaller.
 * This increases the performance of lookup.
 *
 * Each segment accompanies a segment index
 * which contains all the key and offset to the segment
 */

type Segment struct {
	Id          string
	Level       int    // reference from levelDB, using level indicate the compaction process
	smallestKey string // indicates the smallest key in current segment
	timestamp   int64  // the time that segment was created
	keyCount    int
	pIndex      *PrimaryIndex
}

func NewSegment(id string, pIndex *PrimaryIndex, smallesKey string, keyCount int) Segment {
	// the reason of adding segcount is that
	// the creation of a segment is too fast that even nano seconds
	// could not distinguish between segments order
	return Segment{Id: id, Level: 0, smallestKey: smallesKey, keyCount: keyCount, timestamp: time.Now().UnixNano(), pIndex: pIndex}
}

func (s *Segment) Get(k dao.NilString) (dao.Base, bool) {

	filePath := getSegmentFilePath(s.Id)
	fd, err := os.Open(filePath)
	if err != nil {
		panic("Cannot open segment file")
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()
		pair, err := dao.DeSerialize(line) // Figure out a better way to split between keys
		if err != nil {
			panic("Something went wrong while deserializing data")
		}

		if pair.Key.IsEqual(k) {
			return pair.Val, true
		}
	}

	return nil, false
}

func (s *Segment) GetFromPrimaryIndex(key dao.NilString) (dao.Base, bool) {
	offsetLen, ok := s.pIndex.Get(key)
	if !ok {
		return nil, false
	}
	offset, datalen := offsetLen.Offset, offsetLen.Len

	filePath := getSegmentFilePath(s.Id)
	fd, err := os.Open(filePath)
	if err != nil {
		panic("Something went wrong while opening segment file")
	}
	defer fd.Close()

	byteBuffer := make([]byte, datalen)

	fd.Seek(int64(offset), 0)
	n, err := fd.Read(byteBuffer)
	if err != nil {
		panic("Something went wrong while reading segment file")
	}

	if n != datalen {
		panic("something went wrong wuth the segment data, length doesn't match")
	}
	pair, err := dao.DeSerialize(string(byteBuffer))

	if err != nil {
		panic("is the data valid?")
	}

	// validate key match
	if !pair.Key.IsEqual(key) {
		panic("Key does not match the value")
	}

	return pair.Val, true
}

func (s *Segment) GetPrimayIndex() *PrimaryIndex {
	return s.pIndex
}

type Manager struct {
	sync.Mutex
	levelMap [][]Segment // using 2-d array, index of segments
	maxLevel int         // whenever a compaction starts, adjust this maxLevel
}

func NewSegmentManager() Manager {
	return Manager{levelMap: make([][]Segment, 10), maxLevel: 0}
}

func (s *Manager) Add(seg Segment) {
	s.Lock()

	level := seg.Level

	if level >= len(s.levelMap) {
		newLevelMap := make([][]Segment, 2*len(s.levelMap))
		copy(newLevelMap, s.levelMap)
		s.levelMap = newLevelMap
	}

	if s.levelMap[level] == nil {
		s.levelMap[level] = []Segment{seg}
		s.maxLevel = level
	}

	s.levelMap[level] = append(s.levelMap[level], seg)
	s.Unlock()
}

func (s *Manager) CompactionCondition() bool {
	panic("not implemented yet")
}

func (s *Manager) Compaction() {
	panic("not implemented yet")
}

func (s *Manager) GetValue(k dao.NilString) (val dao.Base, status bool) {
	s.Lock()
	defer s.Unlock()
	/**
	 * We search the segements in each level.
	 * The problem with first level (level 0) is that their might be duplicate keys.
	 * Therefore we need to sort the segments by timestamp (the larger timestamp
	 * comes to the front). This part could be optimized by adding a specific field
	 * that stores un-compressed segements.
	 *
	 * segments in other levels are simpler. No duplicate keys.
	 */
	for level, segments := range s.levelMap {

		if level == 0 {
			sort.Slice(segments, func(i, j int) bool {
				return segments[i].timestamp > segments[j].timestamp
			})
		}
		for _, seg := range segments {

			if k.GetVal().(string) < seg.smallestKey {
				continue
			}

			val, status := seg.GetFromPrimaryIndex(k)
			if status {
				return val, status
			}
			// since segIndex is primary index, we assume that
			// this index will always be consistent with segment
		}
	}
	return nil, false
}
