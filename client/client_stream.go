package main

import (
	"context"
	"log"
	"time"

	pb "github.com/akshaypatil3096/gRPC/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("client streaming started")
	stream, err := client.SayHelloClientStream(context.Background())
	if err != nil {
		log.Fatalf("could not able to send the names, got an error: ", err)
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}

		if err := stream.SendMsg(req); err != nil {
			log.Fatalf("could not able to send the names, got an error: ", err)
		}

		log.Printf("sent the request with name: %v", name)
		time.Sleep(time.Second * 2)
	}

	res, err := stream.CloseAndRecv()
	log.Printf("Client Streaming finished")
	if err != nil {
		log.Fatalf("Error while receiving %v", err)
	}
	log.Printf("%v", res.Messages)
}
