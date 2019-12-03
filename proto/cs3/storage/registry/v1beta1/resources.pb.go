// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs3/storage/registry/v1beta1/resources.proto

package registryv1beta1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	v1beta1 "github.com/refs/cs3-micro/proto/cs3/types/v1beta1"
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

// The information for the storage provider.
type ProviderInfo struct {
	// OPTIONAL.
	// Opaque information (containing storage-specific information).
	// For example, additional metadata attached to the resource.
	Opaque *v1beta1.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The storage provider id that will become part of the
	// resource id.
	// For example, if the provider_id is "home", resources obtained
	// from this storage provider will have a resource id like "home:1234".
	ProviderId string `protobuf:"bytes,2,opt,name=provider_id,json=providerId,proto3" json:"provider_id,omitempty"`
	// REQUIRED.
	// The public path prefix this storage provider handles.
	// In Unix literature, the mount path.
	// For example, if the provider_path is "/home", resources obtained
	// from this storage provier will have a resource path lik "/home/docs".
	ProviderPath string `protobuf:"bytes,3,opt,name=provider_path,json=providerPath,proto3" json:"provider_path,omitempty"`
	// REQUIRED.
	// The address where the storage provider can be reached.
	// For example, tcp://localhost:1099.
	Address string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	// OPTIONAL.
	// Information to describe the functionalities
	// offered by the storage provider. Meant to be read
	// by humans.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// REQUIRED.
	// List of available methods.
	Features             *ProviderInfo_Features `protobuf:"bytes,6,opt,name=features,proto3" json:"features,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ProviderInfo) Reset()         { *m = ProviderInfo{} }
func (m *ProviderInfo) String() string { return proto.CompactTextString(m) }
func (*ProviderInfo) ProtoMessage()    {}
func (*ProviderInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_50a67fe50917e260, []int{0}
}

func (m *ProviderInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProviderInfo.Unmarshal(m, b)
}
func (m *ProviderInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProviderInfo.Marshal(b, m, deterministic)
}
func (m *ProviderInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProviderInfo.Merge(m, src)
}
func (m *ProviderInfo) XXX_Size() int {
	return xxx_messageInfo_ProviderInfo.Size(m)
}
func (m *ProviderInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ProviderInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ProviderInfo proto.InternalMessageInfo

func (m *ProviderInfo) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *ProviderInfo) GetProviderId() string {
	if m != nil {
		return m.ProviderId
	}
	return ""
}

func (m *ProviderInfo) GetProviderPath() string {
	if m != nil {
		return m.ProviderPath
	}
	return ""
}

func (m *ProviderInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ProviderInfo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ProviderInfo) GetFeatures() *ProviderInfo_Features {
	if m != nil {
		return m.Features
	}
	return nil
}

// REQUIRED.
// Represents the list of features available
// on this storage provider. If the feature is not supported,
// the related service methods MUST return CODE_UNIMPLEMENTED.
type ProviderInfo_Features struct {
	Recycle              bool     `protobuf:"varint,1,opt,name=recycle,proto3" json:"recycle,omitempty"`
	FileVersions         bool     `protobuf:"varint,2,opt,name=file_versions,json=fileVersions,proto3" json:"file_versions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProviderInfo_Features) Reset()         { *m = ProviderInfo_Features{} }
func (m *ProviderInfo_Features) String() string { return proto.CompactTextString(m) }
func (*ProviderInfo_Features) ProtoMessage()    {}
func (*ProviderInfo_Features) Descriptor() ([]byte, []int) {
	return fileDescriptor_50a67fe50917e260, []int{0, 0}
}

func (m *ProviderInfo_Features) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProviderInfo_Features.Unmarshal(m, b)
}
func (m *ProviderInfo_Features) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProviderInfo_Features.Marshal(b, m, deterministic)
}
func (m *ProviderInfo_Features) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProviderInfo_Features.Merge(m, src)
}
func (m *ProviderInfo_Features) XXX_Size() int {
	return xxx_messageInfo_ProviderInfo_Features.Size(m)
}
func (m *ProviderInfo_Features) XXX_DiscardUnknown() {
	xxx_messageInfo_ProviderInfo_Features.DiscardUnknown(m)
}

var xxx_messageInfo_ProviderInfo_Features proto.InternalMessageInfo

func (m *ProviderInfo_Features) GetRecycle() bool {
	if m != nil {
		return m.Recycle
	}
	return false
}

func (m *ProviderInfo_Features) GetFileVersions() bool {
	if m != nil {
		return m.FileVersions
	}
	return false
}

func init() {
	proto.RegisterType((*ProviderInfo)(nil), "cs3.storage.registry.v1beta1.ProviderInfo")
	proto.RegisterType((*ProviderInfo_Features)(nil), "cs3.storage.registry.v1beta1.ProviderInfo.Features")
}

func init() {
	proto.RegisterFile("cs3/storage/registry/v1beta1/resources.proto", fileDescriptor_50a67fe50917e260)
}

var fileDescriptor_50a67fe50917e260 = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x4f, 0x6a, 0xc2, 0x40,
	0x14, 0xc6, 0x49, 0x6c, 0x6d, 0x3a, 0xda, 0x16, 0xb2, 0x4a, 0xc5, 0xd2, 0xd0, 0x6e, 0x5c, 0x94,
	0x09, 0x69, 0x6e, 0xa0, 0x50, 0x70, 0x65, 0x18, 0xc1, 0x45, 0x11, 0x24, 0x26, 0x4f, 0x1d, 0xb0,
	0x4e, 0x3a, 0x6f, 0x14, 0x3c, 0x42, 0xaf, 0xd1, 0x65, 0x8f, 0xd2, 0x4b, 0xf4, 0x2a, 0x65, 0xc6,
	0x19, 0x71, 0xe5, 0xf2, 0x7d, 0xdf, 0xef, 0x7d, 0xbc, 0x3f, 0xe4, 0xa5, 0xc4, 0x2c, 0x41, 0x25,
	0x64, 0xb1, 0x84, 0x44, 0xc2, 0x92, 0xa3, 0x92, 0xfb, 0x64, 0x97, 0xce, 0x41, 0x15, 0x69, 0x22,
	0x01, 0xc5, 0x56, 0x96, 0x80, 0xb4, 0x96, 0x42, 0x89, 0xb0, 0x5b, 0x62, 0x46, 0x2d, 0x4d, 0x1d,
	0x4d, 0x2d, 0xdd, 0x79, 0xd0, 0x59, 0x6a, 0x5f, 0x03, 0x1e, 0x03, 0x4c, 0x75, 0x68, 0x7e, 0xfa,
	0xf3, 0x49, 0x3b, 0x97, 0x62, 0xc7, 0x2b, 0x90, 0xc3, 0xcd, 0x42, 0x84, 0x29, 0x69, 0x8a, 0xba,
	0xf8, 0xdc, 0x42, 0xe4, 0xc5, 0x5e, 0xaf, 0xf5, 0x7a, 0x4f, 0x75, 0xfc, 0xa1, 0xc5, 0x06, 0xd0,
	0x91, 0x01, 0x98, 0x05, 0xc3, 0x47, 0xd2, 0xaa, 0x6d, 0xc4, 0x8c, 0x57, 0x91, 0x1f, 0x7b, 0xbd,
	0x6b, 0x46, 0x9c, 0x34, 0xac, 0xc2, 0x67, 0x72, 0x73, 0x04, 0xea, 0x42, 0xad, 0xa2, 0x86, 0x41,
	0xda, 0x4e, 0xcc, 0x0b, 0xb5, 0x0a, 0x23, 0x72, 0x55, 0x54, 0x95, 0x04, 0xc4, 0xe8, 0xc2, 0xd8,
	0xae, 0x0c, 0x63, 0xd2, 0xaa, 0x00, 0x4b, 0xc9, 0x6b, 0xc5, 0xc5, 0x26, 0xba, 0x34, 0xee, 0xa9,
	0x14, 0x8e, 0x48, 0xb0, 0x80, 0x42, 0x6d, 0x25, 0x60, 0xd4, 0x34, 0x63, 0x67, 0xf4, 0xdc, 0x55,
	0xe8, 0xe9, 0xca, 0xf4, 0xcd, 0xb6, 0xb2, 0x63, 0x48, 0x67, 0x48, 0x02, 0xa7, 0xea, 0xc1, 0x24,
	0x94, 0xfb, 0x72, 0x7d, 0x38, 0x49, 0xc0, 0x5c, 0xa9, 0xf7, 0x5a, 0xf0, 0x35, 0xcc, 0x76, 0x20,
	0x91, 0x8b, 0x0d, 0x9a, 0xd5, 0x03, 0xd6, 0xd6, 0xe2, 0xc4, 0x6a, 0xfd, 0x2f, 0x8f, 0xc4, 0xa5,
	0xf8, 0x38, 0x3b, 0x4f, 0xff, 0x96, 0xb9, 0xa7, 0xe6, 0xfa, 0x2d, 0xb9, 0xf7, 0x7e, 0xe7, 0x18,
	0x8b, 0x7c, 0xfb, 0x8d, 0xc1, 0x98, 0xfd, 0xf8, 0xdd, 0x01, 0x66, 0x74, 0x6c, 0x73, 0x98, 0xcb,
	0x99, 0xa4, 0x7d, 0x0d, 0xfd, 0x1a, 0x7b, 0x6a, 0xed, 0xa9, 0xb3, 0xa7, 0xd6, 0x9e, 0x37, 0xcd,
	0xd3, 0xb3, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6e, 0x80, 0x52, 0x5a, 0x61, 0x02, 0x00, 0x00,
}
