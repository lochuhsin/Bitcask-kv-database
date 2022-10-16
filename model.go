package rebitcask

type memoryMap struct {
	keyvalue  map[string]string
	memoLimit int
}

type diskMap struct {
	bytePositionMap map[string]int
	byteLengthMap   map[string]int
	byteFileLength int
}
