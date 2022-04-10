package database

import (
	"database/sql/driver"
	"streamer/repositories/database/user"

	"github.com/jmoiron/sqlx"
)

type Storage interface {
	driver.Tx
	user.UserStorage
}

type storage struct {
	*sqlx.Tx
	*user.UserTransaction
}

// NewStorage starts a new object that can be used to access all internal
// storages
//
// NOTE
//
// If you
func NewStorage(tx *sqlx.Tx) Storage {
	return &storage{tx, user.NewTransaction(tx)}
}
