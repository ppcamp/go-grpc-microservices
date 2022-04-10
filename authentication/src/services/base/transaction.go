package base

type BaseTransactionBusiness[In, Out any] interface {
	BaseBusiness[In, Out]

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
