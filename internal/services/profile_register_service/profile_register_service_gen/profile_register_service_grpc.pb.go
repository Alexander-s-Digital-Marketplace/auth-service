// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.2
// source: profile_register_service.proto

package profile_register_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ProfileRegisterService_ProfileRegister_FullMethodName = "/auth.ProfileRegisterService/ProfileRegister"
)

// ProfileRegisterServiceClient is the client API for ProfileRegisterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProfileRegisterServiceClient interface {
	ProfileRegister(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type profileRegisterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProfileRegisterServiceClient(cc grpc.ClientConnInterface) ProfileRegisterServiceClient {
	return &profileRegisterServiceClient{cc}
}

func (c *profileRegisterServiceClient) ProfileRegister(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, ProfileRegisterService_ProfileRegister_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileRegisterServiceServer is the server API for ProfileRegisterService service.
// All implementations must embed UnimplementedProfileRegisterServiceServer
// for forward compatibility.
type ProfileRegisterServiceServer interface {
	ProfileRegister(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedProfileRegisterServiceServer()
}

// UnimplementedProfileRegisterServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProfileRegisterServiceServer struct{}

func (UnimplementedProfileRegisterServiceServer) ProfileRegister(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProfileRegister not implemented")
}
func (UnimplementedProfileRegisterServiceServer) mustEmbedUnimplementedProfileRegisterServiceServer() {
}
func (UnimplementedProfileRegisterServiceServer) testEmbeddedByValue() {}

// UnsafeProfileRegisterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfileRegisterServiceServer will
// result in compilation errors.
type UnsafeProfileRegisterServiceServer interface {
	mustEmbedUnimplementedProfileRegisterServiceServer()
}

func RegisterProfileRegisterServiceServer(s grpc.ServiceRegistrar, srv ProfileRegisterServiceServer) {
	// If the following call pancis, it indicates UnimplementedProfileRegisterServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ProfileRegisterService_ServiceDesc, srv)
}

func _ProfileRegisterService_ProfileRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileRegisterServiceServer).ProfileRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProfileRegisterService_ProfileRegister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileRegisterServiceServer).ProfileRegister(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// ProfileRegisterService_ServiceDesc is the grpc.ServiceDesc for ProfileRegisterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProfileRegisterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.ProfileRegisterService",
	HandlerType: (*ProfileRegisterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProfileRegister",
			Handler:    _ProfileRegisterService_ProfileRegister_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "profile_register_service.proto",
}
