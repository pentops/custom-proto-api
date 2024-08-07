// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: j5/auth/v1/method.proto

package auth_j5pb

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

type MethodAuthType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//
	//	*MethodAuthType_None_
	//	*MethodAuthType_JwtBearer
	//	*MethodAuthType_Custom_
	//	*MethodAuthType_Cookie_
	Type isMethodAuthType_Type `protobuf_oneof:"type"`
}

func (x *MethodAuthType) Reset() {
	*x = MethodAuthType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_j5_auth_v1_method_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MethodAuthType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MethodAuthType) ProtoMessage() {}

func (x *MethodAuthType) ProtoReflect() protoreflect.Message {
	mi := &file_j5_auth_v1_method_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MethodAuthType.ProtoReflect.Descriptor instead.
func (*MethodAuthType) Descriptor() ([]byte, []int) {
	return file_j5_auth_v1_method_proto_rawDescGZIP(), []int{0}
}

func (m *MethodAuthType) GetType() isMethodAuthType_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *MethodAuthType) GetNone() *MethodAuthType_None {
	if x, ok := x.GetType().(*MethodAuthType_None_); ok {
		return x.None
	}
	return nil
}

func (x *MethodAuthType) GetJwtBearer() *MethodAuthType_JWTBearer {
	if x, ok := x.GetType().(*MethodAuthType_JwtBearer); ok {
		return x.JwtBearer
	}
	return nil
}

func (x *MethodAuthType) GetCustom() *MethodAuthType_Custom {
	if x, ok := x.GetType().(*MethodAuthType_Custom_); ok {
		return x.Custom
	}
	return nil
}

func (x *MethodAuthType) GetCookie() *MethodAuthType_Cookie {
	if x, ok := x.GetType().(*MethodAuthType_Cookie_); ok {
		return x.Cookie
	}
	return nil
}

type isMethodAuthType_Type interface {
	isMethodAuthType_Type()
}

type MethodAuthType_None_ struct {
	None *MethodAuthType_None `protobuf:"bytes,10,opt,name=none,proto3,oneof"`
}

type MethodAuthType_JwtBearer struct {
	JwtBearer *MethodAuthType_JWTBearer `protobuf:"bytes,11,opt,name=jwt_bearer,json=jwtBearer,proto3,oneof"`
}

type MethodAuthType_Custom_ struct {
	Custom *MethodAuthType_Custom `protobuf:"bytes,12,opt,name=custom,proto3,oneof"`
}

type MethodAuthType_Cookie_ struct {
	Cookie *MethodAuthType_Cookie `protobuf:"bytes,13,opt,name=cookie,proto3,oneof"`
}

func (*MethodAuthType_None_) isMethodAuthType_Type() {}

func (*MethodAuthType_JwtBearer) isMethodAuthType_Type() {}

func (*MethodAuthType_Custom_) isMethodAuthType_Type() {}

func (*MethodAuthType_Cookie_) isMethodAuthType_Type() {}

type MethodAuthType_None struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MethodAuthType_None) Reset() {
	*x = MethodAuthType_None{}
	if protoimpl.UnsafeEnabled {
		mi := &file_j5_auth_v1_method_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MethodAuthType_None) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MethodAuthType_None) ProtoMessage() {}

func (x *MethodAuthType_None) ProtoReflect() protoreflect.Message {
	mi := &file_j5_auth_v1_method_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MethodAuthType_None.ProtoReflect.Descriptor instead.
func (*MethodAuthType_None) Descriptor() ([]byte, []int) {
	return file_j5_auth_v1_method_proto_rawDescGZIP(), []int{0, 0}
}

type MethodAuthType_JWTBearer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MethodAuthType_JWTBearer) Reset() {
	*x = MethodAuthType_JWTBearer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_j5_auth_v1_method_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MethodAuthType_JWTBearer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MethodAuthType_JWTBearer) ProtoMessage() {}

func (x *MethodAuthType_JWTBearer) ProtoReflect() protoreflect.Message {
	mi := &file_j5_auth_v1_method_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MethodAuthType_JWTBearer.ProtoReflect.Descriptor instead.
func (*MethodAuthType_JWTBearer) Descriptor() ([]byte, []int) {
	return file_j5_auth_v1_method_proto_rawDescGZIP(), []int{0, 1}
}

type MethodAuthType_Cookie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MethodAuthType_Cookie) Reset() {
	*x = MethodAuthType_Cookie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_j5_auth_v1_method_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MethodAuthType_Cookie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MethodAuthType_Cookie) ProtoMessage() {}

func (x *MethodAuthType_Cookie) ProtoReflect() protoreflect.Message {
	mi := &file_j5_auth_v1_method_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MethodAuthType_Cookie.ProtoReflect.Descriptor instead.
func (*MethodAuthType_Cookie) Descriptor() ([]byte, []int) {
	return file_j5_auth_v1_method_proto_rawDescGZIP(), []int{0, 2}
}

type MethodAuthType_Custom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PassThroughHeaders []string `protobuf:"bytes,1,rep,name=pass_through_headers,json=passThroughHeaders,proto3" json:"pass_through_headers,omitempty"`
}

func (x *MethodAuthType_Custom) Reset() {
	*x = MethodAuthType_Custom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_j5_auth_v1_method_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MethodAuthType_Custom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MethodAuthType_Custom) ProtoMessage() {}

func (x *MethodAuthType_Custom) ProtoReflect() protoreflect.Message {
	mi := &file_j5_auth_v1_method_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MethodAuthType_Custom.ProtoReflect.Descriptor instead.
func (*MethodAuthType_Custom) Descriptor() ([]byte, []int) {
	return file_j5_auth_v1_method_proto_rawDescGZIP(), []int{0, 3}
}

func (x *MethodAuthType_Custom) GetPassThroughHeaders() []string {
	if x != nil {
		return x.PassThroughHeaders
	}
	return nil
}

var File_j5_auth_v1_method_proto protoreflect.FileDescriptor

var file_j5_auth_v1_method_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6a, 0x35, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6a, 0x35, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x76, 0x31, 0x22, 0xeb, 0x02, 0x0a, 0x0e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x41, 0x75, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x6e, 0x6f, 0x6e, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6a, 0x35, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x41, 0x75, 0x74, 0x68, 0x54, 0x79,
	0x70, 0x65, 0x2e, 0x4e, 0x6f, 0x6e, 0x65, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x6f, 0x6e, 0x65, 0x12,
	0x45, 0x0a, 0x0a, 0x6a, 0x77, 0x74, 0x5f, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6a, 0x35, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x41, 0x75, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x2e,
	0x4a, 0x57, 0x54, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x48, 0x00, 0x52, 0x09, 0x6a, 0x77, 0x74,
	0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6a, 0x35, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x41, 0x75, 0x74, 0x68, 0x54, 0x79,
	0x70, 0x65, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x48, 0x00, 0x52, 0x06, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x12, 0x3b, 0x0a, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6a, 0x35, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x41, 0x75, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x2e,
	0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x48, 0x00, 0x52, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65,
	0x1a, 0x06, 0x0a, 0x04, 0x4e, 0x6f, 0x6e, 0x65, 0x1a, 0x0b, 0x0a, 0x09, 0x4a, 0x57, 0x54, 0x42,
	0x65, 0x61, 0x72, 0x65, 0x72, 0x1a, 0x08, 0x0a, 0x06, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x1a,
	0x3a, 0x0a, 0x06, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x12, 0x30, 0x0a, 0x14, 0x70, 0x61, 0x73,
	0x73, 0x5f, 0x74, 0x68, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x12, 0x70, 0x61, 0x73, 0x73, 0x54, 0x68, 0x72,
	0x6f, 0x75, 0x67, 0x68, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x42, 0x06, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x70, 0x65, 0x6e, 0x74, 0x6f, 0x70, 0x73, 0x2f, 0x6a, 0x35, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x6a, 0x35, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x6a, 0x35, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_j5_auth_v1_method_proto_rawDescOnce sync.Once
	file_j5_auth_v1_method_proto_rawDescData = file_j5_auth_v1_method_proto_rawDesc
)

func file_j5_auth_v1_method_proto_rawDescGZIP() []byte {
	file_j5_auth_v1_method_proto_rawDescOnce.Do(func() {
		file_j5_auth_v1_method_proto_rawDescData = protoimpl.X.CompressGZIP(file_j5_auth_v1_method_proto_rawDescData)
	})
	return file_j5_auth_v1_method_proto_rawDescData
}

var file_j5_auth_v1_method_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_j5_auth_v1_method_proto_goTypes = []any{
	(*MethodAuthType)(nil),           // 0: j5.auth.v1.MethodAuthType
	(*MethodAuthType_None)(nil),      // 1: j5.auth.v1.MethodAuthType.None
	(*MethodAuthType_JWTBearer)(nil), // 2: j5.auth.v1.MethodAuthType.JWTBearer
	(*MethodAuthType_Cookie)(nil),    // 3: j5.auth.v1.MethodAuthType.Cookie
	(*MethodAuthType_Custom)(nil),    // 4: j5.auth.v1.MethodAuthType.Custom
}
var file_j5_auth_v1_method_proto_depIdxs = []int32{
	1, // 0: j5.auth.v1.MethodAuthType.none:type_name -> j5.auth.v1.MethodAuthType.None
	2, // 1: j5.auth.v1.MethodAuthType.jwt_bearer:type_name -> j5.auth.v1.MethodAuthType.JWTBearer
	4, // 2: j5.auth.v1.MethodAuthType.custom:type_name -> j5.auth.v1.MethodAuthType.Custom
	3, // 3: j5.auth.v1.MethodAuthType.cookie:type_name -> j5.auth.v1.MethodAuthType.Cookie
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_j5_auth_v1_method_proto_init() }
func file_j5_auth_v1_method_proto_init() {
	if File_j5_auth_v1_method_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_j5_auth_v1_method_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*MethodAuthType); i {
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
		file_j5_auth_v1_method_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*MethodAuthType_None); i {
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
		file_j5_auth_v1_method_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*MethodAuthType_JWTBearer); i {
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
		file_j5_auth_v1_method_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*MethodAuthType_Cookie); i {
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
		file_j5_auth_v1_method_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*MethodAuthType_Custom); i {
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
	file_j5_auth_v1_method_proto_msgTypes[0].OneofWrappers = []any{
		(*MethodAuthType_None_)(nil),
		(*MethodAuthType_JwtBearer)(nil),
		(*MethodAuthType_Custom_)(nil),
		(*MethodAuthType_Cookie_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_j5_auth_v1_method_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_j5_auth_v1_method_proto_goTypes,
		DependencyIndexes: file_j5_auth_v1_method_proto_depIdxs,
		MessageInfos:      file_j5_auth_v1_method_proto_msgTypes,
	}.Build()
	File_j5_auth_v1_method_proto = out.File
	file_j5_auth_v1_method_proto_rawDesc = nil
	file_j5_auth_v1_method_proto_goTypes = nil
	file_j5_auth_v1_method_proto_depIdxs = nil
}