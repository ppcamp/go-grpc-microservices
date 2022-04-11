package user_password

import (
	"time"
)

type UserLoginInput struct {
	User     string
	Password string
}

type UserLoginOutput struct {
	Token string
	Exp   time.Time
}
