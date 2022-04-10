package base

import (
	"context"
)

type BaseBusiness[In, Out any] interface {
	// Execute is the function that will be executed by the controller
	Execute(in In) (response *Out, err error)

	SetContext(ctx context.Context)
}

type BaseBusinessImpl struct {
	Context context.Context
}

func (bc *BaseBusinessImpl) SetContext(ctx context.Context) {
	bc.Context = ctx
}
