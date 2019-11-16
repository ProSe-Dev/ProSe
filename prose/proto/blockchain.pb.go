// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/blockchain.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AddBlockRequest struct {
	Data                 *BlockData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *AddBlockRequest) Reset()         { *m = AddBlockRequest{} }
func (m *AddBlockRequest) String() string { return proto.CompactTextString(m) }
func (*AddBlockRequest) ProtoMessage()    {}
func (*AddBlockRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a884755cc01e392a, []int{0}
}

func (m *AddBlockRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddBlockRequest.Unmarshal(m, b)
}
func (m *AddBlockRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddBlockRequest.Marshal(b, m, deterministic)
}
func (m *AddBlockRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddBlockRequest.Merge(m, src)
}
func (m *AddBlockRequest) XXX_Size() int {
	return xxx_messageInfo_AddBlockRequest.Size(m)
}
func (m *AddBlockRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddBlockRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddBlockRequest proto.InternalMessageInfo

func (m *AddBlockRequest) GetData() *BlockData {
	if m != nil {
		return m.Data
	}
	return nil
}

type AddBlockResponse struct {
	ACK                  bool     `protobuf:"varint,1,opt,name=ACK,proto3" json:"ACK,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddBlockResponse) Reset()         { *m = AddBlockResponse{} }
func (m *AddBlockResponse) String() string { return proto.CompactTextString(m) }
func (*AddBlockResponse) ProtoMessage()    {}
func (*AddBlockResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a884755cc01e392a, []int{1}
}

func (m *AddBlockResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddBlockResponse.Unmarshal(m, b)
}
func (m *AddBlockResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddBlockResponse.Marshal(b, m, deterministic)
}
func (m *AddBlockResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddBlockResponse.Merge(m, src)
}
func (m *AddBlockResponse) XXX_Size() int {
	return xxx_messageInfo_AddBlockResponse.Size(m)
}
func (m *AddBlockResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddBlockResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddBlockResponse proto.InternalMessageInfo

func (m *AddBlockResponse) GetACK() bool {
	if m != nil {
		return m.ACK
	}
	return false
}

type BlockData struct {
	Timestamp            string            `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Author               string            `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	Identity             string            `protobuf:"bytes,3,opt,name=identity,proto3" json:"identity,omitempty"`
	ProjectID            string            `protobuf:"bytes,4,opt,name=projectID,proto3" json:"projectID,omitempty"`
	CommitHash           string            `protobuf:"bytes,5,opt,name=commitHash,proto3" json:"commitHash,omitempty"`
	FileHashes           map[string]string `protobuf:"bytes,6,rep,name=fileHashes,proto3" json:"fileHashes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *BlockData) Reset()         { *m = BlockData{} }
func (m *BlockData) String() string { return proto.CompactTextString(m) }
func (*BlockData) ProtoMessage()    {}
func (*BlockData) Descriptor() ([]byte, []int) {
	return fileDescriptor_a884755cc01e392a, []int{2}
}

func (m *BlockData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockData.Unmarshal(m, b)
}
func (m *BlockData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockData.Marshal(b, m, deterministic)
}
func (m *BlockData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockData.Merge(m, src)
}
func (m *BlockData) XXX_Size() int {
	return xxx_messageInfo_BlockData.Size(m)
}
func (m *BlockData) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockData.DiscardUnknown(m)
}

var xxx_messageInfo_BlockData proto.InternalMessageInfo

func (m *BlockData) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *BlockData) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *BlockData) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

func (m *BlockData) GetProjectID() string {
	if m != nil {
		return m.ProjectID
	}
	return ""
}

func (m *BlockData) GetCommitHash() string {
	if m != nil {
		return m.CommitHash
	}
	return ""
}

func (m *BlockData) GetFileHashes() map[string]string {
	if m != nil {
		return m.FileHashes
	}
	return nil
}

type Block struct {
	PrevBlockHash        string     `protobuf:"bytes,1,opt,name=prevBlockHash,proto3" json:"prevBlockHash,omitempty"`
	Data                 *BlockData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Hash                 string     `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_a884755cc01e392a, []int{3}
}

func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (m *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(m, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetPrevBlockHash() string {
	if m != nil {
		return m.PrevBlockHash
	}
	return ""
}

func (m *Block) GetData() *BlockData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Block) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type GetBlockchainRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlockchainRequest) Reset()         { *m = GetBlockchainRequest{} }
func (m *GetBlockchainRequest) String() string { return proto.CompactTextString(m) }
func (*GetBlockchainRequest) ProtoMessage()    {}
func (*GetBlockchainRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a884755cc01e392a, []int{4}
}

func (m *GetBlockchainRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlockchainRequest.Unmarshal(m, b)
}
func (m *GetBlockchainRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlockchainRequest.Marshal(b, m, deterministic)
}
func (m *GetBlockchainRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockchainRequest.Merge(m, src)
}
func (m *GetBlockchainRequest) XXX_Size() int {
	return xxx_messageInfo_GetBlockchainRequest.Size(m)
}
func (m *GetBlockchainRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockchainRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockchainRequest proto.InternalMessageInfo

type GetBlockchainResponse struct {
	Blocks               []*Block `protobuf:"bytes,1,rep,name=blocks,proto3" json:"blocks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlockchainResponse) Reset()         { *m = GetBlockchainResponse{} }
func (m *GetBlockchainResponse) String() string { return proto.CompactTextString(m) }
func (*GetBlockchainResponse) ProtoMessage()    {}
func (*GetBlockchainResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a884755cc01e392a, []int{5}
}

func (m *GetBlockchainResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlockchainResponse.Unmarshal(m, b)
}
func (m *GetBlockchainResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlockchainResponse.Marshal(b, m, deterministic)
}
func (m *GetBlockchainResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockchainResponse.Merge(m, src)
}
func (m *GetBlockchainResponse) XXX_Size() int {
	return xxx_messageInfo_GetBlockchainResponse.Size(m)
}
func (m *GetBlockchainResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockchainResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockchainResponse proto.InternalMessageInfo

func (m *GetBlockchainResponse) GetBlocks() []*Block {
	if m != nil {
		return m.Blocks
	}
	return nil
}

func init() {
	proto.RegisterType((*AddBlockRequest)(nil), "proto.AddBlockRequest")
	proto.RegisterType((*AddBlockResponse)(nil), "proto.AddBlockResponse")
	proto.RegisterType((*BlockData)(nil), "proto.BlockData")
	proto.RegisterMapType((map[string]string)(nil), "proto.BlockData.FileHashesEntry")
	proto.RegisterType((*Block)(nil), "proto.Block")
	proto.RegisterType((*GetBlockchainRequest)(nil), "proto.GetBlockchainRequest")
	proto.RegisterType((*GetBlockchainResponse)(nil), "proto.GetBlockchainResponse")
}

func init() { proto.RegisterFile("proto/blockchain.proto", fileDescriptor_a884755cc01e392a) }

var fileDescriptor_a884755cc01e392a = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xdb, 0x6e, 0xda, 0x40,
	0x10, 0xad, 0x0d, 0xb6, 0x60, 0x28, 0x02, 0x8d, 0xa8, 0x6b, 0xb9, 0xa8, 0x42, 0x96, 0x1f, 0x78,
	0xa2, 0x12, 0x7d, 0x68, 0x55, 0x09, 0xa9, 0x50, 0x9a, 0x8b, 0x92, 0x27, 0xff, 0xc1, 0x62, 0x6f,
	0x62, 0x07, 0x7c, 0x89, 0xbd, 0x20, 0xf1, 0x29, 0x91, 0xf2, 0xb1, 0x91, 0xc7, 0x8b, 0x01, 0x07,
	0xe5, 0x69, 0x67, 0xce, 0x99, 0x33, 0x3b, 0x37, 0x30, 0xd2, 0x2c, 0x11, 0xc9, 0x8f, 0xd5, 0x26,
	0xf1, 0xd6, 0x5e, 0xc0, 0xc2, 0x78, 0x42, 0x00, 0x6a, 0xf4, 0xd8, 0xbf, 0xa0, 0x37, 0xf7, 0xfd,
	0x45, 0xc1, 0xba, 0xfc, 0x79, 0xcb, 0x73, 0x81, 0x0e, 0x34, 0x7d, 0x26, 0x98, 0xa9, 0x8c, 0x94,
	0x71, 0x67, 0xda, 0x2f, 0xe3, 0x27, 0x14, 0xb2, 0x64, 0x82, 0xb9, 0xc4, 0xda, 0x0e, 0xf4, 0x8f,
	0xc2, 0x3c, 0x4d, 0xe2, 0x9c, 0x63, 0x1f, 0x1a, 0xf3, 0x7f, 0x77, 0x24, 0x6c, 0xb9, 0x85, 0x69,
	0xbf, 0xaa, 0xd0, 0xae, 0x94, 0x38, 0x84, 0xb6, 0x08, 0x23, 0x9e, 0x0b, 0x16, 0xa5, 0x14, 0xd5,
	0x76, 0x8f, 0x00, 0x1a, 0xa0, 0xb3, 0xad, 0x08, 0x92, 0xcc, 0x54, 0x89, 0x92, 0x1e, 0x5a, 0xd0,
	0x0a, 0x7d, 0x1e, 0x8b, 0x50, 0xec, 0xcd, 0x06, 0x31, 0x95, 0x5f, 0x64, 0x4c, 0xb3, 0xe4, 0x89,
	0x7b, 0xe2, 0x76, 0x69, 0x36, 0xcb, 0x8c, 0x15, 0x80, 0xdf, 0x01, 0xbc, 0x24, 0x8a, 0x42, 0x71,
	0xc3, 0xf2, 0xc0, 0xd4, 0x88, 0x3e, 0x41, 0xf0, 0x2f, 0xc0, 0x43, 0xb8, 0xe1, 0x85, 0xcd, 0x73,
	0x53, 0x1f, 0x35, 0xc6, 0x9d, 0xe9, 0xa8, 0xde, 0xef, 0xe4, 0xaa, 0x0a, 0xf9, 0x1f, 0x8b, 0x6c,
	0xef, 0x9e, 0x68, 0xac, 0x19, 0xf4, 0x6a, 0x74, 0x31, 0x84, 0x35, 0xdf, 0xcb, 0xf6, 0x0a, 0x13,
	0x07, 0xa0, 0xed, 0xd8, 0x66, 0xcb, 0x65, 0x5f, 0xa5, 0xf3, 0x47, 0xfd, 0xad, 0xd8, 0x8f, 0xa0,
	0xd1, 0x3f, 0xe8, 0x40, 0x37, 0xcd, 0xf8, 0x8e, 0x1c, 0x2a, 0xb6, 0x94, 0x9f, 0x83, 0xd5, 0x66,
	0xd4, 0x8f, 0x36, 0x83, 0x08, 0xcd, 0xa0, 0x48, 0x51, 0xce, 0x8a, 0x6c, 0xdb, 0x80, 0xc1, 0x35,
	0x17, 0x8b, 0xea, 0x08, 0xe4, 0xae, 0xed, 0x19, 0x7c, 0xa9, 0xe1, 0x72, 0x95, 0x0e, 0xe8, 0x74,
	0x32, 0xb9, 0xa9, 0xd0, 0x58, 0x3e, 0x9f, 0x7e, 0xe6, 0x4a, 0x6e, 0xfa, 0xa2, 0x00, 0x1c, 0xc5,
	0x38, 0x83, 0xd6, 0xe1, 0x26, 0xd0, 0x90, 0x82, 0xda, 0x75, 0x59, 0x5f, 0xdf, 0xe1, 0xe5, 0x8f,
	0xf6, 0x27, 0xbc, 0x87, 0xee, 0x59, 0x31, 0xf8, 0x4d, 0xc6, 0x5e, 0x2a, 0xdd, 0x1a, 0x5e, 0x26,
	0x0f, 0xd9, 0x56, 0x3a, 0xd1, 0x3f, 0xdf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x4f, 0x7a, 0xe3, 0xad,
	0x01, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BlockchainClient is the client API for Blockchain service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlockchainClient interface {
	AddBlock(ctx context.Context, in *AddBlockRequest, opts ...grpc.CallOption) (*AddBlockResponse, error)
	GetBlockchain(ctx context.Context, in *GetBlockchainRequest, opts ...grpc.CallOption) (*GetBlockchainResponse, error)
}

type blockchainClient struct {
	cc *grpc.ClientConn
}

func NewBlockchainClient(cc *grpc.ClientConn) BlockchainClient {
	return &blockchainClient{cc}
}

func (c *blockchainClient) AddBlock(ctx context.Context, in *AddBlockRequest, opts ...grpc.CallOption) (*AddBlockResponse, error) {
	out := new(AddBlockResponse)
	err := c.cc.Invoke(ctx, "/proto.Blockchain/AddBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainClient) GetBlockchain(ctx context.Context, in *GetBlockchainRequest, opts ...grpc.CallOption) (*GetBlockchainResponse, error) {
	out := new(GetBlockchainResponse)
	err := c.cc.Invoke(ctx, "/proto.Blockchain/GetBlockchain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlockchainServer is the server API for Blockchain service.
type BlockchainServer interface {
	AddBlock(context.Context, *AddBlockRequest) (*AddBlockResponse, error)
	GetBlockchain(context.Context, *GetBlockchainRequest) (*GetBlockchainResponse, error)
}

// UnimplementedBlockchainServer can be embedded to have forward compatible implementations.
type UnimplementedBlockchainServer struct {
}

func (*UnimplementedBlockchainServer) AddBlock(ctx context.Context, req *AddBlockRequest) (*AddBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBlock not implemented")
}
func (*UnimplementedBlockchainServer) GetBlockchain(ctx context.Context, req *GetBlockchainRequest) (*GetBlockchainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockchain not implemented")
}

func RegisterBlockchainServer(s *grpc.Server, srv BlockchainServer) {
	s.RegisterService(&_Blockchain_serviceDesc, srv)
}

func _Blockchain_AddBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServer).AddBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Blockchain/AddBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServer).AddBlock(ctx, req.(*AddBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blockchain_GetBlockchain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockchainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServer).GetBlockchain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Blockchain/GetBlockchain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServer).GetBlockchain(ctx, req.(*GetBlockchainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Blockchain_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Blockchain",
	HandlerType: (*BlockchainServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddBlock",
			Handler:    _Blockchain_AddBlock_Handler,
		},
		{
			MethodName: "GetBlockchain",
			Handler:    _Blockchain_GetBlockchain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/blockchain.proto",
}
