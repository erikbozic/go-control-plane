// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/access_loggers/grpc/v3/als.proto

package envoy_extensions_access_loggers_grpc_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type HttpGrpcAccessLogConfig struct {
	CommonConfig                    *CommonGrpcAccessLogConfig `protobuf:"bytes,1,opt,name=common_config,json=commonConfig,proto3" json:"common_config,omitempty"`
	AdditionalRequestHeadersToLog   []string                   `protobuf:"bytes,2,rep,name=additional_request_headers_to_log,json=additionalRequestHeadersToLog,proto3" json:"additional_request_headers_to_log,omitempty"`
	AdditionalResponseHeadersToLog  []string                   `protobuf:"bytes,3,rep,name=additional_response_headers_to_log,json=additionalResponseHeadersToLog,proto3" json:"additional_response_headers_to_log,omitempty"`
	AdditionalResponseTrailersToLog []string                   `protobuf:"bytes,4,rep,name=additional_response_trailers_to_log,json=additionalResponseTrailersToLog,proto3" json:"additional_response_trailers_to_log,omitempty"`
	XXX_NoUnkeyedLiteral            struct{}                   `json:"-"`
	XXX_unrecognized                []byte                     `json:"-"`
	XXX_sizecache                   int32                      `json:"-"`
}

func (m *HttpGrpcAccessLogConfig) Reset()         { *m = HttpGrpcAccessLogConfig{} }
func (m *HttpGrpcAccessLogConfig) String() string { return proto.CompactTextString(m) }
func (*HttpGrpcAccessLogConfig) ProtoMessage()    {}
func (*HttpGrpcAccessLogConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3685c6e6be71fba, []int{0}
}

func (m *HttpGrpcAccessLogConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpGrpcAccessLogConfig.Unmarshal(m, b)
}
func (m *HttpGrpcAccessLogConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpGrpcAccessLogConfig.Marshal(b, m, deterministic)
}
func (m *HttpGrpcAccessLogConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpGrpcAccessLogConfig.Merge(m, src)
}
func (m *HttpGrpcAccessLogConfig) XXX_Size() int {
	return xxx_messageInfo_HttpGrpcAccessLogConfig.Size(m)
}
func (m *HttpGrpcAccessLogConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpGrpcAccessLogConfig.DiscardUnknown(m)
}

var xxx_messageInfo_HttpGrpcAccessLogConfig proto.InternalMessageInfo

func (m *HttpGrpcAccessLogConfig) GetCommonConfig() *CommonGrpcAccessLogConfig {
	if m != nil {
		return m.CommonConfig
	}
	return nil
}

func (m *HttpGrpcAccessLogConfig) GetAdditionalRequestHeadersToLog() []string {
	if m != nil {
		return m.AdditionalRequestHeadersToLog
	}
	return nil
}

func (m *HttpGrpcAccessLogConfig) GetAdditionalResponseHeadersToLog() []string {
	if m != nil {
		return m.AdditionalResponseHeadersToLog
	}
	return nil
}

func (m *HttpGrpcAccessLogConfig) GetAdditionalResponseTrailersToLog() []string {
	if m != nil {
		return m.AdditionalResponseTrailersToLog
	}
	return nil
}

type TcpGrpcAccessLogConfig struct {
	CommonConfig         *CommonGrpcAccessLogConfig `protobuf:"bytes,1,opt,name=common_config,json=commonConfig,proto3" json:"common_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *TcpGrpcAccessLogConfig) Reset()         { *m = TcpGrpcAccessLogConfig{} }
func (m *TcpGrpcAccessLogConfig) String() string { return proto.CompactTextString(m) }
func (*TcpGrpcAccessLogConfig) ProtoMessage()    {}
func (*TcpGrpcAccessLogConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3685c6e6be71fba, []int{1}
}

func (m *TcpGrpcAccessLogConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpGrpcAccessLogConfig.Unmarshal(m, b)
}
func (m *TcpGrpcAccessLogConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpGrpcAccessLogConfig.Marshal(b, m, deterministic)
}
func (m *TcpGrpcAccessLogConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpGrpcAccessLogConfig.Merge(m, src)
}
func (m *TcpGrpcAccessLogConfig) XXX_Size() int {
	return xxx_messageInfo_TcpGrpcAccessLogConfig.Size(m)
}
func (m *TcpGrpcAccessLogConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpGrpcAccessLogConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TcpGrpcAccessLogConfig proto.InternalMessageInfo

func (m *TcpGrpcAccessLogConfig) GetCommonConfig() *CommonGrpcAccessLogConfig {
	if m != nil {
		return m.CommonConfig
	}
	return nil
}

type CommonGrpcAccessLogConfig struct {
	LogName                 string                `protobuf:"bytes,1,opt,name=log_name,json=logName,proto3" json:"log_name,omitempty"`
	GrpcService             *v3.GrpcService       `protobuf:"bytes,2,opt,name=grpc_service,json=grpcService,proto3" json:"grpc_service,omitempty"`
	BufferFlushInterval     *duration.Duration    `protobuf:"bytes,3,opt,name=buffer_flush_interval,json=bufferFlushInterval,proto3" json:"buffer_flush_interval,omitempty"`
	BufferSizeBytes         *wrappers.UInt32Value `protobuf:"bytes,4,opt,name=buffer_size_bytes,json=bufferSizeBytes,proto3" json:"buffer_size_bytes,omitempty"`
	FilterStateObjectsToLog []string              `protobuf:"bytes,5,rep,name=filter_state_objects_to_log,json=filterStateObjectsToLog,proto3" json:"filter_state_objects_to_log,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}              `json:"-"`
	XXX_unrecognized        []byte                `json:"-"`
	XXX_sizecache           int32                 `json:"-"`
}

func (m *CommonGrpcAccessLogConfig) Reset()         { *m = CommonGrpcAccessLogConfig{} }
func (m *CommonGrpcAccessLogConfig) String() string { return proto.CompactTextString(m) }
func (*CommonGrpcAccessLogConfig) ProtoMessage()    {}
func (*CommonGrpcAccessLogConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3685c6e6be71fba, []int{2}
}

func (m *CommonGrpcAccessLogConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonGrpcAccessLogConfig.Unmarshal(m, b)
}
func (m *CommonGrpcAccessLogConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonGrpcAccessLogConfig.Marshal(b, m, deterministic)
}
func (m *CommonGrpcAccessLogConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonGrpcAccessLogConfig.Merge(m, src)
}
func (m *CommonGrpcAccessLogConfig) XXX_Size() int {
	return xxx_messageInfo_CommonGrpcAccessLogConfig.Size(m)
}
func (m *CommonGrpcAccessLogConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonGrpcAccessLogConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CommonGrpcAccessLogConfig proto.InternalMessageInfo

func (m *CommonGrpcAccessLogConfig) GetLogName() string {
	if m != nil {
		return m.LogName
	}
	return ""
}

func (m *CommonGrpcAccessLogConfig) GetGrpcService() *v3.GrpcService {
	if m != nil {
		return m.GrpcService
	}
	return nil
}

func (m *CommonGrpcAccessLogConfig) GetBufferFlushInterval() *duration.Duration {
	if m != nil {
		return m.BufferFlushInterval
	}
	return nil
}

func (m *CommonGrpcAccessLogConfig) GetBufferSizeBytes() *wrappers.UInt32Value {
	if m != nil {
		return m.BufferSizeBytes
	}
	return nil
}

func (m *CommonGrpcAccessLogConfig) GetFilterStateObjectsToLog() []string {
	if m != nil {
		return m.FilterStateObjectsToLog
	}
	return nil
}

func init() {
	proto.RegisterType((*HttpGrpcAccessLogConfig)(nil), "envoy.extensions.access_loggers.grpc.v3.HttpGrpcAccessLogConfig")
	proto.RegisterType((*TcpGrpcAccessLogConfig)(nil), "envoy.extensions.access_loggers.grpc.v3.TcpGrpcAccessLogConfig")
	proto.RegisterType((*CommonGrpcAccessLogConfig)(nil), "envoy.extensions.access_loggers.grpc.v3.CommonGrpcAccessLogConfig")
}

func init() {
	proto.RegisterFile("envoy/extensions/access_loggers/grpc/v3/als.proto", fileDescriptor_d3685c6e6be71fba)
}

var fileDescriptor_d3685c6e6be71fba = []byte{
	// 616 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x93, 0xd1, 0x4e, 0x14, 0x3f,
	0x14, 0xc6, 0xff, 0xb3, 0x0b, 0x7f, 0xa0, 0x60, 0xd4, 0x31, 0xca, 0x82, 0x8a, 0xb0, 0x5e, 0x48,
	0xbc, 0xe8, 0xe8, 0x6e, 0x88, 0x86, 0x78, 0xc3, 0x60, 0x74, 0x31, 0x04, 0xc9, 0x82, 0x7a, 0x39,
	0xe9, 0xce, 0x9e, 0x2d, 0x35, 0xdd, 0x76, 0x68, 0x3b, 0x23, 0xf0, 0x04, 0xc6, 0x47, 0xf0, 0x11,
	0x88, 0x8f, 0xa1, 0x2f, 0xc5, 0x95, 0x69, 0x3b, 0xc8, 0xae, 0x30, 0x09, 0x97, 0xde, 0xed, 0x4e,
	0xbf, 0xef, 0x77, 0xfa, 0xf5, 0x9c, 0x83, 0x9e, 0x83, 0x28, 0xe4, 0x71, 0x04, 0x47, 0x06, 0x84,
	0x66, 0x52, 0xe8, 0x88, 0xa4, 0x29, 0x68, 0x9d, 0x70, 0x49, 0x29, 0x28, 0x1d, 0x51, 0x95, 0xa5,
	0x51, 0xd1, 0x8e, 0x08, 0xd7, 0x38, 0x53, 0xd2, 0xc8, 0xf0, 0x89, 0xb3, 0xe0, 0x0b, 0x0b, 0x1e,
	0xb7, 0x60, 0x6b, 0xc1, 0x45, 0x7b, 0xd1, 0x0b, 0xa3, 0x54, 0x8a, 0x01, 0xa3, 0x51, 0x2a, 0x15,
	0x58, 0x90, 0x3d, 0x4d, 0x34, 0xa8, 0x82, 0xa5, 0xe0, 0x89, 0x8b, 0x4b, 0x54, 0x4a, 0xca, 0x21,
	0x72, 0xff, 0x7a, 0xf9, 0x20, 0xea, 0xe7, 0x8a, 0x18, 0x26, 0x45, 0xd5, 0xf9, 0x17, 0x45, 0xb2,
	0xcc, 0x96, 0xf2, 0xe7, 0x2b, 0x79, 0x3f, 0x23, 0x11, 0x11, 0x42, 0x1a, 0x67, 0xd3, 0x51, 0x01,
	0xca, 0x5e, 0x8d, 0x09, 0x5a, 0x4a, 0xe6, 0x0b, 0xc2, 0x59, 0x9f, 0x18, 0x88, 0xce, 0x7f, 0xf8,
	0x83, 0xe6, 0x69, 0x1d, 0xcd, 0x77, 0x8c, 0xc9, 0xde, 0xaa, 0x2c, 0xdd, 0x70, 0x39, 0xb6, 0x25,
	0xdd, 0x74, 0x57, 0x0e, 0x0f, 0xd1, 0x8d, 0x54, 0x0e, 0x87, 0x52, 0x24, 0x3e, 0x43, 0x23, 0x58,
	0x0e, 0x56, 0x67, 0x5b, 0x31, 0xbe, 0xe6, 0x0b, 0xe0, 0x4d, 0xe7, 0xbe, 0x02, 0x1d, 0x4f, 0x9f,
	0xc5, 0x93, 0xdf, 0x82, 0xda, 0xad, 0xa0, 0x3b, 0xe7, 0x4b, 0x94, 0x25, 0x3b, 0x68, 0x85, 0xf4,
	0xfb, 0xcc, 0xa6, 0x20, 0x3c, 0x51, 0x70, 0x98, 0x83, 0x36, 0xc9, 0x01, 0x90, 0x3e, 0x28, 0x9d,
	0x18, 0x69, 0x4b, 0x34, 0x6a, 0xcb, 0xf5, 0xd5, 0x99, 0xee, 0xc3, 0x0b, 0x61, 0xd7, 0xeb, 0x3a,
	0x5e, 0xb6, 0x2f, 0xb7, 0x25, 0x0d, 0xdf, 0xa1, 0xe6, 0x18, 0x49, 0x67, 0x52, 0x68, 0xf8, 0x1b,
	0x55, 0x77, 0xa8, 0xa5, 0x51, 0x94, 0x17, 0x8e, 0xb1, 0xb6, 0xd1, 0xe3, 0xab, 0x58, 0x46, 0x11,
	0xc6, 0x47, 0x60, 0x13, 0x0e, 0xf6, 0xe8, 0x32, 0x6c, 0xbf, 0x14, 0x3a, 0xda, 0xfa, 0xcb, 0xef,
	0xbf, 0xbe, 0x2e, 0xb5, 0xcb, 0xd1, 0xc3, 0xfe, 0x69, 0xcb, 0x17, 0xe4, 0x92, 0xe2, 0xa2, 0x85,
	0x2b, 0x1a, 0xd2, 0xfc, 0x19, 0xa0, 0x7b, 0xfb, 0xe9, 0x3f, 0xd2, 0xab, 0xf5, 0x17, 0x36, 0x47,
	0x0b, 0x3d, 0xab, 0xce, 0x71, 0xf5, 0x5d, 0x9b, 0x3f, 0xea, 0x68, 0xa1, 0xb2, 0x5c, 0xd8, 0x44,
	0xd3, 0x5c, 0xd2, 0x44, 0x90, 0x21, 0xb8, 0x10, 0x33, 0xf1, 0xd4, 0x59, 0x3c, 0xa1, 0x6a, 0xcb,
	0x41, 0x77, 0x8a, 0x4b, 0xba, 0x43, 0x86, 0x10, 0xee, 0xa0, 0xb9, 0xd1, 0x3d, 0x6a, 0xd4, 0x5c,
	0xd8, 0x15, 0x3c, 0x76, 0x15, 0xbb, 0x71, 0x36, 0x99, 0x2d, 0xb2, 0xe7, 0x85, 0x23, 0x59, 0x66,
	0xe9, 0xc5, 0xe7, 0xf0, 0x13, 0xba, 0xdb, 0xcb, 0x07, 0x03, 0x50, 0xc9, 0x80, 0xe7, 0xfa, 0x20,
	0x61, 0xc2, 0x80, 0x2a, 0x08, 0x6f, 0xd4, 0x1d, 0x78, 0x01, 0xfb, 0x0d, 0xc4, 0xe7, 0x1b, 0x88,
	0x5f, 0x97, 0x1b, 0xea, 0x80, 0xa7, 0x41, 0xed, 0xe9, 0x7f, 0xdd, 0x3b, 0x9e, 0xf0, 0xc6, 0x02,
	0xb6, 0x4a, 0x7f, 0xd8, 0x41, 0xb7, 0x4b, 0xb0, 0x66, 0x27, 0x90, 0xf4, 0x8e, 0x0d, 0xe8, 0xc6,
	0x84, 0x83, 0x3e, 0xb8, 0x04, 0xfd, 0xb0, 0x25, 0x4c, 0xbb, 0xf5, 0x91, 0xf0, 0x1c, 0xba, 0x37,
	0xbd, 0x6d, 0x8f, 0x9d, 0x40, 0x6c, 0x4d, 0xe1, 0x2b, 0x74, 0x7f, 0xc0, 0xb8, 0xb1, 0x24, 0x43,
	0x0c, 0x24, 0xb2, 0xf7, 0x19, 0x52, 0xf3, 0x67, 0xf6, 0x26, 0xdd, 0xec, 0xcd, 0x7b, 0xc9, 0x9e,
	0x55, 0xbc, 0xf7, 0x02, 0x3f, 0x73, 0xeb, 0xb6, 0x57, 0x6b, 0xa8, 0x5d, 0xdd, 0xab, 0xea, 0xfe,
	0x6f, 0xa2, 0x35, 0x26, 0xfd, 0xd3, 0x66, 0x4a, 0x1e, 0x1d, 0x5f, 0x77, 0xa4, 0xe2, 0xe9, 0x0d,
	0xae, 0x77, 0x6d, 0xb8, 0xdd, 0xa0, 0xf7, 0xbf, 0x4b, 0xd9, 0xfe, 0x1d, 0x00, 0x00, 0xff, 0xff,
	0xf1, 0xb5, 0xbb, 0xf0, 0x71, 0x05, 0x00, 0x00,
}
