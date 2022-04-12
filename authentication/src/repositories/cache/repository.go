package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

const pingTime time.Duration = 1

type CacheConfig redis.Options

type Cache interface {
	Auth
}

type cache struct {
	*auth
}

func NewCacheRepository(options CacheConfig, identifier string) (Cache, error) {
	ops := redis.Options(options)
	client := redis.NewClient(&ops)

	timedCtx, cancel := context.WithTimeout(context.Background(), pingTime)
	defer cancel()

	if resp := client.Ping(timedCtx); resp.Err() != nil {
		logrus.Warn("redis connection failed %v", resp.Err())
		return nil, resp.Err()
	}

	return &cache{
		&auth{client, identifier},
	}, nil
}
