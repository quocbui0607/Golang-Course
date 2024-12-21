package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Wong-Bui/grpc-go-course/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GreetWithDeadline(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("GreetWithDeadline was invoked with %v\n", req)

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("The client canceled the request!")

			return nil, status.Error(codes.Canceled, "The client canceled the request")
		}
	}

	time.Sleep(1 * time.Second)

	return &pb.GreetResponse{
		Result: "Xin chào " + req.FirstName,
	}, nil
}
