// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: proto/fetcher/v1/trades.proto

package providers

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

// FetcherServiceClient is the client API for FetcherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FetcherServiceClient interface {
	FetchTrades(ctx context.Context, in *FetchTradesRequest, opts ...grpc.CallOption) (*FetchTradesResponse, error)
	StopFetchTrades(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StopFetchTradesResponse, error)
}

type fetcherServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFetcherServiceClient(cc grpc.ClientConnInterface) FetcherServiceClient {
	return &fetcherServiceClient{cc}
}

func (c *fetcherServiceClient) FetchTrades(ctx context.Context, in *FetchTradesRequest, opts ...grpc.CallOption) (*FetchTradesResponse, error) {
	out := new(FetchTradesResponse)
	err := c.cc.Invoke(ctx, "/fetcher.v1.FetcherService/FetchTrades", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fetcherServiceClient) StopFetchTrades(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StopFetchTradesResponse, error) {
	out := new(StopFetchTradesResponse)
	err := c.cc.Invoke(ctx, "/fetcher.v1.FetcherService/StopFetchTrades", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FetcherServiceServer is the server API for FetcherService service.
// All implementations must embed UnimplementedFetcherServiceServer
// for forward compatibility
type FetcherServiceServer interface {
	FetchTrades(context.Context, *FetchTradesRequest) (*FetchTradesResponse, error)
	StopFetchTrades(context.Context, *Empty) (*StopFetchTradesResponse, error)
	mustEmbedUnimplementedFetcherServiceServer()
}

// UnimplementedFetcherServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFetcherServiceServer struct {
}

func (UnimplementedFetcherServiceServer) FetchTrades(context.Context, *FetchTradesRequest) (*FetchTradesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchTrades not implemented")
}
func (UnimplementedFetcherServiceServer) StopFetchTrades(context.Context, *Empty) (*StopFetchTradesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopFetchTrades not implemented")
}
func (UnimplementedFetcherServiceServer) mustEmbedUnimplementedFetcherServiceServer() {}

// UnsafeFetcherServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FetcherServiceServer will
// result in compilation errors.
type UnsafeFetcherServiceServer interface {
	mustEmbedUnimplementedFetcherServiceServer()
}

func RegisterFetcherServiceServer(s grpc.ServiceRegistrar, srv FetcherServiceServer) {
	s.RegisterService(&FetcherService_ServiceDesc, srv)
}

func _FetcherService_FetchTrades_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchTradesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FetcherServiceServer).FetchTrades(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fetcher.v1.FetcherService/FetchTrades",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FetcherServiceServer).FetchTrades(ctx, req.(*FetchTradesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FetcherService_StopFetchTrades_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FetcherServiceServer).StopFetchTrades(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fetcher.v1.FetcherService/StopFetchTrades",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FetcherServiceServer).StopFetchTrades(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// FetcherService_ServiceDesc is the grpc.ServiceDesc for FetcherService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FetcherService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fetcher.v1.FetcherService",
	HandlerType: (*FetcherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchTrades",
			Handler:    _FetcherService_FetchTrades_Handler,
		},
		{
			MethodName: "StopFetchTrades",
			Handler:    _FetcherService_StopFetchTrades_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/fetcher/v1/trades.proto",
}
