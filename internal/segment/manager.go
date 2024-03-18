package segment

import (
	"bufio"
	"os"
	"rebitcask/internal/dao"
	"rebitcask/internal/util"
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
 * we only need to start looking at files that Segment key's who were smaller.
 * This increases the performance of lookup.
 *
 * Each segment accompanies a segment index
 * which contains all the key and offset to the segment
 */

type Segment struct {
	Id          string
	Level       int    // reference from levelDB, using level indicate the compaction process
	smallestKey []byte // indicates the smallest key in current segment
	timestamp   int64  // the time that segment was created
	keyCount    int
	pIndex      *PrimaryIndex
}

func NewSegment(id string, pIndex *PrimaryIndex, smallestKey []byte, keyCount int) Segment {
	// the reason of adding segcount is that
	// the creation of a segment is too fast that even nano seconds
	// could not distinguish between segments order
	return Segment{Id: id, Level: 0, smallestKey: smallestKey, keyCount: keyCount, timestamp: time.Now().UnixNano(), pIndex: pIndex}
}

func (s *Segment) Get(k []byte) (dao.Entry, bool) {

	filePath := util.GetSegmentFilePath(s.Id)
	fd, err := os.Open(filePath)
	if err != nil {
		panic("Cannot open segment file")
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()
		entry, err := dao.DeSerialize(line) // Figure out a better way to split between keys
		if err != nil {
			panic("Something went wrong while deserializing data")
		}
		entryString := util.BytesToString(entry.Key)
		kString := util.BytesToString(k)
		if entryString == kString {
			return entry, true
		}
	}

	return dao.Entry{}, false
}

func (s *Segment) GetFromPrimaryIndex(key []byte) (dao.Entry, bool) {
	offsetLen, ok := s.pIndex.Get(key)
	if !ok {
		return dao.Entry{}, false
	}
	offset, datalen := offsetLen.Offset, offsetLen.Len

	filePath := util.GetSegmentFilePath(s.Id)
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
		panic("something went wrong with the segment data, length doesn't match")
	}
	entry, err := dao.DeSerialize(string(byteBuffer))

	if err != nil {
		panic("is the data valid?")
	}

	entryString := util.BytesToString(entry.Key)
	keyString := util.BytesToString(key)
	// validate key match
	if entryString != keyString {
		panic("Key does not match the value")
	}

	return entry, true
}

func (s *Segment) GetPrimaryIndex() *PrimaryIndex {
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
	defer s.Unlock()
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

}

func (s *Manager) Compaction() {
	panic("not implemented yet")
}

func (s *Manager) GetValue(k []byte) (val dao.Entry, status bool) {
	s.Lock()
	defer s.Unlock()
	/**
	 * We search the segments in each level.
	 * The problem with first level (level 0) is that their might be duplicate keys.
	 * Therefore we need to sort the segments by timestamp (the larger timestamp
	 * comes to the front). This part could be optimized by adding a specific field
	 * that stores un-compressed segments.
	 *
	 * segments in other levels are simpler. No duplicate keys.
	 */
	keyString := util.BytesToString(k)
	for level, segments := range s.levelMap {

		if level == 0 {
			sort.Slice(segments, func(i, j int) bool {
				return segments[i].timestamp > segments[j].timestamp
			})
		}
		for _, seg := range segments {
			segSmallestKeyString := util.BytesToString(seg.smallestKey)
			if keyString < segSmallestKeyString {
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
	return dao.Entry{}, false
}
