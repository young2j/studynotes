// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: gateway/v1/gateway.proto

package gatewayv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_v1_gateway_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_v1_gateway_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_gateway_v1_gateway_proto_rawDescGZIP(), []int{0}
}

func (x *PingRequest) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_v1_gateway_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_v1_gateway_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_gateway_v1_gateway_proto_rawDescGZIP(), []int{1}
}

func (x *PingResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type DetectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DetectRequest) Reset() {
	*x = DetectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_v1_gateway_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetectRequest) ProtoMessage() {}

func (x *DetectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_v1_gateway_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetectRequest.ProtoReflect.Descriptor instead.
func (*DetectRequest) Descriptor() ([]byte, []int) {
	return file_gateway_v1_gateway_proto_rawDescGZIP(), []int{2}
}

func (x *DetectRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DetectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DetectResponse) Reset() {
	*x = DetectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_v1_gateway_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetectResponse) ProtoMessage() {}

func (x *DetectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_v1_gateway_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetectResponse.ProtoReflect.Descriptor instead.
func (*DetectResponse) Descriptor() ([]byte, []int) {
	return file_gateway_v1_gateway_proto_rawDescGZIP(), []int{3}
}

func (x *DetectResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_gateway_v1_gateway_proto protoreflect.FileDescriptor

var file_gateway_v1_gateway_proto_rawDesc = []byte{
	0x0a, 0x18, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x0b, 0x50, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x18, 0x05, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x22, 0x2b, 0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x05, 0x18, 0x0a, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x22, 0x28, 0x0a, 0x0d, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x1a, 0x02, 0x28, 0x0a, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2b, 0x0a, 0x0e, 0x44,
	0x65, 0x74, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x1a, 0x04,
	0x10, 0x0a, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x32, 0xad, 0x01, 0x0a, 0x0c, 0x50, 0x72, 0x6f,
	0x62, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5c, 0x0a, 0x04, 0x50, 0x69, 0x6e,
	0x67, 0x12, 0x17, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22, 0x16, 0x2f, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x2f,
	0x70, 0x69, 0x6e, 0x67, 0x3a, 0x01, 0x2a, 0x12, 0x3f, 0x0a, 0x06, 0x44, 0x65, 0x74, 0x65, 0x63,
	0x74, 0x12, 0x19, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x74, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xe3, 0x01, 0x5a, 0x32, 0x67, 0x72, 0x70,
	0x63, 0x2d, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x76, 0x31, 0x92,
	0x41, 0xab, 0x01, 0x12, 0x1a, 0x0a, 0x12, 0x53, 0x57, 0x41, 0x47, 0x47, 0x45, 0x52, 0x20, 0x41,
	0x50, 0x49, 0x20, 0xe6, 0x96, 0x87, 0xe6, 0xa1, 0xa3, 0x32, 0x04, 0x76, 0x31, 0x2e, 0x30, 0x5a,
	0x8c, 0x01, 0x0a, 0x89, 0x01, 0x0a, 0x06, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x12, 0x7f, 0x08,
	0x03, 0x12, 0x0e, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x20, 0x7b, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x7d, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x20, 0x02, 0x28, 0x01, 0x32, 0x2b, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x68, 0x6f, 0x73, 0x74, 0x3a, 0x35, 0x32, 0x30, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x2f, 0x76, 0x31, 0x2f, 0x6a, 0x77, 0x74, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x3a, 0x2b, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68,
	0x6f, 0x73, 0x74, 0x3a, 0x35, 0x32, 0x30, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31,
	0x2f, 0x6a, 0x77, 0x74, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gateway_v1_gateway_proto_rawDescOnce sync.Once
	file_gateway_v1_gateway_proto_rawDescData = file_gateway_v1_gateway_proto_rawDesc
)

func file_gateway_v1_gateway_proto_rawDescGZIP() []byte {
	file_gateway_v1_gateway_proto_rawDescOnce.Do(func() {
		file_gateway_v1_gateway_proto_rawDescData = protoimpl.X.CompressGZIP(file_gateway_v1_gateway_proto_rawDescData)
	})
	return file_gateway_v1_gateway_proto_rawDescData
}

var file_gateway_v1_gateway_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_gateway_v1_gateway_proto_goTypes = []interface{}{
	(*PingRequest)(nil),    // 0: gateway.v1.PingRequest
	(*PingResponse)(nil),   // 1: gateway.v1.PingResponse
	(*DetectRequest)(nil),  // 2: gateway.v1.DetectRequest
	(*DetectResponse)(nil), // 3: gateway.v1.DetectResponse
}
var file_gateway_v1_gateway_proto_depIdxs = []int32{
	0, // 0: gateway.v1.ProbeService.Ping:input_type -> gateway.v1.PingRequest
	2, // 1: gateway.v1.ProbeService.Detect:input_type -> gateway.v1.DetectRequest
	1, // 2: gateway.v1.ProbeService.Ping:output_type -> gateway.v1.PingResponse
	3, // 3: gateway.v1.ProbeService.Detect:output_type -> gateway.v1.DetectResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gateway_v1_gateway_proto_init() }
func file_gateway_v1_gateway_proto_init() {
	if File_gateway_v1_gateway_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gateway_v1_gateway_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gateway_v1_gateway_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gateway_v1_gateway_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetectRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gateway_v1_gateway_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetectResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gateway_v1_gateway_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gateway_v1_gateway_proto_goTypes,
		DependencyIndexes: file_gateway_v1_gateway_proto_depIdxs,
		MessageInfos:      file_gateway_v1_gateway_proto_msgTypes,
	}.Build()
	File_gateway_v1_gateway_proto = out.File
	file_gateway_v1_gateway_proto_rawDesc = nil
	file_gateway_v1_gateway_proto_goTypes = nil
	file_gateway_v1_gateway_proto_depIdxs = nil
}