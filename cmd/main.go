package main

import (
	"rebitcask"
	"rebitcask/internal/setting"

	"github.com/sirupsen/logrus"
)

func main() {
	flags := ParseFlags()
	rebitcask.Setup(flags.envPaths...)

	if setting.Config.MODE == setting.CLUSTER {
		clusterSetup()
		logrus.Info("Cluster setup complete")
	}
	for {
	}
	// go grpcServerSetup(setting.Config.GRPC_PORT)
	// httpServerSetup(setting.Config.HTTP_PORT)
}
