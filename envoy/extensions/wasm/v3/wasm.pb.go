// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/wasm/v3/wasm.proto

package envoy_extensions_wasm_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

type VmConfig struct {
	VmId                 string              `protobuf:"bytes,1,opt,name=vm_id,json=vmId,proto3" json:"vm_id,omitempty"`
	Runtime              string              `protobuf:"bytes,2,opt,name=runtime,proto3" json:"runtime,omitempty"`
	Code                 *v3.AsyncDataSource `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Configuration        *any.Any            `protobuf:"bytes,4,opt,name=configuration,proto3" json:"configuration,omitempty"`
	AllowPrecompiled     bool                `protobuf:"varint,5,opt,name=allow_precompiled,json=allowPrecompiled,proto3" json:"allow_precompiled,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *VmConfig) Reset()         { *m = VmConfig{} }
func (m *VmConfig) String() string { return proto.CompactTextString(m) }
func (*VmConfig) ProtoMessage()    {}
func (*VmConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb6accf9c32c7c2b, []int{0}
}

func (m *VmConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VmConfig.Unmarshal(m, b)
}
func (m *VmConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VmConfig.Marshal(b, m, deterministic)
}
func (m *VmConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VmConfig.Merge(m, src)
}
func (m *VmConfig) XXX_Size() int {
	return xxx_messageInfo_VmConfig.Size(m)
}
func (m *VmConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_VmConfig.DiscardUnknown(m)
}

var xxx_messageInfo_VmConfig proto.InternalMessageInfo

func (m *VmConfig) GetVmId() string {
	if m != nil {
		return m.VmId
	}
	return ""
}

func (m *VmConfig) GetRuntime() string {
	if m != nil {
		return m.Runtime
	}
	return ""
}

func (m *VmConfig) GetCode() *v3.AsyncDataSource {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *VmConfig) GetConfiguration() *any.Any {
	if m != nil {
		return m.Configuration
	}
	return nil
}

func (m *VmConfig) GetAllowPrecompiled() bool {
	if m != nil {
		return m.AllowPrecompiled
	}
	return false
}

type PluginConfig struct {
	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	GroupName string `protobuf:"bytes,2,opt,name=group_name,json=groupName,proto3" json:"group_name,omitempty"`
	// Types that are valid to be assigned to VmConfig:
	//	*PluginConfig_InlineVmConfig
	VmConfig             isPluginConfig_VmConfig `protobuf_oneof:"vm_config"`
	Configuration        *any.Any                `protobuf:"bytes,5,opt,name=configuration,proto3" json:"configuration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *PluginConfig) Reset()         { *m = PluginConfig{} }
func (m *PluginConfig) String() string { return proto.CompactTextString(m) }
func (*PluginConfig) ProtoMessage()    {}
func (*PluginConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb6accf9c32c7c2b, []int{1}
}

func (m *PluginConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginConfig.Unmarshal(m, b)
}
func (m *PluginConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginConfig.Marshal(b, m, deterministic)
}
func (m *PluginConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginConfig.Merge(m, src)
}
func (m *PluginConfig) XXX_Size() int {
	return xxx_messageInfo_PluginConfig.Size(m)
}
func (m *PluginConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginConfig.DiscardUnknown(m)
}

var xxx_messageInfo_PluginConfig proto.InternalMessageInfo

func (m *PluginConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PluginConfig) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

type isPluginConfig_VmConfig interface {
	isPluginConfig_VmConfig()
}

type PluginConfig_InlineVmConfig struct {
	InlineVmConfig *VmConfig `protobuf:"bytes,3,opt,name=inline_vm_config,json=inlineVmConfig,proto3,oneof"`
}

func (*PluginConfig_InlineVmConfig) isPluginConfig_VmConfig() {}

func (m *PluginConfig) GetVmConfig() isPluginConfig_VmConfig {
	if m != nil {
		return m.VmConfig
	}
	return nil
}

func (m *PluginConfig) GetInlineVmConfig() *VmConfig {
	if x, ok := m.GetVmConfig().(*PluginConfig_InlineVmConfig); ok {
		return x.InlineVmConfig
	}
	return nil
}

func (m *PluginConfig) GetConfiguration() *any.Any {
	if m != nil {
		return m.Configuration
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PluginConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PluginConfig_InlineVmConfig)(nil),
	}
}

type WasmService struct {
	Config               *PluginConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	Singleton            bool          `protobuf:"varint,2,opt,name=singleton,proto3" json:"singleton,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *WasmService) Reset()         { *m = WasmService{} }
func (m *WasmService) String() string { return proto.CompactTextString(m) }
func (*WasmService) ProtoMessage()    {}
func (*WasmService) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb6accf9c32c7c2b, []int{2}
}

func (m *WasmService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WasmService.Unmarshal(m, b)
}
func (m *WasmService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WasmService.Marshal(b, m, deterministic)
}
func (m *WasmService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WasmService.Merge(m, src)
}
func (m *WasmService) XXX_Size() int {
	return xxx_messageInfo_WasmService.Size(m)
}
func (m *WasmService) XXX_DiscardUnknown() {
	xxx_messageInfo_WasmService.DiscardUnknown(m)
}

var xxx_messageInfo_WasmService proto.InternalMessageInfo

func (m *WasmService) GetConfig() *PluginConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *WasmService) GetSingleton() bool {
	if m != nil {
		return m.Singleton
	}
	return false
}

func init() {
	proto.RegisterType((*VmConfig)(nil), "envoy.extensions.wasm.v3.VmConfig")
	proto.RegisterType((*PluginConfig)(nil), "envoy.extensions.wasm.v3.PluginConfig")
	proto.RegisterType((*WasmService)(nil), "envoy.extensions.wasm.v3.WasmService")
}

func init() {
	proto.RegisterFile("envoy/extensions/wasm/v3/wasm.proto", fileDescriptor_fb6accf9c32c7c2b)
}

var fileDescriptor_fb6accf9c32c7c2b = []byte{
	// 505 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x6e, 0xd3, 0x4c,
	0x14, 0xc5, 0x3f, 0xe7, 0x4b, 0xda, 0xe4, 0x06, 0x50, 0x19, 0x90, 0x30, 0x15, 0x7f, 0xd2, 0x40,
	0x43, 0x10, 0x30, 0x96, 0x1a, 0xb1, 0x20, 0x0b, 0xa4, 0x1a, 0x16, 0xb0, 0xa9, 0x22, 0x57, 0x82,
	0xa5, 0x35, 0xb1, 0xa7, 0x66, 0x24, 0x7b, 0xc6, 0x1a, 0xdb, 0xd3, 0xfa, 0x0d, 0x60, 0xc9, 0x12,
	0xde, 0x87, 0x97, 0x62, 0x85, 0x66, 0xc6, 0x4e, 0x52, 0x21, 0x4b, 0xac, 0x3c, 0xbe, 0xf7, 0x9e,
	0xb9, 0xe7, 0xfc, 0x6c, 0x78, 0x42, 0xb9, 0x12, 0xb5, 0x47, 0xaf, 0x4a, 0xca, 0x0b, 0x26, 0x78,
	0xe1, 0x5d, 0x92, 0x22, 0xf3, 0xd4, 0xc2, 0x3c, 0x71, 0x2e, 0x45, 0x29, 0x90, 0x6b, 0x86, 0xf0,
	0x76, 0x08, 0x9b, 0xa6, 0x5a, 0x1c, 0x3e, 0xb6, 0xf2, 0x48, 0xf0, 0x0b, 0x96, 0x78, 0x91, 0x90,
	0x54, 0x4b, 0xd7, 0xa4, 0xa0, 0x56, 0x7a, 0x78, 0x3f, 0x11, 0x22, 0x49, 0xa9, 0x67, 0xde, 0xd6,
	0xd5, 0x85, 0x47, 0x78, 0xdd, 0xb4, 0x8e, 0xaa, 0x38, 0x27, 0x1e, 0xe1, 0x5c, 0x94, 0xa4, 0x34,
	0xab, 0x15, 0x95, 0xfa, 0x7a, 0xc6, 0x93, 0x66, 0xe4, 0x9e, 0x22, 0x29, 0x8b, 0x49, 0x49, 0xbd,
	0xf6, 0x60, 0x1b, 0xd3, 0xef, 0x3d, 0x18, 0x7e, 0xca, 0xde, 0x99, 0xb5, 0xe8, 0x0e, 0x0c, 0x54,
	0x16, 0xb2, 0xd8, 0x75, 0x26, 0xce, 0x7c, 0x14, 0xf4, 0x55, 0xf6, 0x31, 0x46, 0x47, 0xb0, 0x2f,
	0x2b, 0x5e, 0xb2, 0x8c, 0xba, 0x3d, 0x5d, 0xf6, 0xf7, 0x7f, 0xfb, 0x7d, 0xd9, 0x9b, 0x38, 0x41,
	0x5b, 0x47, 0x6f, 0xa0, 0x1f, 0x89, 0x98, 0xba, 0xff, 0x4f, 0x9c, 0xf9, 0xf8, 0xe4, 0x18, 0xdb,
	0x94, 0x36, 0x0b, 0xd6, 0x59, 0xb0, 0x5a, 0xe0, 0xd3, 0xa2, 0xe6, 0xd1, 0x7b, 0x52, 0x92, 0x73,
	0x51, 0xc9, 0x88, 0x06, 0x46, 0x82, 0x96, 0x70, 0xd3, 0xce, 0x55, 0xd2, 0x78, 0x77, 0xfb, 0xe6,
	0x8e, 0xbb, 0xd8, 0xc6, 0xc5, 0x6d, 0x5c, 0x7c, 0xca, 0xeb, 0xe0, 0xfa, 0x28, 0x7a, 0x01, 0xb7,
	0x49, 0x9a, 0x8a, 0xcb, 0x30, 0x97, 0x34, 0x12, 0x59, 0xce, 0x52, 0x1a, 0xbb, 0x83, 0x89, 0x33,
	0x1f, 0x06, 0x07, 0xa6, 0xb1, 0xda, 0xd6, 0x97, 0xcf, 0x7f, 0xfe, 0xfa, 0xfa, 0xe8, 0x29, 0x4c,
	0xaf, 0x79, 0xb3, 0xf4, 0x4f, 0x48, 0x9a, 0x7f, 0x21, 0xb8, 0xc5, 0x30, 0xfd, 0xd6, 0x83, 0x1b,
	0xab, 0xb4, 0x4a, 0x18, 0x6f, 0xb8, 0x20, 0xe8, 0x73, 0x92, 0xd1, 0x16, 0x8b, 0x3e, 0xa3, 0x87,
	0x00, 0x89, 0x14, 0x55, 0x1e, 0x9a, 0x8e, 0x21, 0x13, 0x8c, 0x4c, 0xe5, 0x4c, 0xb7, 0xcf, 0xe0,
	0x80, 0xf1, 0x94, 0x71, 0x1a, 0xaa, 0x2c, 0xb4, 0xdb, 0x1a, 0x3c, 0x53, 0xdc, 0xf5, 0x13, 0x6c,
	0x1c, 0x7c, 0xf8, 0x2f, 0xb8, 0x65, 0xd5, 0x9b, 0x4f, 0xf3, 0x17, 0xa7, 0xc1, 0x3f, 0x73, 0x5a,
	0xbe, 0xd2, 0xd1, 0xe7, 0x30, 0xeb, 0x8e, 0xbe, 0x9b, 0xd6, 0x1f, 0xc3, 0x68, 0xe3, 0x79, 0xfa,
	0xc3, 0x81, 0xf1, 0x67, 0x52, 0x64, 0xe7, 0x54, 0x2a, 0x16, 0x51, 0xf4, 0x16, 0xf6, 0x9a, 0x34,
	0x8e, 0x31, 0x30, 0xeb, 0x4e, 0xb3, 0x7b, 0x69, 0xd0, 0xa8, 0xd0, 0x03, 0x18, 0x15, 0x8c, 0x27,
	0x29, 0x2d, 0x05, 0x37, 0xd4, 0x86, 0xc1, 0xb6, 0xb0, 0x7c, 0xa9, 0x9d, 0x3e, 0x83, 0xe3, 0x6e,
	0xa7, 0x3b, 0x5e, 0xfc, 0xd7, 0x30, 0x63, 0xc2, 0xee, 0xcf, 0xa5, 0xb8, 0xaa, 0x3b, 0xad, 0xf8,
	0x23, 0x2d, 0x5b, 0x69, 0x44, 0x2b, 0x67, 0xbd, 0x67, 0x58, 0x2d, 0xfe, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x0e, 0xc9, 0xd7, 0x81, 0xb2, 0x03, 0x00, 0x00,
}
