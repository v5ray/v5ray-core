package transport

import (
	internet "github.com/v4fly/v4ray-core/v0/transport/internet"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Global transport settings. This affects all type of connections that go
// through V2Ray. Deprecated. Use each settings in StreamConfig.
//
// Deprecated: Marked as deprecated in transport/config.proto.
type Config struct {
	state             protoimpl.MessageState      `protogen:"open.v1"`
	TransportSettings []*internet.TransportConfig `protobuf:"bytes,1,rep,name=transport_settings,json=transportSettings,proto3" json:"transport_settings,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *Config) Reset() {
	*x = Config{}
	mi := &file_transport_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_transport_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_transport_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetTransportSettings() []*internet.TransportConfig {
	if x != nil {
		return x.TransportSettings
	}
	return nil
}

var File_transport_config_proto protoreflect.FileDescriptor

const file_transport_config_proto_rawDesc = "" +
	"\n" +
	"\x16transport/config.proto\x12\x14v2ray.core.transport\x1a\x1ftransport/internet/config.proto\"k\n" +
	"\x06Config\x12]\n" +
	"\x12transport_settings\x18\x01 \x03(\v2..v2ray.core.transport.internet.TransportConfigR\x11transportSettings:\x02\x18\x01B]\n" +
	"\x18com.v2ray.core.transportP\x01Z(github.com/v4fly/v4ray-core/v0/transport\xaa\x02\x14V2Ray.Core.Transportb\x06proto3"

var (
	file_transport_config_proto_rawDescOnce sync.Once
	file_transport_config_proto_rawDescData []byte
)

func file_transport_config_proto_rawDescGZIP() []byte {
	file_transport_config_proto_rawDescOnce.Do(func() {
		file_transport_config_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_transport_config_proto_rawDesc), len(file_transport_config_proto_rawDesc)))
	})
	return file_transport_config_proto_rawDescData
}

var file_transport_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_transport_config_proto_goTypes = []any{
	(*Config)(nil),                   // 0: v2ray.core.transport.Config
	(*internet.TransportConfig)(nil), // 1: v2ray.core.transport.internet.TransportConfig
}
var file_transport_config_proto_depIdxs = []int32{
	1, // 0: v2ray.core.transport.Config.transport_settings:type_name -> v2ray.core.transport.internet.TransportConfig
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_transport_config_proto_init() }
func file_transport_config_proto_init() {
	if File_transport_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_transport_config_proto_rawDesc), len(file_transport_config_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transport_config_proto_goTypes,
		DependencyIndexes: file_transport_config_proto_depIdxs,
		MessageInfos:      file_transport_config_proto_msgTypes,
	}.Build()
	File_transport_config_proto = out.File
	file_transport_config_proto_goTypes = nil
	file_transport_config_proto_depIdxs = nil
}
