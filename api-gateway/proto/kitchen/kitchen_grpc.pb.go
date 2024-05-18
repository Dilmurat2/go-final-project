// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: proto/kitchen/kitchen.proto

package kitchen

import (
	order "api-gateway/proto/order"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// KitchenServiceClient is the client API for KitchenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KitchenServiceClient interface {
	ProcessOrder(ctx context.Context, in *order.Order, opts ...grpc.CallOption) (*ProcessOrderResponse, error)
	ChangeOrderStatus(ctx context.Context, in *order.ChangeOrderStatusRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type kitchenServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKitchenServiceClient(cc grpc.ClientConnInterface) KitchenServiceClient {
	return &kitchenServiceClient{cc}
}

func (c *kitchenServiceClient) ProcessOrder(ctx context.Context, in *order.Order, opts ...grpc.CallOption) (*ProcessOrderResponse, error) {
	out := new(ProcessOrderResponse)
	err := c.cc.Invoke(ctx, "/kitchen.KitchenService/ProcessOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kitchenServiceClient) ChangeOrderStatus(ctx context.Context, in *order.ChangeOrderStatusRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/kitchen.KitchenService/ChangeOrderStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KitchenServiceServer is the server API for KitchenService service.
// All implementations must embed UnimplementedKitchenServiceServer
// for forward compatibility
type KitchenServiceServer interface {
	ProcessOrder(context.Context, *order.Order) (*ProcessOrderResponse, error)
	ChangeOrderStatus(context.Context, *order.ChangeOrderStatusRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedKitchenServiceServer()
}

// UnimplementedKitchenServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKitchenServiceServer struct {
}

func (UnimplementedKitchenServiceServer) ProcessOrder(context.Context, *order.Order) (*ProcessOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessOrder not implemented")
}
func (UnimplementedKitchenServiceServer) ChangeOrderStatus(context.Context, *order.ChangeOrderStatusRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeOrderStatus not implemented")
}
func (UnimplementedKitchenServiceServer) mustEmbedUnimplementedKitchenServiceServer() {}

// UnsafeKitchenServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KitchenServiceServer will
// result in compilation errors.
type UnsafeKitchenServiceServer interface {
	mustEmbedUnimplementedKitchenServiceServer()
}

func RegisterKitchenServiceServer(s grpc.ServiceRegistrar, srv KitchenServiceServer) {
	s.RegisterService(&KitchenService_ServiceDesc, srv)
}

func _KitchenService_ProcessOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(order.Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KitchenServiceServer).ProcessOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kitchen.KitchenService/ProcessOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KitchenServiceServer).ProcessOrder(ctx, req.(*order.Order))
	}
	return interceptor(ctx, in, info, handler)
}

func _KitchenService_ChangeOrderStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(order.ChangeOrderStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KitchenServiceServer).ChangeOrderStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kitchen.KitchenService/ChangeOrderStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KitchenServiceServer).ChangeOrderStatus(ctx, req.(*order.ChangeOrderStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KitchenService_ServiceDesc is the grpc.ServiceDesc for KitchenService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KitchenService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kitchen.KitchenService",
	HandlerType: (*KitchenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessOrder",
			Handler:    _KitchenService_ProcessOrder_Handler,
		},
		{
			MethodName: "ChangeOrderStatus",
			Handler:    _KitchenService_ChangeOrderStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/kitchen/kitchen.proto",
}
