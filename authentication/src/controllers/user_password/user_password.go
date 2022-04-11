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

func (u *UserPasswordService) Create(context.Context, *CreatePasswordInput) (*empty.Empty, error) {
	return nil, nil
}

func (u *UserPasswordService) Recover(context.Context, *RecoverInput) (*RecoverPayload, error) {
	return nil, nil
}

func (u *UserPasswordService) Update(context.Context, *UpdatePasswordInput) (*empty.Empty, error) {
	return nil, nil
}
