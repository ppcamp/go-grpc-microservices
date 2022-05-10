import * as jspb from 'google-protobuf'

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';


export class CreateInput extends jspb.Message {
  getUser(): string;
  setUser(value: string): CreateInput;

  getPassword(): string;
  setPassword(value: string): CreateInput;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateInput.AsObject;
  static toObject(includeInstance: boolean, msg: CreateInput): CreateInput.AsObject;
  static serializeBinaryToWriter(message: CreateInput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateInput;
  static deserializeBinaryFromReader(message: CreateInput, reader: jspb.BinaryReader): CreateInput;
}

export namespace CreateInput {
  export type AsObject = {
    user: string,
    password: string,
  }
}

export class RecoverInput extends jspb.Message {
  getEmail(): string;
  setEmail(value: string): RecoverInput;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RecoverInput.AsObject;
  static toObject(includeInstance: boolean, msg: RecoverInput): RecoverInput.AsObject;
  static serializeBinaryToWriter(message: RecoverInput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RecoverInput;
  static deserializeBinaryFromReader(message: RecoverInput, reader: jspb.BinaryReader): RecoverInput;
}

export namespace RecoverInput {
  export type AsObject = {
    email: string,
  }
}

export class RecoverOutput extends jspb.Message {
  getSecret(): string;
  setSecret(value: string): RecoverOutput;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RecoverOutput.AsObject;
  static toObject(includeInstance: boolean, msg: RecoverOutput): RecoverOutput.AsObject;
  static serializeBinaryToWriter(message: RecoverOutput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RecoverOutput;
  static deserializeBinaryFromReader(message: RecoverOutput, reader: jspb.BinaryReader): RecoverOutput;
}

export namespace RecoverOutput {
  export type AsObject = {
    secret: string,
  }
}

export class CreateOutput extends jspb.Message {
  getSecret(): string;
  setSecret(value: string): CreateOutput;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateOutput.AsObject;
  static toObject(includeInstance: boolean, msg: CreateOutput): CreateOutput.AsObject;
  static serializeBinaryToWriter(message: CreateOutput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateOutput;
  static deserializeBinaryFromReader(message: CreateOutput, reader: jspb.BinaryReader): CreateOutput;
}

export namespace CreateOutput {
  export type AsObject = {
    secret: string,
  }
}

export class ActivateInput extends jspb.Message {
  getSecret(): string;
  setSecret(value: string): ActivateInput;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ActivateInput.AsObject;
  static toObject(includeInstance: boolean, msg: ActivateInput): ActivateInput.AsObject;
  static serializeBinaryToWriter(message: ActivateInput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ActivateInput;
  static deserializeBinaryFromReader(message: ActivateInput, reader: jspb.BinaryReader): ActivateInput;
}

export namespace ActivateInput {
  export type AsObject = {
    secret: string,
  }
}

export class UpdateInput extends jspb.Message {
  getSecret(): string;
  setSecret(value: string): UpdateInput;

  getPassword(): string;
  setPassword(value: string): UpdateInput;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateInput.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateInput): UpdateInput.AsObject;
  static serializeBinaryToWriter(message: UpdateInput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateInput;
  static deserializeBinaryFromReader(message: UpdateInput, reader: jspb.BinaryReader): UpdateInput;
}

export namespace UpdateInput {
  export type AsObject = {
    secret: string,
    password: string,
  }
}

export class DeleteInput extends jspb.Message {
  getToken(): string;
  setToken(value: string): DeleteInput;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteInput.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteInput): DeleteInput.AsObject;
  static serializeBinaryToWriter(message: DeleteInput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteInput;
  static deserializeBinaryFromReader(message: DeleteInput, reader: jspb.BinaryReader): DeleteInput;
}

export namespace DeleteInput {
  export type AsObject = {
    token: string,
  }
}

