package setting

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
/**
 * IMPORTANT NOTE: This is a design flaws, for current implementation,
 * This value should be set as large as possible to prevent deadlock
 * Why ? Because in memory manager, with set operation, a lock is acquired.
 * Within the locked code, if memory reaches a certain amount of value, it sends
 * a block id to this buffer. And the buffer is consumed by the scheduler.
 *
 * The major problem occurs when the buffer is full, meaning that the scheduler
 * is trying to acquire the lock to proceed the process, but the memory manager
 * is holding the lock, since it cannot finish sending the block id through the queue.
 *
 * This causes memory manager is waiting for spaces in buffer, but the scheduler
 * is waiting for lock to be released. Causes deadlock.
 *
 * So I'm settings this value to be reasonable large that the buffer
 * should never be full. (Extreme test should be considered, i.e benchmarks)
 */
const MEMORY_BLOCK_BUFFER_COUNT = 1000000

/**
 * Scheduler related settings
 */

const MEMORY_CONVERT_WORKER_COUNT = 50

/**
 * Transaction related parameters
 */

const COMMIT_LOG_FOLDER = "commit/"

const COMMIT_LOG_FILE_EXT = ".comt"
