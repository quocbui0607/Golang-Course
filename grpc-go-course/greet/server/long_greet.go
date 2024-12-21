package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/Wong-Bui/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
)

func (s *Server) LongGreet(stream grpc.ClientStreamingServer[pb.GreetRequest, pb.GreetResponse]) error {
	log.Printf("LongGreet function was invoked")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		log.Printf("Receiving: %v\n", req)

		res += fmt.Sprintf("Hello %s!\n", req.FirstName)
	}

}
