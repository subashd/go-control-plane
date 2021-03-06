// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/aws_request_signing/v2alpha/aws_request_signing.proto

package envoy_config_filter_http_aws_request_signing_v2alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

type AwsRequestSigning struct {
	ServiceName          string   `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	Region               string   `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AwsRequestSigning) Reset()         { *m = AwsRequestSigning{} }
func (m *AwsRequestSigning) String() string { return proto.CompactTextString(m) }
func (*AwsRequestSigning) ProtoMessage()    {}
func (*AwsRequestSigning) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebccdae368886250, []int{0}
}

func (m *AwsRequestSigning) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AwsRequestSigning.Unmarshal(m, b)
}
func (m *AwsRequestSigning) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AwsRequestSigning.Marshal(b, m, deterministic)
}
func (m *AwsRequestSigning) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AwsRequestSigning.Merge(m, src)
}
func (m *AwsRequestSigning) XXX_Size() int {
	return xxx_messageInfo_AwsRequestSigning.Size(m)
}
func (m *AwsRequestSigning) XXX_DiscardUnknown() {
	xxx_messageInfo_AwsRequestSigning.DiscardUnknown(m)
}

var xxx_messageInfo_AwsRequestSigning proto.InternalMessageInfo

func (m *AwsRequestSigning) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *AwsRequestSigning) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func init() {
	proto.RegisterType((*AwsRequestSigning)(nil), "envoy.config.filter.http.aws_request_signing.v2alpha.AwsRequestSigning")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/aws_request_signing/v2alpha/aws_request_signing.proto", fileDescriptor_ebccdae368886250)
}

var fileDescriptor_ebccdae368886250 = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x95, 0x0a, 0x8a, 0x08, 0x2c, 0x64, 0x80, 0xaa, 0x03, 0x54, 0x4c, 0x88, 0xc1, 0x96,
	0x28, 0x62, 0x27, 0x0f, 0x50, 0x55, 0xe5, 0x01, 0xc2, 0xd1, 0x5e, 0xd3, 0x93, 0x92, 0xb3, 0xb1,
	0xaf, 0x69, 0xfb, 0x14, 0xac, 0x88, 0xc7, 0x64, 0x64, 0x40, 0x28, 0x71, 0x3a, 0x20, 0x2a, 0x86,
	0x6e, 0x96, 0x7f, 0xff, 0x9f, 0xfc, 0xfd, 0xf1, 0x08, 0xb9, 0x32, 0x1b, 0x3d, 0x35, 0x3c, 0xa7,
	0x5c, 0xcf, 0xa9, 0x10, 0x74, 0x7a, 0x21, 0x62, 0x35, 0xac, 0x7c, 0xe6, 0xf0, 0x75, 0x89, 0x5e,
	0x32, 0x4f, 0x39, 0x13, 0xe7, 0xba, 0xba, 0x83, 0xc2, 0x2e, 0x60, 0x57, 0xa6, 0xac, 0x33, 0x62,
	0x92, 0xfb, 0x86, 0xa7, 0x02, 0x4f, 0x05, 0x9e, 0xaa, 0x79, 0x6a, 0x57, 0xa7, 0xe5, 0xf5, 0x2f,
	0x97, 0x33, 0x0b, 0x1a, 0x98, 0x8d, 0x80, 0x90, 0x61, 0xaf, 0x4b, 0xca, 0x1d, 0x08, 0x06, 0x6a,
	0xff, 0xa2, 0x82, 0x82, 0x66, 0x20, 0xa8, 0xb7, 0x87, 0x10, 0x5c, 0x3f, 0xc7, 0x67, 0x8f, 0x2b,
	0x3f, 0x09, 0xd8, 0xa7, 0x40, 0x4d, 0x6e, 0xe3, 0x53, 0x8f, 0xae, 0xa2, 0x29, 0x66, 0x0c, 0x25,
	0xf6, 0xa2, 0x41, 0x74, 0x73, 0x9c, 0x1e, 0x7d, 0xa5, 0x07, 0xae, 0x33, 0x88, 0x26, 0x27, 0x6d,
	0x38, 0x82, 0x12, 0x93, 0xab, 0xb8, 0xeb, 0x30, 0x27, 0xc3, 0xbd, 0xce, 0xef, 0x57, 0xed, 0x75,
	0xfa, 0x11, 0x7d, 0xbe, 0x7f, 0xbf, 0x1d, 0x3e, 0x6c, 0xcd, 0x70, 0x2d, 0xc8, 0xbe, 0xfe, 0x63,
	0x6b, 0xe7, 0xff, 0xd1, 0x1b, 0xc6, 0x29, 0x19, 0xd5, 0x14, 0xad, 0x33, 0xeb, 0x8d, 0xda, 0x67,
	0x9d, 0xf4, 0xfc, 0x8f, 0xe2, 0xb8, 0x96, 0x1f, 0x47, 0x2f, 0xdd, 0x66, 0x85, 0xe1, 0x4f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xd5, 0x41, 0xc2, 0xaa, 0xc6, 0x01, 0x00, 0x00,
}
