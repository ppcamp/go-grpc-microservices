package user_password

import (
	context "context"
	"streamer/services"

	empty "github.com/golang/protobuf/ptypes/empty"
)

type UserPasswordService struct {
	UnsafeUserPasswordServiceServer

	handler *services.Handler
}

func NewUserPasswordService(handler *services.Handler) UserPasswordServiceServer {
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
