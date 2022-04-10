package auth

import (
	"context"

	empty "github.com/golang/protobuf/ptypes/empty"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
)

type AuthService struct {
	UnimplementedAuthServiceServer
}

func NewAuthService() AuthServiceServer {
	return &AuthService{}
}

func (s *AuthService) Login(ctx context.Context, in *LoginInput) (*AuthOutput, error) {

	return nil, nil
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
