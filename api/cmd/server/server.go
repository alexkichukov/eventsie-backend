package main

import (
	"context"
	"events/api/pb"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnsafeEventsServer
}

func (s *server) GetEvents(ctx context.Context, in *pb.GetEventsRequest) (*pb.GetEventsResponse, error) {
	fmt.Println("Received request with message:", in.Message)

	return &pb.GetEventsResponse{
		Status: 200,
		Error:  "",
		Events: "first event, second event, even third event",
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterEventsServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		panic(err)
	}

	fmt.Println("gRPC server up and running")
}
