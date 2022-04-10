package database

import (
	"github.com/jmoiron/sqlx"
)

type Connection interface {
	// StartTransaction starts a transaction and assign it to every storage
	StartTransaction() (Storage, error)
}

type store struct {
	DB *sqlx.DB
}

// NewStore creates a connection with database and its responsible to assign
// the database to the storages
func NewStore(driverName string, dataSource string) (Connection, error) {
	db, err := sqlx.Connect(driverName, dataSource)
	if err != nil {
		return nil, err
	}
	return &store{db}, nil
}

func (s *store) StartTransaction() (Storage, error) {
	tx, err := s.DB.Beginx()
	if err != nil {
		return nil, err
	}
	return NewStorage(tx), nil
}
