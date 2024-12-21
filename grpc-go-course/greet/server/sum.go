package main

import (
	"context"
	"log"

	pb "github.com/Wong-Bui/grpc-go-course/greet/proto"
)

func (s *Server) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Greet function was invoked with %v\n", req)

	return &pb.SumResponse{
		Result: req.First + req.Second,
	}, nil
}
