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
	OffsetMap map[string]OffsetLen
}

func NewSegmentIndex(sid string) PrimaryIndex {
	return PrimaryIndex{OffsetMap: map[string]OffsetLen{}}
}

func (s *PrimaryIndex) Set(k []byte, offset, len int) {
	s.OffsetMap[util.BytesToString(k)] = OffsetLen{offset, len}
}

func (s *PrimaryIndex) Get(k []byte) (OffsetLen, bool) {
	offset, ok := s.OffsetMap[util.BytesToString(k)]
	return offset, ok
}
