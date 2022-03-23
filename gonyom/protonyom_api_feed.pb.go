// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: protonyom_api_feed.proto

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

type AddFeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Feed *Feed `protobuf:"bytes,1,opt,name=feed,proto3" json:"feed,omitempty"`
}

func (x *AddFeedRequest) Reset() {
	*x = AddFeedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protonyom_api_feed_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFeedRequest) ProtoMessage() {}

func (x *AddFeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protonyom_api_feed_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFeedRequest.ProtoReflect.Descriptor instead.
func (*AddFeedRequest) Descriptor() ([]byte, []int) {
	return file_protonyom_api_feed_proto_rawDescGZIP(), []int{0}
}

func (x *AddFeedRequest) GetFeed() *Feed {
	if x != nil {
		return x.Feed
	}
	return nil
}

type AddFeedReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Feed *Feed `protobuf:"bytes,1,opt,name=feed,proto3" json:"feed,omitempty"`
}

func (x *AddFeedReply) Reset() {
	*x = AddFeedReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protonyom_api_feed_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFeedReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFeedReply) ProtoMessage() {}

func (x *AddFeedReply) ProtoReflect() protoreflect.Message {
	mi := &file_protonyom_api_feed_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFeedReply.ProtoReflect.Descriptor instead.
func (*AddFeedReply) Descriptor() ([]byte, []int) {
	return file_protonyom_api_feed_proto_rawDescGZIP(), []int{1}
}

func (x *AddFeedReply) GetFeed() *Feed {
	if x != nil {
		return x.Feed
	}
	return nil
}

type GetFeedsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PetId      string `protobuf:"bytes,1,opt,name=petId,proto3" json:"petId,omitempty"`
	StartAfter int64  `protobuf:"varint,2,opt,name=startAfter,proto3" json:"startAfter,omitempty"` // last feed timestamp
	Limit      int32  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetFeedsRequest) Reset() {
	*x = GetFeedsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protonyom_api_feed_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFeedsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeedsRequest) ProtoMessage() {}

func (x *GetFeedsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protonyom_api_feed_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFeedsRequest.ProtoReflect.Descriptor instead.
func (*GetFeedsRequest) Descriptor() ([]byte, []int) {
	return file_protonyom_api_feed_proto_rawDescGZIP(), []int{2}
}

func (x *GetFeedsRequest) GetPetId() string {
	if x != nil {
		return x.PetId
	}
	return ""
}

func (x *GetFeedsRequest) GetStartAfter() int64 {
	if x != nil {
		return x.StartAfter
	}
	return 0
}

func (x *GetFeedsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetFeedsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Feeds []*Feed `protobuf:"bytes,1,rep,name=feeds,proto3" json:"feeds,omitempty"`
}

func (x *GetFeedsReply) Reset() {
	*x = GetFeedsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protonyom_api_feed_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFeedsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeedsReply) ProtoMessage() {}

func (x *GetFeedsReply) ProtoReflect() protoreflect.Message {
	mi := &file_protonyom_api_feed_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFeedsReply.ProtoReflect.Descriptor instead.
func (*GetFeedsReply) Descriptor() ([]byte, []int) {
	return file_protonyom_api_feed_proto_rawDescGZIP(), []int{3}
}

func (x *GetFeedsReply) GetFeeds() []*Feed {
	if x != nil {
		return x.Feeds
	}
	return nil
}

type DeleteFeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PetId  string `protobuf:"bytes,1,opt,name=petId,proto3" json:"petId,omitempty"`
	FeedId string `protobuf:"bytes,2,opt,name=feedId,proto3" json:"feedId,omitempty"`
}

func (x *DeleteFeedRequest) Reset() {
	*x = DeleteFeedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protonyom_api_feed_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFeedRequest) ProtoMessage() {}

func (x *DeleteFeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protonyom_api_feed_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFeedRequest.ProtoReflect.Descriptor instead.
func (*DeleteFeedRequest) Descriptor() ([]byte, []int) {
	return file_protonyom_api_feed_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteFeedRequest) GetPetId() string {
	if x != nil {
		return x.PetId
	}
	return ""
}

func (x *DeleteFeedRequest) GetFeedId() string {
	if x != nil {
		return x.FeedId
	}
	return ""
}

type DeleteFeedReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteFeedReply) Reset() {
	*x = DeleteFeedReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protonyom_api_feed_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFeedReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFeedReply) ProtoMessage() {}

func (x *DeleteFeedReply) ProtoReflect() protoreflect.Message {
	mi := &file_protonyom_api_feed_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFeedReply.ProtoReflect.Descriptor instead.
func (*DeleteFeedReply) Descriptor() ([]byte, []int) {
	return file_protonyom_api_feed_proto_rawDescGZIP(), []int{5}
}

type UpdateFeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Feed *Feed `protobuf:"bytes,1,opt,name=feed,proto3" json:"feed,omitempty"`
}

func (x *UpdateFeedRequest) Reset() {
	*x = UpdateFeedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protonyom_api_feed_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFeedRequest) ProtoMessage() {}

func (x *UpdateFeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protonyom_api_feed_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFeedRequest.ProtoReflect.Descriptor instead.
func (*UpdateFeedRequest) Descriptor() ([]byte, []int) {
	return file_protonyom_api_feed_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateFeedRequest) GetFeed() *Feed {
	if x != nil {
		return x.Feed
	}
	return nil
}

type UpdateFeedReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Feed *Feed `protobuf:"bytes,1,opt,name=feed,proto3" json:"feed,omitempty"`
}

func (x *UpdateFeedReply) Reset() {
	*x = UpdateFeedReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protonyom_api_feed_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFeedReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFeedReply) ProtoMessage() {}

func (x *UpdateFeedReply) ProtoReflect() protoreflect.Message {
	mi := &file_protonyom_api_feed_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFeedReply.ProtoReflect.Descriptor instead.
func (*UpdateFeedReply) Descriptor() ([]byte, []int) {
	return file_protonyom_api_feed_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateFeedReply) GetFeed() *Feed {
	if x != nil {
		return x.Feed
	}
	return nil
}

var File_protonyom_api_feed_proto protoreflect.FileDescriptor

var file_protonyom_api_feed_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x79, 0x6f, 0x6d, 0x5f, 0x61, 0x70, 0x69, 0x5f,
	0x66, 0x65, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x6e, 0x79, 0x6f, 0x6d, 0x1a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x79, 0x6f, 0x6d,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a,
	0x0e, 0x41, 0x64, 0x64, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x23, 0x0a, 0x04, 0x66, 0x65, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x79, 0x6f, 0x6d, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x52, 0x04,
	0x66, 0x65, 0x65, 0x64, 0x22, 0x33, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x46, 0x65, 0x65, 0x64, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x23, 0x0a, 0x04, 0x66, 0x65, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x79, 0x6f, 0x6d, 0x2e, 0x46,
	0x65, 0x65, 0x64, 0x52, 0x04, 0x66, 0x65, 0x65, 0x64, 0x22, 0x5d, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x46, 0x65, 0x65, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x65, 0x74,
	0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x41, 0x66, 0x74, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x41, 0x66, 0x74,
	0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x36, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x46,
	0x65, 0x65, 0x64, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x25, 0x0a, 0x05, 0x66, 0x65, 0x65,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x6e, 0x79, 0x6f, 0x6d, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x52, 0x05, 0x66, 0x65, 0x65, 0x64, 0x73,
	0x22, 0x41, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x65, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x66,
	0x65, 0x65, 0x64, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x65, 0x65,
	0x64, 0x49, 0x64, 0x22, 0x11, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x65, 0x65,
	0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x38, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x04, 0x66,
	0x65, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x6e, 0x79, 0x6f, 0x6d, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x52, 0x04, 0x66, 0x65, 0x65, 0x64,
	0x22, 0x36, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x23, 0x0a, 0x04, 0x66, 0x65, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x79, 0x6f, 0x6d, 0x2e, 0x46, 0x65,
	0x65, 0x64, 0x52, 0x04, 0x66, 0x65, 0x65, 0x64, 0x32, 0xa2, 0x02, 0x0a, 0x07, 0x46, 0x65, 0x65,
	0x64, 0x41, 0x70, 0x69, 0x12, 0x3f, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x46, 0x65, 0x65, 0x64, 0x12,
	0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x79, 0x6f, 0x6d, 0x2e, 0x41, 0x64, 0x64, 0x46,
	0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x6e, 0x79, 0x6f, 0x6d, 0x2e, 0x41, 0x64, 0x64, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x64,
	0x73, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x79, 0x6f, 0x6d, 0x2e, 0x47, 0x65,
	0x74, 0x46, 0x65, 0x65, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x79, 0x6f, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65,
	0x64, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0a, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e,
	0x79, 0x6f, 0x6d, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x79, 0x6f,
	0x6d, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65,
	0x64, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x79, 0x6f, 0x6d, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x79, 0x6f, 0x6d, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x24, 0x5a,
	0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x69, 0x63, 0x65,
	0x72, 0x75, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x79, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6e,
	0x79, 0x6f, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protonyom_api_feed_proto_rawDescOnce sync.Once
	file_protonyom_api_feed_proto_rawDescData = file_protonyom_api_feed_proto_rawDesc
)

func file_protonyom_api_feed_proto_rawDescGZIP() []byte {
	file_protonyom_api_feed_proto_rawDescOnce.Do(func() {
		file_protonyom_api_feed_proto_rawDescData = protoimpl.X.CompressGZIP(file_protonyom_api_feed_proto_rawDescData)
	})
	return file_protonyom_api_feed_proto_rawDescData
}

var file_protonyom_api_feed_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_protonyom_api_feed_proto_goTypes = []interface{}{
	(*AddFeedRequest)(nil),    // 0: protonyom.AddFeedRequest
	(*AddFeedReply)(nil),      // 1: protonyom.AddFeedReply
	(*GetFeedsRequest)(nil),   // 2: protonyom.GetFeedsRequest
	(*GetFeedsReply)(nil),     // 3: protonyom.GetFeedsReply
	(*DeleteFeedRequest)(nil), // 4: protonyom.DeleteFeedRequest
	(*DeleteFeedReply)(nil),   // 5: protonyom.DeleteFeedReply
	(*UpdateFeedRequest)(nil), // 6: protonyom.UpdateFeedRequest
	(*UpdateFeedReply)(nil),   // 7: protonyom.UpdateFeedReply
	(*Feed)(nil),              // 8: protonyom.Feed
}
var file_protonyom_api_feed_proto_depIdxs = []int32{
	8, // 0: protonyom.AddFeedRequest.feed:type_name -> protonyom.Feed
	8, // 1: protonyom.AddFeedReply.feed:type_name -> protonyom.Feed
	8, // 2: protonyom.GetFeedsReply.feeds:type_name -> protonyom.Feed
	8, // 3: protonyom.UpdateFeedRequest.feed:type_name -> protonyom.Feed
	8, // 4: protonyom.UpdateFeedReply.feed:type_name -> protonyom.Feed
	0, // 5: protonyom.FeedApi.AddFeed:input_type -> protonyom.AddFeedRequest
	2, // 6: protonyom.FeedApi.GetFeeds:input_type -> protonyom.GetFeedsRequest
	4, // 7: protonyom.FeedApi.DeleteFeed:input_type -> protonyom.DeleteFeedRequest
	6, // 8: protonyom.FeedApi.UpdateFeed:input_type -> protonyom.UpdateFeedRequest
	1, // 9: protonyom.FeedApi.AddFeed:output_type -> protonyom.AddFeedReply
	3, // 10: protonyom.FeedApi.GetFeeds:output_type -> protonyom.GetFeedsReply
	5, // 11: protonyom.FeedApi.DeleteFeed:output_type -> protonyom.DeleteFeedReply
	7, // 12: protonyom.FeedApi.UpdateFeed:output_type -> protonyom.UpdateFeedReply
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_protonyom_api_feed_proto_init() }
func file_protonyom_api_feed_proto_init() {
	if File_protonyom_api_feed_proto != nil {
		return
	}
	file_protonyom_models_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protonyom_api_feed_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFeedRequest); i {
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
		file_protonyom_api_feed_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFeedReply); i {
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
		file_protonyom_api_feed_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFeedsRequest); i {
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
		file_protonyom_api_feed_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFeedsReply); i {
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
		file_protonyom_api_feed_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFeedRequest); i {
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
		file_protonyom_api_feed_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFeedReply); i {
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
		file_protonyom_api_feed_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFeedRequest); i {
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
		file_protonyom_api_feed_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFeedReply); i {
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
			RawDescriptor: file_protonyom_api_feed_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protonyom_api_feed_proto_goTypes,
		DependencyIndexes: file_protonyom_api_feed_proto_depIdxs,
		MessageInfos:      file_protonyom_api_feed_proto_msgTypes,
	}.Build()
	File_protonyom_api_feed_proto = out.File
	file_protonyom_api_feed_proto_rawDesc = nil
	file_protonyom_api_feed_proto_goTypes = nil
	file_protonyom_api_feed_proto_depIdxs = nil
}