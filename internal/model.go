package internal

type SegmentMap struct {
	bytePositionMap  map[string]int
	byteLengthMap    map[string]int
	byteFileLength   int
	CurrentSegmentNo int // for file naming
}

type SegmentContainer struct {
	memo     []SegmentMap
	segCount int
}

type envVariables struct {
	logFolder           string
	segmentFolder       string
	tombstone           string
	memoryKeyCountLimit int
	fileByteLimit       int
	segFileCountLimit   int
}
