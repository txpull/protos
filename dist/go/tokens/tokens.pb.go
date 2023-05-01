// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: tokens.proto

package tokens

import (
	any1 "github.com/golang/protobuf/ptypes/any"
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

type TokenStatus int32

const (
	TokenStatus_ANY        TokenStatus = 0
	TokenStatus_QUEUED     TokenStatus = 1
	TokenStatus_PROCESSING TokenStatus = 2
	TokenStatus_PROCESSED  TokenStatus = 3
	TokenStatus_FAILED     TokenStatus = 4
)

// Enum value maps for TokenStatus.
var (
	TokenStatus_name = map[int32]string{
		0: "ANY",
		1: "QUEUED",
		2: "PROCESSING",
		3: "PROCESSED",
		4: "FAILED",
	}
	TokenStatus_value = map[string]int32{
		"ANY":        0,
		"QUEUED":     1,
		"PROCESSING": 2,
		"PROCESSED":  3,
		"FAILED":     4,
	}
)

func (x TokenStatus) Enum() *TokenStatus {
	p := new(TokenStatus)
	*p = x
	return p
}

func (x TokenStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TokenStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_tokens_proto_enumTypes[0].Descriptor()
}

func (TokenStatus) Type() protoreflect.EnumType {
	return &file_tokens_proto_enumTypes[0]
}

func (x TokenStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TokenStatus.Descriptor instead.
func (TokenStatus) EnumDescriptor() ([]byte, []int) {
	return file_tokens_proto_rawDescGZIP(), []int{0}
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tokens_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_tokens_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_tokens_proto_rawDescGZIP(), []int{0}
}

// The `Status` type defines a logical error model that is suitable for
// different programming environments, including REST APIs and RPC APIs.
type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A simple error code that can be easily handled by the client. The
	// actual error code is defined by `google.rpc.Code`.
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// A developer-facing human-readable error message in English. It should
	// both explain the error and offer an actionable resolution to it.
	Reason string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	// A developer-facing human-readable error message in English. It should
	// both explain the error and offer an actionable resolution to it.
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	// Additional error information that the client code can use to handle
	// the error, such as retry info or a help link.
	Details []*any1.Any `protobuf:"bytes,4,rep,name=details,proto3" json:"details,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tokens_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_tokens_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_tokens_proto_rawDescGZIP(), []int{1}
}

func (x *Status) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Status) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *Status) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Status) GetDetails() []*any1.Any {
	if x != nil {
		return x.Details
	}
	return nil
}

type FilterTokensRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartBlock uint64 `protobuf:"varint,1,opt,name=start_block,json=startBlock,proto3" json:"start_block,omitempty"`
	EndBlock   uint64 `protobuf:"varint,2,opt,name=end_block,json=endBlock,proto3" json:"end_block,omitempty"`
	Page       uint64 `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	PageSize   uint64 `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Status     uint32 `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *FilterTokensRequest) Reset() {
	*x = FilterTokensRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tokens_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterTokensRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterTokensRequest) ProtoMessage() {}

func (x *FilterTokensRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tokens_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterTokensRequest.ProtoReflect.Descriptor instead.
func (*FilterTokensRequest) Descriptor() ([]byte, []int) {
	return file_tokens_proto_rawDescGZIP(), []int{2}
}

func (x *FilterTokensRequest) GetStartBlock() uint64 {
	if x != nil {
		return x.StartBlock
	}
	return 0
}

func (x *FilterTokensRequest) GetEndBlock() uint64 {
	if x != nil {
		return x.EndBlock
	}
	return 0
}

func (x *FilterTokensRequest) GetPage() uint64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *FilterTokensRequest) GetPageSize() uint64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *FilterTokensRequest) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type FilterTokensResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Token  []*Token `protobuf:"bytes,2,rep,name=token,proto3" json:"token,omitempty"`
}

func (x *FilterTokensResponse) Reset() {
	*x = FilterTokensResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tokens_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterTokensResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterTokensResponse) ProtoMessage() {}

func (x *FilterTokensResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tokens_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterTokensResponse.ProtoReflect.Descriptor instead.
func (*FilterTokensResponse) Descriptor() ([]byte, []int) {
	return file_tokens_proto_rawDescGZIP(), []int{3}
}

func (x *FilterTokensResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *FilterTokensResponse) GetToken() []*Token {
	if x != nil {
		return x.Token
	}
	return nil
}

type GetTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenId uint64 `protobuf:"varint,1,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *GetTokenRequest) Reset() {
	*x = GetTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tokens_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTokenRequest) ProtoMessage() {}

func (x *GetTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tokens_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTokenRequest.ProtoReflect.Descriptor instead.
func (*GetTokenRequest) Descriptor() ([]byte, []int) {
	return file_tokens_proto_rawDescGZIP(), []int{4}
}

func (x *GetTokenRequest) GetTokenId() uint64 {
	if x != nil {
		return x.TokenId
	}
	return 0
}

func (x *GetTokenRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type GetTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Block *Token `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
}

func (x *GetTokenResponse) Reset() {
	*x = GetTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tokens_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTokenResponse) ProtoMessage() {}

func (x *GetTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tokens_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTokenResponse.ProtoReflect.Descriptor instead.
func (*GetTokenResponse) Descriptor() ([]byte, []int) {
	return file_tokens_proto_rawDescGZIP(), []int{5}
}

func (x *GetTokenResponse) GetBlock() *Token {
	if x != nil {
		return x.Block
	}
	return nil
}

type SubscribeTokensRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartBlock uint64      `protobuf:"varint,1,opt,name=start_block,json=startBlock,proto3" json:"start_block,omitempty"`
	Status     TokenStatus `protobuf:"varint,2,opt,name=status,proto3,enum=txpull.v1.tokens.TokenStatus" json:"status,omitempty"`
}

func (x *SubscribeTokensRequest) Reset() {
	*x = SubscribeTokensRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tokens_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeTokensRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeTokensRequest) ProtoMessage() {}

func (x *SubscribeTokensRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tokens_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeTokensRequest.ProtoReflect.Descriptor instead.
func (*SubscribeTokensRequest) Descriptor() ([]byte, []int) {
	return file_tokens_proto_rawDescGZIP(), []int{6}
}

func (x *SubscribeTokensRequest) GetStartBlock() uint64 {
	if x != nil {
		return x.StartBlock
	}
	return 0
}

func (x *SubscribeTokensRequest) GetStatus() TokenStatus {
	if x != nil {
		return x.Status
	}
	return TokenStatus_ANY
}

var File_tokens_proto protoreflect.FileDescriptor

var file_tokens_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x74, 0x78, 0x70, 0x75, 0x6c, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x7e, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x22, 0x9c, 0x01, 0x0a, 0x13, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1b, 0x0a, 0x09, 0x65,
	0x6e, 0x64, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08,
	0x65, 0x6e, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x77, 0x0a, 0x14, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x78, 0x70, 0x75,
	0x6c, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2d, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x78, 0x70,
	0x75, 0x6c, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x46, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x22, 0x41, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x78, 0x70, 0x75, 0x6c, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x05,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x70, 0x0a, 0x16, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x12, 0x35, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1d, 0x2e, 0x74, 0x78, 0x70, 0x75, 0x6c, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x4d, 0x0a, 0x0b, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x4e, 0x59, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x51, 0x55, 0x45, 0x55, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x50,
	0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x50,
	0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41,
	0x49, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x32, 0xb6, 0x02, 0x0a, 0x06, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x12, 0x6b, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x25, 0x2e, 0x74, 0x78,
	0x70, 0x75, 0x6c, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2e, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x26, 0x2e, 0x74, 0x78, 0x70, 0x75, 0x6c, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x6b,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x21, 0x2e, 0x74, 0x78, 0x70, 0x75, 0x6c, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x74, 0x78, 0x70, 0x75, 0x6c,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x2f, 0x7b, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x52, 0x0a, 0x09, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x28, 0x2e, 0x74, 0x78, 0x70, 0x75, 0x6c,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2e, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x74, 0x78, 0x70, 0x75, 0x6c, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x00, 0x30, 0x01, 0x42,
	0x1d, 0x5a, 0x1b, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x2f, 0x67, 0x6f,
	0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x3b, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tokens_proto_rawDescOnce sync.Once
	file_tokens_proto_rawDescData = file_tokens_proto_rawDesc
)

func file_tokens_proto_rawDescGZIP() []byte {
	file_tokens_proto_rawDescOnce.Do(func() {
		file_tokens_proto_rawDescData = protoimpl.X.CompressGZIP(file_tokens_proto_rawDescData)
	})
	return file_tokens_proto_rawDescData
}

var file_tokens_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tokens_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_tokens_proto_goTypes = []interface{}{
	(TokenStatus)(0),               // 0: txpull.v1.tokens.TokenStatus
	(*Token)(nil),                  // 1: txpull.v1.tokens.Token
	(*Status)(nil),                 // 2: txpull.v1.tokens.Status
	(*FilterTokensRequest)(nil),    // 3: txpull.v1.tokens.FilterTokensRequest
	(*FilterTokensResponse)(nil),   // 4: txpull.v1.tokens.FilterTokensResponse
	(*GetTokenRequest)(nil),        // 5: txpull.v1.tokens.GetTokenRequest
	(*GetTokenResponse)(nil),       // 6: txpull.v1.tokens.GetTokenResponse
	(*SubscribeTokensRequest)(nil), // 7: txpull.v1.tokens.SubscribeTokensRequest
	(*any1.Any)(nil),               // 8: google.protobuf.Any
}
var file_tokens_proto_depIdxs = []int32{
	8, // 0: txpull.v1.tokens.Status.details:type_name -> google.protobuf.Any
	2, // 1: txpull.v1.tokens.FilterTokensResponse.status:type_name -> txpull.v1.tokens.Status
	1, // 2: txpull.v1.tokens.FilterTokensResponse.token:type_name -> txpull.v1.tokens.Token
	1, // 3: txpull.v1.tokens.GetTokenResponse.block:type_name -> txpull.v1.tokens.Token
	0, // 4: txpull.v1.tokens.SubscribeTokensRequest.status:type_name -> txpull.v1.tokens.TokenStatus
	3, // 5: txpull.v1.tokens.Tokens.Filter:input_type -> txpull.v1.tokens.FilterTokensRequest
	5, // 6: txpull.v1.tokens.Tokens.Get:input_type -> txpull.v1.tokens.GetTokenRequest
	7, // 7: txpull.v1.tokens.Tokens.Subscribe:input_type -> txpull.v1.tokens.SubscribeTokensRequest
	4, // 8: txpull.v1.tokens.Tokens.Filter:output_type -> txpull.v1.tokens.FilterTokensResponse
	6, // 9: txpull.v1.tokens.Tokens.Get:output_type -> txpull.v1.tokens.GetTokenResponse
	1, // 10: txpull.v1.tokens.Tokens.Subscribe:output_type -> txpull.v1.tokens.Token
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_tokens_proto_init() }
func file_tokens_proto_init() {
	if File_tokens_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tokens_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
		file_tokens_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_tokens_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterTokensRequest); i {
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
		file_tokens_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterTokensResponse); i {
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
		file_tokens_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTokenRequest); i {
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
		file_tokens_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTokenResponse); i {
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
		file_tokens_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeTokensRequest); i {
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
			RawDescriptor: file_tokens_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tokens_proto_goTypes,
		DependencyIndexes: file_tokens_proto_depIdxs,
		EnumInfos:         file_tokens_proto_enumTypes,
		MessageInfos:      file_tokens_proto_msgTypes,
	}.Build()
	File_tokens_proto = out.File
	file_tokens_proto_rawDesc = nil
	file_tokens_proto_goTypes = nil
	file_tokens_proto_depIdxs = nil
}
