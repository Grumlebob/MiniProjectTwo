// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: grpc/interface.proto

package grpc

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

type Bid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId    int32 `protobuf:"varint,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
	UniqueBidId int32 `protobuf:"varint,2,opt,name=uniqueBidId,proto3" json:"uniqueBidId,omitempty"`
	Amount      int32 `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *Bid) Reset() {
	*x = Bid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_interface_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bid) ProtoMessage() {}

func (x *Bid) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_interface_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bid.ProtoReflect.Descriptor instead.
func (*Bid) Descriptor() ([]byte, []int) {
	return file_grpc_interface_proto_rawDescGZIP(), []int{0}
}

func (x *Bid) GetClientId() int32 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

func (x *Bid) GetUniqueBidId() int32 {
	if x != nil {
		return x.UniqueBidId
	}
	return 0
}

func (x *Bid) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type Acknowledgement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ack string `protobuf:"bytes,1,opt,name=ack,proto3" json:"ack,omitempty"`
}

func (x *Acknowledgement) Reset() {
	*x = Acknowledgement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_interface_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Acknowledgement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Acknowledgement) ProtoMessage() {}

func (x *Acknowledgement) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_interface_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Acknowledgement.ProtoReflect.Descriptor instead.
func (*Acknowledgement) Descriptor() ([]byte, []int) {
	return file_grpc_interface_proto_rawDescGZIP(), []int{1}
}

func (x *Acknowledgement) GetAck() string {
	if x != nil {
		return x.Ack
	}
	return ""
}

type Outcome struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuctionStatus string `protobuf:"bytes,1,opt,name=auctionStatus,proto3" json:"auctionStatus,omitempty"`
	HighestBid    int32  `protobuf:"varint,2,opt,name=highestBid,proto3" json:"highestBid,omitempty"`
}

func (x *Outcome) Reset() {
	*x = Outcome{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_interface_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Outcome) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Outcome) ProtoMessage() {}

func (x *Outcome) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_interface_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Outcome.ProtoReflect.Descriptor instead.
func (*Outcome) Descriptor() ([]byte, []int) {
	return file_grpc_interface_proto_rawDescGZIP(), []int{2}
}

func (x *Outcome) GetAuctionStatus() string {
	if x != nil {
		return x.AuctionStatus
	}
	return ""
}

func (x *Outcome) GetHighestBid() int32 {
	if x != nil {
		return x.HighestBid
	}
	return 0
}

type Replicate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuctionStatus              string `protobuf:"bytes,1,opt,name=auctionStatus,proto3" json:"auctionStatus,omitempty"`
	HighestBidOnCurrentAuction int32  `protobuf:"varint,2,opt,name=highestBidOnCurrentAuction,proto3" json:"highestBidOnCurrentAuction,omitempty"`
	ResponseForRequest         string `protobuf:"bytes,3,opt,name=responseForRequest,proto3" json:"responseForRequest,omitempty"`
	UniqueIdentifierForRequest int32  `protobuf:"varint,4,opt,name=uniqueIdentifierForRequest,proto3" json:"uniqueIdentifierForRequest,omitempty"`
	CurrentItem                string `protobuf:"bytes,5,opt,name=currentItem,proto3" json:"currentItem,omitempty"`
	WinnerId                   int32  `protobuf:"varint,6,opt,name=winnerId,proto3" json:"winnerId,omitempty"`
}

func (x *Replicate) Reset() {
	*x = Replicate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_interface_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Replicate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Replicate) ProtoMessage() {}

func (x *Replicate) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_interface_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Replicate.ProtoReflect.Descriptor instead.
func (*Replicate) Descriptor() ([]byte, []int) {
	return file_grpc_interface_proto_rawDescGZIP(), []int{3}
}

func (x *Replicate) GetAuctionStatus() string {
	if x != nil {
		return x.AuctionStatus
	}
	return ""
}

func (x *Replicate) GetHighestBidOnCurrentAuction() int32 {
	if x != nil {
		return x.HighestBidOnCurrentAuction
	}
	return 0
}

func (x *Replicate) GetResponseForRequest() string {
	if x != nil {
		return x.ResponseForRequest
	}
	return ""
}

func (x *Replicate) GetUniqueIdentifierForRequest() int32 {
	if x != nil {
		return x.UniqueIdentifierForRequest
	}
	return 0
}

func (x *Replicate) GetCurrentItem() string {
	if x != nil {
		return x.CurrentItem
	}
	return ""
}

func (x *Replicate) GetWinnerId() int32 {
	if x != nil {
		return x.WinnerId
	}
	return 0
}

var File_grpc_interface_proto protoreflect.FileDescriptor

var file_grpc_interface_proto_rawDesc = []byte{
	0x0a, 0x14, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x03, 0x42, 0x69, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x42, 0x69, 0x64, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x42, 0x69, 0x64, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x23, 0x0a, 0x0f, 0x41, 0x63, 0x6b, 0x6e, 0x6f, 0x77,
	0x6c, 0x65, 0x64, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x63, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x63, 0x6b, 0x22, 0x4f, 0x0a, 0x07, 0x4f,
	0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a,
	0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x22, 0x9f, 0x02, 0x0a,
	0x09, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x3e, 0x0a, 0x1a, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x4f, 0x6e,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x1a, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64,
	0x4f, 0x6e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2e, 0x0a, 0x12, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x6f, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x3e, 0x0a, 0x1a, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x1a, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x32, 0xbd,
	0x02, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x53, 0x0a, 0x27, 0x68, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x41, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x72, 0x6f, 0x6d, 0x4c, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x12, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x1a, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x63, 0x6b, 0x6e, 0x6f,
	0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0a,
	0x70, 0x69, 0x6e, 0x67, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x63, 0x6b, 0x6e, 0x6f, 0x77,
	0x6c, 0x65, 0x64, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x10, 0x62,
	0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x29, 0x0a, 0x03, 0x62, 0x69, 0x64, 0x12, 0x09, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x42,
	0x69, 0x64, 0x1a, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x63, 0x6b, 0x6e, 0x6f, 0x77,
	0x6c, 0x65, 0x64, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0d, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x22, 0x00, 0x42, 0x3c,
	0x5a, 0x3a, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x72, 0x75, 0x6d, 0x6c, 0x65, 0x62, 0x6f, 0x62, 0x2f, 0x41,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_interface_proto_rawDescOnce sync.Once
	file_grpc_interface_proto_rawDescData = file_grpc_interface_proto_rawDesc
)

func file_grpc_interface_proto_rawDescGZIP() []byte {
	file_grpc_interface_proto_rawDescOnce.Do(func() {
		file_grpc_interface_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_interface_proto_rawDescData)
	})
	return file_grpc_interface_proto_rawDescData
}

var file_grpc_interface_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_grpc_interface_proto_goTypes = []interface{}{
	(*Bid)(nil),             // 0: grpc.Bid
	(*Acknowledgement)(nil), // 1: grpc.Acknowledgement
	(*Outcome)(nil),         // 2: grpc.Outcome
	(*Replicate)(nil),       // 3: grpc.Replicate
	(*emptypb.Empty)(nil),   // 4: google.protobuf.Empty
}
var file_grpc_interface_proto_depIdxs = []int32{
	3, // 0: grpc.Node.handleAgreementAndReplicationFromLeader:input_type -> grpc.Replicate
	4, // 1: grpc.Node.pingLeader:input_type -> google.protobuf.Empty
	1, // 2: grpc.Node.broadcastMessage:input_type -> grpc.Acknowledgement
	0, // 3: grpc.Node.bid:input_type -> grpc.Bid
	4, // 4: grpc.Node.result:input_type -> google.protobuf.Empty
	1, // 5: grpc.Node.handleAgreementAndReplicationFromLeader:output_type -> grpc.Acknowledgement
	1, // 6: grpc.Node.pingLeader:output_type -> grpc.Acknowledgement
	4, // 7: grpc.Node.broadcastMessage:output_type -> google.protobuf.Empty
	1, // 8: grpc.Node.bid:output_type -> grpc.Acknowledgement
	2, // 9: grpc.Node.result:output_type -> grpc.Outcome
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_interface_proto_init() }
func file_grpc_interface_proto_init() {
	if File_grpc_interface_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_interface_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bid); i {
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
		file_grpc_interface_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Acknowledgement); i {
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
		file_grpc_interface_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Outcome); i {
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
		file_grpc_interface_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Replicate); i {
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
			RawDescriptor: file_grpc_interface_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_interface_proto_goTypes,
		DependencyIndexes: file_grpc_interface_proto_depIdxs,
		MessageInfos:      file_grpc_interface_proto_msgTypes,
	}.Build()
	File_grpc_interface_proto = out.File
	file_grpc_interface_proto_rawDesc = nil
	file_grpc_interface_proto_goTypes = nil
	file_grpc_interface_proto_depIdxs = nil
}