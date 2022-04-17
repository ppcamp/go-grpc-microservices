package user

import (
	context "context"
	"streamer/helpers/handlers"
	"streamer/services/create_user"
	"streamer/services/delete_user"
	"streamer/services/update_user"

	empty "github.com/golang/protobuf/ptypes/empty"
)

type userService struct {
	UnsafeUserServiceServer

	handler *handlers.Handler
}

func NewUserService(handler *handlers.Handler) UserServiceServer {
	return &userService{handler: handler}
}

func (s *userService) Create(
	ctx context.Context,
	in *CreateUserInput,
) (*empty.Empty, error) {
	pl := create_user.Input{
		FirstName:  in.FirstName,
		MiddleName: in.MiddleName,
		LastName:   in.MiddleName,
		Nick:       in.Nickname,
		BirthDate:  in.Birthdate.AsTime(),
		Email:      in.Email,
		Password:   in.Password,
	}

	_, err := handlers.Handle[create_user.Input, create_user.Output](
		ctx,
		pl,
		create_user.NewService(s.handler.Auth),
	)

	return new(empty.Empty), err
}

func (s *userService) Update(
	ctx context.Context,
	in *UpdateUserInput,
) (*empty.Empty, error) {
	pl := update_user.Input{
		FirstName:  in.UserData.GetFirstName(),
		MiddleName: in.UserData.GetMiddleName(),
		LastName:   in.UserData.GetLastName(),
		Nick:       in.UserData.GetNickname(),
		BirthDate:  in.UserData.GetBirthdate().AsTime(),
		Email:      in.UserData.GetEmail(),
		Password:   in.UserData.GetPassword(),
	}

	_, err := handlers.Handle[update_user.Input, update_user.Output](
		ctx,
		pl,
		update_user.NewService(s.handler.Auth),
	)

	return new(empty.Empty), err
}

func (s *userService) Delete(
	ctx context.Context,
	in *DeleteUserInput,
) (*empty.Empty, error) {
	pl := delete_user.Input{Token: in.Token}

	_, err := handlers.Handle[delete_user.Input, delete_user.Output](
		ctx,
		pl,
		delete_user.NewService(s.handler.Auth),
	)

	return new(empty.Empty), err
}
