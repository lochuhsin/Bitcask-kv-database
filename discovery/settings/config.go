package settings

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

var Config config

type config struct {
	CLUSTER_MEMBER_COUNT        int
	HTTP_PORT                   string
	REDIS_HOST                  string
	REDIS_PASSWORD              string
	REDIS_DEFAULT_DB            int
	CLUSTER_WAIT_MEMBER_TIMEOUT time.Duration
}

type Option func(*config)

func NewDefaultConfiguration() config {
	return config{
		CLUSTER_MEMBER_COUNT:        3,
		HTTP_PORT:                   ":8765",
		REDIS_HOST:                  "redis:6379",
		REDIS_PASSWORD:              "",
		REDIS_DEFAULT_DB:            0,
		CLUSTER_WAIT_MEMBER_TIMEOUT: time.Second * 30,
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

func setClusterMemberCount() Option {
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

func setHttpPort() Option {
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

func setClusterWaitMemberTimeout() Option {
	return func(c *config) {
		if timeout := os.Getenv("CLUSTER_WAIT_MEMBER_TIMEOUT"); timeout != "" {
			t, err := strconv.Atoi(timeout)
			if err != nil || t < 0 {
				panic(fmt.Sprintf("invalid timeout value: %v", timeout))
			}
			c.CLUSTER_WAIT_MEMBER_TIMEOUT = time.Second * time.Duration(t)
		}
	}
}
