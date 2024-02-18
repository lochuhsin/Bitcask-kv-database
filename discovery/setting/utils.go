package setting

import (
	"encoding/json"
	"fmt"
)

func SetupConfig(envPaths ...string) {
	Config = NewConfiguration(
		envPaths,
		setClusterMemberCount(),
		setUDPServerPort(),
	)
	configString, _ := json.MarshalIndent(Config, "", "\t")
	fmt.Println(string(configString))
}
