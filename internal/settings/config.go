package settings

import (
	"fmt"
	"os"
	"strconv"

	"github.com/goccy/go-json"
	"github.com/joho/godotenv"
)

/**
 * Every developer should be considered as a grown man that should not ever
 * change the config in runtime.
 */
var Config config

type Opt func(*config)

type config struct {
	DataFolderPath    string
	Tombstone         string
	NilData           string
	MemoryCountLimit  int
	MemoryModel       string
	SegFileCountLimit int // used for merge segments or change to other
	HttpPort          string
	GrpcPort          string
}

func NewDefaultConfiguration() config {
	return config{
		DataFolderPath:    "./rbData/",
		Tombstone:         "!@#$%^&*()_+",
		NilData:           ")(*&^)!@!@#$%^&*()",
		MemoryModel:       "hash",
		MemoryCountLimit:  1000000,
		SegFileCountLimit: 100,
		HttpPort:          ":8080",
		GrpcPort:          ":9090",
	}
}

func NewConfiguration(envPaths []string, options ...Opt) config {
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

func setDataFolderPath() Opt {
	return func(conf *config) {
		if DataPath := os.Getenv("DATA_FOLDER_PATH"); DataPath != "" {
			conf.DataFolderPath = DataPath
		}
	}
}

func setTombstone() Opt {
	return func(conf *config) {
		if tombstone := os.Getenv("TOMBSTONE"); tombstone != "" {
			conf.Tombstone = tombstone
		}
	}
}

func setNilData() Opt {
	return func(conf *config) {
		if nilData := os.Getenv("NILDATA"); nilData != "" {
			conf.NilData = nilData
		}
	}
}

func setMemoryCountLimit() Opt {
	return func(conf *config) {
		if memoryCountLimit := os.Getenv("MEMORY_COUNT_LIMIT"); memoryCountLimit != "" {
			limit, err := strconv.Atoi(memoryCountLimit)
			if err != nil {
				panic(err)
			}
			conf.MemoryCountLimit = limit
		}
	}
}

func setMemoryModel() Opt {
	return func(conf *config) {
		if memoryModel := os.Getenv("MEMORY_MODEL"); memoryModel != "" {
			conf.MemoryModel = memoryModel
		}
	}
}

func setSegmentFileCountLimit() Opt {
	return func(conf *config) {
		if segmentFileCountLimit := os.Getenv("SEGMENT_FILE_COUNT_LIMIT"); segmentFileCountLimit != "" {
			limit, err := strconv.Atoi(segmentFileCountLimit)
			if err != nil {
				panic(err)
			}
			conf.SegFileCountLimit = limit
		}
	}
}

func setHttpPort() Opt {
	return func(conf *config) {
		if port := os.Getenv("HTTP_PORT"); port != "" {
			if port[0] != ':' {
				conf.HttpPort = ":" + port
			} else {
				conf.HttpPort = port
			}
		}
	}
}

func setGrpcPort() Opt {
	return func(conf *config) {
		if port := os.Getenv("GRPC_PORT"); port != "" {
			if port[0] != ':' {
				conf.GrpcPort = ":" + "port"
			} else {
				conf.GrpcPort = "port"
			}
		}
	}
}
