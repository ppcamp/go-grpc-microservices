package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

const pingTime time.Duration = 1 * time.Second

type CacheConfig redis.Options

type Cache interface {
	Auth
	UserData
}

type cache struct {
	*auth
	*user
}

func NewCacheRepository(options CacheConfig, identifier string) (Cache, error) {
	ops := redis.Options(options)
	client := redis.NewClient(&ops)

	if resp := client.Ping(context.Background()); resp.Err() != nil {
		logrus.Warnf("redis connection failed %v", resp.Err())
		return nil, resp.Err()
	}

	return &cache{
		&auth{client, identifier},
		&user{client, identifier},
	}, nil
}
