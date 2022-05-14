package create_password

import (
	"database/sql"
	"time"

	"authentication/helpers/services"
	"authentication/repositories/cache"
	"authentication/repositories/database"
	"authentication/utils"

	"github.com/ppcamp/go-lib/random"
	"github.com/sirupsen/logrus"
)

type CreatePasswordService struct {
	services.TransactionBusiness[database.UserStorage]

	cache cache.UserData
	exp   time.Duration
}

// NewService creates a service that creates a new login
func NewService(cache cache.UserData) services.ITransactionBusiness[Input, Output] {
	return &CreatePasswordService{cache: cache, exp: time.Hour * 24}
}

func (s *CreatePasswordService) Execute(in Input) (*Output, error) {
	logrus.Info("getting password")
	_, err := s.Storage.GetUserPassword(in.User)
	if err == nil {
		return nil, ErrUserAlreadyExist
	} else if err != sql.ErrNoRows {
		return nil, err
	}

	logrus.Info("hashing password")
	hashedPassword, err := utils.HashPassword(in.Password)
	if err != nil {
		return nil, err
	}

	logrus.Info("storing password at database")
	err = s.Storage.CreateUserPassword(in.User, hashedPassword)
	if err != nil {
		return nil, err
	}

	unlockSecret := random.String(30)

	logrus.Info("storing password at cache")
	err = s.cache.StoreSecret(s.Context, in.User, unlockSecret, s.exp)
	if err != nil {
		return nil, err
	}

	logrus.Info("returning")
	return &Output{unlockSecret}, nil
}
