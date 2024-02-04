package settings

import (
	"encoding/json"
	"fmt"
)

func SetupConfig(envPaths ...string) {
	Config = NewConfiguration(
		envPaths,
		setClusterMemberCount(),
		setHttpPort(),
		setClusterWaitMemberTimeout(),
	)
	configString, _ := json.MarshalIndent(Config, "", "\t")
	fmt.Println(string(configString))
}
