package main

import (
	"fmt"
	"net"
	"rebitcask/api/chore"
	"rebitcask/api/core"
	_ "rebitcask/docs"
	"rebitcask/internal/discovery"
	"rebitcask/internal/settings"
	"rebitcask/internal/util"
	"rebitcask/server/rebitcaskpb"
	"time"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
)

func waitingClusterStatus(client *discovery.Client, status discovery.ClusterStatus) {
	retryCount := 10
	backOffFactor := 2
	sleep := time.Second
	for retryCount > 0 {
		obj, _ := client.GetClusterStatus()
		if obj.Status == status {
			break
		}

		retryCount -= 1
		time.Sleep(sleep)
		sleep *= time.Duration(backOffFactor)
	}

	if retryCount < 0 {
		panic(fmt.Sprintf("unable to wait cluster status to valid status: %v, shutting down server", status))
	}
}

func clusterSetup() {
	/**
	 * 1. register to discovery server
	 * 2. wait the server status to become yellow
	 * 3. retrieve all existing peer list from server
	 * 4. request back to server that everything is ok
	 * 5. wait the server to become green
	 * 6. start running raft ...
	 */

	client := discovery.NewClient(settings.Config.DISCOVERY_HOST, 10, 2)
	client.Register(
		settings.Config.SERVER_NAME,
		util.GetOutboundIP().String(),
	)

	waitingClusterStatus(client, discovery.Yellow)

	// Store peer list somewhere else
	_, _ = client.GetPeers()

	waitingClusterStatus(client, discovery.GREEN)

}

func httpServerSetup(port string) {
	r := gin.Default()
	core.Routes(r)
	chore.Routes(r)
	// starts swagger at localhost:port/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(port) // listen and serve on 0.0.0.0:8080
}

func grpcServerSetup(port string) {
	rbServer := grpcServer{}
	lst, err := net.Listen("tcp", port)
	if err != nil {
		panic("unable to listen port " + port)
	}
	s := grpc.NewServer()
	rebitcaskpb.RegisterRebitcaskServiceServer(s, &rbServer)
	fmt.Println("Starting grpc server on " + port)
	s.Serve(lst)
}
