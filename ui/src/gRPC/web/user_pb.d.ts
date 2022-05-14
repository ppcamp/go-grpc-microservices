import * as jspb from 'google-protobuf'

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';


export class CreateUserInput extends jspb.Message {
  getFirstName(): string;
  setFirstName(value: string): CreateUserInput;

  getMiddleName(): string;
  setMiddleName(value: string): CreateUserInput;

  getLastName(): string;
  setLastName(value: string): CreateUserInput;

  getNickname(): string;
  setNickname(value: string): CreateUserInput;

  getEmail(): string;
  setEmail(value: string): CreateUserInput;

  getPassword(): string;
  setPassword(value: string): CreateUserInput;

  getBirthdate(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setBirthdate(value?: google_protobuf_timestamp_pb.Timestamp): CreateUserInput;
  hasBirthdate(): boolean;
  clearBirthdate(): CreateUserInput;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateUserInput.AsObject;
  static toObject(includeInstance: boolean, msg: CreateUserInput): CreateUserInput.AsObject;
  static serializeBinaryToWriter(message: CreateUserInput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateUserInput;
  static deserializeBinaryFromReader(message: CreateUserInput, reader: jspb.BinaryReader): CreateUserInput;
}

export namespace CreateUserInput {
  export type AsObject = {
    firstName: string,
    middleName: string,
    lastName: string,
    nickname: string,
    email: string,
    password: string,
    birthdate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class UpdateUserInput extends jspb.Message {
  getToken(): string;
  setToken(value: string): UpdateUserInput;

  getUserData(): CreateUserInput | undefined;
  setUserData(value?: CreateUserInput): UpdateUserInput;
  hasUserData(): boolean;
  clearUserData(): UpdateUserInput;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateUserInput.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateUserInput): UpdateUserInput.AsObject;
  static serializeBinaryToWriter(message: UpdateUserInput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateUserInput;
  static deserializeBinaryFromReader(message: UpdateUserInput, reader: jspb.BinaryReader): UpdateUserInput;
}

export namespace UpdateUserInput {
  export type AsObject = {
    token: string,
    userData?: CreateUserInput.AsObject,
  }
}

export class DeleteUserInput extends jspb.Message {
  getToken(): string;
  setToken(value: string): DeleteUserInput;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteUserInput.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteUserInput): DeleteUserInput.AsObject;
  static serializeBinaryToWriter(message: DeleteUserInput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteUserInput;
  static deserializeBinaryFromReader(message: DeleteUserInput, reader: jspb.BinaryReader): DeleteUserInput;
}

export namespace DeleteUserInput {
  export type AsObject = {
    token: string,
  }
}

