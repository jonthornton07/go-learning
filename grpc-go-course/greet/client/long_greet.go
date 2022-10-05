package main

import (
	"context"
	pb "grpc/greet/proto"
	"io"
	"log"
	"time"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("goLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Jon"},
		{FirstName: "Jon1"},
		{FirstName: "Jon2"},
	}

	stream, err := c.LongGreet(context.Background())

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

	log.Printf("GreetManyTimes: %s\n", res.Result)
}
