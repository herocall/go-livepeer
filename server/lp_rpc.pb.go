// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server/lp_rpc.proto

package server

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// This request is sent by the broadcaster in `GetTranscoder` to request
// information on which transcoder to use.
type TranscoderRequest struct {
	// ID of the job that the broadcaster needs a transcoder for
	JobId int64 `protobuf:"varint,1,opt,name=jobId" json:"jobId,omitempty"`
	// Broadcaster's signature over the jobId
	Sig                  []byte   `protobuf:"bytes,2,opt,name=sig,proto3" json:"sig,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TranscoderRequest) Reset()         { *m = TranscoderRequest{} }
func (m *TranscoderRequest) String() string { return proto.CompactTextString(m) }
func (*TranscoderRequest) ProtoMessage()    {}
func (*TranscoderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_lp_rpc_0a91ed59a4f20ecf, []int{0}
}
func (m *TranscoderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TranscoderRequest.Unmarshal(m, b)
}
func (m *TranscoderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TranscoderRequest.Marshal(b, m, deterministic)
}
func (dst *TranscoderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TranscoderRequest.Merge(dst, src)
}
func (m *TranscoderRequest) XXX_Size() int {
	return xxx_messageInfo_TranscoderRequest.Size(m)
}
func (m *TranscoderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TranscoderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TranscoderRequest proto.InternalMessageInfo

func (m *TranscoderRequest) GetJobId() int64 {
	if m != nil {
		return m.JobId
	}
	return 0
}

func (m *TranscoderRequest) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

// The orchestrator sends this in response to `GetTranscoder`, containing the
// transcoder URI, associated credentials authorizing the broadcaster to
// use the transcoder, and miscellaneous data related to the job.
type TranscoderInfo struct {
	// URI of the transcoder to use for submitting segments
	Transcoder string `protobuf:"bytes,1,opt,name=transcoder" json:"transcoder,omitempty"`
	// Signals the authentication method to expect within `credentials`. This
	// field is opaque to the broadcaster, and should be passed to the transcoder.
	AuthType string `protobuf:"bytes,2,opt,name=authType" json:"authType,omitempty"`
	// Credentials to verify the request has been authorized by an orchestrator.
	// This field is opaque to the broadcaster.
	Credentials string `protobuf:"bytes,3,opt,name=credentials" json:"credentials,omitempty"`
	// Transcoded streamId list to update the master manifest on the broadcaster.
	StreamIds            map[string]string `protobuf:"bytes,16,rep,name=streamIds" json:"streamIds,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TranscoderInfo) Reset()         { *m = TranscoderInfo{} }
func (m *TranscoderInfo) String() string { return proto.CompactTextString(m) }
func (*TranscoderInfo) ProtoMessage()    {}
func (*TranscoderInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_lp_rpc_0a91ed59a4f20ecf, []int{1}
}
func (m *TranscoderInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TranscoderInfo.Unmarshal(m, b)
}
func (m *TranscoderInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TranscoderInfo.Marshal(b, m, deterministic)
}
func (dst *TranscoderInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TranscoderInfo.Merge(dst, src)
}
func (m *TranscoderInfo) XXX_Size() int {
	return xxx_messageInfo_TranscoderInfo.Size(m)
}
func (m *TranscoderInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TranscoderInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TranscoderInfo proto.InternalMessageInfo

func (m *TranscoderInfo) GetTranscoder() string {
	if m != nil {
		return m.Transcoder
	}
	return ""
}

func (m *TranscoderInfo) GetAuthType() string {
	if m != nil {
		return m.AuthType
	}
	return ""
}

func (m *TranscoderInfo) GetCredentials() string {
	if m != nil {
		return m.Credentials
	}
	return ""
}

func (m *TranscoderInfo) GetStreamIds() map[string]string {
	if m != nil {
		return m.StreamIds
	}
	return nil
}

// AuthToken is sent by the orchestrator and encoded in the `credentials` field
// This record is opaque to the broadcaster and is only relevant between the
// orchestrator and the transcoder.
type AuthToken struct {
	// Signature of the orchestrator over the remaining fields
	Sig                  []byte   `protobuf:"bytes,1,opt,name=sig,proto3" json:"sig,omitempty"`
	JobId                int64    `protobuf:"varint,16,opt,name=jobId" json:"jobId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthToken) Reset()         { *m = AuthToken{} }
func (m *AuthToken) String() string { return proto.CompactTextString(m) }
func (*AuthToken) ProtoMessage()    {}
func (*AuthToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_lp_rpc_0a91ed59a4f20ecf, []int{2}
}
func (m *AuthToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthToken.Unmarshal(m, b)
}
func (m *AuthToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthToken.Marshal(b, m, deterministic)
}
func (dst *AuthToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthToken.Merge(dst, src)
}
func (m *AuthToken) XXX_Size() int {
	return xxx_messageInfo_AuthToken.Size(m)
}
func (m *AuthToken) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthToken.DiscardUnknown(m)
}

var xxx_messageInfo_AuthToken proto.InternalMessageInfo

func (m *AuthToken) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

func (m *AuthToken) GetJobId() int64 {
	if m != nil {
		return m.JobId
	}
	return 0
}

// Data included by the broadcaster when submitting a segment for transcoding.
type SegData struct {
	// Sequence number of the segment to be transcoded
	Seq int64 `protobuf:"varint,1,opt,name=seq" json:"seq,omitempty"`
	// Hash of the segment data to be transcoded
	Hash []byte `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	// Broadcaster signature for the segment. Corresponds to:
	// broadcaster.sign(streamId | seqNo | dataHash)
	// where streamId is derived from the jobId
	Sig                  []byte   `protobuf:"bytes,3,opt,name=sig,proto3" json:"sig,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SegData) Reset()         { *m = SegData{} }
func (m *SegData) String() string { return proto.CompactTextString(m) }
func (*SegData) ProtoMessage()    {}
func (*SegData) Descriptor() ([]byte, []int) {
	return fileDescriptor_lp_rpc_0a91ed59a4f20ecf, []int{3}
}
func (m *SegData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SegData.Unmarshal(m, b)
}
func (m *SegData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SegData.Marshal(b, m, deterministic)
}
func (dst *SegData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SegData.Merge(dst, src)
}
func (m *SegData) XXX_Size() int {
	return xxx_messageInfo_SegData.Size(m)
}
func (m *SegData) XXX_DiscardUnknown() {
	xxx_messageInfo_SegData.DiscardUnknown(m)
}

var xxx_messageInfo_SegData proto.InternalMessageInfo

func (m *SegData) GetSeq() int64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *SegData) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *SegData) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

func init() {
	proto.RegisterType((*TranscoderRequest)(nil), "server.TranscoderRequest")
	proto.RegisterType((*TranscoderInfo)(nil), "server.TranscoderInfo")
	proto.RegisterMapType((map[string]string)(nil), "server.TranscoderInfo.StreamIdsEntry")
	proto.RegisterType((*AuthToken)(nil), "server.AuthToken")
	proto.RegisterType((*SegData)(nil), "server.SegData")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OrchestratorClient is the client API for Orchestrator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrchestratorClient interface {
	// Called by the broadcaster to request transcoder info from an orchestrator.
	GetTranscoder(ctx context.Context, in *TranscoderRequest, opts ...grpc.CallOption) (*TranscoderInfo, error)
}

type orchestratorClient struct {
	cc *grpc.ClientConn
}

func NewOrchestratorClient(cc *grpc.ClientConn) OrchestratorClient {
	return &orchestratorClient{cc}
}

func (c *orchestratorClient) GetTranscoder(ctx context.Context, in *TranscoderRequest, opts ...grpc.CallOption) (*TranscoderInfo, error) {
	out := new(TranscoderInfo)
	err := c.cc.Invoke(ctx, "/server.Orchestrator/GetTranscoder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrchestratorServer is the server API for Orchestrator service.
type OrchestratorServer interface {
	// Called by the broadcaster to request transcoder info from an orchestrator.
	GetTranscoder(context.Context, *TranscoderRequest) (*TranscoderInfo, error)
}

func RegisterOrchestratorServer(s *grpc.Server, srv OrchestratorServer) {
	s.RegisterService(&_Orchestrator_serviceDesc, srv)
}

func _Orchestrator_GetTranscoder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TranscoderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrchestratorServer).GetTranscoder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Orchestrator/GetTranscoder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrchestratorServer).GetTranscoder(ctx, req.(*TranscoderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Orchestrator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "server.Orchestrator",
	HandlerType: (*OrchestratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTranscoder",
			Handler:    _Orchestrator_GetTranscoder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/lp_rpc.proto",
}

func init() { proto.RegisterFile("server/lp_rpc.proto", fileDescriptor_lp_rpc_0a91ed59a4f20ecf) }

var fileDescriptor_lp_rpc_0a91ed59a4f20ecf = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4d, 0x4f, 0xfa, 0x40,
	0x10, 0xc6, 0x53, 0xfa, 0xff, 0xa3, 0x1d, 0x90, 0xe0, 0x6a, 0x4c, 0xe5, 0x60, 0x9a, 0x26, 0x26,
	0x9c, 0x6a, 0x02, 0x17, 0xa3, 0x5e, 0xf0, 0x25, 0x86, 0x93, 0xc9, 0xc2, 0xdd, 0x2c, 0xed, 0x48,
	0x11, 0xec, 0x96, 0xd9, 0x85, 0x84, 0x8f, 0xed, 0x37, 0x30, 0xdd, 0xbe, 0x41, 0xf4, 0x36, 0xf3,
	0x4c, 0x9f, 0xc9, 0x33, 0xbf, 0x2e, 0x9c, 0x29, 0xa4, 0x2d, 0xd2, 0xcd, 0x2a, 0x7d, 0xa7, 0x34,
	0x0c, 0x52, 0x92, 0x5a, 0xb2, 0x66, 0x2e, 0xfa, 0xf7, 0x70, 0x3a, 0x25, 0x91, 0xa8, 0x50, 0x46,
	0x48, 0x1c, 0xd7, 0x1b, 0x54, 0x9a, 0x9d, 0xc3, 0xff, 0x4f, 0x39, 0x1b, 0x47, 0xae, 0xe5, 0x59,
	0x7d, 0x9b, 0xe7, 0x0d, 0xeb, 0x82, 0xad, 0x16, 0x73, 0xb7, 0xe1, 0x59, 0xfd, 0x36, 0xcf, 0x4a,
	0xff, 0xdb, 0x82, 0x4e, 0xed, 0x1e, 0x27, 0x1f, 0x92, 0x5d, 0x01, 0xe8, 0x4a, 0x31, 0x7e, 0x87,
	0xef, 0x29, 0xac, 0x07, 0xc7, 0x62, 0xa3, 0xe3, 0xe9, 0x2e, 0x45, 0xb3, 0xc9, 0xe1, 0x55, 0xcf,
	0x3c, 0x68, 0x85, 0x84, 0x11, 0x26, 0x7a, 0x21, 0x56, 0xca, 0xb5, 0xcd, 0x78, 0x5f, 0x62, 0x4f,
	0xe0, 0x28, 0x4d, 0x28, 0xbe, 0xc6, 0x91, 0x72, 0xbb, 0x9e, 0xdd, 0x6f, 0x0d, 0xae, 0x83, 0xfc,
	0x92, 0xe0, 0x30, 0x48, 0x30, 0x29, 0xbf, 0x7b, 0x49, 0x34, 0xed, 0x78, 0xed, 0xeb, 0x3d, 0x40,
	0xe7, 0x70, 0x98, 0x5d, 0xb6, 0xc4, 0x5d, 0x91, 0x36, 0x2b, 0x33, 0x02, 0x5b, 0xb1, 0xda, 0x94,
	0x19, 0xf3, 0xe6, 0xae, 0x71, 0x6b, 0xf9, 0x43, 0x70, 0x46, 0x59, 0x60, 0xb9, 0xc4, 0xa4, 0x44,
	0x62, 0x55, 0x48, 0x6a, 0x74, 0xdd, 0x3d, 0x74, 0xfe, 0x08, 0x8e, 0x26, 0x38, 0x7f, 0x16, 0x5a,
	0x18, 0x0b, 0xae, 0x0b, 0xb2, 0x59, 0xc9, 0x18, 0xfc, 0x8b, 0x85, 0x8a, 0x0b, 0xb0, 0xa6, 0x2e,
	0x17, 0xdb, 0xd5, 0xe2, 0x01, 0x87, 0xf6, 0x1b, 0x85, 0x31, 0x2a, 0x4d, 0x42, 0x4b, 0x62, 0x8f,
	0x70, 0xf2, 0x8a, 0xba, 0x3e, 0x9a, 0x5d, 0xfe, 0x06, 0x51, 0xfc, 0xcf, 0xde, 0xc5, 0xdf, 0x8c,
	0x66, 0x4d, 0xf3, 0x16, 0x86, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb0, 0xd4, 0x13, 0xb6, 0x22,
	0x02, 0x00, 0x00,
}
