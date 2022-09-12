package caching

import (
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/eko/gocache/v3/cache"
	"github.com/eko/gocache/v3/store"
)

var cacheManager *cache.Cache[[]byte]

func InitCacheManager() error {
	bigcacheClient, err := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	bigcacheStore := store.NewBigcache(bigcacheClient)
	cacheManager = cache.New[[]byte](bigcacheStore)

	return err
}

func GetCacheManager() *cache.Cache[[]byte] {
	return cacheManager
}
