package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type auth struct {
	client *redis.Client
	appId  string
}

type internalSession struct {
	Exp time.Time `json:"exp"`
}

type Auth interface {
	// SetSession get the objects from redis, remove those who had expired and
	// update the expiration for the user
	SetSession(ctx context.Context, userId, jwtToken string, exp time.Time) error
	GetSession(key string) error
}

func (s *auth) getKeyForUser(userId string) string {
	return fmt.Sprintf("%s-auth-%s", s.appId, userId)
}

func (s *auth) removeExpTokensAndPipe(
	ctx context.Context,
	key string,
) (redis.Pipeliner, error) {
	tokens, err := s.client.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var session internalSession
	pipe := s.client.TxPipeline()

	for field, value := range tokens {
		err := json.Unmarshal([]byte(value), &session)
		if err != nil {
			return nil, err
		}

		if session.Exp.Before(time.Now()) {
			pipe.HDel(ctx, key, field)
		}
	}

	return pipe, nil
}

func (s *auth) SetSession(
	ctx context.Context,
	userId, jwtToken string,
	exp time.Time,
) error {
	key := s.getKeyForUser(userId)

	session := internalSession{exp}

	pipe, err := s.removeExpTokensAndPipe(ctx, key)
	if err != nil {
		return err
	}

	marshaledSession, err := json.Marshal(session)
	if err != nil {
		return err
	}

	pipe.HSet(ctx, key, jwtToken, marshaledSession)
	pipe.ExpireAt(ctx, key, exp)

	_, err = pipe.Exec(ctx)
	return err
}

func (s *auth) GetSession(key string) error {
	return nil
}
