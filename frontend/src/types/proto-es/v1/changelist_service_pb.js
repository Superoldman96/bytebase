// @generated by protoc-gen-es v2.5.2
// @generated from file v1/changelist_service.proto (package bytebase.v1, syntax proto3)
/* eslint-disable */

import { fileDesc, messageDesc, serviceDesc } from "@bufbuild/protobuf/codegenv2";
import { file_google_api_annotations } from "../google/api/annotations_pb";
import { file_google_api_client } from "../google/api/client_pb";
import { file_google_api_field_behavior } from "../google/api/field_behavior_pb";
import { file_google_api_resource } from "../google/api/resource_pb";
import { file_google_protobuf_empty, file_google_protobuf_field_mask, file_google_protobuf_timestamp } from "@bufbuild/protobuf/wkt";
import { file_v1_annotation } from "./annotation_pb";

/**
 * Describes the file v1/changelist_service.proto.
 */
export const file_v1_changelist_service = /*@__PURE__*/
  fileDesc("Cht2MS9jaGFuZ2VsaXN0X3NlcnZpY2UucHJvdG8SC2J5dGViYXNlLnYxIpUBChdDcmVhdGVDaGFuZ2VsaXN0UmVxdWVzdBIsCgZwYXJlbnQYASABKAlCHOBBAvpBFgoUYnl0ZWJhc2UuY29tL1Byb2plY3QSMAoKY2hhbmdlbGlzdBgCIAEoCzIXLmJ5dGViYXNlLnYxLkNoYW5nZWxpc3RCA+BBAhIaCg1jaGFuZ2VsaXN0X2lkGAMgASgJQgPgQQIiRQoUR2V0Q2hhbmdlbGlzdFJlcXVlc3QSLQoEbmFtZRgBIAEoCUIf4EEC+kEZChdieXRlYmFzZS5jb20vQ2hhbmdlbGlzdCJtChZMaXN0Q2hhbmdlbGlzdHNSZXF1ZXN0EiwKBnBhcmVudBgBIAEoCUIc4EEC+kEWChRieXRlYmFzZS5jb20vUHJvamVjdBIRCglwYWdlX3NpemUYAiABKAUSEgoKcGFnZV90b2tlbhgDIAEoCSJgChdMaXN0Q2hhbmdlbGlzdHNSZXNwb25zZRIsCgtjaGFuZ2VsaXN0cxgBIAMoCzIXLmJ5dGViYXNlLnYxLkNoYW5nZWxpc3QSFwoPbmV4dF9wYWdlX3Rva2VuGAIgASgJInwKF1VwZGF0ZUNoYW5nZWxpc3RSZXF1ZXN0EjAKCmNoYW5nZWxpc3QYASABKAsyFy5ieXRlYmFzZS52MS5DaGFuZ2VsaXN0QgPgQQISLwoLdXBkYXRlX21hc2sYAiABKAsyGi5nb29nbGUucHJvdG9idWYuRmllbGRNYXNrIkgKF0RlbGV0ZUNoYW5nZWxpc3RSZXF1ZXN0Ei0KBG5hbWUYASABKAlCH+BBAvpBGQoXYnl0ZWJhc2UuY29tL0NoYW5nZWxpc3QiqAIKCkNoYW5nZWxpc3QSFAoEbmFtZRgBIAEoCUIG4EEC4EEFEhMKC2Rlc2NyaXB0aW9uGAIgASgJEhQKB2NyZWF0b3IYAyABKAlCA+BBAxI0Cgt1cGRhdGVfdGltZRgGIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5UaW1lc3RhbXBCA+BBAxIvCgdjaGFuZ2VzGAcgAygLMh4uYnl0ZWJhc2UudjEuQ2hhbmdlbGlzdC5DaGFuZ2UaJwoGQ2hhbmdlEg0KBXNoZWV0GAEgASgJEg4KBnNvdXJjZRgCIAEoCTpJ6kFGChdieXRlYmFzZS5jb20vQ2hhbmdlbGlzdBIrcHJvamVjdHMve3Byb2plY3R9L2NoYW5nZWxpc3RzL3tjaGFuZ2VsaXN0fTKPBwoRQ2hhbmdlbGlzdFNlcnZpY2USuwEKEENyZWF0ZUNoYW5nZWxpc3QSJC5ieXRlYmFzZS52MS5DcmVhdGVDaGFuZ2VsaXN0UmVxdWVzdBoXLmJ5dGViYXNlLnYxLkNoYW5nZWxpc3QiaNpBEXBhcmVudCxjaGFuZ2VsaXN0iuowFWJiLmNoYW5nZWxpc3RzLmNyZWF0ZZDqMAGC0+STAjE6CmNoYW5nZWxpc3QiIy92MS97cGFyZW50PXByb2plY3RzLyp9L2NoYW5nZWxpc3RzEpkBCg1HZXRDaGFuZ2VsaXN0EiEuYnl0ZWJhc2UudjEuR2V0Q2hhbmdlbGlzdFJlcXVlc3QaFy5ieXRlYmFzZS52MS5DaGFuZ2VsaXN0IkzaQQRuYW1liuowEmJiLmNoYW5nZWxpc3RzLmdldJDqMAGC0+STAiUSIy92MS97bmFtZT1wcm9qZWN0cy8qL2NoYW5nZWxpc3RzLyp9Eq0BCg9MaXN0Q2hhbmdlbGlzdHMSIy5ieXRlYmFzZS52MS5MaXN0Q2hhbmdlbGlzdHNSZXF1ZXN0GiQuYnl0ZWJhc2UudjEuTGlzdENoYW5nZWxpc3RzUmVzcG9uc2UiT9pBBnBhcmVudIrqMBNiYi5jaGFuZ2VsaXN0cy5saXN0kOowAYLT5JMCJRIjL3YxL3twYXJlbnQ9cHJvamVjdHMvKn0vY2hhbmdlbGlzdHMSywEKEFVwZGF0ZUNoYW5nZWxpc3QSJC5ieXRlYmFzZS52MS5VcGRhdGVDaGFuZ2VsaXN0UmVxdWVzdBoXLmJ5dGViYXNlLnYxLkNoYW5nZWxpc3QieNpBFmNoYW5nZWxpc3QsdXBkYXRlX21hc2uK6jAVYmIuY2hhbmdlbGlzdHMudXBkYXRlkOowAYLT5JMCPDoKY2hhbmdlbGlzdDIuL3YxL3tjaGFuZ2VsaXN0Lm5hbWU9cHJvamVjdHMvKi9jaGFuZ2VsaXN0cy8qfRKhAQoQRGVsZXRlQ2hhbmdlbGlzdBIkLmJ5dGViYXNlLnYxLkRlbGV0ZUNoYW5nZWxpc3RSZXF1ZXN0GhYuZ29vZ2xlLnByb3RvYnVmLkVtcHR5Ik/aQQRuYW1liuowFWJiLmNoYW5nZWxpc3RzLmRlbGV0ZZDqMAGC0+STAiUqIy92MS97bmFtZT1wcm9qZWN0cy8qL2NoYW5nZWxpc3RzLyp9QjZaNGdpdGh1Yi5jb20vYnl0ZWJhc2UvYnl0ZWJhc2UvYmFja2VuZC9nZW5lcmF0ZWQtZ28vdjFiBnByb3RvMw", [file_google_api_annotations, file_google_api_client, file_google_api_field_behavior, file_google_api_resource, file_google_protobuf_empty, file_google_protobuf_field_mask, file_google_protobuf_timestamp, file_v1_annotation]);

/**
 * Describes the message bytebase.v1.CreateChangelistRequest.
 * Use `create(CreateChangelistRequestSchema)` to create a new message.
 */
export const CreateChangelistRequestSchema = /*@__PURE__*/
  messageDesc(file_v1_changelist_service, 0);

/**
 * Describes the message bytebase.v1.GetChangelistRequest.
 * Use `create(GetChangelistRequestSchema)` to create a new message.
 */
export const GetChangelistRequestSchema = /*@__PURE__*/
  messageDesc(file_v1_changelist_service, 1);

/**
 * Describes the message bytebase.v1.ListChangelistsRequest.
 * Use `create(ListChangelistsRequestSchema)` to create a new message.
 */
export const ListChangelistsRequestSchema = /*@__PURE__*/
  messageDesc(file_v1_changelist_service, 2);

/**
 * Describes the message bytebase.v1.ListChangelistsResponse.
 * Use `create(ListChangelistsResponseSchema)` to create a new message.
 */
export const ListChangelistsResponseSchema = /*@__PURE__*/
  messageDesc(file_v1_changelist_service, 3);

/**
 * Describes the message bytebase.v1.UpdateChangelistRequest.
 * Use `create(UpdateChangelistRequestSchema)` to create a new message.
 */
export const UpdateChangelistRequestSchema = /*@__PURE__*/
  messageDesc(file_v1_changelist_service, 4);

/**
 * Describes the message bytebase.v1.DeleteChangelistRequest.
 * Use `create(DeleteChangelistRequestSchema)` to create a new message.
 */
export const DeleteChangelistRequestSchema = /*@__PURE__*/
  messageDesc(file_v1_changelist_service, 5);

/**
 * Describes the message bytebase.v1.Changelist.
 * Use `create(ChangelistSchema)` to create a new message.
 */
export const ChangelistSchema = /*@__PURE__*/
  messageDesc(file_v1_changelist_service, 6);

/**
 * Describes the message bytebase.v1.Changelist.Change.
 * Use `create(Changelist_ChangeSchema)` to create a new message.
 */
export const Changelist_ChangeSchema = /*@__PURE__*/
  messageDesc(file_v1_changelist_service, 6, 0);

/**
 * @generated from service bytebase.v1.ChangelistService
 */
export const ChangelistService = /*@__PURE__*/
  serviceDesc(file_v1_changelist_service, 0);

