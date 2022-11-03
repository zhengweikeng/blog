// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: order/order.proto

package order

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServiceClient interface {
	QueryOrders(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (OrderService_QueryOrdersClient, error)
	UpdateOrders(ctx context.Context, opts ...grpc.CallOption) (OrderService_UpdateOrdersClient, error)
	ProcessOrders(ctx context.Context, opts ...grpc.CallOption) (OrderService_ProcessOrdersClient, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) QueryOrders(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (OrderService_QueryOrdersClient, error) {
	stream, err := c.cc.NewStream(ctx, &OrderService_ServiceDesc.Streams[0], "/order.OrderService/queryOrders", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderServiceQueryOrdersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OrderService_QueryOrdersClient interface {
	Recv() (*Order, error)
	grpc.ClientStream
}

type orderServiceQueryOrdersClient struct {
	grpc.ClientStream
}

func (x *orderServiceQueryOrdersClient) Recv() (*Order, error) {
	m := new(Order)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *orderServiceClient) UpdateOrders(ctx context.Context, opts ...grpc.CallOption) (OrderService_UpdateOrdersClient, error) {
	stream, err := c.cc.NewStream(ctx, &OrderService_ServiceDesc.Streams[1], "/order.OrderService/updateOrders", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderServiceUpdateOrdersClient{stream}
	return x, nil
}

type OrderService_UpdateOrdersClient interface {
	Send(*Order) error
	CloseAndRecv() (*wrapperspb.StringValue, error)
	grpc.ClientStream
}

type orderServiceUpdateOrdersClient struct {
	grpc.ClientStream
}

func (x *orderServiceUpdateOrdersClient) Send(m *Order) error {
	return x.ClientStream.SendMsg(m)
}

func (x *orderServiceUpdateOrdersClient) CloseAndRecv() (*wrapperspb.StringValue, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(wrapperspb.StringValue)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *orderServiceClient) ProcessOrders(ctx context.Context, opts ...grpc.CallOption) (OrderService_ProcessOrdersClient, error) {
	stream, err := c.cc.NewStream(ctx, &OrderService_ServiceDesc.Streams[2], "/order.OrderService/processOrders", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderServiceProcessOrdersClient{stream}
	return x, nil
}

type OrderService_ProcessOrdersClient interface {
	Send(*wrapperspb.Int32Value) error
	Recv() (*ShipmentOrder, error)
	grpc.ClientStream
}

type orderServiceProcessOrdersClient struct {
	grpc.ClientStream
}

func (x *orderServiceProcessOrdersClient) Send(m *wrapperspb.Int32Value) error {
	return x.ClientStream.SendMsg(m)
}

func (x *orderServiceProcessOrdersClient) Recv() (*ShipmentOrder, error) {
	m := new(ShipmentOrder)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OrderServiceServer is the server API for OrderService service.
// All implementations should embed UnimplementedOrderServiceServer
// for forward compatibility
type OrderServiceServer interface {
	QueryOrders(*wrapperspb.StringValue, OrderService_QueryOrdersServer) error
	UpdateOrders(OrderService_UpdateOrdersServer) error
	ProcessOrders(OrderService_ProcessOrdersServer) error
}

// UnimplementedOrderServiceServer should be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (UnimplementedOrderServiceServer) QueryOrders(*wrapperspb.StringValue, OrderService_QueryOrdersServer) error {
	return status.Errorf(codes.Unimplemented, "method QueryOrders not implemented")
}
func (UnimplementedOrderServiceServer) UpdateOrders(OrderService_UpdateOrdersServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateOrders not implemented")
}
func (UnimplementedOrderServiceServer) ProcessOrders(OrderService_ProcessOrdersServer) error {
	return status.Errorf(codes.Unimplemented, "method ProcessOrders not implemented")
}

// UnsafeOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServiceServer will
// result in compilation errors.
type UnsafeOrderServiceServer interface {
	mustEmbedUnimplementedOrderServiceServer()
}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

func _OrderService_QueryOrders_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(wrapperspb.StringValue)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OrderServiceServer).QueryOrders(m, &orderServiceQueryOrdersServer{stream})
}

type OrderService_QueryOrdersServer interface {
	Send(*Order) error
	grpc.ServerStream
}

type orderServiceQueryOrdersServer struct {
	grpc.ServerStream
}

func (x *orderServiceQueryOrdersServer) Send(m *Order) error {
	return x.ServerStream.SendMsg(m)
}

func _OrderService_UpdateOrders_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OrderServiceServer).UpdateOrders(&orderServiceUpdateOrdersServer{stream})
}

type OrderService_UpdateOrdersServer interface {
	SendAndClose(*wrapperspb.StringValue) error
	Recv() (*Order, error)
	grpc.ServerStream
}

type orderServiceUpdateOrdersServer struct {
	grpc.ServerStream
}

func (x *orderServiceUpdateOrdersServer) SendAndClose(m *wrapperspb.StringValue) error {
	return x.ServerStream.SendMsg(m)
}

func (x *orderServiceUpdateOrdersServer) Recv() (*Order, error) {
	m := new(Order)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _OrderService_ProcessOrders_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OrderServiceServer).ProcessOrders(&orderServiceProcessOrdersServer{stream})
}

type OrderService_ProcessOrdersServer interface {
	Send(*ShipmentOrder) error
	Recv() (*wrapperspb.Int32Value, error)
	grpc.ServerStream
}

type orderServiceProcessOrdersServer struct {
	grpc.ServerStream
}

func (x *orderServiceProcessOrdersServer) Send(m *ShipmentOrder) error {
	return x.ServerStream.SendMsg(m)
}

func (x *orderServiceProcessOrdersServer) Recv() (*wrapperspb.Int32Value, error) {
	m := new(wrapperspb.Int32Value)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OrderService_ServiceDesc is the grpc.ServiceDesc for OrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "queryOrders",
			Handler:       _OrderService_QueryOrders_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "updateOrders",
			Handler:       _OrderService_UpdateOrders_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "processOrders",
			Handler:       _OrderService_ProcessOrders_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "order/order.proto",
}
