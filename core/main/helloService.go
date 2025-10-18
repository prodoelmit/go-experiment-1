package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/prodoelmit/go-experiment-1/proto/hello"
)

func (s *orchestratorServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", req)
	message := fmt.Sprintf("Hello, %s", (*req).Name)
	return &pb.HelloResponse{Message: message}, nil
}
