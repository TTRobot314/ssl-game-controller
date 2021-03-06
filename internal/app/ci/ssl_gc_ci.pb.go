// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ssl_gc_ci.proto

package ci

import (
	fmt "fmt"
	api "github.com/RoboCup-SSL/ssl-game-controller/internal/app/api"
	state "github.com/RoboCup-SSL/ssl-game-controller/internal/app/state"
	tracker "github.com/RoboCup-SSL/ssl-game-controller/internal/app/tracker"
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

// The input format to the GC
type CiInput struct {
	// New unix timestamp in [ns] for the GC
	Timestamp *int64 `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	// New tracker packet with ball and robot data
	TrackerPacket *tracker.TrackerWrapperPacket `protobuf:"bytes,2,opt,name=tracker_packet,json=trackerPacket" json:"tracker_packet,omitempty"`
	// (UI) API input
	ApiInputs            []*api.Input `protobuf:"bytes,3,rep,name=api_inputs,json=apiInputs" json:"api_inputs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CiInput) Reset()         { *m = CiInput{} }
func (m *CiInput) String() string { return proto.CompactTextString(m) }
func (*CiInput) ProtoMessage()    {}
func (*CiInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb3c537990502ac5, []int{0}
}

func (m *CiInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CiInput.Unmarshal(m, b)
}
func (m *CiInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CiInput.Marshal(b, m, deterministic)
}
func (m *CiInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CiInput.Merge(m, src)
}
func (m *CiInput) XXX_Size() int {
	return xxx_messageInfo_CiInput.Size(m)
}
func (m *CiInput) XXX_DiscardUnknown() {
	xxx_messageInfo_CiInput.DiscardUnknown(m)
}

var xxx_messageInfo_CiInput proto.InternalMessageInfo

func (m *CiInput) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *CiInput) GetTrackerPacket() *tracker.TrackerWrapperPacket {
	if m != nil {
		return m.TrackerPacket
	}
	return nil
}

func (m *CiInput) GetApiInputs() []*api.Input {
	if m != nil {
		return m.ApiInputs
	}
	return nil
}

// The output format of the GC response
type CiOutput struct {
	// Latest referee message
	RefereeMsg           *state.Referee `protobuf:"bytes,1,opt,name=referee_msg,json=refereeMsg" json:"referee_msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CiOutput) Reset()         { *m = CiOutput{} }
func (m *CiOutput) String() string { return proto.CompactTextString(m) }
func (*CiOutput) ProtoMessage()    {}
func (*CiOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb3c537990502ac5, []int{1}
}

func (m *CiOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CiOutput.Unmarshal(m, b)
}
func (m *CiOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CiOutput.Marshal(b, m, deterministic)
}
func (m *CiOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CiOutput.Merge(m, src)
}
func (m *CiOutput) XXX_Size() int {
	return xxx_messageInfo_CiOutput.Size(m)
}
func (m *CiOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_CiOutput.DiscardUnknown(m)
}

var xxx_messageInfo_CiOutput proto.InternalMessageInfo

func (m *CiOutput) GetRefereeMsg() *state.Referee {
	if m != nil {
		return m.RefereeMsg
	}
	return nil
}

func init() {
	proto.RegisterType((*CiInput)(nil), "CiInput")
	proto.RegisterType((*CiOutput)(nil), "CiOutput")
}

func init() {
	proto.RegisterFile("ssl_gc_ci.proto", fileDescriptor_fb3c537990502ac5)
}

var fileDescriptor_fb3c537990502ac5 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8e, 0x4d, 0x4b, 0xf4, 0x30,
	0x14, 0x85, 0xe9, 0x3b, 0xf0, 0x3a, 0x93, 0xe2, 0x07, 0x05, 0xa1, 0xc8, 0x2c, 0xca, 0x80, 0x50,
	0x17, 0x6d, 0xa1, 0xe0, 0x46, 0x66, 0x65, 0x57, 0x82, 0xa2, 0x64, 0x04, 0xc1, 0x4d, 0xc8, 0xc4,
	0x6b, 0x0c, 0xb6, 0xcd, 0x25, 0x37, 0xd5, 0x3f, 0xe1, 0x8f, 0x96, 0x36, 0x55, 0x57, 0x21, 0xcf,
	0xb9, 0x87, 0xf3, 0xb0, 0x63, 0xa2, 0x56, 0x68, 0x25, 0x94, 0x29, 0xd1, 0x59, 0x6f, 0xcf, 0xb2,
	0x11, 0x7c, 0x18, 0x32, 0xb6, 0x17, 0x9f, 0x4e, 0x22, 0x82, 0x13, 0xde, 0x49, 0xf5, 0x0e, 0x2f,
	0xf3, 0xc5, 0xc9, 0x5c, 0x91, 0xf8, 0xd3, 0x59, 0xcf, 0xc4, 0xc1, 0x2b, 0x38, 0x00, 0xd1, 0x01,
	0x91, 0xd4, 0x10, 0xd2, 0xcd, 0x57, 0xc4, 0x0e, 0x1a, 0x73, 0xd3, 0xe3, 0xe0, 0x93, 0x35, 0x5b,
	0x79, 0xd3, 0x01, 0x79, 0xd9, 0x61, 0x1a, 0x65, 0x51, 0xbe, 0xe0, 0x7f, 0x20, 0xd9, 0xb2, 0xa3,
	0x30, 0xe5, 0x04, 0x8e, 0x8f, 0x4f, 0xff, 0x65, 0x51, 0x1e, 0xd7, 0xa7, 0xe5, 0x63, 0xc0, 0x4f,
	0x41, 0xe8, 0x61, 0x0a, 0xf9, 0xe1, 0x7c, 0x1c, 0xbe, 0xc9, 0x39, 0x63, 0x12, 0x8d, 0x30, 0xe3,
	0x10, 0xa5, 0x8b, 0x6c, 0x91, 0xc7, 0xf5, 0xff, 0x72, 0xda, 0xe5, 0x2b, 0x89, 0xc1, 0x80, 0x36,
	0x97, 0x6c, 0xd9, 0x98, 0xfb, 0xc1, 0x8f, 0x3a, 0x17, 0x2c, 0xfe, 0x75, 0x26, 0x3d, 0x09, 0xc5,
	0xf5, 0xb2, 0xe4, 0x81, 0x71, 0x36, 0x87, 0x77, 0xa4, 0xaf, 0xb7, 0xcf, 0x57, 0xda, 0xf8, 0xb7,
	0x61, 0x5f, 0x2a, 0xdb, 0x55, 0xdc, 0xee, 0x6d, 0x33, 0x60, 0xb1, 0xdb, 0xdd, 0x56, 0x44, 0x6d,
	0xa1, 0x65, 0x07, 0x85, 0xb2, 0xbd, 0x77, 0xb6, 0x6d, 0xc1, 0x55, 0xa6, 0xf7, 0xe0, 0x7a, 0xd9,
	0x56, 0x12, 0xb1, 0x52, 0xe6, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x3c, 0xb7, 0xc6, 0x21, 0x67, 0x01,
	0x00, 0x00,
}
