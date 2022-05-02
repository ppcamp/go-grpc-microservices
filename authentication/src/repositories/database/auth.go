package database

import (
	"database/sql/driver"

	"github.com/jmoiron/sqlx"
)

type AuthTransaction struct {
	*sqlx.Tx
}

type AuthStorage interface {
	driver.Tx
	GetUserPassword(userId string) (string, error)
}

func NewAuthTransaction(tx *sqlx.Tx) *AuthTransaction {
	return &AuthTransaction{tx}
}

func (u *AuthTransaction) GetUserPassword(userId string) (string, error) {
	var password string
	var err error

	sql := `
	SELECT user_password
	FROM passwords
	WHERE user_id = $1
		AND active = TRUE
	ORDER BY created_at DESC
	LIMIT 1
	`
	err = u.Get(&password, sql, userId)
	return password, err
}
