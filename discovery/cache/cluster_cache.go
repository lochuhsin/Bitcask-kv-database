package cache

import (
	"context"
	"fmt"
	"time"
)

type ClusterInfoKey string

const (
	Status ClusterInfoKey = "status"
)

type clusterCache string

var ClusterCache clusterCache = "cluster"

func (c *clusterCache) Set(ctx context.Context, key ClusterInfoKey, value any) {
	client := GetClient()
	err := client.Set(ctx, fmt.Sprintf("%v::%v", *c, key), value, time.Hour*24).Err()
	if err != nil {
		panic(err)
	}
}

func (c *clusterCache) Get(ctx context.Context, key ClusterInfoKey) string {
	client := GetClient()
	return client.Get(ctx, fmt.Sprintf("%v::%v", *c, key)).Val()
}
