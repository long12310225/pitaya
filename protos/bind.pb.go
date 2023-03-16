// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bind.proto

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BindMsg struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid" json:"uid,omitempty"`
	Fid                  string   `protobuf:"bytes,2,opt,name=fid" json:"fid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BindMsg) Reset()         { *m = BindMsg{} }
func (m *BindMsg) String() string { return proto.CompactTextString(m) }
func (*BindMsg) ProtoMessage()    {}
func (*BindMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_bind_9cf772af172ead7d, []int{0}
}
func (m *BindMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BindMsg.Unmarshal(m, b)
}
func (m *BindMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BindMsg.Marshal(b, m, deterministic)
}
func (dst *BindMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindMsg.Merge(dst, src)
}
func (m *BindMsg) XXX_Size() int {
	return xxx_messageInfo_BindMsg.Size(m)
}
func (m *BindMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_BindMsg.DiscardUnknown(m)
}

var xxx_messageInfo_BindMsg proto.InternalMessageInfo

func (m *BindMsg) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *BindMsg) GetFid() string {
	if m != nil {
		return m.Fid
	}
	return ""
}

func init() {
	proto.RegisterType((*BindMsg)(nil), "protos.BindMsg")
}

func init() { proto.RegisterFile("bind.proto", fileDescriptor_bind_9cf772af172ead7d) }

var fileDescriptor_bind_9cf772af172ead7d = []byte{
	// 83 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xca, 0xcc, 0x4b,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x53, 0xc5, 0x4a, 0xba, 0x5c, 0xec, 0x4e,
	0x99, 0x79, 0x29, 0xbe, 0xc5, 0xe9, 0x42, 0x02, 0x5c, 0xcc, 0xa5, 0x99, 0x29, 0x12, 0x8c, 0x0a,
	0x8c, 0x1a, 0x9c, 0x41, 0x20, 0x26, 0x48, 0x24, 0x2d, 0x33, 0x45, 0x82, 0x09, 0x22, 0x92, 0x96,
	0x99, 0x92, 0x04, 0xd1, 0x66, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xce, 0x99, 0xf6, 0x32, 0x4b,
	0x00, 0x00, 0x00,
}
