// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: internal/subscribe/sub.proto

package subscribe

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SubscrRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sport  []string `protobuf:"bytes,1,rep,name=sport,proto3" json:"sport,omitempty"`
	Repeat string   `protobuf:"bytes,2,opt,name=repeat,proto3" json:"repeat,omitempty"`
}

func (x *SubscrRequest) Reset() {
	*x = SubscrRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_subscribe_sub_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscrRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscrRequest) ProtoMessage() {}

func (x *SubscrRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_subscribe_sub_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscrRequest.ProtoReflect.Descriptor instead.
func (*SubscrRequest) Descriptor() ([]byte, []int) {
	return file_internal_subscribe_sub_proto_rawDescGZIP(), []int{0}
}

func (x *SubscrRequest) GetSport() []string {
	if x != nil {
		return x.Sport
	}
	return nil
}

func (x *SubscrRequest) GetRepeat() string {
	if x != nil {
		return x.Repeat
	}
	return ""
}

type SubscrResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*Result `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *SubscrResponse) Reset() {
	*x = SubscrResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_subscribe_sub_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscrResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscrResponse) ProtoMessage() {}

func (x *SubscrResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_subscribe_sub_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscrResponse.ProtoReflect.Descriptor instead.
func (*SubscrResponse) Descriptor() ([]byte, []int) {
	return file_internal_subscribe_sub_proto_rawDescGZIP(), []int{1}
}

func (x *SubscrResponse) GetResults() []*Result {
	if x != nil {
		return x.Results
	}
	return nil
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sport string  `protobuf:"bytes,1,opt,name=sport,proto3" json:"sport,omitempty"`
	Delta float64 `protobuf:"fixed64,2,opt,name=delta,proto3" json:"delta,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_subscribe_sub_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_internal_subscribe_sub_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_internal_subscribe_sub_proto_rawDescGZIP(), []int{2}
}

func (x *Result) GetSport() string {
	if x != nil {
		return x.Sport
	}
	return ""
}

func (x *Result) GetDelta() float64 {
	if x != nil {
		return x.Delta
	}
	return 0
}

var File_internal_subscribe_sub_proto protoreflect.FileDescriptor

var file_internal_subscribe_sub_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x2f, 0x73, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x22, 0x3d, 0x0a, 0x0d, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x22, 0x3d, 0x0a, 0x0e, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x34, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x6c, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x32, 0x62, 0x0a,
	0x0d, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51,
	0x0a, 0x16, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4f, 0x6e, 0x53, 0x70, 0x6f,
	0x72, 0x74, 0x73, 0x4c, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x2e, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30,
	0x01, 0x42, 0x14, 0x5a, 0x12, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_subscribe_sub_proto_rawDescOnce sync.Once
	file_internal_subscribe_sub_proto_rawDescData = file_internal_subscribe_sub_proto_rawDesc
)

func file_internal_subscribe_sub_proto_rawDescGZIP() []byte {
	file_internal_subscribe_sub_proto_rawDescOnce.Do(func() {
		file_internal_subscribe_sub_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_subscribe_sub_proto_rawDescData)
	})
	return file_internal_subscribe_sub_proto_rawDescData
}

var file_internal_subscribe_sub_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_internal_subscribe_sub_proto_goTypes = []interface{}{
	(*SubscrRequest)(nil),  // 0: subscribe.SubscrRequest
	(*SubscrResponse)(nil), // 1: subscribe.SubscrResponse
	(*Result)(nil),         // 2: subscribe.Result
}
var file_internal_subscribe_sub_proto_depIdxs = []int32{
	2, // 0: subscribe.SubscrResponse.results:type_name -> subscribe.Result
	0, // 1: subscribe.SubscrService.SubscribeOnSportsLines:input_type -> subscribe.SubscrRequest
	1, // 2: subscribe.SubscrService.SubscribeOnSportsLines:output_type -> subscribe.SubscrResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_subscribe_sub_proto_init() }
func file_internal_subscribe_sub_proto_init() {
	if File_internal_subscribe_sub_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_subscribe_sub_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscrRequest); i {
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
		file_internal_subscribe_sub_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscrResponse); i {
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
		file_internal_subscribe_sub_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
			RawDescriptor: file_internal_subscribe_sub_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_subscribe_sub_proto_goTypes,
		DependencyIndexes: file_internal_subscribe_sub_proto_depIdxs,
		MessageInfos:      file_internal_subscribe_sub_proto_msgTypes,
	}.Build()
	File_internal_subscribe_sub_proto = out.File
	file_internal_subscribe_sub_proto_rawDesc = nil
	file_internal_subscribe_sub_proto_goTypes = nil
	file_internal_subscribe_sub_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SubscrServiceClient is the client API for SubscrService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SubscrServiceClient interface {
	SubscribeOnSportsLines(ctx context.Context, opts ...grpc.CallOption) (SubscrService_SubscribeOnSportsLinesClient, error)
}

type subscrServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSubscrServiceClient(cc grpc.ClientConnInterface) SubscrServiceClient {
	return &subscrServiceClient{cc}
}

func (c *subscrServiceClient) SubscribeOnSportsLines(ctx context.Context, opts ...grpc.CallOption) (SubscrService_SubscribeOnSportsLinesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SubscrService_serviceDesc.Streams[0], "/subscribe.SubscrService/SubscribeOnSportsLines", opts...)
	if err != nil {
		return nil, err
	}
	x := &subscrServiceSubscribeOnSportsLinesClient{stream}
	return x, nil
}

type SubscrService_SubscribeOnSportsLinesClient interface {
	Send(*SubscrRequest) error
	Recv() (*SubscrResponse, error)
	grpc.ClientStream
}

type subscrServiceSubscribeOnSportsLinesClient struct {
	grpc.ClientStream
}

func (x *subscrServiceSubscribeOnSportsLinesClient) Send(m *SubscrRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *subscrServiceSubscribeOnSportsLinesClient) Recv() (*SubscrResponse, error) {
	m := new(SubscrResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SubscrServiceServer is the server API for SubscrService service.
type SubscrServiceServer interface {
	SubscribeOnSportsLines(SubscrService_SubscribeOnSportsLinesServer) error
}

// UnimplementedSubscrServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSubscrServiceServer struct {
}

func (*UnimplementedSubscrServiceServer) SubscribeOnSportsLines(SubscrService_SubscribeOnSportsLinesServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeOnSportsLines not implemented")
}

func RegisterSubscrServiceServer(s *grpc.Server, srv SubscrServiceServer) {
	s.RegisterService(&_SubscrService_serviceDesc, srv)
}

func _SubscrService_SubscribeOnSportsLines_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SubscrServiceServer).SubscribeOnSportsLines(&subscrServiceSubscribeOnSportsLinesServer{stream})
}

type SubscrService_SubscribeOnSportsLinesServer interface {
	Send(*SubscrResponse) error
	Recv() (*SubscrRequest, error)
	grpc.ServerStream
}

type subscrServiceSubscribeOnSportsLinesServer struct {
	grpc.ServerStream
}

func (x *subscrServiceSubscribeOnSportsLinesServer) Send(m *SubscrResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *subscrServiceSubscribeOnSportsLinesServer) Recv() (*SubscrRequest, error) {
	m := new(SubscrRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _SubscrService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "subscribe.SubscrService",
	HandlerType: (*SubscrServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeOnSportsLines",
			Handler:       _SubscrService_SubscribeOnSportsLines_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "internal/subscribe/sub.proto",
}
