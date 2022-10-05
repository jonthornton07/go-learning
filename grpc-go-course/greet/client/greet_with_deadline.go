package main

import (
	"context"
	pb "grpc/greet/proto"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("doGreet was invoked")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{
		FirstName: "Jon",
	}

	res, err := c.GreetWithDeadline(ctx, req)

	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline Exceeded!")
				return
			}
			log.Fatalf("Error Received message from server:%v\n", e.Message())
			log.Fatalf("Error Received code from server:%v\n", e.Code())
		} else {
			log.Fatalf("A non grpc error:%v\n", err)
		}
	}

	log.Printf("Greeting: %s\n", res.Result)
}
