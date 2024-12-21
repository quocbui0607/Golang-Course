package main

import (
	"io"
	"log"

	pb "github.com/Wong-Bui/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
)

func (s *Server) Max(stream grpc.BidiStreamingServer[pb.MaxRequest, pb.MaxResponse]) error {
	log.Println("GreetEveryone was invoked")
	max := int64(0)
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		if max < req.Number {
			max = req.Number

			err = stream.Send(&pb.MaxResponse{
				Result: max,
			})
			if err != nil {
				log.Fatalf("Error while sending data to client: %v\n", err)
			}
		}

		log.Printf("Number %v is lower than max %v", req.Number, max)

	}
}
