// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var jwt_v1_jwt_pb = require('../../jwt/v1/jwt_pb.js');
var google_api_annotations_pb = require('../../google/api/annotations_pb.js');

function serialize_jwt_v1_GetTokenRequest(arg) {
  if (!(arg instanceof jwt_v1_jwt_pb.GetTokenRequest)) {
    throw new Error('Expected argument of type jwt.v1.GetTokenRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_jwt_v1_GetTokenRequest(buffer_arg) {
  return jwt_v1_jwt_pb.GetTokenRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_jwt_v1_GetTokenResponse(arg) {
  if (!(arg instanceof jwt_v1_jwt_pb.GetTokenResponse)) {
    throw new Error('Expected argument of type jwt.v1.GetTokenResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_jwt_v1_GetTokenResponse(buffer_arg) {
  return jwt_v1_jwt_pb.GetTokenResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_jwt_v1_RefreshTokenRequest(arg) {
  if (!(arg instanceof jwt_v1_jwt_pb.RefreshTokenRequest)) {
    throw new Error('Expected argument of type jwt.v1.RefreshTokenRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_jwt_v1_RefreshTokenRequest(buffer_arg) {
  return jwt_v1_jwt_pb.RefreshTokenRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_jwt_v1_RefreshTokenResponse(arg) {
  if (!(arg instanceof jwt_v1_jwt_pb.RefreshTokenResponse)) {
    throw new Error('Expected argument of type jwt.v1.RefreshTokenResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_jwt_v1_RefreshTokenResponse(buffer_arg) {
  return jwt_v1_jwt_pb.RefreshTokenResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var JWTServiceService = exports.JWTServiceService = {
  getToken: {
    path: '/jwt.v1.JWTService/GetToken',
    requestStream: false,
    responseStream: false,
    requestType: jwt_v1_jwt_pb.GetTokenRequest,
    responseType: jwt_v1_jwt_pb.GetTokenResponse,
    requestSerialize: serialize_jwt_v1_GetTokenRequest,
    requestDeserialize: deserialize_jwt_v1_GetTokenRequest,
    responseSerialize: serialize_jwt_v1_GetTokenResponse,
    responseDeserialize: deserialize_jwt_v1_GetTokenResponse,
  },
  refreshToken: {
    path: '/jwt.v1.JWTService/RefreshToken',
    requestStream: false,
    responseStream: false,
    requestType: jwt_v1_jwt_pb.RefreshTokenRequest,
    responseType: jwt_v1_jwt_pb.RefreshTokenResponse,
    requestSerialize: serialize_jwt_v1_RefreshTokenRequest,
    requestDeserialize: deserialize_jwt_v1_RefreshTokenRequest,
    responseSerialize: serialize_jwt_v1_RefreshTokenResponse,
    responseDeserialize: deserialize_jwt_v1_RefreshTokenResponse,
  },
};

exports.JWTServiceClient = grpc.makeGenericClientConstructor(JWTServiceService);
