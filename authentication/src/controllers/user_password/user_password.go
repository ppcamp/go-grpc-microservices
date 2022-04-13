package user_password

import (
	context "context"
	"streamer/controllers"
	"streamer/services/user/create_password"

	empty "github.com/golang/protobuf/ptypes/empty"
)

type UserPasswordService struct {
	UnsafeUserPasswordServiceServer

	handler *controllers.Handler
}

func NewUserPasswordService(handler *controllers.Handler) UserPasswordServiceServer {
	return &UserPasswordService{handler: handler}
}

func (u *UserPasswordService) Create(ctx context.Context, in *CreateInput) (*CreateOutput, error) {
	pl := create_password.Input{User: in.User, Password: in.Password}

	service := create_password.NewService(u.handler.Cache)
	response, err := controllers.
		Handle[create_password.Input, create_password.Output](ctx, pl, service)
	if err != nil {
		return &CreateOutput{}, err
	}

	return &CreateOutput{Secret: response.ActivateToken}, nil
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
