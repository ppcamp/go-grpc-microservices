package user

import (
	"github.com/jmoiron/sqlx"
)

type UserTransaction struct {
	*sqlx.Tx
}

type UserStorage interface {
	GetUser()
}

func NewTransaction(tx *sqlx.Tx) *UserTransaction {
	return &UserTransaction{tx}
}
