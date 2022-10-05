package main

import (
	"context"
	pb "grpc/calculator/proto"
	"io"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func doRunningMax(c pb.CalculatorServiceClient) {
	log.Println("doGreetEveryone was invoked")

	stream, err := c.RunningMax(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	reqs := []*pb.RunningMaxRequest{
		{Number: 1},
		{Number: 5},
		{Number: 3},
		{Number: 6},
		{Number: 2},
		{Number: 20},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Send Request: %s\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			msg, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error Reading Stream: %v\n", err)
			}

			log.Printf("Current Max: %d\n", msg.Result)
		}
		close(waitc)
	}()

	<-waitc
}

func doSqRt(c pb.CalculatorServiceClient, n int64) {
	log.Println("doSqRt was invoked")

	res, err := c.SqRt(context.Background(), &pb.SqRtRequest{
		Number: n,
	})

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			log.Fatalf("Error Received message from server:%v\n", e.Message())
			log.Fatalf("Error Received code from server:%v\n", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Println("We probably sent a negative number!")
			}
		} else {
			log.Fatalf("A non grpc error:%v\n", err)
		}
	}

	log.Printf("SqRt: %f\n", res.Result)
}
