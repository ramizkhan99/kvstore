package store

import (
	"fmt"
	"log"

	"github.com/cockroachdb/pebble"
)

type database struct {
	db *pebble.DB
}

func connectDB(port uint8) (*database, error) {
	db, err := pebble.Open(fmt.Sprintf("store-%d", port), &pebble.Options{})
	if err != nil {
		log.Fatalf("failed to open connection to db: %v", err)
		return nil, err
	}

	log.Println("connected to db")

	// Return the database connection, do not close it here
	return &database{db: db}, nil
}

func (db *database) Close() error {
	if err := db.db.Close(); err != nil {
		log.Fatal("failed to close db")
		return err
	}
	return nil
}

func (db *database) SetKey(key string, value string) error {
	if err := db.db.Set([]byte(key), []byte(value), nil); err != nil {
		log.Fatal("failed to set key")
		return err
	}

	return nil
}

func (db *database) GetKey(key string) (string, error) {
	value, closer, err := db.db.Get([]byte(key))
	if err != nil {
		log.Fatal("failed to get key")
		return "", err
	}
	defer closer.Close()

	return string(value), nil
}
