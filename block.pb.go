// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.1
// source: block.proto

package block_proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ReturnCode int32

const (
	ReturnCode_OK                    ReturnCode = 0
	ReturnCode_INTERNAL_ERROR        ReturnCode = 1
	ReturnCode_REQUEST_PARAM_ILLEGAL ReturnCode = 2
	ReturnCode_OTHER_ERROR           ReturnCode = 99
)

// Enum value maps for ReturnCode.
var (
	ReturnCode_name = map[int32]string{
		0:  "OK",
		1:  "INTERNAL_ERROR",
		2:  "REQUEST_PARAM_ILLEGAL",
		99: "OTHER_ERROR",
	}
	ReturnCode_value = map[string]int32{
		"OK":                    0,
		"INTERNAL_ERROR":        1,
		"REQUEST_PARAM_ILLEGAL": 2,
		"OTHER_ERROR":           99,
	}
)

func (x ReturnCode) Enum() *ReturnCode {
	p := new(ReturnCode)
	*p = x
	return p
}

func (x ReturnCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReturnCode) Descriptor() protoreflect.EnumDescriptor {
	return file_block_proto_enumTypes[0].Descriptor()
}

func (ReturnCode) Type() protoreflect.EnumType {
	return &file_block_proto_enumTypes[0]
}

func (x ReturnCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReturnCode.Descriptor instead.
func (ReturnCode) EnumDescriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{0}
}

type CountTransactionByBlockNumberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockNumber uint64 `protobuf:"varint,1,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
}

func (x *CountTransactionByBlockNumberRequest) Reset() {
	*x = CountTransactionByBlockNumberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountTransactionByBlockNumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountTransactionByBlockNumberRequest) ProtoMessage() {}

func (x *CountTransactionByBlockNumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountTransactionByBlockNumberRequest.ProtoReflect.Descriptor instead.
func (*CountTransactionByBlockNumberRequest) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{0}
}

func (x *CountTransactionByBlockNumberRequest) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

type CountTransactionByBlockNumberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code             ReturnCode `protobuf:"varint,1,opt,name=code,proto3,enum=block_proto.ReturnCode" json:"code,omitempty"`
	Message          string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	TransactionCount uint64     `protobuf:"varint,3,opt,name=transaction_count,json=transactionCount,proto3" json:"transaction_count,omitempty"`
}

func (x *CountTransactionByBlockNumberResponse) Reset() {
	*x = CountTransactionByBlockNumberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountTransactionByBlockNumberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountTransactionByBlockNumberResponse) ProtoMessage() {}

func (x *CountTransactionByBlockNumberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountTransactionByBlockNumberResponse.ProtoReflect.Descriptor instead.
func (*CountTransactionByBlockNumberResponse) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{1}
}

func (x *CountTransactionByBlockNumberResponse) GetCode() ReturnCode {
	if x != nil {
		return x.Code
	}
	return ReturnCode_OK
}

func (x *CountTransactionByBlockNumberResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CountTransactionByBlockNumberResponse) GetTransactionCount() uint64 {
	if x != nil {
		return x.TransactionCount
	}
	return 0
}

type GetBlocksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit uint32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetBlocksRequest) Reset() {
	*x = GetBlocksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlocksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlocksRequest) ProtoMessage() {}

func (x *GetBlocksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlocksRequest.ProtoReflect.Descriptor instead.
func (*GetBlocksRequest) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{2}
}

func (x *GetBlocksRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetBlocksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code           ReturnCode                     `protobuf:"varint,1,opt,name=code,proto3,enum=block_proto.ReturnCode" json:"code,omitempty"`
	Message        string                         `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	BlockItemArray []*GetBlocksResponse_BlockItem `protobuf:"bytes,3,rep,name=block_item_array,json=blockItemArray,proto3" json:"block_item_array,omitempty"`
}

func (x *GetBlocksResponse) Reset() {
	*x = GetBlocksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlocksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlocksResponse) ProtoMessage() {}

func (x *GetBlocksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlocksResponse.ProtoReflect.Descriptor instead.
func (*GetBlocksResponse) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{3}
}

func (x *GetBlocksResponse) GetCode() ReturnCode {
	if x != nil {
		return x.Code
	}
	return ReturnCode_OK
}

func (x *GetBlocksResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetBlocksResponse) GetBlockItemArray() []*GetBlocksResponse_BlockItem {
	if x != nil {
		return x.BlockItemArray
	}
	return nil
}

type CountByTimeRangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BeginTimestamp uint64 `protobuf:"varint,1,opt,name=begin_timestamp,json=beginTimestamp,proto3" json:"begin_timestamp,omitempty"`
	EndTimestamp   uint64 `protobuf:"varint,2,opt,name=end_timestamp,json=endTimestamp,proto3" json:"end_timestamp,omitempty"`
}

func (x *CountByTimeRangeRequest) Reset() {
	*x = CountByTimeRangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountByTimeRangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountByTimeRangeRequest) ProtoMessage() {}

func (x *CountByTimeRangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountByTimeRangeRequest.ProtoReflect.Descriptor instead.
func (*CountByTimeRangeRequest) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{4}
}

func (x *CountByTimeRangeRequest) GetBeginTimestamp() uint64 {
	if x != nil {
		return x.BeginTimestamp
	}
	return 0
}

func (x *CountByTimeRangeRequest) GetEndTimestamp() uint64 {
	if x != nil {
		return x.EndTimestamp
	}
	return 0
}

type CountByTimeRangeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    ReturnCode `protobuf:"varint,1,opt,name=code,proto3,enum=block_proto.ReturnCode" json:"code,omitempty"`
	Message string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Count   uint64     `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *CountByTimeRangeResponse) Reset() {
	*x = CountByTimeRangeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountByTimeRangeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountByTimeRangeResponse) ProtoMessage() {}

func (x *CountByTimeRangeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountByTimeRangeResponse.ProtoReflect.Descriptor instead.
func (*CountByTimeRangeResponse) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{5}
}

func (x *CountByTimeRangeResponse) GetCode() ReturnCode {
	if x != nil {
		return x.Code
	}
	return ReturnCode_OK
}

func (x *CountByTimeRangeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CountByTimeRangeResponse) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type GetBlocksResponse_BlockItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockNumber uint64 `protobuf:"varint,1,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	BlockTime   uint64 `protobuf:"varint,2,opt,name=block_time,json=blockTime,proto3" json:"block_time,omitempty"`
	BlockHash   string `protobuf:"bytes,3,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
	ParentHash  string `protobuf:"bytes,4,opt,name=parent_hash,json=parentHash,proto3" json:"parent_hash,omitempty"`
	Coinbase    string `protobuf:"bytes,5,opt,name=coinbase,proto3" json:"coinbase,omitempty"`
	GasLimit    uint64 `protobuf:"varint,6,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
}

func (x *GetBlocksResponse_BlockItem) Reset() {
	*x = GetBlocksResponse_BlockItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlocksResponse_BlockItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlocksResponse_BlockItem) ProtoMessage() {}

func (x *GetBlocksResponse_BlockItem) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlocksResponse_BlockItem.ProtoReflect.Descriptor instead.
func (*GetBlocksResponse_BlockItem) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{3, 0}
}

func (x *GetBlocksResponse_BlockItem) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *GetBlocksResponse_BlockItem) GetBlockTime() uint64 {
	if x != nil {
		return x.BlockTime
	}
	return 0
}

func (x *GetBlocksResponse_BlockItem) GetBlockHash() string {
	if x != nil {
		return x.BlockHash
	}
	return ""
}

func (x *GetBlocksResponse_BlockItem) GetParentHash() string {
	if x != nil {
		return x.ParentHash
	}
	return ""
}

func (x *GetBlocksResponse_BlockItem) GetCoinbase() string {
	if x != nil {
		return x.Coinbase
	}
	return ""
}

func (x *GetBlocksResponse_BlockItem) GetGasLimit() uint64 {
	if x != nil {
		return x.GasLimit
	}
	return 0
}

var File_block_proto protoreflect.FileDescriptor

var file_block_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a, 0x24, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x9b, 0x01, 0x0a, 0x25, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2b, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x28, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0xf7, 0x02,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x17, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x52, 0x0a, 0x10, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x61, 0x72, 0x72, 0x61, 0x79, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0e,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x74, 0x65, 0x6d, 0x41, 0x72, 0x72, 0x61, 0x79, 0x1a, 0xc6,
	0x01, 0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x21, 0x0a, 0x0c,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61,
	0x73, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x67,
	0x61, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x67, 0x0a, 0x17, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x42, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x62, 0x65, 0x67,
	0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x23, 0x0a, 0x0d, 0x65,
	0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0c, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x22, 0x77, 0x0a, 0x18, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x52,
	0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2a, 0x54, 0x0a, 0x0a, 0x52, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12,
	0x12, 0x0a, 0x0e, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x50,
	0x41, 0x52, 0x41, 0x4d, 0x5f, 0x49, 0x4c, 0x4c, 0x45, 0x47, 0x41, 0x4c, 0x10, 0x02, 0x12, 0x0f,
	0x0a, 0x0b, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x63, 0x32,
	0xae, 0x04, 0x0a, 0x10, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x6c, 0x0a, 0x1b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x12, 0x24, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x54,
	0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x4c, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12,
	0x1d, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x88, 0x01, 0x0a, 0x1d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x31, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x79, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x68, 0x0a, 0x17, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x54, 0x69, 0x6d,
	0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x24, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x54, 0x69, 0x6d, 0x65,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x42, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x69, 0x0a, 0x18, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x42, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x12, 0x24, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x54, 0x69, 0x6d,
	0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_block_proto_rawDescOnce sync.Once
	file_block_proto_rawDescData = file_block_proto_rawDesc
)

func file_block_proto_rawDescGZIP() []byte {
	file_block_proto_rawDescOnce.Do(func() {
		file_block_proto_rawDescData = protoimpl.X.CompressGZIP(file_block_proto_rawDescData)
	})
	return file_block_proto_rawDescData
}

var file_block_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_block_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_block_proto_goTypes = []interface{}{
	(ReturnCode)(0), // 0: block_proto.ReturnCode
	(*CountTransactionByBlockNumberRequest)(nil),  // 1: block_proto.CountTransactionByBlockNumberRequest
	(*CountTransactionByBlockNumberResponse)(nil), // 2: block_proto.CountTransactionByBlockNumberResponse
	(*GetBlocksRequest)(nil),                      // 3: block_proto.GetBlocksRequest
	(*GetBlocksResponse)(nil),                     // 4: block_proto.GetBlocksResponse
	(*CountByTimeRangeRequest)(nil),               // 5: block_proto.CountByTimeRangeRequest
	(*CountByTimeRangeResponse)(nil),              // 6: block_proto.CountByTimeRangeResponse
	(*GetBlocksResponse_BlockItem)(nil),           // 7: block_proto.GetBlocksResponse.BlockItem
}
var file_block_proto_depIdxs = []int32{
	0, // 0: block_proto.CountTransactionByBlockNumberResponse.code:type_name -> block_proto.ReturnCode
	0, // 1: block_proto.GetBlocksResponse.code:type_name -> block_proto.ReturnCode
	7, // 2: block_proto.GetBlocksResponse.block_item_array:type_name -> block_proto.GetBlocksResponse.BlockItem
	0, // 3: block_proto.CountByTimeRangeResponse.code:type_name -> block_proto.ReturnCode
	5, // 4: block_proto.BlockGrpcService.CountTransactionByTimeRange:input_type -> block_proto.CountByTimeRangeRequest
	3, // 5: block_proto.BlockGrpcService.GetBlocks:input_type -> block_proto.GetBlocksRequest
	1, // 6: block_proto.BlockGrpcService.CountTransactionByBlockNumber:input_type -> block_proto.CountTransactionByBlockNumberRequest
	5, // 7: block_proto.BlockGrpcService.CountAccountByTimeRange:input_type -> block_proto.CountByTimeRangeRequest
	5, // 8: block_proto.BlockGrpcService.CountContractByTimeRange:input_type -> block_proto.CountByTimeRangeRequest
	6, // 9: block_proto.BlockGrpcService.CountTransactionByTimeRange:output_type -> block_proto.CountByTimeRangeResponse
	4, // 10: block_proto.BlockGrpcService.GetBlocks:output_type -> block_proto.GetBlocksResponse
	2, // 11: block_proto.BlockGrpcService.CountTransactionByBlockNumber:output_type -> block_proto.CountTransactionByBlockNumberResponse
	6, // 12: block_proto.BlockGrpcService.CountAccountByTimeRange:output_type -> block_proto.CountByTimeRangeResponse
	6, // 13: block_proto.BlockGrpcService.CountContractByTimeRange:output_type -> block_proto.CountByTimeRangeResponse
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_block_proto_init() }
func file_block_proto_init() {
	if File_block_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_block_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountTransactionByBlockNumberRequest); i {
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
		file_block_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountTransactionByBlockNumberResponse); i {
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
		file_block_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlocksRequest); i {
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
		file_block_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlocksResponse); i {
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
		file_block_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountByTimeRangeRequest); i {
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
		file_block_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountByTimeRangeResponse); i {
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
		file_block_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlocksResponse_BlockItem); i {
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
			RawDescriptor: file_block_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_block_proto_goTypes,
		DependencyIndexes: file_block_proto_depIdxs,
		EnumInfos:         file_block_proto_enumTypes,
		MessageInfos:      file_block_proto_msgTypes,
	}.Build()
	File_block_proto = out.File
	file_block_proto_rawDesc = nil
	file_block_proto_goTypes = nil
	file_block_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BlockGrpcServiceClient is the client API for BlockGrpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlockGrpcServiceClient interface {
	CountTransactionByTimeRange(ctx context.Context, in *CountByTimeRangeRequest, opts ...grpc.CallOption) (*CountByTimeRangeResponse, error)
	GetBlocks(ctx context.Context, in *GetBlocksRequest, opts ...grpc.CallOption) (*GetBlocksResponse, error)
	CountTransactionByBlockNumber(ctx context.Context, in *CountTransactionByBlockNumberRequest, opts ...grpc.CallOption) (*CountTransactionByBlockNumberResponse, error)
	CountAccountByTimeRange(ctx context.Context, in *CountByTimeRangeRequest, opts ...grpc.CallOption) (*CountByTimeRangeResponse, error)
	CountContractByTimeRange(ctx context.Context, in *CountByTimeRangeRequest, opts ...grpc.CallOption) (*CountByTimeRangeResponse, error)
}

type blockGrpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlockGrpcServiceClient(cc grpc.ClientConnInterface) BlockGrpcServiceClient {
	return &blockGrpcServiceClient{cc}
}

func (c *blockGrpcServiceClient) CountTransactionByTimeRange(ctx context.Context, in *CountByTimeRangeRequest, opts ...grpc.CallOption) (*CountByTimeRangeResponse, error) {
	out := new(CountByTimeRangeResponse)
	err := c.cc.Invoke(ctx, "/block_proto.BlockGrpcService/CountTransactionByTimeRange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockGrpcServiceClient) GetBlocks(ctx context.Context, in *GetBlocksRequest, opts ...grpc.CallOption) (*GetBlocksResponse, error) {
	out := new(GetBlocksResponse)
	err := c.cc.Invoke(ctx, "/block_proto.BlockGrpcService/GetBlocks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockGrpcServiceClient) CountTransactionByBlockNumber(ctx context.Context, in *CountTransactionByBlockNumberRequest, opts ...grpc.CallOption) (*CountTransactionByBlockNumberResponse, error) {
	out := new(CountTransactionByBlockNumberResponse)
	err := c.cc.Invoke(ctx, "/block_proto.BlockGrpcService/CountTransactionByBlockNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockGrpcServiceClient) CountAccountByTimeRange(ctx context.Context, in *CountByTimeRangeRequest, opts ...grpc.CallOption) (*CountByTimeRangeResponse, error) {
	out := new(CountByTimeRangeResponse)
	err := c.cc.Invoke(ctx, "/block_proto.BlockGrpcService/CountAccountByTimeRange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockGrpcServiceClient) CountContractByTimeRange(ctx context.Context, in *CountByTimeRangeRequest, opts ...grpc.CallOption) (*CountByTimeRangeResponse, error) {
	out := new(CountByTimeRangeResponse)
	err := c.cc.Invoke(ctx, "/block_proto.BlockGrpcService/CountContractByTimeRange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlockGrpcServiceServer is the server API for BlockGrpcService service.
type BlockGrpcServiceServer interface {
	CountTransactionByTimeRange(context.Context, *CountByTimeRangeRequest) (*CountByTimeRangeResponse, error)
	GetBlocks(context.Context, *GetBlocksRequest) (*GetBlocksResponse, error)
	CountTransactionByBlockNumber(context.Context, *CountTransactionByBlockNumberRequest) (*CountTransactionByBlockNumberResponse, error)
	CountAccountByTimeRange(context.Context, *CountByTimeRangeRequest) (*CountByTimeRangeResponse, error)
	CountContractByTimeRange(context.Context, *CountByTimeRangeRequest) (*CountByTimeRangeResponse, error)
}

// UnimplementedBlockGrpcServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBlockGrpcServiceServer struct {
}

func (*UnimplementedBlockGrpcServiceServer) CountTransactionByTimeRange(context.Context, *CountByTimeRangeRequest) (*CountByTimeRangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountTransactionByTimeRange not implemented")
}
func (*UnimplementedBlockGrpcServiceServer) GetBlocks(context.Context, *GetBlocksRequest) (*GetBlocksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlocks not implemented")
}
func (*UnimplementedBlockGrpcServiceServer) CountTransactionByBlockNumber(context.Context, *CountTransactionByBlockNumberRequest) (*CountTransactionByBlockNumberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountTransactionByBlockNumber not implemented")
}
func (*UnimplementedBlockGrpcServiceServer) CountAccountByTimeRange(context.Context, *CountByTimeRangeRequest) (*CountByTimeRangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountAccountByTimeRange not implemented")
}
func (*UnimplementedBlockGrpcServiceServer) CountContractByTimeRange(context.Context, *CountByTimeRangeRequest) (*CountByTimeRangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountContractByTimeRange not implemented")
}

func RegisterBlockGrpcServiceServer(s *grpc.Server, srv BlockGrpcServiceServer) {
	s.RegisterService(&_BlockGrpcService_serviceDesc, srv)
}

func _BlockGrpcService_CountTransactionByTimeRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountByTimeRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockGrpcServiceServer).CountTransactionByTimeRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/block_proto.BlockGrpcService/CountTransactionByTimeRange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockGrpcServiceServer).CountTransactionByTimeRange(ctx, req.(*CountByTimeRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockGrpcService_GetBlocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlocksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockGrpcServiceServer).GetBlocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/block_proto.BlockGrpcService/GetBlocks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockGrpcServiceServer).GetBlocks(ctx, req.(*GetBlocksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockGrpcService_CountTransactionByBlockNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountTransactionByBlockNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockGrpcServiceServer).CountTransactionByBlockNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/block_proto.BlockGrpcService/CountTransactionByBlockNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockGrpcServiceServer).CountTransactionByBlockNumber(ctx, req.(*CountTransactionByBlockNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockGrpcService_CountAccountByTimeRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountByTimeRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockGrpcServiceServer).CountAccountByTimeRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/block_proto.BlockGrpcService/CountAccountByTimeRange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockGrpcServiceServer).CountAccountByTimeRange(ctx, req.(*CountByTimeRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockGrpcService_CountContractByTimeRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountByTimeRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockGrpcServiceServer).CountContractByTimeRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/block_proto.BlockGrpcService/CountContractByTimeRange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockGrpcServiceServer).CountContractByTimeRange(ctx, req.(*CountByTimeRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BlockGrpcService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "block_proto.BlockGrpcService",
	HandlerType: (*BlockGrpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CountTransactionByTimeRange",
			Handler:    _BlockGrpcService_CountTransactionByTimeRange_Handler,
		},
		{
			MethodName: "GetBlocks",
			Handler:    _BlockGrpcService_GetBlocks_Handler,
		},
		{
			MethodName: "CountTransactionByBlockNumber",
			Handler:    _BlockGrpcService_CountTransactionByBlockNumber_Handler,
		},
		{
			MethodName: "CountAccountByTimeRange",
			Handler:    _BlockGrpcService_CountAccountByTimeRange_Handler,
		},
		{
			MethodName: "CountContractByTimeRange",
			Handler:    _BlockGrpcService_CountContractByTimeRange_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "block.proto",
}
