package store

import "sync"

var cache sync.Map

func GetKeyFromCache(key string) (string, bool) {
	if value, ok := cache.Load(key); ok {
		return value.(string), true
	}
	return "", false
}

func SetKeyToCache(key string, value string) {
	cache.Store(key, value)
}

func DeleteKeyFromCache(key string) {
	cache.Delete(key)
}
