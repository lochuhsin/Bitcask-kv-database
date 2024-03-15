package setting

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

type ConnectionTuple struct {
	Ip   string `json:"Ip"`
	Port string `json:"Port`
}

type PeerList struct {
	Peers []ConnectionTuple `json:"Peers"`
}

type config struct {
	DATA_FOLDER_PATH         string
	TOMBSTONE                string
	NIL_DATA_REP             string
	MEMORY_COUNT_LIMIT       int
	MEMORY_MODEL             string
	SEGMENT_FILE_COUNT_LIMIT int // used for merge segments or change to other
	PORT                     string
	CLUSTER_SETUP_HOST       string
	SERVER_ID                string // used for cluster register
	MODE                     Mode
	PEERS                    []ConnectionTuple
}

func NewDefaultConfiguration() config {
	return config{
		DATA_FOLDER_PATH:         "./rbData/",
		TOMBSTONE:                "!@#$%^&*()_+",
		NIL_DATA_REP:             ")(*&^)!@!@#$%^&*()",
		MEMORY_MODEL:             "hash",
		MEMORY_COUNT_LIMIT:       1000000,
		SEGMENT_FILE_COUNT_LIMIT: 100,
		PORT:                     ":8080",
		CLUSTER_SETUP_HOST:       "discovery-app:9000",
		SERVER_ID:                uuid.New().String(),
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

func SetDataFolderPath() Option {
	return func(conf *config) {
		if DataPath := os.Getenv("DATA_FOLDER_PATH"); DataPath != "" {
			conf.DATA_FOLDER_PATH = DataPath
		}
	}
}

func SetTombstone() Option {
	return func(conf *config) {
		if tombstone := os.Getenv("TOMBSTONE"); tombstone != "" {
			conf.TOMBSTONE = tombstone
		}
	}
}

func SetNilData() Option {
	return func(conf *config) {
		if nilData := os.Getenv("NILDATA"); nilData != "" {
			conf.NIL_DATA_REP = nilData
		}
	}
}

func SetMemoryCountLimit() Option {
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

func SetMemoryModel() Option {
	return func(conf *config) {
		if memoryModel := os.Getenv("MEMORY_MODEL"); memoryModel != "" {
			conf.MEMORY_MODEL = memoryModel
		}
	}
}

func SetSegmentFileCountLimit() Option {
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

func SetPort() Option {
	return func(conf *config) {
		if port := os.Getenv("PORT"); port != "" {
			if port[0] != ':' {
				conf.PORT = ":" + port
			} else {
				conf.PORT = port
			}
		}
	}
}

func SetDiscoveryHost() Option {
	return func(conf *config) {
		if host := os.Getenv("CLUSTER_SETUP_HOST"); host != "" {
			conf.CLUSTER_SETUP_HOST = host
		}
	}
}

func SetServerName() Option {
	return func(conf *config) {
		if name := os.Getenv("SERVER_ID"); name != "" {
			conf.SERVER_ID = name
		}
	}
}

func SetMode() Option {
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

// During after udp setup
func SetPeerList(peers PeerList) Option {
	return func(conf *config) {
		conf.PEERS = peers.Peers
	}
}
