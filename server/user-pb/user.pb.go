// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: server/user-pb/user.proto

package user

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email           string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password        string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Name            string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Age             uint32                 `protobuf:"varint,4,opt,name=age,proto3" json:"age,omitempty"`
	Gender          uint32                 `protobuf:"varint,5,opt,name=gender,proto3" json:"gender,omitempty"`
	CountryOrigin   string                 `protobuf:"bytes,6,opt,name=country_origin,json=countryOrigin,proto3" json:"country_origin,omitempty"`
	ProfilePic      string                 `protobuf:"bytes,7,opt,name=profile_pic,json=profilePic,proto3" json:"profile_pic,omitempty"`
	LoginLengthTime uint32                 `protobuf:"varint,8,opt,name=login_length_time,json=loginLengthTime,proto3" json:"login_length_time,omitempty"`
	Validated       bool                   `protobuf:"varint,9,opt,name=validated,proto3" json:"validated,omitempty"`
	ValidationCode  string                 `protobuf:"bytes,10,opt,name=validation_code,json=validationCode,proto3" json:"validation_code,omitempty"`
	Admin           bool                   `protobuf:"varint,11,opt,name=admin,proto3" json:"admin,omitempty"`
	SuperAdmin      bool                   `protobuf:"varint,12,opt,name=super_admin,json=superAdmin,proto3" json:"super_admin,omitempty"`
	Code            string                 `protobuf:"bytes,13,opt,name=code,proto3" json:"code,omitempty"`
	CodeExpire      *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=code_expire,json=codeExpire,proto3" json:"code_expire,omitempty"`
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	mi := &file_server_user_pb_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_user_pb_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_server_user_pb_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserRequest) GetAge() uint32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *UserRequest) GetGender() uint32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *UserRequest) GetCountryOrigin() string {
	if x != nil {
		return x.CountryOrigin
	}
	return ""
}

func (x *UserRequest) GetProfilePic() string {
	if x != nil {
		return x.ProfilePic
	}
	return ""
}

func (x *UserRequest) GetLoginLengthTime() uint32 {
	if x != nil {
		return x.LoginLengthTime
	}
	return 0
}

func (x *UserRequest) GetValidated() bool {
	if x != nil {
		return x.Validated
	}
	return false
}

func (x *UserRequest) GetValidationCode() string {
	if x != nil {
		return x.ValidationCode
	}
	return ""
}

func (x *UserRequest) GetAdmin() bool {
	if x != nil {
		return x.Admin
	}
	return false
}

func (x *UserRequest) GetSuperAdmin() bool {
	if x != nil {
		return x.SuperAdmin
	}
	return false
}

func (x *UserRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UserRequest) GetCodeExpire() *timestamppb.Timestamp {
	if x != nil {
		return x.CodeExpire
	}
	return nil
}

type UserIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserIDRequest) Reset() {
	*x = UserIDRequest{}
	mi := &file_server_user_pb_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserIDRequest) ProtoMessage() {}

func (x *UserIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_user_pb_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserIDRequest.ProtoReflect.Descriptor instead.
func (*UserIDRequest) Descriptor() ([]byte, []int) {
	return file_server_user_pb_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserIDRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint32       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	User *UserRequest `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UpdateUserRequest) Reset() {
	*x = UpdateUserRequest{}
	mi := &file_server_user_pb_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRequest) ProtoMessage() {}

func (x *UpdateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_user_pb_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return file_server_user_pb_user_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateUserRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateUserRequest) GetUser() *UserRequest {
	if x != nil {
		return x.User
	}
	return nil
}

type UserIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserIDResponse) Reset() {
	*x = UserIDResponse{}
	mi := &file_server_user_pb_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserIDResponse) ProtoMessage() {}

func (x *UserIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_user_pb_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserIDResponse.ProtoReflect.Descriptor instead.
func (*UserIDResponse) Descriptor() ([]byte, []int) {
	return file_server_user_pb_user_proto_rawDescGZIP(), []int{3}
}

func (x *UserIDResponse) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ListUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*UserResponse `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *ListUsersResponse) Reset() {
	*x = ListUsersResponse{}
	mi := &file_server_user_pb_user_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUsersResponse) ProtoMessage() {}

func (x *ListUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_user_pb_user_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUsersResponse.ProtoReflect.Descriptor instead.
func (*ListUsersResponse) Descriptor() ([]byte, []int) {
	return file_server_user_pb_user_proto_rawDescGZIP(), []int{4}
}

func (x *ListUsersResponse) GetUsers() []*UserResponse {
	if x != nil {
		return x.Users
	}
	return nil
}

type UserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email           string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Name            string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age             uint32                 `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	Gender          uint32                 `protobuf:"varint,4,opt,name=gender,proto3" json:"gender,omitempty"`
	CountryOrigin   string                 `protobuf:"bytes,5,opt,name=country_origin,json=countryOrigin,proto3" json:"country_origin,omitempty"`
	ProfilePic      string                 `protobuf:"bytes,6,opt,name=profile_pic,json=profilePic,proto3" json:"profile_pic,omitempty"`
	LoginLengthTime uint32                 `protobuf:"varint,7,opt,name=login_length_time,json=loginLengthTime,proto3" json:"login_length_time,omitempty"`
	Validated       bool                   `protobuf:"varint,8,opt,name=validated,proto3" json:"validated,omitempty"`
	ValidationCode  string                 `protobuf:"bytes,9,opt,name=validation_code,json=validationCode,proto3" json:"validation_code,omitempty"`
	Admin           bool                   `protobuf:"varint,10,opt,name=admin,proto3" json:"admin,omitempty"`
	SuperAdmin      bool                   `protobuf:"varint,11,opt,name=super_admin,json=superAdmin,proto3" json:"super_admin,omitempty"`
	Code            string                 `protobuf:"bytes,12,opt,name=code,proto3" json:"code,omitempty"`
	CodeExpire      *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=code_expire,json=codeExpire,proto3" json:"code_expire,omitempty"`
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	mi := &file_server_user_pb_user_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_user_pb_user_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_server_user_pb_user_proto_rawDescGZIP(), []int{5}
}

func (x *UserResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserResponse) GetAge() uint32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *UserResponse) GetGender() uint32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *UserResponse) GetCountryOrigin() string {
	if x != nil {
		return x.CountryOrigin
	}
	return ""
}

func (x *UserResponse) GetProfilePic() string {
	if x != nil {
		return x.ProfilePic
	}
	return ""
}

func (x *UserResponse) GetLoginLengthTime() uint32 {
	if x != nil {
		return x.LoginLengthTime
	}
	return 0
}

func (x *UserResponse) GetValidated() bool {
	if x != nil {
		return x.Validated
	}
	return false
}

func (x *UserResponse) GetValidationCode() string {
	if x != nil {
		return x.ValidationCode
	}
	return ""
}

func (x *UserResponse) GetAdmin() bool {
	if x != nil {
		return x.Admin
	}
	return false
}

func (x *UserResponse) GetSuperAdmin() bool {
	if x != nil {
		return x.SuperAdmin
	}
	return false
}

func (x *UserResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UserResponse) GetCodeExpire() *timestamppb.Timestamp {
	if x != nil {
		return x.CodeExpire
	}
	return nil
}

var File_server_user_pb_user_proto protoreflect.FileDescriptor

var file_server_user_pb_user_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x70, 0x62,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc0, 0x03, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12,
	0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x70, 0x69, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x12, 0x2a, 0x0a, 0x11, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x27, 0x0a, 0x0f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x75, 0x70, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x73, 0x75, 0x70, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x6f, 0x64, 0x65, 0x45, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x22, 0x1f, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x4a, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22,
	0x20, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x3d, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x22, 0xa5, 0x03, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x67,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x1f, 0x0a, 0x0b,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x69, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x12, 0x2a, 0x0a,
	0x11, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x4c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x75, 0x70, 0x65, 0x72, 0x5f,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x73, 0x75, 0x70,
	0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x63,
	0x6f, 0x64, 0x65, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x6f,
	0x64, 0x65, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x32, 0x99, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x32, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x35, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x04, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x76, 0x61, 0x72, 0x6f, 0x74, 0x6f, 0x72, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2d, 0x67, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_server_user_pb_user_proto_rawDescOnce sync.Once
	file_server_user_pb_user_proto_rawDescData = file_server_user_pb_user_proto_rawDesc
)

func file_server_user_pb_user_proto_rawDescGZIP() []byte {
	file_server_user_pb_user_proto_rawDescOnce.Do(func() {
		file_server_user_pb_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_user_pb_user_proto_rawDescData)
	})
	return file_server_user_pb_user_proto_rawDescData
}

var file_server_user_pb_user_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_server_user_pb_user_proto_goTypes = []any{
	(*UserRequest)(nil),           // 0: user.UserRequest
	(*UserIDRequest)(nil),         // 1: user.UserIDRequest
	(*UpdateUserRequest)(nil),     // 2: user.UpdateUserRequest
	(*UserIDResponse)(nil),        // 3: user.UserIDResponse
	(*ListUsersResponse)(nil),     // 4: user.ListUsersResponse
	(*UserResponse)(nil),          // 5: user.UserResponse
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 7: google.protobuf.Empty
}
var file_server_user_pb_user_proto_depIdxs = []int32{
	6, // 0: user.UserRequest.code_expire:type_name -> google.protobuf.Timestamp
	0, // 1: user.UpdateUserRequest.user:type_name -> user.UserRequest
	5, // 2: user.ListUsersResponse.users:type_name -> user.UserResponse
	6, // 3: user.UserResponse.code_expire:type_name -> google.protobuf.Timestamp
	0, // 4: user.User.Create:input_type -> user.UserRequest
	1, // 5: user.User.Get:input_type -> user.UserIDRequest
	2, // 6: user.User.Update:input_type -> user.UpdateUserRequest
	1, // 7: user.User.Delete:input_type -> user.UserIDRequest
	7, // 8: user.User.List:input_type -> google.protobuf.Empty
	1, // 9: user.User.Create:output_type -> user.UserIDRequest
	5, // 10: user.User.Get:output_type -> user.UserResponse
	3, // 11: user.User.Update:output_type -> user.UserIDResponse
	3, // 12: user.User.Delete:output_type -> user.UserIDResponse
	4, // 13: user.User.List:output_type -> user.ListUsersResponse
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_server_user_pb_user_proto_init() }
func file_server_user_pb_user_proto_init() {
	if File_server_user_pb_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_server_user_pb_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_user_pb_user_proto_goTypes,
		DependencyIndexes: file_server_user_pb_user_proto_depIdxs,
		MessageInfos:      file_server_user_pb_user_proto_msgTypes,
	}.Build()
	File_server_user_pb_user_proto = out.File
	file_server_user_pb_user_proto_rawDesc = nil
	file_server_user_pb_user_proto_goTypes = nil
	file_server_user_pb_user_proto_depIdxs = nil
}
