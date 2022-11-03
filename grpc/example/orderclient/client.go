package main

import (
	"context"
	pb "example/order"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:10000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	client := pb.NewOrderServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	stream, err := client.QueryOrders(ctx, &wrapperspb.StringValue{Value: "iphone"})
	if err != nil {
		log.Fatalf("query order fail:%v", err)
		return
	}

	for {
		order, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("receive order stream error:%v", err)
			continue
		}
		log.Printf("query order: [%+v]", order)
	}

	ctx2, cancel2 := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel2()

	stream2, err := client.UpdateOrders(ctx2)
	if err != nil {
		log.Fatalf("update order fail:%v", err)
		return
	}

	if err := stream2.Send(&pb.Order{
		Id:    1,
		Goods: []string{"iphone14", "ipad", "airpods"},
		Price: 10000,
	}); err != nil {
		log.Fatalf("send order fail:%v", err)
		return
	}

	if err := stream2.Send(&pb.Order{
		Id:    2,
		Goods: []string{"iphone14", "macbook pro"},
		Price: 20000,
	}); err != nil {
		log.Fatalf("send order fail:%v", err)
		return
	}

	response, err := stream2.CloseAndRecv()
	if err != nil {
		log.Fatalf("close stream fail:%v", err)
		return
	}
	log.Printf("update order resp:%v", response)
}
