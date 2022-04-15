package update_password

import (
	"streamer/helpers/services"
	"streamer/repositories/cache"
	"streamer/repositories/database"
	"streamer/utils/jwt"
)

type UpdateLoggedService struct {
	services.TransactionBusiness[database.UserStorage]

	cache cache.UserData
}

// NewService creates a service that get the current logged user, check if is
// valid, and update its password
func NewService(cache cache.UserData) services.ITransactionBusiness[Input, Output] {
	return &UpdateLoggedService{cache: cache}
}

func (s *UpdateLoggedService) Execute(in Input) (*Output, error) {
	session, err := jwt.Signer.Session(in.JwtToken)
	if err != nil {
		return nil, err
	}

	if err = s.Storage.UpdatePassword(*session.UserId, in.Password); err != nil {
		return nil, err
	}

	return &Output{}, nil
}
