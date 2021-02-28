package main

import (
	"context"
	"fmt"
	"log"
	greetpb "protobuf/greet"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client is starting...")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Could not connect: ", err)
	}

	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)
	//fmt.Println("Created a client:", c)
	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting Unary RPC service...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Kaushik",
			LastName:  "Biswas",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalln("Error while calling Greet RPC: ", err)
	}
	log.Println("Response from Greet: ", res.Result)
}
