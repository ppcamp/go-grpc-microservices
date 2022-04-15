package auth

import (
	"context"
	"errors"
	"streamer/helpers/handlers"
	"streamer/services/auth/check_token"
	"streamer/services/auth/invalidate_tokens"
	"streamer/services/auth/login"
	"streamer/services/auth/refresh_token"
	"time"

	empty "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
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

func (s *AuthService) Login(ctx context.Context, in *LoginInput) (*AuthOutput, error) {
	pl := login.Input{User: in.User, Password: in.Password}

	response, err := handlers.Handle[login.Input, login.Output](
		ctx, pl, login.NewService(s.handler.Cache, s.tokenExp),
	)
	if err != nil {
		return nil, err
	}

	return &AuthOutput{
		Token: response.Token,
		Exp:   timestamppb.New(response.Exp),
	}, nil
}

func (s *AuthService) Refresh(ctx context.Context, in *TokenInput) (*AuthOutput, error) {
	pl := refresh_token.Input{Token: in.Token}

	response, err := handlers.Handle(
		ctx, pl, refresh_token.NewService(s.handler.Cache, s.tokenExp),
	)
	if err != nil {
		return nil, err
	}

	return &AuthOutput{
		Token: response.Token,
		Exp:   timestamppb.New(response.Exp),
	}, nil
}

func (s *AuthService) Invalidate(ctx context.Context, in *TokenInput) (*empty.Empty, error) {
	return nil, ErrNotImplemented
}

func (s *AuthService) InvalidateAll(ctx context.Context, in *SessionsInput) (*empty.Empty, error) {
	pl := invalidate_tokens.Input{User: in.User, Token: in.Token}

	_, err := handlers.Handle(
		ctx, pl, invalidate_tokens.NewService(s.handler.Cache),
	)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *AuthService) IsValid(ctx context.Context, in *TokenInput) (*empty.Empty, error) {
	pl := check_token.Input{Token: in.Token}

	_, err := handlers.Handle(ctx, pl, check_token.NewService(s.handler.Cache))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *AuthService) ActiveSessions(context.Context, *SessionsInput) (*SessionsOutput, error) {
	return nil, ErrNotImplemented
}
