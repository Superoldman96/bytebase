// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: v1/audit_log_service.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuditLog_Severity int32

const (
	AuditLog_DEFAULT   AuditLog_Severity = 0
	AuditLog_DEBUG     AuditLog_Severity = 1
	AuditLog_INFO      AuditLog_Severity = 2
	AuditLog_NOTICE    AuditLog_Severity = 3
	AuditLog_WARNING   AuditLog_Severity = 4
	AuditLog_ERROR     AuditLog_Severity = 5
	AuditLog_CRITICAL  AuditLog_Severity = 6
	AuditLog_ALERT     AuditLog_Severity = 7
	AuditLog_EMERGENCY AuditLog_Severity = 8
)

// Enum value maps for AuditLog_Severity.
var (
	AuditLog_Severity_name = map[int32]string{
		0: "DEFAULT",
		1: "DEBUG",
		2: "INFO",
		3: "NOTICE",
		4: "WARNING",
		5: "ERROR",
		6: "CRITICAL",
		7: "ALERT",
		8: "EMERGENCY",
	}
	AuditLog_Severity_value = map[string]int32{
		"DEFAULT":   0,
		"DEBUG":     1,
		"INFO":      2,
		"NOTICE":    3,
		"WARNING":   4,
		"ERROR":     5,
		"CRITICAL":  6,
		"ALERT":     7,
		"EMERGENCY": 8,
	}
)

func (x AuditLog_Severity) Enum() *AuditLog_Severity {
	p := new(AuditLog_Severity)
	*p = x
	return p
}

func (x AuditLog_Severity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuditLog_Severity) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_audit_log_service_proto_enumTypes[0].Descriptor()
}

func (AuditLog_Severity) Type() protoreflect.EnumType {
	return &file_v1_audit_log_service_proto_enumTypes[0]
}

func (x AuditLog_Severity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuditLog_Severity.Descriptor instead.
func (AuditLog_Severity) EnumDescriptor() ([]byte, []int) {
	return file_v1_audit_log_service_proto_rawDescGZIP(), []int{4, 0}
}

type SearchAuditLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parent string `protobuf:"bytes,5,opt,name=parent,proto3" json:"parent,omitempty"`
	// The filter of the log. It should be a valid CEL expression.
	// For example:
	//   - filter = "method == '/bytebase.v1.SQLService/Query'"
	//   - filter = "method == '/bytebase.v1.SQLService/Query' && severity == 'ERROR'"
	//   - filter = "method == '/bytebase.v1.SQLService/Query' && severity == 'ERROR' && user == 'users/bb@bytebase.com'"
	//   - filter = "method == '/bytebase.v1.SQLService/Query' && severity == 'ERROR' && create_time <= '2021-01-01T00:00:00Z' && create_time >= '2020-01-01T00:00:00Z'"
	Filter string `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	// The order by of the log.
	// Only support order by create_time.
	// For example:
	//   - order_by = "create_time asc"
	//   - order_by = "create_time desc"
	OrderBy string `protobuf:"bytes,2,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	// The maximum number of logs to return.
	// The service may return fewer than this value.
	// If unspecified, at most 10 log entries will be returned.
	// The maximum value is 5000; values above 5000 will be coerced to 5000.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// A page token, received from a previous `SearchLogs` call.
	// Provide this to retrieve the subsequent page.
	PageToken string `protobuf:"bytes,4,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *SearchAuditLogsRequest) Reset() {
	*x = SearchAuditLogsRequest{}
	mi := &file_v1_audit_log_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchAuditLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchAuditLogsRequest) ProtoMessage() {}

func (x *SearchAuditLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_audit_log_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchAuditLogsRequest.ProtoReflect.Descriptor instead.
func (*SearchAuditLogsRequest) Descriptor() ([]byte, []int) {
	return file_v1_audit_log_service_proto_rawDescGZIP(), []int{0}
}

func (x *SearchAuditLogsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *SearchAuditLogsRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *SearchAuditLogsRequest) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

func (x *SearchAuditLogsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *SearchAuditLogsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type SearchAuditLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuditLogs []*AuditLog `protobuf:"bytes,1,rep,name=audit_logs,json=auditLogs,proto3" json:"audit_logs,omitempty"`
	// A token to retrieve next page of log entities.
	// Pass this value in the page_token field in the subsequent call
	// to retrieve the next page of log entities.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *SearchAuditLogsResponse) Reset() {
	*x = SearchAuditLogsResponse{}
	mi := &file_v1_audit_log_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchAuditLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchAuditLogsResponse) ProtoMessage() {}

func (x *SearchAuditLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_audit_log_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchAuditLogsResponse.ProtoReflect.Descriptor instead.
func (*SearchAuditLogsResponse) Descriptor() ([]byte, []int) {
	return file_v1_audit_log_service_proto_rawDescGZIP(), []int{1}
}

func (x *SearchAuditLogsResponse) GetAuditLogs() []*AuditLog {
	if x != nil {
		return x.AuditLogs
	}
	return nil
}

func (x *SearchAuditLogsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type ExportAuditLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parent string `protobuf:"bytes,4,opt,name=parent,proto3" json:"parent,omitempty"`
	// The filter of the log. It should be a valid CEL expression.
	// For example:
	//   - filter = "method == '/bytebase.v1.SQLService/Query'"
	//   - filter = "method == '/bytebase.v1.SQLService/Query' && severity == 'ERROR'"
	//   - filter = "method == '/bytebase.v1.SQLService/Query' && severity == 'ERROR' && user == 'users/bb@bytebase.com'"
	//   - filter = "method == '/bytebase.v1.SQLService/Query' && severity == 'ERROR' && create_time <= '2021-01-01T00:00:00Z' && create_time >= '2020-01-01T00:00:00Z'"
	Filter string `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	// The order by of the log.
	// Only support order by create_time.
	// For example:
	//   - order_by = "create_time asc"
	//   - order_by = "create_time desc"
	OrderBy string `protobuf:"bytes,2,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	// The export format.
	Format ExportFormat `protobuf:"varint,3,opt,name=format,proto3,enum=bytebase.v1.ExportFormat" json:"format,omitempty"`
	// The maximum number of logs to return.
	// The service may return fewer than this value.
	// If unspecified, at most 10 log entries will be returned.
	// The maximum value is 5000; values above 5000 will be coerced to 5000.
	PageSize int32 `protobuf:"varint,5,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// A page token, received from a previous `ExportAuditLogs` call.
	// Provide this to retrieve the subsequent page.
	PageToken string `protobuf:"bytes,6,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ExportAuditLogsRequest) Reset() {
	*x = ExportAuditLogsRequest{}
	mi := &file_v1_audit_log_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExportAuditLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportAuditLogsRequest) ProtoMessage() {}

func (x *ExportAuditLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_audit_log_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportAuditLogsRequest.ProtoReflect.Descriptor instead.
func (*ExportAuditLogsRequest) Descriptor() ([]byte, []int) {
	return file_v1_audit_log_service_proto_rawDescGZIP(), []int{2}
}

func (x *ExportAuditLogsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ExportAuditLogsRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *ExportAuditLogsRequest) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

func (x *ExportAuditLogsRequest) GetFormat() ExportFormat {
	if x != nil {
		return x.Format
	}
	return ExportFormat_FORMAT_UNSPECIFIED
}

func (x *ExportAuditLogsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ExportAuditLogsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ExportAuditLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	// A token to retrieve next page of log entities.
	// Pass this value in the page_token field in the subsequent call
	// to retrieve the next page of log entities.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ExportAuditLogsResponse) Reset() {
	*x = ExportAuditLogsResponse{}
	mi := &file_v1_audit_log_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExportAuditLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportAuditLogsResponse) ProtoMessage() {}

func (x *ExportAuditLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_audit_log_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportAuditLogsResponse.ProtoReflect.Descriptor instead.
func (*ExportAuditLogsResponse) Descriptor() ([]byte, []int) {
	return file_v1_audit_log_service_proto_rawDescGZIP(), []int{3}
}

func (x *ExportAuditLogsResponse) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *ExportAuditLogsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type AuditLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the log.
	// Formats:
	// - projects/{project}/auditLogs/{uid}
	// - workspaces/{workspace}/auditLogs/{uid}
	Name       string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Format: users/d@d.com
	User string `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	// e.g. `/bytebase.v1.SQLService/Query`, `bb.project.repository.push`
	Method   string            `protobuf:"bytes,4,opt,name=method,proto3" json:"method,omitempty"`
	Severity AuditLog_Severity `protobuf:"varint,5,opt,name=severity,proto3,enum=bytebase.v1.AuditLog_Severity" json:"severity,omitempty"`
	// The associated resource.
	Resource string `protobuf:"bytes,6,opt,name=resource,proto3" json:"resource,omitempty"`
	// JSON-encoded request.
	Request string `protobuf:"bytes,7,opt,name=request,proto3" json:"request,omitempty"`
	// JSON-encoded response.
	// Some fields are omitted because they are too large or contain sensitive information.
	Response string         `protobuf:"bytes,8,opt,name=response,proto3" json:"response,omitempty"`
	Status   *status.Status `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
	// service-specific data about the request, response, and other activities.
	ServiceData *anypb.Any `protobuf:"bytes,10,opt,name=service_data,json=serviceData,proto3" json:"service_data,omitempty"`
	// Metadata about the operation.
	RequestMetadata *RequestMetadata `protobuf:"bytes,11,opt,name=request_metadata,json=requestMetadata,proto3" json:"request_metadata,omitempty"`
}

func (x *AuditLog) Reset() {
	*x = AuditLog{}
	mi := &file_v1_audit_log_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuditLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditLog) ProtoMessage() {}

func (x *AuditLog) ProtoReflect() protoreflect.Message {
	mi := &file_v1_audit_log_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditLog.ProtoReflect.Descriptor instead.
func (*AuditLog) Descriptor() ([]byte, []int) {
	return file_v1_audit_log_service_proto_rawDescGZIP(), []int{4}
}

func (x *AuditLog) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AuditLog) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *AuditLog) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *AuditLog) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *AuditLog) GetSeverity() AuditLog_Severity {
	if x != nil {
		return x.Severity
	}
	return AuditLog_DEFAULT
}

func (x *AuditLog) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *AuditLog) GetRequest() string {
	if x != nil {
		return x.Request
	}
	return ""
}

func (x *AuditLog) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

func (x *AuditLog) GetStatus() *status.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *AuditLog) GetServiceData() *anypb.Any {
	if x != nil {
		return x.ServiceData
	}
	return nil
}

func (x *AuditLog) GetRequestMetadata() *RequestMetadata {
	if x != nil {
		return x.RequestMetadata
	}
	return nil
}

type AuditData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PolicyDelta *PolicyDelta `protobuf:"bytes,1,opt,name=policy_delta,json=policyDelta,proto3" json:"policy_delta,omitempty"`
}

func (x *AuditData) Reset() {
	*x = AuditData{}
	mi := &file_v1_audit_log_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuditData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditData) ProtoMessage() {}

func (x *AuditData) ProtoReflect() protoreflect.Message {
	mi := &file_v1_audit_log_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditData.ProtoReflect.Descriptor instead.
func (*AuditData) Descriptor() ([]byte, []int) {
	return file_v1_audit_log_service_proto_rawDescGZIP(), []int{5}
}

func (x *AuditData) GetPolicyDelta() *PolicyDelta {
	if x != nil {
		return x.PolicyDelta
	}
	return nil
}

// Metadata about the request.
type RequestMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The IP address of the caller.
	CallerIp string `protobuf:"bytes,1,opt,name=caller_ip,json=callerIp,proto3" json:"caller_ip,omitempty"`
	// The user agent of the caller.
	// This information is not authenticated and should be treated accordingly.
	CallerSuppliedUserAgent string `protobuf:"bytes,2,opt,name=caller_supplied_user_agent,json=callerSuppliedUserAgent,proto3" json:"caller_supplied_user_agent,omitempty"`
}

func (x *RequestMetadata) Reset() {
	*x = RequestMetadata{}
	mi := &file_v1_audit_log_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequestMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestMetadata) ProtoMessage() {}

func (x *RequestMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_v1_audit_log_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestMetadata.ProtoReflect.Descriptor instead.
func (*RequestMetadata) Descriptor() ([]byte, []int) {
	return file_v1_audit_log_service_proto_rawDescGZIP(), []int{6}
}

func (x *RequestMetadata) GetCallerIp() string {
	if x != nil {
		return x.CallerIp
	}
	return ""
}

func (x *RequestMetadata) GetCallerSuppliedUserAgent() string {
	if x != nil {
		return x.CallerSuppliedUserAgent
	}
	return ""
}

var File_v1_audit_log_service_proto protoreflect.FileDescriptor

var file_v1_audit_log_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x62, 0x79,
	0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13,
	0x76, 0x31, 0x2f, 0x69, 0x61, 0x6d, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x01, 0x0a, 0x16, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x75,
	0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36,
	0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1e,
	0xe2, 0x41, 0x01, 0x02, 0xfa, 0x41, 0x17, 0x12, 0x15, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x06,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x19,
	0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x77, 0x0a, 0x17, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41,
	0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x34, 0x0a, 0x0a, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x09, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xf2,
	0x01, 0x0a, 0x16, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f,
	0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x06, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1e, 0xe2, 0x41, 0x01, 0x02, 0xfa,
	0x41, 0x17, 0x12, 0x15, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x42, 0x79, 0x12, 0x31, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52,
	0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x5b, 0x0a, 0x17, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x41, 0x75, 0x64,
	0x69, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0xc9, 0x04, 0x0a, 0x08, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x12, 0x18, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01,
	0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x3a, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69,
	0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x2e,
	0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69,
	0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x37, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x47, 0x0a, 0x10, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x78, 0x0a, 0x08, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x0b,
	0x0a, 0x07, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x44,
	0x45, 0x42, 0x55, 0x47, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x02,
	0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x54, 0x49, 0x43, 0x45, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07,
	0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52, 0x49, 0x54, 0x49, 0x43, 0x41, 0x4c,
	0x10, 0x06, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x4c, 0x45, 0x52, 0x54, 0x10, 0x07, 0x12, 0x0d, 0x0a,
	0x09, 0x45, 0x4d, 0x45, 0x52, 0x47, 0x45, 0x4e, 0x43, 0x59, 0x10, 0x08, 0x22, 0x48, 0x0a, 0x09,
	0x41, 0x75, 0x64, 0x69, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x3b, 0x0a, 0x0c, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x5f, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x52, 0x0b, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x22, 0x6b, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x61, 0x6c,
	0x6c, 0x65, 0x72, 0x5f, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61,
	0x6c, 0x6c, 0x65, 0x72, 0x49, 0x70, 0x12, 0x3b, 0x0a, 0x1a, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72,
	0x5f, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x63, 0x61, 0x6c, 0x6c,
	0x65, 0x72, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x32, 0xa5, 0x03, 0x0a, 0x0f, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xc7, 0x01, 0x0a, 0x0f, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x23, 0x2e, 0x62, 0x79,
	0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x69, 0x8a, 0xea, 0x30, 0x13, 0x62, 0x62, 0x2e, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x90,
	0xea, 0x30, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x48, 0x3a, 0x01, 0x2a, 0x5a, 0x19, 0x3a, 0x01,
	0x2a, 0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x73,
	0x3a, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x28, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x7d,
	0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x3a, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x12, 0xc7, 0x01, 0x0a, 0x0f, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x41, 0x75, 0x64, 0x69,
	0x74, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x23, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c,
	0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x62, 0x79, 0x74,
	0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x41,
	0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x69, 0x8a, 0xea, 0x30, 0x13, 0x62, 0x62, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f,
	0x67, 0x73, 0x2e, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x90, 0xea, 0x30, 0x01, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x48, 0x3a, 0x01, 0x2a, 0x5a, 0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x3a, 0x65, 0x78, 0x70, 0x6f, 0x72,
	0x74, 0x22, 0x28, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x4c, 0x6f, 0x67, 0x73, 0x3a, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x11, 0x5a, 0x0f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_audit_log_service_proto_rawDescOnce sync.Once
	file_v1_audit_log_service_proto_rawDescData = file_v1_audit_log_service_proto_rawDesc
)

func file_v1_audit_log_service_proto_rawDescGZIP() []byte {
	file_v1_audit_log_service_proto_rawDescOnce.Do(func() {
		file_v1_audit_log_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_audit_log_service_proto_rawDescData)
	})
	return file_v1_audit_log_service_proto_rawDescData
}

var file_v1_audit_log_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_audit_log_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_v1_audit_log_service_proto_goTypes = []any{
	(AuditLog_Severity)(0),          // 0: bytebase.v1.AuditLog.Severity
	(*SearchAuditLogsRequest)(nil),  // 1: bytebase.v1.SearchAuditLogsRequest
	(*SearchAuditLogsResponse)(nil), // 2: bytebase.v1.SearchAuditLogsResponse
	(*ExportAuditLogsRequest)(nil),  // 3: bytebase.v1.ExportAuditLogsRequest
	(*ExportAuditLogsResponse)(nil), // 4: bytebase.v1.ExportAuditLogsResponse
	(*AuditLog)(nil),                // 5: bytebase.v1.AuditLog
	(*AuditData)(nil),               // 6: bytebase.v1.AuditData
	(*RequestMetadata)(nil),         // 7: bytebase.v1.RequestMetadata
	(ExportFormat)(0),               // 8: bytebase.v1.ExportFormat
	(*timestamppb.Timestamp)(nil),   // 9: google.protobuf.Timestamp
	(*status.Status)(nil),           // 10: google.rpc.Status
	(*anypb.Any)(nil),               // 11: google.protobuf.Any
	(*PolicyDelta)(nil),             // 12: bytebase.v1.PolicyDelta
}
var file_v1_audit_log_service_proto_depIdxs = []int32{
	5,  // 0: bytebase.v1.SearchAuditLogsResponse.audit_logs:type_name -> bytebase.v1.AuditLog
	8,  // 1: bytebase.v1.ExportAuditLogsRequest.format:type_name -> bytebase.v1.ExportFormat
	9,  // 2: bytebase.v1.AuditLog.create_time:type_name -> google.protobuf.Timestamp
	0,  // 3: bytebase.v1.AuditLog.severity:type_name -> bytebase.v1.AuditLog.Severity
	10, // 4: bytebase.v1.AuditLog.status:type_name -> google.rpc.Status
	11, // 5: bytebase.v1.AuditLog.service_data:type_name -> google.protobuf.Any
	7,  // 6: bytebase.v1.AuditLog.request_metadata:type_name -> bytebase.v1.RequestMetadata
	12, // 7: bytebase.v1.AuditData.policy_delta:type_name -> bytebase.v1.PolicyDelta
	1,  // 8: bytebase.v1.AuditLogService.SearchAuditLogs:input_type -> bytebase.v1.SearchAuditLogsRequest
	3,  // 9: bytebase.v1.AuditLogService.ExportAuditLogs:input_type -> bytebase.v1.ExportAuditLogsRequest
	2,  // 10: bytebase.v1.AuditLogService.SearchAuditLogs:output_type -> bytebase.v1.SearchAuditLogsResponse
	4,  // 11: bytebase.v1.AuditLogService.ExportAuditLogs:output_type -> bytebase.v1.ExportAuditLogsResponse
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_v1_audit_log_service_proto_init() }
func file_v1_audit_log_service_proto_init() {
	if File_v1_audit_log_service_proto != nil {
		return
	}
	file_v1_annotation_proto_init()
	file_v1_common_proto_init()
	file_v1_iam_policy_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_audit_log_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_audit_log_service_proto_goTypes,
		DependencyIndexes: file_v1_audit_log_service_proto_depIdxs,
		EnumInfos:         file_v1_audit_log_service_proto_enumTypes,
		MessageInfos:      file_v1_audit_log_service_proto_msgTypes,
	}.Build()
	File_v1_audit_log_service_proto = out.File
	file_v1_audit_log_service_proto_rawDesc = nil
	file_v1_audit_log_service_proto_goTypes = nil
	file_v1_audit_log_service_proto_depIdxs = nil
}
