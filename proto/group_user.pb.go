// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: group_user.proto

package grpc_board

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GroupUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	GroupId string `protobuf:"bytes,2,opt,name=groupId,proto3" json:"groupId,omitempty"`
	UserId  string `protobuf:"bytes,3,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GroupUser) Reset() {
	*x = GroupUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupUser) ProtoMessage() {}

func (x *GroupUser) ProtoReflect() protoreflect.Message {
	mi := &file_group_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupUser.ProtoReflect.Descriptor instead.
func (*GroupUser) Descriptor() ([]byte, []int) {
	return file_group_user_proto_rawDescGZIP(), []int{0}
}

func (x *GroupUser) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GroupUser) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *GroupUser) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetGroupUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetGroupUserRequest) Reset() {
	*x = GetGroupUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupUserRequest) ProtoMessage() {}

func (x *GetGroupUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_group_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupUserRequest.ProtoReflect.Descriptor instead.
func (*GetGroupUserRequest) Descriptor() ([]byte, []int) {
	return file_group_user_proto_rawDescGZIP(), []int{1}
}

func (x *GetGroupUserRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetGroupUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupUser *GroupUser `protobuf:"bytes,1,opt,name=groupUser,proto3" json:"groupUser,omitempty"`
}

func (x *GetGroupUserResponse) Reset() {
	*x = GetGroupUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupUserResponse) ProtoMessage() {}

func (x *GetGroupUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_group_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupUserResponse.ProtoReflect.Descriptor instead.
func (*GetGroupUserResponse) Descriptor() ([]byte, []int) {
	return file_group_user_proto_rawDescGZIP(), []int{2}
}

func (x *GetGroupUserResponse) GetGroupUser() *GroupUser {
	if x != nil {
		return x.GroupUser
	}
	return nil
}

type CreateGroupUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupUser *GroupUser `protobuf:"bytes,1,opt,name=groupUser,proto3" json:"groupUser,omitempty"`
}

func (x *CreateGroupUserRequest) Reset() {
	*x = CreateGroupUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroupUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupUserRequest) ProtoMessage() {}

func (x *CreateGroupUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_group_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupUserRequest.ProtoReflect.Descriptor instead.
func (*CreateGroupUserRequest) Descriptor() ([]byte, []int) {
	return file_group_user_proto_rawDescGZIP(), []int{3}
}

func (x *CreateGroupUserRequest) GetGroupUser() *GroupUser {
	if x != nil {
		return x.GroupUser
	}
	return nil
}

type CreateGroupUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CreateGroupUserResponse) Reset() {
	*x = CreateGroupUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroupUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupUserResponse) ProtoMessage() {}

func (x *CreateGroupUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_group_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupUserResponse.ProtoReflect.Descriptor instead.
func (*CreateGroupUserResponse) Descriptor() ([]byte, []int) {
	return file_group_user_proto_rawDescGZIP(), []int{4}
}

func (x *CreateGroupUserResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

type UpdateGroupUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupUser *GroupUser `protobuf:"bytes,1,opt,name=groupUser,proto3" json:"groupUser,omitempty"`
}

func (x *UpdateGroupUserRequest) Reset() {
	*x = UpdateGroupUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGroupUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGroupUserRequest) ProtoMessage() {}

func (x *UpdateGroupUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_group_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGroupUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateGroupUserRequest) Descriptor() ([]byte, []int) {
	return file_group_user_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateGroupUserRequest) GetGroupUser() *GroupUser {
	if x != nil {
		return x.GroupUser
	}
	return nil
}

type UpdateGroupUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *UpdateGroupUserResponse) Reset() {
	*x = UpdateGroupUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGroupUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGroupUserResponse) ProtoMessage() {}

func (x *UpdateGroupUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_group_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGroupUserResponse.ProtoReflect.Descriptor instead.
func (*UpdateGroupUserResponse) Descriptor() ([]byte, []int) {
	return file_group_user_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateGroupUserResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

type DeleteGroupUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteGroupUserRequest) Reset() {
	*x = DeleteGroupUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGroupUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGroupUserRequest) ProtoMessage() {}

func (x *DeleteGroupUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_group_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGroupUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteGroupUserRequest) Descriptor() ([]byte, []int) {
	return file_group_user_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteGroupUserRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteGroupUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *DeleteGroupUserResponse) Reset() {
	*x = DeleteGroupUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGroupUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGroupUserResponse) ProtoMessage() {}

func (x *DeleteGroupUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_group_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGroupUserResponse.ProtoReflect.Descriptor instead.
func (*DeleteGroupUserResponse) Descriptor() ([]byte, []int) {
	return file_group_user_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteGroupUserResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

type GroupUsersInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupUsers []*GroupUser `protobuf:"bytes,1,rep,name=groupUsers,proto3" json:"groupUsers,omitempty"`
}

func (x *GroupUsersInfo) Reset() {
	*x = GroupUsersInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupUsersInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupUsersInfo) ProtoMessage() {}

func (x *GroupUsersInfo) ProtoReflect() protoreflect.Message {
	mi := &file_group_user_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupUsersInfo.ProtoReflect.Descriptor instead.
func (*GroupUsersInfo) Descriptor() ([]byte, []int) {
	return file_group_user_proto_rawDescGZIP(), []int{9}
}

func (x *GroupUsersInfo) GetGroupUsers() []*GroupUser {
	if x != nil {
		return x.GroupUsers
	}
	return nil
}

var File_group_user_proto protoreflect.FileDescriptor

var file_group_user_proto_rawDesc = []byte{
	0x0a, 0x10, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x09, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x25, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x4b, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x09, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x55, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x73, 0x65, 0x72, 0x22, 0x4d,
	0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x55, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x73, 0x65, 0x72, 0x22, 0x31, 0x0a,
	0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0x4d, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x09, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x55, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x73, 0x65, 0x72, 0x22,
	0x31, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x28, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x31, 0x0a, 0x17,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x47, 0x0a, 0x0e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x73, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x35, 0x0a, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x73, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x73, 0x65, 0x72, 0x52, 0x0a, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x55, 0x73, 0x65, 0x72, 0x73, 0x32, 0xc3, 0x03, 0x0a, 0x10, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x53, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x0f, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x73, 0x65, 0x72, 0x12, 0x22, 0x2e,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x73, 0x65, 0x72, 0x12, 0x22, 0x2e, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x5a, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x55, 0x73, 0x65, 0x72, 0x12, 0x22, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1e,
	0x5a, 0x1c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4a, 0x49, 0x45,
	0x45, 0x45, 0x4e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_group_user_proto_rawDescOnce sync.Once
	file_group_user_proto_rawDescData = file_group_user_proto_rawDesc
)

func file_group_user_proto_rawDescGZIP() []byte {
	file_group_user_proto_rawDescOnce.Do(func() {
		file_group_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_group_user_proto_rawDescData)
	})
	return file_group_user_proto_rawDescData
}

var file_group_user_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_group_user_proto_goTypes = []interface{}{
	(*GroupUser)(nil),               // 0: group_user.GroupUser
	(*GetGroupUserRequest)(nil),     // 1: group_user.GetGroupUserRequest
	(*GetGroupUserResponse)(nil),    // 2: group_user.GetGroupUserResponse
	(*CreateGroupUserRequest)(nil),  // 3: group_user.CreateGroupUserRequest
	(*CreateGroupUserResponse)(nil), // 4: group_user.CreateGroupUserResponse
	(*UpdateGroupUserRequest)(nil),  // 5: group_user.UpdateGroupUserRequest
	(*UpdateGroupUserResponse)(nil), // 6: group_user.UpdateGroupUserResponse
	(*DeleteGroupUserRequest)(nil),  // 7: group_user.DeleteGroupUserRequest
	(*DeleteGroupUserResponse)(nil), // 8: group_user.DeleteGroupUserResponse
	(*GroupUsersInfo)(nil),          // 9: group_user.GroupUsersInfo
	(*emptypb.Empty)(nil),           // 10: google.protobuf.Empty
}
var file_group_user_proto_depIdxs = []int32{
	0,  // 0: group_user.GetGroupUserResponse.groupUser:type_name -> group_user.GroupUser
	0,  // 1: group_user.CreateGroupUserRequest.groupUser:type_name -> group_user.GroupUser
	0,  // 2: group_user.UpdateGroupUserRequest.groupUser:type_name -> group_user.GroupUser
	0,  // 3: group_user.GroupUsersInfo.groupUsers:type_name -> group_user.GroupUser
	10, // 4: group_user.GroupUserService.GetAllGroupUsers:input_type -> google.protobuf.Empty
	1,  // 5: group_user.GroupUserService.GetGroupUserId:input_type -> group_user.GetGroupUserRequest
	3,  // 6: group_user.GroupUserService.CreateGroupUser:input_type -> group_user.CreateGroupUserRequest
	5,  // 7: group_user.GroupUserService.UpdateGroupUser:input_type -> group_user.UpdateGroupUserRequest
	7,  // 8: group_user.GroupUserService.DeleteGroupUser:input_type -> group_user.DeleteGroupUserRequest
	9,  // 9: group_user.GroupUserService.GetAllGroupUsers:output_type -> group_user.GroupUsersInfo
	2,  // 10: group_user.GroupUserService.GetGroupUserId:output_type -> group_user.GetGroupUserResponse
	4,  // 11: group_user.GroupUserService.CreateGroupUser:output_type -> group_user.CreateGroupUserResponse
	6,  // 12: group_user.GroupUserService.UpdateGroupUser:output_type -> group_user.UpdateGroupUserResponse
	8,  // 13: group_user.GroupUserService.DeleteGroupUser:output_type -> group_user.DeleteGroupUserResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_group_user_proto_init() }
func file_group_user_proto_init() {
	if File_group_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_group_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupUser); i {
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
		file_group_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupUserRequest); i {
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
		file_group_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupUserResponse); i {
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
		file_group_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGroupUserRequest); i {
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
		file_group_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGroupUserResponse); i {
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
		file_group_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGroupUserRequest); i {
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
		file_group_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGroupUserResponse); i {
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
		file_group_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGroupUserRequest); i {
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
		file_group_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGroupUserResponse); i {
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
		file_group_user_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupUsersInfo); i {
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
			RawDescriptor: file_group_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_group_user_proto_goTypes,
		DependencyIndexes: file_group_user_proto_depIdxs,
		MessageInfos:      file_group_user_proto_msgTypes,
	}.Build()
	File_group_user_proto = out.File
	file_group_user_proto_rawDesc = nil
	file_group_user_proto_goTypes = nil
	file_group_user_proto_depIdxs = nil
}
