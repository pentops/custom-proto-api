// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: j5/config/v1/plugin.proto

package config_j5pb

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

type Plugin int32

const (
	Plugin_PLUGIN_UNSPECIFIED Plugin = 0
	Plugin_PLUGIN_PROTO       Plugin = 1
	Plugin_PLUGIN_J5_CLIENT   Plugin = 2
)

// Enum value maps for Plugin.
var (
	Plugin_name = map[int32]string{
		0: "PLUGIN_UNSPECIFIED",
		1: "PLUGIN_PROTO",
		2: "PLUGIN_J5_CLIENT",
	}
	Plugin_value = map[string]int32{
		"PLUGIN_UNSPECIFIED": 0,
		"PLUGIN_PROTO":       1,
		"PLUGIN_J5_CLIENT":   2,
	}
)

func (x Plugin) Enum() *Plugin {
	p := new(Plugin)
	*p = x
	return p
}

func (x Plugin) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Plugin) Descriptor() protoreflect.EnumDescriptor {
	return file_j5_config_v1_plugin_proto_enumTypes[0].Descriptor()
}

func (Plugin) Type() protoreflect.EnumType {
	return &file_j5_config_v1_plugin_proto_enumTypes[0]
}

func (x Plugin) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Plugin.Descriptor instead.
func (Plugin) EnumDescriptor() ([]byte, []int) {
	return file_j5_config_v1_plugin_proto_rawDescGZIP(), []int{0}
}

type BuildPlugin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// extend the base plugin.
	Base *string `protobuf:"bytes,1,opt,name=base,proto3,oneof" json:"base,omitempty"`
	// the name of this plugin
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Specifes the input and output format
	Type Plugin `protobuf:"varint,3,opt,name=type,proto3,enum=j5.config.v1.Plugin" json:"type,omitempty"`
	// Options for the given input type, e.g. protc-gen options.
	// On conflicts, the last option wins in the order:
	// 1. base opts (including the resolved opts of the base's base if applicable)
	// 2. publish or generate opts (i.e. the parent of the plugin)
	// 3. plugin opts here
	Opts map[string]string `protobuf:"bytes,4,rep,name=opts,proto3" json:"opts,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// For non-docker, an executable in $PATH, also serves as the 'name' of the
	// command for logging. Not passed in when using docker comtainers.
	Cmd  string   `protobuf:"bytes,5,opt,name=cmd,proto3" json:"cmd,omitempty"`
	Args []string `protobuf:"bytes,6,rep,name=args,proto3" json:"args,omitempty"` // CLI Arguments, passed as specified
	// Environment Variables for the command or
	// container, Expansion of runtime variables is performed, the available
	// variables are set by the context calling the build,
	Env []string `protobuf:"bytes,7,rep,name=env,proto3" json:"env,omitempty"`
	// a docker container to replace the local $PATH executable.
	Docker *BuildPlugin_DockerSpec `protobuf:"bytes,8,opt,name=docker,proto3" json:"docker,omitempty"`
}

func (x *BuildPlugin) Reset() {
	*x = BuildPlugin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_j5_config_v1_plugin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildPlugin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildPlugin) ProtoMessage() {}

func (x *BuildPlugin) ProtoReflect() protoreflect.Message {
	mi := &file_j5_config_v1_plugin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildPlugin.ProtoReflect.Descriptor instead.
func (*BuildPlugin) Descriptor() ([]byte, []int) {
	return file_j5_config_v1_plugin_proto_rawDescGZIP(), []int{0}
}

func (x *BuildPlugin) GetBase() string {
	if x != nil && x.Base != nil {
		return *x.Base
	}
	return ""
}

func (x *BuildPlugin) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BuildPlugin) GetType() Plugin {
	if x != nil {
		return x.Type
	}
	return Plugin_PLUGIN_UNSPECIFIED
}

func (x *BuildPlugin) GetOpts() map[string]string {
	if x != nil {
		return x.Opts
	}
	return nil
}

func (x *BuildPlugin) GetCmd() string {
	if x != nil {
		return x.Cmd
	}
	return ""
}

func (x *BuildPlugin) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *BuildPlugin) GetEnv() []string {
	if x != nil {
		return x.Env
	}
	return nil
}

func (x *BuildPlugin) GetDocker() *BuildPlugin_DockerSpec {
	if x != nil {
		return x.Docker
	}
	return nil
}

// TODO: This currently floats without a config, we need to decide if it belongs
// in the repo config or builder shared config. The complication is that the
// builder has access to all pulled images on the host, so linking this to the
// repo is a bit misleading.
type DockerRegistryAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Matches which images this auth applies to
	// e.g. ghrc.io/* or *.dkr.ecr.*.amazonaws.com/*
	Registry string `protobuf:"bytes,1,opt,name=registry,proto3" json:"registry,omitempty"`
	// Supplies the method for auth.
	// Not retuired if the registry matches a known pattern.
	//
	// Types that are assignable to Auth:
	//
	//	*DockerRegistryAuth_Basic_
	//	*DockerRegistryAuth_AwsEcs
	//	*DockerRegistryAuth_Github_
	Auth isDockerRegistryAuth_Auth `protobuf_oneof:"auth"`
}

func (x *DockerRegistryAuth) Reset() {
	*x = DockerRegistryAuth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_j5_config_v1_plugin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DockerRegistryAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DockerRegistryAuth) ProtoMessage() {}

func (x *DockerRegistryAuth) ProtoReflect() protoreflect.Message {
	mi := &file_j5_config_v1_plugin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DockerRegistryAuth.ProtoReflect.Descriptor instead.
func (*DockerRegistryAuth) Descriptor() ([]byte, []int) {
	return file_j5_config_v1_plugin_proto_rawDescGZIP(), []int{1}
}

func (x *DockerRegistryAuth) GetRegistry() string {
	if x != nil {
		return x.Registry
	}
	return ""
}

func (m *DockerRegistryAuth) GetAuth() isDockerRegistryAuth_Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (x *DockerRegistryAuth) GetBasic() *DockerRegistryAuth_Basic {
	if x, ok := x.GetAuth().(*DockerRegistryAuth_Basic_); ok {
		return x.Basic
	}
	return nil
}

func (x *DockerRegistryAuth) GetAwsEcs() *DockerRegistryAuth_AWSECS {
	if x, ok := x.GetAuth().(*DockerRegistryAuth_AwsEcs); ok {
		return x.AwsEcs
	}
	return nil
}

func (x *DockerRegistryAuth) GetGithub() *DockerRegistryAuth_Github {
	if x, ok := x.GetAuth().(*DockerRegistryAuth_Github_); ok {
		return x.Github
	}
	return nil
}

type isDockerRegistryAuth_Auth interface {
	isDockerRegistryAuth_Auth()
}

type DockerRegistryAuth_Basic_ struct {
	Basic *DockerRegistryAuth_Basic `protobuf:"bytes,10,opt,name=basic,proto3,oneof"`
}

type DockerRegistryAuth_AwsEcs struct {
	AwsEcs *DockerRegistryAuth_AWSECS `protobuf:"bytes,11,opt,name=aws_ecs,json=awsEcs,proto3,oneof"` // default if *.dkr.ecr.*.amazonaws.com/*
}

type DockerRegistryAuth_Github_ struct {
	Github *DockerRegistryAuth_Github `protobuf:"bytes,12,opt,name=github,proto3,oneof"` // default if ghrc.io/*
}

func (*DockerRegistryAuth_Basic_) isDockerRegistryAuth_Auth() {}

func (*DockerRegistryAuth_AwsEcs) isDockerRegistryAuth_Auth() {}

func (*DockerRegistryAuth_Github_) isDockerRegistryAuth_Auth() {}

type BuildPlugin_DockerSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	// passed as-is to the docker API, does not default to or reference Command.cmd
	Entrypoint []string `protobuf:"bytes,2,rep,name=entrypoint,proto3" json:"entrypoint,omitempty"`
	// passed as-is to the docker API, does not default to or reference Command.cmd
	Cmd []string `protobuf:"bytes,3,rep,name=cmd,proto3" json:"cmd,omitempty"`
}

func (x *BuildPlugin_DockerSpec) Reset() {
	*x = BuildPlugin_DockerSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_j5_config_v1_plugin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildPlugin_DockerSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildPlugin_DockerSpec) ProtoMessage() {}

func (x *BuildPlugin_DockerSpec) ProtoReflect() protoreflect.Message {
	mi := &file_j5_config_v1_plugin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildPlugin_DockerSpec.ProtoReflect.Descriptor instead.
func (*BuildPlugin_DockerSpec) Descriptor() ([]byte, []int) {
	return file_j5_config_v1_plugin_proto_rawDescGZIP(), []int{0, 1}
}

func (x *BuildPlugin_DockerSpec) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *BuildPlugin_DockerSpec) GetEntrypoint() []string {
	if x != nil {
		return x.Entrypoint
	}
	return nil
}

func (x *BuildPlugin_DockerSpec) GetCmd() []string {
	if x != nil {
		return x.Cmd
	}
	return nil
}

type DockerRegistryAuth_Basic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username       string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	PasswordEnvVar string `protobuf:"bytes,2,opt,name=password_env_var,json=passwordEnvVar,proto3" json:"password_env_var,omitempty"`
}

func (x *DockerRegistryAuth_Basic) Reset() {
	*x = DockerRegistryAuth_Basic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_j5_config_v1_plugin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DockerRegistryAuth_Basic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DockerRegistryAuth_Basic) ProtoMessage() {}

func (x *DockerRegistryAuth_Basic) ProtoReflect() protoreflect.Message {
	mi := &file_j5_config_v1_plugin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DockerRegistryAuth_Basic.ProtoReflect.Descriptor instead.
func (*DockerRegistryAuth_Basic) Descriptor() ([]byte, []int) {
	return file_j5_config_v1_plugin_proto_rawDescGZIP(), []int{1, 0}
}

func (x *DockerRegistryAuth_Basic) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *DockerRegistryAuth_Basic) GetPasswordEnvVar() string {
	if x != nil {
		return x.PasswordEnvVar
	}
	return ""
}

type DockerRegistryAuth_AWSECS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DockerRegistryAuth_AWSECS) Reset() {
	*x = DockerRegistryAuth_AWSECS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_j5_config_v1_plugin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DockerRegistryAuth_AWSECS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DockerRegistryAuth_AWSECS) ProtoMessage() {}

func (x *DockerRegistryAuth_AWSECS) ProtoReflect() protoreflect.Message {
	mi := &file_j5_config_v1_plugin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DockerRegistryAuth_AWSECS.ProtoReflect.Descriptor instead.
func (*DockerRegistryAuth_AWSECS) Descriptor() ([]byte, []int) {
	return file_j5_config_v1_plugin_proto_rawDescGZIP(), []int{1, 1}
}

type DockerRegistryAuth_Github struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenEnvVar string `protobuf:"bytes,1,opt,name=token_env_var,json=tokenEnvVar,proto3" json:"token_env_var,omitempty"` // defaults to GITHUB_TOKEN
}

func (x *DockerRegistryAuth_Github) Reset() {
	*x = DockerRegistryAuth_Github{}
	if protoimpl.UnsafeEnabled {
		mi := &file_j5_config_v1_plugin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DockerRegistryAuth_Github) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DockerRegistryAuth_Github) ProtoMessage() {}

func (x *DockerRegistryAuth_Github) ProtoReflect() protoreflect.Message {
	mi := &file_j5_config_v1_plugin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DockerRegistryAuth_Github.ProtoReflect.Descriptor instead.
func (*DockerRegistryAuth_Github) Descriptor() ([]byte, []int) {
	return file_j5_config_v1_plugin_proto_rawDescGZIP(), []int{1, 2}
}

func (x *DockerRegistryAuth_Github) GetTokenEnvVar() string {
	if x != nil {
		return x.TokenEnvVar
	}
	return ""
}

var File_j5_config_v1_plugin_proto protoreflect.FileDescriptor

var file_j5_config_v1_plugin_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6a, 0x35, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6a, 0x35, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x22, 0xab, 0x03, 0x0a, 0x0b, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x17, 0x0a, 0x04, 0x62, 0x61, 0x73,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x6a, 0x35, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x37, 0x0a, 0x04, 0x6f, 0x70, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x6a, 0x35, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x4f, 0x70, 0x74, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x04, 0x6f, 0x70, 0x74, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6d, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x61,
	0x72, 0x67, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e,
	0x76, 0x12, 0x3c, 0x0a, 0x06, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x6a, 0x35, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x44, 0x6f, 0x63,
	0x6b, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x52, 0x06, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x1a,
	0x37, 0x0a, 0x09, 0x4f, 0x70, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x54, 0x0a, 0x0a, 0x44, 0x6f, 0x63, 0x6b,
	0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x65, 0x6e, 0x74, 0x72, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0a, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x63, 0x6d, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x42, 0x07,
	0x0a, 0x05, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x22, 0x86, 0x03, 0x0a, 0x12, 0x44, 0x6f, 0x63, 0x6b,
	0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x41, 0x75, 0x74, 0x68, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x3e, 0x0a, 0x05, 0x62, 0x61,
	0x73, 0x69, 0x63, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6a, 0x35, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x41, 0x75, 0x74, 0x68, 0x2e, 0x42, 0x61, 0x73, 0x69,
	0x63, 0x48, 0x00, 0x52, 0x05, 0x62, 0x61, 0x73, 0x69, 0x63, 0x12, 0x42, 0x0a, 0x07, 0x61, 0x77,
	0x73, 0x5f, 0x65, 0x63, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6a, 0x35,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x63, 0x6b, 0x65,
	0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x41, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x57,
	0x53, 0x45, 0x43, 0x53, 0x48, 0x00, 0x52, 0x06, 0x61, 0x77, 0x73, 0x45, 0x63, 0x73, 0x12, 0x41,
	0x0a, 0x06, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x6a, 0x35, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f,
	0x63, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x41, 0x75, 0x74, 0x68,
	0x2e, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x48, 0x00, 0x52, 0x06, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x1a, 0x4d, 0x0a, 0x05, 0x42, 0x61, 0x73, 0x69, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x5f, 0x65, 0x6e, 0x76, 0x5f, 0x76, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x45, 0x6e, 0x76, 0x56, 0x61, 0x72,
	0x1a, 0x08, 0x0a, 0x06, 0x41, 0x57, 0x53, 0x45, 0x43, 0x53, 0x1a, 0x2c, 0x0a, 0x06, 0x47, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x12, 0x22, 0x0a, 0x0d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x65, 0x6e,
	0x76, 0x5f, 0x76, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x45, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x42, 0x06, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68,
	0x2a, 0x48, 0x0a, 0x06, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x4c,
	0x55, 0x47, 0x49, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x4c, 0x55, 0x47, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f,
	0x54, 0x4f, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x4c, 0x55, 0x47, 0x49, 0x4e, 0x5f, 0x4a,
	0x35, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x6e, 0x74, 0x6f, 0x70, 0x73,
	0x2f, 0x6a, 0x35, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x6a, 0x35, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x6a, 0x35, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_j5_config_v1_plugin_proto_rawDescOnce sync.Once
	file_j5_config_v1_plugin_proto_rawDescData = file_j5_config_v1_plugin_proto_rawDesc
)

func file_j5_config_v1_plugin_proto_rawDescGZIP() []byte {
	file_j5_config_v1_plugin_proto_rawDescOnce.Do(func() {
		file_j5_config_v1_plugin_proto_rawDescData = protoimpl.X.CompressGZIP(file_j5_config_v1_plugin_proto_rawDescData)
	})
	return file_j5_config_v1_plugin_proto_rawDescData
}

var file_j5_config_v1_plugin_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_j5_config_v1_plugin_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_j5_config_v1_plugin_proto_goTypes = []any{
	(Plugin)(0),                       // 0: j5.config.v1.Plugin
	(*BuildPlugin)(nil),               // 1: j5.config.v1.BuildPlugin
	(*DockerRegistryAuth)(nil),        // 2: j5.config.v1.DockerRegistryAuth
	nil,                               // 3: j5.config.v1.BuildPlugin.OptsEntry
	(*BuildPlugin_DockerSpec)(nil),    // 4: j5.config.v1.BuildPlugin.DockerSpec
	(*DockerRegistryAuth_Basic)(nil),  // 5: j5.config.v1.DockerRegistryAuth.Basic
	(*DockerRegistryAuth_AWSECS)(nil), // 6: j5.config.v1.DockerRegistryAuth.AWSECS
	(*DockerRegistryAuth_Github)(nil), // 7: j5.config.v1.DockerRegistryAuth.Github
}
var file_j5_config_v1_plugin_proto_depIdxs = []int32{
	0, // 0: j5.config.v1.BuildPlugin.type:type_name -> j5.config.v1.Plugin
	3, // 1: j5.config.v1.BuildPlugin.opts:type_name -> j5.config.v1.BuildPlugin.OptsEntry
	4, // 2: j5.config.v1.BuildPlugin.docker:type_name -> j5.config.v1.BuildPlugin.DockerSpec
	5, // 3: j5.config.v1.DockerRegistryAuth.basic:type_name -> j5.config.v1.DockerRegistryAuth.Basic
	6, // 4: j5.config.v1.DockerRegistryAuth.aws_ecs:type_name -> j5.config.v1.DockerRegistryAuth.AWSECS
	7, // 5: j5.config.v1.DockerRegistryAuth.github:type_name -> j5.config.v1.DockerRegistryAuth.Github
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_j5_config_v1_plugin_proto_init() }
func file_j5_config_v1_plugin_proto_init() {
	if File_j5_config_v1_plugin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_j5_config_v1_plugin_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*BuildPlugin); i {
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
		file_j5_config_v1_plugin_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*DockerRegistryAuth); i {
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
		file_j5_config_v1_plugin_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*BuildPlugin_DockerSpec); i {
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
		file_j5_config_v1_plugin_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*DockerRegistryAuth_Basic); i {
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
		file_j5_config_v1_plugin_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DockerRegistryAuth_AWSECS); i {
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
		file_j5_config_v1_plugin_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*DockerRegistryAuth_Github); i {
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
	file_j5_config_v1_plugin_proto_msgTypes[0].OneofWrappers = []any{}
	file_j5_config_v1_plugin_proto_msgTypes[1].OneofWrappers = []any{
		(*DockerRegistryAuth_Basic_)(nil),
		(*DockerRegistryAuth_AwsEcs)(nil),
		(*DockerRegistryAuth_Github_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_j5_config_v1_plugin_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_j5_config_v1_plugin_proto_goTypes,
		DependencyIndexes: file_j5_config_v1_plugin_proto_depIdxs,
		EnumInfos:         file_j5_config_v1_plugin_proto_enumTypes,
		MessageInfos:      file_j5_config_v1_plugin_proto_msgTypes,
	}.Build()
	File_j5_config_v1_plugin_proto = out.File
	file_j5_config_v1_plugin_proto_rawDesc = nil
	file_j5_config_v1_plugin_proto_goTypes = nil
	file_j5_config_v1_plugin_proto_depIdxs = nil
}
