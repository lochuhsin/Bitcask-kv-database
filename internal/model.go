package internal

type memoryMap struct {
	keyvalue map[string][]byte
}

type SegmentMap struct {
	bytePositionMap  map[string]int
	byteLengthMap    map[string]int
	byteFileLength   int
	CurrentSegmentNo int // for file naming
}

type SegmentContainer struct {
	memo []SegmentMap
}
