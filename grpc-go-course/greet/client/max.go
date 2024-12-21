package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Wong-Bui/grpc-go-course/greet/proto"
)

func doMax(c pb.CalculatorServiceClient, data []int64) {
	log.Println("doMax was invoked")

	stream, err := c.Max(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	waitc := make(chan struct{})

	go func() {
		for _, num := range data {
			log.Printf("Send request: %v\n", num)
			stream.Send(&pb.MaxRequest{
				Number: num,
			})
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving: %v\n", err)
				break
			}

			log.Printf("Received: %v\n", res.Result)
		}

		close(waitc)
	}()

	<-waitc
}
