package transport

import (
	"context"
	"errors"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"gokit-grpc/internal/endpoint"
	"gokit-grpc/internal/service"
	pb "gokit-grpc/pkg/proto"
)

type grpcServer struct {
	sayHello grpctransport.Handler
	pb.UnimplementedHelloServer
}

func NewGRPCServer(svc service.HelloService) pb.HelloServer {
	e := endpoint.NewHelloEndpoint(svc)
	return &grpcServer{
		sayHello: grpctransport.NewServer(
			e.SayHelloEndpoint(),
			decodeGRPCSumRequest,
			encodeGRPCSumResponse,
		),
	}
}

func (s *grpcServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	_, rep, err := s.sayHello.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.HelloResponse), nil
}

func decodeGRPCSumRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.HelloRequest)
	//TODO seharus nya diconvert dari struct grpc ke struct biasa gokit -> jika mngkin butuh tambah transport http
	return req, nil
}

func encodeGRPCSumResponse(_ context.Context, response interface{}) (interface{}, error) {
	//TODO seharus nya diconvert dari struct biasa ke struct grpc -> jika mngkin butuh tambah transport http
	resp, ok := response.(*pb.HelloResponse)
	if !ok {
		return nil, errors.New("error assertion")
	}
	return resp, nil
}
