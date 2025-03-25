// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: store/plan.proto

package store

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// Type is the database change type.
type PlanConfig_ChangeDatabaseConfig_Type int32

const (
	PlanConfig_ChangeDatabaseConfig_TYPE_UNSPECIFIED PlanConfig_ChangeDatabaseConfig_Type = 0
	// Used for establishing schema baseline, this is used when
	// 1. Onboard the database into Bytebase since Bytebase needs to know the current database schema.
	// 2. Had schema drift and need to re-establish the baseline.
	PlanConfig_ChangeDatabaseConfig_BASELINE PlanConfig_ChangeDatabaseConfig_Type = 1
	// Used for DDL changes including CREATE DATABASE.
	PlanConfig_ChangeDatabaseConfig_MIGRATE PlanConfig_ChangeDatabaseConfig_Type = 2
	// Used for schema changes via state-based schema migration including CREATE DATABASE.
	PlanConfig_ChangeDatabaseConfig_MIGRATE_SDL PlanConfig_ChangeDatabaseConfig_Type = 3
	// Used for DDL changes using gh-ost.
	PlanConfig_ChangeDatabaseConfig_MIGRATE_GHOST PlanConfig_ChangeDatabaseConfig_Type = 4
	// Used for DML change.
	PlanConfig_ChangeDatabaseConfig_DATA PlanConfig_ChangeDatabaseConfig_Type = 6
)

// Enum value maps for PlanConfig_ChangeDatabaseConfig_Type.
var (
	PlanConfig_ChangeDatabaseConfig_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "BASELINE",
		2: "MIGRATE",
		3: "MIGRATE_SDL",
		4: "MIGRATE_GHOST",
		6: "DATA",
	}
	PlanConfig_ChangeDatabaseConfig_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"BASELINE":         1,
		"MIGRATE":          2,
		"MIGRATE_SDL":      3,
		"MIGRATE_GHOST":    4,
		"DATA":             6,
	}
)

func (x PlanConfig_ChangeDatabaseConfig_Type) Enum() *PlanConfig_ChangeDatabaseConfig_Type {
	p := new(PlanConfig_ChangeDatabaseConfig_Type)
	*p = x
	return p
}

func (x PlanConfig_ChangeDatabaseConfig_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlanConfig_ChangeDatabaseConfig_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_store_plan_proto_enumTypes[0].Descriptor()
}

func (PlanConfig_ChangeDatabaseConfig_Type) Type() protoreflect.EnumType {
	return &file_store_plan_proto_enumTypes[0]
}

func (x PlanConfig_ChangeDatabaseConfig_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlanConfig_ChangeDatabaseConfig_Type.Descriptor instead.
func (PlanConfig_ChangeDatabaseConfig_Type) EnumDescriptor() ([]byte, []int) {
	return file_store_plan_proto_rawDescGZIP(), []int{0, 3, 0}
}

type PlanConfig struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	Steps         []*PlanConfig_Step        `protobuf:"bytes,1,rep,name=steps,proto3" json:"steps,omitempty"`
	ReleaseSource *PlanConfig_ReleaseSource `protobuf:"bytes,3,opt,name=release_source,json=releaseSource,proto3" json:"release_source,omitempty"`
	Deployment    *PlanConfig_Deployment    `protobuf:"bytes,4,opt,name=deployment,proto3" json:"deployment,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PlanConfig) Reset() {
	*x = PlanConfig{}
	mi := &file_store_plan_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlanConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlanConfig) ProtoMessage() {}

func (x *PlanConfig) ProtoReflect() protoreflect.Message {
	mi := &file_store_plan_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlanConfig.ProtoReflect.Descriptor instead.
func (*PlanConfig) Descriptor() ([]byte, []int) {
	return file_store_plan_proto_rawDescGZIP(), []int{0}
}

func (x *PlanConfig) GetSteps() []*PlanConfig_Step {
	if x != nil {
		return x.Steps
	}
	return nil
}

func (x *PlanConfig) GetReleaseSource() *PlanConfig_ReleaseSource {
	if x != nil {
		return x.ReleaseSource
	}
	return nil
}

func (x *PlanConfig) GetDeployment() *PlanConfig_Deployment {
	if x != nil {
		return x.Deployment
	}
	return nil
}

type PlanConfig_Step struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Use the title if set.
	// Use a generated title if empty.
	Title         string             `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Specs         []*PlanConfig_Spec `protobuf:"bytes,1,rep,name=specs,proto3" json:"specs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PlanConfig_Step) Reset() {
	*x = PlanConfig_Step{}
	mi := &file_store_plan_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlanConfig_Step) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlanConfig_Step) ProtoMessage() {}

func (x *PlanConfig_Step) ProtoReflect() protoreflect.Message {
	mi := &file_store_plan_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlanConfig_Step.ProtoReflect.Descriptor instead.
func (*PlanConfig_Step) Descriptor() ([]byte, []int) {
	return file_store_plan_proto_rawDescGZIP(), []int{0, 0}
}

func (x *PlanConfig_Step) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PlanConfig_Step) GetSpecs() []*PlanConfig_Spec {
	if x != nil {
		return x.Specs
	}
	return nil
}

type PlanConfig_Spec struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// earliest_allowed_time the earliest execution time of the change.
	EarliestAllowedTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=earliest_allowed_time,json=earliestAllowedTime,proto3" json:"earliest_allowed_time,omitempty"`
	// A UUID4 string that uniquely identifies the Spec.
	Id                string                        `protobuf:"bytes,5,opt,name=id,proto3" json:"id,omitempty"`
	SpecReleaseSource *PlanConfig_SpecReleaseSource `protobuf:"bytes,8,opt,name=spec_release_source,json=specReleaseSource,proto3" json:"spec_release_source,omitempty"`
	// Types that are valid to be assigned to Config:
	//
	//	*PlanConfig_Spec_CreateDatabaseConfig
	//	*PlanConfig_Spec_ChangeDatabaseConfig
	//	*PlanConfig_Spec_ExportDataConfig
	Config        isPlanConfig_Spec_Config `protobuf_oneof:"config"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PlanConfig_Spec) Reset() {
	*x = PlanConfig_Spec{}
	mi := &file_store_plan_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlanConfig_Spec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlanConfig_Spec) ProtoMessage() {}

func (x *PlanConfig_Spec) ProtoReflect() protoreflect.Message {
	mi := &file_store_plan_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlanConfig_Spec.ProtoReflect.Descriptor instead.
func (*PlanConfig_Spec) Descriptor() ([]byte, []int) {
	return file_store_plan_proto_rawDescGZIP(), []int{0, 1}
}

func (x *PlanConfig_Spec) GetEarliestAllowedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EarliestAllowedTime
	}
	return nil
}

func (x *PlanConfig_Spec) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PlanConfig_Spec) GetSpecReleaseSource() *PlanConfig_SpecReleaseSource {
	if x != nil {
		return x.SpecReleaseSource
	}
	return nil
}

func (x *PlanConfig_Spec) GetConfig() isPlanConfig_Spec_Config {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *PlanConfig_Spec) GetCreateDatabaseConfig() *PlanConfig_CreateDatabaseConfig {
	if x != nil {
		if x, ok := x.Config.(*PlanConfig_Spec_CreateDatabaseConfig); ok {
			return x.CreateDatabaseConfig
		}
	}
	return nil
}

func (x *PlanConfig_Spec) GetChangeDatabaseConfig() *PlanConfig_ChangeDatabaseConfig {
	if x != nil {
		if x, ok := x.Config.(*PlanConfig_Spec_ChangeDatabaseConfig); ok {
			return x.ChangeDatabaseConfig
		}
	}
	return nil
}

func (x *PlanConfig_Spec) GetExportDataConfig() *PlanConfig_ExportDataConfig {
	if x != nil {
		if x, ok := x.Config.(*PlanConfig_Spec_ExportDataConfig); ok {
			return x.ExportDataConfig
		}
	}
	return nil
}

type isPlanConfig_Spec_Config interface {
	isPlanConfig_Spec_Config()
}

type PlanConfig_Spec_CreateDatabaseConfig struct {
	CreateDatabaseConfig *PlanConfig_CreateDatabaseConfig `protobuf:"bytes,1,opt,name=create_database_config,json=createDatabaseConfig,proto3,oneof"`
}

type PlanConfig_Spec_ChangeDatabaseConfig struct {
	ChangeDatabaseConfig *PlanConfig_ChangeDatabaseConfig `protobuf:"bytes,2,opt,name=change_database_config,json=changeDatabaseConfig,proto3,oneof"`
}

type PlanConfig_Spec_ExportDataConfig struct {
	ExportDataConfig *PlanConfig_ExportDataConfig `protobuf:"bytes,7,opt,name=export_data_config,json=exportDataConfig,proto3,oneof"`
}

func (*PlanConfig_Spec_CreateDatabaseConfig) isPlanConfig_Spec_Config() {}

func (*PlanConfig_Spec_ChangeDatabaseConfig) isPlanConfig_Spec_Config() {}

func (*PlanConfig_Spec_ExportDataConfig) isPlanConfig_Spec_Config() {}

type PlanConfig_CreateDatabaseConfig struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The resource name of the instance on which the database is created.
	// Format: instances/{instance}
	Target string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	// The name of the database to create.
	Database string `protobuf:"bytes,2,opt,name=database,proto3" json:"database,omitempty"`
	// table is the name of the table, if it is not empty, Bytebase should create a table after creating the database.
	// For example, in MongoDB, it only creates the database when we first store data in that database.
	Table string `protobuf:"bytes,3,opt,name=table,proto3" json:"table,omitempty"`
	// character_set is the character set of the database.
	CharacterSet string `protobuf:"bytes,4,opt,name=character_set,json=characterSet,proto3" json:"character_set,omitempty"`
	// collation is the collation of the database.
	Collation string `protobuf:"bytes,5,opt,name=collation,proto3" json:"collation,omitempty"`
	// cluster is the cluster of the database. This is only applicable to ClickHouse for "ON CLUSTER <<cluster>>".
	Cluster string `protobuf:"bytes,6,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// owner is the owner of the database. This is only applicable to Postgres for "WITH OWNER <<owner>>".
	Owner string `protobuf:"bytes,7,opt,name=owner,proto3" json:"owner,omitempty"`
	// backup is the resource name of the backup.
	// Format: instances/{instance}/databases/{database}/backups/{backup-name}
	Backup string `protobuf:"bytes,8,opt,name=backup,proto3" json:"backup,omitempty"`
	// The environment resource.
	// Format: environments/prod where prod is the environment resource ID.
	Environment   string `protobuf:"bytes,9,opt,name=environment,proto3" json:"environment,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PlanConfig_CreateDatabaseConfig) Reset() {
	*x = PlanConfig_CreateDatabaseConfig{}
	mi := &file_store_plan_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlanConfig_CreateDatabaseConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlanConfig_CreateDatabaseConfig) ProtoMessage() {}

func (x *PlanConfig_CreateDatabaseConfig) ProtoReflect() protoreflect.Message {
	mi := &file_store_plan_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlanConfig_CreateDatabaseConfig.ProtoReflect.Descriptor instead.
func (*PlanConfig_CreateDatabaseConfig) Descriptor() ([]byte, []int) {
	return file_store_plan_proto_rawDescGZIP(), []int{0, 2}
}

func (x *PlanConfig_CreateDatabaseConfig) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *PlanConfig_CreateDatabaseConfig) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *PlanConfig_CreateDatabaseConfig) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *PlanConfig_CreateDatabaseConfig) GetCharacterSet() string {
	if x != nil {
		return x.CharacterSet
	}
	return ""
}

func (x *PlanConfig_CreateDatabaseConfig) GetCollation() string {
	if x != nil {
		return x.Collation
	}
	return ""
}

func (x *PlanConfig_CreateDatabaseConfig) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *PlanConfig_CreateDatabaseConfig) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *PlanConfig_CreateDatabaseConfig) GetBackup() string {
	if x != nil {
		return x.Backup
	}
	return ""
}

func (x *PlanConfig_CreateDatabaseConfig) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

type PlanConfig_ChangeDatabaseConfig struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The resource name of the target.
	// Format: instances/{instance-id}/databases/{database-name}.
	// Format: projects/{project}/databaseGroups/{databaseGroup}.
	Target string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	// The resource name of the sheet.
	// Format: projects/{project}/sheets/{sheet}
	Sheet string                               `protobuf:"bytes,2,opt,name=sheet,proto3" json:"sheet,omitempty"`
	Type  PlanConfig_ChangeDatabaseConfig_Type `protobuf:"varint,3,opt,name=type,proto3,enum=bytebase.store.PlanConfig_ChangeDatabaseConfig_Type" json:"type,omitempty"`
	// schema_version is parsed from file name.
	// It is automatically generated in the UI workflow.
	SchemaVersion string            `protobuf:"bytes,4,opt,name=schema_version,json=schemaVersion,proto3" json:"schema_version,omitempty"`
	GhostFlags    map[string]string `protobuf:"bytes,7,rep,name=ghost_flags,json=ghostFlags,proto3" json:"ghost_flags,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// If set, a backup of the modified data will be created automatically before any changes are applied.
	PreUpdateBackupDetail *PreUpdateBackupDetail `protobuf:"bytes,8,opt,name=pre_update_backup_detail,json=preUpdateBackupDetail,proto3,oneof" json:"pre_update_backup_detail,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *PlanConfig_ChangeDatabaseConfig) Reset() {
	*x = PlanConfig_ChangeDatabaseConfig{}
	mi := &file_store_plan_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlanConfig_ChangeDatabaseConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlanConfig_ChangeDatabaseConfig) ProtoMessage() {}

func (x *PlanConfig_ChangeDatabaseConfig) ProtoReflect() protoreflect.Message {
	mi := &file_store_plan_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlanConfig_ChangeDatabaseConfig.ProtoReflect.Descriptor instead.
func (*PlanConfig_ChangeDatabaseConfig) Descriptor() ([]byte, []int) {
	return file_store_plan_proto_rawDescGZIP(), []int{0, 3}
}

func (x *PlanConfig_ChangeDatabaseConfig) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *PlanConfig_ChangeDatabaseConfig) GetSheet() string {
	if x != nil {
		return x.Sheet
	}
	return ""
}

func (x *PlanConfig_ChangeDatabaseConfig) GetType() PlanConfig_ChangeDatabaseConfig_Type {
	if x != nil {
		return x.Type
	}
	return PlanConfig_ChangeDatabaseConfig_TYPE_UNSPECIFIED
}

func (x *PlanConfig_ChangeDatabaseConfig) GetSchemaVersion() string {
	if x != nil {
		return x.SchemaVersion
	}
	return ""
}

func (x *PlanConfig_ChangeDatabaseConfig) GetGhostFlags() map[string]string {
	if x != nil {
		return x.GhostFlags
	}
	return nil
}

func (x *PlanConfig_ChangeDatabaseConfig) GetPreUpdateBackupDetail() *PreUpdateBackupDetail {
	if x != nil {
		return x.PreUpdateBackupDetail
	}
	return nil
}

type PlanConfig_ExportDataConfig struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The resource name of the target.
	// Format: instances/{instance-id}/databases/{database-name}
	Target string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	// The resource name of the sheet.
	// Format: projects/{project}/sheets/{sheet}
	Sheet string `protobuf:"bytes,2,opt,name=sheet,proto3" json:"sheet,omitempty"`
	// The format of the exported file.
	Format ExportFormat `protobuf:"varint,3,opt,name=format,proto3,enum=bytebase.store.ExportFormat" json:"format,omitempty"`
	// The zip password provide by users.
	// Leave it empty if no needs to encrypt the zip file.
	Password      *string `protobuf:"bytes,4,opt,name=password,proto3,oneof" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PlanConfig_ExportDataConfig) Reset() {
	*x = PlanConfig_ExportDataConfig{}
	mi := &file_store_plan_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlanConfig_ExportDataConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlanConfig_ExportDataConfig) ProtoMessage() {}

func (x *PlanConfig_ExportDataConfig) ProtoReflect() protoreflect.Message {
	mi := &file_store_plan_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlanConfig_ExportDataConfig.ProtoReflect.Descriptor instead.
func (*PlanConfig_ExportDataConfig) Descriptor() ([]byte, []int) {
	return file_store_plan_proto_rawDescGZIP(), []int{0, 4}
}

func (x *PlanConfig_ExportDataConfig) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *PlanConfig_ExportDataConfig) GetSheet() string {
	if x != nil {
		return x.Sheet
	}
	return ""
}

func (x *PlanConfig_ExportDataConfig) GetFormat() ExportFormat {
	if x != nil {
		return x.Format
	}
	return ExportFormat_FORMAT_UNSPECIFIED
}

func (x *PlanConfig_ExportDataConfig) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

type PlanConfig_ReleaseSource struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The release.
	// Format: projects/{project}/releases/{release}
	Release       string `protobuf:"bytes,1,opt,name=release,proto3" json:"release,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PlanConfig_ReleaseSource) Reset() {
	*x = PlanConfig_ReleaseSource{}
	mi := &file_store_plan_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlanConfig_ReleaseSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlanConfig_ReleaseSource) ProtoMessage() {}

func (x *PlanConfig_ReleaseSource) ProtoReflect() protoreflect.Message {
	mi := &file_store_plan_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlanConfig_ReleaseSource.ProtoReflect.Descriptor instead.
func (*PlanConfig_ReleaseSource) Descriptor() ([]byte, []int) {
	return file_store_plan_proto_rawDescGZIP(), []int{0, 5}
}

func (x *PlanConfig_ReleaseSource) GetRelease() string {
	if x != nil {
		return x.Release
	}
	return ""
}

type PlanConfig_SpecReleaseSource struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Format: projects/{project}/releases/{release}/files/{id}
	File          string `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PlanConfig_SpecReleaseSource) Reset() {
	*x = PlanConfig_SpecReleaseSource{}
	mi := &file_store_plan_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlanConfig_SpecReleaseSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlanConfig_SpecReleaseSource) ProtoMessage() {}

func (x *PlanConfig_SpecReleaseSource) ProtoReflect() protoreflect.Message {
	mi := &file_store_plan_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlanConfig_SpecReleaseSource.ProtoReflect.Descriptor instead.
func (*PlanConfig_SpecReleaseSource) Descriptor() ([]byte, []int) {
	return file_store_plan_proto_rawDescGZIP(), []int{0, 6}
}

func (x *PlanConfig_SpecReleaseSource) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

type PlanConfig_Deployment struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The environments deploy order.
	Environments []string `protobuf:"bytes,1,rep,name=environments,proto3" json:"environments,omitempty"`
	// The database group mapping.
	DatabaseGroupMappings []*PlanConfig_Deployment_DatabaseGroupMapping `protobuf:"bytes,2,rep,name=database_group_mappings,json=databaseGroupMappings,proto3" json:"database_group_mappings,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *PlanConfig_Deployment) Reset() {
	*x = PlanConfig_Deployment{}
	mi := &file_store_plan_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlanConfig_Deployment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlanConfig_Deployment) ProtoMessage() {}

func (x *PlanConfig_Deployment) ProtoReflect() protoreflect.Message {
	mi := &file_store_plan_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlanConfig_Deployment.ProtoReflect.Descriptor instead.
func (*PlanConfig_Deployment) Descriptor() ([]byte, []int) {
	return file_store_plan_proto_rawDescGZIP(), []int{0, 7}
}

func (x *PlanConfig_Deployment) GetEnvironments() []string {
	if x != nil {
		return x.Environments
	}
	return nil
}

func (x *PlanConfig_Deployment) GetDatabaseGroupMappings() []*PlanConfig_Deployment_DatabaseGroupMapping {
	if x != nil {
		return x.DatabaseGroupMappings
	}
	return nil
}

type PlanConfig_Deployment_DatabaseGroupMapping struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Format: projects/{project}/databaseGroups/{databaseGroup}.
	DatabaseGroup string `protobuf:"bytes,1,opt,name=database_group,json=databaseGroup,proto3" json:"database_group,omitempty"`
	// Format: instances/{instance-id}/databases/{database-name}.
	Databases     []string `protobuf:"bytes,2,rep,name=databases,proto3" json:"databases,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PlanConfig_Deployment_DatabaseGroupMapping) Reset() {
	*x = PlanConfig_Deployment_DatabaseGroupMapping{}
	mi := &file_store_plan_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlanConfig_Deployment_DatabaseGroupMapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlanConfig_Deployment_DatabaseGroupMapping) ProtoMessage() {}

func (x *PlanConfig_Deployment_DatabaseGroupMapping) ProtoReflect() protoreflect.Message {
	mi := &file_store_plan_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlanConfig_Deployment_DatabaseGroupMapping.ProtoReflect.Descriptor instead.
func (*PlanConfig_Deployment_DatabaseGroupMapping) Descriptor() ([]byte, []int) {
	return file_store_plan_proto_rawDescGZIP(), []int{0, 7, 0}
}

func (x *PlanConfig_Deployment_DatabaseGroupMapping) GetDatabaseGroup() string {
	if x != nil {
		return x.DatabaseGroup
	}
	return ""
}

func (x *PlanConfig_Deployment_DatabaseGroupMapping) GetDatabases() []string {
	if x != nil {
		return x.Databases
	}
	return nil
}

var File_store_plan_proto protoreflect.FileDescriptor

const file_store_plan_proto_rawDesc = "" +
	"\n" +
	"\x10store/plan.proto\x12\x0ebytebase.store\x1a\x1fgoogle/api/field_behavior.proto\x1a\x19google/api/resource.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x12store/common.proto\x1a\x1astore/plan_check_run.proto\"\xde\x11\n" +
	"\n" +
	"PlanConfig\x125\n" +
	"\x05steps\x18\x01 \x03(\v2\x1f.bytebase.store.PlanConfig.StepR\x05steps\x12O\n" +
	"\x0erelease_source\x18\x03 \x01(\v2(.bytebase.store.PlanConfig.ReleaseSourceR\rreleaseSource\x12E\n" +
	"\n" +
	"deployment\x18\x04 \x01(\v2%.bytebase.store.PlanConfig.DeploymentR\n" +
	"deployment\x1aS\n" +
	"\x04Step\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x125\n" +
	"\x05specs\x18\x01 \x03(\v2\x1f.bytebase.store.PlanConfig.SpecR\x05specs\x1a\xfd\x03\n" +
	"\x04Spec\x12N\n" +
	"\x15earliest_allowed_time\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\x13earliestAllowedTime\x12\x0e\n" +
	"\x02id\x18\x05 \x01(\tR\x02id\x12\\\n" +
	"\x13spec_release_source\x18\b \x01(\v2,.bytebase.store.PlanConfig.SpecReleaseSourceR\x11specReleaseSource\x12g\n" +
	"\x16create_database_config\x18\x01 \x01(\v2/.bytebase.store.PlanConfig.CreateDatabaseConfigH\x00R\x14createDatabaseConfig\x12g\n" +
	"\x16change_database_config\x18\x02 \x01(\v2/.bytebase.store.PlanConfig.ChangeDatabaseConfigH\x00R\x14changeDatabaseConfig\x12[\n" +
	"\x12export_data_config\x18\a \x01(\v2+.bytebase.store.PlanConfig.ExportDataConfigH\x00R\x10exportDataConfigB\b\n" +
	"\x06config\x1a\xc3\x02\n" +
	"\x14CreateDatabaseConfig\x12\x1c\n" +
	"\x06target\x18\x01 \x01(\tB\x04\xe2A\x01\x02R\x06target\x12 \n" +
	"\bdatabase\x18\x02 \x01(\tB\x04\xe2A\x01\x02R\bdatabase\x12\x1a\n" +
	"\x05table\x18\x03 \x01(\tB\x04\xe2A\x01\x01R\x05table\x12)\n" +
	"\rcharacter_set\x18\x04 \x01(\tB\x04\xe2A\x01\x01R\fcharacterSet\x12\"\n" +
	"\tcollation\x18\x05 \x01(\tB\x04\xe2A\x01\x01R\tcollation\x12\x1e\n" +
	"\acluster\x18\x06 \x01(\tB\x04\xe2A\x01\x01R\acluster\x12\x1a\n" +
	"\x05owner\x18\a \x01(\tB\x04\xe2A\x01\x01R\x05owner\x12\x1c\n" +
	"\x06backup\x18\b \x01(\tB\x04\xe2A\x01\x01R\x06backup\x12&\n" +
	"\venvironment\x18\t \x01(\tB\x04\xe2A\x01\x01R\venvironment\x1a\xcb\x04\n" +
	"\x14ChangeDatabaseConfig\x12\x16\n" +
	"\x06target\x18\x01 \x01(\tR\x06target\x12\x14\n" +
	"\x05sheet\x18\x02 \x01(\tR\x05sheet\x12H\n" +
	"\x04type\x18\x03 \x01(\x0e24.bytebase.store.PlanConfig.ChangeDatabaseConfig.TypeR\x04type\x12%\n" +
	"\x0eschema_version\x18\x04 \x01(\tR\rschemaVersion\x12`\n" +
	"\vghost_flags\x18\a \x03(\v2?.bytebase.store.PlanConfig.ChangeDatabaseConfig.GhostFlagsEntryR\n" +
	"ghostFlags\x12c\n" +
	"\x18pre_update_backup_detail\x18\b \x01(\v2%.bytebase.store.PreUpdateBackupDetailH\x00R\x15preUpdateBackupDetail\x88\x01\x01\x1a=\n" +
	"\x0fGhostFlagsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"e\n" +
	"\x04Type\x12\x14\n" +
	"\x10TYPE_UNSPECIFIED\x10\x00\x12\f\n" +
	"\bBASELINE\x10\x01\x12\v\n" +
	"\aMIGRATE\x10\x02\x12\x0f\n" +
	"\vMIGRATE_SDL\x10\x03\x12\x11\n" +
	"\rMIGRATE_GHOST\x10\x04\x12\b\n" +
	"\x04DATA\x10\x06B\x1b\n" +
	"\x19_pre_update_backup_detailJ\x04\b\x05\x10\x06J\x04\b\x06\x10\a\x1a\xa4\x01\n" +
	"\x10ExportDataConfig\x12\x16\n" +
	"\x06target\x18\x01 \x01(\tR\x06target\x12\x14\n" +
	"\x05sheet\x18\x02 \x01(\tR\x05sheet\x124\n" +
	"\x06format\x18\x03 \x01(\x0e2\x1c.bytebase.store.ExportFormatR\x06format\x12\x1f\n" +
	"\bpassword\x18\x04 \x01(\tH\x00R\bpassword\x88\x01\x01B\v\n" +
	"\t_password\x1aD\n" +
	"\rReleaseSource\x123\n" +
	"\arelease\x18\x01 \x01(\tB\x19\xfaA\x16\n" +
	"\x14bytebase.com/ReleaseR\arelease\x1a'\n" +
	"\x11SpecReleaseSource\x12\x12\n" +
	"\x04file\x18\x01 \x01(\tR\x04file\x1a\x81\x02\n" +
	"\n" +
	"Deployment\x12\"\n" +
	"\fenvironments\x18\x01 \x03(\tR\fenvironments\x12r\n" +
	"\x17database_group_mappings\x18\x02 \x03(\v2:.bytebase.store.PlanConfig.Deployment.DatabaseGroupMappingR\x15databaseGroupMappings\x1a[\n" +
	"\x14DatabaseGroupMapping\x12%\n" +
	"\x0edatabase_group\x18\x01 \x01(\tR\rdatabaseGroup\x12\x1c\n" +
	"\tdatabases\x18\x02 \x03(\tR\tdatabasesB\x14Z\x12generated-go/storeb\x06proto3"

var (
	file_store_plan_proto_rawDescOnce sync.Once
	file_store_plan_proto_rawDescData []byte
)

func file_store_plan_proto_rawDescGZIP() []byte {
	file_store_plan_proto_rawDescOnce.Do(func() {
		file_store_plan_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_store_plan_proto_rawDesc), len(file_store_plan_proto_rawDesc)))
	})
	return file_store_plan_proto_rawDescData
}

var file_store_plan_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_store_plan_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_store_plan_proto_goTypes = []any{
	(PlanConfig_ChangeDatabaseConfig_Type)(0),          // 0: bytebase.store.PlanConfig.ChangeDatabaseConfig.Type
	(*PlanConfig)(nil),                                 // 1: bytebase.store.PlanConfig
	(*PlanConfig_Step)(nil),                            // 2: bytebase.store.PlanConfig.Step
	(*PlanConfig_Spec)(nil),                            // 3: bytebase.store.PlanConfig.Spec
	(*PlanConfig_CreateDatabaseConfig)(nil),            // 4: bytebase.store.PlanConfig.CreateDatabaseConfig
	(*PlanConfig_ChangeDatabaseConfig)(nil),            // 5: bytebase.store.PlanConfig.ChangeDatabaseConfig
	(*PlanConfig_ExportDataConfig)(nil),                // 6: bytebase.store.PlanConfig.ExportDataConfig
	(*PlanConfig_ReleaseSource)(nil),                   // 7: bytebase.store.PlanConfig.ReleaseSource
	(*PlanConfig_SpecReleaseSource)(nil),               // 8: bytebase.store.PlanConfig.SpecReleaseSource
	(*PlanConfig_Deployment)(nil),                      // 9: bytebase.store.PlanConfig.Deployment
	nil,                                                // 10: bytebase.store.PlanConfig.ChangeDatabaseConfig.GhostFlagsEntry
	(*PlanConfig_Deployment_DatabaseGroupMapping)(nil), // 11: bytebase.store.PlanConfig.Deployment.DatabaseGroupMapping
	(*timestamppb.Timestamp)(nil),                      // 12: google.protobuf.Timestamp
	(*PreUpdateBackupDetail)(nil),                      // 13: bytebase.store.PreUpdateBackupDetail
	(ExportFormat)(0),                                  // 14: bytebase.store.ExportFormat
}
var file_store_plan_proto_depIdxs = []int32{
	2,  // 0: bytebase.store.PlanConfig.steps:type_name -> bytebase.store.PlanConfig.Step
	7,  // 1: bytebase.store.PlanConfig.release_source:type_name -> bytebase.store.PlanConfig.ReleaseSource
	9,  // 2: bytebase.store.PlanConfig.deployment:type_name -> bytebase.store.PlanConfig.Deployment
	3,  // 3: bytebase.store.PlanConfig.Step.specs:type_name -> bytebase.store.PlanConfig.Spec
	12, // 4: bytebase.store.PlanConfig.Spec.earliest_allowed_time:type_name -> google.protobuf.Timestamp
	8,  // 5: bytebase.store.PlanConfig.Spec.spec_release_source:type_name -> bytebase.store.PlanConfig.SpecReleaseSource
	4,  // 6: bytebase.store.PlanConfig.Spec.create_database_config:type_name -> bytebase.store.PlanConfig.CreateDatabaseConfig
	5,  // 7: bytebase.store.PlanConfig.Spec.change_database_config:type_name -> bytebase.store.PlanConfig.ChangeDatabaseConfig
	6,  // 8: bytebase.store.PlanConfig.Spec.export_data_config:type_name -> bytebase.store.PlanConfig.ExportDataConfig
	0,  // 9: bytebase.store.PlanConfig.ChangeDatabaseConfig.type:type_name -> bytebase.store.PlanConfig.ChangeDatabaseConfig.Type
	10, // 10: bytebase.store.PlanConfig.ChangeDatabaseConfig.ghost_flags:type_name -> bytebase.store.PlanConfig.ChangeDatabaseConfig.GhostFlagsEntry
	13, // 11: bytebase.store.PlanConfig.ChangeDatabaseConfig.pre_update_backup_detail:type_name -> bytebase.store.PreUpdateBackupDetail
	14, // 12: bytebase.store.PlanConfig.ExportDataConfig.format:type_name -> bytebase.store.ExportFormat
	11, // 13: bytebase.store.PlanConfig.Deployment.database_group_mappings:type_name -> bytebase.store.PlanConfig.Deployment.DatabaseGroupMapping
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_store_plan_proto_init() }
func file_store_plan_proto_init() {
	if File_store_plan_proto != nil {
		return
	}
	file_store_common_proto_init()
	file_store_plan_check_run_proto_init()
	file_store_plan_proto_msgTypes[2].OneofWrappers = []any{
		(*PlanConfig_Spec_CreateDatabaseConfig)(nil),
		(*PlanConfig_Spec_ChangeDatabaseConfig)(nil),
		(*PlanConfig_Spec_ExportDataConfig)(nil),
	}
	file_store_plan_proto_msgTypes[4].OneofWrappers = []any{}
	file_store_plan_proto_msgTypes[5].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_store_plan_proto_rawDesc), len(file_store_plan_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_plan_proto_goTypes,
		DependencyIndexes: file_store_plan_proto_depIdxs,
		EnumInfos:         file_store_plan_proto_enumTypes,
		MessageInfos:      file_store_plan_proto_msgTypes,
	}.Build()
	File_store_plan_proto = out.File
	file_store_plan_proto_goTypes = nil
	file_store_plan_proto_depIdxs = nil
}
