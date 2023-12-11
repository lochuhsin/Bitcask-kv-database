package settings

type envVar struct {
	LogPath           string
	SegmentFolder     string
	Tombstone         string
	NilData           string
	MemoryCountLimit  int
	MemoryLogFolder   string
	MemoryLogFile     string
	SegLineLimit      int
	SegFileCountLimit int // used for merge segments or change to other
	SegmentLogFolder  string
	SegmentLogFile    string
	SegmentFileExt    string
}

// TODO: Convert this to singleton
var ENV envVar

const ENVPATH = "./rebitcask.env"
