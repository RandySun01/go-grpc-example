// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: 04clientStream/proto/clientStream.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// 定义发送请求信息
type StreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 定义发送的参数，采用驼峰命名方式，小写加下划线，如：student_name
	// 请求参数
	StreamData string `protobuf:"bytes,1,opt,name=stream_data,json=streamData,proto3" json:"stream_data,omitempty"`
}

func (x *StreamRequest) Reset() {
	*x = StreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file__04clientStream_proto_clientStream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamRequest) ProtoMessage() {}

func (x *StreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file__04clientStream_proto_clientStream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamRequest.ProtoReflect.Descriptor instead.
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return file__04clientStream_proto_clientStream_proto_rawDescGZIP(), []int{0}
}

func (x *StreamRequest) GetStreamData() string {
	if x != nil {
		return x.StreamData
	}
	return ""
}

// 定义流式响应信息
type SimpleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//响应码
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	//响应值
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SimpleResponse) Reset() {
	*x = SimpleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file__04clientStream_proto_clientStream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleResponse) ProtoMessage() {}

func (x *SimpleResponse) ProtoReflect() protoreflect.Message {
	mi := &file__04clientStream_proto_clientStream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleResponse.ProtoReflect.Descriptor instead.
func (*SimpleResponse) Descriptor() ([]byte, []int) {
	return file__04clientStream_proto_clientStream_proto_rawDescGZIP(), []int{1}
}

func (x *SimpleResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SimpleResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File__04clientStream_proto_clientStream_proto protoreflect.FileDescriptor

var file__04clientStream_proto_clientStream_proto_rawDesc = []byte{
	0x0a, 0x27, 0x30, 0x34, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x30, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x44, 0x61,
	0x74, 0x61, 0x22, 0x3a, 0x0a, 0x0e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x4c,
	0x0a, 0x0c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x3c,
	0x0a, 0x09, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file__04clientStream_proto_clientStream_proto_rawDescOnce sync.Once
	file__04clientStream_proto_clientStream_proto_rawDescData = file__04clientStream_proto_clientStream_proto_rawDesc
)

func file__04clientStream_proto_clientStream_proto_rawDescGZIP() []byte {
	file__04clientStream_proto_clientStream_proto_rawDescOnce.Do(func() {
		file__04clientStream_proto_clientStream_proto_rawDescData = protoimpl.X.CompressGZIP(file__04clientStream_proto_clientStream_proto_rawDescData)
	})
	return file__04clientStream_proto_clientStream_proto_rawDescData
}

var file__04clientStream_proto_clientStream_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file__04clientStream_proto_clientStream_proto_goTypes = []interface{}{
	(*StreamRequest)(nil),  // 0: proto.StreamRequest
	(*SimpleResponse)(nil), // 1: proto.SimpleResponse
}
var file__04clientStream_proto_clientStream_proto_depIdxs = []int32{
	0, // 0: proto.StreamClient.RouteList:input_type -> proto.StreamRequest
	1, // 1: proto.StreamClient.RouteList:output_type -> proto.SimpleResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file__04clientStream_proto_clientStream_proto_init() }
func file__04clientStream_proto_clientStream_proto_init() {
	if File__04clientStream_proto_clientStream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file__04clientStream_proto_clientStream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamRequest); i {
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
		file__04clientStream_proto_clientStream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleResponse); i {
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
			RawDescriptor: file__04clientStream_proto_clientStream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file__04clientStream_proto_clientStream_proto_goTypes,
		DependencyIndexes: file__04clientStream_proto_clientStream_proto_depIdxs,
		MessageInfos:      file__04clientStream_proto_clientStream_proto_msgTypes,
	}.Build()
	File__04clientStream_proto_clientStream_proto = out.File
	file__04clientStream_proto_clientStream_proto_rawDesc = nil
	file__04clientStream_proto_clientStream_proto_goTypes = nil
	file__04clientStream_proto_clientStream_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StreamClientClient is the client API for StreamClient service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StreamClientClient interface {
	// 客户端流式rpc，在请求的参数前添加stream
	RouteList(ctx context.Context, opts ...grpc.CallOption) (StreamClient_RouteListClient, error)
}

type streamClientClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamClientClient(cc grpc.ClientConnInterface) StreamClientClient {
	return &streamClientClient{cc}
}

func (c *streamClientClient) RouteList(ctx context.Context, opts ...grpc.CallOption) (StreamClient_RouteListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamClient_serviceDesc.Streams[0], "/proto.StreamClient/RouteList", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamClientRouteListClient{stream}
	return x, nil
}

type StreamClient_RouteListClient interface {
	Send(*StreamRequest) error
	CloseAndRecv() (*SimpleResponse, error)
	grpc.ClientStream
}

type streamClientRouteListClient struct {
	grpc.ClientStream
}

func (x *streamClientRouteListClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamClientRouteListClient) CloseAndRecv() (*SimpleResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SimpleResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamClientServer is the server API for StreamClient service.
type StreamClientServer interface {
	// 客户端流式rpc，在请求的参数前添加stream
	RouteList(StreamClient_RouteListServer) error
}

// UnimplementedStreamClientServer can be embedded to have forward compatible implementations.
type UnimplementedStreamClientServer struct {
}

func (*UnimplementedStreamClientServer) RouteList(StreamClient_RouteListServer) error {
	return status.Errorf(codes.Unimplemented, "method RouteList not implemented")
}

func RegisterStreamClientServer(s *grpc.Server, srv StreamClientServer) {
	s.RegisterService(&_StreamClient_serviceDesc, srv)
}

func _StreamClient_RouteList_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamClientServer).RouteList(&streamClientRouteListServer{stream})
}

type StreamClient_RouteListServer interface {
	SendAndClose(*SimpleResponse) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type streamClientRouteListServer struct {
	grpc.ServerStream
}

func (x *streamClientRouteListServer) SendAndClose(m *SimpleResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamClientRouteListServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _StreamClient_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.StreamClient",
	HandlerType: (*StreamClientServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RouteList",
			Handler:       _StreamClient_RouteList_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "04clientStream/proto/clientStream.proto",
}
