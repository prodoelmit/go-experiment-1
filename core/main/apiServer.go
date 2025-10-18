package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/prodoelmit/go-experiment-1/proto/hello"
	"google.golang.org/grpc"
)

type orchestratorServer struct {
	pb.HelloServiceServer
}

func Start() {
	port := 50051
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServer, &orchestratorServer{})

	log.Printf("server listening on port %d", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve on port %d: %v", port, err)
	}
}
