package main

import (
	"context"
	"rebitcask/discovery/cache"
	"rebitcask/discovery/settings"
)

func Init() {
	ctx := context.Background()
	settings.SetupConfig()
	cache.InitClient()
	cache.ClusterCache.Set(ctx, cache.Status, "red") //should be red initially
}
