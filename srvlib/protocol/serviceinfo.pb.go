// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.19.4
// source: serviceinfo.proto

package protocol

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ServiceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AreaId          int32    `protobuf:"varint,1,opt,name=AreaId,proto3" json:"AreaId,omitempty"`
	SrvId           int32    `protobuf:"varint,2,opt,name=SrvId,proto3" json:"SrvId,omitempty"`
	SrvType         int32    `protobuf:"varint,3,opt,name=SrvType,proto3" json:"SrvType,omitempty"`
	SrvPID          int32    `protobuf:"varint,4,opt,name=SrvPID,proto3" json:"SrvPID,omitempty"`
	SrvName         string   `protobuf:"bytes,5,opt,name=SrvName,proto3" json:"SrvName,omitempty"`
	NetworkType     string   `protobuf:"bytes,6,opt,name=NetworkType,proto3" json:"NetworkType,omitempty"`
	Ip              string   `protobuf:"bytes,7,opt,name=Ip,proto3" json:"Ip,omitempty"`
	Port            int32    `protobuf:"varint,8,opt,name=Port,proto3" json:"Port,omitempty"`
	WriteTimeOut    int32    `protobuf:"varint,9,opt,name=WriteTimeOut,proto3" json:"WriteTimeOut,omitempty"`
	ReadTimeOut     int32    `protobuf:"varint,10,opt,name=ReadTimeOut,proto3" json:"ReadTimeOut,omitempty"`
	IdleTimeOut     int32    `protobuf:"varint,11,opt,name=IdleTimeOut,proto3" json:"IdleTimeOut,omitempty"`
	MaxDone         int32    `protobuf:"varint,12,opt,name=MaxDone,proto3" json:"MaxDone,omitempty"`
	MaxPend         int32    `protobuf:"varint,13,opt,name=MaxPend,proto3" json:"MaxPend,omitempty"`
	MaxPacket       int32    `protobuf:"varint,14,opt,name=MaxPacket,proto3" json:"MaxPacket,omitempty"`
	RcvBuff         int32    `protobuf:"varint,15,opt,name=RcvBuff,proto3" json:"RcvBuff,omitempty"`
	SndBuff         int32    `protobuf:"varint,16,opt,name=SndBuff,proto3" json:"SndBuff,omitempty"`
	SoLinger        int32    `protobuf:"varint,17,opt,name=SoLinger,proto3" json:"SoLinger,omitempty"`
	IsAuth          bool     `protobuf:"varint,18,opt,name=IsAuth,proto3" json:"IsAuth,omitempty"`
	KeepAlive       bool     `protobuf:"varint,19,opt,name=KeepAlive,proto3" json:"KeepAlive,omitempty"`
	NoDelay         bool     `protobuf:"varint,20,opt,name=NoDelay,proto3" json:"NoDelay,omitempty"`
	IsAutoReconn    bool     `protobuf:"varint,21,opt,name=IsAutoReconn,proto3" json:"IsAutoReconn,omitempty"`
	IsInnerLink     bool     `protobuf:"varint,22,opt,name=IsInnerLink,proto3" json:"IsInnerLink,omitempty"`
	SupportFragment bool     `protobuf:"varint,23,opt,name=SupportFragment,proto3" json:"SupportFragment,omitempty"`
	AllowMultiConn  bool     `protobuf:"varint,24,opt,name=AllowMultiConn,proto3" json:"AllowMultiConn,omitempty"`
	AuthKey         string   `protobuf:"bytes,25,opt,name=AuthKey,proto3" json:"AuthKey,omitempty"`
	EncoderName     string   `protobuf:"bytes,26,opt,name=EncoderName,proto3" json:"EncoderName,omitempty"`
	DecoderName     string   `protobuf:"bytes,27,opt,name=DecoderName,proto3" json:"DecoderName,omitempty"`
	FilterChain     []string `protobuf:"bytes,28,rep,name=FilterChain,proto3" json:"FilterChain,omitempty"`
	HandlerChain    []string `protobuf:"bytes,29,rep,name=HandlerChain,proto3" json:"HandlerChain,omitempty"`
	Protocol        string   `protobuf:"bytes,30,opt,name=Protocol,proto3" json:"Protocol,omitempty"`
	Path            string   `protobuf:"bytes,31,opt,name=Path,proto3" json:"Path,omitempty"`
	OuterIp         string   `protobuf:"bytes,32,opt,name=OuterIp,proto3" json:"OuterIp,omitempty"`
}

func (x *ServiceInfo) Reset() {
	*x = ServiceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serviceinfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceInfo) ProtoMessage() {}

func (x *ServiceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_serviceinfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceInfo.ProtoReflect.Descriptor instead.
func (*ServiceInfo) Descriptor() ([]byte, []int) {
	return file_serviceinfo_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceInfo) GetAreaId() int32 {
	if x != nil {
		return x.AreaId
	}
	return 0
}

func (x *ServiceInfo) GetSrvId() int32 {
	if x != nil {
		return x.SrvId
	}
	return 0
}

func (x *ServiceInfo) GetSrvType() int32 {
	if x != nil {
		return x.SrvType
	}
	return 0
}

func (x *ServiceInfo) GetSrvPID() int32 {
	if x != nil {
		return x.SrvPID
	}
	return 0
}

func (x *ServiceInfo) GetSrvName() string {
	if x != nil {
		return x.SrvName
	}
	return ""
}

func (x *ServiceInfo) GetNetworkType() string {
	if x != nil {
		return x.NetworkType
	}
	return ""
}

func (x *ServiceInfo) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *ServiceInfo) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *ServiceInfo) GetWriteTimeOut() int32 {
	if x != nil {
		return x.WriteTimeOut
	}
	return 0
}

func (x *ServiceInfo) GetReadTimeOut() int32 {
	if x != nil {
		return x.ReadTimeOut
	}
	return 0
}

func (x *ServiceInfo) GetIdleTimeOut() int32 {
	if x != nil {
		return x.IdleTimeOut
	}
	return 0
}

func (x *ServiceInfo) GetMaxDone() int32 {
	if x != nil {
		return x.MaxDone
	}
	return 0
}

func (x *ServiceInfo) GetMaxPend() int32 {
	if x != nil {
		return x.MaxPend
	}
	return 0
}

func (x *ServiceInfo) GetMaxPacket() int32 {
	if x != nil {
		return x.MaxPacket
	}
	return 0
}

func (x *ServiceInfo) GetRcvBuff() int32 {
	if x != nil {
		return x.RcvBuff
	}
	return 0
}

func (x *ServiceInfo) GetSndBuff() int32 {
	if x != nil {
		return x.SndBuff
	}
	return 0
}

func (x *ServiceInfo) GetSoLinger() int32 {
	if x != nil {
		return x.SoLinger
	}
	return 0
}

func (x *ServiceInfo) GetIsAuth() bool {
	if x != nil {
		return x.IsAuth
	}
	return false
}

func (x *ServiceInfo) GetKeepAlive() bool {
	if x != nil {
		return x.KeepAlive
	}
	return false
}

func (x *ServiceInfo) GetNoDelay() bool {
	if x != nil {
		return x.NoDelay
	}
	return false
}

func (x *ServiceInfo) GetIsAutoReconn() bool {
	if x != nil {
		return x.IsAutoReconn
	}
	return false
}

func (x *ServiceInfo) GetIsInnerLink() bool {
	if x != nil {
		return x.IsInnerLink
	}
	return false
}

func (x *ServiceInfo) GetSupportFragment() bool {
	if x != nil {
		return x.SupportFragment
	}
	return false
}

func (x *ServiceInfo) GetAllowMultiConn() bool {
	if x != nil {
		return x.AllowMultiConn
	}
	return false
}

func (x *ServiceInfo) GetAuthKey() string {
	if x != nil {
		return x.AuthKey
	}
	return ""
}

func (x *ServiceInfo) GetEncoderName() string {
	if x != nil {
		return x.EncoderName
	}
	return ""
}

func (x *ServiceInfo) GetDecoderName() string {
	if x != nil {
		return x.DecoderName
	}
	return ""
}

func (x *ServiceInfo) GetFilterChain() []string {
	if x != nil {
		return x.FilterChain
	}
	return nil
}

func (x *ServiceInfo) GetHandlerChain() []string {
	if x != nil {
		return x.HandlerChain
	}
	return nil
}

func (x *ServiceInfo) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *ServiceInfo) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ServiceInfo) GetOuterIp() string {
	if x != nil {
		return x.OuterIp
	}
	return ""
}

type SSServiceRegiste struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Services []*ServiceInfo `protobuf:"bytes,1,rep,name=Services,proto3" json:"Services,omitempty"`
}

func (x *SSServiceRegiste) Reset() {
	*x = SSServiceRegiste{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serviceinfo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSServiceRegiste) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSServiceRegiste) ProtoMessage() {}

func (x *SSServiceRegiste) ProtoReflect() protoreflect.Message {
	mi := &file_serviceinfo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSServiceRegiste.ProtoReflect.Descriptor instead.
func (*SSServiceRegiste) Descriptor() ([]byte, []int) {
	return file_serviceinfo_proto_rawDescGZIP(), []int{1}
}

func (x *SSServiceRegiste) GetServices() []*ServiceInfo {
	if x != nil {
		return x.Services
	}
	return nil
}

type SSServiceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service *ServiceInfo `protobuf:"bytes,1,opt,name=Service,proto3" json:"Service,omitempty"`
}

func (x *SSServiceInfo) Reset() {
	*x = SSServiceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serviceinfo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSServiceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSServiceInfo) ProtoMessage() {}

func (x *SSServiceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_serviceinfo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSServiceInfo.ProtoReflect.Descriptor instead.
func (*SSServiceInfo) Descriptor() ([]byte, []int) {
	return file_serviceinfo_proto_rawDescGZIP(), []int{2}
}

func (x *SSServiceInfo) GetService() *ServiceInfo {
	if x != nil {
		return x.Service
	}
	return nil
}

type SSServiceShut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service *ServiceInfo `protobuf:"bytes,1,opt,name=Service,proto3" json:"Service,omitempty"`
}

func (x *SSServiceShut) Reset() {
	*x = SSServiceShut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serviceinfo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSServiceShut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSServiceShut) ProtoMessage() {}

func (x *SSServiceShut) ProtoReflect() protoreflect.Message {
	mi := &file_serviceinfo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSServiceShut.ProtoReflect.Descriptor instead.
func (*SSServiceShut) Descriptor() ([]byte, []int) {
	return file_serviceinfo_proto_rawDescGZIP(), []int{3}
}

func (x *SSServiceShut) GetService() *ServiceInfo {
	if x != nil {
		return x.Service
	}
	return nil
}

var File_serviceinfo_proto protoreflect.FileDescriptor

var file_serviceinfo_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0xad, 0x07,
	0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a,
	0x06, 0x41, 0x72, 0x65, 0x61, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x41,
	0x72, 0x65, 0x61, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x72, 0x76, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x53, 0x72, 0x76, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x53,
	0x72, 0x76, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x53, 0x72,
	0x76, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x72, 0x76, 0x50, 0x49, 0x44, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x53, 0x72, 0x76, 0x50, 0x49, 0x44, 0x12, 0x18, 0x0a,
	0x07, 0x53, 0x72, 0x76, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x53, 0x72, 0x76, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x70, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x6f, 0x72,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x22, 0x0a,
	0x0c, 0x57, 0x72, 0x69, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x4f, 0x75, 0x74, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x57, 0x72, 0x69, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x4f, 0x75,
	0x74, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x4f, 0x75, 0x74,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x52, 0x65, 0x61, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x4f, 0x75, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x64, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x4f,
	0x75, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x49, 0x64, 0x6c, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x4f, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x61, 0x78, 0x44, 0x6f, 0x6e, 0x65,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x4d, 0x61, 0x78, 0x44, 0x6f, 0x6e, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x4d, 0x61, 0x78, 0x50, 0x65, 0x6e, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x4d, 0x61, 0x78, 0x50, 0x65, 0x6e, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x61, 0x78,
	0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x4d, 0x61,
	0x78, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x63, 0x76, 0x42, 0x75,
	0x66, 0x66, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x52, 0x63, 0x76, 0x42, 0x75, 0x66,
	0x66, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x6e, 0x64, 0x42, 0x75, 0x66, 0x66, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x53, 0x6e, 0x64, 0x42, 0x75, 0x66, 0x66, 0x12, 0x1a, 0x0a, 0x08, 0x53,
	0x6f, 0x4c, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x53,
	0x6f, 0x4c, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x73, 0x41, 0x75, 0x74,
	0x68, 0x18, 0x12, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x49, 0x73, 0x41, 0x75, 0x74, 0x68, 0x12,
	0x1c, 0x0a, 0x09, 0x4b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x18, 0x13, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x4b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x4e, 0x6f, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x14, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x4e, 0x6f, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x49, 0x73, 0x41, 0x75, 0x74,
	0x6f, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x6e, 0x18, 0x15, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x49,
	0x73, 0x41, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x49,
	0x73, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x16, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x49, 0x73, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x28, 0x0a,
	0x0f, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x17, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x46,
	0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x41, 0x6c, 0x6c, 0x6f, 0x77,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x43, 0x6f, 0x6e, 0x6e, 0x18, 0x18, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0e, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x43, 0x6f, 0x6e, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x45, 0x6e, 0x63,
	0x6f, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44,
	0x65, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x1c, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0b, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12,
	0x22, 0x0a, 0x0c, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x18,
	0x1d, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x43, 0x68,
	0x61, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18,
	0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x50, 0x61, 0x74, 0x68, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x75, 0x74, 0x65, 0x72, 0x49, 0x70, 0x18, 0x20,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4f, 0x75, 0x74, 0x65, 0x72, 0x49, 0x70, 0x22, 0x45, 0x0a,
	0x10, 0x53, 0x53, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x12, 0x31, 0x0a, 0x08, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x22, 0x40, 0x0a, 0x0d, 0x53, 0x53, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2f, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x40, 0x0a, 0x0d, 0x53, 0x53, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x53, 0x68, 0x75, 0x74, 0x12, 0x2f, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_serviceinfo_proto_rawDescOnce sync.Once
	file_serviceinfo_proto_rawDescData = file_serviceinfo_proto_rawDesc
)

func file_serviceinfo_proto_rawDescGZIP() []byte {
	file_serviceinfo_proto_rawDescOnce.Do(func() {
		file_serviceinfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_serviceinfo_proto_rawDescData)
	})
	return file_serviceinfo_proto_rawDescData
}

var file_serviceinfo_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_serviceinfo_proto_goTypes = []interface{}{
	(*ServiceInfo)(nil),      // 0: protocol.ServiceInfo
	(*SSServiceRegiste)(nil), // 1: protocol.SSServiceRegiste
	(*SSServiceInfo)(nil),    // 2: protocol.SSServiceInfo
	(*SSServiceShut)(nil),    // 3: protocol.SSServiceShut
}
var file_serviceinfo_proto_depIdxs = []int32{
	0, // 0: protocol.SSServiceRegiste.Services:type_name -> protocol.ServiceInfo
	0, // 1: protocol.SSServiceInfo.Service:type_name -> protocol.ServiceInfo
	0, // 2: protocol.SSServiceShut.Service:type_name -> protocol.ServiceInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_serviceinfo_proto_init() }
func file_serviceinfo_proto_init() {
	if File_serviceinfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_serviceinfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_serviceinfo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSServiceRegiste); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_serviceinfo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSServiceInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_serviceinfo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSServiceShut); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_serviceinfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_serviceinfo_proto_goTypes,
		DependencyIndexes: file_serviceinfo_proto_depIdxs,
		MessageInfos:      file_serviceinfo_proto_msgTypes,
	}.Build()
	File_serviceinfo_proto = out.File
	file_serviceinfo_proto_rawDesc = nil
	file_serviceinfo_proto_goTypes = nil
	file_serviceinfo_proto_depIdxs = nil
}
