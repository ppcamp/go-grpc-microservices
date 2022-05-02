package database

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
	ActivateUser(userId string) error
	CreateUserPassword(userId string, hashedPassword string) error
	UpdatePassword(userId, password string) error
}

func NewUserTransaction(tx *sqlx.Tx) *UserTransaction {
	return &UserTransaction{tx}
}

func (u *UserTransaction) CreateUserPassword(userId string, hashedPassword string) error {
	sql := `
	INSERT INTO passwords(user_id, user_password)
	VALUES($1, $2);
	`
	_, err := u.Query(sql, userId, hashedPassword)
	return err
}

func (u *UserTransaction) ActivateUser(userId string) error {
	sql := `
	UPDATE passwords
	SET activate = true
	WHERE user_id = $1;
	`
	_, err := u.Query(sql, userId)
	return err
}

func (u *UserTransaction) UpdatePassword(userId, password string) error {
	sql := `
	UPDATE passwords
	SET user_password = $2
	WHERE user_id = $1;
	`
	_, err := u.Query(sql, userId, password)
	return err
}
