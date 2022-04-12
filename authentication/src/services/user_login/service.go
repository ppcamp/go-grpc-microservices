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

	cache     cache.Auth
	signer    jwt.Jwt
	signerExp time.Duration
}

// NewUserLoginService creates a service that get user password, check it, and
// return a valid JWT token
func NewUserLoginService(
	repo cache.Auth,
	signer jwt.Jwt,
) base.IBaseBusiness[UserLoginInput, UserLoginOutput] {
	return &UserLoginService[UserLoginInput, UserLoginOutput]{
		cache:     repo,
		signer:    signer,
		signerExp: time.Second * 60 * 60 * 60,
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

	exp := time.Now().Add(u.signerExp)
	token, err := u.signer.Generate(&jwt.Session{}, u.signerExp)
	if err != nil {
		return nil, err
	}

	response := &UserLoginOutput{token, exp}

	if err = u.cache.SetSession(u.Context, in.User, token, exp); err != nil {
		return nil, err
	}

	return response, nil
}
