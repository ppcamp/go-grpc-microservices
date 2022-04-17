package update_password

import (
	"authentication/helpers/services"
	"authentication/repositories/cache"
	"authentication/repositories/database"
)

type UpdatePasswordService struct {
	services.TransactionBusiness[database.UserStorage]

	cache cache.UserData
}

// NewService creates a service that get user password, check it, and
// return a valid JWT token
func NewService(cache cache.UserData) services.ITransactionBusiness[Input, Output] {
	return &UpdatePasswordService{cache: cache}
}

func (s *UpdatePasswordService) Execute(in Input) (*Output, error) {
	user, err := s.cache.UserFromSecret(in.RecoverToken)
	if err != nil {
		return nil, err
	}

	err = s.Storage.UpdatePassword(user, in.Password)
	return new(Output), err
}
