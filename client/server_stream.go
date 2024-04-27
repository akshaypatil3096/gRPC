package main

import (
	"context"
	"io"
	"log"

	pb "github.com/akshaypatil3096/gRPC/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("streaming started")

	stream, err := client.SayHelloServerStream(context.Background(), names)
	if err != nil {
		log.Fatalf("could not able to send names: %v", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error while streaming: %v", err)
		}

		log.Printf(msg.Message)
	}

	log.Printf("streaming finished")
}
