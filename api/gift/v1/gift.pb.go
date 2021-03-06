// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.13.0
// source: api/gift/v1/gift.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GiveGiftRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderId   uint64 `protobuf:"varint,1,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	ReceiverId uint64 `protobuf:"varint,2,opt,name=receiver_id,json=receiverId,proto3" json:"receiver_id,omitempty"`
	GiftId     uint64 `protobuf:"varint,3,opt,name=gift_id,json=giftId,proto3" json:"gift_id,omitempty"`
	Num        uint64 `protobuf:"varint,4,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *GiveGiftRequest) Reset() {
	*x = GiveGiftRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gift_v1_gift_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GiveGiftRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GiveGiftRequest) ProtoMessage() {}

func (x *GiveGiftRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_gift_v1_gift_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GiveGiftRequest.ProtoReflect.Descriptor instead.
func (*GiveGiftRequest) Descriptor() ([]byte, []int) {
	return file_api_gift_v1_gift_proto_rawDescGZIP(), []int{0}
}

func (x *GiveGiftRequest) GetSenderId() uint64 {
	if x != nil {
		return x.SenderId
	}
	return 0
}

func (x *GiveGiftRequest) GetReceiverId() uint64 {
	if x != nil {
		return x.ReceiverId
	}
	return 0
}

func (x *GiveGiftRequest) GetGiftId() uint64 {
	if x != nil {
		return x.GiftId
	}
	return 0
}

func (x *GiveGiftRequest) GetNum() uint64 {
	if x != nil {
		return x.Num
	}
	return 0
}

type GiveGiftReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrCode int32 `protobuf:"varint,1,opt,name=err_code,json=errCode,proto3" json:"err_code,omitempty"`
}

func (x *GiveGiftReply) Reset() {
	*x = GiveGiftReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gift_v1_gift_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GiveGiftReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GiveGiftReply) ProtoMessage() {}

func (x *GiveGiftReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_gift_v1_gift_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GiveGiftReply.ProtoReflect.Descriptor instead.
func (*GiveGiftReply) Descriptor() ([]byte, []int) {
	return file_api_gift_v1_gift_proto_rawDescGZIP(), []int{1}
}

func (x *GiveGiftReply) GetErrCode() int32 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

type GiftHistory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GiftId           uint64  `protobuf:"varint,1,opt,name=gift_id,json=giftId,proto3" json:"gift_id,omitempty"`
	GiftName         string  `protobuf:"bytes,2,opt,name=gift_name,json=giftName,proto3" json:"gift_name,omitempty"`
	GiftValue        float32 `protobuf:"fixed32,3,opt,name=gift_value,json=giftValue,proto3" json:"gift_value,omitempty"`
	SenderId         uint64  `protobuf:"varint,4,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	ReceiverId       uint64  `protobuf:"varint,5,opt,name=receiver_id,json=receiverId,proto3" json:"receiver_id,omitempty"`
	SenderNickname   string  `protobuf:"bytes,6,opt,name=sender_nickname,json=senderNickname,proto3" json:"sender_nickname,omitempty"`
	ReceiverNickname string  `protobuf:"bytes,7,opt,name=receiver_nickname,json=receiverNickname,proto3" json:"receiver_nickname,omitempty"`
	SendTime         uint64  `protobuf:"varint,8,opt,name=send_time,json=sendTime,proto3" json:"send_time,omitempty"`
	Num              uint64  `protobuf:"varint,9,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *GiftHistory) Reset() {
	*x = GiftHistory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gift_v1_gift_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GiftHistory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GiftHistory) ProtoMessage() {}

func (x *GiftHistory) ProtoReflect() protoreflect.Message {
	mi := &file_api_gift_v1_gift_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GiftHistory.ProtoReflect.Descriptor instead.
func (*GiftHistory) Descriptor() ([]byte, []int) {
	return file_api_gift_v1_gift_proto_rawDescGZIP(), []int{2}
}

func (x *GiftHistory) GetGiftId() uint64 {
	if x != nil {
		return x.GiftId
	}
	return 0
}

func (x *GiftHistory) GetGiftName() string {
	if x != nil {
		return x.GiftName
	}
	return ""
}

func (x *GiftHistory) GetGiftValue() float32 {
	if x != nil {
		return x.GiftValue
	}
	return 0
}

func (x *GiftHistory) GetSenderId() uint64 {
	if x != nil {
		return x.SenderId
	}
	return 0
}

func (x *GiftHistory) GetReceiverId() uint64 {
	if x != nil {
		return x.ReceiverId
	}
	return 0
}

func (x *GiftHistory) GetSenderNickname() string {
	if x != nil {
		return x.SenderNickname
	}
	return ""
}

func (x *GiftHistory) GetReceiverNickname() string {
	if x != nil {
		return x.ReceiverNickname
	}
	return ""
}

func (x *GiftHistory) GetSendTime() uint64 {
	if x != nil {
		return x.SendTime
	}
	return 0
}

func (x *GiftHistory) GetNum() uint64 {
	if x != nil {
		return x.Num
	}
	return 0
}

type GetSendingHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetSendingHistoryRequest) Reset() {
	*x = GetSendingHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gift_v1_gift_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSendingHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSendingHistoryRequest) ProtoMessage() {}

func (x *GetSendingHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_gift_v1_gift_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSendingHistoryRequest.ProtoReflect.Descriptor instead.
func (*GetSendingHistoryRequest) Descriptor() ([]byte, []int) {
	return file_api_gift_v1_gift_proto_rawDescGZIP(), []int{3}
}

func (x *GetSendingHistoryRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetSendingHistoryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrCode int32          `protobuf:"varint,1,opt,name=err_code,json=errCode,proto3" json:"err_code,omitempty"`
	History []*GiftHistory `protobuf:"bytes,2,rep,name=history,proto3" json:"history,omitempty"`
}

func (x *GetSendingHistoryReply) Reset() {
	*x = GetSendingHistoryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gift_v1_gift_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSendingHistoryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSendingHistoryReply) ProtoMessage() {}

func (x *GetSendingHistoryReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_gift_v1_gift_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSendingHistoryReply.ProtoReflect.Descriptor instead.
func (*GetSendingHistoryReply) Descriptor() ([]byte, []int) {
	return file_api_gift_v1_gift_proto_rawDescGZIP(), []int{4}
}

func (x *GetSendingHistoryReply) GetErrCode() int32 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

func (x *GetSendingHistoryReply) GetHistory() []*GiftHistory {
	if x != nil {
		return x.History
	}
	return nil
}

type GetReceivingHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetReceivingHistoryRequest) Reset() {
	*x = GetReceivingHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gift_v1_gift_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReceivingHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReceivingHistoryRequest) ProtoMessage() {}

func (x *GetReceivingHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_gift_v1_gift_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReceivingHistoryRequest.ProtoReflect.Descriptor instead.
func (*GetReceivingHistoryRequest) Descriptor() ([]byte, []int) {
	return file_api_gift_v1_gift_proto_rawDescGZIP(), []int{5}
}

func (x *GetReceivingHistoryRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetReceivingHistoryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrCode int32          `protobuf:"varint,1,opt,name=err_code,json=errCode,proto3" json:"err_code,omitempty"`
	History []*GiftHistory `protobuf:"bytes,2,rep,name=history,proto3" json:"history,omitempty"`
}

func (x *GetReceivingHistoryReply) Reset() {
	*x = GetReceivingHistoryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gift_v1_gift_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReceivingHistoryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReceivingHistoryReply) ProtoMessage() {}

func (x *GetReceivingHistoryReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_gift_v1_gift_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReceivingHistoryReply.ProtoReflect.Descriptor instead.
func (*GetReceivingHistoryReply) Descriptor() ([]byte, []int) {
	return file_api_gift_v1_gift_proto_rawDescGZIP(), []int{6}
}

func (x *GetReceivingHistoryReply) GetErrCode() int32 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

func (x *GetReceivingHistoryReply) GetHistory() []*GiftHistory {
	if x != nil {
		return x.History
	}
	return nil
}

var File_api_gift_v1_gift_proto protoreflect.FileDescriptor

var file_api_gift_v1_gift_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x69, 0x66, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x69,
	0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x69,
	0x66, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x01, 0x0a,
	0x0f, 0x47, 0x69, 0x76, 0x65, 0x47, 0x69, 0x66, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x24, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x28, 0x01, 0x52, 0x08, 0x73, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x32, 0x02, 0x28, 0x01, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x20, 0x0a, 0x07, 0x67, 0x69, 0x66, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x28, 0x01, 0x52, 0x06, 0x67, 0x69, 0x66, 0x74,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x28, 0x01, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x2a, 0x0a,
	0x0d, 0x47, 0x69, 0x76, 0x65, 0x47, 0x69, 0x66, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x19,
	0x0a, 0x08, 0x65, 0x72, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x22, 0xa5, 0x02, 0x0a, 0x0b, 0x47, 0x69,
	0x66, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x69, 0x66,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x67, 0x69, 0x66, 0x74,
	0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x69, 0x66, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x69, 0x66, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x67, 0x69, 0x66, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x09, 0x67, 0x69, 0x66, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0a, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f,
	0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4e, 0x69, 0x63,
	0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x72, 0x5f, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6e, 0x75,
	0x6d, 0x22, 0x33, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x67, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x53, 0x65, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x19, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x32, 0x0a, 0x07, 0x68,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x67, 0x69, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x69, 0x66, 0x74, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x07, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x22,
	0x35, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x69, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x32, 0x0a,
	0x07, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x69, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x69, 0x66,
	0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x07, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x32, 0x86, 0x03, 0x0a, 0x04, 0x47, 0x69, 0x66, 0x74, 0x12, 0x60, 0x0a, 0x08, 0x47, 0x69,
	0x76, 0x65, 0x47, 0x69, 0x66, 0x74, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x69, 0x66,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x69, 0x76, 0x65, 0x47, 0x69, 0x66, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x69, 0x66, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x69, 0x76, 0x65, 0x47, 0x69, 0x66, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x67, 0x69, 0x66, 0x74, 0x2f,
	0x67, 0x69, 0x76, 0x65, 0x5f, 0x67, 0x69, 0x66, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x88, 0x01, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x12, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x69, 0x66, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x67, 0x69, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x27,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x12, 0x1f, 0x2f, 0x67, 0x69, 0x66, 0x74, 0x2f, 0x73, 0x65,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x7b, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x90, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12,
	0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x69, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67,
	0x69, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x69, 0x6e, 0x67, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x12, 0x21, 0x2f, 0x67, 0x69, 0x66, 0x74, 0x2f, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x2e, 0x0a, 0x0b, 0x61, 0x70,
	0x69, 0x2e, 0x67, 0x69, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x1d, 0x65, 0x63, 0x6f,
	0x6e, 0x6f, 0x6d, 0x79, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x69, 0x66, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_gift_v1_gift_proto_rawDescOnce sync.Once
	file_api_gift_v1_gift_proto_rawDescData = file_api_gift_v1_gift_proto_rawDesc
)

func file_api_gift_v1_gift_proto_rawDescGZIP() []byte {
	file_api_gift_v1_gift_proto_rawDescOnce.Do(func() {
		file_api_gift_v1_gift_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_gift_v1_gift_proto_rawDescData)
	})
	return file_api_gift_v1_gift_proto_rawDescData
}

var file_api_gift_v1_gift_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_gift_v1_gift_proto_goTypes = []interface{}{
	(*GiveGiftRequest)(nil),            // 0: api.gift.v1.GiveGiftRequest
	(*GiveGiftReply)(nil),              // 1: api.gift.v1.GiveGiftReply
	(*GiftHistory)(nil),                // 2: api.gift.v1.GiftHistory
	(*GetSendingHistoryRequest)(nil),   // 3: api.gift.v1.GetSendingHistoryRequest
	(*GetSendingHistoryReply)(nil),     // 4: api.gift.v1.GetSendingHistoryReply
	(*GetReceivingHistoryRequest)(nil), // 5: api.gift.v1.GetReceivingHistoryRequest
	(*GetReceivingHistoryReply)(nil),   // 6: api.gift.v1.GetReceivingHistoryReply
}
var file_api_gift_v1_gift_proto_depIdxs = []int32{
	2, // 0: api.gift.v1.GetSendingHistoryReply.history:type_name -> api.gift.v1.GiftHistory
	2, // 1: api.gift.v1.GetReceivingHistoryReply.history:type_name -> api.gift.v1.GiftHistory
	0, // 2: api.gift.v1.Gift.GiveGift:input_type -> api.gift.v1.GiveGiftRequest
	3, // 3: api.gift.v1.Gift.GetSendingHistory:input_type -> api.gift.v1.GetSendingHistoryRequest
	5, // 4: api.gift.v1.Gift.GetReceivingHistory:input_type -> api.gift.v1.GetReceivingHistoryRequest
	1, // 5: api.gift.v1.Gift.GiveGift:output_type -> api.gift.v1.GiveGiftReply
	4, // 6: api.gift.v1.Gift.GetSendingHistory:output_type -> api.gift.v1.GetSendingHistoryReply
	6, // 7: api.gift.v1.Gift.GetReceivingHistory:output_type -> api.gift.v1.GetReceivingHistoryReply
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_gift_v1_gift_proto_init() }
func file_api_gift_v1_gift_proto_init() {
	if File_api_gift_v1_gift_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_gift_v1_gift_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GiveGiftRequest); i {
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
		file_api_gift_v1_gift_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GiveGiftReply); i {
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
		file_api_gift_v1_gift_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GiftHistory); i {
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
		file_api_gift_v1_gift_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSendingHistoryRequest); i {
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
		file_api_gift_v1_gift_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSendingHistoryReply); i {
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
		file_api_gift_v1_gift_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReceivingHistoryRequest); i {
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
		file_api_gift_v1_gift_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReceivingHistoryReply); i {
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
			RawDescriptor: file_api_gift_v1_gift_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_gift_v1_gift_proto_goTypes,
		DependencyIndexes: file_api_gift_v1_gift_proto_depIdxs,
		MessageInfos:      file_api_gift_v1_gift_proto_msgTypes,
	}.Build()
	File_api_gift_v1_gift_proto = out.File
	file_api_gift_v1_gift_proto_rawDesc = nil
	file_api_gift_v1_gift_proto_goTypes = nil
	file_api_gift_v1_gift_proto_depIdxs = nil
}
