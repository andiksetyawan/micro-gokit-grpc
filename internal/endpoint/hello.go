package endpoint

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"gokit-grpc/internal/service"
	pb "gokit-grpc/pkg/proto"
)

type HelloEndpoint interface {
	SayHelloEndpoint() endpoint.Endpoint
}

type helloEndpoint struct {
	svc service.HelloService
}

func NewHelloEndpoint(svc service.HelloService) HelloEndpoint {
	return &helloEndpoint{svc: svc}
}

func (h *helloEndpoint) SayHelloEndpoint() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*pb.HelloRequest)
		if !ok {
			return nil, errors.New("error assertion endpoint")
		}

		return h.svc.SayHello(req)
	}
}
