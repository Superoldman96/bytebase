// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: store/changelist.proto

package store

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

type Changelist struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string               `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	Changes     []*Changelist_Change `protobuf:"bytes,2,rep,name=changes,proto3" json:"changes,omitempty"`
}

func (x *Changelist) Reset() {
	*x = Changelist{}
	mi := &file_store_changelist_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Changelist) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Changelist) ProtoMessage() {}

func (x *Changelist) ProtoReflect() protoreflect.Message {
	mi := &file_store_changelist_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Changelist.ProtoReflect.Descriptor instead.
func (*Changelist) Descriptor() ([]byte, []int) {
	return file_store_changelist_proto_rawDescGZIP(), []int{0}
}

func (x *Changelist) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Changelist) GetChanges() []*Changelist_Change {
	if x != nil {
		return x.Changes
	}
	return nil
}

type Changelist_Change struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of a sheet.
	Sheet string `protobuf:"bytes,1,opt,name=sheet,proto3" json:"sheet,omitempty"`
	// The source of origin.
	// 1) change history: instances/{instance}/databases/{database}/changeHistories/{changeHistory}.
	// 2) branch: projects/{project}/branches/{branch}.
	// 3) raw SQL if empty.
	Source string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	// The migration version for a change.
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Changelist_Change) Reset() {
	*x = Changelist_Change{}
	mi := &file_store_changelist_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Changelist_Change) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Changelist_Change) ProtoMessage() {}

func (x *Changelist_Change) ProtoReflect() protoreflect.Message {
	mi := &file_store_changelist_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Changelist_Change.ProtoReflect.Descriptor instead.
func (*Changelist_Change) Descriptor() ([]byte, []int) {
	return file_store_changelist_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Changelist_Change) GetSheet() string {
	if x != nil {
		return x.Sheet
	}
	return ""
}

func (x *Changelist_Change) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Changelist_Change) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_store_changelist_proto protoreflect.FileDescriptor

var file_store_changelist_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x69,
	0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x22, 0xbd, 0x01, 0x0a, 0x0a, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x07, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x62, 0x79, 0x74,
	0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x07, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x1a, 0x50, 0x0a, 0x06, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x65, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x73, 0x68, 0x65, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x14, 0x5a, 0x12, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_store_changelist_proto_rawDescOnce sync.Once
	file_store_changelist_proto_rawDescData = file_store_changelist_proto_rawDesc
)

func file_store_changelist_proto_rawDescGZIP() []byte {
	file_store_changelist_proto_rawDescOnce.Do(func() {
		file_store_changelist_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_changelist_proto_rawDescData)
	})
	return file_store_changelist_proto_rawDescData
}

var file_store_changelist_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_store_changelist_proto_goTypes = []any{
	(*Changelist)(nil),        // 0: bytebase.store.Changelist
	(*Changelist_Change)(nil), // 1: bytebase.store.Changelist.Change
}
var file_store_changelist_proto_depIdxs = []int32{
	1, // 0: bytebase.store.Changelist.changes:type_name -> bytebase.store.Changelist.Change
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_store_changelist_proto_init() }
func file_store_changelist_proto_init() {
	if File_store_changelist_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_store_changelist_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_changelist_proto_goTypes,
		DependencyIndexes: file_store_changelist_proto_depIdxs,
		MessageInfos:      file_store_changelist_proto_msgTypes,
	}.Build()
	File_store_changelist_proto = out.File
	file_store_changelist_proto_rawDesc = nil
	file_store_changelist_proto_goTypes = nil
	file_store_changelist_proto_depIdxs = nil
}
