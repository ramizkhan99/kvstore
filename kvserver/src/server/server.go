package server

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/ramizkhan99/kvserver/src/generated"
	"github.com/ramizkhan99/kvserver/src/store"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedStoreServer
}

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

func Connect() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterStoreServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
