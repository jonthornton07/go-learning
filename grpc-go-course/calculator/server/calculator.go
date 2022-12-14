package main

import (
	"context"
	"fmt"
	pb "grpc/calculator/proto"
	"io"
	"log"
	"math"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Add(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Add function was invoked %v\n", in)

	return &pb.SumResponse{
		Result: in.Num1 + in.Num2,
	}, nil
}

func (s *Server) Primes(in *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("Primes function was invoked %v\n", in)

	number := in.Number
	divisor := int64(2)

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: divisor,
			})
			number /= divisor
		} else {
			divisor++
		}
	}

	return nil
}

func (s *Server) Average(stream pb.CalculatorService_AverageServer) error {
	log.Printf("Average function was invoked")

	sum := 0
	count := 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AverageResponse{
				Result: float64(sum) / float64(count),
			})
		}

		if err != nil {
			log.Fatalf("Error reading client stream: %v\n", err)
		}

		sum += int(req.Number)
		count++
	}
}

func (s *Server) RunningMax(stream pb.CalculatorService_RunningMaxServer) error {
	log.Printf("RunningMax function was invoked")

	runningMax := int64(0)

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error reading client stream: %v\n", err)
		}

		if req.Number > runningMax {
			runningMax = req.Number

			err = stream.Send(&pb.RunningMaxResponse{
				Result: runningMax,
			})

			if err != nil {
				log.Fatalf("Error while sending data to client: %v\n", err)
			}
		}

	}
}

func (s *Server) SqRt(ctx context.Context, in *pb.SqRtRequest) (*pb.SqRtRespone, error) {
	log.Printf("SqRt function was invoked %v\n", in)

	number := in.Number

	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %d", number),
		)
	}

	return &pb.SqRtRespone{
		Result: math.Sqrt(float64(number)),
	}, nil
}
