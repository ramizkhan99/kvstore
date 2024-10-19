package main

import (
	"context"
	"flag"
	"log"
	"sync"

	pb "github.com/ramizkhan99/kvclient/src/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	defaultAddr = "localhost:50051"
)

var (
	addr    = flag.String("addr", defaultAddr, "The address to connect to")
	key     = flag.String("key", "key", "The key to set")
	clients = flag.Int("clients", 1, "The number of clients to connect to")
)

func main() {
	flag.Parse()

	var wg sync.WaitGroup
	wg.Add(*clients)

	for i := 0; i < *clients; i++ {
		go func() {
			defer wg.Done()
			Connect()
		}()
	}
	wg.Wait()
}

func Connect() {
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewStoreClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r, err := c.Set(ctx, &pb.SetRequest{Key: *key, Value: "value"})
	if err != nil {
		log.Fatalf("could not set: %v", err)
	}
	log.Printf("Set: %s", r.GetResult())

	res, err := c.Get(ctx, &pb.GetRequest{Key: *key})
	if err != nil {
		log.Fatalf("could not get: %v", err)
	}
	log.Printf("Get: %s", res.GetValue())
}
