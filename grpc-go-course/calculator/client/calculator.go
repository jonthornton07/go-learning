package main

import (
	"context"
	pb "grpc/calculator/proto"
	"io"
	"log"
	"time"
)

func doAdd(c pb.CalculatorServiceClient) {
	log.Println("doAdd was invoked")
	res, err := c.Add(context.Background(), &pb.SumRequest{
		Num1: 3,
		Num2: 10,
	})

	if err != nil {
		log.Fatalf("Could not greet :%v\n", err)
	}

	log.Printf("Sum: %d\n", res.Result)
}

func doPrimes(c pb.CalculatorServiceClient) {
	log.Println("doPrimes was invoked")

	req := &pb.PrimeRequest{
		Number: 12390392840,
	}

	stream, err := c.Primes(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not calculate primes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error Reading Stream: %v\n", err)
		}

		log.Printf("Prime Result: %d\n", msg.Result)
	}
}

func doAverage(c pb.CalculatorServiceClient) {
	log.Println("doAverage was invoked")

	reqs := []*pb.AverageRequest{
		{Number: 2},
		{Number: 4},
		{Number: 6},
	}

	stream, err := c.Average(context.Background())

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)

		if err == io.EOF {
			break
		}
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error Reading Stream: %v\n", err)
	}

	log.Printf("doAverage: %f\n", res.Result)
}
