// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/keyword_plan_campaign_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v0/errors"

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

// Enum describing possible errors from applying a keyword plan campaign.
type KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError int32

const (
	// Enum unspecified.
	KeywordPlanCampaignErrorEnum_UNSPECIFIED KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError = 0
	// The received error code is not known in this version.
	KeywordPlanCampaignErrorEnum_UNKNOWN KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError = 1
	// A keyword plan campaign name is missing, empty, longer than allowed limit
	// or contains invalid chars.
	KeywordPlanCampaignErrorEnum_INVALID_NAME KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError = 2
	// A keyword plan campaign contains one or more untargetable languages.
	KeywordPlanCampaignErrorEnum_INVALID_LANGUAGES KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError = 3
	// A keyword plan campaign contains one or more invalid geo targets.
	KeywordPlanCampaignErrorEnum_INVALID_GEOS KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError = 4
	// The keyword plan campaign name is duplicate to an existing keyword plan
	// campaign name or other keyword plan campaign name in the request.
	KeywordPlanCampaignErrorEnum_DUPLICATE_NAME KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError = 5
	// The number of geo targets in the keyword plan campaign exceeds limits.
	KeywordPlanCampaignErrorEnum_MAX_GEOS_EXCEEDED KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError = 6
)

var KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "INVALID_NAME",
	3: "INVALID_LANGUAGES",
	4: "INVALID_GEOS",
	5: "DUPLICATE_NAME",
	6: "MAX_GEOS_EXCEEDED",
}
var KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError_value = map[string]int32{
	"UNSPECIFIED":       0,
	"UNKNOWN":           1,
	"INVALID_NAME":      2,
	"INVALID_LANGUAGES": 3,
	"INVALID_GEOS":      4,
	"DUPLICATE_NAME":    5,
	"MAX_GEOS_EXCEEDED": 6,
}

func (x KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError) String() string {
	return proto.EnumName(KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError_name, int32(x))
}
func (KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_keyword_plan_campaign_error_2abe709d47e6a6ba, []int{0, 0}
}

// Container for enum describing possible errors from applying a keyword plan
// campaign.
type KeywordPlanCampaignErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeywordPlanCampaignErrorEnum) Reset()         { *m = KeywordPlanCampaignErrorEnum{} }
func (m *KeywordPlanCampaignErrorEnum) String() string { return proto.CompactTextString(m) }
func (*KeywordPlanCampaignErrorEnum) ProtoMessage()    {}
func (*KeywordPlanCampaignErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_keyword_plan_campaign_error_2abe709d47e6a6ba, []int{0}
}
func (m *KeywordPlanCampaignErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPlanCampaignErrorEnum.Unmarshal(m, b)
}
func (m *KeywordPlanCampaignErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPlanCampaignErrorEnum.Marshal(b, m, deterministic)
}
func (dst *KeywordPlanCampaignErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPlanCampaignErrorEnum.Merge(dst, src)
}
func (m *KeywordPlanCampaignErrorEnum) XXX_Size() int {
	return xxx_messageInfo_KeywordPlanCampaignErrorEnum.Size(m)
}
func (m *KeywordPlanCampaignErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPlanCampaignErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPlanCampaignErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*KeywordPlanCampaignErrorEnum)(nil), "google.ads.googleads.v0.errors.KeywordPlanCampaignErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError", KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError_name, KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/keyword_plan_campaign_error.proto", fileDescriptor_keyword_plan_campaign_error_2abe709d47e6a6ba)
}

var fileDescriptor_keyword_plan_campaign_error_2abe709d47e6a6ba = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xed, 0xa6, 0x13, 0x32, 0xd1, 0x18, 0x10, 0x3c, 0xe8, 0x0e, 0xfb, 0x00, 0x69, 0xc1,
	0xa3, 0x17, 0xb3, 0x36, 0x96, 0xb2, 0x2d, 0x2b, 0xcc, 0xce, 0x21, 0x85, 0x12, 0xd7, 0x12, 0x86,
	0x5d, 0x53, 0x12, 0x9d, 0xf8, 0x65, 0x3c, 0x78, 0xf4, 0x4b, 0x78, 0xf7, 0xea, 0x17, 0x92, 0x36,
	0xdb, 0xf0, 0x52, 0x4f, 0xfd, 0xd3, 0xf7, 0x7b, 0xbf, 0xc7, 0x7b, 0x01, 0x37, 0x42, 0x4a, 0x91,
	0x67, 0x36, 0x4f, 0xb5, 0x6d, 0x62, 0x95, 0xd6, 0x8e, 0x9d, 0x29, 0x25, 0x95, 0xb6, 0x9f, 0xb2,
	0xb7, 0x57, 0xa9, 0xd2, 0xa4, 0xcc, 0x79, 0x91, 0x2c, 0xf8, 0xaa, 0xe4, 0x4b, 0x51, 0x24, 0x75,
	0x11, 0x97, 0x4a, 0x3e, 0x4b, 0xd4, 0x33, 0x6d, 0x98, 0xa7, 0x1a, 0xef, 0x0c, 0x78, 0xed, 0x60,
	0x63, 0xe8, 0x7f, 0x59, 0xe0, 0x62, 0x68, 0x2c, 0x61, 0xce, 0x0b, 0x77, 0xe3, 0xa0, 0x55, 0x95,
	0x16, 0x2f, 0xab, 0xfe, 0xbb, 0x05, 0xce, 0x9b, 0x00, 0x74, 0x02, 0xba, 0x11, 0x9b, 0x86, 0xd4,
	0x0d, 0x6e, 0x03, 0xea, 0xc1, 0x3d, 0xd4, 0x05, 0x87, 0x11, 0x1b, 0xb2, 0xc9, 0x3d, 0x83, 0x16,
	0x82, 0xe0, 0x28, 0x60, 0x33, 0x32, 0x0a, 0xbc, 0x84, 0x91, 0x31, 0x85, 0x2d, 0x74, 0x06, 0x4e,
	0xb7, 0x7f, 0x46, 0x84, 0xf9, 0x11, 0xf1, 0xe9, 0x14, 0xb6, 0xff, 0x82, 0x3e, 0x9d, 0x4c, 0xe1,
	0x3e, 0x42, 0xe0, 0xd8, 0x8b, 0xc2, 0x51, 0xe0, 0x92, 0x3b, 0x6a, 0x9a, 0x0f, 0xaa, 0xe6, 0x31,
	0x99, 0xd7, 0x44, 0x42, 0xe7, 0x2e, 0xa5, 0x1e, 0xf5, 0x60, 0x67, 0xf0, 0x63, 0x81, 0xfe, 0x42,
	0xae, 0xf0, 0xff, 0x8b, 0x0e, 0x2e, 0x9b, 0x96, 0x08, 0xab, 0x3b, 0x85, 0xd6, 0x83, 0xb7, 0x11,
	0x08, 0x99, 0xf3, 0x42, 0x60, 0xa9, 0x84, 0x2d, 0xb2, 0xa2, 0xbe, 0xe2, 0xf6, 0xf6, 0xe5, 0x52,
	0x37, 0x3d, 0xc5, 0xb5, 0xf9, 0x7c, 0xb4, 0xda, 0x3e, 0x21, 0x9f, 0xad, 0x9e, 0x6f, 0x64, 0x24,
	0xd5, 0xd8, 0xc4, 0x2a, 0xcd, 0x1c, 0x5c, 0x8f, 0xd4, 0xdf, 0x5b, 0x20, 0x26, 0xa9, 0x8e, 0x77,
	0x40, 0x3c, 0x73, 0x62, 0x03, 0x3c, 0x76, 0xea, 0xc1, 0x57, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x73, 0x6c, 0xa6, 0x10, 0x02, 0x02, 0x00, 0x00,
}
