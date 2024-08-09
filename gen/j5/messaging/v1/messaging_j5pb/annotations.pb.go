// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: j5/messaging/v1/annotations.proto

package messaging_j5pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ServiceConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TopicName *string `protobuf:"bytes,1,opt,name=topic_name,json=topicName,proto3,oneof" json:"topic_name,omitempty"`
	// Types that are assignable to Role:
	//
	//	*ServiceConfig_Publish_
	//	*ServiceConfig_Request_
	//	*ServiceConfig_Reply_
	Role isServiceConfig_Role `protobuf_oneof:"role"`
}

func (x *ServiceConfig) Reset() {
	*x = ServiceConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_j5_messaging_v1_annotations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceConfig) ProtoMessage() {}

func (x *ServiceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_j5_messaging_v1_annotations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceConfig.ProtoReflect.Descriptor instead.
func (*ServiceConfig) Descriptor() ([]byte, []int) {
	return file_j5_messaging_v1_annotations_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceConfig) GetTopicName() string {
	if x != nil && x.TopicName != nil {
		return *x.TopicName
	}
	return ""
}

func (m *ServiceConfig) GetRole() isServiceConfig_Role {
	if m != nil {
		return m.Role
	}
	return nil
}

func (x *ServiceConfig) GetPublish() *ServiceConfig_Publish {
	if x, ok := x.GetRole().(*ServiceConfig_Publish_); ok {
		return x.Publish
	}
	return nil
}

func (x *ServiceConfig) GetRequest() *ServiceConfig_Request {
	if x, ok := x.GetRole().(*ServiceConfig_Request_); ok {
		return x.Request
	}
	return nil
}

func (x *ServiceConfig) GetReply() *ServiceConfig_Reply {
	if x, ok := x.GetRole().(*ServiceConfig_Reply_); ok {
		return x.Reply
	}
	return nil
}

type isServiceConfig_Role interface {
	isServiceConfig_Role()
}

type ServiceConfig_Publish_ struct {
	Publish *ServiceConfig_Publish `protobuf:"bytes,10,opt,name=publish,proto3,oneof"`
}

type ServiceConfig_Request_ struct {
	Request *ServiceConfig_Request `protobuf:"bytes,11,opt,name=request,proto3,oneof"`
}

type ServiceConfig_Reply_ struct {
	Reply *ServiceConfig_Reply `protobuf:"bytes,12,opt,name=reply,proto3,oneof"`
}

func (*ServiceConfig_Publish_) isServiceConfig_Role() {}

func (*ServiceConfig_Request_) isServiceConfig_Role() {}

func (*ServiceConfig_Reply_) isServiceConfig_Role() {}

type PublishMethod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PublishMethod) Reset() {
	*x = PublishMethod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_j5_messaging_v1_annotations_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishMethod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishMethod) ProtoMessage() {}

func (x *PublishMethod) ProtoReflect() protoreflect.Message {
	mi := &file_j5_messaging_v1_annotations_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishMethod.ProtoReflect.Descriptor instead.
func (*PublishMethod) Descriptor() ([]byte, []int) {
	return file_j5_messaging_v1_annotations_proto_rawDescGZIP(), []int{1}
}

type RequestMethod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RequestMethod) Reset() {
	*x = RequestMethod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_j5_messaging_v1_annotations_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestMethod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestMethod) ProtoMessage() {}

func (x *RequestMethod) ProtoReflect() protoreflect.Message {
	mi := &file_j5_messaging_v1_annotations_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestMethod.ProtoReflect.Descriptor instead.
func (*RequestMethod) Descriptor() ([]byte, []int) {
	return file_j5_messaging_v1_annotations_proto_rawDescGZIP(), []int{2}
}

type ReplyMethod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReplyMethod) Reset() {
	*x = ReplyMethod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_j5_messaging_v1_annotations_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyMethod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyMethod) ProtoMessage() {}

func (x *ReplyMethod) ProtoReflect() protoreflect.Message {
	mi := &file_j5_messaging_v1_annotations_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyMethod.ProtoReflect.Descriptor instead.
func (*ReplyMethod) Descriptor() ([]byte, []int) {
	return file_j5_messaging_v1_annotations_proto_rawDescGZIP(), []int{3}
}

// Deprecated, use ServiceConfig instead
type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//
	//	*Config_Broadcast
	//	*Config_Unicast
	//	*Config_Request
	//	*Config_Reply
	Type isConfig_Type `protobuf_oneof:"type"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_j5_messaging_v1_annotations_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_j5_messaging_v1_annotations_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_j5_messaging_v1_annotations_proto_rawDescGZIP(), []int{4}
}

func (m *Config) GetType() isConfig_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *Config) GetBroadcast() *BroadcastConfig {
	if x, ok := x.GetType().(*Config_Broadcast); ok {
		return x.Broadcast
	}
	return nil
}

func (x *Config) GetUnicast() *UnicastConfig {
	if x, ok := x.GetType().(*Config_Unicast); ok {
		return x.Unicast
	}
	return nil
}

func (x *Config) GetRequest() *RequestConfig {
	if x, ok := x.GetType().(*Config_Request); ok {
		return x.Request
	}
	return nil
}

func (x *Config) GetReply() *ReplyConfig {
	if x, ok := x.GetType().(*Config_Reply); ok {
		return x.Reply
	}
	return nil
}

type isConfig_Type interface {
	isConfig_Type()
}

type Config_Broadcast struct {
	Broadcast *BroadcastConfig `protobuf:"bytes,1,opt,name=broadcast,proto3,oneof"`
}

type Config_Unicast struct {
	Unicast *UnicastConfig `protobuf:"bytes,2,opt,name=unicast,proto3,oneof"`
}

type Config_Request struct {
	Request *RequestConfig `protobuf:"bytes,3,opt,name=request,proto3,oneof"`
}

type Config_Reply struct {
	Reply *ReplyConfig `protobuf:"bytes,4,opt,name=reply,proto3,oneof"`
}

func (*Config_Broadcast) isConfig_Type() {}

func (*Config_Unicast) isConfig_Type() {}

func (*Config_Request) isConfig_Type() {}

func (*Config_Reply) isConfig_Type() {}

type BroadcastConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *BroadcastConfig) Reset() {
	*x = BroadcastConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_j5_messaging_v1_annotations_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastConfig) ProtoMessage() {}

func (x *BroadcastConfig) ProtoReflect() protoreflect.Message {
	mi := &file_j5_messaging_v1_annotations_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastConfig.ProtoReflect.Descriptor instead.
func (*BroadcastConfig) Descriptor() ([]byte, []int) {
	return file_j5_messaging_v1_annotations_proto_rawDescGZIP(), []int{5}
}

func (x *BroadcastConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UnicastConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UnicastConfig) Reset() {
	*x = UnicastConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_j5_messaging_v1_annotations_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnicastConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnicastConfig) ProtoMessage() {}

func (x *UnicastConfig) ProtoReflect() protoreflect.Message {
	mi := &file_j5_messaging_v1_annotations_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnicastConfig.ProtoReflect.Descriptor instead.
func (*UnicastConfig) Descriptor() ([]byte, []int) {
	return file_j5_messaging_v1_annotations_proto_rawDescGZIP(), []int{6}
}

func (x *UnicastConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RequestConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of the queue group (i.e. prefix), should match ReplyConfig.name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *RequestConfig) Reset() {
	*x = RequestConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_j5_messaging_v1_annotations_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestConfig) ProtoMessage() {}

func (x *RequestConfig) ProtoReflect() protoreflect.Message {
	mi := &file_j5_messaging_v1_annotations_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestConfig.ProtoReflect.Descriptor instead.
func (*RequestConfig) Descriptor() ([]byte, []int) {
	return file_j5_messaging_v1_annotations_proto_rawDescGZIP(), []int{7}
}

func (x *RequestConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ReplyConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of the queue group (i.e. prefix), should match RequestConfig.name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ReplyConfig) Reset() {
	*x = ReplyConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_j5_messaging_v1_annotations_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyConfig) ProtoMessage() {}

func (x *ReplyConfig) ProtoReflect() protoreflect.Message {
	mi := &file_j5_messaging_v1_annotations_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyConfig.ProtoReflect.Descriptor instead.
func (*ReplyConfig) Descriptor() ([]byte, []int) {
	return file_j5_messaging_v1_annotations_proto_rawDescGZIP(), []int{8}
}

func (x *ReplyConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type FieldConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Designates the field as the primary message ID, which will be mapped to the
	// infra message_id field.
	// When no message ID field is set, the ID must be explicitly set at runtime.
	MessageId bool `protobuf:"varint,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
}

func (x *FieldConfig) Reset() {
	*x = FieldConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_j5_messaging_v1_annotations_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldConfig) ProtoMessage() {}

func (x *FieldConfig) ProtoReflect() protoreflect.Message {
	mi := &file_j5_messaging_v1_annotations_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldConfig.ProtoReflect.Descriptor instead.
func (*FieldConfig) Descriptor() ([]byte, []int) {
	return file_j5_messaging_v1_annotations_proto_rawDescGZIP(), []int{9}
}

func (x *FieldConfig) GetMessageId() bool {
	if x != nil {
		return x.MessageId
	}
	return false
}

type ServiceConfig_Publish struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ServiceConfig_Publish) Reset() {
	*x = ServiceConfig_Publish{}
	if protoimpl.UnsafeEnabled {
		mi := &file_j5_messaging_v1_annotations_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceConfig_Publish) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceConfig_Publish) ProtoMessage() {}

func (x *ServiceConfig_Publish) ProtoReflect() protoreflect.Message {
	mi := &file_j5_messaging_v1_annotations_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceConfig_Publish.ProtoReflect.Descriptor instead.
func (*ServiceConfig_Publish) Descriptor() ([]byte, []int) {
	return file_j5_messaging_v1_annotations_proto_rawDescGZIP(), []int{0, 0}
}

type ServiceConfig_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ServiceConfig_Request) Reset() {
	*x = ServiceConfig_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_j5_messaging_v1_annotations_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceConfig_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceConfig_Request) ProtoMessage() {}

func (x *ServiceConfig_Request) ProtoReflect() protoreflect.Message {
	mi := &file_j5_messaging_v1_annotations_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceConfig_Request.ProtoReflect.Descriptor instead.
func (*ServiceConfig_Request) Descriptor() ([]byte, []int) {
	return file_j5_messaging_v1_annotations_proto_rawDescGZIP(), []int{0, 1}
}

type ServiceConfig_Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ServiceConfig_Reply) Reset() {
	*x = ServiceConfig_Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_j5_messaging_v1_annotations_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceConfig_Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceConfig_Reply) ProtoMessage() {}

func (x *ServiceConfig_Reply) ProtoReflect() protoreflect.Message {
	mi := &file_j5_messaging_v1_annotations_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceConfig_Reply.ProtoReflect.Descriptor instead.
func (*ServiceConfig_Reply) Descriptor() ([]byte, []int) {
	return file_j5_messaging_v1_annotations_proto_rawDescGZIP(), []int{0, 2}
}

var file_j5_messaging_v1_annotations_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*Config)(nil),
		Field:         93563434,
		Name:          "j5.messaging.v1.config",
		Tag:           "bytes,93563434,opt,name=config",
		Filename:      "j5/messaging/v1/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*ServiceConfig)(nil),
		Field:         93563435,
		Name:          "j5.messaging.v1.service",
		Tag:           "bytes,93563435,opt,name=service",
		Filename:      "j5/messaging/v1/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*FieldConfig)(nil),
		Field:         93563434,
		Name:          "j5.messaging.v1.field",
		Tag:           "bytes,93563434,opt,name=field",
		Filename:      "j5/messaging/v1/annotations.proto",
	},
}

// Extension fields to descriptorpb.ServiceOptions.
var (
	// optional j5.messaging.v1.Config config = 93563434;
	E_Config = &file_j5_messaging_v1_annotations_proto_extTypes[0]
	// optional j5.messaging.v1.ServiceConfig service = 93563435;
	E_Service = &file_j5_messaging_v1_annotations_proto_extTypes[1]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional j5.messaging.v1.FieldConfig field = 93563434;
	E_Field = &file_j5_messaging_v1_annotations_proto_extTypes[2]
)

var File_j5_messaging_v1_annotations_proto protoreflect.FileDescriptor

var file_j5_messaging_v1_annotations_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6a, 0x35, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x6a, 0x35, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x02, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x22, 0x0a, 0x0a, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x09,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x42, 0x0a, 0x07,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x6a, 0x35, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x48, 0x00, 0x52, 0x07, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x12, 0x42, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x6a, 0x35, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x07, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6a, 0x35, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x48, 0x00, 0x52, 0x05, 0x72, 0x65, 0x70,
	0x6c, 0x79, 0x1a, 0x09, 0x0a, 0x07, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x1a, 0x09, 0x0a,
	0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x42, 0x06, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0x0f, 0x0a, 0x0d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0x0d, 0x0a, 0x0b, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0x80, 0x02, 0x0a, 0x06, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x40, 0x0a, 0x09, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6a, 0x35, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63,
	0x61, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x09, 0x62, 0x72, 0x6f,
	0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x07, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x73,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6a, 0x35, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x69, 0x63, 0x61, 0x73,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x07, 0x75, 0x6e, 0x69, 0x63, 0x61,
	0x73, 0x74, 0x12, 0x3a, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6a, 0x35, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34,
	0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x6a, 0x35, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x05, 0x72,
	0x65, 0x70, 0x6c, 0x79, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x25, 0x0a, 0x0f,
	0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x23, 0x0a, 0x0d, 0x55, 0x6e, 0x69, 0x63, 0x61, 0x73, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x23, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x21, 0x0a,
	0x0b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x2c, 0x0a, 0x0b, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x3a, 0x53,
	0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xaa, 0xd4, 0xce, 0x2c, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6a, 0x35, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x3a, 0x5c, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1f,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xab, 0xd4, 0xce, 0x2c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6a, 0x35, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x3a, 0x54, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xaa, 0xd4, 0xce, 0x2c, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6a, 0x35, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x6e, 0x74, 0x6f, 0x70, 0x73, 0x2f, 0x6a, 0x35,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x6a, 0x35, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e,
	0x67, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x6a,
	0x35, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_j5_messaging_v1_annotations_proto_rawDescOnce sync.Once
	file_j5_messaging_v1_annotations_proto_rawDescData = file_j5_messaging_v1_annotations_proto_rawDesc
)

func file_j5_messaging_v1_annotations_proto_rawDescGZIP() []byte {
	file_j5_messaging_v1_annotations_proto_rawDescOnce.Do(func() {
		file_j5_messaging_v1_annotations_proto_rawDescData = protoimpl.X.CompressGZIP(file_j5_messaging_v1_annotations_proto_rawDescData)
	})
	return file_j5_messaging_v1_annotations_proto_rawDescData
}

var file_j5_messaging_v1_annotations_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_j5_messaging_v1_annotations_proto_goTypes = []any{
	(*ServiceConfig)(nil),               // 0: j5.messaging.v1.ServiceConfig
	(*PublishMethod)(nil),               // 1: j5.messaging.v1.PublishMethod
	(*RequestMethod)(nil),               // 2: j5.messaging.v1.RequestMethod
	(*ReplyMethod)(nil),                 // 3: j5.messaging.v1.ReplyMethod
	(*Config)(nil),                      // 4: j5.messaging.v1.Config
	(*BroadcastConfig)(nil),             // 5: j5.messaging.v1.BroadcastConfig
	(*UnicastConfig)(nil),               // 6: j5.messaging.v1.UnicastConfig
	(*RequestConfig)(nil),               // 7: j5.messaging.v1.RequestConfig
	(*ReplyConfig)(nil),                 // 8: j5.messaging.v1.ReplyConfig
	(*FieldConfig)(nil),                 // 9: j5.messaging.v1.FieldConfig
	(*ServiceConfig_Publish)(nil),       // 10: j5.messaging.v1.ServiceConfig.Publish
	(*ServiceConfig_Request)(nil),       // 11: j5.messaging.v1.ServiceConfig.Request
	(*ServiceConfig_Reply)(nil),         // 12: j5.messaging.v1.ServiceConfig.Reply
	(*descriptorpb.ServiceOptions)(nil), // 13: google.protobuf.ServiceOptions
	(*descriptorpb.FieldOptions)(nil),   // 14: google.protobuf.FieldOptions
}
var file_j5_messaging_v1_annotations_proto_depIdxs = []int32{
	10, // 0: j5.messaging.v1.ServiceConfig.publish:type_name -> j5.messaging.v1.ServiceConfig.Publish
	11, // 1: j5.messaging.v1.ServiceConfig.request:type_name -> j5.messaging.v1.ServiceConfig.Request
	12, // 2: j5.messaging.v1.ServiceConfig.reply:type_name -> j5.messaging.v1.ServiceConfig.Reply
	5,  // 3: j5.messaging.v1.Config.broadcast:type_name -> j5.messaging.v1.BroadcastConfig
	6,  // 4: j5.messaging.v1.Config.unicast:type_name -> j5.messaging.v1.UnicastConfig
	7,  // 5: j5.messaging.v1.Config.request:type_name -> j5.messaging.v1.RequestConfig
	8,  // 6: j5.messaging.v1.Config.reply:type_name -> j5.messaging.v1.ReplyConfig
	13, // 7: j5.messaging.v1.config:extendee -> google.protobuf.ServiceOptions
	13, // 8: j5.messaging.v1.service:extendee -> google.protobuf.ServiceOptions
	14, // 9: j5.messaging.v1.field:extendee -> google.protobuf.FieldOptions
	4,  // 10: j5.messaging.v1.config:type_name -> j5.messaging.v1.Config
	0,  // 11: j5.messaging.v1.service:type_name -> j5.messaging.v1.ServiceConfig
	9,  // 12: j5.messaging.v1.field:type_name -> j5.messaging.v1.FieldConfig
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	10, // [10:13] is the sub-list for extension type_name
	7,  // [7:10] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_j5_messaging_v1_annotations_proto_init() }
func file_j5_messaging_v1_annotations_proto_init() {
	if File_j5_messaging_v1_annotations_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_j5_messaging_v1_annotations_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ServiceConfig); i {
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
		file_j5_messaging_v1_annotations_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PublishMethod); i {
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
		file_j5_messaging_v1_annotations_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*RequestMethod); i {
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
		file_j5_messaging_v1_annotations_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ReplyMethod); i {
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
		file_j5_messaging_v1_annotations_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Config); i {
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
		file_j5_messaging_v1_annotations_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*BroadcastConfig); i {
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
		file_j5_messaging_v1_annotations_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*UnicastConfig); i {
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
		file_j5_messaging_v1_annotations_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*RequestConfig); i {
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
		file_j5_messaging_v1_annotations_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*ReplyConfig); i {
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
		file_j5_messaging_v1_annotations_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*FieldConfig); i {
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
		file_j5_messaging_v1_annotations_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*ServiceConfig_Publish); i {
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
		file_j5_messaging_v1_annotations_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*ServiceConfig_Request); i {
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
		file_j5_messaging_v1_annotations_proto_msgTypes[12].Exporter = func(v any, i int) any {
			switch v := v.(*ServiceConfig_Reply); i {
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
	file_j5_messaging_v1_annotations_proto_msgTypes[0].OneofWrappers = []any{
		(*ServiceConfig_Publish_)(nil),
		(*ServiceConfig_Request_)(nil),
		(*ServiceConfig_Reply_)(nil),
	}
	file_j5_messaging_v1_annotations_proto_msgTypes[4].OneofWrappers = []any{
		(*Config_Broadcast)(nil),
		(*Config_Unicast)(nil),
		(*Config_Request)(nil),
		(*Config_Reply)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_j5_messaging_v1_annotations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 3,
			NumServices:   0,
		},
		GoTypes:           file_j5_messaging_v1_annotations_proto_goTypes,
		DependencyIndexes: file_j5_messaging_v1_annotations_proto_depIdxs,
		MessageInfos:      file_j5_messaging_v1_annotations_proto_msgTypes,
		ExtensionInfos:    file_j5_messaging_v1_annotations_proto_extTypes,
	}.Build()
	File_j5_messaging_v1_annotations_proto = out.File
	file_j5_messaging_v1_annotations_proto_rawDesc = nil
	file_j5_messaging_v1_annotations_proto_goTypes = nil
	file_j5_messaging_v1_annotations_proto_depIdxs = nil
}