// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: v1/hospital.proto

package v1

import (
	context "context"
	pagination "github.com/yrcs/yuneto/third_party/pagination"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// HospitalClient is the client API for Hospital service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HospitalClient interface {
	ListHospitalSetting(ctx context.Context, in *pagination.PagingRequest, opts ...grpc.CallOption) (*pagination.PagingReply, error)
	AddHospitalSetting(ctx context.Context, in *AddHospitalSettingRequest, opts ...grpc.CallOption) (*CommonAddReply, error)
	EditHospitalSetting(ctx context.Context, in *EditHospitalSettingRequest, opts ...grpc.CallOption) (*CommonEditReply, error)
	DeleteHospitalSetting(ctx context.Context, in *DeleteHospitalSettingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteHospitalSettings(ctx context.Context, in *DeleteHospitalSettingsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	LockHospitalSetting(ctx context.Context, in *LockHospitalSettingRequest, opts ...grpc.CallOption) (*CommonEditReply, error)
}

type hospitalClient struct {
	cc grpc.ClientConnInterface
}

func NewHospitalClient(cc grpc.ClientConnInterface) HospitalClient {
	return &hospitalClient{cc}
}

func (c *hospitalClient) ListHospitalSetting(ctx context.Context, in *pagination.PagingRequest, opts ...grpc.CallOption) (*pagination.PagingReply, error) {
	out := new(pagination.PagingReply)
	err := c.cc.Invoke(ctx, "/hospital.v1.Hospital/ListHospitalSetting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hospitalClient) AddHospitalSetting(ctx context.Context, in *AddHospitalSettingRequest, opts ...grpc.CallOption) (*CommonAddReply, error) {
	out := new(CommonAddReply)
	err := c.cc.Invoke(ctx, "/hospital.v1.Hospital/AddHospitalSetting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hospitalClient) EditHospitalSetting(ctx context.Context, in *EditHospitalSettingRequest, opts ...grpc.CallOption) (*CommonEditReply, error) {
	out := new(CommonEditReply)
	err := c.cc.Invoke(ctx, "/hospital.v1.Hospital/EditHospitalSetting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hospitalClient) DeleteHospitalSetting(ctx context.Context, in *DeleteHospitalSettingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/hospital.v1.Hospital/DeleteHospitalSetting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hospitalClient) DeleteHospitalSettings(ctx context.Context, in *DeleteHospitalSettingsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/hospital.v1.Hospital/DeleteHospitalSettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hospitalClient) LockHospitalSetting(ctx context.Context, in *LockHospitalSettingRequest, opts ...grpc.CallOption) (*CommonEditReply, error) {
	out := new(CommonEditReply)
	err := c.cc.Invoke(ctx, "/hospital.v1.Hospital/LockHospitalSetting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HospitalServer is the server API for Hospital service.
// All implementations must embed UnimplementedHospitalServer
// for forward compatibility
type HospitalServer interface {
	ListHospitalSetting(context.Context, *pagination.PagingRequest) (*pagination.PagingReply, error)
	AddHospitalSetting(context.Context, *AddHospitalSettingRequest) (*CommonAddReply, error)
	EditHospitalSetting(context.Context, *EditHospitalSettingRequest) (*CommonEditReply, error)
	DeleteHospitalSetting(context.Context, *DeleteHospitalSettingRequest) (*emptypb.Empty, error)
	DeleteHospitalSettings(context.Context, *DeleteHospitalSettingsRequest) (*emptypb.Empty, error)
	LockHospitalSetting(context.Context, *LockHospitalSettingRequest) (*CommonEditReply, error)
	mustEmbedUnimplementedHospitalServer()
}

// UnimplementedHospitalServer must be embedded to have forward compatible implementations.
type UnimplementedHospitalServer struct {
}

func (UnimplementedHospitalServer) ListHospitalSetting(context.Context, *pagination.PagingRequest) (*pagination.PagingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHospitalSetting not implemented")
}
func (UnimplementedHospitalServer) AddHospitalSetting(context.Context, *AddHospitalSettingRequest) (*CommonAddReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddHospitalSetting not implemented")
}
func (UnimplementedHospitalServer) EditHospitalSetting(context.Context, *EditHospitalSettingRequest) (*CommonEditReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditHospitalSetting not implemented")
}
func (UnimplementedHospitalServer) DeleteHospitalSetting(context.Context, *DeleteHospitalSettingRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHospitalSetting not implemented")
}
func (UnimplementedHospitalServer) DeleteHospitalSettings(context.Context, *DeleteHospitalSettingsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHospitalSettings not implemented")
}
func (UnimplementedHospitalServer) LockHospitalSetting(context.Context, *LockHospitalSettingRequest) (*CommonEditReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LockHospitalSetting not implemented")
}
func (UnimplementedHospitalServer) mustEmbedUnimplementedHospitalServer() {}

// UnsafeHospitalServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HospitalServer will
// result in compilation errors.
type UnsafeHospitalServer interface {
	mustEmbedUnimplementedHospitalServer()
}

func RegisterHospitalServer(s grpc.ServiceRegistrar, srv HospitalServer) {
	s.RegisterService(&Hospital_ServiceDesc, srv)
}

func _Hospital_ListHospitalSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pagination.PagingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HospitalServer).ListHospitalSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hospital.v1.Hospital/ListHospitalSetting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HospitalServer).ListHospitalSetting(ctx, req.(*pagination.PagingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hospital_AddHospitalSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddHospitalSettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HospitalServer).AddHospitalSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hospital.v1.Hospital/AddHospitalSetting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HospitalServer).AddHospitalSetting(ctx, req.(*AddHospitalSettingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hospital_EditHospitalSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditHospitalSettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HospitalServer).EditHospitalSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hospital.v1.Hospital/EditHospitalSetting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HospitalServer).EditHospitalSetting(ctx, req.(*EditHospitalSettingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hospital_DeleteHospitalSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteHospitalSettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HospitalServer).DeleteHospitalSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hospital.v1.Hospital/DeleteHospitalSetting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HospitalServer).DeleteHospitalSetting(ctx, req.(*DeleteHospitalSettingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hospital_DeleteHospitalSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteHospitalSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HospitalServer).DeleteHospitalSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hospital.v1.Hospital/DeleteHospitalSettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HospitalServer).DeleteHospitalSettings(ctx, req.(*DeleteHospitalSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hospital_LockHospitalSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LockHospitalSettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HospitalServer).LockHospitalSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hospital.v1.Hospital/LockHospitalSetting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HospitalServer).LockHospitalSetting(ctx, req.(*LockHospitalSettingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Hospital_ServiceDesc is the grpc.ServiceDesc for Hospital service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Hospital_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hospital.v1.Hospital",
	HandlerType: (*HospitalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListHospitalSetting",
			Handler:    _Hospital_ListHospitalSetting_Handler,
		},
		{
			MethodName: "AddHospitalSetting",
			Handler:    _Hospital_AddHospitalSetting_Handler,
		},
		{
			MethodName: "EditHospitalSetting",
			Handler:    _Hospital_EditHospitalSetting_Handler,
		},
		{
			MethodName: "DeleteHospitalSetting",
			Handler:    _Hospital_DeleteHospitalSetting_Handler,
		},
		{
			MethodName: "DeleteHospitalSettings",
			Handler:    _Hospital_DeleteHospitalSettings_Handler,
		},
		{
			MethodName: "LockHospitalSetting",
			Handler:    _Hospital_LockHospitalSetting_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/hospital.proto",
}
