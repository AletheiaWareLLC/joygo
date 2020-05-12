// Code generated by protoc-gen-go. DO NOT EDIT.
// source: joy.proto

package joygo

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

type Mesh struct {
	Name                 string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Vertices             uint32    `protobuf:"varint,2,opt,name=vertices,proto3" json:"vertices,omitempty"`
	Vertex               []float64 `protobuf:"fixed64,3,rep,packed,name=vertex,proto3" json:"vertex,omitempty"`
	Normal               []float64 `protobuf:"fixed64,4,rep,packed,name=normal,proto3" json:"normal,omitempty"`
	TexCoord             []float64 `protobuf:"fixed64,5,rep,packed,name=tex_coord,json=texCoord,proto3" json:"tex_coord,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Mesh) Reset()         { *m = Mesh{} }
func (m *Mesh) String() string { return proto.CompactTextString(m) }
func (*Mesh) ProtoMessage()    {}
func (*Mesh) Descriptor() ([]byte, []int) {
	return fileDescriptor_be7dde79f0a37eab, []int{0}
}

func (m *Mesh) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mesh.Unmarshal(m, b)
}
func (m *Mesh) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mesh.Marshal(b, m, deterministic)
}
func (m *Mesh) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mesh.Merge(m, src)
}
func (m *Mesh) XXX_Size() int {
	return xxx_messageInfo_Mesh.Size(m)
}
func (m *Mesh) XXX_DiscardUnknown() {
	xxx_messageInfo_Mesh.DiscardUnknown(m)
}

var xxx_messageInfo_Mesh proto.InternalMessageInfo

func (m *Mesh) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Mesh) GetVertices() uint32 {
	if m != nil {
		return m.Vertices
	}
	return 0
}

func (m *Mesh) GetVertex() []float64 {
	if m != nil {
		return m.Vertex
	}
	return nil
}

func (m *Mesh) GetNormal() []float64 {
	if m != nil {
		return m.Normal
	}
	return nil
}

func (m *Mesh) GetTexCoord() []float64 {
	if m != nil {
		return m.TexCoord
	}
	return nil
}

type Shader struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	VertexSource         string   `protobuf:"bytes,2,opt,name=vertex_source,json=vertexSource,proto3" json:"vertex_source,omitempty"`
	FragmentSource       string   `protobuf:"bytes,3,opt,name=fragment_source,json=fragmentSource,proto3" json:"fragment_source,omitempty"`
	Attributes           []string `protobuf:"bytes,4,rep,name=attributes,proto3" json:"attributes,omitempty"`
	Uniforms             []string `protobuf:"bytes,5,rep,name=uniforms,proto3" json:"uniforms,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Shader) Reset()         { *m = Shader{} }
func (m *Shader) String() string { return proto.CompactTextString(m) }
func (*Shader) ProtoMessage()    {}
func (*Shader) Descriptor() ([]byte, []int) {
	return fileDescriptor_be7dde79f0a37eab, []int{1}
}

func (m *Shader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Shader.Unmarshal(m, b)
}
func (m *Shader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Shader.Marshal(b, m, deterministic)
}
func (m *Shader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Shader.Merge(m, src)
}
func (m *Shader) XXX_Size() int {
	return xxx_messageInfo_Shader.Size(m)
}
func (m *Shader) XXX_DiscardUnknown() {
	xxx_messageInfo_Shader.DiscardUnknown(m)
}

var xxx_messageInfo_Shader proto.InternalMessageInfo

func (m *Shader) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Shader) GetVertexSource() string {
	if m != nil {
		return m.VertexSource
	}
	return ""
}

func (m *Shader) GetFragmentSource() string {
	if m != nil {
		return m.FragmentSource
	}
	return ""
}

func (m *Shader) GetAttributes() []string {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *Shader) GetUniforms() []string {
	if m != nil {
		return m.Uniforms
	}
	return nil
}

func init() {
	proto.RegisterType((*Mesh)(nil), "joy.Mesh")
	proto.RegisterType((*Shader)(nil), "joy.Shader")
}

func init() { proto.RegisterFile("joy.proto", fileDescriptor_be7dde79f0a37eab) }

var fileDescriptor_be7dde79f0a37eab = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0xe5, 0xa6, 0x54, 0xf5, 0x89, 0x02, 0xb2, 0x18, 0xa2, 0x0e, 0x10, 0x95, 0x81, 0x4e,
	0xe9, 0xc0, 0x13, 0x90, 0x6e, 0xa8, 0x48, 0xc8, 0x1d, 0x90, 0x58, 0x2a, 0x37, 0xbd, 0x36, 0x89,
	0xea, 0x1c, 0x72, 0x1c, 0x48, 0x9e, 0x82, 0xb7, 0xe0, 0x39, 0x51, 0x9c, 0x34, 0x30, 0xb0, 0xdd,
	0x7d, 0xff, 0x2f, 0xf9, 0xd3, 0x19, 0x78, 0x46, 0x75, 0xf8, 0x6e, 0xc8, 0x92, 0xf0, 0x32, 0xaa,
	0x67, 0x5f, 0x0c, 0x86, 0xcf, 0x58, 0x24, 0x42, 0xc0, 0x30, 0x57, 0x1a, 0x7d, 0x16, 0xb0, 0x39,
	0x97, 0x6e, 0x16, 0x53, 0x18, 0x7f, 0xa0, 0xb1, 0x69, 0x8c, 0x85, 0x3f, 0x08, 0xd8, 0x7c, 0x22,
	0xfb, 0x5d, 0x4c, 0x61, 0xd4, 0xcc, 0x58, 0xf9, 0x5e, 0xe0, 0xcd, 0x59, 0x34, 0xb8, 0x62, 0xb2,
	0x23, 0x4d, 0x96, 0x93, 0xd1, 0xea, 0xe8, 0x0f, 0x7f, 0xb3, 0x96, 0x88, 0x5b, 0xe0, 0x16, 0xab,
	0x4d, 0x4c, 0x64, 0x76, 0xfe, 0x59, 0x1f, 0x8f, 0x2d, 0x56, 0xcb, 0x86, 0xcd, 0xbe, 0x19, 0x8c,
	0xd6, 0x89, 0xda, 0xa1, 0xf9, 0xd7, 0xe9, 0x0e, 0x26, 0xed, 0x2b, 0x9b, 0x82, 0x4a, 0x13, 0xa3,
	0x13, 0xe3, 0xf2, 0xbc, 0x85, 0x6b, 0xc7, 0xc4, 0x3d, 0x5c, 0xee, 0x8d, 0x3a, 0x68, 0xcc, 0xed,
	0xa9, 0xe6, 0xb9, 0xda, 0xc5, 0x09, 0x77, 0xc5, 0x1b, 0x00, 0x65, 0xad, 0x49, 0xb7, 0xa5, 0xc5,
	0xc2, 0xd9, 0x72, 0xf9, 0x87, 0x34, 0x17, 0x28, 0xf3, 0x74, 0x4f, 0x46, 0x17, 0x4e, 0x96, 0xcb,
	0x7e, 0x8f, 0x22, 0xb8, 0x8e, 0x49, 0x87, 0xea, 0x88, 0x36, 0xc1, 0x54, 0x7d, 0x2a, 0x83, 0x61,
	0x46, 0x75, 0x34, 0x7e, 0xa2, 0xfa, 0xa5, 0xb9, 0xf0, 0x5b, 0x70, 0x48, 0x6d, 0x52, 0x6e, 0xc3,
	0x98, 0xf4, 0xe2, 0xb1, 0xab, 0xbd, 0x2a, 0x83, 0xab, 0xd5, 0x72, 0x91, 0x51, 0x7d, 0xa0, 0xed,
	0xc8, 0x7d, 0xc5, 0xc3, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd2, 0xf4, 0xc3, 0x8f, 0x97, 0x01,
	0x00, 0x00,
}
