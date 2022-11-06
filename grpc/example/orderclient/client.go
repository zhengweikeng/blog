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
	conn, err := grpc.Dial("127.0.0.1:10000",
		grpc.WithUnaryInterceptor(orderUnaryClientInterceptor),
		grpc.WithStreamInterceptor(orderStreamInterceptor),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
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

	stream3, err := client.ProcessOrders(context.Background())
	if err != nil {
		log.Fatalf("process order fail:%v", err)
		return
	}

	if err := stream3.Send(wrapperspb.Int32(1)); err != nil {
		log.Fatalf("send order fail:%v", err)
		return
	}
	if err := stream3.Send(wrapperspb.Int32(2)); err != nil {
		log.Fatalf("send order fail:%v", err)
		return
	}
	if err := stream3.Send(wrapperspb.Int32(3)); err != nil {
		log.Fatalf("send order fail:%v", err)
		return
	}

	ch := make(chan struct{})
	go func() {
		for {
			shipOrder, err := stream3.Recv()
			if err != nil {
				if err == io.EOF {
					close(ch)
					break
				}
				continue
			}
			log.Printf("ship order: [%v]\n", shipOrder)
		}
	}()

	if err := stream3.CloseSend(); err != nil {
		log.Fatalf("close stream error:%v", err)
	}
	<-ch
}

func orderUnaryClientInterceptor(ctx context.Context,
	method string,
	req, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	log.Printf("invoke remote method:%s", method)

	err := invoker(ctx, method, req, reply, cc, opts...)
	if err != nil {
		log.Printf("invoke err:%v", err)
	}

	return nil
}

func orderStreamInterceptor(ctx context.Context,
	desc *grpc.StreamDesc,
	cc *grpc.ClientConn,
	method string,
	streamer grpc.Streamer,
	opts ...grpc.CallOption) (grpc.ClientStream, error) {
	log.Printf("invoke remote method:%s", method)

	stream, err := streamer(ctx, desc, cc, method, opts...)
	if err != nil {
		log.Printf("streamer err:%v", err)
	}
	return stream, err
}
