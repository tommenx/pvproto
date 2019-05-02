// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/executor.proto

package executorpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ErrorType int32

const (
	ErrorType_OK                ErrorType = 0
	ErrorType_UNKNOWN           ErrorType = 1
	ErrorType_NO_NODE           ErrorType = 2
	ErrorType_NO_PV             ErrorType = 3
	ErrorType_INTERNAL_ERROR    ErrorType = 4
	ErrorType_INVALID_PARAMETER ErrorType = 5
)

var ErrorType_name = map[int32]string{
	0: "OK",
	1: "UNKNOWN",
	2: "NO_NODE",
	3: "NO_PV",
	4: "INTERNAL_ERROR",
	5: "INVALID_PARAMETER",
}

var ErrorType_value = map[string]int32{
	"OK":                0,
	"UNKNOWN":           1,
	"NO_NODE":           2,
	"NO_PV":             3,
	"INTERNAL_ERROR":    4,
	"INVALID_PARAMETER": 5,
}

func (x ErrorType) String() string {
	return proto.EnumName(ErrorType_name, int32(x))
}

func (ErrorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9866e87da4c4a057, []int{0}
}

type StorageType int32

const (
	StorageType_STORAGE StorageType = 0
	StorageType_LIMIT   StorageType = 1
)

var StorageType_name = map[int32]string{
	0: "STORAGE",
	1: "LIMIT",
}

var StorageType_value = map[string]int32{
	"STORAGE": 0,
	"LIMIT":   1,
}

func (x StorageType) String() string {
	return proto.EnumName(StorageType_name, int32(x))
}

func (StorageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9866e87da4c4a057, []int{1}
}

type Unit int32

const (
	Unit_B  Unit = 0
	Unit_KB Unit = 1
	Unit_MB Unit = 2
	Unit_GB Unit = 3
	Unit_C  Unit = 4
)

var Unit_name = map[int32]string{
	0: "B",
	1: "KB",
	2: "MB",
	3: "GB",
	4: "C",
}

var Unit_value = map[string]int32{
	"B":  0,
	"KB": 1,
	"MB": 2,
	"GB": 3,
	"C":  4,
}

func (x Unit) String() string {
	return proto.EnumName(Unit_name, int32(x))
}

func (Unit) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9866e87da4c4a057, []int{2}
}

type RequestHeader struct {
	// the node id of the sender
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestHeader) Reset()         { *m = RequestHeader{} }
func (m *RequestHeader) String() string { return proto.CompactTextString(m) }
func (*RequestHeader) ProtoMessage()    {}
func (*RequestHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_9866e87da4c4a057, []int{0}
}

func (m *RequestHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestHeader.Unmarshal(m, b)
}
func (m *RequestHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestHeader.Marshal(b, m, deterministic)
}
func (m *RequestHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestHeader.Merge(m, src)
}
func (m *RequestHeader) XXX_Size() int {
	return xxx_messageInfo_RequestHeader.Size(m)
}
func (m *RequestHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestHeader.DiscardUnknown(m)
}

var xxx_messageInfo_RequestHeader proto.InternalMessageInfo

func (m *RequestHeader) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type ResponseHeader struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseHeader) Reset()         { *m = ResponseHeader{} }
func (m *ResponseHeader) String() string { return proto.CompactTextString(m) }
func (*ResponseHeader) ProtoMessage()    {}
func (*ResponseHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_9866e87da4c4a057, []int{1}
}

func (m *ResponseHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseHeader.Unmarshal(m, b)
}
func (m *ResponseHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseHeader.Marshal(b, m, deterministic)
}
func (m *ResponseHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseHeader.Merge(m, src)
}
func (m *ResponseHeader) XXX_Size() int {
	return xxx_messageInfo_ResponseHeader.Size(m)
}
func (m *ResponseHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseHeader.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseHeader proto.InternalMessageInfo

func (m *ResponseHeader) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *ResponseHeader) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type Error struct {
	Type                 ErrorType `protobuf:"varint,1,opt,name=type,proto3,enum=executorpb.ErrorType" json:"type,omitempty"`
	Message              string    `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_9866e87da4c4a057, []int{2}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetType() ErrorType {
	if m != nil {
		return m.Type
	}
	return ErrorType_OK
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Resource struct {
	Type                 StorageType `protobuf:"varint,1,opt,name=type,proto3,enum=executorpb.StorageType" json:"type,omitempty"`
	Kind                 string      `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	Size                 uint64      `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Unit                 Unit        `protobuf:"varint,4,opt,name=unit,proto3,enum=executorpb.Unit" json:"unit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_9866e87da4c4a057, []int{3}
}

func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetType() StorageType {
	if m != nil {
		return m.Type
	}
	return StorageType_STORAGE
}

func (m *Resource) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *Resource) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Resource) GetUnit() Unit {
	if m != nil {
		return m.Unit
	}
	return Unit_B
}

type Device struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Maj                  uint64   `protobuf:"varint,2,opt,name=maj,proto3" json:"maj,omitempty"`
	Min                  uint64   `protobuf:"varint,3,opt,name=min,proto3" json:"min,omitempty"`
	Id                   string   `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	Path                 string   `protobuf:"bytes,5,opt,name=path,proto3" json:"path,omitempty"`
	Vg                   string   `protobuf:"bytes,6,opt,name=vg,proto3" json:"vg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Device) Reset()         { *m = Device{} }
func (m *Device) String() string { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()    {}
func (*Device) Descriptor() ([]byte, []int) {
	return fileDescriptor_9866e87da4c4a057, []int{4}
}

func (m *Device) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Device.Unmarshal(m, b)
}
func (m *Device) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Device.Marshal(b, m, deterministic)
}
func (m *Device) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Device.Merge(m, src)
}
func (m *Device) XXX_Size() int {
	return xxx_messageInfo_Device.Size(m)
}
func (m *Device) XXX_DiscardUnknown() {
	xxx_messageInfo_Device.DiscardUnknown(m)
}

var xxx_messageInfo_Device proto.InternalMessageInfo

func (m *Device) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Device) GetMaj() uint64 {
	if m != nil {
		return m.Maj
	}
	return 0
}

func (m *Device) GetMin() uint64 {
	if m != nil {
		return m.Min
	}
	return 0
}

func (m *Device) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Device) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Device) GetVg() string {
	if m != nil {
		return m.Vg
	}
	return ""
}

type PutIsolationRequest struct {
	Header               *RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Deivice              *Device        `protobuf:"bytes,2,opt,name=deivice,proto3" json:"deivice,omitempty"`
	Resource             []*Resource    `protobuf:"bytes,3,rep,name=Resource,proto3" json:"Resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PutIsolationRequest) Reset()         { *m = PutIsolationRequest{} }
func (m *PutIsolationRequest) String() string { return proto.CompactTextString(m) }
func (*PutIsolationRequest) ProtoMessage()    {}
func (*PutIsolationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9866e87da4c4a057, []int{5}
}

func (m *PutIsolationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutIsolationRequest.Unmarshal(m, b)
}
func (m *PutIsolationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutIsolationRequest.Marshal(b, m, deterministic)
}
func (m *PutIsolationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutIsolationRequest.Merge(m, src)
}
func (m *PutIsolationRequest) XXX_Size() int {
	return xxx_messageInfo_PutIsolationRequest.Size(m)
}
func (m *PutIsolationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PutIsolationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PutIsolationRequest proto.InternalMessageInfo

func (m *PutIsolationRequest) GetHeader() *RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PutIsolationRequest) GetDeivice() *Device {
	if m != nil {
		return m.Deivice
	}
	return nil
}

func (m *PutIsolationRequest) GetResource() []*Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

type PutIsolationResponse struct {
	Header               *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *PutIsolationResponse) Reset()         { *m = PutIsolationResponse{} }
func (m *PutIsolationResponse) String() string { return proto.CompactTextString(m) }
func (*PutIsolationResponse) ProtoMessage()    {}
func (*PutIsolationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9866e87da4c4a057, []int{6}
}

func (m *PutIsolationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutIsolationResponse.Unmarshal(m, b)
}
func (m *PutIsolationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutIsolationResponse.Marshal(b, m, deterministic)
}
func (m *PutIsolationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutIsolationResponse.Merge(m, src)
}
func (m *PutIsolationResponse) XXX_Size() int {
	return xxx_messageInfo_PutIsolationResponse.Size(m)
}
func (m *PutIsolationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PutIsolationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PutIsolationResponse proto.InternalMessageInfo

func (m *PutIsolationResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func init() {
	proto.RegisterEnum("executorpb.ErrorType", ErrorType_name, ErrorType_value)
	proto.RegisterEnum("executorpb.StorageType", StorageType_name, StorageType_value)
	proto.RegisterEnum("executorpb.Unit", Unit_name, Unit_value)
	proto.RegisterType((*RequestHeader)(nil), "executorpb.RequestHeader")
	proto.RegisterType((*ResponseHeader)(nil), "executorpb.ResponseHeader")
	proto.RegisterType((*Error)(nil), "executorpb.Error")
	proto.RegisterType((*Resource)(nil), "executorpb.Resource")
	proto.RegisterType((*Device)(nil), "executorpb.Device")
	proto.RegisterType((*PutIsolationRequest)(nil), "executorpb.PutIsolationRequest")
	proto.RegisterType((*PutIsolationResponse)(nil), "executorpb.PutIsolationResponse")
}

func init() { proto.RegisterFile("proto/executor.proto", fileDescriptor_9866e87da4c4a057) }

var fileDescriptor_9866e87da4c4a057 = []byte{
	// 558 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x5d, 0x73, 0xd2, 0x4c,
	0x14, 0x26, 0x9f, 0x94, 0xc3, 0xfb, 0x32, 0xdb, 0x23, 0x9d, 0xc6, 0xde, 0xc8, 0x64, 0x74, 0x44,
	0xec, 0xa0, 0xc6, 0x5f, 0x00, 0x92, 0xa9, 0x11, 0x48, 0x70, 0x4b, 0xeb, 0x95, 0xc3, 0xa4, 0x64,
	0x87, 0x46, 0x25, 0x89, 0x49, 0x60, 0xac, 0xf7, 0xfe, 0x17, 0x7f, 0xa6, 0xb3, 0x9b, 0xd0, 0x86,
	0xaa, 0xe3, 0xd5, 0x9e, 0xb3, 0xe7, 0xd9, 0xe7, 0x3c, 0xe7, 0x63, 0xa1, 0x9d, 0xa4, 0x71, 0x1e,
	0xbf, 0x60, 0xdf, 0xd8, 0x72, 0x93, 0xc7, 0x69, 0x5f, 0xb8, 0x08, 0x3b, 0x3f, 0xb9, 0x32, 0xbb,
	0xf0, 0x3f, 0x65, 0x5f, 0x37, 0x2c, 0xcb, 0xdf, 0x32, 0x3f, 0x60, 0x29, 0x1e, 0x43, 0x3d, 0x8a,
	0x03, 0xb6, 0x08, 0x03, 0x43, 0xea, 0x48, 0xdd, 0x06, 0xd5, 0xb9, 0xeb, 0x04, 0x26, 0x85, 0x16,
	0x65, 0x59, 0x12, 0x47, 0x19, 0xfb, 0x07, 0x14, 0x9f, 0x82, 0xc6, 0xd2, 0x34, 0x4e, 0x0d, 0xb9,
	0x23, 0x75, 0x9b, 0xd6, 0x61, 0xff, 0x2e, 0x61, 0xdf, 0xe6, 0x01, 0x5a, 0xc4, 0xcd, 0x09, 0x68,
	0xc2, 0xc7, 0x67, 0xa0, 0xe6, 0x37, 0x09, 0x13, 0x3c, 0x2d, 0xeb, 0xe8, 0xb7, 0x07, 0xf3, 0x9b,
	0x84, 0x51, 0x01, 0x41, 0x03, 0xea, 0x6b, 0x96, 0x65, 0xfe, 0x8a, 0x09, 0xfa, 0x06, 0xdd, 0xb9,
	0xe6, 0x0f, 0x09, 0x0e, 0x28, 0xcb, 0xe2, 0x4d, 0xba, 0x64, 0xf8, 0x7c, 0x8f, 0xf1, 0xb8, 0xca,
	0x78, 0x9e, 0xc7, 0xa9, 0xbf, 0x62, 0x15, 0x4e, 0x04, 0xf5, 0x73, 0x18, 0x05, 0x25, 0xa1, 0xb0,
	0xf9, 0x5d, 0x16, 0x7e, 0x67, 0x86, 0xd2, 0x91, 0xba, 0x2a, 0x15, 0x36, 0x3e, 0x06, 0x75, 0x13,
	0x85, 0xb9, 0xa1, 0x0a, 0x52, 0x52, 0x25, 0xbd, 0x88, 0xc2, 0x9c, 0x8a, 0xa8, 0x99, 0x83, 0x3e,
	0x62, 0xdb, 0x70, 0x29, 0x78, 0x23, 0x7f, 0xcd, 0xca, 0xf6, 0x08, 0x1b, 0x09, 0x28, 0x6b, 0xff,
	0x93, 0x48, 0xa5, 0x52, 0x6e, 0x8a, 0x9b, 0x30, 0x2a, 0x13, 0x71, 0x13, 0x5b, 0x20, 0x87, 0x81,
	0xc8, 0xd2, 0xa0, 0x72, 0x28, 0xb4, 0x24, 0x7e, 0x7e, 0x6d, 0x68, 0x05, 0x0f, 0xb7, 0x39, 0x66,
	0xbb, 0x32, 0xf4, 0x02, 0xb3, 0x5d, 0x99, 0x3f, 0x25, 0x78, 0x30, 0xdb, 0xe4, 0x4e, 0x16, 0x7f,
	0xf1, 0xf3, 0x30, 0x8e, 0xca, 0xb1, 0xe2, 0x2b, 0xd0, 0xaf, 0xc5, 0xbc, 0x84, 0x8a, 0xa6, 0xf5,
	0xb0, 0xaa, 0x7a, 0x6f, 0xf6, 0xb4, 0x04, 0xe2, 0x29, 0xd4, 0x03, 0x16, 0xf2, 0x0a, 0xca, 0x09,
	0x62, 0xf5, 0x4d, 0x51, 0x1b, 0xdd, 0x41, 0xf0, 0xe5, 0x5d, 0xd7, 0x0d, 0xa5, 0xa3, 0x74, 0x9b,
	0x56, 0x7b, 0x3f, 0x45, 0x11, 0xa3, 0xb7, 0x28, 0xf3, 0x1d, 0xb4, 0xf7, 0x95, 0x16, 0x6b, 0x85,
	0xd6, 0x3d, 0xa9, 0x27, 0xf7, 0x78, 0x2a, 0xcb, 0xb7, 0xd3, 0xda, 0x5b, 0x42, 0xe3, 0x76, 0x43,
	0x50, 0x07, 0xd9, 0x1b, 0x93, 0x1a, 0x36, 0xa1, 0x7e, 0xe1, 0x8e, 0x5d, 0xef, 0x83, 0x4b, 0x24,
	0xee, 0xb8, 0xde, 0xc2, 0xf5, 0x46, 0x36, 0x91, 0xb1, 0x01, 0x9a, 0xeb, 0x2d, 0x66, 0x97, 0x44,
	0x41, 0x84, 0x96, 0xe3, 0xce, 0x6d, 0xea, 0x0e, 0x26, 0x0b, 0x9b, 0x52, 0x8f, 0x12, 0x15, 0x8f,
	0xe0, 0xd0, 0x71, 0x2f, 0x07, 0x13, 0x67, 0xb4, 0x98, 0x0d, 0xe8, 0x60, 0x6a, 0xcf, 0x6d, 0x4a,
	0xb4, 0xde, 0x13, 0x68, 0x56, 0x96, 0x86, 0x33, 0x9e, 0xcf, 0x3d, 0x3a, 0x38, 0xb3, 0x49, 0x8d,
	0x33, 0x4e, 0x9c, 0xa9, 0x33, 0x27, 0x52, 0xef, 0x14, 0x54, 0xbe, 0x06, 0xa8, 0x81, 0x34, 0x24,
	0x35, 0xae, 0x66, 0x3c, 0x24, 0x12, 0x3f, 0xa7, 0x43, 0x22, 0xf3, 0xf3, 0x6c, 0x48, 0x14, 0x1e,
	0x7e, 0x43, 0x54, 0xeb, 0x23, 0x1c, 0xec, 0xca, 0xc3, 0xf7, 0xf0, 0x5f, 0xb5, 0x23, 0xf8, 0xa8,
	0x5a, 0xf9, 0x1f, 0xa6, 0x7a, 0xd2, 0xf9, 0x3b, 0xa0, 0x68, 0xd3, 0x95, 0x2e, 0x3e, 0xfb, 0xeb,
	0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x50, 0xed, 0xea, 0x1d, 0x04, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExecutorClient is the client API for Executor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExecutorClient interface {
	PutIsolation(ctx context.Context, in *PutIsolationRequest, opts ...grpc.CallOption) (*PutIsolationResponse, error)
}

type executorClient struct {
	cc *grpc.ClientConn
}

func NewExecutorClient(cc *grpc.ClientConn) ExecutorClient {
	return &executorClient{cc}
}

func (c *executorClient) PutIsolation(ctx context.Context, in *PutIsolationRequest, opts ...grpc.CallOption) (*PutIsolationResponse, error) {
	out := new(PutIsolationResponse)
	err := c.cc.Invoke(ctx, "/executorpb.executor/PutIsolation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExecutorServer is the server API for Executor service.
type ExecutorServer interface {
	PutIsolation(context.Context, *PutIsolationRequest) (*PutIsolationResponse, error)
}

// UnimplementedExecutorServer can be embedded to have forward compatible implementations.
type UnimplementedExecutorServer struct {
}

func (*UnimplementedExecutorServer) PutIsolation(ctx context.Context, req *PutIsolationRequest) (*PutIsolationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutIsolation not implemented")
}

func RegisterExecutorServer(s *grpc.Server, srv ExecutorServer) {
	s.RegisterService(&_Executor_serviceDesc, srv)
}

func _Executor_PutIsolation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutIsolationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).PutIsolation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/executorpb.executor/PutIsolation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).PutIsolation(ctx, req.(*PutIsolationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Executor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "executorpb.executor",
	HandlerType: (*ExecutorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PutIsolation",
			Handler:    _Executor_PutIsolation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/executor.proto",
}