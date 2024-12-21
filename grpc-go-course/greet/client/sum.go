package main

import (
	"context"
	"log"

	pb "github.com/Wong-Bui/grpc-go-course/greet/proto"
)

func doSum(c pb.CalculatorServiceClient, first int64, second int64) {
	log.Println("doSum was invoked")

	res, err := c.Sum(context.Background(), &pb.SumRequest{
		First:  first,
		Second: second,
	})

	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}

	log.Printf("Sum: %v\n", res.Result)
}
