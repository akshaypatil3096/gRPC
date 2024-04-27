package main

import (
	"log"
	"time"

	pb "github.com/akshaypatil3096/gRPC/proto"
)

func (s *helloServer) SayHelloServerStream(req *pb.NameList, stream pb.GreetService_SayHelloServerStreamServer) error {
	log.Printf("got request from client with names: %v", req.Names)
	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello " + name,
		}

		if err := stream.SendMsg(res); err != nil {
			return err
		}

		time.Sleep(time.Second * 2)
	}

	return nil
}
