package policy

import (
	_ "github.com/v4fly/v4ray-core/v0/common/protoext"
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

type Second struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         uint32                 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Second) Reset() {
	*x = Second{}
	mi := &file_app_policy_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Second) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Second) ProtoMessage() {}

func (x *Second) ProtoReflect() protoreflect.Message {
	mi := &file_app_policy_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Second.ProtoReflect.Descriptor instead.
func (*Second) Descriptor() ([]byte, []int) {
	return file_app_policy_config_proto_rawDescGZIP(), []int{0}
}

func (x *Second) GetValue() uint32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Policy struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Timeout       *Policy_Timeout        `protobuf:"bytes,1,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Stats         *Policy_Stats          `protobuf:"bytes,2,opt,name=stats,proto3" json:"stats,omitempty"`
	Buffer        *Policy_Buffer         `protobuf:"bytes,3,opt,name=buffer,proto3" json:"buffer,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Policy) Reset() {
	*x = Policy{}
	mi := &file_app_policy_config_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Policy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policy) ProtoMessage() {}

func (x *Policy) ProtoReflect() protoreflect.Message {
	mi := &file_app_policy_config_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policy.ProtoReflect.Descriptor instead.
func (*Policy) Descriptor() ([]byte, []int) {
	return file_app_policy_config_proto_rawDescGZIP(), []int{1}
}

func (x *Policy) GetTimeout() *Policy_Timeout {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *Policy) GetStats() *Policy_Stats {
	if x != nil {
		return x.Stats
	}
	return nil
}

func (x *Policy) GetBuffer() *Policy_Buffer {
	if x != nil {
		return x.Buffer
	}
	return nil
}

type SystemPolicy struct {
	state                 protoimpl.MessageState `protogen:"open.v1"`
	Stats                 *SystemPolicy_Stats    `protobuf:"bytes,1,opt,name=stats,proto3" json:"stats,omitempty"`
	OverrideAccessLogDest bool                   `protobuf:"varint,2,opt,name=override_access_log_dest,json=overrideAccessLogDest,proto3" json:"override_access_log_dest,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *SystemPolicy) Reset() {
	*x = SystemPolicy{}
	mi := &file_app_policy_config_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SystemPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemPolicy) ProtoMessage() {}

func (x *SystemPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_app_policy_config_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemPolicy.ProtoReflect.Descriptor instead.
func (*SystemPolicy) Descriptor() ([]byte, []int) {
	return file_app_policy_config_proto_rawDescGZIP(), []int{2}
}

func (x *SystemPolicy) GetStats() *SystemPolicy_Stats {
	if x != nil {
		return x.Stats
	}
	return nil
}

func (x *SystemPolicy) GetOverrideAccessLogDest() bool {
	if x != nil {
		return x.OverrideAccessLogDest
	}
	return false
}

type Config struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Level         map[uint32]*Policy     `protobuf:"bytes,1,rep,name=level,proto3" json:"level,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	System        *SystemPolicy          `protobuf:"bytes,2,opt,name=system,proto3" json:"system,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Config) Reset() {
	*x = Config{}
	mi := &file_app_policy_config_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_app_policy_config_proto_msgTypes[3]
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
	return file_app_policy_config_proto_rawDescGZIP(), []int{3}
}

func (x *Config) GetLevel() map[uint32]*Policy {
	if x != nil {
		return x.Level
	}
	return nil
}

func (x *Config) GetSystem() *SystemPolicy {
	if x != nil {
		return x.System
	}
	return nil
}

// Timeout is a message for timeout settings in various stages, in seconds.
type Policy_Timeout struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Handshake      *Second                `protobuf:"bytes,1,opt,name=handshake,proto3" json:"handshake,omitempty"`
	ConnectionIdle *Second                `protobuf:"bytes,2,opt,name=connection_idle,json=connectionIdle,proto3" json:"connection_idle,omitempty"`
	UplinkOnly     *Second                `protobuf:"bytes,3,opt,name=uplink_only,json=uplinkOnly,proto3" json:"uplink_only,omitempty"`
	DownlinkOnly   *Second                `protobuf:"bytes,4,opt,name=downlink_only,json=downlinkOnly,proto3" json:"downlink_only,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *Policy_Timeout) Reset() {
	*x = Policy_Timeout{}
	mi := &file_app_policy_config_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Policy_Timeout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policy_Timeout) ProtoMessage() {}

func (x *Policy_Timeout) ProtoReflect() protoreflect.Message {
	mi := &file_app_policy_config_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policy_Timeout.ProtoReflect.Descriptor instead.
func (*Policy_Timeout) Descriptor() ([]byte, []int) {
	return file_app_policy_config_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Policy_Timeout) GetHandshake() *Second {
	if x != nil {
		return x.Handshake
	}
	return nil
}

func (x *Policy_Timeout) GetConnectionIdle() *Second {
	if x != nil {
		return x.ConnectionIdle
	}
	return nil
}

func (x *Policy_Timeout) GetUplinkOnly() *Second {
	if x != nil {
		return x.UplinkOnly
	}
	return nil
}

func (x *Policy_Timeout) GetDownlinkOnly() *Second {
	if x != nil {
		return x.DownlinkOnly
	}
	return nil
}

type Policy_Stats struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserUplink    bool                   `protobuf:"varint,1,opt,name=user_uplink,json=userUplink,proto3" json:"user_uplink,omitempty"`
	UserDownlink  bool                   `protobuf:"varint,2,opt,name=user_downlink,json=userDownlink,proto3" json:"user_downlink,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Policy_Stats) Reset() {
	*x = Policy_Stats{}
	mi := &file_app_policy_config_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Policy_Stats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policy_Stats) ProtoMessage() {}

func (x *Policy_Stats) ProtoReflect() protoreflect.Message {
	mi := &file_app_policy_config_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policy_Stats.ProtoReflect.Descriptor instead.
func (*Policy_Stats) Descriptor() ([]byte, []int) {
	return file_app_policy_config_proto_rawDescGZIP(), []int{1, 1}
}

func (x *Policy_Stats) GetUserUplink() bool {
	if x != nil {
		return x.UserUplink
	}
	return false
}

func (x *Policy_Stats) GetUserDownlink() bool {
	if x != nil {
		return x.UserDownlink
	}
	return false
}

type Policy_Buffer struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Buffer size per connection, in bytes. -1 for unlimited buffer.
	Connection    int32 `protobuf:"varint,1,opt,name=connection,proto3" json:"connection,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Policy_Buffer) Reset() {
	*x = Policy_Buffer{}
	mi := &file_app_policy_config_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Policy_Buffer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policy_Buffer) ProtoMessage() {}

func (x *Policy_Buffer) ProtoReflect() protoreflect.Message {
	mi := &file_app_policy_config_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policy_Buffer.ProtoReflect.Descriptor instead.
func (*Policy_Buffer) Descriptor() ([]byte, []int) {
	return file_app_policy_config_proto_rawDescGZIP(), []int{1, 2}
}

func (x *Policy_Buffer) GetConnection() int32 {
	if x != nil {
		return x.Connection
	}
	return 0
}

type SystemPolicy_Stats struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	InboundUplink    bool                   `protobuf:"varint,1,opt,name=inbound_uplink,json=inboundUplink,proto3" json:"inbound_uplink,omitempty"`
	InboundDownlink  bool                   `protobuf:"varint,2,opt,name=inbound_downlink,json=inboundDownlink,proto3" json:"inbound_downlink,omitempty"`
	OutboundUplink   bool                   `protobuf:"varint,3,opt,name=outbound_uplink,json=outboundUplink,proto3" json:"outbound_uplink,omitempty"`
	OutboundDownlink bool                   `protobuf:"varint,4,opt,name=outbound_downlink,json=outboundDownlink,proto3" json:"outbound_downlink,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *SystemPolicy_Stats) Reset() {
	*x = SystemPolicy_Stats{}
	mi := &file_app_policy_config_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SystemPolicy_Stats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemPolicy_Stats) ProtoMessage() {}

func (x *SystemPolicy_Stats) ProtoReflect() protoreflect.Message {
	mi := &file_app_policy_config_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemPolicy_Stats.ProtoReflect.Descriptor instead.
func (*SystemPolicy_Stats) Descriptor() ([]byte, []int) {
	return file_app_policy_config_proto_rawDescGZIP(), []int{2, 0}
}

func (x *SystemPolicy_Stats) GetInboundUplink() bool {
	if x != nil {
		return x.InboundUplink
	}
	return false
}

func (x *SystemPolicy_Stats) GetInboundDownlink() bool {
	if x != nil {
		return x.InboundDownlink
	}
	return false
}

func (x *SystemPolicy_Stats) GetOutboundUplink() bool {
	if x != nil {
		return x.OutboundUplink
	}
	return false
}

func (x *SystemPolicy_Stats) GetOutboundDownlink() bool {
	if x != nil {
		return x.OutboundDownlink
	}
	return false
}

var File_app_policy_config_proto protoreflect.FileDescriptor

const file_app_policy_config_proto_rawDesc = "" +
	"\n" +
	"\x17app/policy/config.proto\x12\x15v2ray.core.app.policy\x1a common/protoext/extensions.proto\"\x1e\n" +
	"\x06Second\x12\x14\n" +
	"\x05value\x18\x01 \x01(\rR\x05value\"\xd0\x04\n" +
	"\x06Policy\x12?\n" +
	"\atimeout\x18\x01 \x01(\v2%.v2ray.core.app.policy.Policy.TimeoutR\atimeout\x129\n" +
	"\x05stats\x18\x02 \x01(\v2#.v2ray.core.app.policy.Policy.StatsR\x05stats\x12<\n" +
	"\x06buffer\x18\x03 \x01(\v2$.v2ray.core.app.policy.Policy.BufferR\x06buffer\x1a\x92\x02\n" +
	"\aTimeout\x12;\n" +
	"\thandshake\x18\x01 \x01(\v2\x1d.v2ray.core.app.policy.SecondR\thandshake\x12F\n" +
	"\x0fconnection_idle\x18\x02 \x01(\v2\x1d.v2ray.core.app.policy.SecondR\x0econnectionIdle\x12>\n" +
	"\vuplink_only\x18\x03 \x01(\v2\x1d.v2ray.core.app.policy.SecondR\n" +
	"uplinkOnly\x12B\n" +
	"\rdownlink_only\x18\x04 \x01(\v2\x1d.v2ray.core.app.policy.SecondR\fdownlinkOnly\x1aM\n" +
	"\x05Stats\x12\x1f\n" +
	"\vuser_uplink\x18\x01 \x01(\bR\n" +
	"userUplink\x12#\n" +
	"\ruser_downlink\x18\x02 \x01(\bR\fuserDownlink\x1a(\n" +
	"\x06Buffer\x12\x1e\n" +
	"\n" +
	"connection\x18\x01 \x01(\x05R\n" +
	"connection\"\xba\x02\n" +
	"\fSystemPolicy\x12?\n" +
	"\x05stats\x18\x01 \x01(\v2).v2ray.core.app.policy.SystemPolicy.StatsR\x05stats\x127\n" +
	"\x18override_access_log_dest\x18\x02 \x01(\bR\x15overrideAccessLogDest\x1a\xaf\x01\n" +
	"\x05Stats\x12%\n" +
	"\x0einbound_uplink\x18\x01 \x01(\bR\rinboundUplink\x12)\n" +
	"\x10inbound_downlink\x18\x02 \x01(\bR\x0finboundDownlink\x12'\n" +
	"\x0foutbound_uplink\x18\x03 \x01(\bR\x0eoutboundUplink\x12+\n" +
	"\x11outbound_downlink\x18\x04 \x01(\bR\x10outboundDownlink\"\xf5\x01\n" +
	"\x06Config\x12>\n" +
	"\x05level\x18\x01 \x03(\v2(.v2ray.core.app.policy.Config.LevelEntryR\x05level\x12;\n" +
	"\x06system\x18\x02 \x01(\v2#.v2ray.core.app.policy.SystemPolicyR\x06system\x1aW\n" +
	"\n" +
	"LevelEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\rR\x03key\x123\n" +
	"\x05value\x18\x02 \x01(\v2\x1d.v2ray.core.app.policy.PolicyR\x05value:\x028\x01:\x15\x82\xb5\x18\x11\n" +
	"\aservice\x12\x06policyB`\n" +
	"\x19com.v2ray.core.app.policyP\x01Z)github.com/v4fly/v4ray-core/v0/app/policy\xaa\x02\x15V2Ray.Core.App.Policyb\x06proto3"

var (
	file_app_policy_config_proto_rawDescOnce sync.Once
	file_app_policy_config_proto_rawDescData []byte
)

func file_app_policy_config_proto_rawDescGZIP() []byte {
	file_app_policy_config_proto_rawDescOnce.Do(func() {
		file_app_policy_config_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_app_policy_config_proto_rawDesc), len(file_app_policy_config_proto_rawDesc)))
	})
	return file_app_policy_config_proto_rawDescData
}

var file_app_policy_config_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_app_policy_config_proto_goTypes = []any{
	(*Second)(nil),             // 0: v2ray.core.app.policy.Second
	(*Policy)(nil),             // 1: v2ray.core.app.policy.Policy
	(*SystemPolicy)(nil),       // 2: v2ray.core.app.policy.SystemPolicy
	(*Config)(nil),             // 3: v2ray.core.app.policy.Config
	(*Policy_Timeout)(nil),     // 4: v2ray.core.app.policy.Policy.Timeout
	(*Policy_Stats)(nil),       // 5: v2ray.core.app.policy.Policy.Stats
	(*Policy_Buffer)(nil),      // 6: v2ray.core.app.policy.Policy.Buffer
	(*SystemPolicy_Stats)(nil), // 7: v2ray.core.app.policy.SystemPolicy.Stats
	nil,                        // 8: v2ray.core.app.policy.Config.LevelEntry
}
var file_app_policy_config_proto_depIdxs = []int32{
	4,  // 0: v2ray.core.app.policy.Policy.timeout:type_name -> v2ray.core.app.policy.Policy.Timeout
	5,  // 1: v2ray.core.app.policy.Policy.stats:type_name -> v2ray.core.app.policy.Policy.Stats
	6,  // 2: v2ray.core.app.policy.Policy.buffer:type_name -> v2ray.core.app.policy.Policy.Buffer
	7,  // 3: v2ray.core.app.policy.SystemPolicy.stats:type_name -> v2ray.core.app.policy.SystemPolicy.Stats
	8,  // 4: v2ray.core.app.policy.Config.level:type_name -> v2ray.core.app.policy.Config.LevelEntry
	2,  // 5: v2ray.core.app.policy.Config.system:type_name -> v2ray.core.app.policy.SystemPolicy
	0,  // 6: v2ray.core.app.policy.Policy.Timeout.handshake:type_name -> v2ray.core.app.policy.Second
	0,  // 7: v2ray.core.app.policy.Policy.Timeout.connection_idle:type_name -> v2ray.core.app.policy.Second
	0,  // 8: v2ray.core.app.policy.Policy.Timeout.uplink_only:type_name -> v2ray.core.app.policy.Second
	0,  // 9: v2ray.core.app.policy.Policy.Timeout.downlink_only:type_name -> v2ray.core.app.policy.Second
	1,  // 10: v2ray.core.app.policy.Config.LevelEntry.value:type_name -> v2ray.core.app.policy.Policy
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_app_policy_config_proto_init() }
func file_app_policy_config_proto_init() {
	if File_app_policy_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_app_policy_config_proto_rawDesc), len(file_app_policy_config_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_policy_config_proto_goTypes,
		DependencyIndexes: file_app_policy_config_proto_depIdxs,
		MessageInfos:      file_app_policy_config_proto_msgTypes,
	}.Build()
	File_app_policy_config_proto = out.File
	file_app_policy_config_proto_goTypes = nil
	file_app_policy_config_proto_depIdxs = nil
}
