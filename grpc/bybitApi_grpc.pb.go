// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.6
// source: bybitApi.proto

package grpc

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
	BybitService_FuturesTransactions_FullMethodName = "/grpc.BybitService/FuturesTransactions"
	BybitService_SpotAccountInfo_FullMethodName     = "/grpc.BybitService/SpotAccountInfo"
)

// BybitServiceClient is the client API for BybitService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BybitServiceClient interface {
	FuturesTransactions(ctx context.Context, in *FuturesTransactionsRequest, opts ...grpc.CallOption) (*TransactionResult, error)
	SpotAccountInfo(ctx context.Context, in *SpotAccountRequest, opts ...grpc.CallOption) (*SpotAccountResponse, error)
}

type bybitServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBybitServiceClient(cc grpc.ClientConnInterface) BybitServiceClient {
	return &bybitServiceClient{cc}
}

func (c *bybitServiceClient) FuturesTransactions(ctx context.Context, in *FuturesTransactionsRequest, opts ...grpc.CallOption) (*TransactionResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TransactionResult)
	err := c.cc.Invoke(ctx, BybitService_FuturesTransactions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bybitServiceClient) SpotAccountInfo(ctx context.Context, in *SpotAccountRequest, opts ...grpc.CallOption) (*SpotAccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SpotAccountResponse)
	err := c.cc.Invoke(ctx, BybitService_SpotAccountInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BybitServiceServer is the server API for BybitService service.
// All implementations must embed UnimplementedBybitServiceServer
// for forward compatibility.
type BybitServiceServer interface {
	FuturesTransactions(context.Context, *FuturesTransactionsRequest) (*TransactionResult, error)
	SpotAccountInfo(context.Context, *SpotAccountRequest) (*SpotAccountResponse, error)
	mustEmbedUnimplementedBybitServiceServer()
}

// UnimplementedBybitServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBybitServiceServer struct{}

func (UnimplementedBybitServiceServer) FuturesTransactions(context.Context, *FuturesTransactionsRequest) (*TransactionResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FuturesTransactions not implemented")
}
func (UnimplementedBybitServiceServer) SpotAccountInfo(context.Context, *SpotAccountRequest) (*SpotAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SpotAccountInfo not implemented")
}
func (UnimplementedBybitServiceServer) mustEmbedUnimplementedBybitServiceServer() {}
func (UnimplementedBybitServiceServer) testEmbeddedByValue()                      {}

// UnsafeBybitServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BybitServiceServer will
// result in compilation errors.
type UnsafeBybitServiceServer interface {
	mustEmbedUnimplementedBybitServiceServer()
}

func RegisterBybitServiceServer(s grpc.ServiceRegistrar, srv BybitServiceServer) {
	// If the following call pancis, it indicates UnimplementedBybitServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BybitService_ServiceDesc, srv)
}

func _BybitService_FuturesTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FuturesTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BybitServiceServer).FuturesTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BybitService_FuturesTransactions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BybitServiceServer).FuturesTransactions(ctx, req.(*FuturesTransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BybitService_SpotAccountInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpotAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BybitServiceServer).SpotAccountInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BybitService_SpotAccountInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BybitServiceServer).SpotAccountInfo(ctx, req.(*SpotAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BybitService_ServiceDesc is the grpc.ServiceDesc for BybitService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BybitService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.BybitService",
	HandlerType: (*BybitServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FuturesTransactions",
			Handler:    _BybitService_FuturesTransactions_Handler,
		},
		{
			MethodName: "SpotAccountInfo",
			Handler:    _BybitService_SpotAccountInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bybitApi.proto",
}
