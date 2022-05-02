package activate_account

import (
	"authentication/helpers/services"
	"authentication/repositories/cache"
	"authentication/repositories/database"
)

type ActivateAccountService struct {
	services.TransactionBusiness[database.UserStorage]

	cache cache.UserData
}

// NewService creates an UserActivateAccout service, which is responsible to
// validate some secret and turn the user on
func NewService(cache cache.UserData) services.ITransactionBusiness[Input, Output] {
	return &ActivateAccountService{cache: cache}
}

func (s *ActivateAccountService) Execute(in Input) (*Output, error) {
	user, err := s.cache.UserFromSecret(in.Secret)
	if err != nil {
		return nil, err
	}

	err = s.Storage.ActivateUser(user)
	return new(Output), err
}
