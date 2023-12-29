package segment

import (
	"bufio"
	"fmt"
	"os"
	"rebitcask/internal/storage/dao"
	"time"

	"github.com/google/uuid"
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
 *
 * Then we have another
 */

// Design a segment index structure that match the upper condition
// implemented using binary search tree
type Segment struct {
	id          string
	level       int    // reference from levelDB, using level indicate the compaction process
	smallestKey string // indicates the smallest key in current segment
	timestamp   int64  // the time that segment was created
	keyCount    int
}

func InitSegment(segCount int64) Segment {
	// the reason of adding segcount is that
	// the creation of a segment is too fast that even nano seconds
	// could not distinguish between segments order
	return Segment{id: uuid.New().String(), level: 0, smallestKey: "", timestamp: time.Now().UnixNano() + segCount}
}

func (s *Segment) Get(k dao.NilString) (dao.Base, bool) {

	filePath := getSegmentFilePath(s.id)
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

func (s *Segment) GetbyOffset(key dao.NilString, offset int, datalen int) (dao.Base, bool) {
	filePath := getSegmentFilePath(s.id)
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

type SegmentStack struct {
	stack []Segment
}

func InitSegmentStack() SegmentStack {
	return SegmentStack{stack: []Segment{}}
}

func (s *SegmentStack) Add(seg Segment) {
	s.stack = append(s.stack, seg)
}

func (s *SegmentStack) Pop() (Segment, bool) {
	for len(s.stack) > 0 {
		seg := s.stack[len(s.stack)-1]
		s.stack = s.stack[:len(s.stack)-1]
		return seg, true
	}
	return *new(Segment), false
}

func (s *SegmentStack) Size() int {
	return len(s.stack)
}

func (s *SegmentStack) list() *[]Segment {

	newSeg := []Segment{}

	for i := len(s.stack) - 1; i >= 0; i-- {
		newSeg = append(newSeg, s.stack[i])
	}
	return &newSeg
}

// order_by timestamp
type SegmentCollection struct {
	/**
	 * We are using stack to get the native characteristics
	 * of first in last out, which meets the requirements of
	 * order by timestamp
	 */
	noneLevelSeg SegmentStack
	level        map[int][]Segment
	maxLevel     int // whenever a compaction starts, adjust this maxLevel
}

func InitSegmentCollection() SegmentCollection {
	stack := InitSegmentStack()
	return SegmentCollection{level: map[int][]Segment{}, noneLevelSeg: stack, maxLevel: 0}
}

func (s *SegmentCollection) Add(seg Segment) {
	s.noneLevelSeg.Add(seg)
}

func (s *SegmentCollection) CompactionCondition() bool {
	/**
	 * Implement the compaction condtion for manager to determine
	 * When we are starts to compact
	 */
	return false
}

func (s *SegmentCollection) Compaction() {
	panic("not implemented yet")
}

type OffsetLen struct {
	Offset int
	Len    int
}

func InitOffsetLen(offset, Len int) OffsetLen {
	return OffsetLen{offset, Len}
}

func (o OffsetLen) Format() string {
	return fmt.Sprintf("%v::%v", o.Offset, o.Len)
}

type SegmentIndex struct {
	id          string
	smallestKey string
	offsetMap   map[dao.NilString]OffsetLen
}

func InitSegmentIndex(sid string) SegmentIndex {
	return SegmentIndex{id: sid, smallestKey: "", offsetMap: map[dao.NilString]OffsetLen{}}
}

func (s *SegmentIndex) Set(k dao.NilString, offset, len int) {
	s.offsetMap[k] = InitOffsetLen(offset, len)
}

func (s *SegmentIndex) Get(k dao.NilString) (OffsetLen, bool) {
	offset, ok := s.offsetMap[k]
	return offset, ok
}

type SegmentIndexCollection struct {
	indexMap map[string]*SegmentIndex
}

func InitSegmentIndexCollection() SegmentIndexCollection {
	//TODO:  1. try to initialize from .koshint files
	// if none of the exists, create an empty one

	// TODO: possibly, we could do without using pointer ?
	return SegmentIndexCollection{map[string]*SegmentIndex{}}
}

func (s *SegmentIndexCollection) Add(sid string, segIndex *SegmentIndex) {
	s.indexMap[sid] = segIndex
}

func (s *SegmentIndexCollection) Get(sid string) (*SegmentIndex, bool) {
	segIndex, ok := s.indexMap[sid]
	return segIndex, ok
}
