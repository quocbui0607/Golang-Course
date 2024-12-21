package main

import (
	"context"
	"log"

	pb "github.com/Wong-Bui/grpc-go-course/greet/proto"
)

func doAverage(c pb.CalculatorServiceClient, data []int64) {
	log.Println("doAverage was invoke")

	stream, err := c.Average(context.Background())

	if err != nil {
		log.Fatalf("Error while calling Average: %v\n", err)
	}

	for _, number := range data {
		log.Printf("Sending req: %v\n", number)
		stream.Send(&pb.AverageRequest{
			Number: number,
		})
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from Average: %v\n", err)
	}

	log.Printf("doAverage: %v\n", res.Result)
}
