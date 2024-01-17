package main

import (
	"context"
	"rebitcask/server/chorepb"
	"rebitcask/server/rebitcaskpb"
)

type Server struct {
	rebitcaskpb.UnimplementedRebitcaskServiceServer
}

func (s *Server) GetHeartBeat(context.Context, *chorepb.GetHeartBeatRequest) (*chorepb.GetHeartBeatResponse, error) {
	return &chorepb.GetHeartBeatResponse{
		Status: 200,
	}, nil
}
