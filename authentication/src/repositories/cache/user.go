package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type user struct {
	client *redis.Client
	appId  string
}

type UserData interface {
	// StoreSecret into the user logged
	StoreSecret(ctx context.Context, userId, secret string, exp time.Duration) error

	// UserFromSecret check if the user is logged in or not
	UserFromSecret(key string) (string, error)
}

func (s *user) getKeyForSecret(secret string) string {
	return fmt.Sprintf("%s-userdata-%s", s.appId, secret)
}

func (s *user) StoreSecret(ctx context.Context, userId, secret string, exp time.Duration) error {
	key := s.getKeyForSecret(secret)

	pipe := s.client.TxPipeline()
	pipe.Set(ctx, key, userId, exp)

	_, err := pipe.Exec(ctx)
	return err
}

func (s *user) UserFromSecret(secret string) (string, error) {
	key := s.getKeyForSecret(secret)
	response := s.client.Get(context.Background(), key)
	return response.Val(), response.Err()
}
