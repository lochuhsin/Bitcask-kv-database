package cache

import (
	"rebitcask/discovery/settings"
)

func InitClient() {
	cOnce.Do(func() {
		client = NewClient(
			settings.Config.REDIS_HOST,
			settings.Config.REDIS_PASSWORD,
			settings.Config.REDIS_DEFAULT_DB,
		)
	})
}
