package auth

import (
	"context"
	"eventsie/api/client"
	"eventsie/api/models"
	pb "eventsie/pb/auth"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Register(svc *client.Services) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		user := &models.RegisterBody{}

		if err := c.BodyParser(user); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid data"})
		}

		request := &pb.RegisterRequest{
			Password:  user.Password,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		}

		resp, _ := svc.Auth.Register(context.TODO(), request)

		if resp.Error {
			return c.Status(int(resp.Status)).JSON(fiber.Map{"message": resp.Message})
		}

		return c.Status(int(resp.Status)).JSON(fiber.Map{"token": resp.Token})
	}
}
