package delete_user

import (
	"streamer/helpers/services"
	"streamer/microservices/auth"
	"streamer/repositories/database"
)

type deleteUserService struct {
	services.TransactionBusiness[database.UserStorage]

	auth auth.IUserPassword
}

func NewService(auth auth.IUserPassword) services.ITransactionBusiness[Input, Output] {
	return &deleteUserService{auth: auth}
}

func (s *deleteUserService) Execute(in Input) (*Output, error) {
	pl := &auth.DeleteInput{Token: in.Token}
	if err := s.auth.DeleteUser(s.Context, pl); err != nil {
		return nil, err
	}

	err := s.Storage.Delete(in.Token)
	return new(Output), err
}
