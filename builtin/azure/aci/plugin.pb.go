// Code generated by protoc-gen-go. DO NOT EDIT.
// source: waypoint/builtin/azure/aci/plugin.proto

package aci

import (
	fmt "fmt"
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

type Deployment struct {
	Url                  string                     `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Id                   string                     `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	ContainerGroup       *Deployment_ContainerGroup `protobuf:"bytes,3,opt,name=container_group,json=containerGroup,proto3" json:"container_group,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *Deployment) Reset()         { *m = Deployment{} }
func (m *Deployment) String() string { return proto.CompactTextString(m) }
func (*Deployment) ProtoMessage()    {}
func (*Deployment) Descriptor() ([]byte, []int) {
	return fileDescriptor_936458bb1651e71c, []int{0}
}

func (m *Deployment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Deployment.Unmarshal(m, b)
}
func (m *Deployment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Deployment.Marshal(b, m, deterministic)
}
func (m *Deployment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deployment.Merge(m, src)
}
func (m *Deployment) XXX_Size() int {
	return xxx_messageInfo_Deployment.Size(m)
}
func (m *Deployment) XXX_DiscardUnknown() {
	xxx_messageInfo_Deployment.DiscardUnknown(m)
}

var xxx_messageInfo_Deployment proto.InternalMessageInfo

func (m *Deployment) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Deployment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Deployment) GetContainerGroup() *Deployment_ContainerGroup {
	if m != nil {
		return m.ContainerGroup
	}
	return nil
}

type Deployment_ContainerGroup struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ResourceGroup        string   `protobuf:"bytes,2,opt,name=resource_group,json=resourceGroup,proto3" json:"resource_group,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Deployment_ContainerGroup) Reset()         { *m = Deployment_ContainerGroup{} }
func (m *Deployment_ContainerGroup) String() string { return proto.CompactTextString(m) }
func (*Deployment_ContainerGroup) ProtoMessage()    {}
func (*Deployment_ContainerGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_936458bb1651e71c, []int{0, 0}
}

func (m *Deployment_ContainerGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Deployment_ContainerGroup.Unmarshal(m, b)
}
func (m *Deployment_ContainerGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Deployment_ContainerGroup.Marshal(b, m, deterministic)
}
func (m *Deployment_ContainerGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deployment_ContainerGroup.Merge(m, src)
}
func (m *Deployment_ContainerGroup) XXX_Size() int {
	return xxx_messageInfo_Deployment_ContainerGroup.Size(m)
}
func (m *Deployment_ContainerGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_Deployment_ContainerGroup.DiscardUnknown(m)
}

var xxx_messageInfo_Deployment_ContainerGroup proto.InternalMessageInfo

func (m *Deployment_ContainerGroup) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Deployment_ContainerGroup) GetResourceGroup() string {
	if m != nil {
		return m.ResourceGroup
	}
	return ""
}

type Release struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Release) Reset()         { *m = Release{} }
func (m *Release) String() string { return proto.CompactTextString(m) }
func (*Release) ProtoMessage()    {}
func (*Release) Descriptor() ([]byte, []int) {
	return fileDescriptor_936458bb1651e71c, []int{1}
}

func (m *Release) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Release.Unmarshal(m, b)
}
func (m *Release) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Release.Marshal(b, m, deterministic)
}
func (m *Release) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Release.Merge(m, src)
}
func (m *Release) XXX_Size() int {
	return xxx_messageInfo_Release.Size(m)
}
func (m *Release) XXX_DiscardUnknown() {
	xxx_messageInfo_Release.DiscardUnknown(m)
}

var xxx_messageInfo_Release proto.InternalMessageInfo

func (m *Release) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func init() {
	proto.RegisterType((*Deployment)(nil), "azure.aci.Deployment")
	proto.RegisterType((*Deployment_ContainerGroup)(nil), "azure.aci.Deployment.ContainerGroup")
	proto.RegisterType((*Release)(nil), "azure.aci.Release")
}

func init() {
	proto.RegisterFile("waypoint/builtin/azure/aci/plugin.proto", fileDescriptor_936458bb1651e71c)
}

var fileDescriptor_936458bb1651e71c = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x50, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x25, 0x5b, 0x51, 0x3a, 0x62, 0x94, 0x9c, 0x96, 0xea, 0xa1, 0x14, 0xc5, 0x9e, 0x12, 0xd0,
	0x3f, 0x50, 0xc1, 0x83, 0x78, 0xd9, 0xa3, 0x17, 0x49, 0xd3, 0xa1, 0x0c, 0xa4, 0x49, 0xc8, 0x26,
	0xc8, 0xfa, 0x89, 0x7e, 0x95, 0x6c, 0xdc, 0x55, 0x16, 0xe9, 0xed, 0xcd, 0xe3, 0xcd, 0x9b, 0x37,
	0x0f, 0x6e, 0x3f, 0x74, 0x17, 0x3c, 0xb9, 0xa4, 0x36, 0x99, 0x6c, 0x22, 0xa7, 0xf4, 0x67, 0x8e,
	0xa8, 0xb4, 0x21, 0x15, 0x6c, 0xde, 0x91, 0x93, 0x21, 0xfa, 0xe4, 0xc5, 0xbc, 0xf0, 0x52, 0x1b,
	0x5a, 0x7d, 0x31, 0x80, 0x27, 0x0c, 0xd6, 0x77, 0x7b, 0x74, 0x49, 0x5c, 0xc0, 0x2c, 0x47, 0x5b,
	0xb3, 0x25, 0x5b, 0xcf, 0x9b, 0x1e, 0x0a, 0x0e, 0x15, 0x6d, 0xeb, 0xaa, 0x10, 0x15, 0x6d, 0xc5,
	0x2b, 0x9c, 0x1b, 0xef, 0x92, 0x26, 0x87, 0xf1, 0x7d, 0x17, 0x7d, 0x0e, 0xf5, 0x6c, 0xc9, 0xd6,
	0xa7, 0x77, 0xd7, 0xf2, 0xd7, 0x55, 0xfe, 0x39, 0xca, 0xc7, 0x51, 0xfc, 0xdc, 0x6b, 0x1b, 0x6e,
	0x26, 0xf3, 0xe2, 0x05, 0xf8, 0x54, 0x21, 0x04, 0x1c, 0x39, 0xbd, 0xc7, 0x21, 0x43, 0xc1, 0xe2,
	0x06, 0x78, 0xc4, 0xd6, 0xe7, 0x68, 0x70, 0xb8, 0xf9, 0x13, 0xe8, 0x6c, 0x64, 0xcb, 0xea, 0xea,
	0x12, 0x4e, 0x1a, 0xb4, 0xa8, 0x5b, 0xfc, 0xff, 0xc8, 0xc3, 0xd5, 0xdb, 0xe2, 0x70, 0x3f, 0x9b,
	0xe3, 0xd2, 0xcc, 0xfd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x78, 0x10, 0x32, 0xf7, 0x44, 0x01,
	0x00, 0x00,
}