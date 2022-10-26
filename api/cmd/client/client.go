package main

import (
	"context"
	"events/api/pb"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewEventsClient(conn)

	resp, err := client.GetEvents(context.TODO(), &pb.GetEventsRequest{Message: "Hello there"})
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Events)
}
