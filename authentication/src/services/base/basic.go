package base

import (
	"context"
)

type BaseBusiness interface {
	// Execute is the function that will be executed by the controller
	Execute(in any) (response *any, err error)

	SetContext(ctx context.Context)
}

type BaseBusinessImpl struct {
	Context context.Context
}

func (bc *BaseBusinessImpl) SetContext(ctx context.Context) {
	bc.Context = ctx
}
