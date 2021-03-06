// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/example/example.proto

package go_micro_srv_GetArea

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

type Message struct {
	Say                  string   `protobuf:"bytes,1,opt,name=say,proto3" json:"say,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{0}
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

func (m *Message) GetSay() string {
	if m != nil {
		return m.Say
	}
	return ""
}

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type Response struct {
	// 错误码
	Error string `protobuf:"bytes,1,opt,name=Error,proto3" json:"Error,omitempty"`
	// 错误信息
	Errmsg string `protobuf:"bytes,2,opt,name=Errmsg,proto3" json:"Errmsg,omitempty"`
	// 返回的地区切片
	Data                 []*Response_Area `protobuf:"bytes,3,rep,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *Response) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func (m *Response) GetData() []*Response_Area {
	if m != nil {
		return m.Data
	}
	return nil
}

// 地区
type Response_Area struct {
	Aid                  int32    `protobuf:"varint,1,opt,name=Aid,proto3" json:"Aid,omitempty"`
	Aname                string   `protobuf:"bytes,2,opt,name=Aname,proto3" json:"Aname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response_Area) Reset()         { *m = Response_Area{} }
func (m *Response_Area) String() string { return proto.CompactTextString(m) }
func (*Response_Area) ProtoMessage()    {}
func (*Response_Area) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{2, 0}
}

func (m *Response_Area) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response_Area.Unmarshal(m, b)
}
func (m *Response_Area) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response_Area.Marshal(b, m, deterministic)
}
func (m *Response_Area) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response_Area.Merge(m, src)
}
func (m *Response_Area) XXX_Size() int {
	return xxx_messageInfo_Response_Area.Size(m)
}
func (m *Response_Area) XXX_DiscardUnknown() {
	xxx_messageInfo_Response_Area.DiscardUnknown(m)
}

var xxx_messageInfo_Response_Area proto.InternalMessageInfo

func (m *Response_Area) GetAid() int32 {
	if m != nil {
		return m.Aid
	}
	return 0
}

func (m *Response_Area) GetAname() string {
	if m != nil {
		return m.Aname
	}
	return ""
}

type StreamingRequest struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingRequest) Reset()         { *m = StreamingRequest{} }
func (m *StreamingRequest) String() string { return proto.CompactTextString(m) }
func (*StreamingRequest) ProtoMessage()    {}
func (*StreamingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{3}
}

func (m *StreamingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingRequest.Unmarshal(m, b)
}
func (m *StreamingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingRequest.Marshal(b, m, deterministic)
}
func (m *StreamingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingRequest.Merge(m, src)
}
func (m *StreamingRequest) XXX_Size() int {
	return xxx_messageInfo_StreamingRequest.Size(m)
}
func (m *StreamingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingRequest proto.InternalMessageInfo

func (m *StreamingRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type StreamingResponse struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingResponse) Reset()         { *m = StreamingResponse{} }
func (m *StreamingResponse) String() string { return proto.CompactTextString(m) }
func (*StreamingResponse) ProtoMessage()    {}
func (*StreamingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{4}
}

func (m *StreamingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingResponse.Unmarshal(m, b)
}
func (m *StreamingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingResponse.Marshal(b, m, deterministic)
}
func (m *StreamingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingResponse.Merge(m, src)
}
func (m *StreamingResponse) XXX_Size() int {
	return xxx_messageInfo_StreamingResponse.Size(m)
}
func (m *StreamingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingResponse proto.InternalMessageInfo

func (m *StreamingResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Ping struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{5}
}

func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (m *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(m, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

func (m *Ping) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

type Pong struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{6}
}

func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (m *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(m, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

func (m *Pong) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

func init() {
	proto.RegisterType((*Message)(nil), "go.micro.srv.GetArea.Message")
	proto.RegisterType((*Request)(nil), "go.micro.srv.GetArea.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.GetArea.Response")
	proto.RegisterType((*Response_Area)(nil), "go.micro.srv.GetArea.Response.Area")
	proto.RegisterType((*StreamingRequest)(nil), "go.micro.srv.GetArea.StreamingRequest")
	proto.RegisterType((*StreamingResponse)(nil), "go.micro.srv.GetArea.StreamingResponse")
	proto.RegisterType((*Ping)(nil), "go.micro.srv.GetArea.Ping")
	proto.RegisterType((*Pong)(nil), "go.micro.srv.GetArea.Pong")
}

func init() { proto.RegisterFile("proto/example/example.proto", fileDescriptor_097b3f5db5cf5789) }

var fileDescriptor_097b3f5db5cf5789 = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0x5e, 0xec, 0xb6, 0x6e, 0xcf, 0xcb, 0x0c, 0x43, 0x46, 0x87, 0x63, 0x44, 0xd0, 0x7a, 0x89,
	0x63, 0x1e, 0x3c, 0x0f, 0x1c, 0x8a, 0x20, 0x48, 0x3d, 0x7b, 0x88, 0xf3, 0x51, 0x8a, 0x36, 0x99,
	0x49, 0x26, 0xfa, 0x73, 0xfc, 0xa3, 0x22, 0x49, 0x33, 0x14, 0x69, 0xf5, 0xd4, 0x7c, 0xf9, 0xbe,
	0xf7, 0xde, 0xf7, 0xbd, 0x06, 0xc6, 0x6b, 0xad, 0xac, 0x3a, 0xc5, 0x37, 0x51, 0xae, 0x9f, 0x71,
	0xfb, 0xe5, 0xfe, 0x96, 0x0e, 0x73, 0xc5, 0xcb, 0x62, 0xa5, 0x15, 0x37, 0xfa, 0x95, 0x5f, 0xa2,
	0x5d, 0x68, 0x14, 0x6c, 0x0c, 0xf1, 0x0d, 0x1a, 0x23, 0x72, 0xa4, 0x03, 0x88, 0x8c, 0x78, 0x1f,
	0x91, 0x29, 0x49, 0xfb, 0x99, 0x3b, 0xb2, 0x3e, 0xc4, 0x19, 0xbe, 0x6c, 0xd0, 0x58, 0xf6, 0x41,
	0xa0, 0x97, 0xa1, 0x59, 0x2b, 0x69, 0x90, 0x0e, 0xa1, 0xb3, 0xd4, 0x5a, 0xe9, 0xa0, 0xad, 0x00,
	0xdd, 0x87, 0xee, 0x52, 0xeb, 0xd2, 0xe4, 0xa3, 0x1d, 0x7f, 0x1d, 0x10, 0x3d, 0x87, 0xf6, 0x85,
	0xb0, 0x62, 0x14, 0x4d, 0xa3, 0x74, 0x77, 0x7e, 0xc8, 0xeb, 0x7c, 0xf0, 0x6d, 0x6f, 0xee, 0x50,
	0xe6, 0x0b, 0x12, 0x0e, 0x6d, 0x87, 0x9c, 0xb1, 0x45, 0xf1, 0xe8, 0x87, 0x75, 0x32, 0x77, 0x74,
	0x06, 0x16, 0x52, 0x94, 0x18, 0x26, 0x55, 0x80, 0xa5, 0x30, 0xb8, 0xb3, 0x1a, 0x45, 0x59, 0xc8,
	0x3c, 0xf8, 0x76, 0xca, 0x95, 0xda, 0x48, 0xeb, 0xab, 0xa3, 0xac, 0x02, 0xec, 0x04, 0xf6, 0x7e,
	0x28, 0xbf, 0x53, 0xd5, 0x48, 0x27, 0xd0, 0xbe, 0x2d, 0x64, 0xee, 0xd2, 0x19, 0xab, 0xd5, 0x13,
	0x06, 0x3a, 0x20, 0xcf, 0xab, 0x66, 0x7e, 0xfe, 0x49, 0x20, 0x5e, 0x56, 0x3f, 0x82, 0x5e, 0x43,
	0x1c, 0xf2, 0xd2, 0x83, 0xa6, 0x35, 0x78, 0xdb, 0xc9, 0xe4, 0xef, 0x2d, 0xb1, 0x16, 0xbd, 0x87,
	0x6e, 0x15, 0x81, 0x1e, 0xd5, 0x6b, 0x7f, 0xaf, 0x22, 0x39, 0xfe, 0x57, 0xb7, 0x6d, 0x3e, 0x23,
	0xf4, 0x0a, 0x7a, 0x2e, 0xb6, 0x8f, 0x96, 0xd4, 0x17, 0x3a, 0x3e, 0x69, 0xe2, 0x94, 0xcc, 0x59,
	0x2b, 0x25, 0x33, 0xf2, 0xd0, 0xf5, 0xcf, 0xef, 0xec, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x94, 0x69,
	0x0c, 0x3e, 0x9d, 0x02, 0x00, 0x00,
}
