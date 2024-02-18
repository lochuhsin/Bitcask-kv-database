package main

import (
	"rebitcask"
	"rebitcask/internal/setting"
)

func main() {
	flags := ParseFlags()
	rebitcask.Setup(flags.envPaths...)

	if setting.Config.MODE == setting.CLUSTER {
		clusterSetup()
	}
	go grpcServerSetup(setting.Config.GRPC_PORT)
	httpServerSetup(setting.Config.HTTP_PORT)
}
