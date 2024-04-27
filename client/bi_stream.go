package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/akshaypatil3096/gRPC/proto"
)

func callSayHelloBiDirectionalStream(client pb.GreetServiceClient, names *pb.NameList) {
	stream, err := client.SayHelloBiDirectionalStream(context.Background())
	if err != nil {
		log.Fatalf("could not able to send the names, got an error: ", err)
	}

	waitC := make(chan struct{})

	go func() {
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("error while streaming: %v", err)
			}

			log.Printf("received msg: %v", msg.Message)
		}
		close(waitC)
	}()

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

	if err := stream.CloseSend(); err != nil {
		log.Fatalf("Error while receiving %v", err)
	}

	<-waitC
	log.Printf("bi directional streaming completed")
}
