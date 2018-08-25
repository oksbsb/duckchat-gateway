// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core/plugin.proto

package core // import "github.com/duckchat/duckchat-gateway/proto/core"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PluginUsageType int32

const (
	PluginUsageType_PluginUsageNone         PluginUsageType = 0
	PluginUsageType_PluginUsageIndex        PluginUsageType = 1
	PluginUsageType_PluginUsageLogin        PluginUsageType = 2
	PluginUsageType_PluginUsageU2Message    PluginUsageType = 3
	PluginUsageType_PluginUsageTmpMessage   PluginUsageType = 4
	PluginUsageType_PluginUsageGroupMessage PluginUsageType = 5
)

var PluginUsageType_name = map[int32]string{
	0: "PluginUsageNone",
	1: "PluginUsageIndex",
	2: "PluginUsageLogin",
	3: "PluginUsageU2Message",
	4: "PluginUsageTmpMessage",
	5: "PluginUsageGroupMessage",
}
var PluginUsageType_value = map[string]int32{
	"PluginUsageNone":         0,
	"PluginUsageIndex":        1,
	"PluginUsageLogin":        2,
	"PluginUsageU2Message":    3,
	"PluginUsageTmpMessage":   4,
	"PluginUsageGroupMessage": 5,
}

func (x PluginUsageType) String() string {
	return proto.EnumName(PluginUsageType_name, int32(x))
}
func (PluginUsageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_plugin_227f98d3e1a3a9d3, []int{0}
}

type PluginPermissionType int32

const (
	PluginPermissionType_PluginPermissionAdminOnly   PluginPermissionType = 0
	PluginPermissionType_PluginPermissionAll         PluginPermissionType = 1
	PluginPermissionType_PluginPermissionGroupMaster PluginPermissionType = 2
)

var PluginPermissionType_name = map[int32]string{
	0: "PluginPermissionAdminOnly",
	1: "PluginPermissionAll",
	2: "PluginPermissionGroupMaster",
}
var PluginPermissionType_value = map[string]int32{
	"PluginPermissionAdminOnly":   0,
	"PluginPermissionAll":         1,
	"PluginPermissionGroupMaster": 2,
}

func (x PluginPermissionType) String() string {
	return proto.EnumName(PluginPermissionType_name, int32(x))
}
func (PluginPermissionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_plugin_227f98d3e1a3a9d3, []int{1}
}

type PluginLoadingType int32

const (
	PluginLoadingType_PluginLoadingNewPage    PluginLoadingType = 0
	PluginLoadingType_PluginLoadingFloat      PluginLoadingType = 1
	PluginLoadingType_PluginLoadingMask       PluginLoadingType = 2
	PluginLoadingType_PluginLoadingChatbox    PluginLoadingType = 3
	PluginLoadingType_PluginLoadingFullScreen PluginLoadingType = 4
)

var PluginLoadingType_name = map[int32]string{
	0: "PluginLoadingNewPage",
	1: "PluginLoadingFloat",
	2: "PluginLoadingMask",
	3: "PluginLoadingChatbox",
	4: "PluginLoadingFullScreen",
}
var PluginLoadingType_value = map[string]int32{
	"PluginLoadingNewPage":    0,
	"PluginLoadingFloat":      1,
	"PluginLoadingMask":       2,
	"PluginLoadingChatbox":    3,
	"PluginLoadingFullScreen": 4,
}

func (x PluginLoadingType) String() string {
	return proto.EnumName(PluginLoadingType_name, int32(x))
}
func (PluginLoadingType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_plugin_227f98d3e1a3a9d3, []int{2}
}

type PluginProfile struct {
	Id                   int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Logo                 string               `protobuf:"bytes,3,opt,name=logo,proto3" json:"logo,omitempty"`
	Order                int32                `protobuf:"varint,4,opt,name=order,proto3" json:"order,omitempty"`
	LandingPageUrl       string               `protobuf:"bytes,5,opt,name=landingPageUrl,proto3" json:"landingPageUrl,omitempty"`
	LandingPageWithProxy bool                 `protobuf:"varint,6,opt,name=landingPageWithProxy,proto3" json:"landingPageWithProxy,omitempty"`
	UsageTypes           []PluginUsageType    `protobuf:"varint,7,rep,packed,name=usageTypes,proto3,enum=core.PluginUsageType" json:"usageTypes,omitempty"`
	LoadingType          PluginLoadingType    `protobuf:"varint,8,opt,name=loadingType,proto3,enum=core.PluginLoadingType" json:"loadingType,omitempty"`
	PermissionType       PluginPermissionType `protobuf:"varint,9,opt,name=permissionType,proto3,enum=core.PluginPermissionType" json:"permissionType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PluginProfile) Reset()         { *m = PluginProfile{} }
func (m *PluginProfile) String() string { return proto.CompactTextString(m) }
func (*PluginProfile) ProtoMessage()    {}
func (*PluginProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_plugin_227f98d3e1a3a9d3, []int{0}
}
func (m *PluginProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginProfile.Unmarshal(m, b)
}
func (m *PluginProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginProfile.Marshal(b, m, deterministic)
}
func (dst *PluginProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginProfile.Merge(dst, src)
}
func (m *PluginProfile) XXX_Size() int {
	return xxx_messageInfo_PluginProfile.Size(m)
}
func (m *PluginProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginProfile.DiscardUnknown(m)
}

var xxx_messageInfo_PluginProfile proto.InternalMessageInfo

func (m *PluginProfile) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PluginProfile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PluginProfile) GetLogo() string {
	if m != nil {
		return m.Logo
	}
	return ""
}

func (m *PluginProfile) GetOrder() int32 {
	if m != nil {
		return m.Order
	}
	return 0
}

func (m *PluginProfile) GetLandingPageUrl() string {
	if m != nil {
		return m.LandingPageUrl
	}
	return ""
}

func (m *PluginProfile) GetLandingPageWithProxy() bool {
	if m != nil {
		return m.LandingPageWithProxy
	}
	return false
}

func (m *PluginProfile) GetUsageTypes() []PluginUsageType {
	if m != nil {
		return m.UsageTypes
	}
	return nil
}

func (m *PluginProfile) GetLoadingType() PluginLoadingType {
	if m != nil {
		return m.LoadingType
	}
	return PluginLoadingType_PluginLoadingNewPage
}

func (m *PluginProfile) GetPermissionType() PluginPermissionType {
	if m != nil {
		return m.PermissionType
	}
	return PluginPermissionType_PluginPermissionAdminOnly
}

func init() {
	proto.RegisterType((*PluginProfile)(nil), "core.PluginProfile")
	proto.RegisterEnum("core.PluginUsageType", PluginUsageType_name, PluginUsageType_value)
	proto.RegisterEnum("core.PluginPermissionType", PluginPermissionType_name, PluginPermissionType_value)
	proto.RegisterEnum("core.PluginLoadingType", PluginLoadingType_name, PluginLoadingType_value)
}

func init() { proto.RegisterFile("core/plugin.proto", fileDescriptor_plugin_227f98d3e1a3a9d3) }

var fileDescriptor_plugin_227f98d3e1a3a9d3 = []byte{
	// 489 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x86, 0x9b, 0xb4, 0x1d, 0xdb, 0x41, 0x74, 0xde, 0x69, 0x4b, 0x33, 0x26, 0x44, 0xc5, 0x05,
	0xaa, 0x2a, 0xd1, 0x4a, 0x45, 0x5c, 0x70, 0xc9, 0x26, 0x81, 0x90, 0xb6, 0x11, 0x05, 0x2a, 0xd0,
	0xb4, 0x1b, 0x37, 0x31, 0xa9, 0x35, 0xc7, 0xae, 0x9c, 0x44, 0x6b, 0x78, 0x0d, 0x5e, 0x81, 0xa7,
	0xe1, 0x65, 0x78, 0x05, 0xe4, 0x64, 0x5d, 0xdd, 0x8c, 0xbb, 0xe3, 0xef, 0xf3, 0x7f, 0x9c, 0xe3,
	0x24, 0x70, 0x14, 0x2a, 0xcd, 0xa6, 0x2b, 0x91, 0xc7, 0x5c, 0x4e, 0x56, 0x5a, 0x65, 0x0a, 0x5b,
	0x06, 0xbd, 0xfc, 0xeb, 0xc2, 0x13, 0xbf, 0xc4, 0xbe, 0x56, 0x3f, 0xb8, 0x60, 0xd8, 0x01, 0x97,
	0x47, 0x9e, 0x33, 0x74, 0x46, 0xed, 0xc0, 0xe5, 0x11, 0x22, 0xb4, 0x24, 0x4d, 0x98, 0xe7, 0x0e,
	0x9d, 0xd1, 0x41, 0x50, 0xd6, 0x86, 0x09, 0x15, 0x2b, 0xaf, 0x59, 0x31, 0x53, 0x63, 0x0f, 0xda,
	0x4a, 0x47, 0x4c, 0x7b, 0xad, 0x32, 0x5a, 0x2d, 0xf0, 0x15, 0x74, 0x04, 0x95, 0x11, 0x97, 0xb1,
	0x4f, 0x63, 0x36, 0xd7, 0xc2, 0x6b, 0x97, 0x99, 0x1a, 0xc5, 0x19, 0xf4, 0x2c, 0xf2, 0x8d, 0x67,
	0x4b, 0x5f, 0xab, 0x75, 0xe1, 0xed, 0x0d, 0x9d, 0xd1, 0x7e, 0xf0, 0x5f, 0x87, 0x6f, 0x01, 0xf2,
	0x94, 0xc6, 0xec, 0x6b, 0xb1, 0x62, 0xa9, 0xf7, 0x68, 0xd8, 0x1c, 0x75, 0x66, 0xfd, 0x89, 0x19,
	0x6b, 0x52, 0x8d, 0x34, 0xdf, 0xd8, 0xc0, 0xda, 0x88, 0xef, 0xe0, 0xb1, 0x50, 0xd4, 0xb4, 0x33,
	0x6b, 0x6f, 0x7f, 0xe8, 0x8c, 0x3a, 0xb3, 0x81, 0x9d, 0x3b, 0xdf, 0xea, 0xc0, 0xde, 0x8b, 0xa7,
	0xd0, 0x59, 0x31, 0x9d, 0xf0, 0x34, 0xe5, 0x4a, 0x96, 0xe9, 0x83, 0x32, 0xfd, 0xcc, 0x4e, 0xfb,
	0x3b, 0x3b, 0x82, 0x5a, 0x62, 0xfc, 0xdb, 0x81, 0xc3, 0xda, 0xe3, 0x61, 0x77, 0x07, 0x5d, 0x2a,
	0xc9, 0x48, 0x03, 0x7b, 0x40, 0x2c, 0xf8, 0x49, 0x46, 0x6c, 0x4d, 0x9c, 0x1a, 0x3d, 0x57, 0x31,
	0x97, 0xc4, 0x45, 0x0f, 0x7a, 0x16, 0x9d, 0xcf, 0x2e, 0x58, 0x6a, 0x0a, 0xd2, 0xc4, 0x63, 0xe8,
	0xdb, 0xa7, 0x25, 0xab, 0x8d, 0x6a, 0xe1, 0x09, 0x0c, 0x2c, 0xf5, 0x51, 0xab, 0xfc, 0x5e, 0xb6,
	0xc7, 0x6a, 0xd3, 0x71, 0x77, 0x1c, 0x7c, 0x0e, 0xc7, 0x75, 0xfe, 0x3e, 0x4a, 0xb8, 0xfc, 0x2c,
	0x45, 0x41, 0x1a, 0x38, 0x80, 0xee, 0x03, 0x2d, 0x04, 0x71, 0xf0, 0x05, 0x9c, 0xd4, 0x45, 0x75,
	0x22, 0x4d, 0x33, 0xa6, 0x89, 0x3b, 0xfe, 0xe5, 0xc0, 0xd1, 0x83, 0xeb, 0xdf, 0x0e, 0x76, 0x07,
	0x2f, 0xd9, 0xad, 0xf9, 0x08, 0x48, 0x03, 0x9f, 0x02, 0xee, 0x98, 0x0f, 0x42, 0xd1, 0x8c, 0x38,
	0xd8, 0xaf, 0xb5, 0xb9, 0xa0, 0xe9, 0x8d, 0x7d, 0x43, 0x77, 0xf8, 0x6c, 0x49, 0xb3, 0x85, 0x5a,
	0x93, 0xe6, 0xf6, 0x1a, 0x36, 0x8d, 0x72, 0x21, 0xbe, 0x84, 0x9a, 0x31, 0x49, 0x5a, 0xa7, 0xdf,
	0xa1, 0x1b, 0xaa, 0x64, 0xf2, 0x93, 0x8a, 0xa2, 0xfa, 0x6f, 0xca, 0x37, 0x7d, 0x35, 0x8d, 0x79,
	0xb6, 0xcc, 0x17, 0x93, 0x50, 0x25, 0xd3, 0x28, 0x0f, 0x6f, 0xc2, 0x25, 0xcd, 0xee, 0x8b, 0xd7,
	0x31, 0xcd, 0xd8, 0x2d, 0x2d, 0xa6, 0x65, 0x60, 0x6a, 0x02, 0x7f, 0xdc, 0xc3, 0x2b, 0x2a, 0x8a,
	0x6b, 0xdf, 0x90, 0xeb, 0x33, 0xa5, 0xd9, 0x62, 0xaf, 0xb4, 0x6f, 0xfe, 0x05, 0x00, 0x00, 0xff,
	0xff, 0x25, 0x0f, 0x1a, 0x88, 0x9b, 0x03, 0x00, 0x00,
}