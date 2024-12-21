package main

import (
	"io"
	"log"

	pb "github.com/Wong-Bui/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
)

func (s *Server) Average(stream grpc.ClientStreamingServer[pb.AverageRequest, pb.AverageResponse]) error {
	log.Printf("Average function was invoked")

	res := float32(0)

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			result := res / float32(4)
			log.Printf("Making response: %v\n", result)

			return stream.SendAndClose(&pb.AverageResponse{
				Result: result,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}
		res += float32(req.Number)

		log.Printf("Receiving but wont return response: %v\n - response: %v", req, res)

	}

}
