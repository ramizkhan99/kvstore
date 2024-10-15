package store

import "log"

var db *database

func Init() {
	var err error
	db, err = connectDB()
	if err != nil {
		log.Fatalf("failed to connect to db %v", err)
	}
}

func CloseDB() {
	db.Close()
}

func SetKey(key string, value string) error {
	return db.SetKey(key, value)
}

func GetKey(key string) (string, error) {
	return db.GetKey(key)
}
