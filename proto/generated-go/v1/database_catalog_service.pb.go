// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: v1/database_catalog_service.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ObjectSchema_Type int32

const (
	ObjectSchema_TYPE_UNSPECIFIED ObjectSchema_Type = 0
	ObjectSchema_STRING           ObjectSchema_Type = 1
	ObjectSchema_NUMBER           ObjectSchema_Type = 2
	ObjectSchema_BOOLEAN          ObjectSchema_Type = 3
	ObjectSchema_OBJECT           ObjectSchema_Type = 4
	ObjectSchema_ARRAY            ObjectSchema_Type = 5
)

// Enum value maps for ObjectSchema_Type.
var (
	ObjectSchema_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "STRING",
		2: "NUMBER",
		3: "BOOLEAN",
		4: "OBJECT",
		5: "ARRAY",
	}
	ObjectSchema_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"STRING":           1,
		"NUMBER":           2,
		"BOOLEAN":          3,
		"OBJECT":           4,
		"ARRAY":            5,
	}
)

func (x ObjectSchema_Type) Enum() *ObjectSchema_Type {
	p := new(ObjectSchema_Type)
	*p = x
	return p
}

func (x ObjectSchema_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ObjectSchema_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_database_catalog_service_proto_enumTypes[0].Descriptor()
}

func (ObjectSchema_Type) Type() protoreflect.EnumType {
	return &file_v1_database_catalog_service_proto_enumTypes[0]
}

func (x ObjectSchema_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ObjectSchema_Type.Descriptor instead.
func (ObjectSchema_Type) EnumDescriptor() ([]byte, []int) {
	return file_v1_database_catalog_service_proto_rawDescGZIP(), []int{6, 0}
}

type GetDatabaseCatalogRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The name of the database catalog to retrieve.
	// Format: instances/{instance}/databases/{database}/catalog
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetDatabaseCatalogRequest) Reset() {
	*x = GetDatabaseCatalogRequest{}
	mi := &file_v1_database_catalog_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDatabaseCatalogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDatabaseCatalogRequest) ProtoMessage() {}

func (x *GetDatabaseCatalogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_database_catalog_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDatabaseCatalogRequest.ProtoReflect.Descriptor instead.
func (*GetDatabaseCatalogRequest) Descriptor() ([]byte, []int) {
	return file_v1_database_catalog_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetDatabaseCatalogRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpdateDatabaseCatalogRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The database catalog to update.
	//
	// The catalog's `name` field is used to identify the database catalog to update.
	// Format: instances/{instance}/databases/{database}/catalog
	Catalog       *DatabaseCatalog `protobuf:"bytes,1,opt,name=catalog,proto3" json:"catalog,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateDatabaseCatalogRequest) Reset() {
	*x = UpdateDatabaseCatalogRequest{}
	mi := &file_v1_database_catalog_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateDatabaseCatalogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDatabaseCatalogRequest) ProtoMessage() {}

func (x *UpdateDatabaseCatalogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_database_catalog_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDatabaseCatalogRequest.ProtoReflect.Descriptor instead.
func (*UpdateDatabaseCatalogRequest) Descriptor() ([]byte, []int) {
	return file_v1_database_catalog_service_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateDatabaseCatalogRequest) GetCatalog() *DatabaseCatalog {
	if x != nil {
		return x.Catalog
	}
	return nil
}

type DatabaseCatalog struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The name of the database catalog.
	// Format: instances/{instance}/databases/{database}/catalog
	Name          string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Schemas       []*SchemaCatalog `protobuf:"bytes,2,rep,name=schemas,proto3" json:"schemas,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DatabaseCatalog) Reset() {
	*x = DatabaseCatalog{}
	mi := &file_v1_database_catalog_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DatabaseCatalog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatabaseCatalog) ProtoMessage() {}

func (x *DatabaseCatalog) ProtoReflect() protoreflect.Message {
	mi := &file_v1_database_catalog_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatabaseCatalog.ProtoReflect.Descriptor instead.
func (*DatabaseCatalog) Descriptor() ([]byte, []int) {
	return file_v1_database_catalog_service_proto_rawDescGZIP(), []int{2}
}

func (x *DatabaseCatalog) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DatabaseCatalog) GetSchemas() []*SchemaCatalog {
	if x != nil {
		return x.Schemas
	}
	return nil
}

type SchemaCatalog struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Tables        []*TableCatalog        `protobuf:"bytes,2,rep,name=tables,proto3" json:"tables,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SchemaCatalog) Reset() {
	*x = SchemaCatalog{}
	mi := &file_v1_database_catalog_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SchemaCatalog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchemaCatalog) ProtoMessage() {}

func (x *SchemaCatalog) ProtoReflect() protoreflect.Message {
	mi := &file_v1_database_catalog_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchemaCatalog.ProtoReflect.Descriptor instead.
func (*SchemaCatalog) Descriptor() ([]byte, []int) {
	return file_v1_database_catalog_service_proto_rawDescGZIP(), []int{3}
}

func (x *SchemaCatalog) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SchemaCatalog) GetTables() []*TableCatalog {
	if x != nil {
		return x.Tables
	}
	return nil
}

type TableCatalog struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Name  string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to Kind:
	//
	//	*TableCatalog_Columns_
	//	*TableCatalog_ObjectSchema
	Kind           isTableCatalog_Kind `protobuf_oneof:"kind"`
	Classification string              `protobuf:"bytes,4,opt,name=classification,proto3" json:"classification,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *TableCatalog) Reset() {
	*x = TableCatalog{}
	mi := &file_v1_database_catalog_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TableCatalog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableCatalog) ProtoMessage() {}

func (x *TableCatalog) ProtoReflect() protoreflect.Message {
	mi := &file_v1_database_catalog_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableCatalog.ProtoReflect.Descriptor instead.
func (*TableCatalog) Descriptor() ([]byte, []int) {
	return file_v1_database_catalog_service_proto_rawDescGZIP(), []int{4}
}

func (x *TableCatalog) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TableCatalog) GetKind() isTableCatalog_Kind {
	if x != nil {
		return x.Kind
	}
	return nil
}

func (x *TableCatalog) GetColumns() *TableCatalog_Columns {
	if x != nil {
		if x, ok := x.Kind.(*TableCatalog_Columns_); ok {
			return x.Columns
		}
	}
	return nil
}

func (x *TableCatalog) GetObjectSchema() *ObjectSchema {
	if x != nil {
		if x, ok := x.Kind.(*TableCatalog_ObjectSchema); ok {
			return x.ObjectSchema
		}
	}
	return nil
}

func (x *TableCatalog) GetClassification() string {
	if x != nil {
		return x.Classification
	}
	return ""
}

type isTableCatalog_Kind interface {
	isTableCatalog_Kind()
}

type TableCatalog_Columns_ struct {
	Columns *TableCatalog_Columns `protobuf:"bytes,2,opt,name=columns,proto3,oneof"`
}

type TableCatalog_ObjectSchema struct {
	ObjectSchema *ObjectSchema `protobuf:"bytes,3,opt,name=object_schema,json=objectSchema,proto3,oneof"`
}

func (*TableCatalog_Columns_) isTableCatalog_Kind() {}

func (*TableCatalog_ObjectSchema) isTableCatalog_Kind() {}

type ColumnCatalog struct {
	state        protoimpl.MessageState `protogen:"open.v1"`
	Name         string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	SemanticType string                 `protobuf:"bytes,2,opt,name=semantic_type,json=semanticType,proto3" json:"semantic_type,omitempty"`
	// The user labels for a column.
	Labels         map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Classification string            `protobuf:"bytes,4,opt,name=classification,proto3" json:"classification,omitempty"`
	ObjectSchema   *ObjectSchema     `protobuf:"bytes,5,opt,name=object_schema,json=objectSchema,proto3,oneof" json:"object_schema,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *ColumnCatalog) Reset() {
	*x = ColumnCatalog{}
	mi := &file_v1_database_catalog_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ColumnCatalog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ColumnCatalog) ProtoMessage() {}

func (x *ColumnCatalog) ProtoReflect() protoreflect.Message {
	mi := &file_v1_database_catalog_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ColumnCatalog.ProtoReflect.Descriptor instead.
func (*ColumnCatalog) Descriptor() ([]byte, []int) {
	return file_v1_database_catalog_service_proto_rawDescGZIP(), []int{5}
}

func (x *ColumnCatalog) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ColumnCatalog) GetSemanticType() string {
	if x != nil {
		return x.SemanticType
	}
	return ""
}

func (x *ColumnCatalog) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *ColumnCatalog) GetClassification() string {
	if x != nil {
		return x.Classification
	}
	return ""
}

func (x *ColumnCatalog) GetObjectSchema() *ObjectSchema {
	if x != nil {
		return x.ObjectSchema
	}
	return nil
}

type ObjectSchema struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Type  ObjectSchema_Type      `protobuf:"varint,1,opt,name=type,proto3,enum=bytebase.v1.ObjectSchema_Type" json:"type,omitempty"`
	// Types that are valid to be assigned to Kind:
	//
	//	*ObjectSchema_StructKind_
	//	*ObjectSchema_ArrayKind_
	Kind          isObjectSchema_Kind `protobuf_oneof:"kind"`
	SemanticType  string              `protobuf:"bytes,4,opt,name=semantic_type,json=semanticType,proto3" json:"semantic_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ObjectSchema) Reset() {
	*x = ObjectSchema{}
	mi := &file_v1_database_catalog_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ObjectSchema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectSchema) ProtoMessage() {}

func (x *ObjectSchema) ProtoReflect() protoreflect.Message {
	mi := &file_v1_database_catalog_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectSchema.ProtoReflect.Descriptor instead.
func (*ObjectSchema) Descriptor() ([]byte, []int) {
	return file_v1_database_catalog_service_proto_rawDescGZIP(), []int{6}
}

func (x *ObjectSchema) GetType() ObjectSchema_Type {
	if x != nil {
		return x.Type
	}
	return ObjectSchema_TYPE_UNSPECIFIED
}

func (x *ObjectSchema) GetKind() isObjectSchema_Kind {
	if x != nil {
		return x.Kind
	}
	return nil
}

func (x *ObjectSchema) GetStructKind() *ObjectSchema_StructKind {
	if x != nil {
		if x, ok := x.Kind.(*ObjectSchema_StructKind_); ok {
			return x.StructKind
		}
	}
	return nil
}

func (x *ObjectSchema) GetArrayKind() *ObjectSchema_ArrayKind {
	if x != nil {
		if x, ok := x.Kind.(*ObjectSchema_ArrayKind_); ok {
			return x.ArrayKind
		}
	}
	return nil
}

func (x *ObjectSchema) GetSemanticType() string {
	if x != nil {
		return x.SemanticType
	}
	return ""
}

type isObjectSchema_Kind interface {
	isObjectSchema_Kind()
}

type ObjectSchema_StructKind_ struct {
	StructKind *ObjectSchema_StructKind `protobuf:"bytes,2,opt,name=struct_kind,json=structKind,proto3,oneof"`
}

type ObjectSchema_ArrayKind_ struct {
	ArrayKind *ObjectSchema_ArrayKind `protobuf:"bytes,3,opt,name=array_kind,json=arrayKind,proto3,oneof"`
}

func (*ObjectSchema_StructKind_) isObjectSchema_Kind() {}

func (*ObjectSchema_ArrayKind_) isObjectSchema_Kind() {}

type TableCatalog_Columns struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Columns       []*ColumnCatalog       `protobuf:"bytes,1,rep,name=columns,proto3" json:"columns,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TableCatalog_Columns) Reset() {
	*x = TableCatalog_Columns{}
	mi := &file_v1_database_catalog_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TableCatalog_Columns) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableCatalog_Columns) ProtoMessage() {}

func (x *TableCatalog_Columns) ProtoReflect() protoreflect.Message {
	mi := &file_v1_database_catalog_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableCatalog_Columns.ProtoReflect.Descriptor instead.
func (*TableCatalog_Columns) Descriptor() ([]byte, []int) {
	return file_v1_database_catalog_service_proto_rawDescGZIP(), []int{4, 0}
}

func (x *TableCatalog_Columns) GetColumns() []*ColumnCatalog {
	if x != nil {
		return x.Columns
	}
	return nil
}

type ObjectSchema_StructKind struct {
	state         protoimpl.MessageState   `protogen:"open.v1"`
	Properties    map[string]*ObjectSchema `protobuf:"bytes,1,rep,name=properties,proto3" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ObjectSchema_StructKind) Reset() {
	*x = ObjectSchema_StructKind{}
	mi := &file_v1_database_catalog_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ObjectSchema_StructKind) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectSchema_StructKind) ProtoMessage() {}

func (x *ObjectSchema_StructKind) ProtoReflect() protoreflect.Message {
	mi := &file_v1_database_catalog_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectSchema_StructKind.ProtoReflect.Descriptor instead.
func (*ObjectSchema_StructKind) Descriptor() ([]byte, []int) {
	return file_v1_database_catalog_service_proto_rawDescGZIP(), []int{6, 0}
}

func (x *ObjectSchema_StructKind) GetProperties() map[string]*ObjectSchema {
	if x != nil {
		return x.Properties
	}
	return nil
}

type ObjectSchema_ArrayKind struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Kind          *ObjectSchema          `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ObjectSchema_ArrayKind) Reset() {
	*x = ObjectSchema_ArrayKind{}
	mi := &file_v1_database_catalog_service_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ObjectSchema_ArrayKind) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectSchema_ArrayKind) ProtoMessage() {}

func (x *ObjectSchema_ArrayKind) ProtoReflect() protoreflect.Message {
	mi := &file_v1_database_catalog_service_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectSchema_ArrayKind.ProtoReflect.Descriptor instead.
func (*ObjectSchema_ArrayKind) Descriptor() ([]byte, []int) {
	return file_v1_database_catalog_service_proto_rawDescGZIP(), []int{6, 1}
}

func (x *ObjectSchema_ArrayKind) GetKind() *ObjectSchema {
	if x != nil {
		return x.Kind
	}
	return nil
}

var File_v1_database_catalog_service_proto protoreflect.FileDescriptor

const file_v1_database_catalog_service_proto_rawDesc = "" +
	"\n" +
	"!v1/database_catalog_service.proto\x12\vbytebase.v1\x1a\x1cgoogle/api/annotations.proto\x1a\x17google/api/client.proto\x1a\x1fgoogle/api/field_behavior.proto\x1a\x19google/api/resource.proto\x1a\x13v1/annotation.proto\"V\n" +
	"\x19GetDatabaseCatalogRequest\x129\n" +
	"\x04name\x18\x01 \x01(\tB%\xe2A\x01\x02\xfaA\x1e\n" +
	"\x1cbytebase.com/DatabaseCatalogR\x04name\"\\\n" +
	"\x1cUpdateDatabaseCatalogRequest\x12<\n" +
	"\acatalog\x18\x01 \x01(\v2\x1c.bytebase.v1.DatabaseCatalogB\x04\xe2A\x01\x02R\acatalog\"\xb1\x01\n" +
	"\x0fDatabaseCatalog\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x124\n" +
	"\aschemas\x18\x02 \x03(\v2\x1a.bytebase.v1.SchemaCatalogR\aschemas:T\xeaAQ\n" +
	"\x1cbytebase.com/DatabaseCatalog\x121instances/{instance}/databases/{database}/catalog\"V\n" +
	"\rSchemaCatalog\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x121\n" +
	"\x06tables\x18\x02 \x03(\v2\x19.bytebase.v1.TableCatalogR\x06tables\"\x94\x02\n" +
	"\fTableCatalog\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12=\n" +
	"\acolumns\x18\x02 \x01(\v2!.bytebase.v1.TableCatalog.ColumnsH\x00R\acolumns\x12@\n" +
	"\robject_schema\x18\x03 \x01(\v2\x19.bytebase.v1.ObjectSchemaH\x00R\fobjectSchema\x12&\n" +
	"\x0eclassification\x18\x04 \x01(\tR\x0eclassification\x1a?\n" +
	"\aColumns\x124\n" +
	"\acolumns\x18\x01 \x03(\v2\x1a.bytebase.v1.ColumnCatalogR\acolumnsB\x06\n" +
	"\x04kind\"\xc2\x02\n" +
	"\rColumnCatalog\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12#\n" +
	"\rsemantic_type\x18\x02 \x01(\tR\fsemanticType\x12>\n" +
	"\x06labels\x18\x03 \x03(\v2&.bytebase.v1.ColumnCatalog.LabelsEntryR\x06labels\x12&\n" +
	"\x0eclassification\x18\x04 \x01(\tR\x0eclassification\x12C\n" +
	"\robject_schema\x18\x05 \x01(\v2\x19.bytebase.v1.ObjectSchemaH\x00R\fobjectSchema\x88\x01\x01\x1a9\n" +
	"\vLabelsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01B\x10\n" +
	"\x0e_object_schema\"\xd3\x04\n" +
	"\fObjectSchema\x122\n" +
	"\x04type\x18\x01 \x01(\x0e2\x1e.bytebase.v1.ObjectSchema.TypeR\x04type\x12G\n" +
	"\vstruct_kind\x18\x02 \x01(\v2$.bytebase.v1.ObjectSchema.StructKindH\x00R\n" +
	"structKind\x12D\n" +
	"\n" +
	"array_kind\x18\x03 \x01(\v2#.bytebase.v1.ObjectSchema.ArrayKindH\x00R\tarrayKind\x12#\n" +
	"\rsemantic_type\x18\x04 \x01(\tR\fsemanticType\x1a\xbc\x01\n" +
	"\n" +
	"StructKind\x12T\n" +
	"\n" +
	"properties\x18\x01 \x03(\v24.bytebase.v1.ObjectSchema.StructKind.PropertiesEntryR\n" +
	"properties\x1aX\n" +
	"\x0fPropertiesEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12/\n" +
	"\x05value\x18\x02 \x01(\v2\x19.bytebase.v1.ObjectSchemaR\x05value:\x028\x01\x1a:\n" +
	"\tArrayKind\x12-\n" +
	"\x04kind\x18\x01 \x01(\v2\x19.bytebase.v1.ObjectSchemaR\x04kind\"X\n" +
	"\x04Type\x12\x14\n" +
	"\x10TYPE_UNSPECIFIED\x10\x00\x12\n" +
	"\n" +
	"\x06STRING\x10\x01\x12\n" +
	"\n" +
	"\x06NUMBER\x10\x02\x12\v\n" +
	"\aBOOLEAN\x10\x03\x12\n" +
	"\n" +
	"\x06OBJECT\x10\x04\x12\t\n" +
	"\x05ARRAY\x10\x05B\x06\n" +
	"\x04kind2\xb3\x03\n" +
	"\x16DatabaseCatalogService\x12\xb4\x01\n" +
	"\x12GetDatabaseCatalog\x12&.bytebase.v1.GetDatabaseCatalogRequest\x1a\x1c.bytebase.v1.DatabaseCatalog\"X\xdaA\x04name\x8a\xea0\x17bb.databaseCatalogs.get\x90\xea0\x01\x82\xd3\xe4\x93\x02,\x12*/v1/{name=instances/*/databases/*/catalog}\x12\xe1\x01\n" +
	"\x15UpdateDatabaseCatalog\x12).bytebase.v1.UpdateDatabaseCatalogRequest\x1a\x1c.bytebase.v1.DatabaseCatalog\"\x7f\xdaA\x13catalog,update_mask\x8a\xea0\x1abb.databaseCatalogs.update\x90\xea0\x01\x98\xea0\x01\x82\xd3\xe4\x93\x02=:\acatalog22/v1/{catalog.name=instances/*/databases/*/catalog}B\x11Z\x0fgenerated-go/v1b\x06proto3"

var (
	file_v1_database_catalog_service_proto_rawDescOnce sync.Once
	file_v1_database_catalog_service_proto_rawDescData []byte
)

func file_v1_database_catalog_service_proto_rawDescGZIP() []byte {
	file_v1_database_catalog_service_proto_rawDescOnce.Do(func() {
		file_v1_database_catalog_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_v1_database_catalog_service_proto_rawDesc), len(file_v1_database_catalog_service_proto_rawDesc)))
	})
	return file_v1_database_catalog_service_proto_rawDescData
}

var file_v1_database_catalog_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_database_catalog_service_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_v1_database_catalog_service_proto_goTypes = []any{
	(ObjectSchema_Type)(0),               // 0: bytebase.v1.ObjectSchema.Type
	(*GetDatabaseCatalogRequest)(nil),    // 1: bytebase.v1.GetDatabaseCatalogRequest
	(*UpdateDatabaseCatalogRequest)(nil), // 2: bytebase.v1.UpdateDatabaseCatalogRequest
	(*DatabaseCatalog)(nil),              // 3: bytebase.v1.DatabaseCatalog
	(*SchemaCatalog)(nil),                // 4: bytebase.v1.SchemaCatalog
	(*TableCatalog)(nil),                 // 5: bytebase.v1.TableCatalog
	(*ColumnCatalog)(nil),                // 6: bytebase.v1.ColumnCatalog
	(*ObjectSchema)(nil),                 // 7: bytebase.v1.ObjectSchema
	(*TableCatalog_Columns)(nil),         // 8: bytebase.v1.TableCatalog.Columns
	nil,                                  // 9: bytebase.v1.ColumnCatalog.LabelsEntry
	(*ObjectSchema_StructKind)(nil),      // 10: bytebase.v1.ObjectSchema.StructKind
	(*ObjectSchema_ArrayKind)(nil),       // 11: bytebase.v1.ObjectSchema.ArrayKind
	nil,                                  // 12: bytebase.v1.ObjectSchema.StructKind.PropertiesEntry
}
var file_v1_database_catalog_service_proto_depIdxs = []int32{
	3,  // 0: bytebase.v1.UpdateDatabaseCatalogRequest.catalog:type_name -> bytebase.v1.DatabaseCatalog
	4,  // 1: bytebase.v1.DatabaseCatalog.schemas:type_name -> bytebase.v1.SchemaCatalog
	5,  // 2: bytebase.v1.SchemaCatalog.tables:type_name -> bytebase.v1.TableCatalog
	8,  // 3: bytebase.v1.TableCatalog.columns:type_name -> bytebase.v1.TableCatalog.Columns
	7,  // 4: bytebase.v1.TableCatalog.object_schema:type_name -> bytebase.v1.ObjectSchema
	9,  // 5: bytebase.v1.ColumnCatalog.labels:type_name -> bytebase.v1.ColumnCatalog.LabelsEntry
	7,  // 6: bytebase.v1.ColumnCatalog.object_schema:type_name -> bytebase.v1.ObjectSchema
	0,  // 7: bytebase.v1.ObjectSchema.type:type_name -> bytebase.v1.ObjectSchema.Type
	10, // 8: bytebase.v1.ObjectSchema.struct_kind:type_name -> bytebase.v1.ObjectSchema.StructKind
	11, // 9: bytebase.v1.ObjectSchema.array_kind:type_name -> bytebase.v1.ObjectSchema.ArrayKind
	6,  // 10: bytebase.v1.TableCatalog.Columns.columns:type_name -> bytebase.v1.ColumnCatalog
	12, // 11: bytebase.v1.ObjectSchema.StructKind.properties:type_name -> bytebase.v1.ObjectSchema.StructKind.PropertiesEntry
	7,  // 12: bytebase.v1.ObjectSchema.ArrayKind.kind:type_name -> bytebase.v1.ObjectSchema
	7,  // 13: bytebase.v1.ObjectSchema.StructKind.PropertiesEntry.value:type_name -> bytebase.v1.ObjectSchema
	1,  // 14: bytebase.v1.DatabaseCatalogService.GetDatabaseCatalog:input_type -> bytebase.v1.GetDatabaseCatalogRequest
	2,  // 15: bytebase.v1.DatabaseCatalogService.UpdateDatabaseCatalog:input_type -> bytebase.v1.UpdateDatabaseCatalogRequest
	3,  // 16: bytebase.v1.DatabaseCatalogService.GetDatabaseCatalog:output_type -> bytebase.v1.DatabaseCatalog
	3,  // 17: bytebase.v1.DatabaseCatalogService.UpdateDatabaseCatalog:output_type -> bytebase.v1.DatabaseCatalog
	16, // [16:18] is the sub-list for method output_type
	14, // [14:16] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_v1_database_catalog_service_proto_init() }
func file_v1_database_catalog_service_proto_init() {
	if File_v1_database_catalog_service_proto != nil {
		return
	}
	file_v1_annotation_proto_init()
	file_v1_database_catalog_service_proto_msgTypes[4].OneofWrappers = []any{
		(*TableCatalog_Columns_)(nil),
		(*TableCatalog_ObjectSchema)(nil),
	}
	file_v1_database_catalog_service_proto_msgTypes[5].OneofWrappers = []any{}
	file_v1_database_catalog_service_proto_msgTypes[6].OneofWrappers = []any{
		(*ObjectSchema_StructKind_)(nil),
		(*ObjectSchema_ArrayKind_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_v1_database_catalog_service_proto_rawDesc), len(file_v1_database_catalog_service_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_database_catalog_service_proto_goTypes,
		DependencyIndexes: file_v1_database_catalog_service_proto_depIdxs,
		EnumInfos:         file_v1_database_catalog_service_proto_enumTypes,
		MessageInfos:      file_v1_database_catalog_service_proto_msgTypes,
	}.Build()
	File_v1_database_catalog_service_proto = out.File
	file_v1_database_catalog_service_proto_goTypes = nil
	file_v1_database_catalog_service_proto_depIdxs = nil
}
