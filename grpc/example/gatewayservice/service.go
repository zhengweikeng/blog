package main

import (
	"context"
	pb "example/gateway"
	"fmt"
	"log"
	"net"
	"net/http"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var (
	reg                = prometheus.NewRegistry()
	defaultGrpcMetrics = grpc_prometheus.NewServerMetrics()

	// Create a customized counter metric.
	sayHelloCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "greeter_sayHello_count",
		Help: "Total number of RPCs handled on the server.",
	}, []string{"name"})
)

func init() {
	reg.MustRegister(defaultGrpcMetrics, sayHelloCount)
}

func main() {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor))
	pb.RegisterGreeterServer(s, &GreeterService{})
	defaultGrpcMetrics.InitializeMetrics(s)

	addr := "127.0.0.1:10000"
	ls, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
		return
	}

	go func() {
		log.Fatalln(s.Serve(ls))
	}()

	go func() {
		http.ListenAndServe("0.0.0.0:9092", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
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
	sayHelloCount.WithLabelValues(req.Name).Inc()
	fmt.Println("hello:", req.Name)
	return &pb.HelloReply{Message: fmt.Sprintf("Hello %s", req.Name)}, nil
}

func (svc *GreeterService) Echo(_ context.Context, req *wrapperspb.StringValue) (*wrapperspb.StringValue, error) {
	fmt.Println("echo:", req.Value)
	return &wrapperspb.StringValue{Value: req.Value}, nil
}
