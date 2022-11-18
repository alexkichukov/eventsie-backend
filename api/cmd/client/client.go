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

	resp, err := client.FindOne(context.TODO(), &pb.FindOneRequest{Id: "kjlansckjasjd"})
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Event)
}
