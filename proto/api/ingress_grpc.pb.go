// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: proto/api/ingress.proto

package api

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

// IngressPluginClient is the client API for IngressPlugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IngressPluginClient interface {
	CreateOrChangeRoute(ctx context.Context, in *AllocID, opts ...grpc.CallOption) (*Response, error)
}

type ingressPluginClient struct {
	cc grpc.ClientConnInterface
}

func NewIngressPluginClient(cc grpc.ClientConnInterface) IngressPluginClient {
	return &ingressPluginClient{cc}
}

func (c *ingressPluginClient) CreateOrChangeRoute(ctx context.Context, in *AllocID, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.IngressPlugin/CreateOrChangeRoute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IngressPluginServer is the server API for IngressPlugin service.
// All implementations must embed UnimplementedIngressPluginServer
// for forward compatibility
type IngressPluginServer interface {
	CreateOrChangeRoute(context.Context, *AllocID) (*Response, error)
	mustEmbedUnimplementedIngressPluginServer()
}

// UnimplementedIngressPluginServer must be embedded to have forward compatible implementations.
type UnimplementedIngressPluginServer struct {
}

func (UnimplementedIngressPluginServer) CreateOrChangeRoute(context.Context, *AllocID) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrChangeRoute not implemented")
}
func (UnimplementedIngressPluginServer) mustEmbedUnimplementedIngressPluginServer() {}

// UnsafeIngressPluginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IngressPluginServer will
// result in compilation errors.
type UnsafeIngressPluginServer interface {
	mustEmbedUnimplementedIngressPluginServer()
}

func RegisterIngressPluginServer(s grpc.ServiceRegistrar, srv IngressPluginServer) {
	s.RegisterService(&IngressPlugin_ServiceDesc, srv)
}

func _IngressPlugin_CreateOrChangeRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllocID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IngressPluginServer).CreateOrChangeRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.IngressPlugin/CreateOrChangeRoute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IngressPluginServer).CreateOrChangeRoute(ctx, req.(*AllocID))
	}
	return interceptor(ctx, in, info, handler)
}

// IngressPlugin_ServiceDesc is the grpc.ServiceDesc for IngressPlugin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IngressPlugin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.IngressPlugin",
	HandlerType: (*IngressPluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrChangeRoute",
			Handler:    _IngressPlugin_CreateOrChangeRoute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/api/ingress.proto",
}
