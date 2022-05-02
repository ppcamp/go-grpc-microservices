package user_password

import (
	context "context"

	"authentication/helpers/handlers"
	"authentication/services/user/activate_account"
	"authentication/services/user/create_password"
	"authentication/services/user/delete_account"
	"authentication/services/user/recover_password"
	"authentication/services/user/update_logged"
	"authentication/services/user/update_password"

	empty "github.com/golang/protobuf/ptypes/empty"
)

type UserPasswordService struct {
	UnsafeUserPasswordServiceServer

	handler *handlers.Handler
}

func NewUserPasswordService(
	handler *handlers.Handler) UserPasswordServiceServer {
	return &UserPasswordService{handler: handler}
}

func (u *UserPasswordService) Create(
	ctx context.Context, in *CreateInput) (*CreateOutput, error) {
	pl := create_password.Input{User: in.GetUser(), Password: in.GetPassword()}

	service := create_password.NewService(u.handler.Cache)
	response, err := handlers.
		Handle[create_password.Input, create_password.Output](ctx, pl, service)
	if err != nil {
		return nil, err
	}

	return &CreateOutput{Secret: response.ActivateToken}, nil
}

func (u *UserPasswordService) Activate(
	ctx context.Context, in *ActivateInput) (*empty.Empty, error) {
	pl := activate_account.Input{Secret: in.GetSecret()}
	service := activate_account.NewService(u.handler.Cache)

	_, err := handlers.
		Handle[activate_account.Input, activate_account.Output](ctx, pl, service)
	if err != nil {
		return nil, err
	}
	return new(empty.Empty), nil
}

func (u *UserPasswordService) Recover(
	ctx context.Context, in *RecoverInput) (*RecoverOutput, error) {
	pl := recover_password.Input{Email: in.GetEmail()}
	service := recover_password.NewService(u.handler.Cache)

	response, err := handlers.
		Handle[recover_password.Input, recover_password.Output](ctx, pl, service)
	if err != nil {
		return nil, err
	}
	return &RecoverOutput{Secret: response.Secret}, nil
}

func (u *UserPasswordService) Update(
	ctx context.Context, in *UpdateInput) (*empty.Empty, error) {
	pl := update_password.Input{
		RecoverToken: in.GetSecret(),
		Password:     in.GetPassword(),
	}
	service := update_password.NewService(u.handler.Cache)

	_, err := handlers.
		Handle[update_password.Input, update_password.Output](ctx, pl, service)
	if err != nil {
		return nil, err
	}
	return new(empty.Empty), nil
}

func (u *UserPasswordService) UpdateByToken(
	ctx context.Context, in *UpdateInput) (*empty.Empty, error) {
	pl := update_logged.Input{
		JwtToken: in.GetSecret(),
		Password: in.GetPassword(),
	}
	service := update_logged.NewService(u.handler.Cache)

	_, err := handlers.
		Handle[update_logged.Input, update_logged.Output](ctx, pl, service)
	if err != nil {
		return nil, err
	}
	return new(empty.Empty), nil
}

func (u *UserPasswordService) Delete(
	ctx context.Context, in *DeleteInput) (*empty.Empty, error) {
	pl := delete_account.Input{Token: in.GetToken()}
	service := delete_account.NewService()

	_, err := handlers.
		Handle[delete_account.Input, delete_account.Output](ctx, pl, service)
	if err != nil {
		return nil, err
	}
	return new(empty.Empty), nil
}
