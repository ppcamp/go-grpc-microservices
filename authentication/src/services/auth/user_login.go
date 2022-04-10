package user_login

import (
	"streamer/repositories/cache"
	"streamer/services/base"
	"streamer/utils/jwt"
	"time"
)

type UserLoginInput struct {
	User     string
	Password string
}

type UserLoginOutput struct {
	Token string
	Exp   time.Time
}

type UserLogin[In, Out any] struct {
	base.BaseTransactionBusinessImpl

	repository cache.Auth
	signer     jwt.Jwt
	signerExp  int64
}

func (u *UserLogin[In, Out]) Execute(in UserLoginInput) (*UserLoginOutput, error) {
	_, err := u.signer.Generate(&jwt.Session{}, u.signerExp)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func NewUserLoginHandler(
	repo cache.Auth,
	signer jwt.Jwt,
) base.BaseBusiness[UserLoginInput, UserLoginOutput] {
	return &UserLogin[UserLoginInput, UserLoginOutput]{
		repository: repo,
		signer:     signer,
		signerExp:  60 * 60 * 60,
	}
}
