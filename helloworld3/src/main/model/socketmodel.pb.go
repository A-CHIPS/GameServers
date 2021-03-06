// Code generated by protoc-gen-go. DO NOT EDIT.
// source: socketmodel.proto

/*
Package model is a generated protocol buffer package.

It is generated from these files:
	socketmodel.proto

It has these top-level messages:
	Basemsg
	GChatContent
	GChatData
	BPlayerData
*/
package model

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

type Basemsg struct {
	Msgcode int32  `protobuf:"varint,1,opt,name=msgcode" json:"msgcode,omitempty"`
	Buff    []byte `protobuf:"bytes,2,opt,name=buff,proto3" json:"buff,omitempty"`
}

func (m *Basemsg) Reset()                    { *m = Basemsg{} }
func (m *Basemsg) String() string            { return proto.CompactTextString(m) }
func (*Basemsg) ProtoMessage()               {}
func (*Basemsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Basemsg) GetMsgcode() int32 {
	if m != nil {
		return m.Msgcode
	}
	return 0
}

func (m *Basemsg) GetBuff() []byte {
	if m != nil {
		return m.Buff
	}
	return nil
}

type GChatContent struct {
	Content map[string]string `protobuf:"bytes,1,rep,name=content" json:"content,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *GChatContent) Reset()                    { *m = GChatContent{} }
func (m *GChatContent) String() string            { return proto.CompactTextString(m) }
func (*GChatContent) ProtoMessage()               {}
func (*GChatContent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GChatContent) GetContent() map[string]string {
	if m != nil {
		return m.Content
	}
	return nil
}

type GChatData struct {
	ChatType   int32    `protobuf:"varint,1,opt,name=ChatType" json:"ChatType,omitempty"`
	SendDate   int64    `protobuf:"varint,2,opt,name=SendDate" json:"SendDate,omitempty"`
	Content    [][]byte `protobuf:"bytes,3,rep,name=Content,proto3" json:"Content,omitempty"`
	FromUserID int64    `protobuf:"varint,4,opt,name=FromUserID" json:"FromUserID,omitempty"`
	FromNick   string   `protobuf:"bytes,5,opt,name=FromNick" json:"FromNick,omitempty"`
	FromExp    int32    `protobuf:"varint,6,opt,name=FromExp" json:"FromExp,omitempty"`
	OperID     string   `protobuf:"bytes,7,opt,name=operID" json:"operID,omitempty"`
	OperParams []byte   `protobuf:"bytes,8,opt,name=operParams,proto3" json:"operParams,omitempty"`
	Icon       string   `protobuf:"bytes,9,opt,name=icon" json:"icon,omitempty"`
	Border     string   `protobuf:"bytes,10,opt,name=border" json:"border,omitempty"`
	OperType   int32    `protobuf:"varint,11,opt,name=operType" json:"operType,omitempty"`
}

func (m *GChatData) Reset()                    { *m = GChatData{} }
func (m *GChatData) String() string            { return proto.CompactTextString(m) }
func (*GChatData) ProtoMessage()               {}
func (*GChatData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GChatData) GetChatType() int32 {
	if m != nil {
		return m.ChatType
	}
	return 0
}

func (m *GChatData) GetSendDate() int64 {
	if m != nil {
		return m.SendDate
	}
	return 0
}

func (m *GChatData) GetContent() [][]byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *GChatData) GetFromUserID() int64 {
	if m != nil {
		return m.FromUserID
	}
	return 0
}

func (m *GChatData) GetFromNick() string {
	if m != nil {
		return m.FromNick
	}
	return ""
}

func (m *GChatData) GetFromExp() int32 {
	if m != nil {
		return m.FromExp
	}
	return 0
}

func (m *GChatData) GetOperID() string {
	if m != nil {
		return m.OperID
	}
	return ""
}

func (m *GChatData) GetOperParams() []byte {
	if m != nil {
		return m.OperParams
	}
	return nil
}

func (m *GChatData) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *GChatData) GetBorder() string {
	if m != nil {
		return m.Border
	}
	return ""
}

func (m *GChatData) GetOperType() int32 {
	if m != nil {
		return m.OperType
	}
	return 0
}

type BPlayerData struct {
	Pid         int64  `protobuf:"varint,1,opt,name=pid" json:"pid,omitempty"`
	Pname       string `protobuf:"bytes,2,opt,name=pname" json:"pname,omitempty"`
	Picon       string `protobuf:"bytes,3,opt,name=picon" json:"picon,omitempty"`
	Piconborder string `protobuf:"bytes,4,opt,name=piconborder" json:"piconborder,omitempty"`
	Pexp        int32  `protobuf:"varint,5,opt,name=pexp" json:"pexp,omitempty"`
	Porganid    int32  `protobuf:"varint,6,opt,name=porganid" json:"porganid,omitempty"`
}

func (m *BPlayerData) Reset()                    { *m = BPlayerData{} }
func (m *BPlayerData) String() string            { return proto.CompactTextString(m) }
func (*BPlayerData) ProtoMessage()               {}
func (*BPlayerData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *BPlayerData) GetPid() int64 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *BPlayerData) GetPname() string {
	if m != nil {
		return m.Pname
	}
	return ""
}

func (m *BPlayerData) GetPicon() string {
	if m != nil {
		return m.Picon
	}
	return ""
}

func (m *BPlayerData) GetPiconborder() string {
	if m != nil {
		return m.Piconborder
	}
	return ""
}

func (m *BPlayerData) GetPexp() int32 {
	if m != nil {
		return m.Pexp
	}
	return 0
}

func (m *BPlayerData) GetPorganid() int32 {
	if m != nil {
		return m.Porganid
	}
	return 0
}

func init() {
	proto.RegisterType((*Basemsg)(nil), "model.basemsg")
	proto.RegisterType((*GChatContent)(nil), "model.GChatContent")
	proto.RegisterType((*GChatData)(nil), "model.GChatData")
	proto.RegisterType((*BPlayerData)(nil), "model.BPlayerData")
}

func init() { proto.RegisterFile("socketmodel.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x92, 0xc1, 0x8e, 0x9b, 0x30,
	0x10, 0x86, 0x45, 0x1c, 0x92, 0x65, 0xe0, 0xb0, 0xb5, 0xaa, 0xca, 0xda, 0x43, 0x85, 0x38, 0x71,
	0xca, 0xa1, 0x3d, 0xb4, 0xca, 0xb1, 0x9b, 0x6d, 0xd5, 0x4b, 0xb5, 0x72, 0xdb, 0x07, 0x30, 0xe0,
	0x4d, 0x51, 0x02, 0xb6, 0x8c, 0xb7, 0x5a, 0x5e, 0xa0, 0x6f, 0xd1, 0xa7, 0xe8, 0x0b, 0x56, 0x33,
	0x06, 0xc4, 0x89, 0xff, 0x1b, 0x33, 0x33, 0xbf, 0x7f, 0x80, 0x57, 0x83, 0xa9, 0x2f, 0xda, 0x77,
	0xa6, 0xd1, 0xd7, 0x83, 0x75, 0xc6, 0x1b, 0x1e, 0x13, 0x14, 0x1f, 0x60, 0x5f, 0xa9, 0x41, 0x77,
	0xc3, 0x99, 0x0b, 0xd8, 0x77, 0xc3, 0xb9, 0x36, 0x8d, 0x16, 0x51, 0x1e, 0x95, 0xb1, 0x9c, 0x91,
	0x73, 0xd8, 0x56, 0xcf, 0x4f, 0x4f, 0x62, 0x93, 0x47, 0x65, 0x26, 0x49, 0x17, 0x7f, 0x22, 0xc8,
	0xbe, 0xdc, 0xff, 0x52, 0xfe, 0xde, 0xf4, 0x5e, 0xf7, 0x9e, 0x1f, 0x61, 0x5f, 0x07, 0x29, 0xa2,
	0x9c, 0x95, 0xe9, 0xbb, 0xfc, 0x10, 0xf6, 0xad, 0xdf, 0x3a, 0x4c, 0xcf, 0x87, 0xde, 0xbb, 0x51,
	0xce, 0x0d, 0x77, 0x47, 0xc8, 0xd6, 0x07, 0xfc, 0x16, 0xd8, 0x45, 0x8f, 0x64, 0x23, 0x91, 0x28,
	0xf9, 0x6b, 0x88, 0x7f, 0xab, 0xeb, 0xb3, 0x26, 0x0f, 0x89, 0x0c, 0x70, 0xdc, 0x7c, 0x8c, 0x8a,
	0x7f, 0x1b, 0x48, 0x68, 0xc5, 0x49, 0x79, 0xc5, 0xef, 0xe0, 0x06, 0xf5, 0x8f, 0xd1, 0xce, 0xb7,
	0x58, 0x18, 0xcf, 0xbe, 0xeb, 0xbe, 0x39, 0x29, 0x1f, 0xc6, 0x30, 0xb9, 0x30, 0x5e, 0x7e, 0x72,
	0x20, 0x58, 0xce, 0xca, 0x4c, 0xce, 0xc8, 0xdf, 0x02, 0x7c, 0x76, 0xa6, 0xfb, 0x39, 0x68, 0xf7,
	0xf5, 0x24, 0xb6, 0xd4, 0xb7, 0xaa, 0xe0, 0x54, 0xa4, 0x6f, 0x6d, 0x7d, 0x11, 0x31, 0x99, 0x5b,
	0x18, 0xa7, 0xa2, 0x7e, 0x78, 0xb1, 0x62, 0x17, 0x22, 0x9d, 0x90, 0xbf, 0x81, 0x9d, 0xb1, 0x34,
	0x71, 0x4f, 0x3d, 0x13, 0xe1, 0x36, 0x54, 0x8f, 0xca, 0xa9, 0x6e, 0x10, 0x37, 0x14, 0xf8, 0xaa,
	0x82, 0x9f, 0xa2, 0xad, 0x4d, 0x2f, 0x12, 0xea, 0x22, 0x8d, 0xb3, 0x2a, 0xe3, 0x1a, 0xed, 0x04,
	0x84, 0x59, 0x81, 0xd0, 0x19, 0x76, 0x52, 0x16, 0x69, 0xc8, 0x62, 0xe6, 0xe2, 0x6f, 0x04, 0xe9,
	0xa7, 0xc7, 0xab, 0x1a, 0xb5, 0xa3, 0xdc, 0x6e, 0x81, 0xd9, 0xb6, 0xa1, 0xc8, 0x98, 0x44, 0x89,
	0x89, 0xdb, 0x5e, 0x75, 0x4b, 0xe2, 0x04, 0x54, 0x25, 0x03, 0x6c, 0xaa, 0x92, 0x83, 0x1c, 0x52,
	0x12, 0x93, 0x8d, 0x2d, 0x9d, 0xad, 0x4b, 0xe8, 0xdb, 0xea, 0x17, 0x4b, 0x09, 0xc5, 0x92, 0x34,
	0xfa, 0xb3, 0xc6, 0x9d, 0x55, 0xdf, 0x36, 0x53, 0x3c, 0x0b, 0x57, 0x3b, 0xfa, 0x4b, 0xdf, 0xff,
	0x0f, 0x00, 0x00, 0xff, 0xff, 0xca, 0x6d, 0xd8, 0x8b, 0xba, 0x02, 0x00, 0x00,
}
