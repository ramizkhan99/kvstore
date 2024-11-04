package server

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	pb "github.com/ramizkhan99/kvserver/src/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	start_port = 50051
	end_port   = 50100
)

func (s *server) discover() {
	log.Println("server discovery started")
	var wg sync.WaitGroup
	for port := start_port; port <= end_port; port++ {
		if uint8(port-port_offset) == s.id {
			continue
		}
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			s.connectServer(port)
		}(port)
	}
	wg.Wait()
	log.Println("server discovery completed")
}

func (s *server) connectServer(port int) {
	conn, err := grpc.NewClient(fmt.Sprintf("localhost:%d", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client := pb.NewServerClient(conn)
	res, err := client.Ping(ctx, &emptypb.Empty{})
	if err != nil {
		return
	}
	log.Printf("server %d is active: %+v", port-port_offset, res)

	c := connectedServer{
		id:        uint8(port - port_offset),
		is_active: true,
		client:    &client,
	}

	s.addToNetwork(port, &c)
}

func (s *server) addToNetwork(port int, c *connectedServer) {
	s.mu.Lock()
	s.network[uint8(port-port_offset)] = c
	s.mu.Unlock()
}
