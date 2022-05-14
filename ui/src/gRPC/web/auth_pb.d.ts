import * as jspb from 'google-protobuf'

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';
import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';


export class LoginInput extends jspb.Message {
  getUser(): string;
  setUser(value: string): LoginInput;

  getPassword(): string;
  setPassword(value: string): LoginInput;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginInput.AsObject;
  static toObject(includeInstance: boolean, msg: LoginInput): LoginInput.AsObject;
  static serializeBinaryToWriter(message: LoginInput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginInput;
  static deserializeBinaryFromReader(message: LoginInput, reader: jspb.BinaryReader): LoginInput;
}

export namespace LoginInput {
  export type AsObject = {
    user: string,
    password: string,
  }
}

export class AuthOutput extends jspb.Message {
  getToken(): string;
  setToken(value: string): AuthOutput;

  getExp(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setExp(value?: google_protobuf_timestamp_pb.Timestamp): AuthOutput;
  hasExp(): boolean;
  clearExp(): AuthOutput;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AuthOutput.AsObject;
  static toObject(includeInstance: boolean, msg: AuthOutput): AuthOutput.AsObject;
  static serializeBinaryToWriter(message: AuthOutput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AuthOutput;
  static deserializeBinaryFromReader(message: AuthOutput, reader: jspb.BinaryReader): AuthOutput;
}

export namespace AuthOutput {
  export type AsObject = {
    token: string,
    exp?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class TokenInput extends jspb.Message {
  getToken(): string;
  setToken(value: string): TokenInput;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TokenInput.AsObject;
  static toObject(includeInstance: boolean, msg: TokenInput): TokenInput.AsObject;
  static serializeBinaryToWriter(message: TokenInput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TokenInput;
  static deserializeBinaryFromReader(message: TokenInput, reader: jspb.BinaryReader): TokenInput;
}

export namespace TokenInput {
  export type AsObject = {
    token: string,
  }
}

export class SessionsInput extends jspb.Message {
  getToken(): string;
  setToken(value: string): SessionsInput;

  getUser(): string;
  setUser(value: string): SessionsInput;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SessionsInput.AsObject;
  static toObject(includeInstance: boolean, msg: SessionsInput): SessionsInput.AsObject;
  static serializeBinaryToWriter(message: SessionsInput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SessionsInput;
  static deserializeBinaryFromReader(message: SessionsInput, reader: jspb.BinaryReader): SessionsInput;
}

export namespace SessionsInput {
  export type AsObject = {
    token: string,
    user: string,
  }
}

export class SessionsOutput extends jspb.Message {
  getUserList(): Array<Sessions>;
  setUserList(value: Array<Sessions>): SessionsOutput;
  clearUserList(): SessionsOutput;
  addUser(value?: Sessions, index?: number): Sessions;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SessionsOutput.AsObject;
  static toObject(includeInstance: boolean, msg: SessionsOutput): SessionsOutput.AsObject;
  static serializeBinaryToWriter(message: SessionsOutput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SessionsOutput;
  static deserializeBinaryFromReader(message: SessionsOutput, reader: jspb.BinaryReader): SessionsOutput;
}

export namespace SessionsOutput {
  export type AsObject = {
    userList: Array<Sessions.AsObject>,
  }
}

export class Sessions extends jspb.Message {
  getToken(): string;
  setToken(value: string): Sessions;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Sessions.AsObject;
  static toObject(includeInstance: boolean, msg: Sessions): Sessions.AsObject;
  static serializeBinaryToWriter(message: Sessions, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Sessions;
  static deserializeBinaryFromReader(message: Sessions, reader: jspb.BinaryReader): Sessions;
}

export namespace Sessions {
  export type AsObject = {
    token: string,
  }
}

