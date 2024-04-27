package main

import (
	"log"
	"net"

	pb "github.com/akshaypatil3096/gRPC/proto"

	"google.golang.org/grpc"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("%v", err)
	}

	defer lis.Close()

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("server started at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		panic("error")
	}
}
