package internal

type memoryMap struct {
	keyvalue map[string]string
}

type CurrentSegmentMap struct {
	bytePositionMap  map[string]int
	byteLengthMap    map[string]int
	byteFileLength   int
	CurrentSegmentNo int
}

// Array used as hashmap
// index is segment file No.
type SegmentContainer struct {
	memo []CurrentSegmentMap
}
