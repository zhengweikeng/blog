package main

import (
	"context"
	pb "example/order"
	"io"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var orders = map[int32]*pb.Order{
	1: {
		Id:    1,
		Goods: []string{"iphone 14", "airpods"},
		Price: 6000,
	},
	2: {
		Id:    2,
		Goods: []string{"iphone 13"},
		Price: 4000,
	},
}

type OrderService struct{}

func (svc OrderService) QueryOrders(search *wrapperspb.StringValue,
	stream pb.OrderService_QueryOrdersServer) error {
	for _, o := range orders {
		for _, g := range o.Goods {
			if !strings.Contains(g, search.Value) {
				continue
			}

			err := stream.Send(o)
			if err != nil {
				return err
			}
			log.Printf("found order:%v", o)
			break
		}
	}

	return nil
}

func (svc OrderService) UpdateOrders(stream pb.OrderService_UpdateOrdersServer) error {
	for {
		order, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				log.Println("receive finish")
				return stream.SendAndClose(&wrapperspb.StringValue{Value: "receive finish"})
			}
			log.Fatalf("receive order error:%v", err)
			continue
		}

		orders[order.Id] = order
	}
}

func (svc OrderService) ProcessOrders(stream pb.OrderService_ProcessOrdersServer) error {
	maxBatch := 2

	var shipOrders []*pb.ShipmentOrder
	for {
		order_id, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				for _, shipOrder := range shipOrders {
					stream.Send(shipOrder)
				}
				return nil
			}
			log.Fatalf("receive order err:[%v]", err)
			return err
		}

		shipOrders = append(shipOrders, &pb.ShipmentOrder{
			OrderId: order_id.Value,
			Status:  "shipped",
		})
		if len(shipOrders) == maxBatch {
			for _, shipOrder := range shipOrders {
				stream.Send(shipOrder)
			}
			shipOrders = []*pb.ShipmentOrder{}
		}
	}
}

func main() {
	s := grpc.NewServer(grpc.UnaryInterceptor(orderInterceptor), grpc.StreamInterceptor(orderStreamInterceptorfunc))

	pb.RegisterOrderServiceServer(s, &OrderService{})
	log.Printf("start listen order service port:%d", 10000)

	ls, err := net.Listen("tcp", "127.0.0.1:10000")
	if err != nil {
		log.Fatal(err)
		return
	}

	if err := s.Serve(ls); err != nil {
		log.Fatalf("order service serve error:%v", err)
	}
}

func orderInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Printf("before handle: %s", info.FullMethod)
	ctx = context.WithValue(ctx, "foo", "bar")

	result, err := handler(ctx, req)

	log.Printf("after handle, result:%v, err:%v", result, err)
	return result, err
}

func orderStreamInterceptorfunc(srv interface{},
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler) error {
	log.Printf("before steaming handle: %s", info.FullMethod)

	err := handler(srv, ss)
	if err != nil {
		log.Printf("handle error:%v", err)
	}

	return err
}
