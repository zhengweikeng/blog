package main

import (
	"context"
	"fmt"
	"goCase/gokit-id-generator-demo/endpoint"
	"goCase/gokit-id-generator-demo/service"
	httptransport "goCase/gokit-id-generator-demo/transport/http"
	"net/http"
)

var port = ":8080"

func main() {
	ctx := context.Background()
	digService := service.DefaultIdGenerator{}

	idGeneratorEndpoint := endpoint.MakeIdGeneratorEndpoint(digService)
	idGeneratorHandler := httptransport.MakeIdGeneratorHandler(ctx, idGeneratorEndpoint)

	fmt.Printf("server start at port %s\n", port)
	http.Handle("/", idGeneratorHandler)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println(err)
	}
}
