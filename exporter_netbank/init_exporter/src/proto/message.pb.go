// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package proto

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

// RawMessage represents a message emitted by the Mainflux adapters layer.
type RawMessage struct {
	Channel              string   `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	Subtopic             string   `protobuf:"bytes,2,opt,name=subtopic,proto3" json:"subtopic,omitempty"`
	Publisher            string   `protobuf:"bytes,3,opt,name=publisher,proto3" json:"publisher,omitempty"`
	Protocol             string   `protobuf:"bytes,4,opt,name=protocol,proto3" json:"protocol,omitempty"`
	ContentType          string   `protobuf:"bytes,5,opt,name=contentType,proto3" json:"contentType,omitempty"`
	Payload              []byte   `protobuf:"bytes,6,opt,name=payload,proto3" json:"payload,omitempty"`
	SrcIp                string   `protobuf:"bytes,7,opt,name=srcIp,proto3" json:"srcIp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RawMessage) Reset()         { *m = RawMessage{} }
func (m *RawMessage) String() string { return proto.CompactTextString(m) }
func (*RawMessage) ProtoMessage()    {}
func (*RawMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *RawMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RawMessage.Unmarshal(m, b)
}
func (m *RawMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RawMessage.Marshal(b, m, deterministic)
}
func (m *RawMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RawMessage.Merge(m, src)
}
func (m *RawMessage) XXX_Size() int {
	return xxx_messageInfo_RawMessage.Size(m)
}
func (m *RawMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_RawMessage.DiscardUnknown(m)
}

var xxx_messageInfo_RawMessage proto.InternalMessageInfo

func (m *RawMessage) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *RawMessage) GetSubtopic() string {
	if m != nil {
		return m.Subtopic
	}
	return ""
}

func (m *RawMessage) GetPublisher() string {
	if m != nil {
		return m.Publisher
	}
	return ""
}

func (m *RawMessage) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *RawMessage) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *RawMessage) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *RawMessage) GetSrcIp() string {
	if m != nil {
		return m.SrcIp
	}
	return ""
}

// Message represents a resolved (normalized) raw message.
type Message struct {
	Channel   string `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	Subtopic  string `protobuf:"bytes,2,opt,name=subtopic,proto3" json:"subtopic,omitempty"`
	Publisher string `protobuf:"bytes,3,opt,name=publisher,proto3" json:"publisher,omitempty"`
	Protocol  string `protobuf:"bytes,4,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Name      string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Unit      string `protobuf:"bytes,6,opt,name=unit,proto3" json:"unit,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*Message_FloatValue
	//	*Message_StringValue
	//	*Message_BoolValue
	//	*Message_DataValue
	Value                isMessage_Value `protobuf_oneof:"value"`
	ValueSum             *SumValue       `protobuf:"bytes,11,opt,name=valueSum,proto3" json:"valueSum,omitempty"`
	Time                 float64         `protobuf:"fixed64,12,opt,name=time,proto3" json:"time,omitempty"`
	UpdateTime           float64         `protobuf:"fixed64,13,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	Link                 string          `protobuf:"bytes,14,opt,name=link,proto3" json:"link,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *Message) GetSubtopic() string {
	if m != nil {
		return m.Subtopic
	}
	return ""
}

func (m *Message) GetPublisher() string {
	if m != nil {
		return m.Publisher
	}
	return ""
}

func (m *Message) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *Message) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Message) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

type isMessage_Value interface {
	isMessage_Value()
}

type Message_FloatValue struct {
	FloatValue float64 `protobuf:"fixed64,7,opt,name=floatValue,proto3,oneof"`
}

type Message_StringValue struct {
	StringValue string `protobuf:"bytes,8,opt,name=stringValue,proto3,oneof"`
}

type Message_BoolValue struct {
	BoolValue bool `protobuf:"varint,9,opt,name=boolValue,proto3,oneof"`
}

type Message_DataValue struct {
	DataValue string `protobuf:"bytes,10,opt,name=dataValue,proto3,oneof"`
}

func (*Message_FloatValue) isMessage_Value() {}

func (*Message_StringValue) isMessage_Value() {}

func (*Message_BoolValue) isMessage_Value() {}

func (*Message_DataValue) isMessage_Value() {}

func (m *Message) GetValue() isMessage_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Message) GetFloatValue() float64 {
	if x, ok := m.GetValue().(*Message_FloatValue); ok {
		return x.FloatValue
	}
	return 0
}

func (m *Message) GetStringValue() string {
	if x, ok := m.GetValue().(*Message_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *Message) GetBoolValue() bool {
	if x, ok := m.GetValue().(*Message_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (m *Message) GetDataValue() string {
	if x, ok := m.GetValue().(*Message_DataValue); ok {
		return x.DataValue
	}
	return ""
}

func (m *Message) GetValueSum() *SumValue {
	if m != nil {
		return m.ValueSum
	}
	return nil
}

func (m *Message) GetTime() float64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Message) GetUpdateTime() float64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

func (m *Message) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Message) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Message_FloatValue)(nil),
		(*Message_StringValue)(nil),
		(*Message_BoolValue)(nil),
		(*Message_DataValue)(nil),
	}
}

// SumValue is a simple wrapper around the double value.
type SumValue struct {
	Value                float64  `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumValue) Reset()         { *m = SumValue{} }
func (m *SumValue) String() string { return proto.CompactTextString(m) }
func (*SumValue) ProtoMessage()    {}
func (*SumValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{2}
}

func (m *SumValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumValue.Unmarshal(m, b)
}
func (m *SumValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumValue.Marshal(b, m, deterministic)
}
func (m *SumValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumValue.Merge(m, src)
}
func (m *SumValue) XXX_Size() int {
	return xxx_messageInfo_SumValue.Size(m)
}
func (m *SumValue) XXX_DiscardUnknown() {
	xxx_messageInfo_SumValue.DiscardUnknown(m)
}

var xxx_messageInfo_SumValue proto.InternalMessageInfo

func (m *SumValue) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*RawMessage)(nil), "mainflux.RawMessage")
	proto.RegisterType((*Message)(nil), "mainflux.Message")
	proto.RegisterType((*SumValue)(nil), "mainflux.SumValue")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x92, 0x3d, 0x6e, 0xe3, 0x30,
	0x10, 0x85, 0xcd, 0xf5, 0x8f, 0xa4, 0x91, 0xbd, 0x05, 0xb1, 0x05, 0xb1, 0x58, 0x18, 0x82, 0x2a,
	0x55, 0x2a, 0x76, 0x6f, 0xb0, 0x95, 0x53, 0xa4, 0xa1, 0x8d, 0xf4, 0x94, 0x4c, 0xdb, 0x42, 0x28,
	0x52, 0x90, 0xc8, 0x24, 0xbe, 0x51, 0x4e, 0x93, 0x33, 0x05, 0x1c, 0x46, 0xb6, 0x6e, 0x90, 0x6e,
	0xde, 0xf7, 0xe6, 0x91, 0x33, 0xa2, 0x60, 0xd3, 0xca, 0x61, 0x10, 0x67, 0x59, 0x76, 0xbd, 0xb1,
	0x86, 0xc6, 0xad, 0x68, 0xf4, 0x49, 0xb9, 0xb7, 0xfc, 0x83, 0x00, 0x70, 0xf1, 0xfa, 0x18, 0x6c,
	0xca, 0x20, 0xaa, 0x2f, 0x42, 0x6b, 0xa9, 0x18, 0xc9, 0x48, 0x91, 0xf0, 0x51, 0xd2, 0xdf, 0x10,
	0x0f, 0xae, 0xb2, 0xa6, 0x6b, 0x6a, 0xf6, 0x03, 0xad, 0x9b, 0xa6, 0x7f, 0x20, 0xe9, 0x5c, 0xa5,
	0x9a, 0xe1, 0x22, 0x7b, 0x36, 0x47, 0xf3, 0x0e, 0x7c, 0x12, 0x6f, 0xad, 0x8d, 0x62, 0x8b, 0x90,
	0x1c, 0x35, 0xcd, 0x20, 0xad, 0x8d, 0xb6, 0x52, 0xdb, 0xc3, 0xb5, 0x93, 0x6c, 0x89, 0xf6, 0x14,
	0xf9, 0x89, 0x3a, 0x71, 0x55, 0x46, 0x1c, 0xd9, 0x2a, 0x23, 0xc5, 0x9a, 0x8f, 0x92, 0xfe, 0x82,
	0xe5, 0xd0, 0xd7, 0x0f, 0x1d, 0x8b, 0x30, 0x15, 0x44, 0xfe, 0x3e, 0x87, 0xe8, 0xbb, 0xb6, 0xa1,
	0xb0, 0xd0, 0xa2, 0x1d, 0xd7, 0xc0, 0xda, 0x33, 0xa7, 0x1b, 0x8b, 0xc3, 0x27, 0x1c, 0x6b, 0x9a,
	0x01, 0x9c, 0x94, 0x11, 0xf6, 0x49, 0x28, 0x27, 0x71, 0x7c, 0xb2, 0x9b, 0xf1, 0x09, 0xa3, 0x39,
	0xa4, 0x83, 0xed, 0x1b, 0x7d, 0x0e, 0x2d, 0xb1, 0x0f, 0xef, 0x66, 0x7c, 0x0a, 0xe9, 0x16, 0x92,
	0xca, 0x18, 0x15, 0x3a, 0x92, 0x8c, 0x14, 0xf1, 0x6e, 0xc6, 0xef, 0xc8, 0xfb, 0x47, 0x61, 0x45,
	0xf0, 0xe1, 0xeb, 0x84, 0x3b, 0xa2, 0x25, 0xc4, 0x2f, 0xbe, 0xd8, 0xbb, 0x96, 0xa5, 0x19, 0x29,
	0xd2, 0xbf, 0xb4, 0x1c, 0xff, 0x8b, 0x72, 0xef, 0x5a, 0xec, 0xe2, 0xb7, 0x1e, 0xbf, 0x89, 0x6d,
	0x5a, 0xc9, 0xd6, 0x7e, 0x5e, 0x8e, 0x35, 0xdd, 0x02, 0xb8, 0xee, 0x28, 0xac, 0x3c, 0x78, 0x67,
	0x83, 0xce, 0x84, 0xf8, 0x8c, 0x6a, 0xf4, 0x33, 0xfb, 0x19, 0xb6, 0xf7, 0xf5, 0xff, 0x08, 0x96,
	0x78, 0x66, 0x9e, 0x41, 0x3c, 0x5e, 0xe3, 0x1f, 0x13, 0x21, 0x3e, 0x14, 0xe1, 0x41, 0x54, 0x2b,
	0xfc, 0xb4, 0xff, 0x3e, 0x03, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x4e, 0x0e, 0xa9, 0xbf, 0x02, 0x00,
	0x00,
}
