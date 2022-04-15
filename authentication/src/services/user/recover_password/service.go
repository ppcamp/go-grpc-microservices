package recover_password

import (
	"streamer/helpers/services"
	"streamer/repositories/cache"
	"streamer/repositories/database"
	"time"

	"github.com/ppcamp/go-lib/random"
)

type RecoverPasswordService struct {
	services.TransactionBusiness[database.UserStorage]

	cache cache.UserData
	exp   time.Duration
}

// NewService creates a service that generates a token to recover some user
// password
func NewService(cache cache.UserData) services.ITransactionBusiness[Input, Output] {
	return &RecoverPasswordService{cache: cache}
}

func (s *RecoverPasswordService) Execute(in Input) (*Output, error) {
	secret := random.String(30)

	err := s.cache.StoreSecret(s.Context, in.User, secret, s.exp)
	if err != nil {
		return nil, err
	}

	return &Output{secret}, err
}
