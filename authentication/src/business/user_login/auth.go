package user_login

import (
	"streamer/business/base"
	"streamer/repositories/cache"
	"streamer/utils/jwt"
)

type UserLoginInput struct{}

type UserLoginOutput struct {
	Token string `json:"token"`
}

type UserLogin struct {
	base.BaseTransactionBusinessImpl

	repository cache.Auth
	signer     jwt.Jwt
	signerExp  int64
}

func (u *UserLogin) Execute(_ UserLoginInput) (*UserLoginOutput, error) {
	token, err := u.signer.Generate(&jwt.Session{}, u.signerExp)
	if err != nil {
		return nil, err
	}

	return &UserLoginOutput{
		Token: token,
	}, nil
}

func NewUserLoginHandler(
	repo cache.Auth,
	signer jwt.Jwt,
) base.BaseBusiness[UserLoginInput, UserLoginOutput] {
	return &UserLogin{
		repository: repo,
		signer:     signer,
		signerExp:  60 * 60 * 60,
	}
}
