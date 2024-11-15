// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: v1/cel_service.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	v1alpha1 "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
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

type BatchParseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Expressions []string `protobuf:"bytes,1,rep,name=expressions,proto3" json:"expressions,omitempty"`
}

func (x *BatchParseRequest) Reset() {
	*x = BatchParseRequest{}
	mi := &file_v1_cel_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BatchParseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchParseRequest) ProtoMessage() {}

func (x *BatchParseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_cel_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchParseRequest.ProtoReflect.Descriptor instead.
func (*BatchParseRequest) Descriptor() ([]byte, []int) {
	return file_v1_cel_service_proto_rawDescGZIP(), []int{0}
}

func (x *BatchParseRequest) GetExpressions() []string {
	if x != nil {
		return x.Expressions
	}
	return nil
}

type BatchParseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Expressions []*v1alpha1.Expr `protobuf:"bytes,1,rep,name=expressions,proto3" json:"expressions,omitempty"`
}

func (x *BatchParseResponse) Reset() {
	*x = BatchParseResponse{}
	mi := &file_v1_cel_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BatchParseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchParseResponse) ProtoMessage() {}

func (x *BatchParseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_cel_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchParseResponse.ProtoReflect.Descriptor instead.
func (*BatchParseResponse) Descriptor() ([]byte, []int) {
	return file_v1_cel_service_proto_rawDescGZIP(), []int{1}
}

func (x *BatchParseResponse) GetExpressions() []*v1alpha1.Expr {
	if x != nil {
		return x.Expressions
	}
	return nil
}

type BatchDeparseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Expressions []*v1alpha1.Expr `protobuf:"bytes,1,rep,name=expressions,proto3" json:"expressions,omitempty"`
}

func (x *BatchDeparseRequest) Reset() {
	*x = BatchDeparseRequest{}
	mi := &file_v1_cel_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BatchDeparseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchDeparseRequest) ProtoMessage() {}

func (x *BatchDeparseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_cel_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchDeparseRequest.ProtoReflect.Descriptor instead.
func (*BatchDeparseRequest) Descriptor() ([]byte, []int) {
	return file_v1_cel_service_proto_rawDescGZIP(), []int{2}
}

func (x *BatchDeparseRequest) GetExpressions() []*v1alpha1.Expr {
	if x != nil {
		return x.Expressions
	}
	return nil
}

type BatchDeparseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Expressions []string `protobuf:"bytes,1,rep,name=expressions,proto3" json:"expressions,omitempty"`
}

func (x *BatchDeparseResponse) Reset() {
	*x = BatchDeparseResponse{}
	mi := &file_v1_cel_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BatchDeparseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchDeparseResponse) ProtoMessage() {}

func (x *BatchDeparseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_cel_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchDeparseResponse.ProtoReflect.Descriptor instead.
func (*BatchDeparseResponse) Descriptor() ([]byte, []int) {
	return file_v1_cel_service_proto_rawDescGZIP(), []int{3}
}

func (x *BatchDeparseResponse) GetExpressions() []string {
	if x != nil {
		return x.Expressions
	}
	return nil
}

var File_v1_cel_service_proto protoreflect.FileDescriptor

var file_v1_cel_service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x76, 0x31, 0x2f, 0x63, 0x65, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x25, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78,
	0x70, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x79, 0x6e, 0x74,
	0x61, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a,
	0x11, 0x42, 0x61, 0x74, 0x63, 0x68, 0x50, 0x61, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x22, 0x56, 0x0a, 0x12, 0x42, 0x61, 0x74, 0x63, 0x68, 0x50, 0x61, 0x72,
	0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x65, 0x78,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x78, 0x70,
	0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x52,
	0x0b, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x57, 0x0a, 0x13,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x70, 0x61, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x52, 0x0b, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x38, 0x0a, 0x14, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65,
	0x70, 0x61, 0x72, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0b, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x32,
	0xf8, 0x01, 0x0a, 0x0a, 0x43, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x70,
	0x0a, 0x0a, 0x42, 0x61, 0x74, 0x63, 0x68, 0x50, 0x61, 0x72, 0x73, 0x65, 0x12, 0x1e, 0x2e, 0x62,
	0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x50, 0x61, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x62,
	0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x50, 0x61, 0x72, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x80,
	0xea, 0x30, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x65, 0x6c, 0x2f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x50, 0x61, 0x72, 0x73, 0x65,
	0x12, 0x78, 0x0a, 0x0c, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x70, 0x61, 0x72, 0x73, 0x65,
	0x12, 0x20, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x70, 0x61, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x70, 0x61, 0x72, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x80, 0xea, 0x30, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x65, 0x6c, 0x2f, 0x62, 0x61,
	0x74, 0x63, 0x68, 0x44, 0x65, 0x70, 0x61, 0x72, 0x73, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_cel_service_proto_rawDescOnce sync.Once
	file_v1_cel_service_proto_rawDescData = file_v1_cel_service_proto_rawDesc
)

func file_v1_cel_service_proto_rawDescGZIP() []byte {
	file_v1_cel_service_proto_rawDescOnce.Do(func() {
		file_v1_cel_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_cel_service_proto_rawDescData)
	})
	return file_v1_cel_service_proto_rawDescData
}

var file_v1_cel_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_v1_cel_service_proto_goTypes = []any{
	(*BatchParseRequest)(nil),    // 0: bytebase.v1.BatchParseRequest
	(*BatchParseResponse)(nil),   // 1: bytebase.v1.BatchParseResponse
	(*BatchDeparseRequest)(nil),  // 2: bytebase.v1.BatchDeparseRequest
	(*BatchDeparseResponse)(nil), // 3: bytebase.v1.BatchDeparseResponse
	(*v1alpha1.Expr)(nil),        // 4: google.api.expr.v1alpha1.Expr
}
var file_v1_cel_service_proto_depIdxs = []int32{
	4, // 0: bytebase.v1.BatchParseResponse.expressions:type_name -> google.api.expr.v1alpha1.Expr
	4, // 1: bytebase.v1.BatchDeparseRequest.expressions:type_name -> google.api.expr.v1alpha1.Expr
	0, // 2: bytebase.v1.CelService.BatchParse:input_type -> bytebase.v1.BatchParseRequest
	2, // 3: bytebase.v1.CelService.BatchDeparse:input_type -> bytebase.v1.BatchDeparseRequest
	1, // 4: bytebase.v1.CelService.BatchParse:output_type -> bytebase.v1.BatchParseResponse
	3, // 5: bytebase.v1.CelService.BatchDeparse:output_type -> bytebase.v1.BatchDeparseResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_v1_cel_service_proto_init() }
func file_v1_cel_service_proto_init() {
	if File_v1_cel_service_proto != nil {
		return
	}
	file_v1_annotation_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_cel_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_cel_service_proto_goTypes,
		DependencyIndexes: file_v1_cel_service_proto_depIdxs,
		MessageInfos:      file_v1_cel_service_proto_msgTypes,
	}.Build()
	File_v1_cel_service_proto = out.File
	file_v1_cel_service_proto_rawDesc = nil
	file_v1_cel_service_proto_goTypes = nil
	file_v1_cel_service_proto_depIdxs = nil
}
