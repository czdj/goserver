// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server.proto

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

type ServerInfo struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	CmdAddr              string   `protobuf:"bytes,4,opt,name=cmdAddr,proto3" json:"cmdAddr,omitempty"`
	RpcAddr              string   `protobuf:"bytes,5,opt,name=rpcAddr,proto3" json:"rpcAddr,omitempty"`
	ExtAddr              string   `protobuf:"bytes,6,opt,name=extAddr,proto3" json:"extAddr,omitempty"`
	LastTick             int64    `protobuf:"varint,7,opt,name=lastTick,proto3" json:"lastTick,omitempty"`
	Type                 string   `protobuf:"bytes,8,opt,name=type,proto3" json:"type,omitempty"`
	Load                 int32    `protobuf:"varint,9,opt,name=load,proto3" json:"load,omitempty"`
	Conf                 string   `protobuf:"bytes,10,opt,name=conf,proto3" json:"conf,omitempty"`
	OnlineCount          int32    `protobuf:"varint,11,opt,name=onlineCount,proto3" json:"onlineCount,omitempty"`
	Md5                  string   `protobuf:"bytes,12,opt,name=md5,proto3" json:"md5,omitempty"`
	PbVer                string   `protobuf:"bytes,13,opt,name=pbVer,proto3" json:"pbVer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerInfo) Reset()         { *m = ServerInfo{} }
func (m *ServerInfo) String() string { return proto.CompactTextString(m) }
func (*ServerInfo) ProtoMessage()    {}
func (*ServerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{0}
}

func (m *ServerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerInfo.Unmarshal(m, b)
}
func (m *ServerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerInfo.Marshal(b, m, deterministic)
}
func (m *ServerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerInfo.Merge(m, src)
}
func (m *ServerInfo) XXX_Size() int {
	return xxx_messageInfo_ServerInfo.Size(m)
}
func (m *ServerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ServerInfo proto.InternalMessageInfo

func (m *ServerInfo) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ServerInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServerInfo) GetCmdAddr() string {
	if m != nil {
		return m.CmdAddr
	}
	return ""
}

func (m *ServerInfo) GetRpcAddr() string {
	if m != nil {
		return m.RpcAddr
	}
	return ""
}

func (m *ServerInfo) GetExtAddr() string {
	if m != nil {
		return m.ExtAddr
	}
	return ""
}

func (m *ServerInfo) GetLastTick() int64 {
	if m != nil {
		return m.LastTick
	}
	return 0
}

func (m *ServerInfo) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ServerInfo) GetLoad() int32 {
	if m != nil {
		return m.Load
	}
	return 0
}

func (m *ServerInfo) GetConf() string {
	if m != nil {
		return m.Conf
	}
	return ""
}

func (m *ServerInfo) GetOnlineCount() int32 {
	if m != nil {
		return m.OnlineCount
	}
	return 0
}

func (m *ServerInfo) GetMd5() string {
	if m != nil {
		return m.Md5
	}
	return ""
}

func (m *ServerInfo) GetPbVer() string {
	if m != nil {
		return m.PbVer
	}
	return ""
}

type ServerWarn struct {
	BadRpc               []int32  `protobuf:"varint,1,rep,packed,name=badRpc,proto3" json:"badRpc,omitempty"`
	LastPanic            int64    `protobuf:"varint,2,opt,name=lastPanic,proto3" json:"lastPanic,omitempty"`
	PanicCount           int32    `protobuf:"varint,3,opt,name=panicCount,proto3" json:"panicCount,omitempty"`
	GoScale              bool     `protobuf:"varint,4,opt,name=goScale,proto3" json:"goScale,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerWarn) Reset()         { *m = ServerWarn{} }
func (m *ServerWarn) String() string { return proto.CompactTextString(m) }
func (*ServerWarn) ProtoMessage()    {}
func (*ServerWarn) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{1}
}

func (m *ServerWarn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerWarn.Unmarshal(m, b)
}
func (m *ServerWarn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerWarn.Marshal(b, m, deterministic)
}
func (m *ServerWarn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerWarn.Merge(m, src)
}
func (m *ServerWarn) XXX_Size() int {
	return xxx_messageInfo_ServerWarn.Size(m)
}
func (m *ServerWarn) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerWarn.DiscardUnknown(m)
}

var xxx_messageInfo_ServerWarn proto.InternalMessageInfo

func (m *ServerWarn) GetBadRpc() []int32 {
	if m != nil {
		return m.BadRpc
	}
	return nil
}

func (m *ServerWarn) GetLastPanic() int64 {
	if m != nil {
		return m.LastPanic
	}
	return 0
}

func (m *ServerWarn) GetPanicCount() int32 {
	if m != nil {
		return m.PanicCount
	}
	return 0
}

func (m *ServerWarn) GetGoScale() bool {
	if m != nil {
		return m.GoScale
	}
	return false
}

type ServerCmd struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	GlobalPath           string   `protobuf:"bytes,2,opt,name=globalPath,proto3" json:"globalPath,omitempty"`
	LogPath              string   `protobuf:"bytes,3,opt,name=logPath,proto3" json:"logPath,omitempty"`
	LogSign              string   `protobuf:"bytes,4,opt,name=logSign,proto3" json:"logSign,omitempty"`
	Daemon               bool     `protobuf:"varint,5,opt,name=daemon,proto3" json:"daemon,omitempty"`
	Lan                  bool     `protobuf:"varint,6,opt,name=lan,proto3" json:"lan,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerCmd) Reset()         { *m = ServerCmd{} }
func (m *ServerCmd) String() string { return proto.CompactTextString(m) }
func (*ServerCmd) ProtoMessage()    {}
func (*ServerCmd) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{2}
}

func (m *ServerCmd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerCmd.Unmarshal(m, b)
}
func (m *ServerCmd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerCmd.Marshal(b, m, deterministic)
}
func (m *ServerCmd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerCmd.Merge(m, src)
}
func (m *ServerCmd) XXX_Size() int {
	return xxx_messageInfo_ServerCmd.Size(m)
}
func (m *ServerCmd) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerCmd.DiscardUnknown(m)
}

var xxx_messageInfo_ServerCmd proto.InternalMessageInfo

func (m *ServerCmd) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *ServerCmd) GetGlobalPath() string {
	if m != nil {
		return m.GlobalPath
	}
	return ""
}

func (m *ServerCmd) GetLogPath() string {
	if m != nil {
		return m.LogPath
	}
	return ""
}

func (m *ServerCmd) GetLogSign() string {
	if m != nil {
		return m.LogSign
	}
	return ""
}

func (m *ServerCmd) GetDaemon() bool {
	if m != nil {
		return m.Daemon
	}
	return false
}

func (m *ServerCmd) GetLan() bool {
	if m != nil {
		return m.Lan
	}
	return false
}

func init() {
	proto.RegisterType((*ServerInfo)(nil), "ServerInfo")
	proto.RegisterType((*ServerWarn)(nil), "ServerWarn")
	proto.RegisterType((*ServerCmd)(nil), "ServerCmd")
}

func init() { proto.RegisterFile("server.proto", fileDescriptor_ad098daeda4239f7) }

var fileDescriptor_ad098daeda4239f7 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x92, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0x69, 0xbb, 0x6e, 0xed, 0xbb, 0x29, 0x12, 0x44, 0x82, 0x88, 0x94, 0x9d, 0x7a, 0xf2,
	0x22, 0x7e, 0x00, 0xd9, 0xc9, 0xdb, 0xc8, 0x44, 0xcf, 0x69, 0x93, 0xd5, 0x62, 0x9a, 0x94, 0xac,
	0x8a, 0x82, 0x5f, 0x45, 0xf0, 0xa3, 0xca, 0xfb, 0xa6, 0xd5, 0xdd, 0x9e, 0xe7, 0xf9, 0xa5, 0xc9,
	0xfb, 0xa7, 0xb0, 0x3a, 0x68, 0xff, 0xae, 0xfd, 0x4d, 0xef, 0xdd, 0xe0, 0xd6, 0x3f, 0x31, 0xc0,
	0x8e, 0x82, 0x07, 0xbb, 0x77, 0xec, 0x14, 0xe2, 0x56, 0xf1, 0xa8, 0x88, 0xca, 0x54, 0xc4, 0xad,
	0x62, 0x0c, 0x66, 0x56, 0x76, 0x9a, 0x27, 0x45, 0x54, 0xe6, 0x82, 0x34, 0xe3, 0xb0, 0xa8, 0x3b,
	0x75, 0xaf, 0x94, 0xe7, 0x33, 0x8a, 0x27, 0x8b, 0xc4, 0xf7, 0x35, 0x91, 0x34, 0x90, 0xd1, 0x22,
	0xd1, 0x1f, 0x03, 0x91, 0x79, 0x20, 0xa3, 0x65, 0x97, 0x90, 0x19, 0x79, 0x18, 0x1e, 0xdb, 0xfa,
	0x95, 0x2f, 0x8a, 0xa8, 0x4c, 0xc4, 0x9f, 0xc7, 0xd7, 0x87, 0xcf, 0x5e, 0xf3, 0x2c, 0xbc, 0x8e,
	0x1a, 0x33, 0xe3, 0xa4, 0xe2, 0x39, 0xd5, 0x48, 0x1a, 0xb3, 0xda, 0xd9, 0x3d, 0x87, 0x70, 0x0e,
	0x35, 0x2b, 0x60, 0xe9, 0xac, 0x69, 0xad, 0xde, 0xb8, 0x37, 0x3b, 0xf0, 0x25, 0x1d, 0x3f, 0x8e,
	0xd8, 0x19, 0x24, 0x9d, 0xba, 0xe3, 0x2b, 0xfa, 0x08, 0x25, 0x3b, 0x87, 0xb4, 0xaf, 0x9e, 0xb4,
	0xe7, 0x27, 0x94, 0x05, 0xb3, 0xfe, 0x9a, 0x26, 0xf4, 0x2c, 0xbd, 0x65, 0x17, 0x30, 0xaf, 0xa4,
	0x12, 0x7d, 0xcd, 0xa3, 0x22, 0x29, 0x53, 0x31, 0x3a, 0x76, 0x05, 0x39, 0xd6, 0xbd, 0x95, 0xb6,
	0xad, 0x79, 0x4c, 0x8d, 0xfc, 0x07, 0xec, 0x1a, 0xa0, 0x47, 0x11, 0x8a, 0x49, 0xa8, 0x98, 0xa3,
	0x04, 0xe7, 0xd3, 0xb8, 0x5d, 0x2d, 0x8d, 0xa6, 0x99, 0x66, 0x62, 0xb2, 0xeb, 0xef, 0x08, 0xf2,
	0xf0, 0xfc, 0xa6, 0xa3, 0x4e, 0x7b, 0x39, 0xbc, 0xd0, 0x86, 0x72, 0x41, 0x1a, 0xef, 0x6e, 0x8c,
	0xab, 0xa4, 0xd9, 0x22, 0x89, 0x89, 0x1c, 0x25, 0x78, 0xb7, 0x71, 0x0d, 0xc1, 0xb0, 0xc6, 0xc9,
	0x8e, 0x64, 0xd7, 0x36, 0x76, 0xda, 0xe4, 0x68, 0xb1, 0x4b, 0x25, 0x75, 0xe7, 0x2c, 0x2d, 0x32,
	0x13, 0xa3, 0xc3, 0x99, 0x19, 0x69, 0x69, 0x87, 0x99, 0x40, 0x59, 0xcd, 0xe9, 0x3f, 0xba, 0xfd,
	0x0d, 0x00, 0x00, 0xff, 0xff, 0xa2, 0xd2, 0x99, 0x9c, 0x57, 0x02, 0x00, 0x00,
}
