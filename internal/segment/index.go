package segment

import (
	"fmt"
	"rebitcask/internal/dao"
)

type OffsetLen struct {
	Offset int
	Len    int
}

func (o OffsetLen) Format() string {
	return fmt.Sprintf("%v::%v", o.Offset, o.Len)
}

// contains the key offset in segment
// physical position
type PrimaryIndex struct {
	OffsetMap map[dao.NilString]OffsetLen
}

func NewSegmentIndex(sid string) PrimaryIndex {
	return PrimaryIndex{OffsetMap: map[dao.NilString]OffsetLen{}}
}

func (s *PrimaryIndex) Set(k dao.NilString, offset, len int) {
	s.OffsetMap[k] = OffsetLen{offset, len}
}

func (s *PrimaryIndex) Get(k dao.NilString) (OffsetLen, bool) {
	offset, ok := s.OffsetMap[k]
	return offset, ok
}
