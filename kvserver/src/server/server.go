package server

import (
	"flag"
	"fmt"
	"log"
	"net"
	"sync"

	pb "github.com/ramizkhan99/kvserver/src/generated"
	"google.golang.org/grpc"
)

const (
	port_offset = 50051
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	id             uint8
	is_coordinator bool
	coordinator    uint8
	mu             sync.Mutex
	network        map[uint8]*connectedServer
	port           int
	Wg             sync.WaitGroup
	grpc_server    *grpc.Server
	pb.UnimplementedStoreServer
	pb.UnimplementedServerServer
}

type connectedServer struct {
	id             uint8
	is_active      bool
	is_coordinator bool
	client         *pb.ServerClient
}

func Start() (*server, error) {
	flag.Parse()

	s, err := createServer(*port)
	if err != nil {
		return nil, err
	}

	s.port = *port
	s.id = uint8(*port - port_offset)
	s.is_coordinator = s.id == 0
	s.network = make(map[uint8]*connectedServer)
	log.Printf("Server %d initialized", s.id)

	s.discover()

	return s, nil
}

func createServer(port int) (*server, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}

	s := &server{}
	gs := grpc.NewServer()
	s.grpc_server = gs

	log.Printf("server starting at %v", lis.Addr())

	s.Wg.Add(1)
	go func() {
		defer s.Wg.Done()
		pb.RegisterStoreServer(gs, s)
		pb.RegisterServerServer(gs, s)
		if err := gs.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	log.Println("server started")

	return s, nil
}

func (s *server) Shutdown() {
	s.mu.Lock()
	s.grpc_server.GracefulStop()
	s.mu.Unlock()
}

func (s *server) GetPort() int {
	return s.port
}

func (s *server) GetId() uint8 {
	return s.id
}
