package user_password

import (
	database "streamer/repositories/database/user"
	"streamer/services/base"
)

type UserCreatePassword[In, Out any] struct {
	base.TransactionBusiness[database.UserStorage]
}

// NewUserCreatePasswordService creates a service that get user password, check it, and
// return a valid JWT token
func NewUserCreatePasswordService() base.IBaseBusiness[UserLoginInput, UserLoginOutput] {
	return &UserCreatePassword[UserLoginInput, UserLoginOutput]{}
}

func (u *UserCreatePassword[In, Out]) Execute(in UserLoginInput) (*UserLoginOutput, error) {
	return nil, nil
}
