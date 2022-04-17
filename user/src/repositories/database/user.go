package database

import (
	"database/sql/driver"
	"time"

	"github.com/jmoiron/sqlx"
)

type UserDataEntity struct {
	FirstName  string    `db:"first_name"`
	MiddleName string    `db:"middle_name"`
	LastName   string    `db:"last_name"`
	Nick       string    `db:"nick"`
	BirthDate  time.Time `db:"birthdate"`
	Email      string    `db:"email"`
}

type UserStorage interface {
	driver.Tx
	Create(in UserDataEntity) (userId string, err error)
	Delete(userId string) error
}

type userTransaction struct{ *sqlx.Tx }

func NewUserStorage(tx *sqlx.Tx) *userTransaction { return &userTransaction{tx} }

func (s *userTransaction) Create(in UserDataEntity) (userId string, err error) {
	sql := `
	INSERT INTO users(first_name, middle_name, last_name, nick, birthdate, email)
	VALUES(:first_name, :middle_name, :last_name, :nick, :birthdate, :email)
	RETURNING id;
	`
	row, err := s.NamedQuery(sql, in)
	if err != nil {
		return "", err
	}
	err = row.Scan(&userId)
	return "", err
}

func (s *userTransaction) Delete(userId string) error {
	sql := `
	DELETE FROM users
	WHERE id = $1;
	`
	_, err := s.Query(sql, userId)
	return err
}
