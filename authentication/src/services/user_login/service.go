package user_login

import (
	"streamer/repositories/cache"
	database "streamer/repositories/database/user"
	"streamer/services/base"
	"streamer/utils/jwt"
	"time"
)

type UserLoginService[In, Out any] struct {
	base.TransactionBusiness[database.UserStorage]

	repository cache.Auth
	signer     jwt.Jwt
	signerExp  time.Duration
}

// NewUserLoginService creates a service that get user password, check it, and
// return a valid JWT token
func NewUserLoginService(
	repo cache.Auth,
	signer jwt.Jwt,
) base.IBaseBusiness[UserLoginInput, UserLoginOutput] {
	return &UserLoginService[UserLoginInput, UserLoginOutput]{
		repository: repo,
		signer:     signer,
		signerExp:  60 * 60 * 60,
	}
}

func (u *UserLoginService[In, Out]) Execute(in UserLoginInput) (*UserLoginOutput, error) {
	hash, err := u.Transaction.GetUserPassword(in.User)
	if err != nil {
		return nil, err
	}

	if err = validatePassword(in.Password, hash); err != nil {
		return nil, err
	}

	token, err := u.signer.Generate(&jwt.Session{}, u.signerExp)
	if err != nil {
		return nil, err
	}

	exp := time.Now().Add(time.Second * time.Duration(u.signerExp))

	return &UserLoginOutput{token, exp}, nil
}
