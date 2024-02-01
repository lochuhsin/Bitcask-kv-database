package settings

import "sync"

var cOnce sync.Once

func InitConfig(envPaths ...string) {
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
			)
		},
	)
}
