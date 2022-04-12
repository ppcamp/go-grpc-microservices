// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: protos/user_password.proto

package user_password

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User     string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`         // user's id
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"` // user's password
}

func (x *CreateInput) Reset() {
	*x = CreateInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_password_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInput) ProtoMessage() {}

func (x *CreateInput) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_password_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateInput.ProtoReflect.Descriptor instead.
func (*CreateInput) Descriptor() ([]byte, []int) {
	return file_protos_user_password_proto_rawDescGZIP(), []int{0}
}

func (x *CreateInput) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *CreateInput) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type RecoverInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"` // user's id
}

func (x *RecoverInput) Reset() {
	*x = RecoverInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_password_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecoverInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecoverInput) ProtoMessage() {}

func (x *RecoverInput) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_password_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecoverInput.ProtoReflect.Descriptor instead.
func (*RecoverInput) Descriptor() ([]byte, []int) {
	return file_protos_user_password_proto_rawDescGZIP(), []int{1}
}

func (x *RecoverInput) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

type RecoverOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Secret string `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"` // some random secret sent to user
}

func (x *RecoverOutput) Reset() {
	*x = RecoverOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_password_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecoverOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecoverOutput) ProtoMessage() {}

func (x *RecoverOutput) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_password_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecoverOutput.ProtoReflect.Descriptor instead.
func (*RecoverOutput) Descriptor() ([]byte, []int) {
	return file_protos_user_password_proto_rawDescGZIP(), []int{2}
}

func (x *RecoverOutput) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

type CreateOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Secret string `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"` // some random secret sent to user
}

func (x *CreateOutput) Reset() {
	*x = CreateOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_password_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOutput) ProtoMessage() {}

func (x *CreateOutput) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_password_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOutput.ProtoReflect.Descriptor instead.
func (*CreateOutput) Descriptor() ([]byte, []int) {
	return file_protos_user_password_proto_rawDescGZIP(), []int{3}
}

func (x *CreateOutput) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

type ActivateInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Secret string `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"` // the secret
}

func (x *ActivateInput) Reset() {
	*x = ActivateInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_password_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivateInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivateInput) ProtoMessage() {}

func (x *ActivateInput) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_password_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivateInput.ProtoReflect.Descriptor instead.
func (*ActivateInput) Descriptor() ([]byte, []int) {
	return file_protos_user_password_proto_rawDescGZIP(), []int{4}
}

func (x *ActivateInput) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

type UpdateInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The scret can be the same secret from RecoverOutput (Recover) or a valid
	// Token(UpdateByToken)
	Secret   string `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"` // the new password
}

func (x *UpdateInput) Reset() {
	*x = UpdateInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_password_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateInput) ProtoMessage() {}

func (x *UpdateInput) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_password_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateInput.ProtoReflect.Descriptor instead.
func (*UpdateInput) Descriptor() ([]byte, []int) {
	return file_protos_user_password_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateInput) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *UpdateInput) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_protos_user_password_proto protoreflect.FileDescriptor

var file_protos_user_password_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x0b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x27, 0x0a, 0x0d,
	0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x26, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x27, 0x0a,
	0x0d, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x41, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x32, 0x8b, 0x02, 0x0a, 0x13, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x27, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0c, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x0d, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x08, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x2a, 0x0a, 0x07, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x0d, 0x2e, 0x52, 0x65,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x0e, 0x2e, 0x52, 0x65, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x06,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0c, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x37,
	0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x0c, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x20, 0x5a, 0x1e, 0x2f, 0x73, 0x72, 0x63, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_protos_user_password_proto_rawDescOnce sync.Once
	file_protos_user_password_proto_rawDescData = file_protos_user_password_proto_rawDesc
)

func file_protos_user_password_proto_rawDescGZIP() []byte {
	file_protos_user_password_proto_rawDescOnce.Do(func() {
		file_protos_user_password_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_user_password_proto_rawDescData)
	})
	return file_protos_user_password_proto_rawDescData
}

var file_protos_user_password_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protos_user_password_proto_goTypes = []interface{}{
	(*CreateInput)(nil),   // 0: CreateInput
	(*RecoverInput)(nil),  // 1: RecoverInput
	(*RecoverOutput)(nil), // 2: RecoverOutput
	(*CreateOutput)(nil),  // 3: CreateOutput
	(*ActivateInput)(nil), // 4: ActivateInput
	(*UpdateInput)(nil),   // 5: UpdateInput
	(*empty.Empty)(nil),   // 6: google.protobuf.Empty
}
var file_protos_user_password_proto_depIdxs = []int32{
	0, // 0: UserPasswordService.Create:input_type -> CreateInput
	4, // 1: UserPasswordService.Activate:input_type -> ActivateInput
	1, // 2: UserPasswordService.Recover:input_type -> RecoverInput
	5, // 3: UserPasswordService.Update:input_type -> UpdateInput
	5, // 4: UserPasswordService.UpdateByToken:input_type -> UpdateInput
	3, // 5: UserPasswordService.Create:output_type -> CreateOutput
	6, // 6: UserPasswordService.Activate:output_type -> google.protobuf.Empty
	2, // 7: UserPasswordService.Recover:output_type -> RecoverOutput
	6, // 8: UserPasswordService.Update:output_type -> google.protobuf.Empty
	6, // 9: UserPasswordService.UpdateByToken:output_type -> google.protobuf.Empty
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_user_password_proto_init() }
func file_protos_user_password_proto_init() {
	if File_protos_user_password_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_user_password_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateInput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_user_password_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecoverInput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_user_password_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecoverOutput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_user_password_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOutput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_user_password_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivateInput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_user_password_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateInput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_user_password_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_user_password_proto_goTypes,
		DependencyIndexes: file_protos_user_password_proto_depIdxs,
		MessageInfos:      file_protos_user_password_proto_msgTypes,
	}.Build()
	File_protos_user_password_proto = out.File
	file_protos_user_password_proto_rawDesc = nil
	file_protos_user_password_proto_goTypes = nil
	file_protos_user_password_proto_depIdxs = nil
}
