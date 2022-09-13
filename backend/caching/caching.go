package caching

import (
	"context"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/eko/gocache/v3/cache"
	"github.com/eko/gocache/v3/metrics"
	"github.com/eko/gocache/v3/store"
	"github.com/vmihailenco/msgpack"
)

var metricCache *cache.MetricCache[[]byte]

func InitCacheManager() error {
	bigcacheClient, err := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	bigcacheStore := store.NewBigcache(bigcacheClient)

	promMetrics := metrics.NewPrometheus("xedule-proxy")
	metricCache = cache.NewMetric[[]byte](promMetrics, cache.New[[]byte](bigcacheStore))

	return err
}

func MarshalCache(key string, object any, options ...store.Option) error {
	bytes, err := msgpack.Marshal(object)
	if err != nil {
		return err
	}

	return metricCache.Set(context.Background(), key, bytes, options...)
}

func UnmarshalCache(key string, object any) error {
	bytes, err := metricCache.Get(context.Background(), key)
	if err != nil {
		return err
	}

	msgpack.Unmarshal(bytes, &object)
	return nil
}
