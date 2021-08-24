package main

import (
	"context"
	"fmt"
	"go-grpc-k8s-starter-server/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	proto.UnimplementedAddServiceServer
}

func (s *server) Compute(_ context.Context, input *proto.AddRequest) (*proto.AddResponse, error) {
	inputA, inputB := input.GetA(), input.GetB()
	response := &proto.AddResponse{Result: inputA + inputB}
	return response, nil
}

func main() {
	// Setup a basic listener via tcp
	// Create a new GRPC Service (server{})
	// Register the Service
	// Have the Service listen
	conn, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal("failed to connect to listener", err)
	}

	s := grpc.NewServer()
	service := &server{}
	proto.RegisterAddServiceServer(s, service)
	fmt.Println("listening on 3000")
	if err := s.Serve(conn); err != nil {
		log.Fatal(err)
	}
}
