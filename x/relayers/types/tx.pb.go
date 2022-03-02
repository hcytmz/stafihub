// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: relayers/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgCreateRelayer struct {
	Creator   string   `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Denom     string   `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
	Addresses []string `protobuf:"bytes,3,rep,name=addresses,proto3" json:"addresses,omitempty"`
}

func (m *MsgCreateRelayer) Reset()         { *m = MsgCreateRelayer{} }
func (m *MsgCreateRelayer) String() string { return proto.CompactTextString(m) }
func (*MsgCreateRelayer) ProtoMessage()    {}
func (*MsgCreateRelayer) Descriptor() ([]byte, []int) {
	return fileDescriptor_858ee6e4c5f310aa, []int{0}
}
func (m *MsgCreateRelayer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateRelayer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateRelayer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateRelayer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateRelayer.Merge(m, src)
}
func (m *MsgCreateRelayer) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateRelayer) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateRelayer.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateRelayer proto.InternalMessageInfo

func (m *MsgCreateRelayer) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateRelayer) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *MsgCreateRelayer) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

type MsgCreateRelayerResponse struct {
}

func (m *MsgCreateRelayerResponse) Reset()         { *m = MsgCreateRelayerResponse{} }
func (m *MsgCreateRelayerResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateRelayerResponse) ProtoMessage()    {}
func (*MsgCreateRelayerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_858ee6e4c5f310aa, []int{1}
}
func (m *MsgCreateRelayerResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateRelayerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateRelayerResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateRelayerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateRelayerResponse.Merge(m, src)
}
func (m *MsgCreateRelayerResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateRelayerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateRelayerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateRelayerResponse proto.InternalMessageInfo

type MsgDeleteRelayer struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Denom   string `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *MsgDeleteRelayer) Reset()         { *m = MsgDeleteRelayer{} }
func (m *MsgDeleteRelayer) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteRelayer) ProtoMessage()    {}
func (*MsgDeleteRelayer) Descriptor() ([]byte, []int) {
	return fileDescriptor_858ee6e4c5f310aa, []int{2}
}
func (m *MsgDeleteRelayer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteRelayer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteRelayer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteRelayer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteRelayer.Merge(m, src)
}
func (m *MsgDeleteRelayer) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteRelayer) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteRelayer.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteRelayer proto.InternalMessageInfo

func (m *MsgDeleteRelayer) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgDeleteRelayer) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *MsgDeleteRelayer) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type MsgDeleteRelayerResponse struct {
}

func (m *MsgDeleteRelayerResponse) Reset()         { *m = MsgDeleteRelayerResponse{} }
func (m *MsgDeleteRelayerResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteRelayerResponse) ProtoMessage()    {}
func (*MsgDeleteRelayerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_858ee6e4c5f310aa, []int{3}
}
func (m *MsgDeleteRelayerResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteRelayerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteRelayerResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteRelayerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteRelayerResponse.Merge(m, src)
}
func (m *MsgDeleteRelayerResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteRelayerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteRelayerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteRelayerResponse proto.InternalMessageInfo

type MsgUpdateThreshold struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Denom   string `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
	Value   uint32 `protobuf:"varint,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *MsgUpdateThreshold) Reset()         { *m = MsgUpdateThreshold{} }
func (m *MsgUpdateThreshold) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateThreshold) ProtoMessage()    {}
func (*MsgUpdateThreshold) Descriptor() ([]byte, []int) {
	return fileDescriptor_858ee6e4c5f310aa, []int{4}
}
func (m *MsgUpdateThreshold) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateThreshold) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateThreshold.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateThreshold) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateThreshold.Merge(m, src)
}
func (m *MsgUpdateThreshold) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateThreshold) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateThreshold.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateThreshold proto.InternalMessageInfo

func (m *MsgUpdateThreshold) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgUpdateThreshold) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *MsgUpdateThreshold) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type MsgUpdateThresholdResponse struct {
}

func (m *MsgUpdateThresholdResponse) Reset()         { *m = MsgUpdateThresholdResponse{} }
func (m *MsgUpdateThresholdResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateThresholdResponse) ProtoMessage()    {}
func (*MsgUpdateThresholdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_858ee6e4c5f310aa, []int{5}
}
func (m *MsgUpdateThresholdResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateThresholdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateThresholdResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateThresholdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateThresholdResponse.Merge(m, src)
}
func (m *MsgUpdateThresholdResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateThresholdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateThresholdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateThresholdResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateRelayer)(nil), "stafihub.stafihub.relayers.MsgCreateRelayer")
	proto.RegisterType((*MsgCreateRelayerResponse)(nil), "stafihub.stafihub.relayers.MsgCreateRelayerResponse")
	proto.RegisterType((*MsgDeleteRelayer)(nil), "stafihub.stafihub.relayers.MsgDeleteRelayer")
	proto.RegisterType((*MsgDeleteRelayerResponse)(nil), "stafihub.stafihub.relayers.MsgDeleteRelayerResponse")
	proto.RegisterType((*MsgUpdateThreshold)(nil), "stafihub.stafihub.relayers.MsgUpdateThreshold")
	proto.RegisterType((*MsgUpdateThresholdResponse)(nil), "stafihub.stafihub.relayers.MsgUpdateThresholdResponse")
}

func init() { proto.RegisterFile("relayers/tx.proto", fileDescriptor_858ee6e4c5f310aa) }

var fileDescriptor_858ee6e4c5f310aa = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x41, 0x4b, 0x02, 0x41,
	0x1c, 0xc5, 0x5d, 0x17, 0x13, 0xff, 0x20, 0xd5, 0xe0, 0x61, 0x58, 0x64, 0x91, 0x3d, 0x79, 0xa8,
	0x15, 0x2a, 0xfa, 0x00, 0x15, 0x74, 0xf2, 0xb2, 0xd4, 0x45, 0x3a, 0x34, 0xba, 0xff, 0x56, 0x61,
	0x75, 0x96, 0xf9, 0x8f, 0xa1, 0xdf, 0xa2, 0x8f, 0xd5, 0xd1, 0x63, 0xc7, 0xd0, 0x5b, 0x9f, 0x22,
	0xdc, 0x75, 0x94, 0xdd, 0x0a, 0xb1, 0xdb, 0xbc, 0xc7, 0xe3, 0xfd, 0xe0, 0x0d, 0x7f, 0x38, 0x55,
	0x18, 0x8b, 0x39, 0x2a, 0xea, 0xe8, 0x99, 0x9f, 0x28, 0xa9, 0x25, 0x73, 0x48, 0x8b, 0x97, 0xd1,
	0x70, 0xda, 0xf7, 0xb7, 0x0f, 0x13, 0xf2, 0x9e, 0xe1, 0xa4, 0x4b, 0xd1, 0xad, 0x42, 0xa1, 0x31,
	0xc8, 0x4c, 0xc6, 0xa1, 0x3a, 0x58, 0x1b, 0x52, 0x71, 0xab, 0x65, 0xb5, 0x6b, 0x81, 0x91, 0xac,
	0x01, 0x95, 0x10, 0x27, 0x72, 0xcc, 0xcb, 0xa9, 0x9f, 0x09, 0xd6, 0x84, 0x9a, 0x08, 0x43, 0x85,
	0x44, 0x48, 0xdc, 0x6e, 0xd9, 0xed, 0x5a, 0xb0, 0x33, 0x3c, 0x07, 0x78, 0x91, 0x10, 0x20, 0x25,
	0x72, 0x42, 0xe8, 0x3d, 0xa5, 0xf4, 0x3b, 0x8c, 0xf1, 0xff, 0x74, 0x0e, 0xd5, 0x0d, 0x8c, 0xdb,
	0x59, 0x7e, 0x23, 0x37, 0xe4, 0x5c, 0xfb, 0x96, 0xdc, 0x03, 0xd6, 0xa5, 0xe8, 0x31, 0x09, 0x85,
	0xc6, 0x87, 0xa1, 0x42, 0x1a, 0xca, 0x38, 0x3c, 0x98, 0xdd, 0x80, 0xca, 0xab, 0x88, 0xa7, 0x98,
	0x92, 0xeb, 0x41, 0x26, 0xbc, 0x26, 0x38, 0x3f, 0xbb, 0x0d, 0xf9, 0xe2, 0xab, 0x0c, 0x76, 0x97,
	0x22, 0x46, 0x50, 0xcf, 0xcf, 0x7e, 0xe6, 0xff, 0xfd, 0x4f, 0x7e, 0x71, 0x42, 0xe7, 0xea, 0x90,
	0xb4, 0x81, 0xaf, 0xa1, 0xf9, 0xb5, 0xf7, 0x41, 0x73, 0xe9, 0xbd, 0xd0, 0x5f, 0xb7, 0x66, 0x73,
	0x38, 0x2e, 0x0e, 0xed, 0xef, 0x29, 0x2a, 0xe4, 0x9d, 0xeb, 0xc3, 0xf2, 0x06, 0x7d, 0x73, 0xff,
	0xbe, 0x74, 0xad, 0xc5, 0xd2, 0xb5, 0x3e, 0x97, 0xae, 0xf5, 0xb6, 0x72, 0x4b, 0x8b, 0x95, 0x5b,
	0xfa, 0x58, 0xb9, 0xa5, 0xde, 0x79, 0x34, 0xd2, 0xeb, 0x96, 0x81, 0x1c, 0x77, 0x4c, 0xe5, 0xee,
	0x31, 0xeb, 0xec, 0xce, 0x68, 0x9e, 0x20, 0xf5, 0x8f, 0xd2, 0x53, 0xba, 0xfc, 0x0e, 0x00, 0x00,
	0xff, 0xff, 0xf1, 0x97, 0x55, 0x14, 0x5f, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	CreateRelayer(ctx context.Context, in *MsgCreateRelayer, opts ...grpc.CallOption) (*MsgCreateRelayerResponse, error)
	DeleteRelayer(ctx context.Context, in *MsgDeleteRelayer, opts ...grpc.CallOption) (*MsgDeleteRelayerResponse, error)
	UpdateThreshold(ctx context.Context, in *MsgUpdateThreshold, opts ...grpc.CallOption) (*MsgUpdateThresholdResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateRelayer(ctx context.Context, in *MsgCreateRelayer, opts ...grpc.CallOption) (*MsgCreateRelayerResponse, error) {
	out := new(MsgCreateRelayerResponse)
	err := c.cc.Invoke(ctx, "/stafihub.stafihub.relayers.Msg/CreateRelayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteRelayer(ctx context.Context, in *MsgDeleteRelayer, opts ...grpc.CallOption) (*MsgDeleteRelayerResponse, error) {
	out := new(MsgDeleteRelayerResponse)
	err := c.cc.Invoke(ctx, "/stafihub.stafihub.relayers.Msg/DeleteRelayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateThreshold(ctx context.Context, in *MsgUpdateThreshold, opts ...grpc.CallOption) (*MsgUpdateThresholdResponse, error) {
	out := new(MsgUpdateThresholdResponse)
	err := c.cc.Invoke(ctx, "/stafihub.stafihub.relayers.Msg/UpdateThreshold", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateRelayer(context.Context, *MsgCreateRelayer) (*MsgCreateRelayerResponse, error)
	DeleteRelayer(context.Context, *MsgDeleteRelayer) (*MsgDeleteRelayerResponse, error)
	UpdateThreshold(context.Context, *MsgUpdateThreshold) (*MsgUpdateThresholdResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateRelayer(ctx context.Context, req *MsgCreateRelayer) (*MsgCreateRelayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRelayer not implemented")
}
func (*UnimplementedMsgServer) DeleteRelayer(ctx context.Context, req *MsgDeleteRelayer) (*MsgDeleteRelayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRelayer not implemented")
}
func (*UnimplementedMsgServer) UpdateThreshold(ctx context.Context, req *MsgUpdateThreshold) (*MsgUpdateThresholdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateThreshold not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateRelayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateRelayer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateRelayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stafihub.stafihub.relayers.Msg/CreateRelayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateRelayer(ctx, req.(*MsgCreateRelayer))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteRelayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteRelayer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteRelayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stafihub.stafihub.relayers.Msg/DeleteRelayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteRelayer(ctx, req.(*MsgDeleteRelayer))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateThreshold_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateThreshold)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateThreshold(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stafihub.stafihub.relayers.Msg/UpdateThreshold",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateThreshold(ctx, req.(*MsgUpdateThreshold))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stafihub.stafihub.relayers.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRelayer",
			Handler:    _Msg_CreateRelayer_Handler,
		},
		{
			MethodName: "DeleteRelayer",
			Handler:    _Msg_DeleteRelayer_Handler,
		},
		{
			MethodName: "UpdateThreshold",
			Handler:    _Msg_UpdateThreshold_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "relayers/tx.proto",
}

func (m *MsgCreateRelayer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateRelayer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateRelayer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Addresses) > 0 {
		for iNdEx := len(m.Addresses) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Addresses[iNdEx])
			copy(dAtA[i:], m.Addresses[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.Addresses[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateRelayerResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateRelayerResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateRelayerResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgDeleteRelayer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteRelayer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteRelayer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDeleteRelayerResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteRelayerResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteRelayerResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUpdateThreshold) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateThreshold) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateThreshold) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Value))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateThresholdResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateThresholdResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateThresholdResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateRelayer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Addresses) > 0 {
		for _, s := range m.Addresses {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgCreateRelayerResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgDeleteRelayer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgDeleteRelayerResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUpdateThreshold) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Value != 0 {
		n += 1 + sovTx(uint64(m.Value))
	}
	return n
}

func (m *MsgUpdateThresholdResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateRelayer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateRelayer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateRelayer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addresses = append(m.Addresses, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgCreateRelayerResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateRelayerResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateRelayerResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgDeleteRelayer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgDeleteRelayer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteRelayer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgDeleteRelayerResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgDeleteRelayerResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteRelayerResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateThreshold) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateThreshold: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateThreshold: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateThresholdResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateThresholdResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateThresholdResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
