package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Wong-Bui/grpc-go-course/greet/proto"
)

func doPrime(c pb.CalculatorServiceClient) {
	log.Println("doPrime was invoked")

	req := &pb.PrimeRequest{
		N: 120,
	}

	stream, err := c.Prime(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling Prime: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("Prime: %d\n", msg.Result)
	}
}
