package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
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
	logrus.Info("getting key")
	key := s.getKeyForSecret(secret)

	logrus.Info("creating pipeline")
	pipe := s.client.TxPipeline()
	pipe.Set(ctx, key, userId, exp)

	logrus.Info("exec pipeline")
	_, err := pipe.Exec(ctx)
	logrus.WithField("werr", err).Info("finish")
	return err
}

func (s *user) UserFromSecret(secret string) (string, error) {
	key := s.getKeyForSecret(secret)
	response := s.client.Get(context.Background(), key)
	return response.Val(), response.Err()
}
