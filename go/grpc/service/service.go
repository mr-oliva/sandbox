package service

import (
	"context"

	"github.com/bookun/sandbox/go/grpc/pb"
)

type MyCatService struct {
}

func (s *MyCatService) GetMyCat(ctx context.Context, message *pb.GetMyCatMessage) (*pb.MyCatResponse, error) {
	return &pb.MyCatResponse{
		Name: message.TargetCat,
		Kind: message.TargetCat,
	}, nil
}
