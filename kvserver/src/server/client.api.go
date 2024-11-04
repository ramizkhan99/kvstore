package server

import (
	"context"
	"log"

	pb "github.com/ramizkhan99/kvserver/src/generated"
	"github.com/ramizkhan99/kvserver/src/store"
)

func (s *server) Set(_ context.Context, in *pb.SetRequest) (*pb.SetResponse, error) {
	log.Printf("Setting pair in db: %v -> %v", in.GetKey(), in.GetValue())
	err := store.SetKey(in.GetKey(), in.GetValue())
	if err != nil {
		return nil, err
	}
	return &pb.SetResponse{Result: "OK"}, nil
}

func (s *server) Get(_ context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	log.Printf("Getting pair from db: %v", in.GetKey())
	value, err := store.GetKey(in.GetKey())
	if err != nil {
		return nil, err
	}
	return &pb.GetResponse{Value: value}, nil
}

func (server *server) GetPrefix(ctx context.Context, in *pb.GetPrefixRequest) (*pb.GetPrefixResponse, error) {
	// TODO: Define get prefix
	log.Printf("Getting prefix from db: %v", in.GetPrefix())
	// values, err := store.GetPrefix(in.GetPrefix())
	// if err != nil {
	// 	return nil, err
	// }
	return &pb.GetPrefixResponse{Value: nil}, nil
}
