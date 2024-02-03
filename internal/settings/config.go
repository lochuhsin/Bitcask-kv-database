package settings

import (
	"fmt"
	"os"
	"strconv"

	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

/**
 * Every developer should be considered as a grown man that should not ever
 * change the config in runtime.
 */
var Config config

type Option func(*config)

type Mode string

const (
	STANDALONE Mode = "standalone"
	CLUSTER    Mode = "cluster"
)

type config struct {
	DATA_FOLDER_PATH         string
	TOMBSTONE                string
	NIL_DATA_REP             string
	MEMORY_COUNT_LIMIT       int
	MEMORY_MODEL             string
	SEGMENT_FILE_COUNT_LIMIT int // used for merge segments or change to other
	HTTP_PORT                string
	GRPC_PORT                string
	DISCOVERY_HOST           string
	SERVER_NAME              string // used for cluster register
	MODE                     Mode
}

func NewDefaultConfiguration() config {
	return config{
		DATA_FOLDER_PATH:         "./rbData/",
		TOMBSTONE:                "!@#$%^&*()_+",
		NIL_DATA_REP:             ")(*&^)!@!@#$%^&*()",
		MEMORY_MODEL:             "hash",
		MEMORY_COUNT_LIMIT:       1000000,
		SEGMENT_FILE_COUNT_LIMIT: 100,
		HTTP_PORT:                ":8080",
		GRPC_PORT:                ":9090",
		DISCOVERY_HOST:           "http://discovery-app:8765",
		SERVER_NAME:              uuid.New().String(),
		MODE:                     STANDALONE,
	}
}

func NewConfiguration(envPaths []string, options ...Option) config {
	newConfig := NewDefaultConfiguration()
	err := godotenv.Load(envPaths...)
	if err == nil {
		for _, fn := range options {
			fn(&newConfig)
		}
	} else {
		fmt.Println("Possibly no .env file or unable to open, using default")

	}
	configString, _ := json.MarshalIndent(newConfig, "", "\t")
	fmt.Println(string(configString))
	return newConfig
}

func setDataFolderPath() Option {
	return func(conf *config) {
		if DataPath := os.Getenv("DATA_FOLDER_PATH"); DataPath != "" {
			conf.DATA_FOLDER_PATH = DataPath
		}
	}
}

func setTombstone() Option {
	return func(conf *config) {
		if tombstone := os.Getenv("TOMBSTONE"); tombstone != "" {
			conf.TOMBSTONE = tombstone
		}
	}
}

func setNilData() Option {
	return func(conf *config) {
		if nilData := os.Getenv("NILDATA"); nilData != "" {
			conf.NIL_DATA_REP = nilData
		}
	}
}

func setMemoryCountLimit() Option {
	return func(conf *config) {
		if memoryCountLimit := os.Getenv("MEMORY_COUNT_LIMIT"); memoryCountLimit != "" {
			limit, err := strconv.Atoi(memoryCountLimit)
			if err != nil {
				panic(err)
			}
			conf.MEMORY_COUNT_LIMIT = limit
		}
	}
}

func setMemoryModel() Option {
	return func(conf *config) {
		if memoryModel := os.Getenv("MEMORY_MODEL"); memoryModel != "" {
			conf.MEMORY_MODEL = memoryModel
		}
	}
}

func setSegmentFileCountLimit() Option {
	return func(conf *config) {
		if segmentFileCountLimit := os.Getenv("SEGMENT_FILE_COUNT_LIMIT"); segmentFileCountLimit != "" {
			limit, err := strconv.Atoi(segmentFileCountLimit)
			if err != nil {
				panic(err)
			}
			conf.SEGMENT_FILE_COUNT_LIMIT = limit
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

func setGrpcPort() Option {
	return func(conf *config) {
		if port := os.Getenv("GRPC_PORT"); port != "" {
			if port[0] != ':' {
				conf.GRPC_PORT = ":" + "port"
			} else {
				conf.GRPC_PORT = "port"
			}
		}
	}
}

func setDiscoveryHost() Option {
	return func(conf *config) {
		if host := os.Getenv("DISCOVERY_HOST"); host != "" {
			conf.DISCOVERY_HOST = host
		}
	}
}

func setServerName() Option {
	return func(conf *config) {
		if name := os.Getenv("SERVER_NAME"); name != "" {
			conf.SERVER_NAME = name
		}
	}
}

func setMode() Option {
	return func(c *config) {
		if mode := os.Getenv("MODE"); mode != "" {
			switch Mode(mode) {
			case STANDALONE:
				c.MODE = STANDALONE
			case CLUSTER:
				c.MODE = CLUSTER
			default:
				fmt.Printf("Invalid mode: %v, using default, %v \n", mode, STANDALONE)
				c.MODE = STANDALONE
			}
		}
	}
}
