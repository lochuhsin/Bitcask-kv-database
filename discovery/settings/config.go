package settings

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Config config

type config struct {
	CLUSTER_MEMBER_COUNT int
	HTTP_PORT            string
}

type Option func(*config)

func NewDefaultConfiguration() config {
	return config{
		CLUSTER_MEMBER_COUNT: 3,
		HTTP_PORT:            ":8765",
	}
}

func NewConfiguration(envPaths []string, options ...Option) config {
	newConfig := NewDefaultConfiguration()
	err := godotenv.Load(envPaths...)
	if err == nil {
		for _, opt := range options {
			opt(&newConfig)
		}
	} else {
		fmt.Println("either unable to locate or open the env file")
		fmt.Println("using default values")
	}
	return newConfig
}

func SetClusterMemberCount() Option {
	return func(conf *config) {
		if clusterMemberCount := os.Getenv("CLUSTER_MEMBER_COUNT"); clusterMemberCount != "" {
			count, err := strconv.Atoi(clusterMemberCount)
			if err != nil {
				panic(err)
			}
			conf.CLUSTER_MEMBER_COUNT = count
		}
	}
}

func SetHttpPort() Option {
	return func(conf *config) {
		if port := os.Getenv("HTTP_PORT"); port != "" {
			if port[0] != ':' {
				conf.HTTP_PORT = ":" + port
			} else {
				conf.HTTP_PORT = port
			}
		}
	}
}
