package main

import (
	"context"
	"fmt"
	"net"
	"rebitcask"
	"rebitcask/api/chore"
	"rebitcask/api/core"
	"rebitcask/server/chorepb"
	"rebitcask/server/rebitcaskpb"

	_ "rebitcask/docs"

	"rebitcask/internal/settings"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"google.golang.org/grpc"
)

type grpcServer struct {
	rebitcaskpb.UnimplementedRebitcaskServiceServer
}

func (s *grpcServer) GetHeartBeat(context.Context, *chorepb.GetHeartBeatRequest) (*chorepb.GetHeartBeatResponse, error) {
	return &chorepb.GetHeartBeatResponse{
		Status: 200,
	}, nil
}

func runGRPC(port string) {
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

func main() {
	rebitcask.Init()
	env := settings.ENV
	r := gin.Default()
	core.Routes(r)
	chore.Routes(r)

	// starts swagger at localhost:port/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// use channel to handle goroutine shut down
	go runGRPC(env.GrpcPort)
	r.Run(env.HttpPort) // listen and serve on 0.0.0.0:8080

}
