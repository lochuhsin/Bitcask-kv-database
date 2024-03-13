package util

import (
	"rebitcask/internal/setting"
	"strings"
)

func GetSegmentFilePath(segId string) string {
	var builder strings.Builder
	builder.WriteString(setting.Config.DATA_FOLDER_PATH)
	builder.WriteString(setting.SEGMENT_FILE_FOLDER)
	builder.WriteString(segId)
	builder.WriteString(setting.SEGMENT_FILE_EXT)
	return builder.String()
}

func GetSegmentIndexFilePath(segId string) string {
	var builder strings.Builder
	builder.WriteString(setting.Config.DATA_FOLDER_PATH)
	builder.WriteString(setting.INDEX_FILE_FOLDER)
	builder.WriteString(segId)
	builder.WriteString(setting.SEGMENT_KEY_OFFSET_FILE_EXT)
	return builder.String()
}

func GetSegmentMetaDataFilePath(segId string) string {
	var builder strings.Builder
	builder.WriteString(setting.Config.DATA_FOLDER_PATH)
	builder.WriteString(setting.SEGMENT_FILE_FOLDER)
	builder.WriteString(segId)
	builder.WriteString(setting.SEGMENT_FILE_METADATA_EXT)
	return builder.String()
}

func GetCommitLogFilePath(commitId string) string {
	var builder strings.Builder
	builder.WriteString(setting.Config.DATA_FOLDER_PATH)
	builder.WriteString(setting.COMMIT_LOG_FOLDER)
	builder.WriteString(commitId)
	builder.WriteString(setting.COMMIT_LOG_FILE_EXT)
	return builder.String()
}
