// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var hello_v1_hello_pb = require('../../hello/v1/hello_pb.js');
var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');

function serialize_hello_v1_SayHelloRequest(arg) {
  if (!(arg instanceof hello_v1_hello_pb.SayHelloRequest)) {
    throw new Error('Expected argument of type hello.v1.SayHelloRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_hello_v1_SayHelloRequest(buffer_arg) {
  return hello_v1_hello_pb.SayHelloRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_hello_v1_SayHelloResponse(arg) {
  if (!(arg instanceof hello_v1_hello_pb.SayHelloResponse)) {
    throw new Error('Expected argument of type hello.v1.SayHelloResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_hello_v1_SayHelloResponse(buffer_arg) {
  return hello_v1_hello_pb.SayHelloResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// ------------定义RPC服务接口service--------------
// 1. protocol buffer编译器会根据所选择的不同语言生成对应的服务接口代码
// 2. 接口方法需要定义请求参数HelloRequest以及返回参数HelloResponse
var HelloServiceService = exports.HelloServiceService = {
  sayHello: {
    path: '/hello.v1.HelloService/SayHello',
    requestStream: false,
    responseStream: false,
    requestType: hello_v1_hello_pb.SayHelloRequest,
    responseType: hello_v1_hello_pb.SayHelloResponse,
    requestSerialize: serialize_hello_v1_SayHelloRequest,
    requestDeserialize: deserialize_hello_v1_SayHelloRequest,
    responseSerialize: serialize_hello_v1_SayHelloResponse,
    responseDeserialize: deserialize_hello_v1_SayHelloResponse,
  },
};

exports.HelloServiceClient = grpc.makeGenericClientConstructor(HelloServiceService);
