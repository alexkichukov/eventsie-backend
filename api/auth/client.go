package auth

import (
	"context"
	"eventsie/api/config"
	"eventsie/api/models"
	pb "eventsie/pb/auth"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.AuthClient
}

func NewServiceClient() *ServiceClient {
	cfg := config.GetConfig()
	serviceURL := fmt.Sprintf("localhost:%d", cfg.AUTH_SERVICE_PORT)

	conn, err := grpc.Dial(serviceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Couldn't connect to service:", err)
	}

	return &ServiceClient{
		Client: pb.NewAuthClient(conn),
	}
}

func (svc *ServiceClient) Register(c *fiber.Ctx) error {
	user := &models.RegisterBody{}

	if err := c.BodyParser(user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid data"})
	}

	request := &pb.RegisterRequest{
		Username:  user.Username,
		Password:  user.Password,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	resp, _ := svc.Client.Register(context.TODO(), request)

	if resp.Error {
		return c.Status(int(resp.Status)).JSON(fiber.Map{"message": resp.Message})
	}

	return c.Status(int(resp.Status)).JSON(fiber.Map{"message": resp.Message})
}
