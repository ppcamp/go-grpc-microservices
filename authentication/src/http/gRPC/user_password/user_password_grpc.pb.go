// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package user_password

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserPasswordServiceClient is the client API for UserPasswordService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserPasswordServiceClient interface {
	// Create a new (disabled) user and return a temp secret to activate it'
	Create(ctx context.Context, in *CreateInput, opts ...grpc.CallOption) (*CreateOutput, error)
	// Use the secret from Create to activate the user
	Activate(ctx context.Context, in *ActivateInput, opts ...grpc.CallOption) (*empty.Empty, error)
	// Create a temp secret that allows user to update its password
	Recover(ctx context.Context, in *RecoverInput, opts ...grpc.CallOption) (*RecoverOutput, error)
	// Use the secret got from Recover to update some user password
	Update(ctx context.Context, in *UpdateInput, opts ...grpc.CallOption) (*empty.Empty, error)
	// Use the token to update the current user password
	UpdateByToken(ctx context.Context, in *UpdateInput, opts ...grpc.CallOption) (*empty.Empty, error)
	// Delete the current user
	Delete(ctx context.Context, in *DeleteInput, opts ...grpc.CallOption) (*empty.Empty, error)
}

type userPasswordServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserPasswordServiceClient(cc grpc.ClientConnInterface) UserPasswordServiceClient {
	return &userPasswordServiceClient{cc}
}

func (c *userPasswordServiceClient) Create(ctx context.Context, in *CreateInput, opts ...grpc.CallOption) (*CreateOutput, error) {
	out := new(CreateOutput)
	err := c.cc.Invoke(ctx, "/UserPasswordService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userPasswordServiceClient) Activate(ctx context.Context, in *ActivateInput, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/UserPasswordService/Activate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userPasswordServiceClient) Recover(ctx context.Context, in *RecoverInput, opts ...grpc.CallOption) (*RecoverOutput, error) {
	out := new(RecoverOutput)
	err := c.cc.Invoke(ctx, "/UserPasswordService/Recover", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userPasswordServiceClient) Update(ctx context.Context, in *UpdateInput, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/UserPasswordService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userPasswordServiceClient) UpdateByToken(ctx context.Context, in *UpdateInput, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/UserPasswordService/UpdateByToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userPasswordServiceClient) Delete(ctx context.Context, in *DeleteInput, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/UserPasswordService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserPasswordServiceServer is the server API for UserPasswordService service.
// All implementations must embed UnimplementedUserPasswordServiceServer
// for forward compatibility
type UserPasswordServiceServer interface {
	// Create a new (disabled) user and return a temp secret to activate it'
	Create(context.Context, *CreateInput) (*CreateOutput, error)
	// Use the secret from Create to activate the user
	Activate(context.Context, *ActivateInput) (*empty.Empty, error)
	// Create a temp secret that allows user to update its password
	Recover(context.Context, *RecoverInput) (*RecoverOutput, error)
	// Use the secret got from Recover to update some user password
	Update(context.Context, *UpdateInput) (*empty.Empty, error)
	// Use the token to update the current user password
	UpdateByToken(context.Context, *UpdateInput) (*empty.Empty, error)
	// Delete the current user
	Delete(context.Context, *DeleteInput) (*empty.Empty, error)
	mustEmbedUnimplementedUserPasswordServiceServer()
}

// UnimplementedUserPasswordServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserPasswordServiceServer struct {
}

func (UnimplementedUserPasswordServiceServer) Create(context.Context, *CreateInput) (*CreateOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedUserPasswordServiceServer) Activate(context.Context, *ActivateInput) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Activate not implemented")
}
func (UnimplementedUserPasswordServiceServer) Recover(context.Context, *RecoverInput) (*RecoverOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Recover not implemented")
}
func (UnimplementedUserPasswordServiceServer) Update(context.Context, *UpdateInput) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedUserPasswordServiceServer) UpdateByToken(context.Context, *UpdateInput) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateByToken not implemented")
}
func (UnimplementedUserPasswordServiceServer) Delete(context.Context, *DeleteInput) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedUserPasswordServiceServer) mustEmbedUnimplementedUserPasswordServiceServer() {}

// UnsafeUserPasswordServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserPasswordServiceServer will
// result in compilation errors.
type UnsafeUserPasswordServiceServer interface {
	mustEmbedUnimplementedUserPasswordServiceServer()
}

func RegisterUserPasswordServiceServer(s grpc.ServiceRegistrar, srv UserPasswordServiceServer) {
	s.RegisterService(&UserPasswordService_ServiceDesc, srv)
}

func _UserPasswordService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserPasswordServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserPasswordService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserPasswordServiceServer).Create(ctx, req.(*CreateInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserPasswordService_Activate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivateInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserPasswordServiceServer).Activate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserPasswordService/Activate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserPasswordServiceServer).Activate(ctx, req.(*ActivateInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserPasswordService_Recover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecoverInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserPasswordServiceServer).Recover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserPasswordService/Recover",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserPasswordServiceServer).Recover(ctx, req.(*RecoverInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserPasswordService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserPasswordServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserPasswordService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserPasswordServiceServer).Update(ctx, req.(*UpdateInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserPasswordService_UpdateByToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserPasswordServiceServer).UpdateByToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserPasswordService/UpdateByToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserPasswordServiceServer).UpdateByToken(ctx, req.(*UpdateInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserPasswordService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserPasswordServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserPasswordService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserPasswordServiceServer).Delete(ctx, req.(*DeleteInput))
	}
	return interceptor(ctx, in, info, handler)
}

// UserPasswordService_ServiceDesc is the grpc.ServiceDesc for UserPasswordService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserPasswordService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UserPasswordService",
	HandlerType: (*UserPasswordServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _UserPasswordService_Create_Handler,
		},
		{
			MethodName: "Activate",
			Handler:    _UserPasswordService_Activate_Handler,
		},
		{
			MethodName: "Recover",
			Handler:    _UserPasswordService_Recover_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserPasswordService_Update_Handler,
		},
		{
			MethodName: "UpdateByToken",
			Handler:    _UserPasswordService_UpdateByToken_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserPasswordService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/user_password.proto",
}
