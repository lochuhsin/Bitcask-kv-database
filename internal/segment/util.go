package segment

import (
	"fmt"
	"rebitcask/internal/setting"
)

func getSegmentFilePath(segId string) string {
	return fmt.Sprintf("%v%v%v%v", setting.Config.DATA_FOLDER_PATH, setting.SEGMENT_FILE_FOLDER, segId, setting.SEGMENT_FILE_EXT)
}
