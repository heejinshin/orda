// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ortoo.grpc.proto

package model

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

type SnapshotRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Collection           string   `protobuf:"bytes,2,opt,name=collection,proto3" json:"collection,omitempty"`
	Duid                 string   `protobuf:"bytes,3,opt,name=duid,proto3" json:"duid,omitempty"`
	Sseq                 uint64   `protobuf:"varint,4,opt,name=sseq,proto3" json:"sseq,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SnapshotRequest) Reset()         { *m = SnapshotRequest{} }
func (m *SnapshotRequest) String() string { return proto.CompactTextString(m) }
func (*SnapshotRequest) ProtoMessage()    {}
func (*SnapshotRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_61fc1459943e44c7, []int{0}
}
func (m *SnapshotRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SnapshotRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SnapshotRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SnapshotRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnapshotRequest.Merge(m, src)
}
func (m *SnapshotRequest) XXX_Size() int {
	return m.Size()
}
func (m *SnapshotRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SnapshotRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SnapshotRequest proto.InternalMessageInfo

func (m *SnapshotRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SnapshotRequest) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

func (m *SnapshotRequest) GetDuid() string {
	if m != nil {
		return m.Duid
	}
	return ""
}

func (m *SnapshotRequest) GetSseq() uint64 {
	if m != nil {
		return m.Sseq
	}
	return 0
}

type SnapshotResponse struct {
	Meta                 *DatatypeMeta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Json                 string        `protobuf:"bytes,2,opt,name=json,proto3" json:"json,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SnapshotResponse) Reset()         { *m = SnapshotResponse{} }
func (m *SnapshotResponse) String() string { return proto.CompactTextString(m) }
func (*SnapshotResponse) ProtoMessage()    {}
func (*SnapshotResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_61fc1459943e44c7, []int{1}
}
func (m *SnapshotResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SnapshotResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SnapshotResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SnapshotResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnapshotResponse.Merge(m, src)
}
func (m *SnapshotResponse) XXX_Size() int {
	return m.Size()
}
func (m *SnapshotResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SnapshotResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SnapshotResponse proto.InternalMessageInfo

func (m *SnapshotResponse) GetMeta() *DatatypeMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *SnapshotResponse) GetJson() string {
	if m != nil {
		return m.Json
	}
	return ""
}

func init() {
	proto.RegisterType((*SnapshotRequest)(nil), "ortoo.SnapshotRequest")
	proto.RegisterType((*SnapshotResponse)(nil), "ortoo.SnapshotResponse")
}

func init() { proto.RegisterFile("ortoo.grpc.proto", fileDescriptor_61fc1459943e44c7) }

var fileDescriptor_61fc1459943e44c7 = []byte{
	// 480 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xc5, 0x89, 0x01, 0xb1, 0x05, 0x35, 0x2c, 0xa8, 0xb5, 0x2c, 0x30, 0xd6, 0x5e, 0x88, 0x90,
	0x1a, 0x37, 0x06, 0x21, 0xd4, 0x72, 0x69, 0xc3, 0x85, 0x43, 0x55, 0xcb, 0xbd, 0x71, 0x5b, 0xec,
	0x91, 0x63, 0xba, 0xf5, 0x6e, 0xbc, 0xeb, 0x54, 0x51, 0x65, 0x21, 0xf1, 0x09, 0x70, 0xe1, 0x93,
	0x38, 0x22, 0xf1, 0x03, 0x28, 0xf0, 0x15, 0x9c, 0x90, 0xd7, 0x76, 0x5c, 0x51, 0xaa, 0x9c, 0xbc,
	0xf3, 0xde, 0x8c, 0xdf, 0x7b, 0xb3, 0x36, 0x1a, 0xf0, 0x5c, 0x71, 0x3e, 0x4a, 0x72, 0x11, 0x8d,
	0x44, 0xce, 0x15, 0xc7, 0x37, 0x35, 0x62, 0x6f, 0xd4, 0x84, 0xc6, 0xec, 0x47, 0x09, 0xe7, 0x09,
	0x03, 0x8f, 0x8a, 0xd4, 0xa3, 0x59, 0xc6, 0x15, 0x55, 0x29, 0xcf, 0x64, 0xc3, 0xd6, 0x8f, 0x68,
	0x27, 0x81, 0x6c, 0x87, 0x0b, 0xc8, 0xa8, 0x48, 0xe7, 0xbe, 0xc7, 0x85, 0xee, 0xb9, 0xda, 0x4f,
	0x38, 0xda, 0x3c, 0xc9, 0xa8, 0x90, 0x53, 0xae, 0x42, 0x98, 0x15, 0x20, 0x15, 0x1e, 0xa0, 0xfe,
	0x29, 0x2c, 0x2c, 0xc3, 0x35, 0x86, 0x77, 0xc2, 0xea, 0x88, 0x1d, 0x84, 0x22, 0xce, 0x18, 0x44,
	0xd5, 0xa4, 0xd5, 0xd3, 0xc4, 0x25, 0x04, 0x63, 0x64, 0xc6, 0x45, 0x1a, 0x5b, 0x7d, 0xcd, 0xe8,
	0x33, 0xde, 0x42, 0xa6, 0x94, 0x30, 0xb3, 0x4c, 0xd7, 0x18, 0x9a, 0x87, 0xbd, 0x5d, 0x23, 0xd4,
	0x35, 0x39, 0x46, 0x83, 0x4e, 0x50, 0x0a, 0x9e, 0x49, 0xc0, 0x4f, 0x91, 0x79, 0x06, 0x8a, 0x6a,
	0xc9, 0x0d, 0xff, 0xc1, 0xa8, 0x8e, 0xfb, 0x86, 0x2a, 0xaa, 0x16, 0x02, 0x8e, 0x40, 0xd1, 0x50,
	0x37, 0x54, 0x42, 0x1f, 0xe4, 0xca, 0x82, 0x3e, 0xfb, 0x7f, 0xfa, 0xe8, 0xee, 0x71, 0x35, 0x70,
	0x02, 0xf9, 0x3c, 0x8d, 0x00, 0x7f, 0x44, 0x9b, 0x41, 0xce, 0x23, 0x90, 0x32, 0x28, 0xe4, 0x34,
	0x28, 0x18, 0xc3, 0x5b, 0xcd, 0x2b, 0x5b, 0xe0, 0x08, 0xa4, 0xa4, 0x09, 0xd8, 0xd7, 0xe0, 0xe4,
	0xf5, 0xa7, 0x1f, 0xbf, 0xbf, 0xf4, 0x5e, 0x92, 0xb1, 0xde, 0xf2, 0x7c, 0xec, 0x75, 0x61, 0xa5,
	0x77, 0xd1, 0x15, 0xa5, 0x27, 0x0a, 0x39, 0x15, 0x05, 0x63, 0x15, 0x5c, 0xa4, 0x71, 0xb9, 0x67,
	0x3c, 0xc3, 0xe7, 0xe8, 0x5e, 0x63, 0x60, 0xc2, 0x52, 0xc8, 0x14, 0x7e, 0xd8, 0xc8, 0xd4, 0x65,
	0x2b, 0xfe, 0x5f, 0x94, 0xec, 0x69, 0xe9, 0x17, 0xc4, 0x5b, 0x2b, 0x1d, 0xe9, 0xb9, 0xcb, 0xc2,
	0x65, 0xb7, 0xdb, 0x76, 0x79, 0xab, 0xe8, 0xff, 0xdc, 0xb2, 0xbd, 0x7d, 0x05, 0xaf, 0x2f, 0x83,
	0xbc, 0xd2, 0x06, 0x7c, 0xbc, 0xbb, 0xd6, 0x40, 0xdc, 0x68, 0x48, 0xef, 0xe2, 0x14, 0x16, 0x25,
	0x9e, 0xa1, 0xfb, 0x93, 0x1c, 0xa8, 0x82, 0x49, 0x37, 0x81, 0xad, 0x36, 0xe5, 0x0a, 0x6b, 0xf3,
	0x5f, 0xcb, 0x90, 0xa1, 0xb6, 0x40, 0x6c, 0x77, 0x9d, 0x85, 0xc3, 0xfd, 0xcf, 0x07, 0x4f, 0xf0,
	0x63, 0xb4, 0xad, 0x3f, 0x00, 0x37, 0x09, 0x83, 0x89, 0x9b, 0x50, 0x05, 0xe7, 0x74, 0xe1, 0x1e,
	0x04, 0x6f, 0xa5, 0xdf, 0x9b, 0x8f, 0xbf, 0x2d, 0x1d, 0xe3, 0xfb, 0xd2, 0x31, 0x7e, 0x2e, 0x1d,
	0xe3, 0xeb, 0x2f, 0xe7, 0xc6, 0xbb, 0xdb, 0xa3, 0xfd, 0x33, 0x1e, 0x03, 0x7b, 0x7f, 0x4b, 0xff,
	0x02, 0xcf, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x91, 0x2c, 0x79, 0x01, 0x78, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OrtooServiceClient is the client API for OrtooService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrtooServiceClient interface {
	ProcessPushPull(ctx context.Context, in *PushPullMessage, opts ...grpc.CallOption) (*PushPullMessage, error)
	ProcessClient(ctx context.Context, in *ClientMessage, opts ...grpc.CallOption) (*ClientMessage, error)
	SnapshotDatatype(ctx context.Context, in *SnapshotRequest, opts ...grpc.CallOption) (*SnapshotResponse, error)
	CreateCollections(ctx context.Context, in *CollectionMessage, opts ...grpc.CallOption) (*CollectionMessage, error)
}

type ortooServiceClient struct {
	cc *grpc.ClientConn
}

func NewOrtooServiceClient(cc *grpc.ClientConn) OrtooServiceClient {
	return &ortooServiceClient{cc}
}

func (c *ortooServiceClient) ProcessPushPull(ctx context.Context, in *PushPullMessage, opts ...grpc.CallOption) (*PushPullMessage, error) {
	out := new(PushPullMessage)
	err := c.cc.Invoke(ctx, "/ortoo.OrtooService/ProcessPushPull", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ortooServiceClient) ProcessClient(ctx context.Context, in *ClientMessage, opts ...grpc.CallOption) (*ClientMessage, error) {
	out := new(ClientMessage)
	err := c.cc.Invoke(ctx, "/ortoo.OrtooService/ProcessClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ortooServiceClient) SnapshotDatatype(ctx context.Context, in *SnapshotRequest, opts ...grpc.CallOption) (*SnapshotResponse, error) {
	out := new(SnapshotResponse)
	err := c.cc.Invoke(ctx, "/ortoo.OrtooService/SnapshotDatatype", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ortooServiceClient) CreateCollections(ctx context.Context, in *CollectionMessage, opts ...grpc.CallOption) (*CollectionMessage, error) {
	out := new(CollectionMessage)
	err := c.cc.Invoke(ctx, "/ortoo.OrtooService/CreateCollections", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrtooServiceServer is the server API for OrtooService service.
type OrtooServiceServer interface {
	ProcessPushPull(context.Context, *PushPullMessage) (*PushPullMessage, error)
	ProcessClient(context.Context, *ClientMessage) (*ClientMessage, error)
	SnapshotDatatype(context.Context, *SnapshotRequest) (*SnapshotResponse, error)
	CreateCollections(context.Context, *CollectionMessage) (*CollectionMessage, error)
}

// UnimplementedOrtooServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOrtooServiceServer struct {
}

func (*UnimplementedOrtooServiceServer) ProcessPushPull(ctx context.Context, req *PushPullMessage) (*PushPullMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessPushPull not implemented")
}
func (*UnimplementedOrtooServiceServer) ProcessClient(ctx context.Context, req *ClientMessage) (*ClientMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessClient not implemented")
}
func (*UnimplementedOrtooServiceServer) SnapshotDatatype(ctx context.Context, req *SnapshotRequest) (*SnapshotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SnapshotDatatype not implemented")
}
func (*UnimplementedOrtooServiceServer) CreateCollections(ctx context.Context, req *CollectionMessage) (*CollectionMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCollections not implemented")
}

func RegisterOrtooServiceServer(s *grpc.Server, srv OrtooServiceServer) {
	s.RegisterService(&_OrtooService_serviceDesc, srv)
}

func _OrtooService_ProcessPushPull_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushPullMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrtooServiceServer).ProcessPushPull(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ortoo.OrtooService/ProcessPushPull",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrtooServiceServer).ProcessPushPull(ctx, req.(*PushPullMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrtooService_ProcessClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrtooServiceServer).ProcessClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ortoo.OrtooService/ProcessClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrtooServiceServer).ProcessClient(ctx, req.(*ClientMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrtooService_SnapshotDatatype_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrtooServiceServer).SnapshotDatatype(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ortoo.OrtooService/SnapshotDatatype",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrtooServiceServer).SnapshotDatatype(ctx, req.(*SnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrtooService_CreateCollections_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectionMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrtooServiceServer).CreateCollections(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ortoo.OrtooService/CreateCollections",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrtooServiceServer).CreateCollections(ctx, req.(*CollectionMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _OrtooService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ortoo.OrtooService",
	HandlerType: (*OrtooServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessPushPull",
			Handler:    _OrtooService_ProcessPushPull_Handler,
		},
		{
			MethodName: "ProcessClient",
			Handler:    _OrtooService_ProcessClient_Handler,
		},
		{
			MethodName: "SnapshotDatatype",
			Handler:    _OrtooService_SnapshotDatatype_Handler,
		},
		{
			MethodName: "CreateCollections",
			Handler:    _OrtooService_CreateCollections_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ortoo.grpc.proto",
}

func (m *SnapshotRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SnapshotRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SnapshotRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Sseq != 0 {
		i = encodeVarintOrtooGrpc(dAtA, i, uint64(m.Sseq))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Duid) > 0 {
		i -= len(m.Duid)
		copy(dAtA[i:], m.Duid)
		i = encodeVarintOrtooGrpc(dAtA, i, uint64(len(m.Duid)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Collection) > 0 {
		i -= len(m.Collection)
		copy(dAtA[i:], m.Collection)
		i = encodeVarintOrtooGrpc(dAtA, i, uint64(len(m.Collection)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintOrtooGrpc(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SnapshotResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SnapshotResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SnapshotResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Json) > 0 {
		i -= len(m.Json)
		copy(dAtA[i:], m.Json)
		i = encodeVarintOrtooGrpc(dAtA, i, uint64(len(m.Json)))
		i--
		dAtA[i] = 0x12
	}
	if m.Meta != nil {
		{
			size, err := m.Meta.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOrtooGrpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOrtooGrpc(dAtA []byte, offset int, v uint64) int {
	offset -= sovOrtooGrpc(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SnapshotRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovOrtooGrpc(uint64(l))
	}
	l = len(m.Collection)
	if l > 0 {
		n += 1 + l + sovOrtooGrpc(uint64(l))
	}
	l = len(m.Duid)
	if l > 0 {
		n += 1 + l + sovOrtooGrpc(uint64(l))
	}
	if m.Sseq != 0 {
		n += 1 + sovOrtooGrpc(uint64(m.Sseq))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SnapshotResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Meta != nil {
		l = m.Meta.Size()
		n += 1 + l + sovOrtooGrpc(uint64(l))
	}
	l = len(m.Json)
	if l > 0 {
		n += 1 + l + sovOrtooGrpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovOrtooGrpc(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOrtooGrpc(x uint64) (n int) {
	return sovOrtooGrpc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SnapshotRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrtooGrpc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SnapshotRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SnapshotRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrtooGrpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrtooGrpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrtooGrpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Collection", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrtooGrpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrtooGrpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrtooGrpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Collection = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrtooGrpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrtooGrpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrtooGrpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Duid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sseq", wireType)
			}
			m.Sseq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrtooGrpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sseq |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOrtooGrpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOrtooGrpc
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOrtooGrpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SnapshotResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrtooGrpc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SnapshotResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SnapshotResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Meta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrtooGrpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOrtooGrpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrtooGrpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Meta == nil {
				m.Meta = &DatatypeMeta{}
			}
			if err := m.Meta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Json", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrtooGrpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrtooGrpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrtooGrpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Json = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOrtooGrpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOrtooGrpc
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOrtooGrpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipOrtooGrpc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOrtooGrpc
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowOrtooGrpc
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowOrtooGrpc
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthOrtooGrpc
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOrtooGrpc
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOrtooGrpc
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOrtooGrpc        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOrtooGrpc          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOrtooGrpc = fmt.Errorf("proto: unexpected end of group")
)
