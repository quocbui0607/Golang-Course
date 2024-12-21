package main

import (
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/Wong-Bui/grpc-go-course/greet/proto"
)

var addr string = "localhost:50051"

func main() {

	tls := false
	opts := []grpc.DialOption{}

	if tls {
		certFile := "/Users/macbook/Downloads/Golang/grpc-go-course/greet/ssl/ca.crt"

		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("Error while loading CA trust certificate: %v\n", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}
	conn, err := grpc.NewClient(addr, opts...)

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)
	calculatorService := pb.NewCalculatorServiceClient(conn)

	doGreet(c)
	doGreetManyTimes(c)
	doLongGreet(c)
	doGreetEveryone(c)
	doGreetWithDeadline(c, 1*time.Second)

	doSum(calculatorService, 3, 5)
	doPrime(calculatorService)
	doAverage(calculatorService, []int64{1, 2, 3, 4})
	doMax(calculatorService, []int64{4, 7, 2, 19, 4, 6, 32})
	doSqrt(calculatorService, -10)
}
