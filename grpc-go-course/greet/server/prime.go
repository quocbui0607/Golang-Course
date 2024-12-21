package main

import (
	"log"

	pb "github.com/Wong-Bui/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
)

func (s *Server) Prime(req *pb.PrimeRequest, stream grpc.ServerStreamingServer[pb.PrimeResponse]) error {
	var k int64 = 2
	N := req.N

	log.Printf("Prime function was invoked with: k=%v and N=%v\n", k, N)

	for N > 1 {
		if N%k == 0 {

			stream.Send(&pb.PrimeResponse{
				Result: k,
			})

			N /= k
		} else {
			k++
		}

	}

	return nil
}
