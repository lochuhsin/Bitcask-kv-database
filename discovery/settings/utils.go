package settings

import (
	"encoding/json"
	"fmt"
)

func SetupConfig(envPaths ...string) {
	Config = NewConfiguration(
		envPaths,
		SetClusterMemberCount(),
		SetHttpPort(),
	)
	configString, _ := json.MarshalIndent(Config, "", "\t")
	fmt.Println(string(configString))
}
