package main

import (
	"context"
	pb "grpc/calculator/proto"
	"log"
)

func (s *Server) Add(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Add function was invoked %v\n", in)

	return &pb.SumResponse{
		Result: in.Num1 + in.Num2,
	}, nil
}
