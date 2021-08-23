package main

import (
	"context"
	pb "gokit-grpc/pkg/proto"
	"google.golang.org/grpc"
	"log"
	"testing"
	"time"
)

const (
	address = "localhost:9001"
)

func TestGrpcClient(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	cGrpc := pb.NewHelloClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := cGrpc.SayHello(ctx, &pb.HelloRequest{Name: "PAEJOOOO"})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("hello ", res)
}
