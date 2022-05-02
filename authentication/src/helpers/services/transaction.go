package services

import (
	"database/sql/driver"

	"authentication/repositories/database"

	"github.com/sirupsen/logrus"
)

type ITransactionBusiness[In, Out any] interface {
	IBaseBusiness[In, Out]

	// Set the current sqlx transaction for the context
	SetTransaction(tr database.Storage)
}

type TransactionBusiness[T driver.Tx] struct {
	BaseBusiness

	Storage T
}

func (tc *TransactionBusiness[T]) SetTransaction(tr database.Storage) {
	ntr, ok := tr.(T)
	if !ok {
		logrus.Panic("fail to parse transaction")
	}
	tc.Storage = ntr
}
