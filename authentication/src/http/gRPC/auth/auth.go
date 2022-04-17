package auth

import (
	"context"
	"errors"
	"time"

	"authentication/helpers/handlers"
	"authentication/services/auth/check_token"
	"authentication/services/auth/invalidate_tokens"
	"authentication/services/auth/login"
	"authentication/services/auth/refresh_token"

	empty "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var ErrNotImplemented error = errors.New("not implemented")

type AuthService struct {
	UnsafeAuthServiceServer

	handler  *handlers.Handler
	tokenExp time.Duration
}

func NewAuthService(handler *handlers.Handler) AuthServiceServer {
	return &AuthService{handler: handler, tokenExp: time.Second * 60 * 60 * 60}
}

func (s *AuthService) Login(
	ctx context.Context, in *LoginInput) (*AuthOutput, error) {
	pl := login.Input{User: in.User, Password: in.Password}
	service := login.NewService(s.handler.Cache, s.tokenExp)

	response, err := handlers.Handle[login.Input, login.Output](ctx, pl, service)
	if err != nil {
		return nil, err
	}

	return &AuthOutput{
		Token: response.Token,
		Exp:   timestamppb.New(response.Exp),
	}, nil
}

func (s *AuthService) Refresh(
	ctx context.Context, in *TokenInput) (*AuthOutput, error) {
	pl := refresh_token.Input{Token: in.Token}
	service := refresh_token.NewService(s.handler.Cache, s.tokenExp)

	response, err := handlers.Handle(ctx, pl, service)
	if err != nil {
		return nil, err
	}

	return &AuthOutput{
		Token: response.Token,
		Exp:   timestamppb.New(response.Exp),
	}, nil
}

func (s *AuthService) InvalidateAll(
	ctx context.Context, in *SessionsInput) (*empty.Empty, error) {
	pl := invalidate_tokens.Input{User: in.User, Token: in.Token}
	service := invalidate_tokens.NewService(s.handler.Cache)

	_, err := handlers.Handle(ctx, pl, service)
	if err != nil {
		return nil, err
	}

	return new(empty.Empty), nil
}

func (s *AuthService) IsValid(
	ctx context.Context, in *TokenInput) (*empty.Empty, error) {
	pl := check_token.Input{Token: in.Token}
	service := check_token.NewService(s.handler.Cache)

	_, err := handlers.Handle(ctx, pl, service)
	if err != nil {
		return nil, err
	}

	return new(empty.Empty), nil
}

func (s *AuthService) ActiveSessions(
	context.Context, *SessionsInput) (*SessionsOutput, error) {
	return nil, ErrNotImplemented
}

func (s *AuthService) Invalidate(
	ctx context.Context, in *TokenInput) (*empty.Empty, error) {
	return new(empty.Empty), ErrNotImplemented
}
