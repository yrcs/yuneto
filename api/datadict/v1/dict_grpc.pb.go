// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: v1/dict.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DataDictClient is the client API for DataDict service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataDictClient interface {
	ListChildren(ctx context.Context, in *ListChildrenRequest, opts ...grpc.CallOption) (*ListChildrenReply, error)
}

type dataDictClient struct {
	cc grpc.ClientConnInterface
}

func NewDataDictClient(cc grpc.ClientConnInterface) DataDictClient {
	return &dataDictClient{cc}
}

func (c *dataDictClient) ListChildren(ctx context.Context, in *ListChildrenRequest, opts ...grpc.CallOption) (*ListChildrenReply, error) {
	out := new(ListChildrenReply)
	err := c.cc.Invoke(ctx, "/dict.v1.DataDict/ListChildren", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataDictServer is the server API for DataDict service.
// All implementations must embed UnimplementedDataDictServer
// for forward compatibility
type DataDictServer interface {
	ListChildren(context.Context, *ListChildrenRequest) (*ListChildrenReply, error)
	mustEmbedUnimplementedDataDictServer()
}

// UnimplementedDataDictServer must be embedded to have forward compatible implementations.
type UnimplementedDataDictServer struct {
}

func (UnimplementedDataDictServer) ListChildren(context.Context, *ListChildrenRequest) (*ListChildrenReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListChildren not implemented")
}
func (UnimplementedDataDictServer) mustEmbedUnimplementedDataDictServer() {}

// UnsafeDataDictServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataDictServer will
// result in compilation errors.
type UnsafeDataDictServer interface {
	mustEmbedUnimplementedDataDictServer()
}

func RegisterDataDictServer(s grpc.ServiceRegistrar, srv DataDictServer) {
	s.RegisterService(&DataDict_ServiceDesc, srv)
}

func _DataDict_ListChildren_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListChildrenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataDictServer).ListChildren(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dict.v1.DataDict/ListChildren",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataDictServer).ListChildren(ctx, req.(*ListChildrenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DataDict_ServiceDesc is the grpc.ServiceDesc for DataDict service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataDict_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dict.v1.DataDict",
	HandlerType: (*DataDictServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListChildren",
			Handler:    _DataDict_ListChildren_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/dict.proto",
}