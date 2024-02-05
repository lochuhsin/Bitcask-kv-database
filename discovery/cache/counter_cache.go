package cache

import (
	"context"
	"fmt"
	"strconv"
)

type counterCache string

var CounterCache counterCache = "counter"

func (c *counterCache) Add(ctx context.Context) {
	client := GetClient()
	err := client.Incr(ctx, fmt.Sprintf("%v::%v", *c, *c)).Err()
	if err != nil {
		panic(err)
	}
}

func (c *counterCache) Count(ctx context.Context) int {
	client := GetClient()
	val, err := client.Get(ctx, fmt.Sprintf("%v::%v", *c, *c)).Result()
	if err != nil {
		panic(err)
	}

	intVal, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}
	return intVal
}
