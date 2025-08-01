package srtp

import (
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

type Config struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Version       uint32                 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Padding       bool                   `protobuf:"varint,2,opt,name=padding,proto3" json:"padding,omitempty"`
	Extension     bool                   `protobuf:"varint,3,opt,name=extension,proto3" json:"extension,omitempty"`
	CsrcCount     uint32                 `protobuf:"varint,4,opt,name=csrc_count,json=csrcCount,proto3" json:"csrc_count,omitempty"`
	Marker        bool                   `protobuf:"varint,5,opt,name=marker,proto3" json:"marker,omitempty"`
	PayloadType   uint32                 `protobuf:"varint,6,opt,name=payload_type,json=payloadType,proto3" json:"payload_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Config) Reset() {
	*x = Config{}
	mi := &file_transport_internet_headers_srtp_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_transport_internet_headers_srtp_config_proto_msgTypes[0]
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
	return file_transport_internet_headers_srtp_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Config) GetPadding() bool {
	if x != nil {
		return x.Padding
	}
	return false
}

func (x *Config) GetExtension() bool {
	if x != nil {
		return x.Extension
	}
	return false
}

func (x *Config) GetCsrcCount() uint32 {
	if x != nil {
		return x.CsrcCount
	}
	return 0
}

func (x *Config) GetMarker() bool {
	if x != nil {
		return x.Marker
	}
	return false
}

func (x *Config) GetPayloadType() uint32 {
	if x != nil {
		return x.PayloadType
	}
	return 0
}

var File_transport_internet_headers_srtp_config_proto protoreflect.FileDescriptor

const file_transport_internet_headers_srtp_config_proto_rawDesc = "" +
	"\n" +
	",transport/internet/headers/srtp/config.proto\x12*v2ray.core.transport.internet.headers.srtp\"\xb4\x01\n" +
	"\x06Config\x12\x18\n" +
	"\aversion\x18\x01 \x01(\rR\aversion\x12\x18\n" +
	"\apadding\x18\x02 \x01(\bR\apadding\x12\x1c\n" +
	"\textension\x18\x03 \x01(\bR\textension\x12\x1d\n" +
	"\n" +
	"csrc_count\x18\x04 \x01(\rR\tcsrcCount\x12\x16\n" +
	"\x06marker\x18\x05 \x01(\bR\x06marker\x12!\n" +
	"\fpayload_type\x18\x06 \x01(\rR\vpayloadTypeB\x9f\x01\n" +
	".com.v2ray.core.transport.internet.headers.srtpP\x01Z>github.com/v4fly/v4ray-core/v0/transport/internet/headers/srtp\xaa\x02*V2Ray.Core.Transport.Internet.Headers.Srtpb\x06proto3"

var (
	file_transport_internet_headers_srtp_config_proto_rawDescOnce sync.Once
	file_transport_internet_headers_srtp_config_proto_rawDescData []byte
)

func file_transport_internet_headers_srtp_config_proto_rawDescGZIP() []byte {
	file_transport_internet_headers_srtp_config_proto_rawDescOnce.Do(func() {
		file_transport_internet_headers_srtp_config_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_transport_internet_headers_srtp_config_proto_rawDesc), len(file_transport_internet_headers_srtp_config_proto_rawDesc)))
	})
	return file_transport_internet_headers_srtp_config_proto_rawDescData
}

var file_transport_internet_headers_srtp_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_transport_internet_headers_srtp_config_proto_goTypes = []any{
	(*Config)(nil), // 0: v2ray.core.transport.internet.headers.srtp.Config
}
var file_transport_internet_headers_srtp_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_transport_internet_headers_srtp_config_proto_init() }
func file_transport_internet_headers_srtp_config_proto_init() {
	if File_transport_internet_headers_srtp_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_transport_internet_headers_srtp_config_proto_rawDesc), len(file_transport_internet_headers_srtp_config_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transport_internet_headers_srtp_config_proto_goTypes,
		DependencyIndexes: file_transport_internet_headers_srtp_config_proto_depIdxs,
		MessageInfos:      file_transport_internet_headers_srtp_config_proto_msgTypes,
	}.Build()
	File_transport_internet_headers_srtp_config_proto = out.File
	file_transport_internet_headers_srtp_config_proto_goTypes = nil
	file_transport_internet_headers_srtp_config_proto_depIdxs = nil
}
