// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs3/app/registry/v1beta1/registry_api.proto

package registryv1beta1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	v1beta12 "github.com/refs/cs3-micro/pkg/cs3/rpc/v1beta1"
	v1beta11 "github.com/refs/cs3-micro/pkg/cs3/storage/provider/v1beta1"
	v1beta1 "github.com/refs/cs3-micro/pkg/cs3/types/v1beta1"
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

type GetAppProvidersRequest struct {
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The resource information.
	ResourceInfo         *v1beta11.ResourceInfo `protobuf:"bytes,2,opt,name=resource_info,json=resourceInfo,proto3" json:"resource_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *GetAppProvidersRequest) Reset()         { *m = GetAppProvidersRequest{} }
func (m *GetAppProvidersRequest) String() string { return proto.CompactTextString(m) }
func (*GetAppProvidersRequest) ProtoMessage()    {}
func (*GetAppProvidersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a26a91cc65771e6, []int{0}
}

func (m *GetAppProvidersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAppProvidersRequest.Unmarshal(m, b)
}
func (m *GetAppProvidersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAppProvidersRequest.Marshal(b, m, deterministic)
}
func (m *GetAppProvidersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAppProvidersRequest.Merge(m, src)
}
func (m *GetAppProvidersRequest) XXX_Size() int {
	return xxx_messageInfo_GetAppProvidersRequest.Size(m)
}
func (m *GetAppProvidersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAppProvidersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAppProvidersRequest proto.InternalMessageInfo

func (m *GetAppProvidersRequest) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *GetAppProvidersRequest) GetResourceInfo() *v1beta11.ResourceInfo {
	if m != nil {
		return m.ResourceInfo
	}
	return nil
}

type GetAppProvidersResponse struct {
	// REQUIRED.
	// The response status.
	Status *v1beta12.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The app providers available for the given resource info.
	Providers            []*ProviderInfo `protobuf:"bytes,3,rep,name=providers,proto3" json:"providers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetAppProvidersResponse) Reset()         { *m = GetAppProvidersResponse{} }
func (m *GetAppProvidersResponse) String() string { return proto.CompactTextString(m) }
func (*GetAppProvidersResponse) ProtoMessage()    {}
func (*GetAppProvidersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a26a91cc65771e6, []int{1}
}

func (m *GetAppProvidersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAppProvidersResponse.Unmarshal(m, b)
}
func (m *GetAppProvidersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAppProvidersResponse.Marshal(b, m, deterministic)
}
func (m *GetAppProvidersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAppProvidersResponse.Merge(m, src)
}
func (m *GetAppProvidersResponse) XXX_Size() int {
	return xxx_messageInfo_GetAppProvidersResponse.Size(m)
}
func (m *GetAppProvidersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAppProvidersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAppProvidersResponse proto.InternalMessageInfo

func (m *GetAppProvidersResponse) GetStatus() *v1beta12.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *GetAppProvidersResponse) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *GetAppProvidersResponse) GetProviders() []*ProviderInfo {
	if m != nil {
		return m.Providers
	}
	return nil
}

type ListAppProvidersRequest struct {
	// OPTIONAL.
	// Opaque information.
	Opaque               *v1beta1.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ListAppProvidersRequest) Reset()         { *m = ListAppProvidersRequest{} }
func (m *ListAppProvidersRequest) String() string { return proto.CompactTextString(m) }
func (*ListAppProvidersRequest) ProtoMessage()    {}
func (*ListAppProvidersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a26a91cc65771e6, []int{2}
}

func (m *ListAppProvidersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAppProvidersRequest.Unmarshal(m, b)
}
func (m *ListAppProvidersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAppProvidersRequest.Marshal(b, m, deterministic)
}
func (m *ListAppProvidersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAppProvidersRequest.Merge(m, src)
}
func (m *ListAppProvidersRequest) XXX_Size() int {
	return xxx_messageInfo_ListAppProvidersRequest.Size(m)
}
func (m *ListAppProvidersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAppProvidersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListAppProvidersRequest proto.InternalMessageInfo

func (m *ListAppProvidersRequest) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

type ListAppProvidersResponse struct {
	// REQUIRED.
	// The response status.
	Status *v1beta12.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The list of app providers this registry knows about.
	Providers            []*ProviderInfo `protobuf:"bytes,3,rep,name=providers,proto3" json:"providers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ListAppProvidersResponse) Reset()         { *m = ListAppProvidersResponse{} }
func (m *ListAppProvidersResponse) String() string { return proto.CompactTextString(m) }
func (*ListAppProvidersResponse) ProtoMessage()    {}
func (*ListAppProvidersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a26a91cc65771e6, []int{3}
}

func (m *ListAppProvidersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAppProvidersResponse.Unmarshal(m, b)
}
func (m *ListAppProvidersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAppProvidersResponse.Marshal(b, m, deterministic)
}
func (m *ListAppProvidersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAppProvidersResponse.Merge(m, src)
}
func (m *ListAppProvidersResponse) XXX_Size() int {
	return xxx_messageInfo_ListAppProvidersResponse.Size(m)
}
func (m *ListAppProvidersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAppProvidersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListAppProvidersResponse proto.InternalMessageInfo

func (m *ListAppProvidersResponse) GetStatus() *v1beta12.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ListAppProvidersResponse) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *ListAppProvidersResponse) GetProviders() []*ProviderInfo {
	if m != nil {
		return m.Providers
	}
	return nil
}

func init() {
	proto.RegisterType((*GetAppProvidersRequest)(nil), "cs3.app.registry.v1beta1.GetAppProvidersRequest")
	proto.RegisterType((*GetAppProvidersResponse)(nil), "cs3.app.registry.v1beta1.GetAppProvidersResponse")
	proto.RegisterType((*ListAppProvidersRequest)(nil), "cs3.app.registry.v1beta1.ListAppProvidersRequest")
	proto.RegisterType((*ListAppProvidersResponse)(nil), "cs3.app.registry.v1beta1.ListAppProvidersResponse")
}

func init() {
	proto.RegisterFile("cs3/app/registry/v1beta1/registry_api.proto", fileDescriptor_9a26a91cc65771e6)
}

var fileDescriptor_9a26a91cc65771e6 = []byte{
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x94, 0x4f, 0x8b, 0xda, 0x40,
	0x18, 0xc6, 0x49, 0x04, 0xa1, 0x63, 0x8b, 0x92, 0x43, 0x4d, 0x83, 0x05, 0xc9, 0xa1, 0x48, 0x5b,
	0x26, 0x4d, 0xf2, 0x09, 0xa2, 0x85, 0x22, 0x08, 0x86, 0x14, 0x7a, 0x28, 0x82, 0xc4, 0x74, 0x94,
	0x1c, 0xea, 0xbc, 0xce, 0x4c, 0x04, 0x4f, 0xed, 0x07, 0xe9, 0xa9, 0xc7, 0x7e, 0x8d, 0xee, 0x69,
	0x3f, 0xd5, 0x92, 0xc9, 0x4c, 0x56, 0x74, 0x23, 0x2b, 0xec, 0x65, 0x8f, 0x93, 0xe7, 0xf7, 0x3c,
	0xf3, 0xfe, 0x49, 0x82, 0x3e, 0x64, 0x3c, 0xf4, 0x52, 0x00, 0x8f, 0x91, 0x4d, 0xce, 0x05, 0x3b,
	0x78, 0x7b, 0x7f, 0x45, 0x44, 0xea, 0xd7, 0x0f, 0x96, 0x29, 0xe4, 0x18, 0x18, 0x15, 0xd4, 0xb2,
	0x33, 0x1e, 0xe2, 0x14, 0x00, 0x6b, 0x0d, 0x2b, 0xd8, 0x19, 0x5d, 0x88, 0xe1, 0xb4, 0x60, 0x19,
	0xe1, 0x55, 0x86, 0x33, 0x28, 0x49, 0x06, 0x59, 0x0d, 0x70, 0x91, 0x8a, 0x42, 0xab, 0x1f, 0x4b,
	0x95, 0x0b, 0xca, 0xd2, 0x0d, 0xf1, 0x80, 0xd1, 0x7d, 0xfe, 0x83, 0xb0, 0xc6, 0xac, 0xb7, 0x25,
	0x2d, 0x0e, 0x40, 0x78, 0x8d, 0xc8, 0x53, 0x25, 0xbb, 0x7f, 0x0c, 0xf4, 0xfa, 0x0b, 0x11, 0x11,
	0x40, 0xac, 0x92, 0x78, 0x42, 0x76, 0x05, 0xe1, 0xc2, 0xf2, 0x51, 0x9b, 0x42, 0xba, 0x2b, 0x88,
	0x6d, 0x0c, 0x8d, 0x51, 0x27, 0x78, 0x83, 0xcb, 0xd6, 0x2a, 0xb3, 0x8a, 0xc2, 0x73, 0x09, 0x24,
	0x0a, 0xb4, 0xe6, 0xe8, 0x95, 0xbe, 0x7f, 0x99, 0x6f, 0xd7, 0xd4, 0x36, 0xa5, 0xf3, 0xbd, 0x74,
	0xaa, 0x92, 0xb1, 0x2e, 0xb9, 0x0e, 0x49, 0x94, 0x65, 0xba, 0x5d, 0xd3, 0xe4, 0x25, 0x3b, 0x3a,
	0xb9, 0xff, 0x0d, 0xd4, 0x3f, 0x2b, 0x8f, 0x03, 0xdd, 0x72, 0x62, 0x79, 0xa8, 0x5d, 0xcd, 0x45,
	0xd5, 0xd7, 0x97, 0xb7, 0x30, 0xc8, 0xea, 0xe0, 0xaf, 0x52, 0x4e, 0x14, 0x76, 0xd4, 0x90, 0xf9,
	0xd8, 0x86, 0x3e, 0xa3, 0x17, 0xba, 0x5c, 0x6e, 0xb7, 0x86, 0xad, 0x51, 0x27, 0x78, 0x87, 0x9b,
	0x36, 0x8c, 0x75, 0x8d, 0xb2, 0x91, 0x7b, 0xa3, 0x3b, 0x43, 0xfd, 0x59, 0xce, 0x9f, 0x68, 0xc8,
	0xee, 0x8d, 0x81, 0xec, 0xf3, 0xb8, 0xe7, 0x36, 0x94, 0xe0, 0xb7, 0x89, 0x3a, 0x89, 0x82, 0xa3,
	0x78, 0x6a, 0xed, 0x51, 0xf7, 0x64, 0xd3, 0xd6, 0xa7, 0xe6, 0xd4, 0x87, 0xdf, 0x59, 0xc7, 0xbf,
	0xc2, 0xa1, 0x26, 0x76, 0x40, 0xbd, 0xd3, 0x69, 0x5a, 0x17, 0x62, 0x1a, 0x16, 0xe9, 0x04, 0xd7,
	0x58, 0xaa, 0xab, 0xc7, 0xbf, 0xd0, 0x20, 0xa3, 0x3f, 0x1b, 0x8d, 0xe3, 0x5e, 0x3d, 0x1f, 0xc8,
	0xe3, 0xf2, 0x73, 0x8d, 0x8d, 0xef, 0x5d, 0x4d, 0x29, 0xe8, 0xaf, 0xd9, 0x9a, 0x44, 0xc9, 0x3f,
	0xd3, 0x9e, 0xf0, 0x10, 0x47, 0x00, 0x58, 0x7b, 0xf0, 0x37, 0x7f, 0x5c, 0x02, 0xb7, 0x52, 0x5a,
	0x44, 0x00, 0x0b, 0x2d, 0x2d, 0x94, 0xb4, 0x6a, 0xcb, 0x9f, 0x40, 0x78, 0x17, 0x00, 0x00, 0xff,
	0xff, 0x7f, 0x47, 0x97, 0x47, 0xe2, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RegistryAPIClient is the client API for RegistryAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RegistryAPIClient interface {
	// Returns the app providers that are capable of handling this resource info.
	// MUST return CODE_NOT_FOUND if no providers are available.
	GetAppProviders(ctx context.Context, in *GetAppProvidersRequest, opts ...grpc.CallOption) (*GetAppProvidersResponse, error)
	// Returns a list of the available app providers known by this registry.
	ListAppProviders(ctx context.Context, in *ListAppProvidersRequest, opts ...grpc.CallOption) (*ListAppProvidersResponse, error)
}

type registryAPIClient struct {
	cc *grpc.ClientConn
}

func NewRegistryAPIClient(cc *grpc.ClientConn) RegistryAPIClient {
	return &registryAPIClient{cc}
}

func (c *registryAPIClient) GetAppProviders(ctx context.Context, in *GetAppProvidersRequest, opts ...grpc.CallOption) (*GetAppProvidersResponse, error) {
	out := new(GetAppProvidersResponse)
	err := c.cc.Invoke(ctx, "/cs3.app.registry.v1beta1.RegistryAPI/GetAppProviders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryAPIClient) ListAppProviders(ctx context.Context, in *ListAppProvidersRequest, opts ...grpc.CallOption) (*ListAppProvidersResponse, error) {
	out := new(ListAppProvidersResponse)
	err := c.cc.Invoke(ctx, "/cs3.app.registry.v1beta1.RegistryAPI/ListAppProviders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistryAPIServer is the server API for RegistryAPI service.
type RegistryAPIServer interface {
	// Returns the app providers that are capable of handling this resource info.
	// MUST return CODE_NOT_FOUND if no providers are available.
	GetAppProviders(context.Context, *GetAppProvidersRequest) (*GetAppProvidersResponse, error)
	// Returns a list of the available app providers known by this registry.
	ListAppProviders(context.Context, *ListAppProvidersRequest) (*ListAppProvidersResponse, error)
}

// UnimplementedRegistryAPIServer can be embedded to have forward compatible implementations.
type UnimplementedRegistryAPIServer struct {
}

func (*UnimplementedRegistryAPIServer) GetAppProviders(ctx context.Context, req *GetAppProvidersRequest) (*GetAppProvidersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppProviders not implemented")
}
func (*UnimplementedRegistryAPIServer) ListAppProviders(ctx context.Context, req *ListAppProvidersRequest) (*ListAppProvidersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAppProviders not implemented")
}

func RegisterRegistryAPIServer(s *grpc.Server, srv RegistryAPIServer) {
	s.RegisterService(&_RegistryAPI_serviceDesc, srv)
}

func _RegistryAPI_GetAppProviders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppProvidersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryAPIServer).GetAppProviders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cs3.app.registry.v1beta1.RegistryAPI/GetAppProviders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryAPIServer).GetAppProviders(ctx, req.(*GetAppProvidersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistryAPI_ListAppProviders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAppProvidersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryAPIServer).ListAppProviders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cs3.app.registry.v1beta1.RegistryAPI/ListAppProviders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryAPIServer).ListAppProviders(ctx, req.(*ListAppProvidersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RegistryAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cs3.app.registry.v1beta1.RegistryAPI",
	HandlerType: (*RegistryAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAppProviders",
			Handler:    _RegistryAPI_GetAppProviders_Handler,
		},
		{
			MethodName: "ListAppProviders",
			Handler:    _RegistryAPI_ListAppProviders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cs3/app/registry/v1beta1/registry_api.proto",
}