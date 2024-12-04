// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.2.0
//   protoc               unknown
// source: store/changelog.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import Long from "long";
import { ChangedResources } from "./instance_change_history";

export const protobufPackage = "bytebase.store";

export interface ChangelogPayload {
  /** Format: projects/{project}/rollouts/{rollout}/stages/{stage}/tasks/{task}/taskruns/{taskrun} */
  taskRun: string;
  /** Format: projects/{project}/issues/{issue} */
  issue: string;
  /**
   * The revision uid.
   * optional
   */
  revision: Long;
  changedResources:
    | ChangedResources
    | undefined;
  /**
   * The sheet that holds the content.
   * Format: projects/{project}/sheets/{sheet}
   */
  sheet: string;
  version: string;
  type: ChangelogPayload_Type;
}

export enum ChangelogPayload_Type {
  TYPE_UNSPECIFIED = "TYPE_UNSPECIFIED",
  BASELINE = "BASELINE",
  MIGRATE = "MIGRATE",
  MIGRATE_SDL = "MIGRATE_SDL",
  MIGRATE_GHOST = "MIGRATE_GHOST",
  DATA = "DATA",
  UNRECOGNIZED = "UNRECOGNIZED",
}

export function changelogPayload_TypeFromJSON(object: any): ChangelogPayload_Type {
  switch (object) {
    case 0:
    case "TYPE_UNSPECIFIED":
      return ChangelogPayload_Type.TYPE_UNSPECIFIED;
    case 1:
    case "BASELINE":
      return ChangelogPayload_Type.BASELINE;
    case 2:
    case "MIGRATE":
      return ChangelogPayload_Type.MIGRATE;
    case 3:
    case "MIGRATE_SDL":
      return ChangelogPayload_Type.MIGRATE_SDL;
    case 4:
    case "MIGRATE_GHOST":
      return ChangelogPayload_Type.MIGRATE_GHOST;
    case 6:
    case "DATA":
      return ChangelogPayload_Type.DATA;
    case -1:
    case "UNRECOGNIZED":
    default:
      return ChangelogPayload_Type.UNRECOGNIZED;
  }
}

export function changelogPayload_TypeToJSON(object: ChangelogPayload_Type): string {
  switch (object) {
    case ChangelogPayload_Type.TYPE_UNSPECIFIED:
      return "TYPE_UNSPECIFIED";
    case ChangelogPayload_Type.BASELINE:
      return "BASELINE";
    case ChangelogPayload_Type.MIGRATE:
      return "MIGRATE";
    case ChangelogPayload_Type.MIGRATE_SDL:
      return "MIGRATE_SDL";
    case ChangelogPayload_Type.MIGRATE_GHOST:
      return "MIGRATE_GHOST";
    case ChangelogPayload_Type.DATA:
      return "DATA";
    case ChangelogPayload_Type.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export function changelogPayload_TypeToNumber(object: ChangelogPayload_Type): number {
  switch (object) {
    case ChangelogPayload_Type.TYPE_UNSPECIFIED:
      return 0;
    case ChangelogPayload_Type.BASELINE:
      return 1;
    case ChangelogPayload_Type.MIGRATE:
      return 2;
    case ChangelogPayload_Type.MIGRATE_SDL:
      return 3;
    case ChangelogPayload_Type.MIGRATE_GHOST:
      return 4;
    case ChangelogPayload_Type.DATA:
      return 6;
    case ChangelogPayload_Type.UNRECOGNIZED:
    default:
      return -1;
  }
}

function createBaseChangelogPayload(): ChangelogPayload {
  return {
    taskRun: "",
    issue: "",
    revision: Long.ZERO,
    changedResources: undefined,
    sheet: "",
    version: "",
    type: ChangelogPayload_Type.TYPE_UNSPECIFIED,
  };
}

export const ChangelogPayload: MessageFns<ChangelogPayload> = {
  encode(message: ChangelogPayload, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.taskRun !== "") {
      writer.uint32(10).string(message.taskRun);
    }
    if (message.issue !== "") {
      writer.uint32(18).string(message.issue);
    }
    if (!message.revision.equals(Long.ZERO)) {
      writer.uint32(24).int64(message.revision.toString());
    }
    if (message.changedResources !== undefined) {
      ChangedResources.encode(message.changedResources, writer.uint32(34).fork()).join();
    }
    if (message.sheet !== "") {
      writer.uint32(42).string(message.sheet);
    }
    if (message.version !== "") {
      writer.uint32(50).string(message.version);
    }
    if (message.type !== ChangelogPayload_Type.TYPE_UNSPECIFIED) {
      writer.uint32(56).int32(changelogPayload_TypeToNumber(message.type));
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): ChangelogPayload {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseChangelogPayload();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.taskRun = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.issue = reader.string();
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.revision = Long.fromString(reader.int64().toString());
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.changedResources = ChangedResources.decode(reader, reader.uint32());
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.sheet = reader.string();
          continue;
        case 6:
          if (tag !== 50) {
            break;
          }

          message.version = reader.string();
          continue;
        case 7:
          if (tag !== 56) {
            break;
          }

          message.type = changelogPayload_TypeFromJSON(reader.int32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ChangelogPayload {
    return {
      taskRun: isSet(object.taskRun) ? globalThis.String(object.taskRun) : "",
      issue: isSet(object.issue) ? globalThis.String(object.issue) : "",
      revision: isSet(object.revision) ? Long.fromValue(object.revision) : Long.ZERO,
      changedResources: isSet(object.changedResources) ? ChangedResources.fromJSON(object.changedResources) : undefined,
      sheet: isSet(object.sheet) ? globalThis.String(object.sheet) : "",
      version: isSet(object.version) ? globalThis.String(object.version) : "",
      type: isSet(object.type) ? changelogPayload_TypeFromJSON(object.type) : ChangelogPayload_Type.TYPE_UNSPECIFIED,
    };
  },

  toJSON(message: ChangelogPayload): unknown {
    const obj: any = {};
    if (message.taskRun !== "") {
      obj.taskRun = message.taskRun;
    }
    if (message.issue !== "") {
      obj.issue = message.issue;
    }
    if (!message.revision.equals(Long.ZERO)) {
      obj.revision = (message.revision || Long.ZERO).toString();
    }
    if (message.changedResources !== undefined) {
      obj.changedResources = ChangedResources.toJSON(message.changedResources);
    }
    if (message.sheet !== "") {
      obj.sheet = message.sheet;
    }
    if (message.version !== "") {
      obj.version = message.version;
    }
    if (message.type !== ChangelogPayload_Type.TYPE_UNSPECIFIED) {
      obj.type = changelogPayload_TypeToJSON(message.type);
    }
    return obj;
  },

  create(base?: DeepPartial<ChangelogPayload>): ChangelogPayload {
    return ChangelogPayload.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<ChangelogPayload>): ChangelogPayload {
    const message = createBaseChangelogPayload();
    message.taskRun = object.taskRun ?? "";
    message.issue = object.issue ?? "";
    message.revision = (object.revision !== undefined && object.revision !== null)
      ? Long.fromValue(object.revision)
      : Long.ZERO;
    message.changedResources = (object.changedResources !== undefined && object.changedResources !== null)
      ? ChangedResources.fromPartial(object.changedResources)
      : undefined;
    message.sheet = object.sheet ?? "";
    message.version = object.version ?? "";
    message.type = object.type ?? ChangelogPayload_Type.TYPE_UNSPECIFIED;
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Long ? string | number | Long : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

export interface MessageFns<T> {
  encode(message: T, writer?: BinaryWriter): BinaryWriter;
  decode(input: BinaryReader | Uint8Array, length?: number): T;
  fromJSON(object: any): T;
  toJSON(message: T): unknown;
  create(base?: DeepPartial<T>): T;
  fromPartial(object: DeepPartial<T>): T;
}
