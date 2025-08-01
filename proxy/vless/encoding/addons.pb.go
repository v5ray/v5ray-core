package encoding

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

type Addons struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Flow          string                 `protobuf:"bytes,1,opt,name=Flow,proto3" json:"Flow,omitempty"`
	Seed          []byte                 `protobuf:"bytes,2,opt,name=Seed,proto3" json:"Seed,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Addons) Reset() {
	*x = Addons{}
	mi := &file_proxy_vless_encoding_addons_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Addons) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Addons) ProtoMessage() {}

func (x *Addons) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_vless_encoding_addons_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Addons.ProtoReflect.Descriptor instead.
func (*Addons) Descriptor() ([]byte, []int) {
	return file_proxy_vless_encoding_addons_proto_rawDescGZIP(), []int{0}
}

func (x *Addons) GetFlow() string {
	if x != nil {
		return x.Flow
	}
	return ""
}

func (x *Addons) GetSeed() []byte {
	if x != nil {
		return x.Seed
	}
	return nil
}

var File_proxy_vless_encoding_addons_proto protoreflect.FileDescriptor

const file_proxy_vless_encoding_addons_proto_rawDesc = "" +
	"\n" +
	"!proxy/vless/encoding/addons.proto\x12\x1fv2ray.core.proxy.vless.encoding\"0\n" +
	"\x06Addons\x12\x12\n" +
	"\x04Flow\x18\x01 \x01(\tR\x04Flow\x12\x12\n" +
	"\x04Seed\x18\x02 \x01(\fR\x04SeedB~\n" +
	"#com.v2ray.core.proxy.vless.encodingP\x01Z3github.com/v4fly/v4ray-core/v0/proxy/vless/encoding\xaa\x02\x1fV2Ray.Core.Proxy.Vless.Encodingb\x06proto3"

var (
	file_proxy_vless_encoding_addons_proto_rawDescOnce sync.Once
	file_proxy_vless_encoding_addons_proto_rawDescData []byte
)

func file_proxy_vless_encoding_addons_proto_rawDescGZIP() []byte {
	file_proxy_vless_encoding_addons_proto_rawDescOnce.Do(func() {
		file_proxy_vless_encoding_addons_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proxy_vless_encoding_addons_proto_rawDesc), len(file_proxy_vless_encoding_addons_proto_rawDesc)))
	})
	return file_proxy_vless_encoding_addons_proto_rawDescData
}

var file_proxy_vless_encoding_addons_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proxy_vless_encoding_addons_proto_goTypes = []any{
	(*Addons)(nil), // 0: v2ray.core.proxy.vless.encoding.Addons
}
var file_proxy_vless_encoding_addons_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proxy_vless_encoding_addons_proto_init() }
func file_proxy_vless_encoding_addons_proto_init() {
	if File_proxy_vless_encoding_addons_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proxy_vless_encoding_addons_proto_rawDesc), len(file_proxy_vless_encoding_addons_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proxy_vless_encoding_addons_proto_goTypes,
		DependencyIndexes: file_proxy_vless_encoding_addons_proto_depIdxs,
		MessageInfos:      file_proxy_vless_encoding_addons_proto_msgTypes,
	}.Build()
	File_proxy_vless_encoding_addons_proto = out.File
	file_proxy_vless_encoding_addons_proto_goTypes = nil
	file_proxy_vless_encoding_addons_proto_depIdxs = nil
}
