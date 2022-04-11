package user

import (
	"github.com/jmoiron/sqlx"
)

type UserTransaction struct {
	*sqlx.Tx
}

type UserStorage interface {
	GetUserPassword(userId string) (string, error)
	CreateUserPassword(userId string, hashedPassword string) error
}

func NewTransaction(tx *sqlx.Tx) *UserTransaction {
	return &UserTransaction{tx}
}
