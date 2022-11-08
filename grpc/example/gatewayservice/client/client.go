package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	pb "example/gateway"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	reg := prometheus.NewRegistry()
	grpcMetrics := grpc_prometheus.NewClientMetrics()
	reg.MustRegister(grpcMetrics)

	conn, err := grpc.Dial(
		"127.0.0.1:10000",
		grpc.WithUnaryInterceptor(grpcMetrics.UnaryClientInterceptor()),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	go func() {
		http.ListenAndServe("0.0.0.0:9094", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	}()

	client := pb.NewGreeterClient(conn)
	fmt.Println("Start to call the method called SayHello every 3 seconds")
	go func() {
		for {
			_, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "Test"})
			if err != nil {
				return
			}
			time.Sleep(3 * time.Second)
		}
	}()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("You can press n or N to stop the process of client")
	for scanner.Scan() {
		if strings.ToLower(scanner.Text()) == "n" {
			os.Exit(0)
		}
	}
}
