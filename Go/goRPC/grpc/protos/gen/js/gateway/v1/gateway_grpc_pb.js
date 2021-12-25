// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var gateway_v1_gateway_pb = require('../../gateway/v1/gateway_pb.js');
var google_api_annotations_pb = require('../../google/api/annotations_pb.js');
var openapiv2_annotations_pb = require('../../openapiv2/annotations_pb.js');
var validate_validate_pb = require('../../validate/validate_pb.js');

function serialize_gateway_v1_DetectRequest(arg) {
  if (!(arg instanceof gateway_v1_gateway_pb.DetectRequest)) {
    throw new Error('Expected argument of type gateway.v1.DetectRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_gateway_v1_DetectRequest(buffer_arg) {
  return gateway_v1_gateway_pb.DetectRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_gateway_v1_DetectResponse(arg) {
  if (!(arg instanceof gateway_v1_gateway_pb.DetectResponse)) {
    throw new Error('Expected argument of type gateway.v1.DetectResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_gateway_v1_DetectResponse(buffer_arg) {
  return gateway_v1_gateway_pb.DetectResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_gateway_v1_PingRequest(arg) {
  if (!(arg instanceof gateway_v1_gateway_pb.PingRequest)) {
    throw new Error('Expected argument of type gateway.v1.PingRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_gateway_v1_PingRequest(buffer_arg) {
  return gateway_v1_gateway_pb.PingRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_gateway_v1_PingResponse(arg) {
  if (!(arg instanceof gateway_v1_gateway_pb.PingResponse)) {
    throw new Error('Expected argument of type gateway.v1.PingResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_gateway_v1_PingResponse(buffer_arg) {
  return gateway_v1_gateway_pb.PingResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var ProbeServiceService = exports.ProbeServiceService = {
  ping: {
    path: '/gateway.v1.ProbeService/Ping',
    requestStream: false,
    responseStream: false,
    requestType: gateway_v1_gateway_pb.PingRequest,
    responseType: gateway_v1_gateway_pb.PingResponse,
    requestSerialize: serialize_gateway_v1_PingRequest,
    requestDeserialize: deserialize_gateway_v1_PingRequest,
    responseSerialize: serialize_gateway_v1_PingResponse,
    responseDeserialize: deserialize_gateway_v1_PingResponse,
  },
  detect: {
    path: '/gateway.v1.ProbeService/Detect',
    requestStream: false,
    responseStream: false,
    requestType: gateway_v1_gateway_pb.DetectRequest,
    responseType: gateway_v1_gateway_pb.DetectResponse,
    requestSerialize: serialize_gateway_v1_DetectRequest,
    requestDeserialize: deserialize_gateway_v1_DetectRequest,
    responseSerialize: serialize_gateway_v1_DetectResponse,
    responseDeserialize: deserialize_gateway_v1_DetectResponse,
  },
};

exports.ProbeServiceClient = grpc.makeGenericClientConstructor(ProbeServiceService);
