package main

import (
	"context"
	"events/api/pb"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedEventsServer
}

func (s *server) FindOne(ctx context.Context, in *pb.FindOneRequest) (*pb.FindOneResponse, error) {
	fmt.Println("Looking for event with ID:", in.Id)

	return &pb.FindOneResponse{
		Status: 200,
		Event: &pb.Event{
			Id:          in.Id,
			Title:       "Example event",
			Date:        "2022",
			Description: "Description here...",
			Tags:        []string{"epic", "dev"},
			Category:    "dev10983mkf-D",
			Price:       &pb.Price{},
			Location: &pb.Location{
				Address:  "Some Street",
				City:     "Example City",
				Postcode: "4029",
			},
		},
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
