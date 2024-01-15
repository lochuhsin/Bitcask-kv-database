package segment

import (
	"fmt"
	"rebitcask/internal/storage/dao"
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
	offsetMap map[dao.NilString]OffsetLen
}

func InitSegmentIndex(sid string) PrimaryIndex {
	return PrimaryIndex{offsetMap: map[dao.NilString]OffsetLen{}}
}

func (s *PrimaryIndex) Set(k dao.NilString, offset, len int) {
	s.offsetMap[k] = OffsetLen{offset, len}
}

func (s *PrimaryIndex) Get(k dao.NilString) (OffsetLen, bool) {
	offset, ok := s.offsetMap[k]
	return offset, ok
}
