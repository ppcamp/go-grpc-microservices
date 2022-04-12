package user_password

import (
	database "streamer/repositories/database/user"
	"streamer/services/base"
)

type UserCreatePassword[In, Out any] struct {
	base.TransactionBusiness[database.UserStorage]
}

// NewService creates a service that get user password, check it, and
// return a valid JWT token
func NewService() base.IBaseBusiness[Input, Output] {

	return &UserCreatePassword[Input, Output]{}
}

func (u *UserCreatePassword[In, Out]) Execute(in Input) (*Output, error) {
	return nil, nil
}
