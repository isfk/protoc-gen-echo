// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: example.proto

package example

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

const (
	ExampleService_Hello_FullMethodName = "/example.ExampleService/Hello"
	ExampleService_Say_FullMethodName   = "/example.ExampleService/Say"
)

// ExampleServiceClient is the client API for ExampleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExampleServiceClient interface {
	// hello
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	// Say ......
	// Say godoc
	// @Summary Say
	// @Description say some thing
	// @Tags      say
	// @Accept    json
	// @Produce   json
	// @Param     args  SayRequest
	// @Success   200   {array}   SayResponse
	// @Failure   400   {object}  httputil.HTTPError
	// @Failure   404   {object}  httputil.HTTPError
	// @Failure   500   {object}  httputil.HTTPError
	// @Router    /say  [get]
	Say(ctx context.Context, in *SayRequest, opts ...grpc.CallOption) (*SayResponse, error)
}

type exampleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExampleServiceClient(cc grpc.ClientConnInterface) ExampleServiceClient {
	return &exampleServiceClient{cc}
}

func (c *exampleServiceClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, ExampleService_Hello_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleServiceClient) Say(ctx context.Context, in *SayRequest, opts ...grpc.CallOption) (*SayResponse, error) {
	out := new(SayResponse)
	err := c.cc.Invoke(ctx, ExampleService_Say_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExampleServiceServer is the server API for ExampleService service.
// All implementations should embed UnimplementedExampleServiceServer
// for forward compatibility
type ExampleServiceServer interface {
	// hello
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	// Say ......
	// Say godoc
	// @Summary Say
	// @Description say some thing
	// @Tags      say
	// @Accept    json
	// @Produce   json
	// @Param     args  SayRequest
	// @Success   200   {array}   SayResponse
	// @Failure   400   {object}  httputil.HTTPError
	// @Failure   404   {object}  httputil.HTTPError
	// @Failure   500   {object}  httputil.HTTPError
	// @Router    /say  [get]
	Say(context.Context, *SayRequest) (*SayResponse, error)
}

// UnimplementedExampleServiceServer should be embedded to have forward compatible implementations.
type UnimplementedExampleServiceServer struct {
}

func (UnimplementedExampleServiceServer) Hello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedExampleServiceServer) Say(context.Context, *SayRequest) (*SayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Say not implemented")
}

// UnsafeExampleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExampleServiceServer will
// result in compilation errors.
type UnsafeExampleServiceServer interface {
	mustEmbedUnimplementedExampleServiceServer()
}

func RegisterExampleServiceServer(s grpc.ServiceRegistrar, srv ExampleServiceServer) {
	s.RegisterService(&ExampleService_ServiceDesc, srv)
}

func _ExampleService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExampleService_Hello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServiceServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleService_Say_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServiceServer).Say(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExampleService_Say_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServiceServer).Say(ctx, req.(*SayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExampleService_ServiceDesc is the grpc.ServiceDesc for ExampleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExampleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "example.ExampleService",
	HandlerType: (*ExampleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _ExampleService_Hello_Handler,
		},
		{
			MethodName: "Say",
			Handler:    _ExampleService_Say_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example.proto",
}
