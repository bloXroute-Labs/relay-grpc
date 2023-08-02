// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.3
// source: blxr-mev.proto

// The gateway service definition.

package relaygrpc

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

type TxHashListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthHeader string   `protobuf:"bytes,1,opt,name=auth_header,json=authHeader,proto3" json:"auth_header,omitempty"`
	TxHashes   [][]byte `protobuf:"bytes,2,rep,name=txHashes,proto3" json:"txHashes,omitempty"`
}

func (x *TxHashListRequest) Reset() {
	*x = TxHashListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blxr_mev_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxHashListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxHashListRequest) ProtoMessage() {}

func (x *TxHashListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blxr_mev_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxHashListRequest.ProtoReflect.Descriptor instead.
func (*TxHashListRequest) Descriptor() ([]byte, []int) {
	return file_blxr_mev_proto_rawDescGZIP(), []int{0}
}

func (x *TxHashListRequest) GetAuthHeader() string {
	if x != nil {
		return x.AuthHeader
	}
	return ""
}

func (x *TxHashListRequest) GetTxHashes() [][]byte {
	if x != nil {
		return x.TxHashes
	}
	return nil
}

type ShortIDListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortIDs []uint32 `protobuf:"varint,1,rep,packed,name=shortIDs,proto3" json:"shortIDs,omitempty"`
}

func (x *ShortIDListReply) Reset() {
	*x = ShortIDListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blxr_mev_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortIDListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortIDListReply) ProtoMessage() {}

func (x *ShortIDListReply) ProtoReflect() protoreflect.Message {
	mi := &file_blxr_mev_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortIDListReply.ProtoReflect.Descriptor instead.
func (*ShortIDListReply) Descriptor() ([]byte, []int) {
	return file_blxr_mev_proto_rawDescGZIP(), []int{1}
}

func (x *ShortIDListReply) GetShortIDs() []uint32 {
	if x != nil {
		return x.ShortIDs
	}
	return nil
}

type ProposedBlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthHeader           string        `protobuf:"bytes,1,opt,name=auth_header,json=authHeader,proto3" json:"auth_header,omitempty"`
	ValidatorHttpAddress string        `protobuf:"bytes,2,opt,name=validatorHttpAddress,proto3" json:"validatorHttpAddress,omitempty"`
	Namespace            string        `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	BlockNumber          uint64        `protobuf:"varint,4,opt,name=blockNumber,proto3" json:"blockNumber,omitempty"`
	PrevBlockHash        string        `protobuf:"bytes,5,opt,name=prevBlockHash,proto3" json:"prevBlockHash,omitempty"`
	BlockReward          string        `protobuf:"bytes,6,opt,name=blockReward,proto3" json:"blockReward,omitempty"`
	GasLimit             uint64        `protobuf:"varint,7,opt,name=gasLimit,proto3" json:"gasLimit,omitempty"`
	GasUsed              uint64        `protobuf:"varint,8,opt,name=gasUsed,proto3" json:"gasUsed,omitempty"`
	Payload              []*CompressTx `protobuf:"bytes,9,rep,name=payload,proto3" json:"payload,omitempty"`
}

func (x *ProposedBlockRequest) Reset() {
	*x = ProposedBlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blxr_mev_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProposedBlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProposedBlockRequest) ProtoMessage() {}

func (x *ProposedBlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blxr_mev_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProposedBlockRequest.ProtoReflect.Descriptor instead.
func (*ProposedBlockRequest) Descriptor() ([]byte, []int) {
	return file_blxr_mev_proto_rawDescGZIP(), []int{2}
}

func (x *ProposedBlockRequest) GetAuthHeader() string {
	if x != nil {
		return x.AuthHeader
	}
	return ""
}

func (x *ProposedBlockRequest) GetValidatorHttpAddress() string {
	if x != nil {
		return x.ValidatorHttpAddress
	}
	return ""
}

func (x *ProposedBlockRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ProposedBlockRequest) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *ProposedBlockRequest) GetPrevBlockHash() string {
	if x != nil {
		return x.PrevBlockHash
	}
	return ""
}

func (x *ProposedBlockRequest) GetBlockReward() string {
	if x != nil {
		return x.BlockReward
	}
	return ""
}

func (x *ProposedBlockRequest) GetGasLimit() uint64 {
	if x != nil {
		return x.GasLimit
	}
	return 0
}

func (x *ProposedBlockRequest) GetGasUsed() uint64 {
	if x != nil {
		return x.GasUsed
	}
	return 0
}

func (x *ProposedBlockRequest) GetPayload() []*CompressTx {
	if x != nil {
		return x.Payload
	}
	return nil
}

type CompressTx struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RawData []byte `protobuf:"bytes,1,opt,name=rawData,proto3" json:"rawData,omitempty"`
	ShortID uint32 `protobuf:"varint,2,opt,name=shortID,proto3" json:"shortID,omitempty"`
}

func (x *CompressTx) Reset() {
	*x = CompressTx{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blxr_mev_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompressTx) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompressTx) ProtoMessage() {}

func (x *CompressTx) ProtoReflect() protoreflect.Message {
	mi := &file_blxr_mev_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompressTx.ProtoReflect.Descriptor instead.
func (*CompressTx) Descriptor() ([]byte, []int) {
	return file_blxr_mev_proto_rawDescGZIP(), []int{3}
}

func (x *CompressTx) GetRawData() []byte {
	if x != nil {
		return x.RawData
	}
	return nil
}

func (x *CompressTx) GetShortID() uint32 {
	if x != nil {
		return x.ShortID
	}
	return 0
}

type ProposedBlockReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ValidatorReply     string `protobuf:"bytes,1,opt,name=validatorReply,proto3" json:"validatorReply,omitempty"`
	ValidatorReplyTime int64  `protobuf:"varint,2,opt,name=validatorReplyTime,proto3" json:"validatorReplyTime,omitempty"`
}

func (x *ProposedBlockReply) Reset() {
	*x = ProposedBlockReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blxr_mev_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProposedBlockReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProposedBlockReply) ProtoMessage() {}

func (x *ProposedBlockReply) ProtoReflect() protoreflect.Message {
	mi := &file_blxr_mev_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProposedBlockReply.ProtoReflect.Descriptor instead.
func (*ProposedBlockReply) Descriptor() ([]byte, []int) {
	return file_blxr_mev_proto_rawDescGZIP(), []int{4}
}

func (x *ProposedBlockReply) GetValidatorReply() string {
	if x != nil {
		return x.ValidatorReply
	}
	return ""
}

func (x *ProposedBlockReply) GetValidatorReplyTime() int64 {
	if x != nil {
		return x.ValidatorReplyTime
	}
	return 0
}

var File_blxr_mev_proto protoreflect.FileDescriptor

var file_blxr_mev_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x62, 0x6c, 0x78, 0x72, 0x2d, 0x6d, 0x65, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x22, 0x50, 0x0a, 0x11, 0x54, 0x78, 0x48,
	0x61, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x1a, 0x0a, 0x08, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0c, 0x52, 0x08, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x22, 0x2e, 0x0a, 0x10, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x49, 0x44, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x1a, 0x0a, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x49, 0x44, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x49, 0x44, 0x73, 0x22, 0xd8, 0x02, 0x0a, 0x14,
	0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x14, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x48, 0x74, 0x74, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x14, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x48, 0x74,
	0x74, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x72, 0x65,
	0x76, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x70, 0x72, 0x65, 0x76, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12,
	0x20, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x61, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x67, 0x61, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x67, 0x61, 0x73, 0x55, 0x73, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07,
	0x67, 0x61, 0x73, 0x55, 0x73, 0x65, 0x64, 0x12, 0x2d, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x54, 0x78, 0x52, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x40, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x54, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x61, 0x77, 0x44, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x72, 0x61, 0x77, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x49, 0x44, 0x22, 0x6c, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x26,
	0x0a, 0x0e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2e, 0x0a, 0x12, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x12, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x32, 0x9d, 0x01, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x12, 0x43, 0x0a, 0x08, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x49, 0x44, 0x73, 0x12, 0x1a,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x54, 0x78, 0x48, 0x61, 0x73, 0x68, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x49, 0x44, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x70, 0x6f,
	0x73, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1d, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x6c, 0x6f, 0x58, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2d, 0x4c,
	0x61, 0x62, 0x73, 0x2f, 0x62, 0x73, 0x63, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x6d,
	0x69, 0x6e, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blxr_mev_proto_rawDescOnce sync.Once
	file_blxr_mev_proto_rawDescData = file_blxr_mev_proto_rawDesc
)

func file_blxr_mev_proto_rawDescGZIP() []byte {
	file_blxr_mev_proto_rawDescOnce.Do(func() {
		file_blxr_mev_proto_rawDescData = protoimpl.X.CompressGZIP(file_blxr_mev_proto_rawDescData)
	})
	return file_blxr_mev_proto_rawDescData
}

var file_blxr_mev_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_blxr_mev_proto_goTypes = []interface{}{
	(*TxHashListRequest)(nil),    // 0: gateway.TxHashListRequest
	(*ShortIDListReply)(nil),     // 1: gateway.ShortIDListReply
	(*ProposedBlockRequest)(nil), // 2: gateway.ProposedBlockRequest
	(*CompressTx)(nil),           // 3: gateway.compressTx
	(*ProposedBlockReply)(nil),   // 4: gateway.ProposedBlockReply
}
var file_blxr_mev_proto_depIdxs = []int32{
	3, // 0: gateway.ProposedBlockRequest.payload:type_name -> gateway.compressTx
	0, // 1: gateway.Gateway.ShortIDs:input_type -> gateway.TxHashListRequest
	2, // 2: gateway.Gateway.ProposedBlock:input_type -> gateway.ProposedBlockRequest
	1, // 3: gateway.Gateway.ShortIDs:output_type -> gateway.ShortIDListReply
	4, // 4: gateway.Gateway.ProposedBlock:output_type -> gateway.ProposedBlockReply
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_blxr_mev_proto_init() }
func file_blxr_mev_proto_init() {
	if File_blxr_mev_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blxr_mev_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxHashListRequest); i {
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
		file_blxr_mev_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortIDListReply); i {
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
		file_blxr_mev_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProposedBlockRequest); i {
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
		file_blxr_mev_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompressTx); i {
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
		file_blxr_mev_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProposedBlockReply); i {
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
			RawDescriptor: file_blxr_mev_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_blxr_mev_proto_goTypes,
		DependencyIndexes: file_blxr_mev_proto_depIdxs,
		MessageInfos:      file_blxr_mev_proto_msgTypes,
	}.Build()
	File_blxr_mev_proto = out.File
	file_blxr_mev_proto_rawDesc = nil
	file_blxr_mev_proto_goTypes = nil
	file_blxr_mev_proto_depIdxs = nil
}
