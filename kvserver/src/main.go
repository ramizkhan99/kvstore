package main

import (
	"log"

	"github.com/ramizkhan99/kvserver/src/server"
	"github.com/ramizkhan99/kvserver/src/store"
)

func main() {
	s, err := server.Start()
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
		return
	}

	store.StartDB(s.GetId())
	defer store.CloseDB()

	defer s.Shutdown()
	defer s.Wg.Wait()
}
