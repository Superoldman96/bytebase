// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.3.0
//   protoc               unknown
// source: store/advice.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import Long from "long";
import { Position } from "./common";

export const protobufPackage = "bytebase.store";

export interface Advice {
  /** The advice status. */
  status: Advice_Status;
  /** The advice code. */
  code: number;
  /** The advice title. */
  title: string;
  /** The advice content. */
  content: string;
  /**
   * The start_position is inclusive and the end_position is exclusive.
   * TODO: use range instead.
   */
  startPosition: Position | undefined;
  endPosition: Position | undefined;
}

export enum Advice_Status {
  /** STATUS_UNSPECIFIED - Unspecified. */
  STATUS_UNSPECIFIED = "STATUS_UNSPECIFIED",
  SUCCESS = "SUCCESS",
  WARNING = "WARNING",
  ERROR = "ERROR",
  UNRECOGNIZED = "UNRECOGNIZED",
}

export function advice_StatusFromJSON(object: any): Advice_Status {
  switch (object) {
    case 0:
    case "STATUS_UNSPECIFIED":
      return Advice_Status.STATUS_UNSPECIFIED;
    case 1:
    case "SUCCESS":
      return Advice_Status.SUCCESS;
    case 2:
    case "WARNING":
      return Advice_Status.WARNING;
    case 3:
    case "ERROR":
      return Advice_Status.ERROR;
    case -1:
    case "UNRECOGNIZED":
    default:
      return Advice_Status.UNRECOGNIZED;
  }
}

export function advice_StatusToJSON(object: Advice_Status): string {
  switch (object) {
    case Advice_Status.STATUS_UNSPECIFIED:
      return "STATUS_UNSPECIFIED";
    case Advice_Status.SUCCESS:
      return "SUCCESS";
    case Advice_Status.WARNING:
      return "WARNING";
    case Advice_Status.ERROR:
      return "ERROR";
    case Advice_Status.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export function advice_StatusToNumber(object: Advice_Status): number {
  switch (object) {
    case Advice_Status.STATUS_UNSPECIFIED:
      return 0;
    case Advice_Status.SUCCESS:
      return 1;
    case Advice_Status.WARNING:
      return 2;
    case Advice_Status.ERROR:
      return 3;
    case Advice_Status.UNRECOGNIZED:
    default:
      return -1;
  }
}

function createBaseAdvice(): Advice {
  return {
    status: Advice_Status.STATUS_UNSPECIFIED,
    code: 0,
    title: "",
    content: "",
    startPosition: undefined,
    endPosition: undefined,
  };
}

export const Advice: MessageFns<Advice> = {
  encode(message: Advice, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.status !== Advice_Status.STATUS_UNSPECIFIED) {
      writer.uint32(8).int32(advice_StatusToNumber(message.status));
    }
    if (message.code !== 0) {
      writer.uint32(16).int32(message.code);
    }
    if (message.title !== "") {
      writer.uint32(26).string(message.title);
    }
    if (message.content !== "") {
      writer.uint32(34).string(message.content);
    }
    if (message.startPosition !== undefined) {
      Position.encode(message.startPosition, writer.uint32(50).fork()).join();
    }
    if (message.endPosition !== undefined) {
      Position.encode(message.endPosition, writer.uint32(58).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): Advice {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAdvice();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 8) {
            break;
          }

          message.status = advice_StatusFromJSON(reader.int32());
          continue;
        }
        case 2: {
          if (tag !== 16) {
            break;
          }

          message.code = reader.int32();
          continue;
        }
        case 3: {
          if (tag !== 26) {
            break;
          }

          message.title = reader.string();
          continue;
        }
        case 4: {
          if (tag !== 34) {
            break;
          }

          message.content = reader.string();
          continue;
        }
        case 6: {
          if (tag !== 50) {
            break;
          }

          message.startPosition = Position.decode(reader, reader.uint32());
          continue;
        }
        case 7: {
          if (tag !== 58) {
            break;
          }

          message.endPosition = Position.decode(reader, reader.uint32());
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

  fromJSON(object: any): Advice {
    return {
      status: isSet(object.status) ? advice_StatusFromJSON(object.status) : Advice_Status.STATUS_UNSPECIFIED,
      code: isSet(object.code) ? globalThis.Number(object.code) : 0,
      title: isSet(object.title) ? globalThis.String(object.title) : "",
      content: isSet(object.content) ? globalThis.String(object.content) : "",
      startPosition: isSet(object.startPosition) ? Position.fromJSON(object.startPosition) : undefined,
      endPosition: isSet(object.endPosition) ? Position.fromJSON(object.endPosition) : undefined,
    };
  },

  toJSON(message: Advice): unknown {
    const obj: any = {};
    if (message.status !== Advice_Status.STATUS_UNSPECIFIED) {
      obj.status = advice_StatusToJSON(message.status);
    }
    if (message.code !== 0) {
      obj.code = Math.round(message.code);
    }
    if (message.title !== "") {
      obj.title = message.title;
    }
    if (message.content !== "") {
      obj.content = message.content;
    }
    if (message.startPosition !== undefined) {
      obj.startPosition = Position.toJSON(message.startPosition);
    }
    if (message.endPosition !== undefined) {
      obj.endPosition = Position.toJSON(message.endPosition);
    }
    return obj;
  },

  create(base?: DeepPartial<Advice>): Advice {
    return Advice.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<Advice>): Advice {
    const message = createBaseAdvice();
    message.status = object.status ?? Advice_Status.STATUS_UNSPECIFIED;
    message.code = object.code ?? 0;
    message.title = object.title ?? "";
    message.content = object.content ?? "";
    message.startPosition = (object.startPosition !== undefined && object.startPosition !== null)
      ? Position.fromPartial(object.startPosition)
      : undefined;
    message.endPosition = (object.endPosition !== undefined && object.endPosition !== null)
      ? Position.fromPartial(object.endPosition)
      : undefined;
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
