package main

import (
	"context"
	"fmt"
	"net"
	"rebitcask/server/chorepb"
	"rebitcask/server/rebitcaskpb"

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
