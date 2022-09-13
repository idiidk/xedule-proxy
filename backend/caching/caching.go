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

func MarshalCache(ctx context.Context, key, object any, options ...store.Option) error {
	bytes, err := msgpack.Marshal(object)
	if err != nil {
		return err
	}

	return metricCache.Set(ctx, key, bytes, options...)
}

func UnmarshalCache(ctx context.Context, key, obj interface{}) error {
	data, err := metricCache.Get(ctx, key)
	if err != nil {
		return err
	}

	msgpack.Unmarshal(data, obj)
	return nil
}
