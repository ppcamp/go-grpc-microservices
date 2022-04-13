package login

import (
	"streamer/repositories/cache"
	"streamer/repositories/database/user"
	"streamer/services"
	"streamer/utils/jwt"
	"time"
)

type UserLoginService[In, Out any] struct {
	services.TransactionBusiness[user.UserStorage]

	Storage   user.UserStorage
	cache     cache.Auth
	signer    jwt.Jwt
	signerExp time.Duration
}

// NewService creates a service that get user password, check it, and
// return a valid JWT token
func NewService(
	repo cache.Auth,
	signer jwt.Jwt,
) services.ITransactionBusiness[Input, Output] {
	return &UserLoginService[Input, Output]{
		cache:     repo,
		signer:    signer,
		signerExp: time.Second * 60 * 60 * 60,
	}
}

func (u *UserLoginService[In, Out]) Execute(in Input) (*Output, error) {
	hash, err := u.Storage.GetUserPassword(in.User)
	if err != nil {
		return nil, err
	}

	if err = validatePassword(in.Password, hash); err != nil {
		return nil, err
	}

	exp := time.Now().Add(u.signerExp)
	token, err := u.signer.Generate(&jwt.Session{}, u.signerExp)
	if err != nil {
		return nil, err
	}

	response := &Output{token, exp}

	if err = u.cache.SetSession(u.Context, in.User, token, exp); err != nil {
		return nil, err
	}

	return response, nil
}
