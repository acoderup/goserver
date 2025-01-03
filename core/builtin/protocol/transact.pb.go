// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.19.4
// source: transact.proto

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

type TransactStart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MyTNP      *TransactParam `protobuf:"bytes,1,opt,name=MyTNP,proto3" json:"MyTNP,omitempty"`
	ParenTNP   *TransactParam `protobuf:"bytes,2,opt,name=ParenTNP,proto3" json:"ParenTNP,omitempty"`
	CustomData []byte         `protobuf:"bytes,3,opt,name=CustomData,proto3" json:"CustomData,omitempty"`
}

func (x *TransactStart) Reset() {
	*x = TransactStart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transact_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactStart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactStart) ProtoMessage() {}

func (x *TransactStart) ProtoReflect() protoreflect.Message {
	mi := &file_transact_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactStart.ProtoReflect.Descriptor instead.
func (*TransactStart) Descriptor() ([]byte, []int) {
	return file_transact_proto_rawDescGZIP(), []int{0}
}

func (x *TransactStart) GetMyTNP() *TransactParam {
	if x != nil {
		return x.MyTNP
	}
	return nil
}

func (x *TransactStart) GetParenTNP() *TransactParam {
	if x != nil {
		return x.ParenTNP
	}
	return nil
}

func (x *TransactStart) GetCustomData() []byte {
	if x != nil {
		return x.CustomData
	}
	return nil
}

type TransactCtrlCmd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TId int64 `protobuf:"varint,1,opt,name=TId,proto3" json:"TId,omitempty"`
	Cmd int32 `protobuf:"varint,2,opt,name=Cmd,proto3" json:"Cmd,omitempty"`
}

func (x *TransactCtrlCmd) Reset() {
	*x = TransactCtrlCmd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transact_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactCtrlCmd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactCtrlCmd) ProtoMessage() {}

func (x *TransactCtrlCmd) ProtoReflect() protoreflect.Message {
	mi := &file_transact_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactCtrlCmd.ProtoReflect.Descriptor instead.
func (*TransactCtrlCmd) Descriptor() ([]byte, []int) {
	return file_transact_proto_rawDescGZIP(), []int{1}
}

func (x *TransactCtrlCmd) GetTId() int64 {
	if x != nil {
		return x.TId
	}
	return 0
}

func (x *TransactCtrlCmd) GetCmd() int32 {
	if x != nil {
		return x.Cmd
	}
	return 0
}

type TransactResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MyTId      int64  `protobuf:"varint,1,opt,name=MyTId,proto3" json:"MyTId,omitempty"`
	ChildTId   int64  `protobuf:"varint,2,opt,name=ChildTId,proto3" json:"ChildTId,omitempty"`
	RetCode    int32  `protobuf:"varint,3,opt,name=RetCode,proto3" json:"RetCode,omitempty"`
	CustomData []byte `protobuf:"bytes,4,opt,name=CustomData,proto3" json:"CustomData,omitempty"`
}

func (x *TransactResult) Reset() {
	*x = TransactResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transact_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactResult) ProtoMessage() {}

func (x *TransactResult) ProtoReflect() protoreflect.Message {
	mi := &file_transact_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactResult.ProtoReflect.Descriptor instead.
func (*TransactResult) Descriptor() ([]byte, []int) {
	return file_transact_proto_rawDescGZIP(), []int{2}
}

func (x *TransactResult) GetMyTId() int64 {
	if x != nil {
		return x.MyTId
	}
	return 0
}

func (x *TransactResult) GetChildTId() int64 {
	if x != nil {
		return x.ChildTId
	}
	return 0
}

func (x *TransactResult) GetRetCode() int32 {
	if x != nil {
		return x.RetCode
	}
	return 0
}

func (x *TransactResult) GetCustomData() []byte {
	if x != nil {
		return x.CustomData
	}
	return nil
}

type TransactParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransNodeID     int64 `protobuf:"varint,1,opt,name=TransNodeID,proto3" json:"TransNodeID,omitempty"`
	TransType       int32 `protobuf:"varint,2,opt,name=TransType,proto3" json:"TransType,omitempty"`
	OwnerType       int32 `protobuf:"varint,3,opt,name=OwnerType,proto3" json:"OwnerType,omitempty"`
	OwnerID         int32 `protobuf:"varint,4,opt,name=OwnerID,proto3" json:"OwnerID,omitempty"`
	SkeletonID      int32 `protobuf:"varint,5,opt,name=SkeletonID,proto3" json:"SkeletonID,omitempty"`
	LevelNo         int32 `protobuf:"varint,6,opt,name=LevelNo,proto3" json:"LevelNo,omitempty"`
	AreaID          int32 `protobuf:"varint,7,opt,name=AreaID,proto3" json:"AreaID,omitempty"`
	TimeOut         int64 `protobuf:"varint,8,opt,name=TimeOut,proto3" json:"TimeOut,omitempty"`
	TransCommitType int32 `protobuf:"varint,9,opt,name=TransCommitType,proto3" json:"TransCommitType,omitempty"`
	ExpiresTs       int64 `protobuf:"varint,10,opt,name=ExpiresTs,proto3" json:"ExpiresTs,omitempty"`
}

func (x *TransactParam) Reset() {
	*x = TransactParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transact_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactParam) ProtoMessage() {}

func (x *TransactParam) ProtoReflect() protoreflect.Message {
	mi := &file_transact_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactParam.ProtoReflect.Descriptor instead.
func (*TransactParam) Descriptor() ([]byte, []int) {
	return file_transact_proto_rawDescGZIP(), []int{3}
}

func (x *TransactParam) GetTransNodeID() int64 {
	if x != nil {
		return x.TransNodeID
	}
	return 0
}

func (x *TransactParam) GetTransType() int32 {
	if x != nil {
		return x.TransType
	}
	return 0
}

func (x *TransactParam) GetOwnerType() int32 {
	if x != nil {
		return x.OwnerType
	}
	return 0
}

func (x *TransactParam) GetOwnerID() int32 {
	if x != nil {
		return x.OwnerID
	}
	return 0
}

func (x *TransactParam) GetSkeletonID() int32 {
	if x != nil {
		return x.SkeletonID
	}
	return 0
}

func (x *TransactParam) GetLevelNo() int32 {
	if x != nil {
		return x.LevelNo
	}
	return 0
}

func (x *TransactParam) GetAreaID() int32 {
	if x != nil {
		return x.AreaID
	}
	return 0
}

func (x *TransactParam) GetTimeOut() int64 {
	if x != nil {
		return x.TimeOut
	}
	return 0
}

func (x *TransactParam) GetTransCommitType() int32 {
	if x != nil {
		return x.TransCommitType
	}
	return 0
}

func (x *TransactParam) GetExpiresTs() int64 {
	if x != nil {
		return x.ExpiresTs
	}
	return 0
}

var File_transact_proto protoreflect.FileDescriptor

var file_transact_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x93, 0x01, 0x0a, 0x0d, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x2d, 0x0a, 0x05,
	0x4d, 0x79, 0x54, 0x4e, 0x50, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x52, 0x05, 0x4d, 0x79, 0x54, 0x4e, 0x50, 0x12, 0x33, 0x0a, 0x08, 0x50,
	0x61, 0x72, 0x65, 0x6e, 0x54, 0x4e, 0x50, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x08, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x54, 0x4e, 0x50,
	0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x61, 0x74, 0x61,
	0x22, 0x35, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x43, 0x74, 0x72, 0x6c,
	0x43, 0x6d, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x54, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x54, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x43, 0x6d, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x43, 0x6d, 0x64, 0x22, 0x7c, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x79, 0x54,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x4d, 0x79, 0x54, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x54, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x54, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x52,
	0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x52, 0x65,
	0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x44, 0x61, 0x74, 0x61, 0x22, 0xbb, 0x02, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x4e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x4f, 0x77, 0x6e, 0x65,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x1e, 0x0a, 0x0a, 0x53, 0x6b, 0x65, 0x6c, 0x65, 0x74, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x53, 0x6b, 0x65, 0x6c, 0x65, 0x74, 0x6f, 0x6e, 0x49, 0x44, 0x12,
	0x18, 0x0a, 0x07, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4e, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4e, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x72, 0x65,
	0x61, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x41, 0x72, 0x65, 0x61, 0x49,
	0x44, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x69, 0x6d, 0x65, 0x4f, 0x75, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x54, 0x69, 0x6d, 0x65, 0x4f, 0x75, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73,
	0x54, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x73, 0x54, 0x73, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transact_proto_rawDescOnce sync.Once
	file_transact_proto_rawDescData = file_transact_proto_rawDesc
)

func file_transact_proto_rawDescGZIP() []byte {
	file_transact_proto_rawDescOnce.Do(func() {
		file_transact_proto_rawDescData = protoimpl.X.CompressGZIP(file_transact_proto_rawDescData)
	})
	return file_transact_proto_rawDescData
}

var file_transact_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_transact_proto_goTypes = []interface{}{
	(*TransactStart)(nil),   // 0: protocol.TransactStart
	(*TransactCtrlCmd)(nil), // 1: protocol.TransactCtrlCmd
	(*TransactResult)(nil),  // 2: protocol.TransactResult
	(*TransactParam)(nil),   // 3: protocol.TransactParam
}
var file_transact_proto_depIdxs = []int32{
	3, // 0: protocol.TransactStart.MyTNP:type_name -> protocol.TransactParam
	3, // 1: protocol.TransactStart.ParenTNP:type_name -> protocol.TransactParam
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_transact_proto_init() }
func file_transact_proto_init() {
	if File_transact_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transact_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactStart); i {
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
		file_transact_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactCtrlCmd); i {
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
		file_transact_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactResult); i {
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
		file_transact_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactParam); i {
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
			RawDescriptor: file_transact_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transact_proto_goTypes,
		DependencyIndexes: file_transact_proto_depIdxs,
		MessageInfos:      file_transact_proto_msgTypes,
	}.Build()
	File_transact_proto = out.File
	file_transact_proto_rawDesc = nil
	file_transact_proto_goTypes = nil
	file_transact_proto_depIdxs = nil
}
