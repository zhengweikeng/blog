package main

import (
	pb "example/order"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var orders = map[int]*pb.Order{
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

func main() {
	s := grpc.NewServer()

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
