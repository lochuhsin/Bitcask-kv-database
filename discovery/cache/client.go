package cache

import (
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	client *redis.Client
	cOnce  sync.Once
)

type Opt func(*redis.Options)

func NewClient(Addr string, Password string, DB int) *redis.Client {
	defaultOpt := redis.Options{
		Addr:     Addr,
		Password: Password, // no password set
		DB:       DB,       // use default DB
	}
	return redis.NewClient(&defaultOpt)
}

func GetClient() *redis.Client {
	return client
}
