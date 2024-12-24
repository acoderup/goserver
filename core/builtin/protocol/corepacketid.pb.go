// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.19.4
// source: corepacketid.proto

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

type CoreBuiltinPacketID int32

const (
	CoreBuiltinPacketID_PACKET_COREBUILTIN_UNKNOW CoreBuiltinPacketID = 0
	CoreBuiltinPacketID_PACKET_SS_TX_START        CoreBuiltinPacketID = -1000
	CoreBuiltinPacketID_PACKET_SS_TX_CMD          CoreBuiltinPacketID = -1001
	CoreBuiltinPacketID_PACKET_SS_TX_RESULT       CoreBuiltinPacketID = -1002
	CoreBuiltinPacketID_PACKET_SS_SLICES          CoreBuiltinPacketID = -1003
	CoreBuiltinPacketID_PACKET_SS_AUTH            CoreBuiltinPacketID = -1004
	CoreBuiltinPacketID_PACKET_SS_KEEPALIVE       CoreBuiltinPacketID = -1005
	CoreBuiltinPacketID_PACKET_SS_AUTH_ACK        CoreBuiltinPacketID = -1006
	CoreBuiltinPacketID_PACKET_SS_RPC_REQ         CoreBuiltinPacketID = -1007
	CoreBuiltinPacketID_PACKET_SS_RPC_RESP        CoreBuiltinPacketID = -1008
)

// Enum value maps for CoreBuiltinPacketID.
var (
	CoreBuiltinPacketID_name = map[int32]string{
		0:     "PACKET_COREBUILTIN_UNKNOW",
		-1000: "PACKET_SS_TX_START",
		-1001: "PACKET_SS_TX_CMD",
		-1002: "PACKET_SS_TX_RESULT",
		-1003: "PACKET_SS_SLICES",
		-1004: "PACKET_SS_AUTH",
		-1005: "PACKET_SS_KEEPALIVE",
		-1006: "PACKET_SS_AUTH_ACK",
		-1007: "PACKET_SS_RPC_REQ",
		-1008: "PACKET_SS_RPC_RESP",
	}
	CoreBuiltinPacketID_value = map[string]int32{
		"PACKET_COREBUILTIN_UNKNOW": 0,
		"PACKET_SS_TX_START":        -1000,
		"PACKET_SS_TX_CMD":          -1001,
		"PACKET_SS_TX_RESULT":       -1002,
		"PACKET_SS_SLICES":          -1003,
		"PACKET_SS_AUTH":            -1004,
		"PACKET_SS_KEEPALIVE":       -1005,
		"PACKET_SS_AUTH_ACK":        -1006,
		"PACKET_SS_RPC_REQ":         -1007,
		"PACKET_SS_RPC_RESP":        -1008,
	}
)

func (x CoreBuiltinPacketID) Enum() *CoreBuiltinPacketID {
	p := new(CoreBuiltinPacketID)
	*p = x
	return p
}

func (x CoreBuiltinPacketID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CoreBuiltinPacketID) Descriptor() protoreflect.EnumDescriptor {
	return file_corepacketid_proto_enumTypes[0].Descriptor()
}

func (CoreBuiltinPacketID) Type() protoreflect.EnumType {
	return &file_corepacketid_proto_enumTypes[0]
}

func (x CoreBuiltinPacketID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CoreBuiltinPacketID.Descriptor instead.
func (CoreBuiltinPacketID) EnumDescriptor() ([]byte, []int) {
	return file_corepacketid_proto_rawDescGZIP(), []int{0}
}

var File_corepacketid_proto protoreflect.FileDescriptor

var file_corepacketid_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2a, 0xd6,
	0x02, 0x0a, 0x13, 0x43, 0x6f, 0x72, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x74, 0x69, 0x6e, 0x50, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x49, 0x44, 0x12, 0x1d, 0x0a, 0x19, 0x50, 0x41, 0x43, 0x4b, 0x45, 0x54,
	0x5f, 0x43, 0x4f, 0x52, 0x45, 0x42, 0x55, 0x49, 0x4c, 0x54, 0x49, 0x4e, 0x5f, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x10, 0x00, 0x12, 0x1f, 0x0a, 0x12, 0x50, 0x41, 0x43, 0x4b, 0x45, 0x54, 0x5f,
	0x53, 0x53, 0x5f, 0x54, 0x58, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x10, 0x98, 0xf8, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x1d, 0x0a, 0x10, 0x50, 0x41, 0x43, 0x4b, 0x45, 0x54,
	0x5f, 0x53, 0x53, 0x5f, 0x54, 0x58, 0x5f, 0x43, 0x4d, 0x44, 0x10, 0x97, 0xf8, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x20, 0x0a, 0x13, 0x50, 0x41, 0x43, 0x4b, 0x45, 0x54, 0x5f,
	0x53, 0x53, 0x5f, 0x54, 0x58, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x10, 0x96, 0xf8, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x1d, 0x0a, 0x10, 0x50, 0x41, 0x43, 0x4b, 0x45,
	0x54, 0x5f, 0x53, 0x53, 0x5f, 0x53, 0x4c, 0x49, 0x43, 0x45, 0x53, 0x10, 0x95, 0xf8, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x1b, 0x0a, 0x0e, 0x50, 0x41, 0x43, 0x4b, 0x45, 0x54,
	0x5f, 0x53, 0x53, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x10, 0x94, 0xf8, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0x01, 0x12, 0x20, 0x0a, 0x13, 0x50, 0x41, 0x43, 0x4b, 0x45, 0x54, 0x5f, 0x53, 0x53,
	0x5f, 0x4b, 0x45, 0x45, 0x50, 0x41, 0x4c, 0x49, 0x56, 0x45, 0x10, 0x93, 0xf8, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x1f, 0x0a, 0x12, 0x50, 0x41, 0x43, 0x4b, 0x45, 0x54, 0x5f,
	0x53, 0x53, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x41, 0x43, 0x4b, 0x10, 0x92, 0xf8, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x1e, 0x0a, 0x11, 0x50, 0x41, 0x43, 0x4b, 0x45, 0x54,
	0x5f, 0x53, 0x53, 0x5f, 0x52, 0x50, 0x43, 0x5f, 0x52, 0x45, 0x51, 0x10, 0x91, 0xf8, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x1f, 0x0a, 0x12, 0x50, 0x41, 0x43, 0x4b, 0x45, 0x54,
	0x5f, 0x53, 0x53, 0x5f, 0x52, 0x50, 0x43, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x10, 0x90, 0xf8, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_corepacketid_proto_rawDescOnce sync.Once
	file_corepacketid_proto_rawDescData = file_corepacketid_proto_rawDesc
)

func file_corepacketid_proto_rawDescGZIP() []byte {
	file_corepacketid_proto_rawDescOnce.Do(func() {
		file_corepacketid_proto_rawDescData = protoimpl.X.CompressGZIP(file_corepacketid_proto_rawDescData)
	})
	return file_corepacketid_proto_rawDescData
}

var file_corepacketid_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_corepacketid_proto_goTypes = []interface{}{
	(CoreBuiltinPacketID)(0), // 0: protocol.CoreBuiltinPacketID
}
var file_corepacketid_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_corepacketid_proto_init() }
func file_corepacketid_proto_init() {
	if File_corepacketid_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_corepacketid_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_corepacketid_proto_goTypes,
		DependencyIndexes: file_corepacketid_proto_depIdxs,
		EnumInfos:         file_corepacketid_proto_enumTypes,
	}.Build()
	File_corepacketid_proto = out.File
	file_corepacketid_proto_rawDesc = nil
	file_corepacketid_proto_goTypes = nil
	file_corepacketid_proto_depIdxs = nil
}
