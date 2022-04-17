package auth

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IUserPassword interface {
	CreatePassword(ctx context.Context, in *CreateInput) (string, error)
	UpdatePassword(ctx context.Context, in *UpdateInput) error
	DeleteUser(ctx context.Context, in *DeleteInput) error
}

func NewUserPasswordClient() (IUserPassword, error) {
	opt := grpc.WithTransportCredentials(insecure.NewCredentials())

	conn, err := grpc.Dial("", opt)
	if err != nil {
		return nil, err
	}

	client := NewUserPasswordServiceClient(conn)
	return &userPassword{client: client}, nil
}

type userPassword struct{ client UserPasswordServiceClient }

func (s *userPassword) CreatePassword(
	ctx context.Context,
	in *CreateInput,
) (string, error) {
	out, err := s.client.Create(ctx, in)
	if err != nil {
		return "", err
	}
	return out.GetSecret(), nil
}

func (s *userPassword) UpdatePassword(
	ctx context.Context,
	in *UpdateInput,
) error {
	_, err := s.client.Update(ctx, in)
	return err
}

func (s *userPassword) DeleteUser(
	ctx context.Context,
	in *DeleteInput,
) error {
	_, err := s.client.Delete(ctx, in)
	return err
}
