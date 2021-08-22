// Code generated by protoc-gen-go. DO NOT EDIT.
// source: postgres.proto

package conns

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

type Postgres struct {
	// Required
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// Required
	Port uint32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	// Optional
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	// Optional
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	// Required
	Database string `protobuf:"bytes,5,opt,name=database,proto3" json:"database,omitempty"`
	// Whether to collect using TLS
	UseTls bool `protobuf:"varint,6,opt,name=use_tls,json=useTls,proto3" json:"use_tls,omitempty"`
	// Specify to not verify server's TLS certificate
	InsecureTls          bool     `protobuf:"varint,7,opt,name=insecure_tls,json=insecureTls,proto3" json:"insecure_tls,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Postgres) Reset()         { *m = Postgres{} }
func (m *Postgres) String() string { return proto.CompactTextString(m) }
func (*Postgres) ProtoMessage()    {}
func (*Postgres) Descriptor() ([]byte, []int) {
	return fileDescriptor_a62b40fab176d0d3, []int{0}
}

func (m *Postgres) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Postgres.Unmarshal(m, b)
}
func (m *Postgres) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Postgres.Marshal(b, m, deterministic)
}
func (m *Postgres) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Postgres.Merge(m, src)
}
func (m *Postgres) XXX_Size() int {
	return xxx_messageInfo_Postgres.Size(m)
}
func (m *Postgres) XXX_DiscardUnknown() {
	xxx_messageInfo_Postgres.DiscardUnknown(m)
}

var xxx_messageInfo_Postgres proto.InternalMessageInfo

func (m *Postgres) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Postgres) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Postgres) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Postgres) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Postgres) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

func (m *Postgres) GetUseTls() bool {
	if m != nil {
		return m.UseTls
	}
	return false
}

func (m *Postgres) GetInsecureTls() bool {
	if m != nil {
		return m.InsecureTls
	}
	return false
}

func init() {
	proto.RegisterType((*Postgres)(nil), "protos.conns.Postgres")
}

func init() { proto.RegisterFile("postgres.proto", fileDescriptor_a62b40fab176d0d3) }

var fileDescriptor_a62b40fab176d0d3 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x90, 0x3d, 0x4e, 0xc4, 0x30,
	0x10, 0x85, 0x65, 0x58, 0x92, 0x60, 0x16, 0x0a, 0x37, 0x58, 0x54, 0x81, 0x2a, 0x0d, 0x71, 0x41,
	0x87, 0xa8, 0x38, 0x01, 0x8a, 0xb6, 0xa2, 0x41, 0xfe, 0x19, 0x25, 0x91, 0x9c, 0xd8, 0x9a, 0xb1,
	0xc5, 0x15, 0x39, 0x16, 0x5a, 0x87, 0x50, 0xcd, 0x7c, 0xef, 0x7b, 0xd5, 0xe3, 0x77, 0x31, 0x50,
	0x1a, 0x11, 0xa8, 0x8f, 0x18, 0x52, 0x10, 0xc7, 0x72, 0xa8, 0xb7, 0x61, 0x5d, 0xe9, 0xe9, 0x87,
	0xf1, 0xe6, 0xe3, 0xaf, 0x20, 0x24, 0xaf, 0xb5, 0x73, 0x08, 0x44, 0x92, 0xb5, 0xac, 0xbb, 0x1e,
	0x76, 0x14, 0x82, 0x1f, 0x62, 0xc0, 0x24, 0x2f, 0x5a, 0xd6, 0xdd, 0x0e, 0xe5, 0x17, 0x0f, 0xbc,
	0xc9, 0x04, 0xb8, 0xea, 0x05, 0xe4, 0x65, 0xa9, 0xff, 0xf3, 0xd9, 0x45, 0x4d, 0xf4, 0x1d, 0xd0,
	0xc9, 0xc3, 0xe6, 0x76, 0x3e, 0x3b, 0xa7, 0x93, 0x36, 0x9a, 0x40, 0x5e, 0x6d, 0x6e, 0x67, 0x71,
	0xcf, 0xeb, 0x4c, 0xf0, 0x95, 0x3c, 0xc9, 0xaa, 0x65, 0x5d, 0x33, 0x54, 0x99, 0xe0, 0xe4, 0x49,
	0x3c, 0xf2, 0xe3, 0xbc, 0x12, 0xd8, 0x8c, 0x9b, 0xad, 0x8b, 0xbd, 0xd9, 0xb3, 0x93, 0xa7, 0xf7,
	0xb7, 0xcf, 0xd7, 0x71, 0x4e, 0x53, 0x36, 0xbd, 0x0d, 0x8b, 0x32, 0x3a, 0xd9, 0xc9, 0x06, 0x8c,
	0x2a, 0xfa, 0xbc, 0x18, 0xc0, 0x67, 0xb2, 0x13, 0x2c, 0x9a, 0x94, 0xc9, 0xb3, 0x77, 0x6a, 0x0c,
	0x6a, 0x1b, 0x42, 0x95, 0x21, 0x4c, 0x55, 0xe8, 0xe5, 0x37, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x14,
	0xa8, 0x6c, 0x2f, 0x01, 0x00, 0x00,
}
