/**
 * @fileoverview gRPC-Web generated client stub for 
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as auth_pb from './auth_pb';
import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';


export class AuthServiceClient {
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

  methodDescriptorLogin = new grpcWeb.MethodDescriptor(
    '/AuthService/Login',
    grpcWeb.MethodType.UNARY,
    auth_pb.LoginInput,
    auth_pb.AuthOutput,
    (request: auth_pb.LoginInput) => {
      return request.serializeBinary();
    },
    auth_pb.AuthOutput.deserializeBinary
  );

  login(
    request: auth_pb.LoginInput,
    metadata: grpcWeb.Metadata | null): Promise<auth_pb.AuthOutput>;

  login(
    request: auth_pb.LoginInput,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: auth_pb.AuthOutput) => void): grpcWeb.ClientReadableStream<auth_pb.AuthOutput>;

  login(
    request: auth_pb.LoginInput,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: auth_pb.AuthOutput) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/AuthService/Login',
        request,
        metadata || {},
        this.methodDescriptorLogin,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/AuthService/Login',
    request,
    metadata || {},
    this.methodDescriptorLogin);
  }

  methodDescriptorRefresh = new grpcWeb.MethodDescriptor(
    '/AuthService/Refresh',
    grpcWeb.MethodType.UNARY,
    auth_pb.TokenInput,
    auth_pb.AuthOutput,
    (request: auth_pb.TokenInput) => {
      return request.serializeBinary();
    },
    auth_pb.AuthOutput.deserializeBinary
  );

  refresh(
    request: auth_pb.TokenInput,
    metadata: grpcWeb.Metadata | null): Promise<auth_pb.AuthOutput>;

  refresh(
    request: auth_pb.TokenInput,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: auth_pb.AuthOutput) => void): grpcWeb.ClientReadableStream<auth_pb.AuthOutput>;

  refresh(
    request: auth_pb.TokenInput,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: auth_pb.AuthOutput) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/AuthService/Refresh',
        request,
        metadata || {},
        this.methodDescriptorRefresh,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/AuthService/Refresh',
    request,
    metadata || {},
    this.methodDescriptorRefresh);
  }

  methodDescriptorInvalidate = new grpcWeb.MethodDescriptor(
    '/AuthService/Invalidate',
    grpcWeb.MethodType.UNARY,
    auth_pb.TokenInput,
    google_protobuf_empty_pb.Empty,
    (request: auth_pb.TokenInput) => {
      return request.serializeBinary();
    },
    google_protobuf_empty_pb.Empty.deserializeBinary
  );

  invalidate(
    request: auth_pb.TokenInput,
    metadata: grpcWeb.Metadata | null): Promise<google_protobuf_empty_pb.Empty>;

  invalidate(
    request: auth_pb.TokenInput,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  invalidate(
    request: auth_pb.TokenInput,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/AuthService/Invalidate',
        request,
        metadata || {},
        this.methodDescriptorInvalidate,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/AuthService/Invalidate',
    request,
    metadata || {},
    this.methodDescriptorInvalidate);
  }

  methodDescriptorInvalidateAll = new grpcWeb.MethodDescriptor(
    '/AuthService/InvalidateAll',
    grpcWeb.MethodType.UNARY,
    auth_pb.SessionsInput,
    google_protobuf_empty_pb.Empty,
    (request: auth_pb.SessionsInput) => {
      return request.serializeBinary();
    },
    google_protobuf_empty_pb.Empty.deserializeBinary
  );

  invalidateAll(
    request: auth_pb.SessionsInput,
    metadata: grpcWeb.Metadata | null): Promise<google_protobuf_empty_pb.Empty>;

  invalidateAll(
    request: auth_pb.SessionsInput,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  invalidateAll(
    request: auth_pb.SessionsInput,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/AuthService/InvalidateAll',
        request,
        metadata || {},
        this.methodDescriptorInvalidateAll,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/AuthService/InvalidateAll',
    request,
    metadata || {},
    this.methodDescriptorInvalidateAll);
  }

  methodDescriptorActiveSessions = new grpcWeb.MethodDescriptor(
    '/AuthService/ActiveSessions',
    grpcWeb.MethodType.UNARY,
    auth_pb.SessionsInput,
    auth_pb.SessionsOutput,
    (request: auth_pb.SessionsInput) => {
      return request.serializeBinary();
    },
    auth_pb.SessionsOutput.deserializeBinary
  );

  activeSessions(
    request: auth_pb.SessionsInput,
    metadata: grpcWeb.Metadata | null): Promise<auth_pb.SessionsOutput>;

  activeSessions(
    request: auth_pb.SessionsInput,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: auth_pb.SessionsOutput) => void): grpcWeb.ClientReadableStream<auth_pb.SessionsOutput>;

  activeSessions(
    request: auth_pb.SessionsInput,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: auth_pb.SessionsOutput) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/AuthService/ActiveSessions',
        request,
        metadata || {},
        this.methodDescriptorActiveSessions,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/AuthService/ActiveSessions',
    request,
    metadata || {},
    this.methodDescriptorActiveSessions);
  }

  methodDescriptorIsValid = new grpcWeb.MethodDescriptor(
    '/AuthService/IsValid',
    grpcWeb.MethodType.UNARY,
    auth_pb.TokenInput,
    google_protobuf_empty_pb.Empty,
    (request: auth_pb.TokenInput) => {
      return request.serializeBinary();
    },
    google_protobuf_empty_pb.Empty.deserializeBinary
  );

  isValid(
    request: auth_pb.TokenInput,
    metadata: grpcWeb.Metadata | null): Promise<google_protobuf_empty_pb.Empty>;

  isValid(
    request: auth_pb.TokenInput,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  isValid(
    request: auth_pb.TokenInput,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/AuthService/IsValid',
        request,
        metadata || {},
        this.methodDescriptorIsValid,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/AuthService/IsValid',
    request,
    metadata || {},
    this.methodDescriptorIsValid);
  }

}

