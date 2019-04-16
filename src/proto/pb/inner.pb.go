// Code generated by protoc-gen-go. DO NOT EDIT.
// source: inner.proto

package pb

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

type InnerGamerLogout struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	LoginOhter           bool     `protobuf:"varint,2,opt,name=loginOhter,proto3" json:"loginOhter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InnerGamerLogout) Reset()         { *m = InnerGamerLogout{} }
func (m *InnerGamerLogout) String() string { return proto.CompactTextString(m) }
func (*InnerGamerLogout) ProtoMessage()    {}
func (*InnerGamerLogout) Descriptor() ([]byte, []int) {
	return fileDescriptor_9778376c45113e49, []int{0}
}

func (m *InnerGamerLogout) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InnerGamerLogout.Unmarshal(m, b)
}
func (m *InnerGamerLogout) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InnerGamerLogout.Marshal(b, m, deterministic)
}
func (m *InnerGamerLogout) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InnerGamerLogout.Merge(m, src)
}
func (m *InnerGamerLogout) XXX_Size() int {
	return xxx_messageInfo_InnerGamerLogout.Size(m)
}
func (m *InnerGamerLogout) XXX_DiscardUnknown() {
	xxx_messageInfo_InnerGamerLogout.DiscardUnknown(m)
}

var xxx_messageInfo_InnerGamerLogout proto.InternalMessageInfo

func (m *InnerGamerLogout) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *InnerGamerLogout) GetLoginOhter() bool {
	if m != nil {
		return m.LoginOhter
	}
	return false
}

type InnerNewGamerFriendReq struct {
	Gid                  int32    `protobuf:"varint,1,opt,name=gid,proto3" json:"gid,omitempty"`
	Oid                  int32    `protobuf:"varint,2,opt,name=oid,proto3" json:"oid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InnerNewGamerFriendReq) Reset()         { *m = InnerNewGamerFriendReq{} }
func (m *InnerNewGamerFriendReq) String() string { return proto.CompactTextString(m) }
func (*InnerNewGamerFriendReq) ProtoMessage()    {}
func (*InnerNewGamerFriendReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9778376c45113e49, []int{1}
}

func (m *InnerNewGamerFriendReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InnerNewGamerFriendReq.Unmarshal(m, b)
}
func (m *InnerNewGamerFriendReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InnerNewGamerFriendReq.Marshal(b, m, deterministic)
}
func (m *InnerNewGamerFriendReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InnerNewGamerFriendReq.Merge(m, src)
}
func (m *InnerNewGamerFriendReq) XXX_Size() int {
	return xxx_messageInfo_InnerNewGamerFriendReq.Size(m)
}
func (m *InnerNewGamerFriendReq) XXX_DiscardUnknown() {
	xxx_messageInfo_InnerNewGamerFriendReq.DiscardUnknown(m)
}

var xxx_messageInfo_InnerNewGamerFriendReq proto.InternalMessageInfo

func (m *InnerNewGamerFriendReq) GetGid() int32 {
	if m != nil {
		return m.Gid
	}
	return 0
}

func (m *InnerNewGamerFriendReq) GetOid() int32 {
	if m != nil {
		return m.Oid
	}
	return 0
}

type InnerNewGamerFriend struct {
	Gid                  int32    `protobuf:"varint,1,opt,name=gid,proto3" json:"gid,omitempty"`
	Oid                  int32    `protobuf:"varint,2,opt,name=oid,proto3" json:"oid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InnerNewGamerFriend) Reset()         { *m = InnerNewGamerFriend{} }
func (m *InnerNewGamerFriend) String() string { return proto.CompactTextString(m) }
func (*InnerNewGamerFriend) ProtoMessage()    {}
func (*InnerNewGamerFriend) Descriptor() ([]byte, []int) {
	return fileDescriptor_9778376c45113e49, []int{2}
}

func (m *InnerNewGamerFriend) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InnerNewGamerFriend.Unmarshal(m, b)
}
func (m *InnerNewGamerFriend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InnerNewGamerFriend.Marshal(b, m, deterministic)
}
func (m *InnerNewGamerFriend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InnerNewGamerFriend.Merge(m, src)
}
func (m *InnerNewGamerFriend) XXX_Size() int {
	return xxx_messageInfo_InnerNewGamerFriend.Size(m)
}
func (m *InnerNewGamerFriend) XXX_DiscardUnknown() {
	xxx_messageInfo_InnerNewGamerFriend.DiscardUnknown(m)
}

var xxx_messageInfo_InnerNewGamerFriend proto.InternalMessageInfo

func (m *InnerNewGamerFriend) GetGid() int32 {
	if m != nil {
		return m.Gid
	}
	return 0
}

func (m *InnerNewGamerFriend) GetOid() int32 {
	if m != nil {
		return m.Oid
	}
	return 0
}

type InnerNewGamerChat struct {
	Gid                  int32     `protobuf:"varint,1,opt,name=gid,proto3" json:"gid,omitempty"`
	Chat                 *ChatData `protobuf:"bytes,2,opt,name=chat,proto3" json:"chat,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *InnerNewGamerChat) Reset()         { *m = InnerNewGamerChat{} }
func (m *InnerNewGamerChat) String() string { return proto.CompactTextString(m) }
func (*InnerNewGamerChat) ProtoMessage()    {}
func (*InnerNewGamerChat) Descriptor() ([]byte, []int) {
	return fileDescriptor_9778376c45113e49, []int{3}
}

func (m *InnerNewGamerChat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InnerNewGamerChat.Unmarshal(m, b)
}
func (m *InnerNewGamerChat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InnerNewGamerChat.Marshal(b, m, deterministic)
}
func (m *InnerNewGamerChat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InnerNewGamerChat.Merge(m, src)
}
func (m *InnerNewGamerChat) XXX_Size() int {
	return xxx_messageInfo_InnerNewGamerChat.Size(m)
}
func (m *InnerNewGamerChat) XXX_DiscardUnknown() {
	xxx_messageInfo_InnerNewGamerChat.DiscardUnknown(m)
}

var xxx_messageInfo_InnerNewGamerChat proto.InternalMessageInfo

func (m *InnerNewGamerChat) GetGid() int32 {
	if m != nil {
		return m.Gid
	}
	return 0
}

func (m *InnerNewGamerChat) GetChat() *ChatData {
	if m != nil {
		return m.Chat
	}
	return nil
}

type InnerNewPVPResult struct {
	Gid                  int32    `protobuf:"varint,1,opt,name=gid,proto3" json:"gid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InnerNewPVPResult) Reset()         { *m = InnerNewPVPResult{} }
func (m *InnerNewPVPResult) String() string { return proto.CompactTextString(m) }
func (*InnerNewPVPResult) ProtoMessage()    {}
func (*InnerNewPVPResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_9778376c45113e49, []int{4}
}

func (m *InnerNewPVPResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InnerNewPVPResult.Unmarshal(m, b)
}
func (m *InnerNewPVPResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InnerNewPVPResult.Marshal(b, m, deterministic)
}
func (m *InnerNewPVPResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InnerNewPVPResult.Merge(m, src)
}
func (m *InnerNewPVPResult) XXX_Size() int {
	return xxx_messageInfo_InnerNewPVPResult.Size(m)
}
func (m *InnerNewPVPResult) XXX_DiscardUnknown() {
	xxx_messageInfo_InnerNewPVPResult.DiscardUnknown(m)
}

var xxx_messageInfo_InnerNewPVPResult proto.InternalMessageInfo

func (m *InnerNewPVPResult) GetGid() int32 {
	if m != nil {
		return m.Gid
	}
	return 0
}

type InnerMatchOk struct {
	Gid                  int32      `protobuf:"varint,1,opt,name=gid,proto3" json:"gid,omitempty"`
	MatchInfo            *MatchInfo `protobuf:"bytes,2,opt,name=matchInfo,proto3" json:"matchInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *InnerMatchOk) Reset()         { *m = InnerMatchOk{} }
func (m *InnerMatchOk) String() string { return proto.CompactTextString(m) }
func (*InnerMatchOk) ProtoMessage()    {}
func (*InnerMatchOk) Descriptor() ([]byte, []int) {
	return fileDescriptor_9778376c45113e49, []int{5}
}

func (m *InnerMatchOk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InnerMatchOk.Unmarshal(m, b)
}
func (m *InnerMatchOk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InnerMatchOk.Marshal(b, m, deterministic)
}
func (m *InnerMatchOk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InnerMatchOk.Merge(m, src)
}
func (m *InnerMatchOk) XXX_Size() int {
	return xxx_messageInfo_InnerMatchOk.Size(m)
}
func (m *InnerMatchOk) XXX_DiscardUnknown() {
	xxx_messageInfo_InnerMatchOk.DiscardUnknown(m)
}

var xxx_messageInfo_InnerMatchOk proto.InternalMessageInfo

func (m *InnerMatchOk) GetGid() int32 {
	if m != nil {
		return m.Gid
	}
	return 0
}

func (m *InnerMatchOk) GetMatchInfo() *MatchInfo {
	if m != nil {
		return m.MatchInfo
	}
	return nil
}

type InnerNewGamerMail struct {
	Gid                  int32    `protobuf:"varint,1,opt,name=gid,proto3" json:"gid,omitempty"`
	Mail                 *Mail    `protobuf:"bytes,2,opt,name=mail,proto3" json:"mail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InnerNewGamerMail) Reset()         { *m = InnerNewGamerMail{} }
func (m *InnerNewGamerMail) String() string { return proto.CompactTextString(m) }
func (*InnerNewGamerMail) ProtoMessage()    {}
func (*InnerNewGamerMail) Descriptor() ([]byte, []int) {
	return fileDescriptor_9778376c45113e49, []int{6}
}

func (m *InnerNewGamerMail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InnerNewGamerMail.Unmarshal(m, b)
}
func (m *InnerNewGamerMail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InnerNewGamerMail.Marshal(b, m, deterministic)
}
func (m *InnerNewGamerMail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InnerNewGamerMail.Merge(m, src)
}
func (m *InnerNewGamerMail) XXX_Size() int {
	return xxx_messageInfo_InnerNewGamerMail.Size(m)
}
func (m *InnerNewGamerMail) XXX_DiscardUnknown() {
	xxx_messageInfo_InnerNewGamerMail.DiscardUnknown(m)
}

var xxx_messageInfo_InnerNewGamerMail proto.InternalMessageInfo

func (m *InnerNewGamerMail) GetGid() int32 {
	if m != nil {
		return m.Gid
	}
	return 0
}

func (m *InnerNewGamerMail) GetMail() *Mail {
	if m != nil {
		return m.Mail
	}
	return nil
}

type Inner struct {
	GamerLogout          *InnerGamerLogout       `protobuf:"bytes,1,opt,name=gamerLogout,proto3" json:"gamerLogout,omitempty"`
	NewFriendReq         *InnerNewGamerFriendReq `protobuf:"bytes,2,opt,name=newFriendReq,proto3" json:"newFriendReq,omitempty"`
	NewFriend            *InnerNewGamerFriend    `protobuf:"bytes,3,opt,name=newFriend,proto3" json:"newFriend,omitempty"`
	NewChat              *InnerNewGamerChat      `protobuf:"bytes,4,opt,name=newChat,proto3" json:"newChat,omitempty"`
	NewPVPResult         *InnerNewPVPResult      `protobuf:"bytes,5,opt,name=newPVPResult,proto3" json:"newPVPResult,omitempty"`
	MatchOk              *InnerMatchOk           `protobuf:"bytes,6,opt,name=matchOk,proto3" json:"matchOk,omitempty"`
	NewGamerMail         *InnerNewGamerMail      `protobuf:"bytes,7,opt,name=newGamerMail,proto3" json:"newGamerMail,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Inner) Reset()         { *m = Inner{} }
func (m *Inner) String() string { return proto.CompactTextString(m) }
func (*Inner) ProtoMessage()    {}
func (*Inner) Descriptor() ([]byte, []int) {
	return fileDescriptor_9778376c45113e49, []int{7}
}

func (m *Inner) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Inner.Unmarshal(m, b)
}
func (m *Inner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Inner.Marshal(b, m, deterministic)
}
func (m *Inner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Inner.Merge(m, src)
}
func (m *Inner) XXX_Size() int {
	return xxx_messageInfo_Inner.Size(m)
}
func (m *Inner) XXX_DiscardUnknown() {
	xxx_messageInfo_Inner.DiscardUnknown(m)
}

var xxx_messageInfo_Inner proto.InternalMessageInfo

func (m *Inner) GetGamerLogout() *InnerGamerLogout {
	if m != nil {
		return m.GamerLogout
	}
	return nil
}

func (m *Inner) GetNewFriendReq() *InnerNewGamerFriendReq {
	if m != nil {
		return m.NewFriendReq
	}
	return nil
}

func (m *Inner) GetNewFriend() *InnerNewGamerFriend {
	if m != nil {
		return m.NewFriend
	}
	return nil
}

func (m *Inner) GetNewChat() *InnerNewGamerChat {
	if m != nil {
		return m.NewChat
	}
	return nil
}

func (m *Inner) GetNewPVPResult() *InnerNewPVPResult {
	if m != nil {
		return m.NewPVPResult
	}
	return nil
}

func (m *Inner) GetMatchOk() *InnerMatchOk {
	if m != nil {
		return m.MatchOk
	}
	return nil
}

func (m *Inner) GetNewGamerMail() *InnerNewGamerMail {
	if m != nil {
		return m.NewGamerMail
	}
	return nil
}

func init() {
	proto.RegisterType((*InnerGamerLogout)(nil), "InnerGamerLogout")
	proto.RegisterType((*InnerNewGamerFriendReq)(nil), "InnerNewGamerFriendReq")
	proto.RegisterType((*InnerNewGamerFriend)(nil), "InnerNewGamerFriend")
	proto.RegisterType((*InnerNewGamerChat)(nil), "InnerNewGamerChat")
	proto.RegisterType((*InnerNewPVPResult)(nil), "InnerNewPVPResult")
	proto.RegisterType((*InnerMatchOk)(nil), "InnerMatchOk")
	proto.RegisterType((*InnerNewGamerMail)(nil), "InnerNewGamerMail")
	proto.RegisterType((*Inner)(nil), "Inner")
}

func init() { proto.RegisterFile("inner.proto", fileDescriptor_9778376c45113e49) }

var fileDescriptor_9778376c45113e49 = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x41, 0x6b, 0xdb, 0x40,
	0x10, 0x85, 0xb1, 0x6c, 0xd9, 0xf5, 0xc8, 0x2d, 0xf6, 0xb6, 0xb4, 0x6a, 0xa1, 0xc5, 0x08, 0x4a,
	0x7d, 0x28, 0x7b, 0xb0, 0x21, 0x10, 0x92, 0x43, 0x48, 0x4c, 0x82, 0x43, 0x1c, 0x9b, 0x3d, 0xe4,
	0xbe, 0xb1, 0x36, 0xd2, 0x12, 0x69, 0x37, 0x51, 0x64, 0xfc, 0xc7, 0xf3, 0x03, 0xc2, 0x8e, 0x65,
	0xc9, 0x72, 0xf6, 0x90, 0x9b, 0x76, 0xe6, 0x7d, 0x6f, 0x86, 0x79, 0x08, 0x3c, 0xa9, 0x94, 0xc8,
	0xe8, 0x53, 0xa6, 0x73, 0xfd, 0xab, 0xb7, 0xd2, 0x69, 0xaa, 0xd5, 0xf6, 0x15, 0x9c, 0x43, 0x7f,
	0x66, 0x9a, 0x57, 0x3c, 0x15, 0xd9, 0x8d, 0x8e, 0xf4, 0x3a, 0x27, 0x5f, 0xc0, 0x91, 0xa1, 0xdf,
	0x18, 0x36, 0x46, 0x2e, 0x73, 0x64, 0x48, 0xfe, 0x00, 0x24, 0x3a, 0x92, 0x6a, 0x11, 0xe7, 0x22,
	0xf3, 0x9d, 0x61, 0x63, 0xf4, 0x89, 0xed, 0x55, 0x82, 0x53, 0xf8, 0x8e, 0x1e, 0xb7, 0x62, 0x83,
	0x36, 0x97, 0x99, 0x14, 0x2a, 0x64, 0xe2, 0x99, 0xf4, 0xa1, 0x19, 0x95, 0x56, 0xe6, 0xd3, 0x54,
	0xb4, 0x0c, 0xd1, 0xc4, 0x65, 0xe6, 0x33, 0x38, 0x86, 0xaf, 0x16, 0xfa, 0x43, 0xe8, 0x14, 0x06,
	0x35, 0xf4, 0x22, 0xe6, 0xb9, 0x05, 0xfc, 0x0d, 0xad, 0x55, 0xcc, 0x73, 0x24, 0xbd, 0x71, 0x97,
	0x1a, 0xd9, 0x94, 0xe7, 0x9c, 0x61, 0x39, 0xf8, 0x5b, 0xb9, 0x2c, 0xef, 0x96, 0x4c, 0xbc, 0xac,
	0x13, 0x8b, 0x4b, 0x70, 0x0d, 0x3d, 0x94, 0xcd, 0x79, 0xbe, 0x8a, 0x17, 0x8f, 0x96, 0x39, 0x23,
	0xe8, 0xa6, 0xa6, 0x39, 0x53, 0x0f, 0xba, 0x18, 0x06, 0x74, 0xbe, 0xab, 0xb0, 0xaa, 0x19, 0x9c,
	0x1d, 0x2c, 0x3e, 0xe7, 0x32, 0xb1, 0x18, 0xfe, 0x84, 0x56, 0xca, 0x65, 0x52, 0x78, 0xb9, 0xd4,
	0xc8, 0x18, 0x96, 0x82, 0x57, 0x07, 0x5c, 0xb4, 0x20, 0x13, 0xf0, 0xa2, 0x2a, 0x3c, 0xc4, 0xbd,
	0xf1, 0x80, 0x1e, 0xa6, 0xca, 0xf6, 0x55, 0xe4, 0x04, 0x7a, 0x4a, 0x6c, 0xca, 0xa0, 0x8a, 0x09,
	0x3f, 0xa8, 0x3d, 0x47, 0x56, 0x13, 0x93, 0x31, 0x74, 0xcb, 0xb7, 0xdf, 0x44, 0xf2, 0x9b, 0x95,
	0xac, 0x64, 0xe4, 0x3f, 0x74, 0x94, 0xd8, 0x98, 0xcb, 0xfb, 0x2d, 0x24, 0x08, 0x7d, 0x17, 0x1d,
	0xdb, 0x49, 0xc8, 0x11, 0xae, 0x57, 0xa6, 0xe1, 0xbb, 0x07, 0x48, 0xd9, 0x61, 0x35, 0x1d, 0xf9,
	0x07, 0x9d, 0x74, 0x1b, 0x8f, 0xdf, 0x46, 0xe4, 0x33, 0xdd, 0xcf, 0x8c, 0xed, 0xba, 0xc5, 0x80,
	0xf2, 0xf6, 0x7e, 0xc7, 0xb6, 0x13, 0x9e, 0xbb, 0xa6, 0xbb, 0x6f, 0xe3, 0x5f, 0x33, 0x79, 0x0b,
	0x00, 0x00, 0xff, 0xff, 0x7b, 0xa9, 0x1a, 0x39, 0x52, 0x03, 0x00, 0x00,
}