package update_user

import (
	"streamer/helpers/services"
	"streamer/microservices/auth"
	"streamer/repositories/database"
)

type updateUserService struct {
	services.TransactionBusiness[database.UserStorage]

	auth auth.IUserPassword
}

func NewService(auth auth.IUserPassword) services.ITransactionBusiness[Input, Output] {
	return &updateUserService{auth: auth}
}

func (s *updateUserService) Execute(in Input) (*Output, error) {
	return &Output{}, nil
}
