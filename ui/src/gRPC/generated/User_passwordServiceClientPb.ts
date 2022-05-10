/**
 * @fileoverview gRPC-Web generated client stub for 
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as user_password_pb from './user_password_pb';
import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';


export class UserPasswordServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorCreate = new grpcWeb.MethodDescriptor(
    '/UserPasswordService/Create',
    grpcWeb.MethodType.UNARY,
    user_password_pb.CreateInput,
    user_password_pb.CreateOutput,
    (request: user_password_pb.CreateInput) => {
      return request.serializeBinary();
    },
    user_password_pb.CreateOutput.deserializeBinary
  );

  create(
    request: user_password_pb.CreateInput,
    metadata: grpcWeb.Metadata | null): Promise<user_password_pb.CreateOutput>;

  create(
    request: user_password_pb.CreateInput,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: user_password_pb.CreateOutput) => void): grpcWeb.ClientReadableStream<user_password_pb.CreateOutput>;

  create(
    request: user_password_pb.CreateInput,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: user_password_pb.CreateOutput) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/UserPasswordService/Create',
        request,
        metadata || {},
        this.methodDescriptorCreate,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/UserPasswordService/Create',
    request,
    metadata || {},
    this.methodDescriptorCreate);
  }

  methodDescriptorActivate = new grpcWeb.MethodDescriptor(
    '/UserPasswordService/Activate',
    grpcWeb.MethodType.UNARY,
    user_password_pb.ActivateInput,
    google_protobuf_empty_pb.Empty,
    (request: user_password_pb.ActivateInput) => {
      return request.serializeBinary();
    },
    google_protobuf_empty_pb.Empty.deserializeBinary
  );

  activate(
    request: user_password_pb.ActivateInput,
    metadata: grpcWeb.Metadata | null): Promise<google_protobuf_empty_pb.Empty>;

  activate(
    request: user_password_pb.ActivateInput,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  activate(
    request: user_password_pb.ActivateInput,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/UserPasswordService/Activate',
        request,
        metadata || {},
        this.methodDescriptorActivate,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/UserPasswordService/Activate',
    request,
    metadata || {},
    this.methodDescriptorActivate);
  }

  methodDescriptorRecover = new grpcWeb.MethodDescriptor(
    '/UserPasswordService/Recover',
    grpcWeb.MethodType.UNARY,
    user_password_pb.RecoverInput,
    user_password_pb.RecoverOutput,
    (request: user_password_pb.RecoverInput) => {
      return request.serializeBinary();
    },
    user_password_pb.RecoverOutput.deserializeBinary
  );

  recover(
    request: user_password_pb.RecoverInput,
    metadata: grpcWeb.Metadata | null): Promise<user_password_pb.RecoverOutput>;

  recover(
    request: user_password_pb.RecoverInput,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: user_password_pb.RecoverOutput) => void): grpcWeb.ClientReadableStream<user_password_pb.RecoverOutput>;

  recover(
    request: user_password_pb.RecoverInput,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: user_password_pb.RecoverOutput) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/UserPasswordService/Recover',
        request,
        metadata || {},
        this.methodDescriptorRecover,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/UserPasswordService/Recover',
    request,
    metadata || {},
    this.methodDescriptorRecover);
  }

  methodDescriptorUpdate = new grpcWeb.MethodDescriptor(
    '/UserPasswordService/Update',
    grpcWeb.MethodType.UNARY,
    user_password_pb.UpdateInput,
    google_protobuf_empty_pb.Empty,
    (request: user_password_pb.UpdateInput) => {
      return request.serializeBinary();
    },
    google_protobuf_empty_pb.Empty.deserializeBinary
  );

  update(
    request: user_password_pb.UpdateInput,
    metadata: grpcWeb.Metadata | null): Promise<google_protobuf_empty_pb.Empty>;

  update(
    request: user_password_pb.UpdateInput,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  update(
    request: user_password_pb.UpdateInput,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/UserPasswordService/Update',
        request,
        metadata || {},
        this.methodDescriptorUpdate,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/UserPasswordService/Update',
    request,
    metadata || {},
    this.methodDescriptorUpdate);
  }

  methodDescriptorUpdateByToken = new grpcWeb.MethodDescriptor(
    '/UserPasswordService/UpdateByToken',
    grpcWeb.MethodType.UNARY,
    user_password_pb.UpdateInput,
    google_protobuf_empty_pb.Empty,
    (request: user_password_pb.UpdateInput) => {
      return request.serializeBinary();
    },
    google_protobuf_empty_pb.Empty.deserializeBinary
  );

  updateByToken(
    request: user_password_pb.UpdateInput,
    metadata: grpcWeb.Metadata | null): Promise<google_protobuf_empty_pb.Empty>;

  updateByToken(
    request: user_password_pb.UpdateInput,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  updateByToken(
    request: user_password_pb.UpdateInput,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/UserPasswordService/UpdateByToken',
        request,
        metadata || {},
        this.methodDescriptorUpdateByToken,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/UserPasswordService/UpdateByToken',
    request,
    metadata || {},
    this.methodDescriptorUpdateByToken);
  }

  methodDescriptorDelete = new grpcWeb.MethodDescriptor(
    '/UserPasswordService/Delete',
    grpcWeb.MethodType.UNARY,
    user_password_pb.DeleteInput,
    google_protobuf_empty_pb.Empty,
    (request: user_password_pb.DeleteInput) => {
      return request.serializeBinary();
    },
    google_protobuf_empty_pb.Empty.deserializeBinary
  );

  delete(
    request: user_password_pb.DeleteInput,
    metadata: grpcWeb.Metadata | null): Promise<google_protobuf_empty_pb.Empty>;

  delete(
    request: user_password_pb.DeleteInput,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  delete(
    request: user_password_pb.DeleteInput,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/UserPasswordService/Delete',
        request,
        metadata || {},
        this.methodDescriptorDelete,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/UserPasswordService/Delete',
    request,
    metadata || {},
    this.methodDescriptorDelete);
  }

}

