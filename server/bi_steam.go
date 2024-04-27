package main

import (
	"io"
	"log"

	pb "github.com/akshaypatil3096/gRPC/proto"
)

func (s *helloServer) SayHelloBiDirectionalStream(stream pb.GreetService_SayHelloBiDirectionalStreamServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("got an error while receiving the msg: %v", err)
			return err
		}

		log.Printf("got request with name: %v", req.Name)
		resp := &pb.HelloResponse{
			Message: "Hello " + req.Name,
		}

		if err := stream.SendMsg(resp); err != nil {
			return err
		}
	}

}
