package main

import (
	"log"
	"net"

	pb "github.com/Wong-Bui/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	// opts := []grpc.ServerOption{}
	// tls := false

	// if tls {
	// 	certFile := "/Users/macbook/Downloads/Golang/grpc-go-course/greet/ssl/server.crt"
	// 	keyFile := "/Users/macbook/Downloads/Golang/grpc-go-course/greet/ssl/server.pem"
	// 	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)

	// 	if err != nil {
	// 		log.Fatalf("Failed to loading certificates: %v\n", err)

	// 	}

	// 	opts = append(opts, grpc.Creds(creds))
	// }

	s := grpc.NewServer()

	pb.RegisterGreetServiceServer(s, &Server{})
	pb.RegisterCalculatorServiceServer(s, &Server{})

	reflection.Register(s)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
