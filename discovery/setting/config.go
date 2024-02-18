package setting

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Config config

type config struct {
	CLUSTER_MEMBER_COUNT int
	UDP_SERVER_PORT      string
}

type Option func(*config)

func NewDefaultConfiguration() config {
	return config{
		CLUSTER_MEMBER_COUNT: 3,
		UDP_SERVER_PORT:      ":9000",
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

func setUDPServerPort() Option {
	return func(conf *config) {
		if udpServerPort := os.Getenv("UDP_SERVER_PORT"); udpServerPort != "" {
			if udpServerPort[0] != ':' {
				udpServerPort = fmt.Sprintf(":%v", udpServerPort)
			}
			conf.UDP_SERVER_PORT = udpServerPort
		}
	}
}
