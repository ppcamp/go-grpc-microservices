package auth

import (
	"context"
	"streamer/controllers/base"
	authS "streamer/services/user_login"

	empty "github.com/golang/protobuf/ptypes/empty"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuthService struct {
	UnsafeAuthServiceServer

	handler *base.Handler
}

func NewAuthService(handler *base.Handler) AuthServiceServer {
	return &AuthService{handler: handler}
}

func (s *AuthService) Login(ctx context.Context, in *LoginInput) (*AuthOutput, error) {
	pl := authS.UserLoginInput{User: in.User, Password: in.Password}

	response, err := base.Handle(ctx, pl,
		authS.NewUserLoginService(s.handler.Cache, s.handler.Signer),
	)
	if err != nil {
		return nil, err
	}

	return &AuthOutput{
		Token: response.Token,
		Exp:   timestamppb.New(response.Exp),
	}, nil
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
