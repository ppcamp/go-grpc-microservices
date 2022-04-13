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
	SetSecret(ctx context.Context, userId, secret string, exp time.Duration) error
	GetSecret(key string) error
}

func (s *user) getKeyForUser(userId string) string {
	return fmt.Sprintf("%s-activate-%s", s.appId, userId)
}

func (s *user) SetSecret(ctx context.Context, userId, secret string, exp time.Duration) error {
	key := s.getKeyForUser(userId)

	pipe := s.client.TxPipeline()
	pipe.Set(ctx, key, secret, exp)

	_, err := pipe.Exec(ctx)
	return err
}

func (s *user) GetSecret(key string) error {
	return nil
}
