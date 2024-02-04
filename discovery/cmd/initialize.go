package main

import (
	"rebitcask/discovery/cache"
	"rebitcask/discovery/settings"
)

func Init() {
	settings.SetupConfig()
	cache.InitClient()
}
