package cache

import (
	"context"
	"fmt"
	"time"
)

type ClusterInfoKey string

const (
	ClusterStatus ClusterInfoKey = "status"
)

type clusterCache string

var ClusterCache clusterCache = "cluster"

func (c *clusterCache) Set(ctx context.Context, key ClusterInfoKey, value any) {
	client := GetClient()
	client.Set(ctx, fmt.Sprintf("%v::%v", *c, key), value, time.Hour*24)
}

func (c *clusterCache) Get(ctx context.Context, key ClusterInfoKey) string {
	client := GetClient()
	return client.Get(ctx, fmt.Sprintf("%v::%v", *c, key)).Val()
}
