package main

import (
	"context"
	"fmt"
	"log"
	"net"
	greetpb "protobuf/greet"

	"google.golang.org/grpc"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Println("Greet function was invoked with: \n", req)
	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()
	result := "Hello! " + firstName + " " + lastName
	resp := &greetpb.GreetResponse{
		Result: result,
	}
	return resp, nil
}

func main() {

	fmt.Println("Starting server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln("Failed to listen: ", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalln("Failed to serve: ", err)
	}
}
