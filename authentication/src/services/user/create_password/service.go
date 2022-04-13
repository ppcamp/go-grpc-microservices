package create_password

import (
	"database/sql"
	"streamer/repositories/cache"
	"streamer/repositories/database/user"
	"streamer/services"
	"streamer/utils"
	"streamer/utils/strings"
	"time"
)

type UserCreatePassword struct {
	cache cache.UserData
	exp   time.Duration

	services.TransactionBusiness[user.UserStorage]
}

// NewService creates a service that get user password, check it, and
// return a valid JWT token
func NewService(cache cache.UserData) services.ITransactionBusiness[Input, Output] {
	return &UserCreatePassword{cache: cache, exp: time.Hour * 24}
}

func (u *UserCreatePassword) Execute(in Input) (*Output, error) {
	_, err := u.Storage.GetUserPassword(in.User)
	if err == nil {
		return nil, ErrUserAlreadyExist
	} else if err != sql.ErrNoRows {
		return nil, err
	}

	hashedPassword, err := utils.HashPassword(in.Password)
	if err != nil {
		return nil, err
	}

	err = u.Storage.CreateUserPassword(in.User, hashedPassword)
	if err != nil {
		return nil, err
	}

	unlockSecret := strings.RandString(30)

	err = u.cache.SetSecret(u.Context, in.User, unlockSecret, u.exp)
	if err != nil {
		return nil, err
	}

	return &Output{unlockSecret}, nil
}
