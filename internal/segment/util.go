package segment

import (
	"fmt"
	"rebitcask/internal/settings"
)

func getSegmentFilePath(segId string) string {
	return fmt.Sprintf("%v%v%v%v", settings.Config.DataFolderPath, settings.SEGMENT_FILE_FOLDER, segId, settings.SEGMENT_FILE_EXT)
}
