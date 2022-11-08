package main

import (
	"context"
	pb "example/gateway"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func main() {
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &GreeterService{})

	addr := "127.0.0.1:10000"
	ls, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
		return
	}

	go func() {
		log.Fatalln(s.Serve(ls))
	}()

	mux := runtime.NewServeMux()
	err = pb.RegisterGreeterHandlerFromEndpoint(context.Background(), mux, addr, []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	})
	if err != nil {
		log.Fatalln(err)
		return
	}

	if err = http.ListenAndServe("127.0.0.1:8080", mux); err != nil {
		log.Fatalln(err)
	}
}

type GreeterService struct{}

func (svc *GreeterService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("hello:", req.Name)
	return &pb.HelloReply{Message: fmt.Sprintf("Hello %s", req.Name)}, nil
}

func (svc *GreeterService) Echo(_ context.Context, req *wrapperspb.StringValue) (*wrapperspb.StringValue, error) {
	fmt.Println("echo:", req.Value)
	return &wrapperspb.StringValue{Value: req.Value}, nil
}
