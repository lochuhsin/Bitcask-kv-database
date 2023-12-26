package settings

type envVar struct {
	LogPath           string
	SegmentFolder     string
	Tombstone         string
	NilData           string
	MemoryCountLimit  int
	MemoryModel       string
	MemoryLogFolder   string
	MemoryLogFile     string
	SegFileCountLimit int // used for merge segments or change to other
	SegmentLogFolder  string
	SegmentLogFile    string
	SegmentFileExt    string
}

// TODO: Convert this to singleton
var ENV envVar

const ENVPATH = "./rebitcask.env"

const DATASAPARATER = ".."
