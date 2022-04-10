package auth

import (
	"context"
	"streamer/services"
	authService "streamer/services/auth"

	empty "github.com/golang/protobuf/ptypes/empty"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuthService struct {
	UnsafeAuthServiceServer

	handler *services.Handler
}

func NewAuthService(handler *services.Handler) AuthServiceServer {
	return &AuthService{handler: handler}
}

func (s *AuthService) Login(ctx context.Context, in *LoginInput) (*AuthOutput, error) {
	_, err := s.handler.Handle(
		ctx,
		nil,
		authService.NewUserLoginHandler(
			s.handler.Cache,
			s.handler.Signer,
		),
	)
	if err != nil {
		return nil, err
	}

	out := &AuthOutput{
		Token: "",                // response.Token,
		Exp:   timestamppb.Now(), //response.Exp,
	}

	return out, nil
}

func (s *AuthService) Refresh(context.Context, *TokenInput) (*AuthOutput, error) {
	return nil, nil
}

func (s *AuthService) Invalidate(context.Context, *TokenInput) (*empty.Empty, error) {
	return nil, nil
}

func (s *AuthService) IsValid(context.Context, *TokenInput) (*wrappers.BoolValue, error) {
	return nil, nil
}
