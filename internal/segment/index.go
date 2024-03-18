package segment

import (
	"rebitcask/internal/util"
	"strconv"
	"strings"
)

type OffsetLen struct {
	Offset int
	Len    int
}

func (o OffsetLen) Format() string {
	var builder strings.Builder
	builder.WriteString(strconv.Itoa(o.Offset))
	builder.WriteString("::")
	builder.WriteString(strconv.Itoa(o.Len))
	return builder.String()
}

// contains the key offset in segment
// physical position
type PrimaryIndex struct {
	offset map[string]OffsetLen
}

func NewSegmentIndex(sid string) PrimaryIndex {
	return PrimaryIndex{offset: map[string]OffsetLen{}}
}

func (s *PrimaryIndex) Set(k []byte, offset, len int) {
	s.offset[util.BytesToString(k)] = OffsetLen{offset, len}
}

func (s *PrimaryIndex) Get(k []byte) (OffsetLen, bool) {
	offset, ok := s.offset[util.BytesToString(k)]
	return offset, ok
}

func (s *PrimaryIndex) GetAllIndex() map[string]OffsetLen {
	return s.offset
}
