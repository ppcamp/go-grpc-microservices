package database

import (
	"github.com/jmoiron/sqlx"
)

type Storage interface {
	UserStorage
	AuthStorage
}

type storage struct {
	*sqlx.Tx
	*UserTransaction
	*AuthTransaction
}

// NewStorage starts a new object that can be used to access all internal
// storages
//
// NOTE
//
// If you
func NewStorage(tx *sqlx.Tx) Storage {
	return &storage{tx, NewUserTransaction(tx), NewAuthTransaction(tx)}
}
