package segment

import (
	"fmt"
	"rebitcask/internal/settings"
)

func getSegmentFilePath(segId string) string {
	return fmt.Sprintf("%v%v%v%v", settings.ENV.LogPath, settings.ENV.SegmentFolder, segId, settings.ENV.SegmentFileExt)
}
