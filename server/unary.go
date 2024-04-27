package main

import (
	"context"

	pb "github.com/akshaypatil3096/gRPC/proto"
)

func (h *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
