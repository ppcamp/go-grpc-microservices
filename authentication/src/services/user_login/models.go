package user_login

import (
	"time"
)

type Input struct {
	User     string
	Password string
}

type Output struct {
	Token string
	Exp   time.Time
}
