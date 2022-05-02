package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

type CacheConfig redis.Options

type Cache interface {
}

type cache struct {
}

func NewCacheRepository(options CacheConfig, identifier string) (Cache, error) {
	ops := redis.Options(options)
	client := redis.NewClient(&ops)

	if resp := client.Ping(context.Background()); resp.Err() != nil {
		logrus.Warnf("redis connection failed %v", resp.Err())
		return nil, resp.Err()
	}

	return &cache{}, nil
}
