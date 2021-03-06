// Code generated by protoc-gen-gogo.
// source: server/pkg/deploy/deploy.proto
// DO NOT EDIT!

/*
Package deploy is a generated protocol buffer package.

It is generated from these files:
	server/pkg/deploy/deploy.proto

It has these top-level messages:
	KubeEndpoint
	Cluster
	ClusterInfo
	ClusterInfos
	CreateClusterRequest
	UpdateClusterRequest
	InspectClusterRequest
	ListClusterRequest
	DeleteClusterRequest
*/
package deploy

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/types"

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

type KubeEndpoint struct {
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
}

func (m *KubeEndpoint) Reset()                    { *m = KubeEndpoint{} }
func (m *KubeEndpoint) String() string            { return proto.CompactTextString(m) }
func (*KubeEndpoint) ProtoMessage()               {}
func (*KubeEndpoint) Descriptor() ([]byte, []int) { return fileDescriptorDeploy, []int{0} }

func (m *KubeEndpoint) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

type Cluster struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *Cluster) Reset()                    { *m = Cluster{} }
func (m *Cluster) String() string            { return proto.CompactTextString(m) }
func (*Cluster) ProtoMessage()               {}
func (*Cluster) Descriptor() ([]byte, []int) { return fileDescriptorDeploy, []int{1} }

func (m *Cluster) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ClusterInfo struct {
	Cluster *Cluster `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Shards  uint64   `protobuf:"varint,2,opt,name=shards,proto3" json:"shards,omitempty"`
}

func (m *ClusterInfo) Reset()                    { *m = ClusterInfo{} }
func (m *ClusterInfo) String() string            { return proto.CompactTextString(m) }
func (*ClusterInfo) ProtoMessage()               {}
func (*ClusterInfo) Descriptor() ([]byte, []int) { return fileDescriptorDeploy, []int{2} }

func (m *ClusterInfo) GetCluster() *Cluster {
	if m != nil {
		return m.Cluster
	}
	return nil
}

func (m *ClusterInfo) GetShards() uint64 {
	if m != nil {
		return m.Shards
	}
	return 0
}

type ClusterInfos struct {
	ClusterInfos []*ClusterInfo `protobuf:"bytes,1,rep,name=cluster_infos,json=clusterInfos" json:"cluster_infos,omitempty"`
}

func (m *ClusterInfos) Reset()                    { *m = ClusterInfos{} }
func (m *ClusterInfos) String() string            { return proto.CompactTextString(m) }
func (*ClusterInfos) ProtoMessage()               {}
func (*ClusterInfos) Descriptor() ([]byte, []int) { return fileDescriptorDeploy, []int{3} }

func (m *ClusterInfos) GetClusterInfos() []*ClusterInfo {
	if m != nil {
		return m.ClusterInfos
	}
	return nil
}

type CreateClusterRequest struct {
	Cluster *Cluster `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Shards  uint64   `protobuf:"varint,2,opt,name=shards,proto3" json:"shards,omitempty"`
}

func (m *CreateClusterRequest) Reset()                    { *m = CreateClusterRequest{} }
func (m *CreateClusterRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateClusterRequest) ProtoMessage()               {}
func (*CreateClusterRequest) Descriptor() ([]byte, []int) { return fileDescriptorDeploy, []int{4} }

func (m *CreateClusterRequest) GetCluster() *Cluster {
	if m != nil {
		return m.Cluster
	}
	return nil
}

func (m *CreateClusterRequest) GetShards() uint64 {
	if m != nil {
		return m.Shards
	}
	return 0
}

type UpdateClusterRequest struct {
	Cluster *Cluster `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Nodes   uint64   `protobuf:"varint,2,opt,name=nodes,proto3" json:"nodes,omitempty"`
}

func (m *UpdateClusterRequest) Reset()                    { *m = UpdateClusterRequest{} }
func (m *UpdateClusterRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateClusterRequest) ProtoMessage()               {}
func (*UpdateClusterRequest) Descriptor() ([]byte, []int) { return fileDescriptorDeploy, []int{5} }

func (m *UpdateClusterRequest) GetCluster() *Cluster {
	if m != nil {
		return m.Cluster
	}
	return nil
}

func (m *UpdateClusterRequest) GetNodes() uint64 {
	if m != nil {
		return m.Nodes
	}
	return 0
}

type InspectClusterRequest struct {
	Cluster *Cluster `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
}

func (m *InspectClusterRequest) Reset()                    { *m = InspectClusterRequest{} }
func (m *InspectClusterRequest) String() string            { return proto.CompactTextString(m) }
func (*InspectClusterRequest) ProtoMessage()               {}
func (*InspectClusterRequest) Descriptor() ([]byte, []int) { return fileDescriptorDeploy, []int{6} }

func (m *InspectClusterRequest) GetCluster() *Cluster {
	if m != nil {
		return m.Cluster
	}
	return nil
}

type ListClusterRequest struct {
}

func (m *ListClusterRequest) Reset()                    { *m = ListClusterRequest{} }
func (m *ListClusterRequest) String() string            { return proto.CompactTextString(m) }
func (*ListClusterRequest) ProtoMessage()               {}
func (*ListClusterRequest) Descriptor() ([]byte, []int) { return fileDescriptorDeploy, []int{7} }

type DeleteClusterRequest struct {
	Cluster *Cluster `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
}

func (m *DeleteClusterRequest) Reset()                    { *m = DeleteClusterRequest{} }
func (m *DeleteClusterRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteClusterRequest) ProtoMessage()               {}
func (*DeleteClusterRequest) Descriptor() ([]byte, []int) { return fileDescriptorDeploy, []int{8} }

func (m *DeleteClusterRequest) GetCluster() *Cluster {
	if m != nil {
		return m.Cluster
	}
	return nil
}

func init() {
	proto.RegisterType((*KubeEndpoint)(nil), "deploy.KubeEndpoint")
	proto.RegisterType((*Cluster)(nil), "deploy.Cluster")
	proto.RegisterType((*ClusterInfo)(nil), "deploy.ClusterInfo")
	proto.RegisterType((*ClusterInfos)(nil), "deploy.ClusterInfos")
	proto.RegisterType((*CreateClusterRequest)(nil), "deploy.CreateClusterRequest")
	proto.RegisterType((*UpdateClusterRequest)(nil), "deploy.UpdateClusterRequest")
	proto.RegisterType((*InspectClusterRequest)(nil), "deploy.InspectClusterRequest")
	proto.RegisterType((*ListClusterRequest)(nil), "deploy.ListClusterRequest")
	proto.RegisterType((*DeleteClusterRequest)(nil), "deploy.DeleteClusterRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for API service

type APIClient interface {
	CreateCluster(ctx context.Context, in *CreateClusterRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	UpdateCluster(ctx context.Context, in *UpdateClusterRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	InspectCluster(ctx context.Context, in *InspectClusterRequest, opts ...grpc.CallOption) (*ClusterInfo, error)
	ListCluster(ctx context.Context, in *ListClusterRequest, opts ...grpc.CallOption) (*ClusterInfos, error)
	DeleteCluster(ctx context.Context, in *DeleteClusterRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) CreateCluster(ctx context.Context, in *CreateClusterRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/deploy.API/CreateCluster", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) UpdateCluster(ctx context.Context, in *UpdateClusterRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/deploy.API/UpdateCluster", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) InspectCluster(ctx context.Context, in *InspectClusterRequest, opts ...grpc.CallOption) (*ClusterInfo, error) {
	out := new(ClusterInfo)
	err := grpc.Invoke(ctx, "/deploy.API/InspectCluster", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListCluster(ctx context.Context, in *ListClusterRequest, opts ...grpc.CallOption) (*ClusterInfos, error) {
	out := new(ClusterInfos)
	err := grpc.Invoke(ctx, "/deploy.API/ListCluster", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) DeleteCluster(ctx context.Context, in *DeleteClusterRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/deploy.API/DeleteCluster", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for API service

type APIServer interface {
	CreateCluster(context.Context, *CreateClusterRequest) (*google_protobuf.Empty, error)
	UpdateCluster(context.Context, *UpdateClusterRequest) (*google_protobuf.Empty, error)
	InspectCluster(context.Context, *InspectClusterRequest) (*ClusterInfo, error)
	ListCluster(context.Context, *ListClusterRequest) (*ClusterInfos, error)
	DeleteCluster(context.Context, *DeleteClusterRequest) (*google_protobuf.Empty, error)
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_CreateCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).CreateCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deploy.API/CreateCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).CreateCluster(ctx, req.(*CreateClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_UpdateCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).UpdateCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deploy.API/UpdateCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).UpdateCluster(ctx, req.(*UpdateClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_InspectCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InspectClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).InspectCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deploy.API/InspectCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).InspectCluster(ctx, req.(*InspectClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_ListCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).ListCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deploy.API/ListCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).ListCluster(ctx, req.(*ListClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_DeleteCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).DeleteCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deploy.API/DeleteCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).DeleteCluster(ctx, req.(*DeleteClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "deploy.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCluster",
			Handler:    _API_CreateCluster_Handler,
		},
		{
			MethodName: "UpdateCluster",
			Handler:    _API_UpdateCluster_Handler,
		},
		{
			MethodName: "InspectCluster",
			Handler:    _API_InspectCluster_Handler,
		},
		{
			MethodName: "ListCluster",
			Handler:    _API_ListCluster_Handler,
		},
		{
			MethodName: "DeleteCluster",
			Handler:    _API_DeleteCluster_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/pkg/deploy/deploy.proto",
}

func init() { proto.RegisterFile("server/pkg/deploy/deploy.proto", fileDescriptorDeploy) }

var fileDescriptorDeploy = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x92, 0x4f, 0x4f, 0xf2, 0x40,
	0x10, 0xc6, 0x29, 0xf0, 0x42, 0xde, 0x29, 0x68, 0xb2, 0x56, 0x42, 0xaa, 0x18, 0xb2, 0x27, 0xbc,
	0xb4, 0x09, 0x5e, 0xbc, 0x22, 0xa2, 0x12, 0x3d, 0x90, 0x26, 0xc6, 0x78, 0x32, 0x40, 0x87, 0x3f,
	0xb1, 0x74, 0x6b, 0x77, 0x6b, 0xe2, 0x67, 0xf6, 0x4b, 0x98, 0xb6, 0xbb, 0x42, 0xb1, 0x7a, 0x10,
	0x4f, 0xdd, 0x99, 0x79, 0xe6, 0xd7, 0x99, 0x27, 0x03, 0x27, 0x1c, 0xc3, 0x57, 0x0c, 0xed, 0xe0,
	0x79, 0x6e, 0xbb, 0x18, 0x78, 0xec, 0x4d, 0x7e, 0xac, 0x20, 0x64, 0x82, 0x91, 0x4a, 0x1a, 0x99,
	0x47, 0x73, 0xc6, 0xe6, 0x1e, 0xda, 0x49, 0x76, 0x12, 0xcd, 0x6c, 0x5c, 0x05, 0x42, 0x8a, 0x28,
	0x85, 0xda, 0x6d, 0x34, 0xc1, 0x81, 0xef, 0x06, 0x6c, 0xe9, 0x0b, 0x42, 0xa0, 0xbc, 0x60, 0x5c,
	0x34, 0xb5, 0xb6, 0xd6, 0xf9, 0xef, 0x24, 0x6f, 0xda, 0x82, 0x6a, 0xdf, 0x8b, 0xb8, 0xc0, 0x30,
	0x2e, 0xfb, 0xe3, 0x15, 0xaa, 0x72, 0xfc, 0xa6, 0x23, 0xd0, 0x65, 0x79, 0xe8, 0xcf, 0x18, 0x39,
	0x85, 0xea, 0x34, 0x0d, 0x13, 0x95, 0xde, 0xdd, 0xb7, 0xe4, 0x58, 0x52, 0xe5, 0xa8, 0x3a, 0x69,
	0x40, 0x85, 0x2f, 0xc6, 0xa1, 0xcb, 0x9b, 0xc5, 0xb6, 0xd6, 0x29, 0x3b, 0x32, 0xa2, 0x37, 0x50,
	0xdb, 0x20, 0x72, 0x72, 0x0e, 0x75, 0xd9, 0xf2, 0xb4, 0x8c, 0x13, 0x4d, 0xad, 0x5d, 0xea, 0xe8,
	0xdd, 0x83, 0x2d, 0x70, 0x2c, 0x76, 0x6a, 0xd3, 0x8d, 0x4e, 0xfa, 0x08, 0x46, 0x3f, 0xc4, 0xb1,
	0x40, 0xf5, 0x6f, 0x7c, 0x89, 0x90, 0x8b, 0xbf, 0x18, 0xf2, 0x01, 0x8c, 0xfb, 0xc0, 0xdd, 0x09,
	0x6d, 0xc0, 0x3f, 0x9f, 0xb9, 0xa8, 0xc8, 0x69, 0x40, 0x2f, 0xe0, 0x70, 0xe8, 0xf3, 0x00, 0xa7,
	0xe2, 0xd7, 0x64, 0x6a, 0x00, 0xb9, 0x5b, 0xf2, 0x2d, 0x00, 0xed, 0x81, 0x71, 0x89, 0x1e, 0xee,
	0x30, 0x72, 0xf7, 0xbd, 0x08, 0xa5, 0xde, 0x68, 0x48, 0xae, 0xa1, 0x9e, 0x31, 0x96, 0x1c, 0x7f,
	0xb6, 0xe4, 0xf8, 0x6d, 0x36, 0xac, 0xf4, 0x08, 0x2d, 0x75, 0x84, 0xd6, 0x20, 0x3e, 0x42, 0x5a,
	0x88, 0x41, 0x19, 0x1b, 0xd7, 0xa0, 0x3c, 0x77, 0x7f, 0x00, 0x5d, 0xc1, 0x5e, 0xd6, 0x36, 0xd2,
	0x52, 0xa4, 0x5c, 0x3b, 0xcd, 0xbc, 0xf3, 0xa1, 0x05, 0xd2, 0x03, 0x7d, 0xc3, 0x3a, 0x62, 0x2a,
	0xd5, 0x57, 0x3f, 0x4d, 0x23, 0x87, 0xc0, 0xd3, 0x9d, 0x32, 0x3e, 0xaf, 0x77, 0xca, 0xb3, 0xff,
	0xfb, 0x9d, 0x26, 0x95, 0x24, 0x73, 0xf6, 0x11, 0x00, 0x00, 0xff, 0xff, 0xa5, 0xaf, 0xcd, 0x04,
	0xeb, 0x03, 0x00, 0x00,
}
