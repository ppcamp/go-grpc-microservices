package database

import (
	"github.com/jmoiron/sqlx"
)

type Storage interface {
	UserStorage
}

type storage struct {
	*userTransaction
}

// NewStorage starts a new object that can be used to access all internal
// storages
//
// NOTE
//
// If you
func NewStorage(tx *sqlx.Tx) Storage {
	return &storage{NewUserStorage(tx)}
}
