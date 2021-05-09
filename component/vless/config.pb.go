// Code generated by protoc-gen-go. DO NOT EDIT.
// source: component/vless/config.proto

package vless

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

type Addons struct {
	Flow                 string   `protobuf:"bytes,1,opt,name=Flow,proto3" json:"Flow,omitempty"`
	Seed                 []byte   `protobuf:"bytes,2,opt,name=Seed,proto3" json:"Seed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Addons) Reset()         { *m = Addons{} }
func (m *Addons) String() string { return proto.CompactTextString(m) }
func (*Addons) ProtoMessage()    {}
func (*Addons) Descriptor() ([]byte, []int) {
	return fileDescriptor_73f48668a00892e0, []int{0}
}

func (m *Addons) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Addons.Unmarshal(m, b)
}
func (m *Addons) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Addons.Marshal(b, m, deterministic)
}
func (m *Addons) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Addons.Merge(m, src)
}
func (m *Addons) XXX_Size() int {
	return xxx_messageInfo_Addons.Size(m)
}
func (m *Addons) XXX_DiscardUnknown() {
	xxx_messageInfo_Addons.DiscardUnknown(m)
}

var xxx_messageInfo_Addons proto.InternalMessageInfo

func (m *Addons) GetFlow() string {
	if m != nil {
		return m.Flow
	}
	return ""
}

func (m *Addons) GetSeed() []byte {
	if m != nil {
		return m.Seed
	}
	return nil
}

func init() {
	proto.RegisterType((*Addons)(nil), "component.vless.Addons")
}

func init() {
	proto.RegisterFile("component/vless/config.proto", fileDescriptor_73f48668a00892e0)
}

var fileDescriptor_73f48668a00892e0 = []byte{
	// 114 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x49, 0xce, 0xcf, 0x2d,
	0xc8, 0xcf, 0x4b, 0xcd, 0x2b, 0xd1, 0x2f, 0xcb, 0x49, 0x2d, 0x2e, 0xd6, 0x4f, 0xce, 0xcf, 0x4b,
	0xcb, 0x4c, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x87, 0xcb, 0xea, 0x81, 0x65, 0x95,
	0x0c, 0xb8, 0xd8, 0x1c, 0x53, 0x52, 0xf2, 0xf3, 0x8a, 0x85, 0x84, 0xb8, 0x58, 0xdc, 0x72, 0xf2,
	0xcb, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x90, 0x58, 0x70, 0x6a, 0x6a, 0x8a,
	0x04, 0x93, 0x02, 0xa3, 0x06, 0x4f, 0x10, 0x98, 0xed, 0x24, 0x18, 0xc5, 0x8f, 0x66, 0x45, 0x12,
	0x1b, 0xd8, 0x70, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x95, 0xe9, 0x83, 0x34, 0x7c, 0x00,
	0x00, 0x00,
}
