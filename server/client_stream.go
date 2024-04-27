package main

import (
	"io"
	"log"

	pb "github.com/akshaypatil3096/gRPC/proto"
)

func (s *helloServer) SayHelloClientStream(stream pb.GreetService_SayHelloClientStreamServer) error {
	var message = make([]string, 0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageList{
				Messages: message,
			})
		}

		if err != nil {
			log.Fatalf("got an error while receiving the msg: %v", err)
			return err
		}

		log.Printf("got request with name: %v", req.Name)
		message = append(message, "Hello ", req.Name)
	}

	log.Printf("messages: %v", message)

	return nil
}
