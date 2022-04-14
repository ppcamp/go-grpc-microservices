package user

import (
	"database/sql/driver"

	"github.com/jmoiron/sqlx"
)

type UserTransaction struct {
	*sqlx.Tx
}

type UserStorage interface {
	driver.Tx
	GetUserPassword(userId string) (string, error)
	CreateUserPassword(userId string, hashedPassword string) error
}

func NewTransaction(tx *sqlx.Tx) *UserTransaction {
	return &UserTransaction{tx}
}
