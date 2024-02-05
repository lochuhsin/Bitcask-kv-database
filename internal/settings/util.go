package settings

import (
	"fmt"
	"os"
	"sync"
)

var cOnce sync.Once

func SetupConfig(envPaths ...string) {
	cOnce.Do(
		func() {
			Config = NewConfiguration(
				envPaths,
				setDataFolderPath(),
				setGrpcPort(),
				setHttpPort(),
				setMemoryCountLimit(),
				setMemoryModel(),
				setNilData(),
				setSegmentFileCountLimit(),
				setTombstone(),
				setDiscoveryHost(),
				setServerName(),
				setMode(),
				setClusterSetup(),
			)
		},
	)
}

func SetUpDirectory() {
	segDir := fmt.Sprintf("%s%s", Config.DATA_FOLDER_PATH, SEGMENT_FILE_FOLDER)
	indexDir := fmt.Sprintf("%s%s", Config.DATA_FOLDER_PATH, INDEX_FILE_FOLDER)
	os.MkdirAll(segDir, os.ModePerm)
	os.MkdirAll(indexDir, os.ModePerm)
}
