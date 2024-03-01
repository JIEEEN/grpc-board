// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: board_group.proto

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

type BoardGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *BoardGroup) Reset() {
	*x = BoardGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_board_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoardGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoardGroup) ProtoMessage() {}

func (x *BoardGroup) ProtoReflect() protoreflect.Message {
	mi := &file_board_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoardGroup.ProtoReflect.Descriptor instead.
func (*BoardGroup) Descriptor() ([]byte, []int) {
	return file_board_group_proto_rawDescGZIP(), []int{0}
}

func (x *BoardGroup) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BoardGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetBoardGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetBoardGroupRequest) Reset() {
	*x = GetBoardGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_board_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBoardGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBoardGroupRequest) ProtoMessage() {}

func (x *GetBoardGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_board_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBoardGroupRequest.ProtoReflect.Descriptor instead.
func (*GetBoardGroupRequest) Descriptor() ([]byte, []int) {
	return file_board_group_proto_rawDescGZIP(), []int{1}
}

func (x *GetBoardGroupRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetBoardGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BoardGroup *BoardGroup `protobuf:"bytes,1,opt,name=boardGroup,proto3" json:"boardGroup,omitempty"`
}

func (x *GetBoardGroupResponse) Reset() {
	*x = GetBoardGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_board_group_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBoardGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBoardGroupResponse) ProtoMessage() {}

func (x *GetBoardGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_board_group_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBoardGroupResponse.ProtoReflect.Descriptor instead.
func (*GetBoardGroupResponse) Descriptor() ([]byte, []int) {
	return file_board_group_proto_rawDescGZIP(), []int{2}
}

func (x *GetBoardGroupResponse) GetBoardGroup() *BoardGroup {
	if x != nil {
		return x.BoardGroup
	}
	return nil
}

type CreateBoardGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BoardGroup *BoardGroup `protobuf:"bytes,1,opt,name=boardGroup,proto3" json:"boardGroup,omitempty"`
}

func (x *CreateBoardGroupRequest) Reset() {
	*x = CreateBoardGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_board_group_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBoardGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBoardGroupRequest) ProtoMessage() {}

func (x *CreateBoardGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_board_group_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBoardGroupRequest.ProtoReflect.Descriptor instead.
func (*CreateBoardGroupRequest) Descriptor() ([]byte, []int) {
	return file_board_group_proto_rawDescGZIP(), []int{3}
}

func (x *CreateBoardGroupRequest) GetBoardGroup() *BoardGroup {
	if x != nil {
		return x.BoardGroup
	}
	return nil
}

type CreateBoardGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CreateBoardGroupResponse) Reset() {
	*x = CreateBoardGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_board_group_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBoardGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBoardGroupResponse) ProtoMessage() {}

func (x *CreateBoardGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_board_group_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBoardGroupResponse.ProtoReflect.Descriptor instead.
func (*CreateBoardGroupResponse) Descriptor() ([]byte, []int) {
	return file_board_group_proto_rawDescGZIP(), []int{4}
}

func (x *CreateBoardGroupResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

type UpdateBoardGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BoardGroup *BoardGroup `protobuf:"bytes,1,opt,name=boardGroup,proto3" json:"boardGroup,omitempty"`
}

func (x *UpdateBoardGroupRequest) Reset() {
	*x = UpdateBoardGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_board_group_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBoardGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBoardGroupRequest) ProtoMessage() {}

func (x *UpdateBoardGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_board_group_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBoardGroupRequest.ProtoReflect.Descriptor instead.
func (*UpdateBoardGroupRequest) Descriptor() ([]byte, []int) {
	return file_board_group_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateBoardGroupRequest) GetBoardGroup() *BoardGroup {
	if x != nil {
		return x.BoardGroup
	}
	return nil
}

type UpdateBoardGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *UpdateBoardGroupResponse) Reset() {
	*x = UpdateBoardGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_board_group_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBoardGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBoardGroupResponse) ProtoMessage() {}

func (x *UpdateBoardGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_board_group_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBoardGroupResponse.ProtoReflect.Descriptor instead.
func (*UpdateBoardGroupResponse) Descriptor() ([]byte, []int) {
	return file_board_group_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateBoardGroupResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

type DeleteBoardGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteBoardGroupRequest) Reset() {
	*x = DeleteBoardGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_board_group_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBoardGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBoardGroupRequest) ProtoMessage() {}

func (x *DeleteBoardGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_board_group_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBoardGroupRequest.ProtoReflect.Descriptor instead.
func (*DeleteBoardGroupRequest) Descriptor() ([]byte, []int) {
	return file_board_group_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteBoardGroupRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteBoardGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *DeleteBoardGroupResponse) Reset() {
	*x = DeleteBoardGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_board_group_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBoardGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBoardGroupResponse) ProtoMessage() {}

func (x *DeleteBoardGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_board_group_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBoardGroupResponse.ProtoReflect.Descriptor instead.
func (*DeleteBoardGroupResponse) Descriptor() ([]byte, []int) {
	return file_board_group_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteBoardGroupResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

type BoardGroupsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BoardGroups []*BoardGroup `protobuf:"bytes,1,rep,name=boardGroups,proto3" json:"boardGroups,omitempty"`
}

func (x *BoardGroupsInfo) Reset() {
	*x = BoardGroupsInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_board_group_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoardGroupsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoardGroupsInfo) ProtoMessage() {}

func (x *BoardGroupsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_board_group_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoardGroupsInfo.ProtoReflect.Descriptor instead.
func (*BoardGroupsInfo) Descriptor() ([]byte, []int) {
	return file_board_group_proto_rawDescGZIP(), []int{9}
}

func (x *BoardGroupsInfo) GetBoardGroups() []*BoardGroup {
	if x != nil {
		return x.BoardGroups
	}
	return nil
}

var File_board_group_proto protoreflect.FileDescriptor

var file_board_group_proto_rawDesc = []byte{
	0x0a, 0x11, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a,
	0x0a, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x26, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x50, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x42, 0x6f,
	0x61, 0x72, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x37, 0x0a, 0x0a, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x0a, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x52, 0x0a, 0x17, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x0a, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x0a, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x32, 0x0a,
	0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x52, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x0a,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x42,
	0x6f, 0x61, 0x72, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x0a, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x32, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42,
	0x6f, 0x61, 0x72, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x29, 0x0a, 0x17, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x32, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f,
	0x61, 0x72, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x4c, 0x0a, 0x0f, 0x42, 0x6f, 0x61, 0x72,
	0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x39, 0x0a, 0x0b, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x42,
	0x6f, 0x61, 0x72, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x0b, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x32, 0xd1, 0x03, 0x0a, 0x11, 0x42, 0x6f, 0x61, 0x72, 0x64,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0c,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x53, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x64,
	0x12, 0x21, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x24, 0x2e, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x42, 0x6f, 0x61, 0x72, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x24, 0x2e, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x25, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a, 0x10, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x24, 0x2e,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1e, 0x5a, 0x1c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4a, 0x49, 0x45, 0x45, 0x45, 0x4e, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2d, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_board_group_proto_rawDescOnce sync.Once
	file_board_group_proto_rawDescData = file_board_group_proto_rawDesc
)

func file_board_group_proto_rawDescGZIP() []byte {
	file_board_group_proto_rawDescOnce.Do(func() {
		file_board_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_board_group_proto_rawDescData)
	})
	return file_board_group_proto_rawDescData
}

var file_board_group_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_board_group_proto_goTypes = []interface{}{
	(*BoardGroup)(nil),               // 0: board_group.BoardGroup
	(*GetBoardGroupRequest)(nil),     // 1: board_group.GetBoardGroupRequest
	(*GetBoardGroupResponse)(nil),    // 2: board_group.GetBoardGroupResponse
	(*CreateBoardGroupRequest)(nil),  // 3: board_group.CreateBoardGroupRequest
	(*CreateBoardGroupResponse)(nil), // 4: board_group.CreateBoardGroupResponse
	(*UpdateBoardGroupRequest)(nil),  // 5: board_group.UpdateBoardGroupRequest
	(*UpdateBoardGroupResponse)(nil), // 6: board_group.UpdateBoardGroupResponse
	(*DeleteBoardGroupRequest)(nil),  // 7: board_group.DeleteBoardGroupRequest
	(*DeleteBoardGroupResponse)(nil), // 8: board_group.DeleteBoardGroupResponse
	(*BoardGroupsInfo)(nil),          // 9: board_group.BoardGroupsInfo
	(*emptypb.Empty)(nil),            // 10: google.protobuf.Empty
}
var file_board_group_proto_depIdxs = []int32{
	0,  // 0: board_group.GetBoardGroupResponse.boardGroup:type_name -> board_group.BoardGroup
	0,  // 1: board_group.CreateBoardGroupRequest.boardGroup:type_name -> board_group.BoardGroup
	0,  // 2: board_group.UpdateBoardGroupRequest.boardGroup:type_name -> board_group.BoardGroup
	0,  // 3: board_group.BoardGroupsInfo.boardGroups:type_name -> board_group.BoardGroup
	10, // 4: board_group.BoardGroupService.GetAllBoards:input_type -> google.protobuf.Empty
	1,  // 5: board_group.BoardGroupService.GetBoardId:input_type -> board_group.GetBoardGroupRequest
	3,  // 6: board_group.BoardGroupService.CreateBoardGroup:input_type -> board_group.CreateBoardGroupRequest
	5,  // 7: board_group.BoardGroupService.UpdateBoardGroup:input_type -> board_group.UpdateBoardGroupRequest
	7,  // 8: board_group.BoardGroupService.DeleteBoardGroup:input_type -> board_group.DeleteBoardGroupRequest
	9,  // 9: board_group.BoardGroupService.GetAllBoards:output_type -> board_group.BoardGroupsInfo
	2,  // 10: board_group.BoardGroupService.GetBoardId:output_type -> board_group.GetBoardGroupResponse
	4,  // 11: board_group.BoardGroupService.CreateBoardGroup:output_type -> board_group.CreateBoardGroupResponse
	6,  // 12: board_group.BoardGroupService.UpdateBoardGroup:output_type -> board_group.UpdateBoardGroupResponse
	8,  // 13: board_group.BoardGroupService.DeleteBoardGroup:output_type -> board_group.DeleteBoardGroupResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_board_group_proto_init() }
func file_board_group_proto_init() {
	if File_board_group_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_board_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoardGroup); i {
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
		file_board_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBoardGroupRequest); i {
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
		file_board_group_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBoardGroupResponse); i {
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
		file_board_group_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBoardGroupRequest); i {
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
		file_board_group_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBoardGroupResponse); i {
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
		file_board_group_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBoardGroupRequest); i {
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
		file_board_group_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBoardGroupResponse); i {
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
		file_board_group_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBoardGroupRequest); i {
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
		file_board_group_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBoardGroupResponse); i {
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
		file_board_group_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoardGroupsInfo); i {
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
			RawDescriptor: file_board_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_board_group_proto_goTypes,
		DependencyIndexes: file_board_group_proto_depIdxs,
		MessageInfos:      file_board_group_proto_msgTypes,
	}.Build()
	File_board_group_proto = out.File
	file_board_group_proto_rawDesc = nil
	file_board_group_proto_goTypes = nil
	file_board_group_proto_depIdxs = nil
}
