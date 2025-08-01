package command

import (
	v5 "github.com/v4fly/v4ray-core/v0"
	protocol "github.com/v4fly/v4ray-core/v0/common/protocol"
	_ "github.com/v4fly/v4ray-core/v0/common/protoext"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
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

type AddUserOperation struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *protocol.User         `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddUserOperation) Reset() {
	*x = AddUserOperation{}
	mi := &file_app_proxyman_command_command_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddUserOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserOperation) ProtoMessage() {}

func (x *AddUserOperation) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserOperation.ProtoReflect.Descriptor instead.
func (*AddUserOperation) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{0}
}

func (x *AddUserOperation) GetUser() *protocol.User {
	if x != nil {
		return x.User
	}
	return nil
}

type RemoveUserOperation struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveUserOperation) Reset() {
	*x = RemoveUserOperation{}
	mi := &file_app_proxyman_command_command_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveUserOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveUserOperation) ProtoMessage() {}

func (x *RemoveUserOperation) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveUserOperation.ProtoReflect.Descriptor instead.
func (*RemoveUserOperation) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{1}
}

func (x *RemoveUserOperation) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type AddInboundRequest struct {
	state         protoimpl.MessageState   `protogen:"open.v1"`
	Inbound       *v5.InboundHandlerConfig `protobuf:"bytes,1,opt,name=inbound,proto3" json:"inbound,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddInboundRequest) Reset() {
	*x = AddInboundRequest{}
	mi := &file_app_proxyman_command_command_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddInboundRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddInboundRequest) ProtoMessage() {}

func (x *AddInboundRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddInboundRequest.ProtoReflect.Descriptor instead.
func (*AddInboundRequest) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{2}
}

func (x *AddInboundRequest) GetInbound() *v5.InboundHandlerConfig {
	if x != nil {
		return x.Inbound
	}
	return nil
}

type AddInboundResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddInboundResponse) Reset() {
	*x = AddInboundResponse{}
	mi := &file_app_proxyman_command_command_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddInboundResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddInboundResponse) ProtoMessage() {}

func (x *AddInboundResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddInboundResponse.ProtoReflect.Descriptor instead.
func (*AddInboundResponse) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{3}
}

type RemoveInboundRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tag           string                 `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveInboundRequest) Reset() {
	*x = RemoveInboundRequest{}
	mi := &file_app_proxyman_command_command_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveInboundRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveInboundRequest) ProtoMessage() {}

func (x *RemoveInboundRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveInboundRequest.ProtoReflect.Descriptor instead.
func (*RemoveInboundRequest) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveInboundRequest) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

type RemoveInboundResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveInboundResponse) Reset() {
	*x = RemoveInboundResponse{}
	mi := &file_app_proxyman_command_command_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveInboundResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveInboundResponse) ProtoMessage() {}

func (x *RemoveInboundResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveInboundResponse.ProtoReflect.Descriptor instead.
func (*RemoveInboundResponse) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{5}
}

type AlterInboundRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tag           string                 `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	Operation     *anypb.Any             `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AlterInboundRequest) Reset() {
	*x = AlterInboundRequest{}
	mi := &file_app_proxyman_command_command_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlterInboundRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlterInboundRequest) ProtoMessage() {}

func (x *AlterInboundRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlterInboundRequest.ProtoReflect.Descriptor instead.
func (*AlterInboundRequest) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{6}
}

func (x *AlterInboundRequest) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *AlterInboundRequest) GetOperation() *anypb.Any {
	if x != nil {
		return x.Operation
	}
	return nil
}

type AlterInboundResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AlterInboundResponse) Reset() {
	*x = AlterInboundResponse{}
	mi := &file_app_proxyman_command_command_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlterInboundResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlterInboundResponse) ProtoMessage() {}

func (x *AlterInboundResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlterInboundResponse.ProtoReflect.Descriptor instead.
func (*AlterInboundResponse) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{7}
}

type AddOutboundRequest struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	Outbound      *v5.OutboundHandlerConfig `protobuf:"bytes,1,opt,name=outbound,proto3" json:"outbound,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddOutboundRequest) Reset() {
	*x = AddOutboundRequest{}
	mi := &file_app_proxyman_command_command_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddOutboundRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddOutboundRequest) ProtoMessage() {}

func (x *AddOutboundRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddOutboundRequest.ProtoReflect.Descriptor instead.
func (*AddOutboundRequest) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{8}
}

func (x *AddOutboundRequest) GetOutbound() *v5.OutboundHandlerConfig {
	if x != nil {
		return x.Outbound
	}
	return nil
}

type AddOutboundResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddOutboundResponse) Reset() {
	*x = AddOutboundResponse{}
	mi := &file_app_proxyman_command_command_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddOutboundResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddOutboundResponse) ProtoMessage() {}

func (x *AddOutboundResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddOutboundResponse.ProtoReflect.Descriptor instead.
func (*AddOutboundResponse) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{9}
}

type RemoveOutboundRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tag           string                 `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveOutboundRequest) Reset() {
	*x = RemoveOutboundRequest{}
	mi := &file_app_proxyman_command_command_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveOutboundRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveOutboundRequest) ProtoMessage() {}

func (x *RemoveOutboundRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveOutboundRequest.ProtoReflect.Descriptor instead.
func (*RemoveOutboundRequest) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{10}
}

func (x *RemoveOutboundRequest) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

type RemoveOutboundResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveOutboundResponse) Reset() {
	*x = RemoveOutboundResponse{}
	mi := &file_app_proxyman_command_command_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveOutboundResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveOutboundResponse) ProtoMessage() {}

func (x *RemoveOutboundResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveOutboundResponse.ProtoReflect.Descriptor instead.
func (*RemoveOutboundResponse) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{11}
}

type AlterOutboundRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tag           string                 `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	Operation     *anypb.Any             `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AlterOutboundRequest) Reset() {
	*x = AlterOutboundRequest{}
	mi := &file_app_proxyman_command_command_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlterOutboundRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlterOutboundRequest) ProtoMessage() {}

func (x *AlterOutboundRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlterOutboundRequest.ProtoReflect.Descriptor instead.
func (*AlterOutboundRequest) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{12}
}

func (x *AlterOutboundRequest) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *AlterOutboundRequest) GetOperation() *anypb.Any {
	if x != nil {
		return x.Operation
	}
	return nil
}

type AlterOutboundResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AlterOutboundResponse) Reset() {
	*x = AlterOutboundResponse{}
	mi := &file_app_proxyman_command_command_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlterOutboundResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlterOutboundResponse) ProtoMessage() {}

func (x *AlterOutboundResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlterOutboundResponse.ProtoReflect.Descriptor instead.
func (*AlterOutboundResponse) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{13}
}

type Config struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Config) Reset() {
	*x = Config{}
	mi := &file_app_proxyman_command_command_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[14]
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
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{14}
}

var File_app_proxyman_command_command_proto protoreflect.FileDescriptor

const file_app_proxyman_command_command_proto_rawDesc = "" +
	"\n" +
	"\"app/proxyman/command/command.proto\x12\x1fv2ray.core.app.proxyman.command\x1a\x1acommon/protocol/user.proto\x1a\x19google/protobuf/any.proto\x1a common/protoext/extensions.proto\x1a\fconfig.proto\"H\n" +
	"\x10AddUserOperation\x124\n" +
	"\x04user\x18\x01 \x01(\v2 .v2ray.core.common.protocol.UserR\x04user\"+\n" +
	"\x13RemoveUserOperation\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email\"O\n" +
	"\x11AddInboundRequest\x12:\n" +
	"\ainbound\x18\x01 \x01(\v2 .v2ray.core.InboundHandlerConfigR\ainbound\"\x14\n" +
	"\x12AddInboundResponse\"(\n" +
	"\x14RemoveInboundRequest\x12\x10\n" +
	"\x03tag\x18\x01 \x01(\tR\x03tag\"\x17\n" +
	"\x15RemoveInboundResponse\"[\n" +
	"\x13AlterInboundRequest\x12\x10\n" +
	"\x03tag\x18\x01 \x01(\tR\x03tag\x122\n" +
	"\toperation\x18\x02 \x01(\v2\x14.google.protobuf.AnyR\toperation\"\x16\n" +
	"\x14AlterInboundResponse\"S\n" +
	"\x12AddOutboundRequest\x12=\n" +
	"\boutbound\x18\x01 \x01(\v2!.v2ray.core.OutboundHandlerConfigR\boutbound\"\x15\n" +
	"\x13AddOutboundResponse\")\n" +
	"\x15RemoveOutboundRequest\x12\x10\n" +
	"\x03tag\x18\x01 \x01(\tR\x03tag\"\x18\n" +
	"\x16RemoveOutboundResponse\"\\\n" +
	"\x14AlterOutboundRequest\x12\x10\n" +
	"\x03tag\x18\x01 \x01(\tR\x03tag\x122\n" +
	"\toperation\x18\x02 \x01(\v2\x14.google.protobuf.AnyR\toperation\"\x17\n" +
	"\x15AlterOutboundResponse\"%\n" +
	"\x06Config:\x1b\x82\xb5\x18\x17\n" +
	"\vgrpcservice\x12\bproxyman2\x90\x06\n" +
	"\x0eHandlerService\x12w\n" +
	"\n" +
	"AddInbound\x122.v2ray.core.app.proxyman.command.AddInboundRequest\x1a3.v2ray.core.app.proxyman.command.AddInboundResponse\"\x00\x12\x80\x01\n" +
	"\rRemoveInbound\x125.v2ray.core.app.proxyman.command.RemoveInboundRequest\x1a6.v2ray.core.app.proxyman.command.RemoveInboundResponse\"\x00\x12}\n" +
	"\fAlterInbound\x124.v2ray.core.app.proxyman.command.AlterInboundRequest\x1a5.v2ray.core.app.proxyman.command.AlterInboundResponse\"\x00\x12z\n" +
	"\vAddOutbound\x123.v2ray.core.app.proxyman.command.AddOutboundRequest\x1a4.v2ray.core.app.proxyman.command.AddOutboundResponse\"\x00\x12\x83\x01\n" +
	"\x0eRemoveOutbound\x126.v2ray.core.app.proxyman.command.RemoveOutboundRequest\x1a7.v2ray.core.app.proxyman.command.RemoveOutboundResponse\"\x00\x12\x80\x01\n" +
	"\rAlterOutbound\x125.v2ray.core.app.proxyman.command.AlterOutboundRequest\x1a6.v2ray.core.app.proxyman.command.AlterOutboundResponse\"\x00B~\n" +
	"#com.v2ray.core.app.proxyman.commandP\x01Z3github.com/v4fly/v4ray-core/v0/app/proxyman/command\xaa\x02\x1fV2Ray.Core.App.Proxyman.Commandb\x06proto3"

var (
	file_app_proxyman_command_command_proto_rawDescOnce sync.Once
	file_app_proxyman_command_command_proto_rawDescData []byte
)

func file_app_proxyman_command_command_proto_rawDescGZIP() []byte {
	file_app_proxyman_command_command_proto_rawDescOnce.Do(func() {
		file_app_proxyman_command_command_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_app_proxyman_command_command_proto_rawDesc), len(file_app_proxyman_command_command_proto_rawDesc)))
	})
	return file_app_proxyman_command_command_proto_rawDescData
}

var file_app_proxyman_command_command_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_app_proxyman_command_command_proto_goTypes = []any{
	(*AddUserOperation)(nil),         // 0: v2ray.core.app.proxyman.command.AddUserOperation
	(*RemoveUserOperation)(nil),      // 1: v2ray.core.app.proxyman.command.RemoveUserOperation
	(*AddInboundRequest)(nil),        // 2: v2ray.core.app.proxyman.command.AddInboundRequest
	(*AddInboundResponse)(nil),       // 3: v2ray.core.app.proxyman.command.AddInboundResponse
	(*RemoveInboundRequest)(nil),     // 4: v2ray.core.app.proxyman.command.RemoveInboundRequest
	(*RemoveInboundResponse)(nil),    // 5: v2ray.core.app.proxyman.command.RemoveInboundResponse
	(*AlterInboundRequest)(nil),      // 6: v2ray.core.app.proxyman.command.AlterInboundRequest
	(*AlterInboundResponse)(nil),     // 7: v2ray.core.app.proxyman.command.AlterInboundResponse
	(*AddOutboundRequest)(nil),       // 8: v2ray.core.app.proxyman.command.AddOutboundRequest
	(*AddOutboundResponse)(nil),      // 9: v2ray.core.app.proxyman.command.AddOutboundResponse
	(*RemoveOutboundRequest)(nil),    // 10: v2ray.core.app.proxyman.command.RemoveOutboundRequest
	(*RemoveOutboundResponse)(nil),   // 11: v2ray.core.app.proxyman.command.RemoveOutboundResponse
	(*AlterOutboundRequest)(nil),     // 12: v2ray.core.app.proxyman.command.AlterOutboundRequest
	(*AlterOutboundResponse)(nil),    // 13: v2ray.core.app.proxyman.command.AlterOutboundResponse
	(*Config)(nil),                   // 14: v2ray.core.app.proxyman.command.Config
	(*protocol.User)(nil),            // 15: v2ray.core.common.protocol.User
	(*v5.InboundHandlerConfig)(nil),  // 16: v2ray.core.InboundHandlerConfig
	(*anypb.Any)(nil),                // 17: google.protobuf.Any
	(*v5.OutboundHandlerConfig)(nil), // 18: v2ray.core.OutboundHandlerConfig
}
var file_app_proxyman_command_command_proto_depIdxs = []int32{
	15, // 0: v2ray.core.app.proxyman.command.AddUserOperation.user:type_name -> v2ray.core.common.protocol.User
	16, // 1: v2ray.core.app.proxyman.command.AddInboundRequest.inbound:type_name -> v2ray.core.InboundHandlerConfig
	17, // 2: v2ray.core.app.proxyman.command.AlterInboundRequest.operation:type_name -> google.protobuf.Any
	18, // 3: v2ray.core.app.proxyman.command.AddOutboundRequest.outbound:type_name -> v2ray.core.OutboundHandlerConfig
	17, // 4: v2ray.core.app.proxyman.command.AlterOutboundRequest.operation:type_name -> google.protobuf.Any
	2,  // 5: v2ray.core.app.proxyman.command.HandlerService.AddInbound:input_type -> v2ray.core.app.proxyman.command.AddInboundRequest
	4,  // 6: v2ray.core.app.proxyman.command.HandlerService.RemoveInbound:input_type -> v2ray.core.app.proxyman.command.RemoveInboundRequest
	6,  // 7: v2ray.core.app.proxyman.command.HandlerService.AlterInbound:input_type -> v2ray.core.app.proxyman.command.AlterInboundRequest
	8,  // 8: v2ray.core.app.proxyman.command.HandlerService.AddOutbound:input_type -> v2ray.core.app.proxyman.command.AddOutboundRequest
	10, // 9: v2ray.core.app.proxyman.command.HandlerService.RemoveOutbound:input_type -> v2ray.core.app.proxyman.command.RemoveOutboundRequest
	12, // 10: v2ray.core.app.proxyman.command.HandlerService.AlterOutbound:input_type -> v2ray.core.app.proxyman.command.AlterOutboundRequest
	3,  // 11: v2ray.core.app.proxyman.command.HandlerService.AddInbound:output_type -> v2ray.core.app.proxyman.command.AddInboundResponse
	5,  // 12: v2ray.core.app.proxyman.command.HandlerService.RemoveInbound:output_type -> v2ray.core.app.proxyman.command.RemoveInboundResponse
	7,  // 13: v2ray.core.app.proxyman.command.HandlerService.AlterInbound:output_type -> v2ray.core.app.proxyman.command.AlterInboundResponse
	9,  // 14: v2ray.core.app.proxyman.command.HandlerService.AddOutbound:output_type -> v2ray.core.app.proxyman.command.AddOutboundResponse
	11, // 15: v2ray.core.app.proxyman.command.HandlerService.RemoveOutbound:output_type -> v2ray.core.app.proxyman.command.RemoveOutboundResponse
	13, // 16: v2ray.core.app.proxyman.command.HandlerService.AlterOutbound:output_type -> v2ray.core.app.proxyman.command.AlterOutboundResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_app_proxyman_command_command_proto_init() }
func file_app_proxyman_command_command_proto_init() {
	if File_app_proxyman_command_command_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_app_proxyman_command_command_proto_rawDesc), len(file_app_proxyman_command_command_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_proxyman_command_command_proto_goTypes,
		DependencyIndexes: file_app_proxyman_command_command_proto_depIdxs,
		MessageInfos:      file_app_proxyman_command_command_proto_msgTypes,
	}.Build()
	File_app_proxyman_command_command_proto = out.File
	file_app_proxyman_command_command_proto_goTypes = nil
	file_app_proxyman_command_command_proto_depIdxs = nil
}
