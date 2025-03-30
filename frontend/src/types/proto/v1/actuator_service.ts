// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.3.0
//   protoc               unknown
// source: v1/actuator_service.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import Long from "long";
import { Empty } from "../google/protobuf/empty";
import { FieldMask } from "../google/protobuf/field_mask";
import { Timestamp } from "../google/protobuf/timestamp";
import { State, stateFromJSON, stateToJSON, stateToNumber } from "./common";
import { PasswordRestrictionSetting } from "./setting_service";
import { UserType, userTypeFromJSON, userTypeToJSON, userTypeToNumber } from "./user_service";

export const protobufPackage = "bytebase.v1";

/** The request message for getting the theme resource. */
export interface GetResourcePackageRequest {
}

/** The theme resources. */
export interface ResourcePackage {
  /** The branding logo. */
  logo: Uint8Array;
}

export interface GetActuatorInfoRequest {
}

export interface UpdateActuatorInfoRequest {
  /** The actuator to update. */
  actuator:
    | ActuatorInfo
    | undefined;
  /** The list of fields to update. */
  updateMask: string[] | undefined;
}

export interface DeleteCacheRequest {
}

/**
 * ServerInfo is the API message for server info.
 * Actuator concept is similar to the Spring Boot Actuator.
 */
export interface ActuatorInfo {
  /** version is the bytebase's server version */
  version: string;
  /** git_commit is the git commit hash of the build */
  gitCommit: string;
  /** readonly flag means if the Bytebase is running in readonly mode. */
  readonly: boolean;
  /** saas flag means if the Bytebase is running in SaaS mode, some features are not allowed to edit by users. */
  saas: boolean;
  /** demo flag means if the Bytebase is running in demo mode. */
  demo: boolean;
  /** host is the Bytebase instance host. */
  host: string;
  /** port is the Bytebase instance port. */
  port: string;
  /** external_url is the URL where user or webhook callback visits Bytebase. */
  externalUrl: string;
  /** need_admin_setup flag means the Bytebase instance doesn't have any end users. */
  needAdminSetup: boolean;
  /** disallow_signup is the flag to disable self-service signup. */
  disallowSignup: boolean;
  /** last_active_time is the service last active time in UTC Time Format, any API calls will refresh this value. */
  lastActiveTime:
    | Timestamp
    | undefined;
  /** require_2fa is the flag to require 2FA for all users. */
  require2fa: boolean;
  /** workspace_id is the identifier for the workspace. */
  workspaceId: string;
  /** debug flag means if the debug mode is enabled. */
  debug: boolean;
  unlicensedFeatures: string[];
  /** disallow_password_signin is the flag to disallow user signin with email&password. (except workspace admins) */
  disallowPasswordSignin: boolean;
  passwordRestriction:
    | PasswordRestrictionSetting
    | undefined;
  /** docker flag means if the Bytebase instance is running in docker. */
  docker: boolean;
  userStats: ActuatorInfo_StatUser[];
  activatedInstanceCount: number;
  totalInstanceCount: number;
}

export interface ActuatorInfo_StatUser {
  userType: UserType;
  state: State;
  count: number;
}

function createBaseGetResourcePackageRequest(): GetResourcePackageRequest {
  return {};
}

export const GetResourcePackageRequest: MessageFns<GetResourcePackageRequest> = {
  encode(_: GetResourcePackageRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): GetResourcePackageRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetResourcePackageRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): GetResourcePackageRequest {
    return {};
  },

  toJSON(_: GetResourcePackageRequest): unknown {
    const obj: any = {};
    return obj;
  },

  create(base?: DeepPartial<GetResourcePackageRequest>): GetResourcePackageRequest {
    return GetResourcePackageRequest.fromPartial(base ?? {});
  },
  fromPartial(_: DeepPartial<GetResourcePackageRequest>): GetResourcePackageRequest {
    const message = createBaseGetResourcePackageRequest();
    return message;
  },
};

function createBaseResourcePackage(): ResourcePackage {
  return { logo: new Uint8Array(0) };
}

export const ResourcePackage: MessageFns<ResourcePackage> = {
  encode(message: ResourcePackage, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.logo.length !== 0) {
      writer.uint32(10).bytes(message.logo);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): ResourcePackage {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseResourcePackage();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.logo = reader.bytes();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ResourcePackage {
    return { logo: isSet(object.logo) ? bytesFromBase64(object.logo) : new Uint8Array(0) };
  },

  toJSON(message: ResourcePackage): unknown {
    const obj: any = {};
    if (message.logo.length !== 0) {
      obj.logo = base64FromBytes(message.logo);
    }
    return obj;
  },

  create(base?: DeepPartial<ResourcePackage>): ResourcePackage {
    return ResourcePackage.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<ResourcePackage>): ResourcePackage {
    const message = createBaseResourcePackage();
    message.logo = object.logo ?? new Uint8Array(0);
    return message;
  },
};

function createBaseGetActuatorInfoRequest(): GetActuatorInfoRequest {
  return {};
}

export const GetActuatorInfoRequest: MessageFns<GetActuatorInfoRequest> = {
  encode(_: GetActuatorInfoRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): GetActuatorInfoRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetActuatorInfoRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): GetActuatorInfoRequest {
    return {};
  },

  toJSON(_: GetActuatorInfoRequest): unknown {
    const obj: any = {};
    return obj;
  },

  create(base?: DeepPartial<GetActuatorInfoRequest>): GetActuatorInfoRequest {
    return GetActuatorInfoRequest.fromPartial(base ?? {});
  },
  fromPartial(_: DeepPartial<GetActuatorInfoRequest>): GetActuatorInfoRequest {
    const message = createBaseGetActuatorInfoRequest();
    return message;
  },
};

function createBaseUpdateActuatorInfoRequest(): UpdateActuatorInfoRequest {
  return { actuator: undefined, updateMask: undefined };
}

export const UpdateActuatorInfoRequest: MessageFns<UpdateActuatorInfoRequest> = {
  encode(message: UpdateActuatorInfoRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.actuator !== undefined) {
      ActuatorInfo.encode(message.actuator, writer.uint32(10).fork()).join();
    }
    if (message.updateMask !== undefined) {
      FieldMask.encode(FieldMask.wrap(message.updateMask), writer.uint32(18).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): UpdateActuatorInfoRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateActuatorInfoRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.actuator = ActuatorInfo.decode(reader, reader.uint32());
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.updateMask = FieldMask.unwrap(FieldMask.decode(reader, reader.uint32()));
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UpdateActuatorInfoRequest {
    return {
      actuator: isSet(object.actuator) ? ActuatorInfo.fromJSON(object.actuator) : undefined,
      updateMask: isSet(object.updateMask) ? FieldMask.unwrap(FieldMask.fromJSON(object.updateMask)) : undefined,
    };
  },

  toJSON(message: UpdateActuatorInfoRequest): unknown {
    const obj: any = {};
    if (message.actuator !== undefined) {
      obj.actuator = ActuatorInfo.toJSON(message.actuator);
    }
    if (message.updateMask !== undefined) {
      obj.updateMask = FieldMask.toJSON(FieldMask.wrap(message.updateMask));
    }
    return obj;
  },

  create(base?: DeepPartial<UpdateActuatorInfoRequest>): UpdateActuatorInfoRequest {
    return UpdateActuatorInfoRequest.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<UpdateActuatorInfoRequest>): UpdateActuatorInfoRequest {
    const message = createBaseUpdateActuatorInfoRequest();
    message.actuator = (object.actuator !== undefined && object.actuator !== null)
      ? ActuatorInfo.fromPartial(object.actuator)
      : undefined;
    message.updateMask = object.updateMask ?? undefined;
    return message;
  },
};

function createBaseDeleteCacheRequest(): DeleteCacheRequest {
  return {};
}

export const DeleteCacheRequest: MessageFns<DeleteCacheRequest> = {
  encode(_: DeleteCacheRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): DeleteCacheRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteCacheRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): DeleteCacheRequest {
    return {};
  },

  toJSON(_: DeleteCacheRequest): unknown {
    const obj: any = {};
    return obj;
  },

  create(base?: DeepPartial<DeleteCacheRequest>): DeleteCacheRequest {
    return DeleteCacheRequest.fromPartial(base ?? {});
  },
  fromPartial(_: DeepPartial<DeleteCacheRequest>): DeleteCacheRequest {
    const message = createBaseDeleteCacheRequest();
    return message;
  },
};

function createBaseActuatorInfo(): ActuatorInfo {
  return {
    version: "",
    gitCommit: "",
    readonly: false,
    saas: false,
    demo: false,
    host: "",
    port: "",
    externalUrl: "",
    needAdminSetup: false,
    disallowSignup: false,
    lastActiveTime: undefined,
    require2fa: false,
    workspaceId: "",
    debug: false,
    unlicensedFeatures: [],
    disallowPasswordSignin: false,
    passwordRestriction: undefined,
    docker: false,
    userStats: [],
    activatedInstanceCount: 0,
    totalInstanceCount: 0,
  };
}

export const ActuatorInfo: MessageFns<ActuatorInfo> = {
  encode(message: ActuatorInfo, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.version !== "") {
      writer.uint32(10).string(message.version);
    }
    if (message.gitCommit !== "") {
      writer.uint32(18).string(message.gitCommit);
    }
    if (message.readonly !== false) {
      writer.uint32(24).bool(message.readonly);
    }
    if (message.saas !== false) {
      writer.uint32(32).bool(message.saas);
    }
    if (message.demo !== false) {
      writer.uint32(40).bool(message.demo);
    }
    if (message.host !== "") {
      writer.uint32(50).string(message.host);
    }
    if (message.port !== "") {
      writer.uint32(58).string(message.port);
    }
    if (message.externalUrl !== "") {
      writer.uint32(66).string(message.externalUrl);
    }
    if (message.needAdminSetup !== false) {
      writer.uint32(72).bool(message.needAdminSetup);
    }
    if (message.disallowSignup !== false) {
      writer.uint32(80).bool(message.disallowSignup);
    }
    if (message.lastActiveTime !== undefined) {
      Timestamp.encode(message.lastActiveTime, writer.uint32(90).fork()).join();
    }
    if (message.require2fa !== false) {
      writer.uint32(96).bool(message.require2fa);
    }
    if (message.workspaceId !== "") {
      writer.uint32(106).string(message.workspaceId);
    }
    if (message.debug !== false) {
      writer.uint32(120).bool(message.debug);
    }
    for (const v of message.unlicensedFeatures) {
      writer.uint32(154).string(v!);
    }
    if (message.disallowPasswordSignin !== false) {
      writer.uint32(160).bool(message.disallowPasswordSignin);
    }
    if (message.passwordRestriction !== undefined) {
      PasswordRestrictionSetting.encode(message.passwordRestriction, writer.uint32(170).fork()).join();
    }
    if (message.docker !== false) {
      writer.uint32(176).bool(message.docker);
    }
    for (const v of message.userStats) {
      ActuatorInfo_StatUser.encode(v!, writer.uint32(186).fork()).join();
    }
    if (message.activatedInstanceCount !== 0) {
      writer.uint32(192).int32(message.activatedInstanceCount);
    }
    if (message.totalInstanceCount !== 0) {
      writer.uint32(200).int32(message.totalInstanceCount);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): ActuatorInfo {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseActuatorInfo();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.version = reader.string();
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.gitCommit = reader.string();
          continue;
        }
        case 3: {
          if (tag !== 24) {
            break;
          }

          message.readonly = reader.bool();
          continue;
        }
        case 4: {
          if (tag !== 32) {
            break;
          }

          message.saas = reader.bool();
          continue;
        }
        case 5: {
          if (tag !== 40) {
            break;
          }

          message.demo = reader.bool();
          continue;
        }
        case 6: {
          if (tag !== 50) {
            break;
          }

          message.host = reader.string();
          continue;
        }
        case 7: {
          if (tag !== 58) {
            break;
          }

          message.port = reader.string();
          continue;
        }
        case 8: {
          if (tag !== 66) {
            break;
          }

          message.externalUrl = reader.string();
          continue;
        }
        case 9: {
          if (tag !== 72) {
            break;
          }

          message.needAdminSetup = reader.bool();
          continue;
        }
        case 10: {
          if (tag !== 80) {
            break;
          }

          message.disallowSignup = reader.bool();
          continue;
        }
        case 11: {
          if (tag !== 90) {
            break;
          }

          message.lastActiveTime = Timestamp.decode(reader, reader.uint32());
          continue;
        }
        case 12: {
          if (tag !== 96) {
            break;
          }

          message.require2fa = reader.bool();
          continue;
        }
        case 13: {
          if (tag !== 106) {
            break;
          }

          message.workspaceId = reader.string();
          continue;
        }
        case 15: {
          if (tag !== 120) {
            break;
          }

          message.debug = reader.bool();
          continue;
        }
        case 19: {
          if (tag !== 154) {
            break;
          }

          message.unlicensedFeatures.push(reader.string());
          continue;
        }
        case 20: {
          if (tag !== 160) {
            break;
          }

          message.disallowPasswordSignin = reader.bool();
          continue;
        }
        case 21: {
          if (tag !== 170) {
            break;
          }

          message.passwordRestriction = PasswordRestrictionSetting.decode(reader, reader.uint32());
          continue;
        }
        case 22: {
          if (tag !== 176) {
            break;
          }

          message.docker = reader.bool();
          continue;
        }
        case 23: {
          if (tag !== 186) {
            break;
          }

          message.userStats.push(ActuatorInfo_StatUser.decode(reader, reader.uint32()));
          continue;
        }
        case 24: {
          if (tag !== 192) {
            break;
          }

          message.activatedInstanceCount = reader.int32();
          continue;
        }
        case 25: {
          if (tag !== 200) {
            break;
          }

          message.totalInstanceCount = reader.int32();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ActuatorInfo {
    return {
      version: isSet(object.version) ? globalThis.String(object.version) : "",
      gitCommit: isSet(object.gitCommit) ? globalThis.String(object.gitCommit) : "",
      readonly: isSet(object.readonly) ? globalThis.Boolean(object.readonly) : false,
      saas: isSet(object.saas) ? globalThis.Boolean(object.saas) : false,
      demo: isSet(object.demo) ? globalThis.Boolean(object.demo) : false,
      host: isSet(object.host) ? globalThis.String(object.host) : "",
      port: isSet(object.port) ? globalThis.String(object.port) : "",
      externalUrl: isSet(object.externalUrl) ? globalThis.String(object.externalUrl) : "",
      needAdminSetup: isSet(object.needAdminSetup) ? globalThis.Boolean(object.needAdminSetup) : false,
      disallowSignup: isSet(object.disallowSignup) ? globalThis.Boolean(object.disallowSignup) : false,
      lastActiveTime: isSet(object.lastActiveTime) ? fromJsonTimestamp(object.lastActiveTime) : undefined,
      require2fa: isSet(object.require2fa) ? globalThis.Boolean(object.require2fa) : false,
      workspaceId: isSet(object.workspaceId) ? globalThis.String(object.workspaceId) : "",
      debug: isSet(object.debug) ? globalThis.Boolean(object.debug) : false,
      unlicensedFeatures: globalThis.Array.isArray(object?.unlicensedFeatures)
        ? object.unlicensedFeatures.map((e: any) => globalThis.String(e))
        : [],
      disallowPasswordSignin: isSet(object.disallowPasswordSignin)
        ? globalThis.Boolean(object.disallowPasswordSignin)
        : false,
      passwordRestriction: isSet(object.passwordRestriction)
        ? PasswordRestrictionSetting.fromJSON(object.passwordRestriction)
        : undefined,
      docker: isSet(object.docker) ? globalThis.Boolean(object.docker) : false,
      userStats: globalThis.Array.isArray(object?.userStats)
        ? object.userStats.map((e: any) => ActuatorInfo_StatUser.fromJSON(e))
        : [],
      activatedInstanceCount: isSet(object.activatedInstanceCount)
        ? globalThis.Number(object.activatedInstanceCount)
        : 0,
      totalInstanceCount: isSet(object.totalInstanceCount) ? globalThis.Number(object.totalInstanceCount) : 0,
    };
  },

  toJSON(message: ActuatorInfo): unknown {
    const obj: any = {};
    if (message.version !== "") {
      obj.version = message.version;
    }
    if (message.gitCommit !== "") {
      obj.gitCommit = message.gitCommit;
    }
    if (message.readonly !== false) {
      obj.readonly = message.readonly;
    }
    if (message.saas !== false) {
      obj.saas = message.saas;
    }
    if (message.demo !== false) {
      obj.demo = message.demo;
    }
    if (message.host !== "") {
      obj.host = message.host;
    }
    if (message.port !== "") {
      obj.port = message.port;
    }
    if (message.externalUrl !== "") {
      obj.externalUrl = message.externalUrl;
    }
    if (message.needAdminSetup !== false) {
      obj.needAdminSetup = message.needAdminSetup;
    }
    if (message.disallowSignup !== false) {
      obj.disallowSignup = message.disallowSignup;
    }
    if (message.lastActiveTime !== undefined) {
      obj.lastActiveTime = fromTimestamp(message.lastActiveTime).toISOString();
    }
    if (message.require2fa !== false) {
      obj.require2fa = message.require2fa;
    }
    if (message.workspaceId !== "") {
      obj.workspaceId = message.workspaceId;
    }
    if (message.debug !== false) {
      obj.debug = message.debug;
    }
    if (message.unlicensedFeatures?.length) {
      obj.unlicensedFeatures = message.unlicensedFeatures;
    }
    if (message.disallowPasswordSignin !== false) {
      obj.disallowPasswordSignin = message.disallowPasswordSignin;
    }
    if (message.passwordRestriction !== undefined) {
      obj.passwordRestriction = PasswordRestrictionSetting.toJSON(message.passwordRestriction);
    }
    if (message.docker !== false) {
      obj.docker = message.docker;
    }
    if (message.userStats?.length) {
      obj.userStats = message.userStats.map((e) => ActuatorInfo_StatUser.toJSON(e));
    }
    if (message.activatedInstanceCount !== 0) {
      obj.activatedInstanceCount = Math.round(message.activatedInstanceCount);
    }
    if (message.totalInstanceCount !== 0) {
      obj.totalInstanceCount = Math.round(message.totalInstanceCount);
    }
    return obj;
  },

  create(base?: DeepPartial<ActuatorInfo>): ActuatorInfo {
    return ActuatorInfo.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<ActuatorInfo>): ActuatorInfo {
    const message = createBaseActuatorInfo();
    message.version = object.version ?? "";
    message.gitCommit = object.gitCommit ?? "";
    message.readonly = object.readonly ?? false;
    message.saas = object.saas ?? false;
    message.demo = object.demo ?? false;
    message.host = object.host ?? "";
    message.port = object.port ?? "";
    message.externalUrl = object.externalUrl ?? "";
    message.needAdminSetup = object.needAdminSetup ?? false;
    message.disallowSignup = object.disallowSignup ?? false;
    message.lastActiveTime = (object.lastActiveTime !== undefined && object.lastActiveTime !== null)
      ? Timestamp.fromPartial(object.lastActiveTime)
      : undefined;
    message.require2fa = object.require2fa ?? false;
    message.workspaceId = object.workspaceId ?? "";
    message.debug = object.debug ?? false;
    message.unlicensedFeatures = object.unlicensedFeatures?.map((e) => e) || [];
    message.disallowPasswordSignin = object.disallowPasswordSignin ?? false;
    message.passwordRestriction = (object.passwordRestriction !== undefined && object.passwordRestriction !== null)
      ? PasswordRestrictionSetting.fromPartial(object.passwordRestriction)
      : undefined;
    message.docker = object.docker ?? false;
    message.userStats = object.userStats?.map((e) => ActuatorInfo_StatUser.fromPartial(e)) || [];
    message.activatedInstanceCount = object.activatedInstanceCount ?? 0;
    message.totalInstanceCount = object.totalInstanceCount ?? 0;
    return message;
  },
};

function createBaseActuatorInfo_StatUser(): ActuatorInfo_StatUser {
  return { userType: UserType.USER_TYPE_UNSPECIFIED, state: State.STATE_UNSPECIFIED, count: 0 };
}

export const ActuatorInfo_StatUser: MessageFns<ActuatorInfo_StatUser> = {
  encode(message: ActuatorInfo_StatUser, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.userType !== UserType.USER_TYPE_UNSPECIFIED) {
      writer.uint32(8).int32(userTypeToNumber(message.userType));
    }
    if (message.state !== State.STATE_UNSPECIFIED) {
      writer.uint32(16).int32(stateToNumber(message.state));
    }
    if (message.count !== 0) {
      writer.uint32(24).int32(message.count);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): ActuatorInfo_StatUser {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseActuatorInfo_StatUser();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 8) {
            break;
          }

          message.userType = userTypeFromJSON(reader.int32());
          continue;
        }
        case 2: {
          if (tag !== 16) {
            break;
          }

          message.state = stateFromJSON(reader.int32());
          continue;
        }
        case 3: {
          if (tag !== 24) {
            break;
          }

          message.count = reader.int32();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ActuatorInfo_StatUser {
    return {
      userType: isSet(object.userType) ? userTypeFromJSON(object.userType) : UserType.USER_TYPE_UNSPECIFIED,
      state: isSet(object.state) ? stateFromJSON(object.state) : State.STATE_UNSPECIFIED,
      count: isSet(object.count) ? globalThis.Number(object.count) : 0,
    };
  },

  toJSON(message: ActuatorInfo_StatUser): unknown {
    const obj: any = {};
    if (message.userType !== UserType.USER_TYPE_UNSPECIFIED) {
      obj.userType = userTypeToJSON(message.userType);
    }
    if (message.state !== State.STATE_UNSPECIFIED) {
      obj.state = stateToJSON(message.state);
    }
    if (message.count !== 0) {
      obj.count = Math.round(message.count);
    }
    return obj;
  },

  create(base?: DeepPartial<ActuatorInfo_StatUser>): ActuatorInfo_StatUser {
    return ActuatorInfo_StatUser.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<ActuatorInfo_StatUser>): ActuatorInfo_StatUser {
    const message = createBaseActuatorInfo_StatUser();
    message.userType = object.userType ?? UserType.USER_TYPE_UNSPECIFIED;
    message.state = object.state ?? State.STATE_UNSPECIFIED;
    message.count = object.count ?? 0;
    return message;
  },
};

export type ActuatorServiceDefinition = typeof ActuatorServiceDefinition;
export const ActuatorServiceDefinition = {
  name: "ActuatorService",
  fullName: "bytebase.v1.ActuatorService",
  methods: {
    getActuatorInfo: {
      name: "GetActuatorInfo",
      requestType: GetActuatorInfoRequest,
      requestStream: false,
      responseType: ActuatorInfo,
      responseStream: false,
      options: {
        _unknownFields: {
          8410: [new Uint8Array([0])],
          800000: [new Uint8Array([1])],
          578365826: [
            new Uint8Array([19, 18, 17, 47, 118, 49, 47, 97, 99, 116, 117, 97, 116, 111, 114, 47, 105, 110, 102, 111]),
          ],
        },
      },
    },
    updateActuatorInfo: {
      name: "UpdateActuatorInfo",
      requestType: UpdateActuatorInfoRequest,
      requestStream: false,
      responseType: ActuatorInfo,
      responseStream: false,
      options: {
        _unknownFields: {
          8410: [
            new Uint8Array([
              20,
              97,
              99,
              116,
              117,
              97,
              116,
              111,
              114,
              44,
              117,
              112,
              100,
              97,
              116,
              101,
              95,
              109,
              97,
              115,
              107,
            ]),
          ],
          800010: [new Uint8Array([15, 98, 98, 46, 115, 101, 116, 116, 105, 110, 103, 115, 46, 115, 101, 116])],
          800016: [new Uint8Array([1])],
          578365826: [
            new Uint8Array([
              29,
              58,
              8,
              97,
              99,
              116,
              117,
              97,
              116,
              111,
              114,
              50,
              17,
              47,
              118,
              49,
              47,
              97,
              99,
              116,
              117,
              97,
              116,
              111,
              114,
              47,
              105,
              110,
              102,
              111,
            ]),
          ],
        },
      },
    },
    deleteCache: {
      name: "DeleteCache",
      requestType: DeleteCacheRequest,
      requestStream: false,
      responseType: Empty,
      responseStream: false,
      options: {
        _unknownFields: {
          800000: [new Uint8Array([1])],
          578365826: [
            new Uint8Array([
              20,
              42,
              18,
              47,
              118,
              49,
              47,
              97,
              99,
              116,
              117,
              97,
              116,
              111,
              114,
              47,
              99,
              97,
              99,
              104,
              101,
            ]),
          ],
        },
      },
    },
    getResourcePackage: {
      name: "GetResourcePackage",
      requestType: GetResourcePackageRequest,
      requestStream: false,
      responseType: ResourcePackage,
      responseStream: false,
      options: {
        _unknownFields: {
          8410: [new Uint8Array([0])],
          800000: [new Uint8Array([1])],
          578365826: [
            new Uint8Array([
              24,
              18,
              22,
              47,
              118,
              49,
              47,
              97,
              99,
              116,
              117,
              97,
              116,
              111,
              114,
              47,
              114,
              101,
              115,
              111,
              117,
              114,
              99,
              101,
              115,
            ]),
          ],
        },
      },
    },
  },
} as const;

function bytesFromBase64(b64: string): Uint8Array {
  const bin = globalThis.atob(b64);
  const arr = new Uint8Array(bin.length);
  for (let i = 0; i < bin.length; ++i) {
    arr[i] = bin.charCodeAt(i);
  }
  return arr;
}

function base64FromBytes(arr: Uint8Array): string {
  const bin: string[] = [];
  arr.forEach((byte) => {
    bin.push(globalThis.String.fromCharCode(byte));
  });
  return globalThis.btoa(bin.join(""));
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Long ? string | number | Long : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function toTimestamp(date: Date): Timestamp {
  const seconds = numberToLong(Math.trunc(date.getTime() / 1_000));
  const nanos = (date.getTime() % 1_000) * 1_000_000;
  return { seconds, nanos };
}

function fromTimestamp(t: Timestamp): Date {
  let millis = (t.seconds.toNumber() || 0) * 1_000;
  millis += (t.nanos || 0) / 1_000_000;
  return new globalThis.Date(millis);
}

function fromJsonTimestamp(o: any): Timestamp {
  if (o instanceof globalThis.Date) {
    return toTimestamp(o);
  } else if (typeof o === "string") {
    return toTimestamp(new globalThis.Date(o));
  } else {
    return Timestamp.fromJSON(o);
  }
}

function numberToLong(number: number) {
  return Long.fromNumber(number);
}

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
