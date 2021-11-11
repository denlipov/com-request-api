// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package com_request_api

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

// ComRequestApiServiceClient is the client API for ComRequestApiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ComRequestApiServiceClient interface {
	// DescribeRequestV1 - Describe a request
	DescribeRequestV1(ctx context.Context, in *DescribeRequestV1Request, opts ...grpc.CallOption) (*DescribeRequestV1Response, error)
	CreateRequestV1(ctx context.Context, in *CreateRequestV1Request, opts ...grpc.CallOption) (*CreateRequestV1Response, error)
	ListRequestV1(ctx context.Context, in *ListRequestV1Request, opts ...grpc.CallOption) (*ListRequestV1Response, error)
	RemoveRequestV1(ctx context.Context, in *RemoveRequestV1Request, opts ...grpc.CallOption) (*RemoveRequestV1Response, error)
}

type comRequestApiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewComRequestApiServiceClient(cc grpc.ClientConnInterface) ComRequestApiServiceClient {
	return &comRequestApiServiceClient{cc}
}

func (c *comRequestApiServiceClient) DescribeRequestV1(ctx context.Context, in *DescribeRequestV1Request, opts ...grpc.CallOption) (*DescribeRequestV1Response, error) {
	out := new(DescribeRequestV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.com_request_api.v1.ComRequestApiService/DescribeRequestV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comRequestApiServiceClient) CreateRequestV1(ctx context.Context, in *CreateRequestV1Request, opts ...grpc.CallOption) (*CreateRequestV1Response, error) {
	out := new(CreateRequestV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.com_request_api.v1.ComRequestApiService/CreateRequestV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comRequestApiServiceClient) ListRequestV1(ctx context.Context, in *ListRequestV1Request, opts ...grpc.CallOption) (*ListRequestV1Response, error) {
	out := new(ListRequestV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.com_request_api.v1.ComRequestApiService/ListRequestV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comRequestApiServiceClient) RemoveRequestV1(ctx context.Context, in *RemoveRequestV1Request, opts ...grpc.CallOption) (*RemoveRequestV1Response, error) {
	out := new(RemoveRequestV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.com_request_api.v1.ComRequestApiService/RemoveRequestV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComRequestApiServiceServer is the server API for ComRequestApiService service.
// All implementations must embed UnimplementedComRequestApiServiceServer
// for forward compatibility
type ComRequestApiServiceServer interface {
	// DescribeRequestV1 - Describe a request
	DescribeRequestV1(context.Context, *DescribeRequestV1Request) (*DescribeRequestV1Response, error)
	CreateRequestV1(context.Context, *CreateRequestV1Request) (*CreateRequestV1Response, error)
	ListRequestV1(context.Context, *ListRequestV1Request) (*ListRequestV1Response, error)
	RemoveRequestV1(context.Context, *RemoveRequestV1Request) (*RemoveRequestV1Response, error)
	mustEmbedUnimplementedComRequestApiServiceServer()
}

// UnimplementedComRequestApiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedComRequestApiServiceServer struct {
}

func (UnimplementedComRequestApiServiceServer) DescribeRequestV1(context.Context, *DescribeRequestV1Request) (*DescribeRequestV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeRequestV1 not implemented")
}
func (UnimplementedComRequestApiServiceServer) CreateRequestV1(context.Context, *CreateRequestV1Request) (*CreateRequestV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRequestV1 not implemented")
}
func (UnimplementedComRequestApiServiceServer) ListRequestV1(context.Context, *ListRequestV1Request) (*ListRequestV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRequestV1 not implemented")
}
func (UnimplementedComRequestApiServiceServer) RemoveRequestV1(context.Context, *RemoveRequestV1Request) (*RemoveRequestV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveRequestV1 not implemented")
}
func (UnimplementedComRequestApiServiceServer) mustEmbedUnimplementedComRequestApiServiceServer() {}

// UnsafeComRequestApiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ComRequestApiServiceServer will
// result in compilation errors.
type UnsafeComRequestApiServiceServer interface {
	mustEmbedUnimplementedComRequestApiServiceServer()
}

func RegisterComRequestApiServiceServer(s grpc.ServiceRegistrar, srv ComRequestApiServiceServer) {
	s.RegisterService(&ComRequestApiService_ServiceDesc, srv)
}

func _ComRequestApiService_DescribeRequestV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeRequestV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComRequestApiServiceServer).DescribeRequestV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.com_request_api.v1.ComRequestApiService/DescribeRequestV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComRequestApiServiceServer).DescribeRequestV1(ctx, req.(*DescribeRequestV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComRequestApiService_CreateRequestV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequestV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComRequestApiServiceServer).CreateRequestV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.com_request_api.v1.ComRequestApiService/CreateRequestV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComRequestApiServiceServer).CreateRequestV1(ctx, req.(*CreateRequestV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComRequestApiService_ListRequestV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequestV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComRequestApiServiceServer).ListRequestV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.com_request_api.v1.ComRequestApiService/ListRequestV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComRequestApiServiceServer).ListRequestV1(ctx, req.(*ListRequestV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComRequestApiService_RemoveRequestV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRequestV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComRequestApiServiceServer).RemoveRequestV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.com_request_api.v1.ComRequestApiService/RemoveRequestV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComRequestApiServiceServer).RemoveRequestV1(ctx, req.(*RemoveRequestV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// ComRequestApiService_ServiceDesc is the grpc.ServiceDesc for ComRequestApiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ComRequestApiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ozonmp.com_request_api.v1.ComRequestApiService",
	HandlerType: (*ComRequestApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribeRequestV1",
			Handler:    _ComRequestApiService_DescribeRequestV1_Handler,
		},
		{
			MethodName: "CreateRequestV1",
			Handler:    _ComRequestApiService_CreateRequestV1_Handler,
		},
		{
			MethodName: "ListRequestV1",
			Handler:    _ComRequestApiService_ListRequestV1_Handler,
		},
		{
			MethodName: "RemoveRequestV1",
			Handler:    _ComRequestApiService_RemoveRequestV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ozonmp/com_request_api/v1/com_request_api.proto",
}
