package http

import (
	_ "github.com/v4fly/v4ray-core/v0/common/protoext"
	http "github.com/v4fly/v4ray-core/v0/transport/internet/headers/http"
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
	Host          []string               `protobuf:"bytes,1,rep,name=host,proto3" json:"host,omitempty"`
	Path          string                 `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Method        string                 `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
	Header        []*http.Header         `protobuf:"bytes,4,rep,name=header,proto3" json:"header,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Config) Reset() {
	*x = Config{}
	mi := &file_transport_internet_http_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_transport_internet_http_config_proto_msgTypes[0]
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
	return file_transport_internet_http_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetHost() []string {
	if x != nil {
		return x.Host
	}
	return nil
}

func (x *Config) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Config) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Config) GetHeader() []*http.Header {
	if x != nil {
		return x.Header
	}
	return nil
}

var File_transport_internet_http_config_proto protoreflect.FileDescriptor

const file_transport_internet_http_config_proto_rawDesc = "" +
	"\n" +
	"$transport/internet/http/config.proto\x12\"v2ray.core.transport.internet.http\x1a common/protoext/extensions.proto\x1a,transport/internet/headers/http/config.proto\"\xb1\x01\n" +
	"\x06Config\x12\x12\n" +
	"\x04host\x18\x01 \x03(\tR\x04host\x12\x12\n" +
	"\x04path\x18\x02 \x01(\tR\x04path\x12\x16\n" +
	"\x06method\x18\x03 \x01(\tR\x06method\x12J\n" +
	"\x06header\x18\x04 \x03(\v22.v2ray.core.transport.internet.headers.http.HeaderR\x06header:\x1b\x82\xb5\x18\x17\n" +
	"\ttransport\x12\x02h2\x8a\xff)\x04httpB\x87\x01\n" +
	"&com.v2ray.core.transport.internet.httpP\x01Z6github.com/v4fly/v4ray-core/v0/transport/internet/http\xaa\x02\"V2Ray.Core.Transport.Internet.Httpb\x06proto3"

var (
	file_transport_internet_http_config_proto_rawDescOnce sync.Once
	file_transport_internet_http_config_proto_rawDescData []byte
)

func file_transport_internet_http_config_proto_rawDescGZIP() []byte {
	file_transport_internet_http_config_proto_rawDescOnce.Do(func() {
		file_transport_internet_http_config_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_transport_internet_http_config_proto_rawDesc), len(file_transport_internet_http_config_proto_rawDesc)))
	})
	return file_transport_internet_http_config_proto_rawDescData
}

var file_transport_internet_http_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_transport_internet_http_config_proto_goTypes = []any{
	(*Config)(nil),      // 0: v2ray.core.transport.internet.http.Config
	(*http.Header)(nil), // 1: v2ray.core.transport.internet.headers.http.Header
}
var file_transport_internet_http_config_proto_depIdxs = []int32{
	1, // 0: v2ray.core.transport.internet.http.Config.header:type_name -> v2ray.core.transport.internet.headers.http.Header
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_transport_internet_http_config_proto_init() }
func file_transport_internet_http_config_proto_init() {
	if File_transport_internet_http_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_transport_internet_http_config_proto_rawDesc), len(file_transport_internet_http_config_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transport_internet_http_config_proto_goTypes,
		DependencyIndexes: file_transport_internet_http_config_proto_depIdxs,
		MessageInfos:      file_transport_internet_http_config_proto_msgTypes,
	}.Build()
	File_transport_internet_http_config_proto = out.File
	file_transport_internet_http_config_proto_goTypes = nil
	file_transport_internet_http_config_proto_depIdxs = nil
}
