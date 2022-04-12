package user_password

import (
	context "context"
	"streamer/controllers/base"

	empty "github.com/golang/protobuf/ptypes/empty"
)

type UserPasswordService struct {
	UnsafeUserPasswordServiceServer

	handler *base.Handler
}

func NewUserPasswordService(handler *base.Handler) UserPasswordServiceServer {
	return &UserPasswordService{handler: handler}
}

func (u *UserPasswordService) Create(context.Context, *CreateInput) (*CreateOutput, error) {
	return nil, nil
}

func (u *UserPasswordService) Activate(context.Context, *ActivateInput) (*empty.Empty, error) {
	return nil, nil
}

func (u *UserPasswordService) Recover(context.Context, *RecoverInput) (*RecoverOutput, error) {
	return nil, nil
}

func (u *UserPasswordService) Update(context.Context, *UpdateInput) (*empty.Empty, error) {
	return nil, nil
}

func (u *UserPasswordService) UpdateByToken(context.Context, *UpdateInput) (*empty.Empty, error) {
	return nil, nil
}
