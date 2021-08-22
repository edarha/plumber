// Code generated by protoc-gen-go. DO NOT EDIT.
// source: azure-service-bus.proto

package backends

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

type AzureServiceBus struct {
	// Required for reads and writes
	Queue string `protobuf:"bytes,1,opt,name=queue,proto3" json:"queue,omitempty"`
	// Required for reads
	// Ignored for writes
	Topic string `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	// Required for reads
	// Ignored for writes
	SubscriptionName     string   `protobuf:"bytes,3,opt,name=subscription_name,json=subscriptionName,proto3" json:"subscription_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AzureServiceBus) Reset()         { *m = AzureServiceBus{} }
func (m *AzureServiceBus) String() string { return proto.CompactTextString(m) }
func (*AzureServiceBus) ProtoMessage()    {}
func (*AzureServiceBus) Descriptor() ([]byte, []int) {
	return fileDescriptor_454428be5f04e702, []int{0}
}

func (m *AzureServiceBus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AzureServiceBus.Unmarshal(m, b)
}
func (m *AzureServiceBus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AzureServiceBus.Marshal(b, m, deterministic)
}
func (m *AzureServiceBus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AzureServiceBus.Merge(m, src)
}
func (m *AzureServiceBus) XXX_Size() int {
	return xxx_messageInfo_AzureServiceBus.Size(m)
}
func (m *AzureServiceBus) XXX_DiscardUnknown() {
	xxx_messageInfo_AzureServiceBus.DiscardUnknown(m)
}

var xxx_messageInfo_AzureServiceBus proto.InternalMessageInfo

func (m *AzureServiceBus) GetQueue() string {
	if m != nil {
		return m.Queue
	}
	return ""
}

func (m *AzureServiceBus) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *AzureServiceBus) GetSubscriptionName() string {
	if m != nil {
		return m.SubscriptionName
	}
	return ""
}

func init() {
	proto.RegisterType((*AzureServiceBus)(nil), "protos.backends.AzureServiceBus")
}

func init() { proto.RegisterFile("azure-service-bus.proto", fileDescriptor_454428be5f04e702) }

var fileDescriptor_454428be5f04e702 = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0x31, 0x4b, 0xc4, 0x40,
	0x10, 0x46, 0x89, 0xa2, 0xe0, 0x36, 0xd1, 0x20, 0x98, 0x52, 0xac, 0x04, 0x49, 0xb6, 0xb0, 0x16,
	0x31, 0x3f, 0xc0, 0x42, 0x3b, 0x1b, 0xd9, 0x99, 0x0c, 0xc9, 0x6a, 0x36, 0xbb, 0xb7, 0xb3, 0x73,
	0xc5, 0xfd, 0xfa, 0x23, 0xbb, 0x1c, 0x5c, 0x35, 0x7c, 0x6f, 0x5e, 0xf1, 0xd4, 0x83, 0x39, 0x48,
	0xa4, 0x8e, 0x29, 0xee, 0x2d, 0x52, 0x07, 0xc2, 0x7d, 0x88, 0x3e, 0xf9, 0xa6, 0xce, 0x87, 0x7b,
	0x30, 0xf8, 0x4f, 0xeb, 0xc8, 0x4f, 0x7f, 0xaa, 0xfe, 0xd8, 0xdc, 0xef, 0xa2, 0x0e, 0xc2, 0xcd,
	0xbd, 0xba, 0xda, 0x09, 0x09, 0xb5, 0xd5, 0x63, 0xf5, 0x7c, 0xf3, 0x55, 0xc6, 0x46, 0x93, 0x0f,
	0x16, 0xdb, 0x8b, 0x42, 0xf3, 0x68, 0x5e, 0xd4, 0x1d, 0x0b, 0x30, 0x46, 0x1b, 0x92, 0xf5, 0xeb,
	0xef, 0x6a, 0x1c, 0xb5, 0x97, 0xd9, 0xb8, 0x3d, 0x7f, 0x7c, 0x1a, 0x47, 0xc3, 0xfb, 0xcf, 0xdb,
	0x64, 0xd3, 0x2c, 0xd0, 0xa3, 0x77, 0x1a, 0x4c, 0xc2, 0x19, 0x7d, 0x0c, 0x3a, 0x2c, 0xe2, 0x80,
	0x62, 0xc7, 0x38, 0x93, 0x33, 0xac, 0x41, 0xec, 0x32, 0xea, 0xc9, 0xeb, 0x12, 0xab, 0x4f, 0xb1,
	0x70, 0x9d, 0xc1, 0xeb, 0x31, 0x00, 0x00, 0xff, 0xff, 0xb6, 0xee, 0xf0, 0xd8, 0xdf, 0x00, 0x00,
	0x00,
}
