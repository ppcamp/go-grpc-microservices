package create_password

import "errors"

var (
	ErrUserAlreadyExist error = errors.New("user already exist and it's active")
)
