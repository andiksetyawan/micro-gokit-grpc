package service

import (
	"fmt"
	pb "gokit-grpc/pkg/proto"
)

type HelloService interface {
	SayHello(req *pb.HelloRequest) (*pb.HelloResponse, error)
}

type helloService struct {
}

func NewHelloService() HelloService {
	return &helloService{}
}

func (s *helloService) SayHello(req *pb.HelloRequest) (*pb.HelloResponse, error) {
	res := fmt.Sprintf("HELLO %v", req.Name)
	return &pb.HelloResponse{Message: res}, nil
}
