package main

import (
	"context"

	"log"

	pb "github.com/akshaypatil3096/gRPC/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx := context.Background()
	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	log.Printf("%s", res.Message)
}
