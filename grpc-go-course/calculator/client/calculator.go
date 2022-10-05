package main

import (
	"context"
	pb "grpc/calculator/proto"
	"log"
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
