package main

import (
	"eventsie/auth/config"
	"eventsie/auth/server"
	pb "eventsie/pb/auth"
	"fmt"
	"net"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.GetConfig()

	mgm.SetDefaultConfig(nil, cfg.MONGO_NAME, options.Client().ApplyURI(cfg.MONGO_URI))

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.PORT))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Auth gRPC server running at port %d\n", cfg.PORT)

	s := grpc.NewServer()
	pb.RegisterAuthServer(s, &server.Server{})

	if err := s.Serve(listener); err != nil {
		panic(err)
	}
}
