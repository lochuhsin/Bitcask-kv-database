package main

import (
	"context"
	"rebitcask/server/chorepb"
	"rebitcask/server/rebitcaskpb"
)

type grpcServer struct {
	rebitcaskpb.UnimplementedRebitcaskServiceServer
}

func (s *grpcServer) GetHeartBeat(context.Context, *chorepb.GetHeartBeatRequest) (*chorepb.GetHeartBeatResponse, error) {
	return &chorepb.GetHeartBeatResponse{
		Status: 200,
	}, nil
}
