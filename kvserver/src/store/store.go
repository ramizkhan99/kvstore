package store

import (
	"fmt"
	"log"
)

var db *database

const (
	maxWidthOfKey = 16
)

func StartDB(port uint8) {
	var err error
	db, err = connectDB(port)
	if err != nil {
		log.Fatalf("failed to connect to db %v", err)
	}
}

func CloseDB() {
	db.Close()
}

func SetKey(key string, value string) error {
	if len(key) > maxWidthOfKey {
		return fmt.Errorf("key %q too long, max width is %d", key, maxWidthOfKey)
	}
	SetKeyToCache(key, value)
	return db.SetKey(key, value)
}

func GetKey(key string) (string, error) {
	value, ok := GetKeyFromCache(key)
	if ok {
		log.Println("key found in cache")
		return value, nil
	}

	if len(key) > maxWidthOfKey {
		return "", fmt.Errorf("key %q too long, max width is %d", key, maxWidthOfKey)
	}
	value, err := db.GetKey(key)
	if err != nil {
		return "", err
	}
	SetKeyToCache(key, value)
	return value, nil
}
