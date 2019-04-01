// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/talent/v4beta1/tenant.proto

package talent // import "google.golang.org/genproto/googleapis/cloud/talent/v4beta1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum that represents how user data owned by the tenant is used.
type Tenant_DataUsageType int32

const (
	// Default value.
	Tenant_DATA_USAGE_TYPE_UNSPECIFIED Tenant_DataUsageType = 0
	// Data owned by this tenant is used to improve search/recommendation
	// quality across tenants.
	Tenant_AGGREGATED Tenant_DataUsageType = 1
	// Data owned by this tenant is used to improve search/recommendation
	// quality for this tenant only.
	Tenant_ISOLATED Tenant_DataUsageType = 2
)

var Tenant_DataUsageType_name = map[int32]string{
	0: "DATA_USAGE_TYPE_UNSPECIFIED",
	1: "AGGREGATED",
	2: "ISOLATED",
}
var Tenant_DataUsageType_value = map[string]int32{
	"DATA_USAGE_TYPE_UNSPECIFIED": 0,
	"AGGREGATED":                  1,
	"ISOLATED":                    2,
}

func (x Tenant_DataUsageType) String() string {
	return proto.EnumName(Tenant_DataUsageType_name, int32(x))
}
func (Tenant_DataUsageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_tenant_a58ed01536a3d383, []int{0, 0}
}

// A Tenant resource represents a tenant in the service. A tenant is a group or
// entity that shares common access with specific privileges for resources like
// profiles. Customer may create multiple tenants to provide data isolation for
// different groups.
type Tenant struct {
	// Required during tenant update.
	//
	// The resource name for a tenant. This is generated by the service when a
	// tenant is created.
	//
	// The format is "projects/{project_id}/tenants/{tenant_id}", for example,
	// "projects/api-test-project/tenants/foo".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required.
	//
	// Client side tenant identifier, used to uniquely identify the tenant.
	//
	// The maximum number of allowed characters is 255.
	ExternalId string `protobuf:"bytes,2,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	// Optional.
	//
	// Indicates whether data owned by this tenant may be used to provide product
	// improvements across other tenants.
	//
	// Defaults behavior is
	// [DataUsageType.ISOLATED][google.cloud.talent.v4beta1.Tenant.DataUsageType.ISOLATED]
	// if it's unset.
	UsageType            Tenant_DataUsageType `protobuf:"varint,3,opt,name=usage_type,json=usageType,proto3,enum=google.cloud.talent.v4beta1.Tenant_DataUsageType" json:"usage_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Tenant) Reset()         { *m = Tenant{} }
func (m *Tenant) String() string { return proto.CompactTextString(m) }
func (*Tenant) ProtoMessage()    {}
func (*Tenant) Descriptor() ([]byte, []int) {
	return fileDescriptor_tenant_a58ed01536a3d383, []int{0}
}
func (m *Tenant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tenant.Unmarshal(m, b)
}
func (m *Tenant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tenant.Marshal(b, m, deterministic)
}
func (dst *Tenant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tenant.Merge(dst, src)
}
func (m *Tenant) XXX_Size() int {
	return xxx_messageInfo_Tenant.Size(m)
}
func (m *Tenant) XXX_DiscardUnknown() {
	xxx_messageInfo_Tenant.DiscardUnknown(m)
}

var xxx_messageInfo_Tenant proto.InternalMessageInfo

func (m *Tenant) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Tenant) GetExternalId() string {
	if m != nil {
		return m.ExternalId
	}
	return ""
}

func (m *Tenant) GetUsageType() Tenant_DataUsageType {
	if m != nil {
		return m.UsageType
	}
	return Tenant_DATA_USAGE_TYPE_UNSPECIFIED
}

func init() {
	proto.RegisterType((*Tenant)(nil), "google.cloud.talent.v4beta1.Tenant")
	proto.RegisterEnum("google.cloud.talent.v4beta1.Tenant_DataUsageType", Tenant_DataUsageType_name, Tenant_DataUsageType_value)
}

func init() {
	proto.RegisterFile("google/cloud/talent/v4beta1/tenant.proto", fileDescriptor_tenant_a58ed01536a3d383)
}

var fileDescriptor_tenant_a58ed01536a3d383 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x41, 0x6b, 0xea, 0x40,
	0x14, 0x85, 0xdf, 0xe8, 0x43, 0x9e, 0xf3, 0xde, 0x13, 0x99, 0x6e, 0x44, 0x0b, 0x8a, 0x2b, 0x57,
	0x33, 0xd8, 0x76, 0xd7, 0x55, 0x34, 0x69, 0x08, 0x14, 0x1b, 0x92, 0xb8, 0x68, 0x37, 0x61, 0x8c,
	0xc3, 0x10, 0x48, 0xe6, 0x86, 0x64, 0x52, 0xea, 0xb2, 0x7f, 0xa5, 0x7f, 0xac, 0x7f, 0xa5, 0x38,
	0xa3, 0x8b, 0x42, 0x71, 0x77, 0xef, 0xdc, 0xef, 0xcc, 0x39, 0x1c, 0xbc, 0x90, 0x00, 0xb2, 0x10,
	0x2c, 0x2b, 0xa0, 0xdd, 0x33, 0xcd, 0x0b, 0xa1, 0x34, 0x7b, 0xbd, 0xdb, 0x09, 0xcd, 0x97, 0x4c,
	0x0b, 0xc5, 0x95, 0xa6, 0x55, 0x0d, 0x1a, 0xc8, 0xc4, 0x92, 0xd4, 0x90, 0xd4, 0x92, 0xf4, 0x44,
	0x8e, 0xaf, 0x4f, 0xdf, 0xf0, 0x2a, 0x67, 0x5c, 0x29, 0xd0, 0x5c, 0xe7, 0xa0, 0x1a, 0x2b, 0x1d,
	0x5f, 0x34, 0xc9, 0xa0, 0x2c, 0x41, 0x59, 0x72, 0xfe, 0x89, 0x70, 0x2f, 0x31, 0xae, 0x84, 0xe0,
	0xdf, 0x8a, 0x97, 0x62, 0x84, 0x66, 0x68, 0xd1, 0x8f, 0xcc, 0x4c, 0xa6, 0xf8, 0xaf, 0x78, 0xd3,
	0xa2, 0x56, 0xbc, 0x48, 0xf3, 0xfd, 0xa8, 0x63, 0x4e, 0xf8, 0xfc, 0x14, 0xec, 0x49, 0x88, 0x71,
	0xdb, 0x70, 0x29, 0x52, 0x7d, 0xa8, 0xc4, 0xa8, 0x3b, 0x43, 0x8b, 0xc1, 0xcd, 0x92, 0x5e, 0x48,
	0x4e, 0xad, 0x1b, 0x75, 0xb9, 0xe6, 0xdb, 0xa3, 0x32, 0x39, 0x54, 0x22, 0xea, 0xb7, 0xe7, 0x71,
	0xbe, 0xc1, 0xff, 0xbf, 0xdd, 0xc8, 0x14, 0x4f, 0x5c, 0x27, 0x71, 0xd2, 0x6d, 0xec, 0xf8, 0x5e,
	0x9a, 0x3c, 0x87, 0x5e, 0xba, 0xdd, 0xc4, 0xa1, 0xb7, 0x0e, 0x1e, 0x02, 0xcf, 0x1d, 0xfe, 0x22,
	0x03, 0x8c, 0x1d, 0xdf, 0x8f, 0x3c, 0xdf, 0x49, 0x3c, 0x77, 0x88, 0xc8, 0x3f, 0xfc, 0x27, 0x88,
	0x9f, 0x1e, 0xcd, 0xd6, 0x59, 0xbd, 0x23, 0x3c, 0xcd, 0xa0, 0xbc, 0x94, 0x69, 0x75, 0x65, 0x43,
	0x45, 0xa2, 0x81, 0xb6, 0xce, 0x44, 0x78, 0xac, 0x26, 0x44, 0x2f, 0xce, 0x49, 0x23, 0xa1, 0xe0,
	0x4a, 0x52, 0xa8, 0x25, 0x93, 0x42, 0x99, 0xe2, 0x98, 0x3d, 0xf1, 0x2a, 0x6f, 0x7e, 0x6c, 0xf9,
	0xde, 0xae, 0x1f, 0x9d, 0xee, 0x3a, 0x89, 0x77, 0x3d, 0xa3, 0xb9, 0xfd, 0x0a, 0x00, 0x00, 0xff,
	0xff, 0xd9, 0xf1, 0x67, 0xd5, 0xfd, 0x01, 0x00, 0x00,
}