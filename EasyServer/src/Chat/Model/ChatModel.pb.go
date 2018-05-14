// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ChatModel.proto

/*
Package ChatModel is a generated protocol buffer package.

It is generated from these files:
	ChatModel.proto

It has these top-level messages:
	ChatMessage
	ChaterInfo
	AllChaterInfo
*/
package Model

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

type ChatMessage struct {
	Content string `protobuf:"bytes,1,opt,name=content" json:"content,omitempty"`
	Id      int32  `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Nick    string `protobuf:"bytes,3,opt,name=nick" json:"nick,omitempty"`
}

func (m *ChatMessage) Reset()                    { *m = ChatMessage{} }
func (m *ChatMessage) String() string            { return proto.CompactTextString(m) }
func (*ChatMessage) ProtoMessage()               {}
func (*ChatMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ChatMessage) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *ChatMessage) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ChatMessage) GetNick() string {
	if m != nil {
		return m.Nick
	}
	return ""
}

type ChaterInfo struct {
	Nick string `protobuf:"bytes,1,opt,name=nick" json:"nick,omitempty"`
	Id   int32  `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
}

func (m *ChaterInfo) Reset()                    { *m = ChaterInfo{} }
func (m *ChaterInfo) String() string            { return proto.CompactTextString(m) }
func (*ChaterInfo) ProtoMessage()               {}
func (*ChaterInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ChaterInfo) GetNick() string {
	if m != nil {
		return m.Nick
	}
	return ""
}

func (m *ChaterInfo) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type AllChaterInfo struct {
	Data []*ChaterInfo `protobuf:"bytes,1,rep,name=data" json:"data,omitempty"`
}

func (m *AllChaterInfo) Reset()                    { *m = AllChaterInfo{} }
func (m *AllChaterInfo) String() string            { return proto.CompactTextString(m) }
func (*AllChaterInfo) ProtoMessage()               {}
func (*AllChaterInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AllChaterInfo) GetData() []*ChaterInfo {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ChatMessage)(nil), "ChatModel.ChatMessage")
	proto.RegisterType((*ChaterInfo)(nil), "ChatModel.ChaterInfo")
	proto.RegisterType((*AllChaterInfo)(nil), "ChatModel.AllChaterInfo")
}

func init() { proto.RegisterFile("ChatModel.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x77, 0xce, 0x48, 0x2c,
	0xf1, 0xcd, 0x4f, 0x49, 0xcd, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x0b, 0x28,
	0x79, 0x73, 0x71, 0x83, 0x39, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x42, 0x12, 0x5c, 0xec, 0xc9,
	0xf9, 0x79, 0x25, 0xa9, 0x79, 0x25, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x30, 0xae, 0x10,
	0x1f, 0x17, 0x53, 0x66, 0x8a, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x53, 0x66, 0x8a, 0x90,
	0x10, 0x17, 0x4b, 0x5e, 0x66, 0x72, 0xb6, 0x04, 0x33, 0x58, 0x19, 0x98, 0xad, 0x64, 0xc0, 0xc5,
	0x05, 0x32, 0x2c, 0xb5, 0xc8, 0x33, 0x2f, 0x2d, 0x1f, 0xae, 0x82, 0x11, 0xa1, 0x02, 0xdd, 0x14,
	0x25, 0x2b, 0x2e, 0x5e, 0xc7, 0x9c, 0x1c, 0x24, 0x4d, 0x9a, 0x5c, 0x2c, 0x29, 0x89, 0x25, 0x89,
	0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0xa2, 0x7a, 0x08, 0xa7, 0x23, 0x14, 0x05, 0x81, 0x95,
	0x24, 0xb1, 0x81, 0x3d, 0x63, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x10, 0x3a, 0xd2, 0xf2, 0xdf,
	0x00, 0x00, 0x00,
}
