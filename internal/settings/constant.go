package settings

const DATA_SEPARATOR = ".."

const SEGMENT_FILE_EXT = ".sst"

const SEGMENT_FILE_METADATA_EXT = ".meta"

const SEGMENT_FILE_FOLDER = "seg/"

/**
 * Index backup files
 */

// format segmentId.hint
const INDEX_FILE_FOLDER = "hint/"

const SEGMENT_KEY_OFFSET_FILE_EXT = ".koshint"

/**
 * Log related global variables
 */

const LOG_FOLDER_PATH = "./log/"

const MEMORY_LOG_FOLDER = "mlog/"

const MEMORY_LOG_FILE = "m.log"

const SEGMENT_LOG_FOLDER = "slog/"

const SEGMENT_LOG_FILE = "s.log"

/**
 * Convert to segment scheduler parameters
 */

const TASK_POOL_SIZE = 100

const WORKER_COUNT = 30
