// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: protonyom_api_account.proto

package gonyom

import (
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

type GetAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAccountRequest) Reset() {
	*x = GetAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protonyom_api_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountRequest) ProtoMessage() {}

func (x *GetAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protonyom_api_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountRequest.ProtoReflect.Descriptor instead.
func (*GetAccountRequest) Descriptor() ([]byte, []int) {
	return file_protonyom_api_account_proto_rawDescGZIP(), []int{0}
}

type GetAccountReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account *Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *GetAccountReply) Reset() {
	*x = GetAccountReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protonyom_api_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountReply) ProtoMessage() {}

func (x *GetAccountReply) ProtoReflect() protoreflect.Message {
	mi := &file_protonyom_api_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountReply.ProtoReflect.Descriptor instead.
func (*GetAccountReply) Descriptor() ([]byte, []int) {
	return file_protonyom_api_account_proto_rawDescGZIP(), []int{1}
}

func (x *GetAccountReply) GetAccount() *Account {
	if x != nil {
		return x.Account
	}
	return nil
}

type UpdateAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path  string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *UpdateAccountRequest) Reset() {
	*x = UpdateAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protonyom_api_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAccountRequest) ProtoMessage() {}

func (x *UpdateAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protonyom_api_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAccountRequest.ProtoReflect.Descriptor instead.
func (*UpdateAccountRequest) Descriptor() ([]byte, []int) {
	return file_protonyom_api_account_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateAccountRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *UpdateAccountRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type UpdateAccountReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account *Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *UpdateAccountReply) Reset() {
	*x = UpdateAccountReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protonyom_api_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAccountReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAccountReply) ProtoMessage() {}

func (x *UpdateAccountReply) ProtoReflect() protoreflect.Message {
	mi := &file_protonyom_api_account_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAccountReply.ProtoReflect.Descriptor instead.
func (*UpdateAccountReply) Descriptor() ([]byte, []int) {
	return file_protonyom_api_account_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateAccountReply) GetAccount() *Account {
	if x != nil {
		return x.Account
	}
	return nil
}

type DeleteAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteAccountRequest) Reset() {
	*x = DeleteAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protonyom_api_account_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAccountRequest) ProtoMessage() {}

func (x *DeleteAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protonyom_api_account_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAccountRequest.ProtoReflect.Descriptor instead.
func (*DeleteAccountRequest) Descriptor() ([]byte, []int) {
	return file_protonyom_api_account_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteAccountRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteAccountReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteAccountReply) Reset() {
	*x = DeleteAccountReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protonyom_api_account_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAccountReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAccountReply) ProtoMessage() {}

func (x *DeleteAccountReply) ProtoReflect() protoreflect.Message {
	mi := &file_protonyom_api_account_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAccountReply.ProtoReflect.Descriptor instead.
func (*DeleteAccountReply) Descriptor() ([]byte, []int) {
	return file_protonyom_api_account_proto_rawDescGZIP(), []int{5}
}

type AcceptInviteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PetId string `protobuf:"bytes,1,opt,name=petId,proto3" json:"petId,omitempty"`
}

func (x *AcceptInviteRequest) Reset() {
	*x = AcceptInviteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protonyom_api_account_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptInviteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptInviteRequest) ProtoMessage() {}

func (x *AcceptInviteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protonyom_api_account_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptInviteRequest.ProtoReflect.Descriptor instead.
func (*AcceptInviteRequest) Descriptor() ([]byte, []int) {
	return file_protonyom_api_account_proto_rawDescGZIP(), []int{6}
}

func (x *AcceptInviteRequest) GetPetId() string {
	if x != nil {
		return x.PetId
	}
	return ""
}

type AcceptInviteReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account *Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *AcceptInviteReply) Reset() {
	*x = AcceptInviteReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protonyom_api_account_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptInviteReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptInviteReply) ProtoMessage() {}

func (x *AcceptInviteReply) ProtoReflect() protoreflect.Message {
	mi := &file_protonyom_api_account_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptInviteReply.ProtoReflect.Descriptor instead.
func (*AcceptInviteReply) Descriptor() ([]byte, []int) {
	return file_protonyom_api_account_proto_rawDescGZIP(), []int{7}
}

func (x *AcceptInviteReply) GetAccount() *Account {
	if x != nil {
		return x.Account
	}
	return nil
}

type UploadProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProfilePhoto       []byte `protobuf:"bytes,1,opt,name=profilePhoto,proto3" json:"profilePhoto,omitempty"`
	ProfileContentType string `protobuf:"bytes,2,opt,name=profileContentType,proto3" json:"profileContentType,omitempty"`
}

func (x *UploadProfileRequest) Reset() {
	*x = UploadProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protonyom_api_account_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadProfileRequest) ProtoMessage() {}

func (x *UploadProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protonyom_api_account_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadProfileRequest.ProtoReflect.Descriptor instead.
func (*UploadProfileRequest) Descriptor() ([]byte, []int) {
	return file_protonyom_api_account_proto_rawDescGZIP(), []int{8}
}

func (x *UploadProfileRequest) GetProfilePhoto() []byte {
	if x != nil {
		return x.ProfilePhoto
	}
	return nil
}

func (x *UploadProfileRequest) GetProfileContentType() string {
	if x != nil {
		return x.ProfileContentType
	}
	return ""
}

type UploadProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account *Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *UploadProfileResponse) Reset() {
	*x = UploadProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protonyom_api_account_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadProfileResponse) ProtoMessage() {}

func (x *UploadProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protonyom_api_account_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadProfileResponse.ProtoReflect.Descriptor instead.
func (*UploadProfileResponse) Descriptor() ([]byte, []int) {
	return file_protonyom_api_account_proto_rawDescGZIP(), []int{9}
}

func (x *UploadProfileResponse) GetAccount() *Account {
	if x != nil {
		return x.Account
	}
	return nil
}

var File_protonyom_api_account_proto protoreflect.FileDescriptor

var file_protonyom_api_account_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x79, 0x6f, 0x6d, 0x5f, 0x61, 0x70, 0x69, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x79, 0x6f, 0x6d, 0x1a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e,
	0x79, 0x6f, 0x6d, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3f, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2c, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x6e, 0x79, 0x6f, 0x6d, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x40, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x42, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2c,
	0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x79, 0x6f, 0x6d, 0x2e, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x26, 0x0a, 0x14,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2b, 0x0a, 0x13, 0x41, 0x63,
	0x63, 0x65, 0x70, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x65, 0x74, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x11, 0x41, 0x63, 0x63, 0x65, 0x70,
	0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2c, 0x0a, 0x07,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x79, 0x6f, 0x6d, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x6a, 0x0a, 0x14, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x68, 0x6f,
	0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x2e, 0x0a, 0x12, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x12, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x45, 0x0a, 0x15, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2c, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x79, 0x6f, 0x6d, 0x2e, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x8d, 0x03,
	0x0a, 0x0a, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x70, 0x69, 0x12, 0x41, 0x0a, 0x03,
	0x47, 0x65, 0x74, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x79, 0x6f, 0x6d, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x79, 0x6f, 0x6d, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12,
	0x4a, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x6e, 0x79, 0x6f, 0x6d, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x6e, 0x79, 0x6f, 0x6d, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x79, 0x6f,
	0x6d, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x79,
	0x6f, 0x6d, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0c, 0x41, 0x63, 0x63, 0x65, 0x70,
	0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e,
	0x79, 0x6f, 0x6d, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e,
	0x79, 0x6f, 0x6d, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x0d, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x6e, 0x79, 0x6f, 0x6d, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x6e, 0x79, 0x6f, 0x6d, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x24, 0x5a,
	0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x69, 0x63, 0x65,
	0x72, 0x75, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x79, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6e,
	0x79, 0x6f, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protonyom_api_account_proto_rawDescOnce sync.Once
	file_protonyom_api_account_proto_rawDescData = file_protonyom_api_account_proto_rawDesc
)

func file_protonyom_api_account_proto_rawDescGZIP() []byte {
	file_protonyom_api_account_proto_rawDescOnce.Do(func() {
		file_protonyom_api_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_protonyom_api_account_proto_rawDescData)
	})
	return file_protonyom_api_account_proto_rawDescData
}

var file_protonyom_api_account_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_protonyom_api_account_proto_goTypes = []interface{}{
	(*GetAccountRequest)(nil),     // 0: protonyom.GetAccountRequest
	(*GetAccountReply)(nil),       // 1: protonyom.GetAccountReply
	(*UpdateAccountRequest)(nil),  // 2: protonyom.UpdateAccountRequest
	(*UpdateAccountReply)(nil),    // 3: protonyom.UpdateAccountReply
	(*DeleteAccountRequest)(nil),  // 4: protonyom.DeleteAccountRequest
	(*DeleteAccountReply)(nil),    // 5: protonyom.DeleteAccountReply
	(*AcceptInviteRequest)(nil),   // 6: protonyom.AcceptInviteRequest
	(*AcceptInviteReply)(nil),     // 7: protonyom.AcceptInviteReply
	(*UploadProfileRequest)(nil),  // 8: protonyom.UploadProfileRequest
	(*UploadProfileResponse)(nil), // 9: protonyom.UploadProfileResponse
	(*Account)(nil),               // 10: protonyom.Account
}
var file_protonyom_api_account_proto_depIdxs = []int32{
	10, // 0: protonyom.GetAccountReply.account:type_name -> protonyom.Account
	10, // 1: protonyom.UpdateAccountReply.account:type_name -> protonyom.Account
	10, // 2: protonyom.AcceptInviteReply.account:type_name -> protonyom.Account
	10, // 3: protonyom.UploadProfileResponse.account:type_name -> protonyom.Account
	0,  // 4: protonyom.AccountApi.Get:input_type -> protonyom.GetAccountRequest
	2,  // 5: protonyom.AccountApi.Update:input_type -> protonyom.UpdateAccountRequest
	4,  // 6: protonyom.AccountApi.Delete:input_type -> protonyom.DeleteAccountRequest
	6,  // 7: protonyom.AccountApi.AcceptInvite:input_type -> protonyom.AcceptInviteRequest
	8,  // 8: protonyom.AccountApi.UploadProfile:input_type -> protonyom.UploadProfileRequest
	1,  // 9: protonyom.AccountApi.Get:output_type -> protonyom.GetAccountReply
	3,  // 10: protonyom.AccountApi.Update:output_type -> protonyom.UpdateAccountReply
	5,  // 11: protonyom.AccountApi.Delete:output_type -> protonyom.DeleteAccountReply
	7,  // 12: protonyom.AccountApi.AcceptInvite:output_type -> protonyom.AcceptInviteReply
	9,  // 13: protonyom.AccountApi.UploadProfile:output_type -> protonyom.UploadProfileResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_protonyom_api_account_proto_init() }
func file_protonyom_api_account_proto_init() {
	if File_protonyom_api_account_proto != nil {
		return
	}
	file_protonyom_models_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protonyom_api_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountRequest); i {
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
		file_protonyom_api_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountReply); i {
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
		file_protonyom_api_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAccountRequest); i {
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
		file_protonyom_api_account_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAccountReply); i {
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
		file_protonyom_api_account_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAccountRequest); i {
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
		file_protonyom_api_account_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAccountReply); i {
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
		file_protonyom_api_account_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptInviteRequest); i {
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
		file_protonyom_api_account_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptInviteReply); i {
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
		file_protonyom_api_account_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadProfileRequest); i {
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
		file_protonyom_api_account_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadProfileResponse); i {
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
			RawDescriptor: file_protonyom_api_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protonyom_api_account_proto_goTypes,
		DependencyIndexes: file_protonyom_api_account_proto_depIdxs,
		MessageInfos:      file_protonyom_api_account_proto_msgTypes,
	}.Build()
	File_protonyom_api_account_proto = out.File
	file_protonyom_api_account_proto_rawDesc = nil
	file_protonyom_api_account_proto_goTypes = nil
	file_protonyom_api_account_proto_depIdxs = nil
}
