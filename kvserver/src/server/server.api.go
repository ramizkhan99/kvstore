package server

import (
	"context"
	"log"

	pb "github.com/ramizkhan99/kvserver/src/generated"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *server) JoinServer(_ context.Context, in *pb.JoinRequest) (*pb.JoinResponse, error) {
	log.Printf("Received join request from client: %+v", in)
	return nil, nil
}

func (s *server) Ping(_ context.Context, in *emptypb.Empty) (*pb.HeartbeatResponse, error) {
	log.Println("Server pinged")
	return &pb.HeartbeatResponse{Status: pb.ServerStatus_OK}, nil
}
