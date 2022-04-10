package database

import (
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	MAX_OPEN_CONNS     = 10
	MAX_IDLE_CONNS     = 1
	CONN_MAX_IDLE_TIME = time.Second * 2
)

type Connection interface {
	// StartTransaction starts a transaction and assign it to every storage
	StartTransaction() (Storage, error)

	Close() error
}

type store struct {
	DB *sqlx.DB
}

// NewStore creates a connection with database and its responsible to assign
// the database to the storages
func NewStore(dataSource string) (Connection, error) {
	db, err := sqlx.Connect("postgres", dataSource)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(MAX_OPEN_CONNS)
	db.SetMaxIdleConns(MAX_IDLE_CONNS)
	db.SetConnMaxIdleTime(CONN_MAX_IDLE_TIME)
	return &store{db}, nil
}

func (s *store) StartTransaction() (Storage, error) {
	tx, err := s.DB.Beginx()
	if err != nil {
		return nil, err
	}
	return NewStorage(tx), nil
}

func (s *store) Close() error {
	return s.DB.Close()
}
