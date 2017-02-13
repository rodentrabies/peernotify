// Code generated by protoc-gen-go.
// source: contact.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	contact.proto
	message.proto

It has these top-level messages:
	Contact
	Method
	Message
*/
package pb

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

// Contact contains public key 'pubkey' which is used to build
// a cryptographic identity inside the system, and notification
// methods.
//
// TODO: currently, only email method is supported, but 'repeated'
// declaration should be considered, along with changing Method
// message declaration to be self-descriptive.
type Contact struct {
	Pubkey string  `protobuf:"bytes,1,opt,name=pubkey" json:"pubkey,omitempty"`
	Email  *Method `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
}

func (m *Contact) Reset()                    { *m = Contact{} }
func (m *Contact) String() string            { return proto.CompactTextString(m) }
func (*Contact) ProtoMessage()               {}
func (*Contact) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Contact) GetPubkey() string {
	if m != nil {
		return m.Pubkey
	}
	return ""
}

func (m *Contact) GetEmail() *Method {
	if m != nil {
		return m.Email
	}
	return nil
}

// Method represents notification method or medium and contains
// abstract address, which is the unique identifier used by the
// medium to identify a particular entity.
//
// TODO: self-desciption should be added to allow Contacts to have
// multiple notification mechanisms enabled; possibly 'enabled' field
// can be used to register multiple methods but have only a subset of
// them activated; also for key-forwarding model, Method should
// have (<address>, <encryption key>) pairs.
type Method struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
}

func (m *Method) Reset()                    { *m = Method{} }
func (m *Method) String() string            { return proto.CompactTextString(m) }
func (*Method) ProtoMessage()               {}
func (*Method) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Method) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*Contact)(nil), "pb.Contact")
	proto.RegisterType((*Method)(nil), "pb.Method")
}

func init() { proto.RegisterFile("contact.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 124 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xce, 0xcf, 0x2b,
	0x49, 0x4c, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x72, 0xe6,
	0x62, 0x77, 0x86, 0x08, 0x0a, 0x89, 0x71, 0xb1, 0x15, 0x94, 0x26, 0x65, 0xa7, 0x56, 0x4a, 0x30,
	0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x79, 0x42, 0x0a, 0x5c, 0xac, 0xa9, 0xb9, 0x89, 0x99, 0x39,
	0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x5c, 0x7a, 0x05, 0x49, 0x7a, 0xbe, 0xa9, 0x25, 0x19,
	0xf9, 0x29, 0x41, 0x10, 0x09, 0x25, 0x25, 0x2e, 0x36, 0x88, 0x80, 0x90, 0x04, 0x17, 0x7b, 0x62,
	0x4a, 0x4a, 0x51, 0x6a, 0x71, 0x31, 0xd4, 0x10, 0x18, 0x37, 0x89, 0x0d, 0x6c, 0xa7, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0xf5, 0x58, 0x8e, 0x26, 0x84, 0x00, 0x00, 0x00,
}
