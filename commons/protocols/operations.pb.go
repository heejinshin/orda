// Code generated by protoc-gen-go. DO NOT EDIT.
// source: operations.proto

package protocols

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type PbOperationId struct {
	Era                  uint32   `protobuf:"varint,1,opt,name=era,proto3" json:"era,omitempty"`
	Lamport              uint64   `protobuf:"varint,2,opt,name=lamport,proto3" json:"lamport,omitempty"`
	Cuid                 *PbUuid  `protobuf:"bytes,3,opt,name=cuid,proto3" json:"cuid,omitempty"`
	Seq                  uint64   `protobuf:"varint,4,opt,name=seq,proto3" json:"seq,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PbOperationId) Reset()         { *m = PbOperationId{} }
func (m *PbOperationId) String() string { return proto.CompactTextString(m) }
func (*PbOperationId) ProtoMessage()    {}
func (*PbOperationId) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b4a5877375e491e, []int{0}
}

func (m *PbOperationId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PbOperationId.Unmarshal(m, b)
}
func (m *PbOperationId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PbOperationId.Marshal(b, m, deterministic)
}
func (m *PbOperationId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PbOperationId.Merge(m, src)
}
func (m *PbOperationId) XXX_Size() int {
	return xxx_messageInfo_PbOperationId.Size(m)
}
func (m *PbOperationId) XXX_DiscardUnknown() {
	xxx_messageInfo_PbOperationId.DiscardUnknown(m)
}

var xxx_messageInfo_PbOperationId proto.InternalMessageInfo

func (m *PbOperationId) GetEra() uint32 {
	if m != nil {
		return m.Era
	}
	return 0
}

func (m *PbOperationId) GetLamport() uint64 {
	if m != nil {
		return m.Lamport
	}
	return 0
}

func (m *PbOperationId) GetCuid() *PbUuid {
	if m != nil {
		return m.Cuid
	}
	return nil
}

func (m *PbOperationId) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

type PbOperation struct {
	Id                   *PbOperationId `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OpType               int32          `protobuf:"varint,2,opt,name=opType,proto3" json:"opType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PbOperation) Reset()         { *m = PbOperation{} }
func (m *PbOperation) String() string { return proto.CompactTextString(m) }
func (*PbOperation) ProtoMessage()    {}
func (*PbOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b4a5877375e491e, []int{1}
}

func (m *PbOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PbOperation.Unmarshal(m, b)
}
func (m *PbOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PbOperation.Marshal(b, m, deterministic)
}
func (m *PbOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PbOperation.Merge(m, src)
}
func (m *PbOperation) XXX_Size() int {
	return xxx_messageInfo_PbOperation.Size(m)
}
func (m *PbOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_PbOperation.DiscardUnknown(m)
}

var xxx_messageInfo_PbOperation proto.InternalMessageInfo

func (m *PbOperation) GetId() *PbOperationId {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *PbOperation) GetOpType() int32 {
	if m != nil {
		return m.OpType
	}
	return 0
}

type PbIncreaseOperation struct {
	Base                 *PbOperation `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Delta                int32        `protobuf:"varint,2,opt,name=delta,proto3" json:"delta,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PbIncreaseOperation) Reset()         { *m = PbIncreaseOperation{} }
func (m *PbIncreaseOperation) String() string { return proto.CompactTextString(m) }
func (*PbIncreaseOperation) ProtoMessage()    {}
func (*PbIncreaseOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b4a5877375e491e, []int{2}
}

func (m *PbIncreaseOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PbIncreaseOperation.Unmarshal(m, b)
}
func (m *PbIncreaseOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PbIncreaseOperation.Marshal(b, m, deterministic)
}
func (m *PbIncreaseOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PbIncreaseOperation.Merge(m, src)
}
func (m *PbIncreaseOperation) XXX_Size() int {
	return xxx_messageInfo_PbIncreaseOperation.Size(m)
}
func (m *PbIncreaseOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_PbIncreaseOperation.DiscardUnknown(m)
}

var xxx_messageInfo_PbIncreaseOperation proto.InternalMessageInfo

func (m *PbIncreaseOperation) GetBase() *PbOperation {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *PbIncreaseOperation) GetDelta() int32 {
	if m != nil {
		return m.Delta
	}
	return 0
}

func init() {
	proto.RegisterType((*PbOperationId)(nil), "protocols.PbOperationId")
	proto.RegisterType((*PbOperation)(nil), "protocols.PbOperation")
	proto.RegisterType((*PbIncreaseOperation)(nil), "protocols.PbIncreaseOperation")
}

func init() { proto.RegisterFile("operations.proto", fileDescriptor_1b4a5877375e491e) }

var fileDescriptor_1b4a5877375e491e = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8e, 0x4f, 0x4b, 0x03, 0x31,
	0x10, 0xc5, 0xc9, 0x36, 0xad, 0x38, 0x4b, 0xa1, 0x46, 0x29, 0xc1, 0xd3, 0xb2, 0x20, 0x04, 0x0f,
	0x7b, 0xa8, 0x9f, 0xa2, 0xa7, 0x2e, 0x41, 0xf1, 0x9c, 0x6c, 0xe6, 0x10, 0x58, 0x9b, 0x98, 0x3f,
	0x07, 0xbf, 0xbd, 0x6c, 0xba, 0xd5, 0x15, 0x3c, 0x65, 0x5e, 0x66, 0xde, 0xef, 0x3d, 0xd8, 0x39,
	0x8f, 0x41, 0x25, 0xeb, 0xce, 0xb1, 0xf3, 0xc1, 0x25, 0xc7, 0x6e, 0xcb, 0x33, 0xb8, 0x31, 0x3e,
	0x42, 0xce, 0xd6, 0x5c, 0xbe, 0xdb, 0x04, 0xdb, 0x5e, 0x9f, 0xae, 0xc7, 0x47, 0xc3, 0x76, 0xb0,
	0xc2, 0xa0, 0x38, 0x69, 0x88, 0xd8, 0xca, 0x69, 0x64, 0x1c, 0x6e, 0x46, 0xf5, 0xe1, 0x5d, 0x48,
	0xbc, 0x6a, 0x88, 0xa0, 0xf2, 0x2a, 0xd9, 0x13, 0xd0, 0x21, 0x5b, 0xc3, 0x57, 0x0d, 0x11, 0xf5,
	0xe1, 0xae, 0xfb, 0x89, 0xe8, 0x7a, 0xfd, 0x96, 0xad, 0x91, 0x65, 0x3d, 0x21, 0x23, 0x7e, 0x72,
	0x5a, 0xcc, 0xd3, 0xd8, 0x9e, 0xa0, 0x5e, 0xa4, 0x32, 0x01, 0x95, 0x35, 0x25, 0xb2, 0x3e, 0xf0,
	0x3f, 0x94, 0x45, 0x33, 0x59, 0x59, 0xc3, 0xf6, 0xb0, 0x71, 0xfe, 0xf5, 0xcb, 0x63, 0xa9, 0xb2,
	0x96, 0xb3, 0x6a, 0xdf, 0xe1, 0xbe, 0xd7, 0xc7, 0xf3, 0x10, 0x50, 0x45, 0xfc, 0x05, 0x3f, 0x03,
	0xd5, 0x2a, 0xe2, 0x8c, 0xde, 0xff, 0x8f, 0x96, 0xe5, 0x86, 0x3d, 0xc0, 0xda, 0xe0, 0x98, 0xd4,
	0x4c, 0xbe, 0x08, 0xbd, 0x29, 0x96, 0x97, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb3, 0xa1, 0xbb,
	0x42, 0x51, 0x01, 0x00, 0x00,
}
