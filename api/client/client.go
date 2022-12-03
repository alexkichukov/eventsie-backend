package client

import (
	"eventsie/api/config"
	authPb "eventsie/pb/auth"
	eventsPb "eventsie/pb/events"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Services struct {
	Events eventsPb.EventsClient
	Auth   authPb.AuthClient
}

func ConnectToService(port int) *grpc.ClientConn {
	serviceURL := fmt.Sprintf("localhost:%d", port)

	conn, err := grpc.Dial(serviceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Couldn't connect to service:", err)
	}

	return conn
}

func NewServicesClient() *Services {
	cfg := config.GetConfig()

	return &Services{
		Events: eventsPb.NewEventsClient(ConnectToService(cfg.EVENTS_SERVICE_PORT)),
		Auth:   authPb.NewAuthClient(ConnectToService(cfg.AUTH_SERVICE_PORT)),
	}
}
