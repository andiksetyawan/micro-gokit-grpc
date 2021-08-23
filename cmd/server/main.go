package main

import (
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"gokit-grpc/internal/service"
	"gokit-grpc/internal/transport"
	"gokit-grpc/internal/util"
	pb "gokit-grpc/pkg/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

var grpcBind = util.GetEnv("GRPC_BIND_ADDRESS", "localhost:9001")

func main() {
	helloService := service.NewHelloService()
	grpcServer := transport.NewGRPCServer(helloService)

	lis, err := net.Listen("tcp", grpcBind)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
	pb.RegisterHelloServer(baseServer, grpcServer)

	log.Println("listening grpc :", grpcBind)
	panic(baseServer.Serve(lis))
}
