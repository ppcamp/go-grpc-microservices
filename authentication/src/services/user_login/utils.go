package user_login

import "golang.org/x/crypto/bcrypt"

func validatePassword(pswd, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pswd))
}
