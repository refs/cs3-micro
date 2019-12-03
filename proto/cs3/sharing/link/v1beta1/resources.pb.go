// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs3/sharing/link/v1beta1/resources.proto

package linkv1beta1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	v1beta11 "github.com/refs/cs3-micro/proto/cs3/identity/user/v1beta1"
	v1beta1 "github.com/refs/cs3-micro/proto/cs3/storage/provider/v1beta1"
	v1beta12 "github.com/refs/cs3-micro/proto/cs3/types/v1beta1"
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

// Public share are relationships between a resource owner
// (usually the authenticated user) who grants permissions to a recipient (grantee)
// on a specified resource (resource_id). UserShares represents both user and groups.
// TODO(labkode): do we need to have  resource_type stored on the share?
// This is not needed if when getting the shares a stat operation is launched against the
// the storage provider.
type PublicShare struct {
	// REQUIRED.
	// Opaque unique identifier of the share.
	Id *PublicShareId `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// REQUIRED.
	// The unlisted token to give public access
	// to the public share.
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	// REQUIRED.
	// Unique identifier of the shared resource.
	ResourceId *v1beta1.ResourceId `protobuf:"bytes,3,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// REQUIRED.
	// Permissions for the grantee to use
	// the resource.
	Permissions *PublicSharePermissions `protobuf:"bytes,4,opt,name=permissions,proto3" json:"permissions,omitempty"`
	// REQUIRED.
	// Uniquely identifies the owner of the share
	// (the resource owner at the time of creating the share).
	// In case the ownership of the underlying resource changes
	// the share owner field MAY change to reflect the change of ownsership.
	Owner *v1beta11.UserId `protobuf:"bytes,5,opt,name=owner,proto3" json:"owner,omitempty"`
	// REQUIRED.
	// Uniquely identifies a principal who initiates the share creation.
	// A creator can create shares on behalf of the owner (because of re-sharing,
	// because belonging to special groups, ...).
	// Creator and owner often result in being the same principal.
	Creator *v1beta11.UserId `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
	// REQUIRED.
	// Creation time of the share.
	Ctime *v1beta12.Timestamp `protobuf:"bytes,7,opt,name=ctime,proto3" json:"ctime,omitempty"`
	// REQUIRED.
	// Last modification time of the share.
	Mtime *v1beta12.Timestamp `protobuf:"bytes,8,opt,name=mtime,proto3" json:"mtime,omitempty"`
	// REQUIRED.
	// Determines if the public share is password protected or not.
	PasswordProtected bool `protobuf:"varint,9,opt,name=password_protected,json=passwordProtected,proto3" json:"password_protected,omitempty"`
	// OPTIONAL.
	// The expiration time for the public share.
	Expiration *v1beta12.Timestamp `protobuf:"bytes,10,opt,name=expiration,proto3" json:"expiration,omitempty"`
	// OPTIONAL.
	// Display name for the shared resource (such as file, directory basename or any
	// user defined name).
	// The display name MAY be different than the actual resource basename.
	// This field is only useful for informational purposes, like for example,
	// setting the window title in a public share HTML page.
	DisplayName          string   `protobuf:"bytes,11,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublicShare) Reset()         { *m = PublicShare{} }
func (m *PublicShare) String() string { return proto.CompactTextString(m) }
func (*PublicShare) ProtoMessage()    {}
func (*PublicShare) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f13c61869aaebc4, []int{0}
}

func (m *PublicShare) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublicShare.Unmarshal(m, b)
}
func (m *PublicShare) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublicShare.Marshal(b, m, deterministic)
}
func (m *PublicShare) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublicShare.Merge(m, src)
}
func (m *PublicShare) XXX_Size() int {
	return xxx_messageInfo_PublicShare.Size(m)
}
func (m *PublicShare) XXX_DiscardUnknown() {
	xxx_messageInfo_PublicShare.DiscardUnknown(m)
}

var xxx_messageInfo_PublicShare proto.InternalMessageInfo

func (m *PublicShare) GetId() *PublicShareId {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *PublicShare) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *PublicShare) GetResourceId() *v1beta1.ResourceId {
	if m != nil {
		return m.ResourceId
	}
	return nil
}

func (m *PublicShare) GetPermissions() *PublicSharePermissions {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *PublicShare) GetOwner() *v1beta11.UserId {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *PublicShare) GetCreator() *v1beta11.UserId {
	if m != nil {
		return m.Creator
	}
	return nil
}

func (m *PublicShare) GetCtime() *v1beta12.Timestamp {
	if m != nil {
		return m.Ctime
	}
	return nil
}

func (m *PublicShare) GetMtime() *v1beta12.Timestamp {
	if m != nil {
		return m.Mtime
	}
	return nil
}

func (m *PublicShare) GetPasswordProtected() bool {
	if m != nil {
		return m.PasswordProtected
	}
	return false
}

func (m *PublicShare) GetExpiration() *v1beta12.Timestamp {
	if m != nil {
		return m.Expiration
	}
	return nil
}

func (m *PublicShare) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

// The permissions for a share.
type PublicSharePermissions struct {
	Permissions          *v1beta1.ResourcePermissions `protobuf:"bytes,1,opt,name=permissions,proto3" json:"permissions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *PublicSharePermissions) Reset()         { *m = PublicSharePermissions{} }
func (m *PublicSharePermissions) String() string { return proto.CompactTextString(m) }
func (*PublicSharePermissions) ProtoMessage()    {}
func (*PublicSharePermissions) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f13c61869aaebc4, []int{1}
}

func (m *PublicSharePermissions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublicSharePermissions.Unmarshal(m, b)
}
func (m *PublicSharePermissions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublicSharePermissions.Marshal(b, m, deterministic)
}
func (m *PublicSharePermissions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublicSharePermissions.Merge(m, src)
}
func (m *PublicSharePermissions) XXX_Size() int {
	return xxx_messageInfo_PublicSharePermissions.Size(m)
}
func (m *PublicSharePermissions) XXX_DiscardUnknown() {
	xxx_messageInfo_PublicSharePermissions.DiscardUnknown(m)
}

var xxx_messageInfo_PublicSharePermissions proto.InternalMessageInfo

func (m *PublicSharePermissions) GetPermissions() *v1beta1.ResourcePermissions {
	if m != nil {
		return m.Permissions
	}
	return nil
}

// A share id identifies uniquely a // share in the share provider namespace.
// A ShareId MUST be unique inside the share provider.
type PublicShareId struct {
	// REQUIRED.
	// The internal id used by service implementor to
	// uniquely identity the share in the internal
	// implementation of the service.
	OpaqueId             string   `protobuf:"bytes,2,opt,name=opaque_id,json=opaqueId,proto3" json:"opaque_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublicShareId) Reset()         { *m = PublicShareId{} }
func (m *PublicShareId) String() string { return proto.CompactTextString(m) }
func (*PublicShareId) ProtoMessage()    {}
func (*PublicShareId) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f13c61869aaebc4, []int{2}
}

func (m *PublicShareId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublicShareId.Unmarshal(m, b)
}
func (m *PublicShareId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublicShareId.Marshal(b, m, deterministic)
}
func (m *PublicShareId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublicShareId.Merge(m, src)
}
func (m *PublicShareId) XXX_Size() int {
	return xxx_messageInfo_PublicShareId.Size(m)
}
func (m *PublicShareId) XXX_DiscardUnknown() {
	xxx_messageInfo_PublicShareId.DiscardUnknown(m)
}

var xxx_messageInfo_PublicShareId proto.InternalMessageInfo

func (m *PublicShareId) GetOpaqueId() string {
	if m != nil {
		return m.OpaqueId
	}
	return ""
}

// The mechanism to identify a share
// in the share provider namespace.
type PublicShareReference struct {
	// REQUIRED.
	// One of the specifications MUST be specified.
	//
	// Types that are valid to be assigned to Spec:
	//	*PublicShareReference_Id
	//	*PublicShareReference_Token
	Spec                 isPublicShareReference_Spec `protobuf_oneof:"spec"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *PublicShareReference) Reset()         { *m = PublicShareReference{} }
func (m *PublicShareReference) String() string { return proto.CompactTextString(m) }
func (*PublicShareReference) ProtoMessage()    {}
func (*PublicShareReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f13c61869aaebc4, []int{3}
}

func (m *PublicShareReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublicShareReference.Unmarshal(m, b)
}
func (m *PublicShareReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublicShareReference.Marshal(b, m, deterministic)
}
func (m *PublicShareReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublicShareReference.Merge(m, src)
}
func (m *PublicShareReference) XXX_Size() int {
	return xxx_messageInfo_PublicShareReference.Size(m)
}
func (m *PublicShareReference) XXX_DiscardUnknown() {
	xxx_messageInfo_PublicShareReference.DiscardUnknown(m)
}

var xxx_messageInfo_PublicShareReference proto.InternalMessageInfo

type isPublicShareReference_Spec interface {
	isPublicShareReference_Spec()
}

type PublicShareReference_Id struct {
	Id *PublicShareId `protobuf:"bytes,1,opt,name=id,proto3,oneof"`
}

type PublicShareReference_Token struct {
	Token string `protobuf:"bytes,2,opt,name=token,proto3,oneof"`
}

func (*PublicShareReference_Id) isPublicShareReference_Spec() {}

func (*PublicShareReference_Token) isPublicShareReference_Spec() {}

func (m *PublicShareReference) GetSpec() isPublicShareReference_Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *PublicShareReference) GetId() *PublicShareId {
	if x, ok := m.GetSpec().(*PublicShareReference_Id); ok {
		return x.Id
	}
	return nil
}

func (m *PublicShareReference) GetToken() string {
	if x, ok := m.GetSpec().(*PublicShareReference_Token); ok {
		return x.Token
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PublicShareReference) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PublicShareReference_Id)(nil),
		(*PublicShareReference_Token)(nil),
	}
}

// Defines the restrictions for the public share.
type Grant struct {
	// REQUIRED.
	// The permissions for the share.
	Permissions *PublicSharePermissions `protobuf:"bytes,1,opt,name=permissions,proto3" json:"permissions,omitempty"`
	// OPTIONAL.
	// A password to protect the access to the public share.
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// OPTIONAL.
	// An expiration date to protect the access to the public share.
	Expiration           *v1beta12.Timestamp `protobuf:"bytes,3,opt,name=expiration,proto3" json:"expiration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Grant) Reset()         { *m = Grant{} }
func (m *Grant) String() string { return proto.CompactTextString(m) }
func (*Grant) ProtoMessage()    {}
func (*Grant) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f13c61869aaebc4, []int{4}
}

func (m *Grant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Grant.Unmarshal(m, b)
}
func (m *Grant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Grant.Marshal(b, m, deterministic)
}
func (m *Grant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Grant.Merge(m, src)
}
func (m *Grant) XXX_Size() int {
	return xxx_messageInfo_Grant.Size(m)
}
func (m *Grant) XXX_DiscardUnknown() {
	xxx_messageInfo_Grant.DiscardUnknown(m)
}

var xxx_messageInfo_Grant proto.InternalMessageInfo

func (m *Grant) GetPermissions() *PublicSharePermissions {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *Grant) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Grant) GetExpiration() *v1beta12.Timestamp {
	if m != nil {
		return m.Expiration
	}
	return nil
}

func init() {
	proto.RegisterType((*PublicShare)(nil), "cs3.sharing.link.v1beta1.PublicShare")
	proto.RegisterType((*PublicSharePermissions)(nil), "cs3.sharing.link.v1beta1.PublicSharePermissions")
	proto.RegisterType((*PublicShareId)(nil), "cs3.sharing.link.v1beta1.PublicShareId")
	proto.RegisterType((*PublicShareReference)(nil), "cs3.sharing.link.v1beta1.PublicShareReference")
	proto.RegisterType((*Grant)(nil), "cs3.sharing.link.v1beta1.Grant")
}

func init() {
	proto.RegisterFile("cs3/sharing/link/v1beta1/resources.proto", fileDescriptor_2f13c61869aaebc4)
}

var fileDescriptor_2f13c61869aaebc4 = []byte{
	// 566 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4d, 0x6b, 0x13, 0x41,
	0x18, 0xee, 0xe6, 0xab, 0xc9, 0xbb, 0x2a, 0x38, 0x94, 0x32, 0xc4, 0x0a, 0x69, 0x2e, 0x46, 0xa8,
	0xbb, 0xa6, 0x39, 0x14, 0xd1, 0x53, 0x7a, 0xb0, 0x81, 0x22, 0x61, 0xa3, 0x1e, 0x24, 0x10, 0x26,
	0x3b, 0xaf, 0xed, 0xd0, 0xec, 0xce, 0x76, 0x66, 0xd2, 0x9a, 0xbf, 0xe3, 0xd1, 0xbb, 0x7f, 0xc2,
	0x7f, 0xe2, 0xbf, 0x90, 0x9d, 0xdd, 0x8d, 0x49, 0x31, 0x12, 0xa5, 0xb7, 0xcc, 0x3e, 0x1f, 0xf3,
	0x0e, 0xcf, 0xf3, 0x06, 0x3a, 0xa1, 0xee, 0xf9, 0xfa, 0x92, 0x29, 0x11, 0x5f, 0xf8, 0x33, 0x11,
	0x5f, 0xf9, 0x37, 0xdd, 0x29, 0x1a, 0xd6, 0xf5, 0x15, 0x6a, 0x39, 0x57, 0x21, 0x6a, 0x2f, 0x51,
	0xd2, 0x48, 0x42, 0x43, 0xdd, 0xf3, 0x72, 0xa6, 0x97, 0x32, 0xbd, 0x9c, 0xd9, 0x7c, 0x9e, 0x7a,
	0x08, 0x8e, 0xb1, 0x11, 0x66, 0xe1, 0xcf, 0x35, 0xaa, 0x4d, 0x26, 0xcd, 0x23, 0x7b, 0x9d, 0x91,
	0x8a, 0x5d, 0xa0, 0x9f, 0x28, 0x79, 0x23, 0xf8, 0x5f, 0xd8, 0x4f, 0x53, 0xb6, 0x59, 0x24, 0xa8,
	0x97, 0x14, 0x7b, 0xca, 0xe0, 0xf6, 0xcf, 0x0a, 0xb8, 0xc3, 0xf9, 0x74, 0x26, 0xc2, 0xd1, 0x25,
	0x53, 0x48, 0x4e, 0xa0, 0x24, 0x38, 0x75, 0x5a, 0x4e, 0xc7, 0x3d, 0x7e, 0xe6, 0x6d, 0x1a, 0xd7,
	0x5b, 0x91, 0x0c, 0x78, 0x50, 0x12, 0x9c, 0xec, 0x41, 0xd5, 0xc8, 0x2b, 0x8c, 0x69, 0xa9, 0xe5,
	0x74, 0x1a, 0x41, 0x76, 0x20, 0x03, 0x70, 0x8b, 0x81, 0x26, 0x82, 0xd3, 0xb2, 0xf5, 0xed, 0x64,
	0xbe, 0xd9, 0x0b, 0xbc, 0xe2, 0x05, 0x4b, 0xef, 0x20, 0x17, 0x0c, 0x78, 0x00, 0x6a, 0xf9, 0x9b,
	0x04, 0xe0, 0x26, 0xa8, 0x22, 0xa1, 0xb5, 0x90, 0xb1, 0xa6, 0x15, 0x6b, 0xf5, 0x72, 0xab, 0x11,
	0x87, 0xbf, 0x75, 0xc1, 0xaa, 0x09, 0x39, 0x81, 0xaa, 0xbc, 0x8d, 0x51, 0xd1, 0xaa, 0x75, 0x3b,
	0xb4, 0x6e, 0x45, 0x0a, 0x5e, 0x9a, 0xc2, 0xd2, 0xee, 0x83, 0x46, 0x35, 0xe0, 0x41, 0xc6, 0x27,
	0xaf, 0x61, 0x37, 0x54, 0xc8, 0x8c, 0x54, 0xb4, 0xb6, 0xad, 0xb4, 0x50, 0x90, 0x63, 0xa8, 0x86,
	0x46, 0x44, 0x48, 0x77, 0xad, 0xf4, 0xc0, 0x4a, 0xb3, 0x50, 0x0a, 0xc9, 0x7b, 0x11, 0xa1, 0x36,
	0x2c, 0x4a, 0x82, 0x8c, 0x9a, 0x6a, 0x22, 0xab, 0xa9, 0x6f, 0xa3, 0xb1, 0x54, 0xf2, 0x02, 0x48,
	0xc2, 0xb4, 0xbe, 0x95, 0x8a, 0x4f, 0xd2, 0xb4, 0x31, 0x34, 0xc8, 0x69, 0xa3, 0xe5, 0x74, 0xea,
	0xc1, 0xe3, 0x02, 0x19, 0x16, 0x00, 0x79, 0x03, 0x80, 0x5f, 0x12, 0xa1, 0x98, 0x11, 0x32, 0xa6,
	0xb0, 0xc5, 0x3d, 0x2b, 0x7c, 0x72, 0x08, 0x0f, 0xb8, 0xd0, 0xc9, 0x8c, 0x2d, 0x26, 0x31, 0x8b,
	0x90, 0xba, 0xb6, 0x06, 0x6e, 0xfe, 0xed, 0x1d, 0x8b, 0xb0, 0x1d, 0xc1, 0xfe, 0x9f, 0x43, 0x21,
	0xa3, 0xf5, 0x6c, 0xb3, 0xfa, 0x75, 0xb7, 0xab, 0xc9, 0xa6, 0x70, 0xdb, 0x47, 0xf0, 0x70, 0xad,
	0xa6, 0xe4, 0x09, 0x34, 0x64, 0xc2, 0xae, 0xe7, 0xb6, 0x8a, 0x59, 0x4d, 0xeb, 0xd9, 0x87, 0x01,
	0x6f, 0x5f, 0xc3, 0xde, 0x0a, 0x3b, 0xc0, 0xcf, 0xa8, 0x30, 0x0e, 0x91, 0xbc, 0xfa, 0x8f, 0x85,
	0x38, 0xdb, 0xb1, 0x2b, 0xb1, 0xbf, 0xb6, 0x12, 0x67, 0x3b, 0xf9, 0x52, 0xf4, 0x6b, 0x50, 0xd1,
	0x09, 0x86, 0xed, 0xef, 0x0e, 0x54, 0xdf, 0x2a, 0x16, 0x9b, 0xbb, 0xdd, 0x76, 0xee, 0xa3, 0xdb,
	0x4d, 0xa8, 0x17, 0x19, 0x17, 0x8f, 0x2d, 0xce, 0x77, 0xa2, 0x2e, 0xff, 0x5b, 0xd4, 0xfd, 0x05,
	0x1c, 0x84, 0x32, 0xda, 0x38, 0x5d, 0xff, 0x51, 0x11, 0x8d, 0x4e, 0xcb, 0x25, 0x87, 0xce, 0x27,
	0x37, 0xc5, 0x73, 0xf8, 0x6b, 0xa9, 0x7c, 0x3a, 0x3a, 0xff, 0x56, 0xa2, 0xa7, 0xba, 0xe7, 0x8d,
	0x72, 0xfd, 0x79, 0xaa, 0xff, 0xd8, 0xed, 0xa7, 0x84, 0x1f, 0x16, 0x1a, 0xe7, 0xd0, 0x38, 0x85,
	0xc6, 0x39, 0x34, 0xad, 0xd9, 0x7f, 0xad, 0xde, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe1, 0x04,
	0x8d, 0x6c, 0x73, 0x05, 0x00, 0x00,
}
