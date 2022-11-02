package main

import (
	"context"
	"errors"
	pb "example/user"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &UserService{})

	log.Printf("start listen user service port:%d", 10000)

	ls, err := net.Listen("tcp", "127.0.0.1:10000")
	if err != nil {
		log.Fatal(err)
		return
	}

	if err := s.Serve(ls); err != nil {
		log.Fatalf("user service serve error:%v", err)
	}
}

var users = map[string]*pb.User{
	"Jerry": {
		Id:     1,
		Name:   "Jerry",
		Age:    21,
		Gender: 1,
	},
	"Jack": {
		Id:     2,
		Name:   "Jack",
		Age:    30,
		Gender: 1,
	},
}

type UserService struct{}

func (svc UserService) QueryUsers(ctx context.Context, userReq *pb.UserRequest) (*pb.UsersResponse, error) {
	u, ok := users[userReq.UserName]
	if !ok {
		return nil, errors.New("user not found")
	}

	resp := &pb.UsersResponse{
		Code:  0,
		Users: []*pb.User{u},
	}
	return resp, nil
}
