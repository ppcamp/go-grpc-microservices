package login

import (
	"time"

	"authentication/helpers/services"
	"authentication/repositories/cache"
	"authentication/repositories/database"
	"authentication/utils/jwt"
)

type LoginService[In, Out any] struct {
	services.TransactionBusiness[database.AuthStorage]

	Storage   database.UserStorage
	cache     cache.Auth
	signerExp time.Duration
}

// NewService creates a service that get user password, check it, and
// return a valid JWT token
func NewService(
	repo cache.Auth,
	exp time.Duration,
) services.ITransactionBusiness[Input, Output] {
	return &LoginService[Input, Output]{
		cache:     repo,
		signerExp: exp,
	}
}

func (s *LoginService[In, Out]) Execute(in Input) (*Output, error) {
	hash, err := s.Storage.GetUserPassword(in.User)
	if err != nil {
		return nil, err
	}

	if err = validatePassword(in.Password, hash); err != nil {
		return nil, err
	}

	exp := time.Now().Add(s.signerExp)
	token, err := jwt.Signer.Generate(&jwt.Session{}, s.signerExp)
	if err != nil {
		return nil, err
	}

	response := &Output{token, exp}

	if err = s.cache.Authorize(s.Context, in.User, token, exp); err != nil {
		return nil, err
	}

	return response, nil
}
