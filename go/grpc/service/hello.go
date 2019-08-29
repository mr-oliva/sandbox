package service

import (
	"context"

	"github.com/bookun/sandbox/go/grpc/pb"
)

type HelloService struct{}

func (h *HelloService) Hello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "hello, " + in.Name,
	}, nil
}
