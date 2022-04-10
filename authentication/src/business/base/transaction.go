package base

type BaseTransactionBusiness[Input any, Output any] interface {
	BaseBusiness[Input, Output]

	// Set the current sqlx transaction for the context
	SetTransaction(tr any)
}

type BaseTransactionBusinessImpl struct {
	BaseBusinessImpl

	Transaction any
}

func (tc *BaseTransactionBusinessImpl) SetTransaction(tr any) {
	tc.Transaction = tr
}
