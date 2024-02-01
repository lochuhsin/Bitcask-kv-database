package segment

import (
	"fmt"
	"rebitcask/internal/settings"
)

func getSegmentFilePath(segId string) string {
	return fmt.Sprintf("%v%v%v%v", settings.Config.DATA_FOLDER_PATH, settings.SEGMENT_FILE_FOLDER, segId, settings.SEGMENT_FILE_EXT)
}
