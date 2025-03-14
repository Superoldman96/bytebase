// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: v1/actuator_service.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

// The request message for getting the theme resource.
type GetResourcePackageRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetResourcePackageRequest) Reset() {
	*x = GetResourcePackageRequest{}
	mi := &file_v1_actuator_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetResourcePackageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourcePackageRequest) ProtoMessage() {}

func (x *GetResourcePackageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_actuator_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourcePackageRequest.ProtoReflect.Descriptor instead.
func (*GetResourcePackageRequest) Descriptor() ([]byte, []int) {
	return file_v1_actuator_service_proto_rawDescGZIP(), []int{0}
}

// The theme resources.
type ResourcePackage struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The branding logo.
	Logo          []byte `protobuf:"bytes,1,opt,name=logo,proto3" json:"logo,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResourcePackage) Reset() {
	*x = ResourcePackage{}
	mi := &file_v1_actuator_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourcePackage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourcePackage) ProtoMessage() {}

func (x *ResourcePackage) ProtoReflect() protoreflect.Message {
	mi := &file_v1_actuator_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourcePackage.ProtoReflect.Descriptor instead.
func (*ResourcePackage) Descriptor() ([]byte, []int) {
	return file_v1_actuator_service_proto_rawDescGZIP(), []int{1}
}

func (x *ResourcePackage) GetLogo() []byte {
	if x != nil {
		return x.Logo
	}
	return nil
}

type GetActuatorInfoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetActuatorInfoRequest) Reset() {
	*x = GetActuatorInfoRequest{}
	mi := &file_v1_actuator_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetActuatorInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetActuatorInfoRequest) ProtoMessage() {}

func (x *GetActuatorInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_actuator_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetActuatorInfoRequest.ProtoReflect.Descriptor instead.
func (*GetActuatorInfoRequest) Descriptor() ([]byte, []int) {
	return file_v1_actuator_service_proto_rawDescGZIP(), []int{2}
}

type UpdateActuatorInfoRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The actuator to update.
	Actuator *ActuatorInfo `protobuf:"bytes,1,opt,name=actuator,proto3" json:"actuator,omitempty"`
	// The list of fields to update.
	UpdateMask    *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateActuatorInfoRequest) Reset() {
	*x = UpdateActuatorInfoRequest{}
	mi := &file_v1_actuator_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateActuatorInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateActuatorInfoRequest) ProtoMessage() {}

func (x *UpdateActuatorInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_actuator_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateActuatorInfoRequest.ProtoReflect.Descriptor instead.
func (*UpdateActuatorInfoRequest) Descriptor() ([]byte, []int) {
	return file_v1_actuator_service_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateActuatorInfoRequest) GetActuator() *ActuatorInfo {
	if x != nil {
		return x.Actuator
	}
	return nil
}

func (x *UpdateActuatorInfoRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type DeleteCacheRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteCacheRequest) Reset() {
	*x = DeleteCacheRequest{}
	mi := &file_v1_actuator_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCacheRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCacheRequest) ProtoMessage() {}

func (x *DeleteCacheRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_actuator_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCacheRequest.ProtoReflect.Descriptor instead.
func (*DeleteCacheRequest) Descriptor() ([]byte, []int) {
	return file_v1_actuator_service_proto_rawDescGZIP(), []int{4}
}

// ServerInfo is the API message for server info.
// Actuator concept is similar to the Spring Boot Actuator.
type ActuatorInfo struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// version is the bytebase's server version
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// git_commit is the git commit hash of the build
	GitCommit string `protobuf:"bytes,2,opt,name=git_commit,json=gitCommit,proto3" json:"git_commit,omitempty"`
	// readonly flag means if the Bytebase is running in readonly mode.
	Readonly bool `protobuf:"varint,3,opt,name=readonly,proto3" json:"readonly,omitempty"`
	// saas flag means if the Bytebase is running in SaaS mode, some features are not allowed to edit by users.
	Saas bool `protobuf:"varint,4,opt,name=saas,proto3" json:"saas,omitempty"`
	// demo flag means if the Bytebase is running in demo mode.
	Demo bool `protobuf:"varint,5,opt,name=demo,proto3" json:"demo,omitempty"`
	// host is the Bytebase instance host.
	Host string `protobuf:"bytes,6,opt,name=host,proto3" json:"host,omitempty"`
	// port is the Bytebase instance port.
	Port string `protobuf:"bytes,7,opt,name=port,proto3" json:"port,omitempty"`
	// external_url is the URL where user or webhook callback visits Bytebase.
	ExternalUrl string `protobuf:"bytes,8,opt,name=external_url,json=externalUrl,proto3" json:"external_url,omitempty"`
	// need_admin_setup flag means the Bytebase instance doesn't have any end users.
	NeedAdminSetup bool `protobuf:"varint,9,opt,name=need_admin_setup,json=needAdminSetup,proto3" json:"need_admin_setup,omitempty"`
	// disallow_signup is the flag to disable self-service signup.
	DisallowSignup bool `protobuf:"varint,10,opt,name=disallow_signup,json=disallowSignup,proto3" json:"disallow_signup,omitempty"`
	// last_active_time is the service last active time in UTC Time Format, any API calls will refresh this value.
	LastActiveTime *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=last_active_time,json=lastActiveTime,proto3" json:"last_active_time,omitempty"`
	// require_2fa is the flag to require 2FA for all users.
	Require_2Fa bool `protobuf:"varint,12,opt,name=require_2fa,json=require2fa,proto3" json:"require_2fa,omitempty"`
	// workspace_id is the identifier for the workspace.
	WorkspaceId string `protobuf:"bytes,13,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	// debug flag means if the debug mode is enabled.
	Debug              bool     `protobuf:"varint,15,opt,name=debug,proto3" json:"debug,omitempty"`
	UnlicensedFeatures []string `protobuf:"bytes,19,rep,name=unlicensed_features,json=unlicensedFeatures,proto3" json:"unlicensed_features,omitempty"`
	// disallow_password_signin is the flag to disallow user signin with email&password. (except workspace admins)
	DisallowPasswordSignin bool                        `protobuf:"varint,20,opt,name=disallow_password_signin,json=disallowPasswordSignin,proto3" json:"disallow_password_signin,omitempty"`
	PasswordRestriction    *PasswordRestrictionSetting `protobuf:"bytes,21,opt,name=password_restriction,json=passwordRestriction,proto3" json:"password_restriction,omitempty"`
	// docker flag means if the Bytebase instance is running in docker.
	Docker        bool `protobuf:"varint,22,opt,name=docker,proto3" json:"docker,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ActuatorInfo) Reset() {
	*x = ActuatorInfo{}
	mi := &file_v1_actuator_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ActuatorInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActuatorInfo) ProtoMessage() {}

func (x *ActuatorInfo) ProtoReflect() protoreflect.Message {
	mi := &file_v1_actuator_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActuatorInfo.ProtoReflect.Descriptor instead.
func (*ActuatorInfo) Descriptor() ([]byte, []int) {
	return file_v1_actuator_service_proto_rawDescGZIP(), []int{5}
}

func (x *ActuatorInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ActuatorInfo) GetGitCommit() string {
	if x != nil {
		return x.GitCommit
	}
	return ""
}

func (x *ActuatorInfo) GetReadonly() bool {
	if x != nil {
		return x.Readonly
	}
	return false
}

func (x *ActuatorInfo) GetSaas() bool {
	if x != nil {
		return x.Saas
	}
	return false
}

func (x *ActuatorInfo) GetDemo() bool {
	if x != nil {
		return x.Demo
	}
	return false
}

func (x *ActuatorInfo) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *ActuatorInfo) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *ActuatorInfo) GetExternalUrl() string {
	if x != nil {
		return x.ExternalUrl
	}
	return ""
}

func (x *ActuatorInfo) GetNeedAdminSetup() bool {
	if x != nil {
		return x.NeedAdminSetup
	}
	return false
}

func (x *ActuatorInfo) GetDisallowSignup() bool {
	if x != nil {
		return x.DisallowSignup
	}
	return false
}

func (x *ActuatorInfo) GetLastActiveTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastActiveTime
	}
	return nil
}

func (x *ActuatorInfo) GetRequire_2Fa() bool {
	if x != nil {
		return x.Require_2Fa
	}
	return false
}

func (x *ActuatorInfo) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *ActuatorInfo) GetDebug() bool {
	if x != nil {
		return x.Debug
	}
	return false
}

func (x *ActuatorInfo) GetUnlicensedFeatures() []string {
	if x != nil {
		return x.UnlicensedFeatures
	}
	return nil
}

func (x *ActuatorInfo) GetDisallowPasswordSignin() bool {
	if x != nil {
		return x.DisallowPasswordSignin
	}
	return false
}

func (x *ActuatorInfo) GetPasswordRestriction() *PasswordRestrictionSetting {
	if x != nil {
		return x.PasswordRestriction
	}
	return nil
}

func (x *ActuatorInfo) GetDocker() bool {
	if x != nil {
		return x.Docker
	}
	return false
}

var File_v1_actuator_service_proto protoreflect.FileDescriptor

var file_v1_actuator_service_proto_rawDesc = string([]byte{
	0x0a, 0x19, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x74, 0x75, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x62, 0x79, 0x74,
	0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x13, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x1b, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x25, 0x0a, 0x0f,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x6c,
	0x6f, 0x67, 0x6f, 0x22, 0x18, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x75, 0x61, 0x74,
	0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x9b, 0x01,
	0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x75, 0x61, 0x74, 0x6f, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x08, 0x61,
	0x63, 0x74, 0x75, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x75,
	0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x08,
	0x61, 0x63, 0x74, 0x75, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x41, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x14, 0x0a, 0x12, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0xf6, 0x05, 0x0a, 0x0c, 0x41, 0x63, 0x74, 0x75, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x1e, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0a, 0x67, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x09, 0x67, 0x69,
	0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x20, 0x0a, 0x08, 0x72, 0x65, 0x61, 0x64, 0x6f,
	0x6e, 0x6c, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52,
	0x08, 0x72, 0x65, 0x61, 0x64, 0x6f, 0x6e, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x04, 0x73, 0x61, 0x61,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x04, 0x73,
	0x61, 0x61, 0x73, 0x12, 0x18, 0x0a, 0x04, 0x64, 0x65, 0x6d, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x04, 0x64, 0x65, 0x6d, 0x6f, 0x12, 0x18, 0x0a,
	0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01,
	0x03, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x27, 0x0a, 0x0c, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x0b, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x12, 0x2e, 0x0a, 0x10, 0x6e, 0x65,
	0x65, 0x64, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x74, 0x75, 0x70, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x08, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x0e, 0x6e, 0x65, 0x65, 0x64,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x74, 0x75, 0x70, 0x12, 0x2d, 0x0a, 0x0f, 0x64, 0x69,
	0x73, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x08, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x0e, 0x64, 0x69, 0x73, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x12, 0x4a, 0x0a, 0x10, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x5f, 0x32, 0x66, 0x61, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03,
	0x52, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x32, 0x66, 0x61, 0x12, 0x27, 0x0a, 0x0c,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x12, 0x2f, 0x0a, 0x13, 0x75,
	0x6e, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x64, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x73, 0x18, 0x13, 0x20, 0x03, 0x28, 0x09, 0x52, 0x12, 0x75, 0x6e, 0x6c, 0x69, 0x63, 0x65,
	0x6e, 0x73, 0x65, 0x64, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x38, 0x0a, 0x18,
	0x64, 0x69, 0x73, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x18, 0x14, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16,
	0x64, 0x69, 0x73, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x12, 0x5a, 0x0a, 0x14, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x74, 0x72,
	0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x13, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x16, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x32, 0x9f, 0x04, 0x0a, 0x0f, 0x41,
	0x63, 0x74, 0x75, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x73,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x75, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x23, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x75, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x75, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x20, 0xda, 0x41, 0x00, 0x80, 0xea, 0x30, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13,
	0x12, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x74, 0x75, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x69,
	0x6e, 0x66, 0x6f, 0x12, 0xaa, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63,
	0x74, 0x75, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x26, 0x2e, 0x62, 0x79, 0x74,
	0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x63, 0x74, 0x75, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x63, 0x74, 0x75, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x51, 0xda,
	0x41, 0x14, 0x61, 0x63, 0x74, 0x75, 0x61, 0x74, 0x6f, 0x72, 0x2c, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x8a, 0xea, 0x30, 0x0f, 0x62, 0x62, 0x2e, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x73, 0x65, 0x74, 0x90, 0xea, 0x30, 0x01, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1d, 0x3a, 0x08, 0x61, 0x63, 0x74, 0x75, 0x61, 0x74, 0x6f, 0x72, 0x32, 0x11, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x63, 0x74, 0x75, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x69, 0x6e, 0x66, 0x6f,
	0x12, 0x66, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12,
	0x1f, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1e, 0x80, 0xea, 0x30, 0x01, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x14, 0x2a, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x74, 0x75, 0x61, 0x74,
	0x6f, 0x72, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x12, 0x81, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12,
	0x26, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x22, 0x25, 0xda, 0x41, 0x00, 0x80, 0xea, 0x30, 0x01, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x74, 0x75, 0x61, 0x74,
	0x6f, 0x72, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x11, 0x5a, 0x0f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_v1_actuator_service_proto_rawDescOnce sync.Once
	file_v1_actuator_service_proto_rawDescData []byte
)

func file_v1_actuator_service_proto_rawDescGZIP() []byte {
	file_v1_actuator_service_proto_rawDescOnce.Do(func() {
		file_v1_actuator_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_v1_actuator_service_proto_rawDesc), len(file_v1_actuator_service_proto_rawDesc)))
	})
	return file_v1_actuator_service_proto_rawDescData
}

var file_v1_actuator_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_v1_actuator_service_proto_goTypes = []any{
	(*GetResourcePackageRequest)(nil),  // 0: bytebase.v1.GetResourcePackageRequest
	(*ResourcePackage)(nil),            // 1: bytebase.v1.ResourcePackage
	(*GetActuatorInfoRequest)(nil),     // 2: bytebase.v1.GetActuatorInfoRequest
	(*UpdateActuatorInfoRequest)(nil),  // 3: bytebase.v1.UpdateActuatorInfoRequest
	(*DeleteCacheRequest)(nil),         // 4: bytebase.v1.DeleteCacheRequest
	(*ActuatorInfo)(nil),               // 5: bytebase.v1.ActuatorInfo
	(*fieldmaskpb.FieldMask)(nil),      // 6: google.protobuf.FieldMask
	(*timestamppb.Timestamp)(nil),      // 7: google.protobuf.Timestamp
	(*PasswordRestrictionSetting)(nil), // 8: bytebase.v1.PasswordRestrictionSetting
	(*emptypb.Empty)(nil),              // 9: google.protobuf.Empty
}
var file_v1_actuator_service_proto_depIdxs = []int32{
	5, // 0: bytebase.v1.UpdateActuatorInfoRequest.actuator:type_name -> bytebase.v1.ActuatorInfo
	6, // 1: bytebase.v1.UpdateActuatorInfoRequest.update_mask:type_name -> google.protobuf.FieldMask
	7, // 2: bytebase.v1.ActuatorInfo.last_active_time:type_name -> google.protobuf.Timestamp
	8, // 3: bytebase.v1.ActuatorInfo.password_restriction:type_name -> bytebase.v1.PasswordRestrictionSetting
	2, // 4: bytebase.v1.ActuatorService.GetActuatorInfo:input_type -> bytebase.v1.GetActuatorInfoRequest
	3, // 5: bytebase.v1.ActuatorService.UpdateActuatorInfo:input_type -> bytebase.v1.UpdateActuatorInfoRequest
	4, // 6: bytebase.v1.ActuatorService.DeleteCache:input_type -> bytebase.v1.DeleteCacheRequest
	0, // 7: bytebase.v1.ActuatorService.GetResourcePackage:input_type -> bytebase.v1.GetResourcePackageRequest
	5, // 8: bytebase.v1.ActuatorService.GetActuatorInfo:output_type -> bytebase.v1.ActuatorInfo
	5, // 9: bytebase.v1.ActuatorService.UpdateActuatorInfo:output_type -> bytebase.v1.ActuatorInfo
	9, // 10: bytebase.v1.ActuatorService.DeleteCache:output_type -> google.protobuf.Empty
	1, // 11: bytebase.v1.ActuatorService.GetResourcePackage:output_type -> bytebase.v1.ResourcePackage
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_v1_actuator_service_proto_init() }
func file_v1_actuator_service_proto_init() {
	if File_v1_actuator_service_proto != nil {
		return
	}
	file_v1_annotation_proto_init()
	file_v1_setting_service_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_v1_actuator_service_proto_rawDesc), len(file_v1_actuator_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_actuator_service_proto_goTypes,
		DependencyIndexes: file_v1_actuator_service_proto_depIdxs,
		MessageInfos:      file_v1_actuator_service_proto_msgTypes,
	}.Build()
	File_v1_actuator_service_proto = out.File
	file_v1_actuator_service_proto_goTypes = nil
	file_v1_actuator_service_proto_depIdxs = nil
}
