//指定proto版本

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: chatRoom.proto

package chatRoom

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

//登入者
type LoginUserVo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *LoginUserVo) Reset() {
	*x = LoginUserVo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatRoom_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginUserVo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginUserVo) ProtoMessage() {}

func (x *LoginUserVo) ProtoReflect() protoreflect.Message {
	mi := &file_chatRoom_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginUserVo.ProtoReflect.Descriptor instead.
func (*LoginUserVo) Descriptor() ([]byte, []int) {
	return file_chatRoom_proto_rawDescGZIP(), []int{0}
}

func (x *LoginUserVo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

//使用者
type UserVo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UserVo) Reset() {
	*x = UserVo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatRoom_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserVo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserVo) ProtoMessage() {}

func (x *UserVo) ProtoReflect() protoreflect.Message {
	mi := &file_chatRoom_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserVo.ProtoReflect.Descriptor instead.
func (*UserVo) Descriptor() ([]byte, []int) {
	return file_chatRoom_proto_rawDescGZIP(), []int{1}
}

func (x *UserVo) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *UserVo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

//聊天視窗裡的訊息
type MessageVo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *MessageVo) Reset() {
	*x = MessageVo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatRoom_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageVo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageVo) ProtoMessage() {}

func (x *MessageVo) ProtoReflect() protoreflect.Message {
	mi := &file_chatRoom_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageVo.ProtoReflect.Descriptor instead.
func (*MessageVo) Descriptor() ([]byte, []int) {
	return file_chatRoom_proto_rawDescGZIP(), []int{2}
}

func (x *MessageVo) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_chatRoom_proto protoreflect.FileDescriptor

var file_chatRoom_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x68, 0x61, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x63, 0x68, 0x61, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x22, 0x21, 0x0a, 0x0b, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x56, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x32, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x56, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x25, 0x0a, 0x09, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x56, 0x6f, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32, 0x79, 0x0a, 0x0f, 0x43, 0x68, 0x61, 0x74,
	0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x12, 0x15, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x2e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x56, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x56, 0x6f, 0x12, 0x34, 0x0a,
	0x04, 0x43, 0x68, 0x61, 0x74, 0x12, 0x13, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x52, 0x6f, 0x6f, 0x6d,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x56, 0x6f, 0x1a, 0x13, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x56, 0x6f, 0x28,
	0x01, 0x30, 0x01, 0x42, 0x0f, 0x5a, 0x0d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x68, 0x61, 0x74,
	0x52, 0x6f, 0x6f, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chatRoom_proto_rawDescOnce sync.Once
	file_chatRoom_proto_rawDescData = file_chatRoom_proto_rawDesc
)

func file_chatRoom_proto_rawDescGZIP() []byte {
	file_chatRoom_proto_rawDescOnce.Do(func() {
		file_chatRoom_proto_rawDescData = protoimpl.X.CompressGZIP(file_chatRoom_proto_rawDescData)
	})
	return file_chatRoom_proto_rawDescData
}

var file_chatRoom_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_chatRoom_proto_goTypes = []interface{}{
	(*LoginUserVo)(nil), // 0: chatRoom.LoginUserVo
	(*UserVo)(nil),      // 1: chatRoom.UserVo
	(*MessageVo)(nil),   // 2: chatRoom.MessageVo
}
var file_chatRoom_proto_depIdxs = []int32{
	0, // 0: chatRoom.ChatRoomService.Login:input_type -> chatRoom.LoginUserVo
	2, // 1: chatRoom.ChatRoomService.Chat:input_type -> chatRoom.MessageVo
	1, // 2: chatRoom.ChatRoomService.Login:output_type -> chatRoom.UserVo
	2, // 3: chatRoom.ChatRoomService.Chat:output_type -> chatRoom.MessageVo
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chatRoom_proto_init() }
func file_chatRoom_proto_init() {
	if File_chatRoom_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chatRoom_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginUserVo); i {
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
		file_chatRoom_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserVo); i {
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
		file_chatRoom_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageVo); i {
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
			RawDescriptor: file_chatRoom_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chatRoom_proto_goTypes,
		DependencyIndexes: file_chatRoom_proto_depIdxs,
		MessageInfos:      file_chatRoom_proto_msgTypes,
	}.Build()
	File_chatRoom_proto = out.File
	file_chatRoom_proto_rawDesc = nil
	file_chatRoom_proto_goTypes = nil
	file_chatRoom_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChatRoomServiceClient is the client API for ChatRoomService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatRoomServiceClient interface {
	//登入, UserVo.Token有值時表示登入成功
	Login(ctx context.Context, in *LoginUserVo, opts ...grpc.CallOption) (*UserVo, error)
	//傳入送出的聊天訊息, 回傳聊天紀錄, 需在context裡指定值, key = "token", value = UserVo.token
	Chat(ctx context.Context, opts ...grpc.CallOption) (ChatRoomService_ChatClient, error)
}

type chatRoomServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatRoomServiceClient(cc grpc.ClientConnInterface) ChatRoomServiceClient {
	return &chatRoomServiceClient{cc}
}

func (c *chatRoomServiceClient) Login(ctx context.Context, in *LoginUserVo, opts ...grpc.CallOption) (*UserVo, error) {
	out := new(UserVo)
	err := c.cc.Invoke(ctx, "/chatRoom.ChatRoomService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatRoomServiceClient) Chat(ctx context.Context, opts ...grpc.CallOption) (ChatRoomService_ChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChatRoomService_serviceDesc.Streams[0], "/chatRoom.ChatRoomService/Chat", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatRoomServiceChatClient{stream}
	return x, nil
}

type ChatRoomService_ChatClient interface {
	Send(*MessageVo) error
	Recv() (*MessageVo, error)
	grpc.ClientStream
}

type chatRoomServiceChatClient struct {
	grpc.ClientStream
}

func (x *chatRoomServiceChatClient) Send(m *MessageVo) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatRoomServiceChatClient) Recv() (*MessageVo, error) {
	m := new(MessageVo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatRoomServiceServer is the server API for ChatRoomService service.
type ChatRoomServiceServer interface {
	//登入, UserVo.Token有值時表示登入成功
	Login(context.Context, *LoginUserVo) (*UserVo, error)
	//傳入送出的聊天訊息, 回傳聊天紀錄, 需在context裡指定值, key = "token", value = UserVo.token
	Chat(ChatRoomService_ChatServer) error
}

// UnimplementedChatRoomServiceServer can be embedded to have forward compatible implementations.
type UnimplementedChatRoomServiceServer struct {
}

func (*UnimplementedChatRoomServiceServer) Login(context.Context, *LoginUserVo) (*UserVo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedChatRoomServiceServer) Chat(ChatRoomService_ChatServer) error {
	return status.Errorf(codes.Unimplemented, "method Chat not implemented")
}

func RegisterChatRoomServiceServer(s *grpc.Server, srv ChatRoomServiceServer) {
	s.RegisterService(&_ChatRoomService_serviceDesc, srv)
}

func _ChatRoomService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserVo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatRoomServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatRoom.ChatRoomService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatRoomServiceServer).Login(ctx, req.(*LoginUserVo))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatRoomService_Chat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatRoomServiceServer).Chat(&chatRoomServiceChatServer{stream})
}

type ChatRoomService_ChatServer interface {
	Send(*MessageVo) error
	Recv() (*MessageVo, error)
	grpc.ServerStream
}

type chatRoomServiceChatServer struct {
	grpc.ServerStream
}

func (x *chatRoomServiceChatServer) Send(m *MessageVo) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatRoomServiceChatServer) Recv() (*MessageVo, error) {
	m := new(MessageVo)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ChatRoomService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chatRoom.ChatRoomService",
	HandlerType: (*ChatRoomServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _ChatRoomService_Login_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Chat",
			Handler:       _ChatRoomService_Chat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "chatRoom.proto",
}
