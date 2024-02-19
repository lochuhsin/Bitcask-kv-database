package main

import (
	"fmt"
	"net"
	"rebitcask/api/chore"
	"rebitcask/api/core"
	_ "rebitcask/docs"
	"rebitcask/server/rebitcaskpb"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
)

func clusterSetup() {
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
