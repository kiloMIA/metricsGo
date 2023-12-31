// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: buses.proto

package pb

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
	BusService_RequestBus_FullMethodName = "/buses.BusService/RequestBus"
)

// BusServiceClient is the client API for BusService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BusServiceClient interface {
	RequestBus(ctx context.Context, in *BusRequest, opts ...grpc.CallOption) (*BusResponse, error)
}

type busServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBusServiceClient(cc grpc.ClientConnInterface) BusServiceClient {
	return &busServiceClient{cc}
}

func (c *busServiceClient) RequestBus(ctx context.Context, in *BusRequest, opts ...grpc.CallOption) (*BusResponse, error) {
	out := new(BusResponse)
	err := c.cc.Invoke(ctx, BusService_RequestBus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BusServiceServer is the server API for BusService service.
// All implementations must embed UnimplementedBusServiceServer
// for forward compatibility
type BusServiceServer interface {
	RequestBus(context.Context, *BusRequest) (*BusResponse, error)
	mustEmbedUnimplementedBusServiceServer()
}

// UnimplementedBusServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBusServiceServer struct {
}

func (UnimplementedBusServiceServer) RequestBus(context.Context, *BusRequest) (*BusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestBus not implemented")
}
func (UnimplementedBusServiceServer) mustEmbedUnimplementedBusServiceServer() {}

// UnsafeBusServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BusServiceServer will
// result in compilation errors.
type UnsafeBusServiceServer interface {
	mustEmbedUnimplementedBusServiceServer()
}

func RegisterBusServiceServer(s grpc.ServiceRegistrar, srv BusServiceServer) {
	s.RegisterService(&BusService_ServiceDesc, srv)
}

func _BusService_RequestBus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusServiceServer).RequestBus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BusService_RequestBus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusServiceServer).RequestBus(ctx, req.(*BusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BusService_ServiceDesc is the grpc.ServiceDesc for BusService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BusService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "buses.BusService",
	HandlerType: (*BusServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestBus",
			Handler:    _BusService_RequestBus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "buses.proto",
}
