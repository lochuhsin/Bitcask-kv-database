package main

import (
	"rebitcask"
	"rebitcask/internal/settings"
)

func main() {
	flags := ParseFlags()
	rebitcask.Setup(flags.envPaths...)

	if settings.Config.MODE == settings.CLUSTER {
		clusterSetup()
	}
	go grpcServerSetup(settings.Config.GRPC_PORT)
	httpServerSetup(settings.Config.HTTP_PORT)
}
