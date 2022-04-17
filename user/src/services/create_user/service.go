package create_user

import (
	"streamer/helpers/services"
	"streamer/microservices/auth"
	"streamer/repositories/database"

	"github.com/sirupsen/logrus"
)

type CreateUserService struct {
	services.TransactionBusiness[database.UserStorage]

	auth auth.IUserPassword
}

func NewService(auth auth.IUserPassword) services.ITransactionBusiness[Input, Output] {
	return &CreateUserService{auth: auth}
}

func (s *CreateUserService) Execute(in Input) (*Output, error) {
	payload := database.UserDataEntity{
		FirstName:  in.FirstName,
		MiddleName: in.MiddleName,
		LastName:   in.LastName,
		Nick:       in.Nick,
		BirthDate:  in.BirthDate,
		Email:      in.Email,
	}

	userId, err := s.Storage.Create(payload)
	if err != nil {
		return nil, err
	}

	authPl := &auth.CreateInput{
		User:     userId,
		Password: in.Password,
	}

	token, err := s.auth.CreatePassword(s.Context, authPl)
	if err != nil {
		return nil, err
	}

	logrus.WithField("token", token).Debug("Send email to user with this token")

	return new(Output), nil
}
