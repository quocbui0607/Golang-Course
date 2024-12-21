package main

import (
	"fmt"
	"log"

	pb "github.com/Wong-Bui/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
)

func (s *Server) GreetManyTimes(req *pb.GreetRequest, stream grpc.ServerStreamingServer[pb.GreetResponse]) error {
	log.Printf("GreetManyTimes function was invoked with: %v\n", req)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d", req.FirstName, i)

		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}

	return nil
}
