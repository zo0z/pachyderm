// Code generated by protoc-gen-gogo.
// source: server/pkg/worker/worker_service.proto
// DO NOT EDIT!

/*
Package worker is a generated protocol buffer package.

It is generated from these files:
	server/pkg/worker/worker_service.proto

It has these top-level messages:
	ProcessRequest
	ProcessResponse
*/
package worker

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import pfs "github.com/pachyderm/pachyderm/src/client/pfs"

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type DatumState int32

const (
	// The user's process started, (possibly) read this datum, produced output,
	// and exited successfully
	DatumState_DATUM_SUCCESS DatumState = 0
	// Worker failed to download the input data, or otherwise couldn't
	// prepare the container to start the user's process (or it couldn't upload
	// the data after the user's process finished successfully)
	DatumState_DATUM_INTERNAL_ERROR DatumState = 1
	// The user's process didn't start, or it exited with an error code
	DatumState_DATUM_FAILURE DatumState = 2
)

var DatumState_name = map[int32]string{
	0: "DATUM_SUCCESS",
	1: "DATUM_INTERNAL_ERROR",
	2: "DATUM_FAILURE",
}
var DatumState_value = map[string]int32{
	"DATUM_SUCCESS":        0,
	"DATUM_INTERNAL_ERROR": 1,
	"DATUM_FAILURE":        2,
}

func (x DatumState) String() string {
	return proto.EnumName(DatumState_name, int32(x))
}
func (DatumState) EnumDescriptor() ([]byte, []int) { return fileDescriptorWorkerService, []int{0} }

type ProcessRequest struct {
	// The file to process
	Data []*pfs.FileInfo `protobuf:"bytes,1,rep,name=data" json:"data,omitempty"`
	// The tag to write the output hashtree under
	// (Possible alternative: return blobref)
	Tag string `protobuf:"bytes,2,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (m *ProcessRequest) Reset()                    { *m = ProcessRequest{} }
func (m *ProcessRequest) String() string            { return proto.CompactTextString(m) }
func (*ProcessRequest) ProtoMessage()               {}
func (*ProcessRequest) Descriptor() ([]byte, []int) { return fileDescriptorWorkerService, []int{0} }

func (m *ProcessRequest) GetData() []*pfs.FileInfo {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ProcessRequest) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

type ProcessResponse struct {
	// The final state after trying to process ProcessRequest.Datum
	State DatumState `protobuf:"varint,1,opt,name=state,proto3,enum=worker.DatumState" json:"state,omitempty"`
}

func (m *ProcessResponse) Reset()                    { *m = ProcessResponse{} }
func (m *ProcessResponse) String() string            { return proto.CompactTextString(m) }
func (*ProcessResponse) ProtoMessage()               {}
func (*ProcessResponse) Descriptor() ([]byte, []int) { return fileDescriptorWorkerService, []int{1} }

func (m *ProcessResponse) GetState() DatumState {
	if m != nil {
		return m.State
	}
	return DatumState_DATUM_SUCCESS
}

func init() {
	proto.RegisterType((*ProcessRequest)(nil), "worker.ProcessRequest")
	proto.RegisterType((*ProcessResponse)(nil), "worker.ProcessResponse")
	proto.RegisterEnum("worker.DatumState", DatumState_name, DatumState_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Worker service

type WorkerClient interface {
	Process(ctx context.Context, in *ProcessRequest, opts ...grpc.CallOption) (*ProcessResponse, error)
}

type workerClient struct {
	cc *grpc.ClientConn
}

func NewWorkerClient(cc *grpc.ClientConn) WorkerClient {
	return &workerClient{cc}
}

func (c *workerClient) Process(ctx context.Context, in *ProcessRequest, opts ...grpc.CallOption) (*ProcessResponse, error) {
	out := new(ProcessResponse)
	err := grpc.Invoke(ctx, "/worker.Worker/Process", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Worker service

type WorkerServer interface {
	Process(context.Context, *ProcessRequest) (*ProcessResponse, error)
}

func RegisterWorkerServer(s *grpc.Server, srv WorkerServer) {
	s.RegisterService(&_Worker_serviceDesc, srv)
}

func _Worker_Process_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Process(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/worker.Worker/Process",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Process(ctx, req.(*ProcessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Worker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "worker.Worker",
	HandlerType: (*WorkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Process",
			Handler:    _Worker_Process_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/pkg/worker/worker_service.proto",
}

func init() { proto.RegisterFile("server/pkg/worker/worker_service.proto", fileDescriptorWorkerService) }

var fileDescriptorWorkerService = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x90, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0xd7, 0x4d, 0x27, 0x3e, 0xd9, 0xac, 0x61, 0x68, 0xd9, 0xa9, 0xf6, 0x20, 0xc5, 0x43,
	0x0b, 0xf5, 0xa8, 0x97, 0xb2, 0xb5, 0x50, 0xa8, 0x53, 0xd2, 0x15, 0x8f, 0xa5, 0xd6, 0xb7, 0x51,
	0x36, 0x9b, 0x9a, 0x64, 0xfa, 0xef, 0x4b, 0x9a, 0xe9, 0x10, 0x0f, 0x21, 0xe1, 0xf7, 0xfb, 0xc8,
	0xfb, 0x12, 0xb8, 0x11, 0xc8, 0x3f, 0x91, 0xfb, 0xed, 0x66, 0xed, 0x7f, 0x31, 0xbe, 0x41, 0xbe,
	0xdf, 0x0a, 0x25, 0xea, 0x0a, 0xbd, 0x96, 0x33, 0xc9, 0xc8, 0x50, 0xd3, 0xe9, 0xa4, 0xda, 0xd6,
	0xd8, 0x48, 0xbf, 0x5d, 0x09, 0xb5, 0xb4, 0x75, 0x22, 0x18, 0x3f, 0x73, 0x56, 0xa1, 0x10, 0x14,
	0x3f, 0x76, 0x28, 0x24, 0xb9, 0x86, 0xa3, 0xb7, 0x52, 0x96, 0x96, 0x61, 0x0f, 0xdc, 0xb3, 0x60,
	0xe4, 0xa9, 0x6c, 0x5c, 0x6f, 0x31, 0x69, 0x56, 0x8c, 0x76, 0x8a, 0x98, 0x30, 0x90, 0xe5, 0xda,
	0xea, 0xdb, 0x86, 0x7b, 0x4a, 0xd5, 0xd1, 0xb9, 0x87, 0xf3, 0xdf, 0x6b, 0x44, 0xcb, 0x1a, 0x81,
	0xc4, 0x85, 0x63, 0x21, 0x4b, 0x89, 0x96, 0x61, 0x1b, 0xee, 0x38, 0x20, 0x9e, 0xee, 0xe1, 0xcd,
	0x4b, 0xb9, 0x7b, 0xcf, 0x94, 0xa1, 0x3a, 0x70, 0x9b, 0x02, 0x1c, 0x20, 0xb9, 0x80, 0xd1, 0x3c,
	0x5c, 0xe6, 0x8f, 0x45, 0x96, 0xcf, 0x66, 0x51, 0x96, 0x99, 0x3d, 0x62, 0xc1, 0x44, 0xa3, 0x64,
	0xb1, 0x8c, 0xe8, 0x22, 0x4c, 0x8b, 0x88, 0xd2, 0x27, 0x6a, 0x1a, 0x87, 0x70, 0x1c, 0x26, 0x69,
	0x4e, 0x23, 0xb3, 0x1f, 0xc4, 0x30, 0x7c, 0xe9, 0x26, 0x91, 0x07, 0x38, 0xd9, 0x97, 0x22, 0x97,
	0x3f, 0xd3, 0xff, 0x3e, 0x76, 0x7a, 0xf5, 0x8f, 0xeb, 0xf6, 0x4e, 0xef, 0x75, 0xd8, 0x7d, 0xd0,
	0xdd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xea, 0x3a, 0x05, 0x0c, 0x68, 0x01, 0x00, 0x00,
}