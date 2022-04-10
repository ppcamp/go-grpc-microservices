package user_login

import (
	"streamer/repositories/cache"
	"streamer/services/base"
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

func (u *UserLogin) Execute(_ any) (*any, error) {
	_, err := u.signer.Generate(&jwt.Session{}, u.signerExp)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func NewUserLoginHandler(
	repo cache.Auth,
	signer jwt.Jwt,
) base.BaseBusiness {
	return &UserLogin{
		repository: repo,
		signer:     signer,
		signerExp:  60 * 60 * 60,
	}
}
