// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

/*
Package user is a generated protocol buffer package.

It is generated from these files:
	user.proto

It has these top-level messages:
	User
*/
package user

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

type User struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	Id    int64  `protobuf:"varint,3,opt,name=id" json:"id,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*User)(nil), "user.User")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 100 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x1c, 0xb8, 0x58, 0x42, 0x8b,
	0x53, 0x8b, 0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38,
	0x83, 0xc0, 0x6c, 0x21, 0x11, 0x2e, 0xd6, 0xd4, 0xdc, 0xc4, 0xcc, 0x1c, 0x09, 0x26, 0xb0, 0x20,
	0x84, 0x23, 0xc4, 0xc7, 0xc5, 0x94, 0x99, 0x22, 0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x1c, 0xc4, 0x94,
	0x99, 0x92, 0xc4, 0x06, 0x36, 0xce, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x85, 0x24, 0x35, 0xfa,
	0x5c, 0x00, 0x00, 0x00,
}