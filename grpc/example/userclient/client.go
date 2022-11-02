package main

import (
	"context"
	"log"
	"time"

	pb "example/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:10000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	userSvcClient := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	userReq := pb.UserRequest{
		UserName: "Jack",
	}
	resp, err := userSvcClient.QueryUsers(ctx, &userReq)
	if err != nil {
		log.Fatalf("query user fail:%v", err)
		return
	}

	log.Printf("%v", resp)
}
