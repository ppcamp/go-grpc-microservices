package base

import "streamer/repositories/database"

type ITransactionBusiness[In, Out any] interface {
	IBaseBusiness[In, Out]

	// Set the current sqlx transaction for the context
	SetTransaction(tr database.Storage)
}

type TransactionBusiness[T any] struct {
	BaseBusiness

	Transaction T
}

func (tc *TransactionBusiness[T]) SetTransaction(tr T) {
	tc.Transaction = tr
}
