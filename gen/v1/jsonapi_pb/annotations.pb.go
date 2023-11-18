// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: jsonapi/v1/annotations.proto

package jsonapi_pb

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

type PackageOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Label  string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	Hidden bool   `protobuf:"varint,2,opt,name=hidden,proto3" json:"hidden,omitempty"`
}

func (x *PackageOptions) Reset() {
	*x = PackageOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jsonapi_v1_annotations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PackageOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PackageOptions) ProtoMessage() {}

func (x *PackageOptions) ProtoReflect() protoreflect.Message {
	mi := &file_jsonapi_v1_annotations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PackageOptions.ProtoReflect.Descriptor instead.
func (*PackageOptions) Descriptor() ([]byte, []int) {
	return file_jsonapi_v1_annotations_proto_rawDescGZIP(), []int{0}
}

func (x *PackageOptions) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *PackageOptions) GetHidden() bool {
	if x != nil {
		return x.Hidden
	}
	return false
}

type MessageOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// When true, all fields in this message should be wrapped in a single oneof
	// field. The message will show in json-schema as-is but with the
	// x-oneof flag set.
	IsOneofWrapper bool `protobuf:"varint,1,opt,name=is_oneof_wrapper,json=isOneofWrapper,proto3" json:"is_oneof_wrapper,omitempty"`
}

func (x *MessageOptions) Reset() {
	*x = MessageOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jsonapi_v1_annotations_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageOptions) ProtoMessage() {}

func (x *MessageOptions) ProtoReflect() protoreflect.Message {
	mi := &file_jsonapi_v1_annotations_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageOptions.ProtoReflect.Descriptor instead.
func (*MessageOptions) Descriptor() ([]byte, []int) {
	return file_jsonapi_v1_annotations_proto_rawDescGZIP(), []int{1}
}

func (x *MessageOptions) GetIsOneofWrapper() bool {
	if x != nil {
		return x.IsOneofWrapper
	}
	return false
}

type OneofOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// When true, the oneof is exposed as a field in the parent message, rather
	// than being a validation rule.
	// Will show in json-schema as an object with the x-oneof flag set.
	Expose bool `protobuf:"varint,1,opt,name=expose,proto3" json:"expose,omitempty"`
}

func (x *OneofOptions) Reset() {
	*x = OneofOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jsonapi_v1_annotations_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OneofOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OneofOptions) ProtoMessage() {}

func (x *OneofOptions) ProtoReflect() protoreflect.Message {
	mi := &file_jsonapi_v1_annotations_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OneofOptions.ProtoReflect.Descriptor instead.
func (*OneofOptions) Descriptor() ([]byte, []int) {
	return file_jsonapi_v1_annotations_proto_rawDescGZIP(), []int{2}
}

func (x *OneofOptions) GetExpose() bool {
	if x != nil {
		return x.Expose
	}
	return false
}

type MethodOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Label  string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	Hidden bool   `protobuf:"varint,2,opt,name=hidden,proto3" json:"hidden,omitempty"`
}

func (x *MethodOptions) Reset() {
	*x = MethodOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jsonapi_v1_annotations_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MethodOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MethodOptions) ProtoMessage() {}

func (x *MethodOptions) ProtoReflect() protoreflect.Message {
	mi := &file_jsonapi_v1_annotations_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MethodOptions.ProtoReflect.Descriptor instead.
func (*MethodOptions) Descriptor() ([]byte, []int) {
	return file_jsonapi_v1_annotations_proto_rawDescGZIP(), []int{3}
}

func (x *MethodOptions) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *MethodOptions) GetHidden() bool {
	if x != nil {
		return x.Hidden
	}
	return false
}

var file_jsonapi_v1_annotations_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*PackageOptions)(nil),
		Field:         90443354,
		Name:          "jsonapi.v1.package",
		Tag:           "bytes,90443354,opt,name=package",
		Filename:      "jsonapi/v1/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*MessageOptions)(nil),
		Field:         90443353,
		Name:          "jsonapi.v1.message",
		Tag:           "bytes,90443353,opt,name=message",
		Filename:      "jsonapi/v1/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.OneofOptions)(nil),
		ExtensionType: (*OneofOptions)(nil),
		Field:         90443355,
		Name:          "jsonapi.v1.oneof",
		Tag:           "bytes,90443355,opt,name=oneof",
		Filename:      "jsonapi/v1/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*MethodOptions)(nil),
		Field:         90443356,
		Name:          "jsonapi.v1.method",
		Tag:           "bytes,90443356,opt,name=method",
		Filename:      "jsonapi/v1/annotations.proto",
	},
}

// Extension fields to descriptorpb.FileOptions.
var (
	// optional jsonapi.v1.PackageOptions package = 90443354;
	E_Package = &file_jsonapi_v1_annotations_proto_extTypes[0]
)

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional jsonapi.v1.MessageOptions message = 90443353;
	E_Message = &file_jsonapi_v1_annotations_proto_extTypes[1]
)

// Extension fields to descriptorpb.OneofOptions.
var (
	// optional jsonapi.v1.OneofOptions oneof = 90443355;
	E_Oneof = &file_jsonapi_v1_annotations_proto_extTypes[2]
)

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional jsonapi.v1.MethodOptions method = 90443356;
	E_Method = &file_jsonapi_v1_annotations_proto_extTypes[3]
)

var File_jsonapi_v1_annotations_proto protoreflect.FileDescriptor

var file_jsonapi_v1_annotations_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6a, 0x73, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x6a, 0x73, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x0e,
	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x22, 0x3a, 0x0a, 0x0e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x28,
	0x0a, 0x10, 0x69, 0x73, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x77, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x69, 0x73, 0x4f, 0x6e, 0x65, 0x6f,
	0x66, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x22, 0x26, 0x0a, 0x0c, 0x4f, 0x6e, 0x65, 0x6f,
	0x66, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x6f,
	0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x65,
	0x22, 0x3d, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x3a,
	0x55, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xda, 0x9c, 0x90, 0x2b, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x70,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x3a, 0x58, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xd9, 0x9c, 0x90, 0x2b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6a, 0x73,
	0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x3a, 0x50, 0x0a, 0x05, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4f, 0x6e, 0x65, 0x6f,
	0x66, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xdb, 0x9c, 0x90, 0x2b, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4f,
	0x6e, 0x65, 0x6f, 0x66, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x05, 0x6f, 0x6e, 0x65,
	0x6f, 0x66, 0x3a, 0x54, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1e, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xdc, 0x9c, 0x90,
	0x2b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x6e, 0x74, 0x6f, 0x70, 0x73, 0x2f, 0x6a,
	0x73, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6a, 0x73,
	0x6f, 0x6e, 0x61, 0x70, 0x69, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_jsonapi_v1_annotations_proto_rawDescOnce sync.Once
	file_jsonapi_v1_annotations_proto_rawDescData = file_jsonapi_v1_annotations_proto_rawDesc
)

func file_jsonapi_v1_annotations_proto_rawDescGZIP() []byte {
	file_jsonapi_v1_annotations_proto_rawDescOnce.Do(func() {
		file_jsonapi_v1_annotations_proto_rawDescData = protoimpl.X.CompressGZIP(file_jsonapi_v1_annotations_proto_rawDescData)
	})
	return file_jsonapi_v1_annotations_proto_rawDescData
}

var file_jsonapi_v1_annotations_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_jsonapi_v1_annotations_proto_goTypes = []interface{}{
	(*PackageOptions)(nil),              // 0: jsonapi.v1.PackageOptions
	(*MessageOptions)(nil),              // 1: jsonapi.v1.MessageOptions
	(*OneofOptions)(nil),                // 2: jsonapi.v1.OneofOptions
	(*MethodOptions)(nil),               // 3: jsonapi.v1.MethodOptions
	(*descriptorpb.FileOptions)(nil),    // 4: google.protobuf.FileOptions
	(*descriptorpb.MessageOptions)(nil), // 5: google.protobuf.MessageOptions
	(*descriptorpb.OneofOptions)(nil),   // 6: google.protobuf.OneofOptions
	(*descriptorpb.MethodOptions)(nil),  // 7: google.protobuf.MethodOptions
}
var file_jsonapi_v1_annotations_proto_depIdxs = []int32{
	4, // 0: jsonapi.v1.package:extendee -> google.protobuf.FileOptions
	5, // 1: jsonapi.v1.message:extendee -> google.protobuf.MessageOptions
	6, // 2: jsonapi.v1.oneof:extendee -> google.protobuf.OneofOptions
	7, // 3: jsonapi.v1.method:extendee -> google.protobuf.MethodOptions
	0, // 4: jsonapi.v1.package:type_name -> jsonapi.v1.PackageOptions
	1, // 5: jsonapi.v1.message:type_name -> jsonapi.v1.MessageOptions
	2, // 6: jsonapi.v1.oneof:type_name -> jsonapi.v1.OneofOptions
	3, // 7: jsonapi.v1.method:type_name -> jsonapi.v1.MethodOptions
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	4, // [4:8] is the sub-list for extension type_name
	0, // [0:4] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_jsonapi_v1_annotations_proto_init() }
func file_jsonapi_v1_annotations_proto_init() {
	if File_jsonapi_v1_annotations_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_jsonapi_v1_annotations_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PackageOptions); i {
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
		file_jsonapi_v1_annotations_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageOptions); i {
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
		file_jsonapi_v1_annotations_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OneofOptions); i {
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
		file_jsonapi_v1_annotations_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MethodOptions); i {
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
			RawDescriptor: file_jsonapi_v1_annotations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 4,
			NumServices:   0,
		},
		GoTypes:           file_jsonapi_v1_annotations_proto_goTypes,
		DependencyIndexes: file_jsonapi_v1_annotations_proto_depIdxs,
		MessageInfos:      file_jsonapi_v1_annotations_proto_msgTypes,
		ExtensionInfos:    file_jsonapi_v1_annotations_proto_extTypes,
	}.Build()
	File_jsonapi_v1_annotations_proto = out.File
	file_jsonapi_v1_annotations_proto_rawDesc = nil
	file_jsonapi_v1_annotations_proto_goTypes = nil
	file_jsonapi_v1_annotations_proto_depIdxs = nil
}
