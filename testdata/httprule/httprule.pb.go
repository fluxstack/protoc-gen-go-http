// Code generated by protoc-gen-go. DO NOT EDIT.
// source: httprule/httprule.proto

package httprulepb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type UpdateMessageRequest struct {
	MessageId            string   `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	Message              *Message `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateMessageRequest) Reset()         { *m = UpdateMessageRequest{} }
func (m *UpdateMessageRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateMessageRequest) ProtoMessage()    {}
func (*UpdateMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_72d547f0661199d3, []int{0}
}

func (m *UpdateMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateMessageRequest.Unmarshal(m, b)
}
func (m *UpdateMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateMessageRequest.Marshal(b, m, deterministic)
}
func (m *UpdateMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateMessageRequest.Merge(m, src)
}
func (m *UpdateMessageRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateMessageRequest.Size(m)
}
func (m *UpdateMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateMessageRequest proto.InternalMessageInfo

func (m *UpdateMessageRequest) GetMessageId() string {
	if m != nil {
		return m.MessageId
	}
	return ""
}

func (m *UpdateMessageRequest) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

type SubFieldMessageRequest struct {
	MessageId            string                             `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	Sub                  *SubFieldMessageRequest_SubMessage `protobuf:"bytes,2,opt,name=sub,proto3" json:"sub,omitempty"`
	Text                 string                             `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *SubFieldMessageRequest) Reset()         { *m = SubFieldMessageRequest{} }
func (m *SubFieldMessageRequest) String() string { return proto.CompactTextString(m) }
func (*SubFieldMessageRequest) ProtoMessage()    {}
func (*SubFieldMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_72d547f0661199d3, []int{1}
}

func (m *SubFieldMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubFieldMessageRequest.Unmarshal(m, b)
}
func (m *SubFieldMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubFieldMessageRequest.Marshal(b, m, deterministic)
}
func (m *SubFieldMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubFieldMessageRequest.Merge(m, src)
}
func (m *SubFieldMessageRequest) XXX_Size() int {
	return xxx_messageInfo_SubFieldMessageRequest.Size(m)
}
func (m *SubFieldMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubFieldMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubFieldMessageRequest proto.InternalMessageInfo

func (m *SubFieldMessageRequest) GetMessageId() string {
	if m != nil {
		return m.MessageId
	}
	return ""
}

func (m *SubFieldMessageRequest) GetSub() *SubFieldMessageRequest_SubMessage {
	if m != nil {
		return m.Sub
	}
	return nil
}

func (m *SubFieldMessageRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type SubFieldMessageRequest_SubMessage struct {
	Subfield             string   `protobuf:"bytes,1,opt,name=subfield,proto3" json:"subfield,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubFieldMessageRequest_SubMessage) Reset()         { *m = SubFieldMessageRequest_SubMessage{} }
func (m *SubFieldMessageRequest_SubMessage) String() string { return proto.CompactTextString(m) }
func (*SubFieldMessageRequest_SubMessage) ProtoMessage()    {}
func (*SubFieldMessageRequest_SubMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_72d547f0661199d3, []int{1, 0}
}

func (m *SubFieldMessageRequest_SubMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubFieldMessageRequest_SubMessage.Unmarshal(m, b)
}
func (m *SubFieldMessageRequest_SubMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubFieldMessageRequest_SubMessage.Marshal(b, m, deterministic)
}
func (m *SubFieldMessageRequest_SubMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubFieldMessageRequest_SubMessage.Merge(m, src)
}
func (m *SubFieldMessageRequest_SubMessage) XXX_Size() int {
	return xxx_messageInfo_SubFieldMessageRequest_SubMessage.Size(m)
}
func (m *SubFieldMessageRequest_SubMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SubFieldMessageRequest_SubMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SubFieldMessageRequest_SubMessage proto.InternalMessageInfo

func (m *SubFieldMessageRequest_SubMessage) GetSubfield() string {
	if m != nil {
		return m.Subfield
	}
	return ""
}

type Message struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_72d547f0661199d3, []int{2}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*UpdateMessageRequest)(nil), "httprule.UpdateMessageRequest")
	proto.RegisterType((*SubFieldMessageRequest)(nil), "httprule.SubFieldMessageRequest")
	proto.RegisterType((*SubFieldMessageRequest_SubMessage)(nil), "httprule.SubFieldMessageRequest.SubMessage")
	proto.RegisterType((*Message)(nil), "httprule.Message")
}

func init() { proto.RegisterFile("httprule/httprule.proto", fileDescriptor_72d547f0661199d3) }

var fileDescriptor_72d547f0661199d3 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcf, 0x28, 0x29, 0x29,
	0x28, 0x2a, 0xcd, 0x49, 0xd5, 0x87, 0x31, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x38, 0x60,
	0x7c, 0x29, 0x99, 0xf4, 0xfc, 0xfc, 0xf4, 0x9c, 0x54, 0xfd, 0xc4, 0x82, 0x4c, 0xfd, 0xc4, 0xbc,
	0xbc, 0xfc, 0x92, 0xc4, 0x92, 0xcc, 0xfc, 0xbc, 0x62, 0x88, 0x3a, 0xa5, 0x24, 0x2e, 0x91, 0xd0,
	0x82, 0x94, 0xc4, 0x92, 0x54, 0xdf, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0xd4, 0xa0, 0xd4, 0xc2, 0xd2,
	0xd4, 0xe2, 0x12, 0x21, 0x59, 0x2e, 0xae, 0x5c, 0x88, 0x48, 0x7c, 0x66, 0x8a, 0x04, 0xa3, 0x02,
	0xa3, 0x06, 0x67, 0x10, 0x27, 0x54, 0xc4, 0x33, 0x45, 0x48, 0x9b, 0x8b, 0x1d, 0xca, 0x91, 0x60,
	0x52, 0x60, 0xd4, 0xe0, 0x36, 0x12, 0xd4, 0x83, 0x3b, 0x00, 0x66, 0x12, 0x4c, 0x85, 0xd2, 0x16,
	0x46, 0x2e, 0xb1, 0xe0, 0xd2, 0x24, 0xb7, 0xcc, 0xd4, 0x9c, 0x14, 0xd2, 0xac, 0xb1, 0xe5, 0x62,
	0x2e, 0x2e, 0x4d, 0x82, 0x5a, 0xa1, 0x8d, 0xb0, 0x02, 0xbb, 0x69, 0x20, 0x61, 0x98, 0x08, 0x48,
	0x9f, 0x90, 0x10, 0x17, 0x4b, 0x49, 0x6a, 0x45, 0x89, 0x04, 0x33, 0xd8, 0x5c, 0x30, 0x5b, 0x4a,
	0x83, 0x8b, 0x0b, 0xa1, 0x4c, 0x48, 0x8a, 0x8b, 0xa3, 0xb8, 0x34, 0x29, 0x0d, 0x64, 0x16, 0xd4,
	0x76, 0x38, 0x5f, 0x49, 0x96, 0x8b, 0x1d, 0xa6, 0x0c, 0x66, 0x10, 0x23, 0xc2, 0x20, 0xa3, 0x4f,
	0x8c, 0x5c, 0x9c, 0x10, 0xf9, 0xcc, 0xbc, 0x74, 0xa1, 0x0c, 0x2e, 0x5e, 0x94, 0x70, 0x14, 0x92,
	0x43, 0xb8, 0x16, 0x5b, 0x00, 0x4b, 0x61, 0x06, 0x98, 0x92, 0x4a, 0xd3, 0xe5, 0x27, 0x93, 0x99,
	0xe4, 0xa4, 0x24, 0xf5, 0xcb, 0x0c, 0xf5, 0xa1, 0x81, 0x50, 0xac, 0x5f, 0x8d, 0x08, 0xa0, 0x5a,
	0x2b, 0x46, 0x2d, 0xa1, 0x6a, 0x2e, 0x7e, 0x34, 0xef, 0x0b, 0x29, 0x10, 0x0a, 0x19, 0x6c, 0xb6,
	0x19, 0x83, 0x6d, 0xd3, 0x55, 0xd2, 0xc0, 0x69, 0x9b, 0x7e, 0x75, 0x71, 0x69, 0x92, 0x1e, 0x2c,
	0x3c, 0x40, 0x96, 0x3b, 0xf1, 0x44, 0x71, 0xc1, 0x0c, 0x2a, 0x48, 0x4a, 0x62, 0x03, 0xa7, 0x21,
	0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x98, 0x0f, 0xa1, 0x83, 0x86, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MessagingClient is the client API for Messaging service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessagingClient interface {
	UpdateMessage(ctx context.Context, in *UpdateMessageRequest, opts ...grpc.CallOption) (*Message, error)
	SubFieldMessage(ctx context.Context, in *SubFieldMessageRequest, opts ...grpc.CallOption) (*Message, error)
}

type messagingClient struct {
	cc *grpc.ClientConn
}

func NewMessagingClient(cc *grpc.ClientConn) MessagingClient {
	return &messagingClient{cc}
}

func (c *messagingClient) UpdateMessage(ctx context.Context, in *UpdateMessageRequest, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/httprule.Messaging/UpdateMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingClient) SubFieldMessage(ctx context.Context, in *SubFieldMessageRequest, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/httprule.Messaging/SubFieldMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessagingServer is the server API for Messaging service.
type MessagingServer interface {
	UpdateMessage(context.Context, *UpdateMessageRequest) (*Message, error)
	SubFieldMessage(context.Context, *SubFieldMessageRequest) (*Message, error)
}

// UnimplementedMessagingServer can be embedded to have forward compatible implementations.
type UnimplementedMessagingServer struct {
}

func (*UnimplementedMessagingServer) UpdateMessage(ctx context.Context, req *UpdateMessageRequest) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMessage not implemented")
}
func (*UnimplementedMessagingServer) SubFieldMessage(ctx context.Context, req *SubFieldMessageRequest) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubFieldMessage not implemented")
}

func RegisterMessagingServer(s *grpc.Server, srv MessagingServer) {
	s.RegisterService(&_Messaging_serviceDesc, srv)
}

func _Messaging_UpdateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServer).UpdateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/httprule.Messaging/UpdateMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServer).UpdateMessage(ctx, req.(*UpdateMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messaging_SubFieldMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubFieldMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServer).SubFieldMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/httprule.Messaging/SubFieldMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServer).SubFieldMessage(ctx, req.(*SubFieldMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Messaging_serviceDesc = grpc.ServiceDesc{
	ServiceName: "httprule.Messaging",
	HandlerType: (*MessagingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateMessage",
			Handler:    _Messaging_UpdateMessage_Handler,
		},
		{
			MethodName: "SubFieldMessage",
			Handler:    _Messaging_SubFieldMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "httprule/httprule.proto",
}
