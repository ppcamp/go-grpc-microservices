package services

import "context"

type IBaseBusiness[In, Out any] interface {
	// Execute is the function that will be executed by the controller
	Execute(in In) (response *Out, err error)

	SetContext(ctx context.Context)
}

type BaseBusiness struct {
	Context context.Context
}

func (bc *BaseBusiness) SetContext(ctx context.Context) {
	bc.Context = ctx
}
