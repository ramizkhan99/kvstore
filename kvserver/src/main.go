package main

import (
	"github.com/ramizkhan99/kvserver/src/server"
	"github.com/ramizkhan99/kvserver/src/store"
)

func main() {
	store.Init()

	server.Connect()

	defer store.CloseDB()
}
