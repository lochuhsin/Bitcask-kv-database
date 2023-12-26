package settings

type envVar struct {
	LogPath           string
	SegmentFolder     string
	Tombstone         string
	NilData           string
	MemoryCountLimit  int
	MemoryModel       string
	SegFileCountLimit int // used for merge segments or change to other
}

// TODO: Convert this to singleton
var ENV envVar

const ENVPATH = "./rebitcask.env"

const DATASAPARATER = ".."

const MEMORY_LOG_FOLDER = "mlog/"

const MEMORY_LOG_FILE = "m.log"

const SEGMENT_LOG_FOLDER = "slog/"

const SEGMENT_LOG_FILE = "s.log"

const SEGMENT_LOG_FILE_EXT = ".sst"
