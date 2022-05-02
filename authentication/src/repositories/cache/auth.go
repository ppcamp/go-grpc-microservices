package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var ErrSessionExpired error = errors.New("session has expired")

type auth struct {
	client *redis.Client
	appId  string
}

type internalSession struct {
	Exp time.Time `json:"exp"`
}

type Auth interface {
	// Authorize get the objects from redis, remove those who had expired and
	// update the expiration for the user.
	//
	// Note
	//
	// This approach is used to get control of multiple logged devices. With this
	// approach, we are able to invalidate all sessions for the user
	Authorize(ctx context.Context, userId, jwtToken string, exp time.Time) error
	// Valid check if the user is still logged
	Valid(ctx context.Context, userId, key string) error
	// InvalidateAll tokens for a given user id
	InvalidateAll(ctx context.Context, userId string) error
	// Invalidate some token for a given user
	Invalidate(ctx context.Context, userId, token string) error
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

func (s *auth) Authorize(
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

func (s *auth) Valid(ctx context.Context, userId, key string) error {
	k := s.getKeyForUser(userId)

	r := s.client.HGet(ctx, k, key)
	if r.Err() != nil {
		return r.Err()
	}

	var session internalSession
	if err := json.Unmarshal([]byte(r.Val()), &session); err != nil {
		return err
	}

	if session.Exp.Before(time.Now()) {
		return ErrSessionExpired
	}

	return nil
}

func (s *auth) InvalidateAll(ctx context.Context, userId string) error {
	key := s.getKeyForUser(userId)
	r := s.client.Del(ctx, key)
	return r.Err()
}

func (s *auth) Invalidate(ctx context.Context, userId, token string) error {
	key := s.getKeyForUser(userId)
	r := s.client.HDel(ctx, key, token)
	return r.Err()
}
