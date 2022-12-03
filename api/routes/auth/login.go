package auth

import (
	"context"
	"eventsie/api/client"
	"eventsie/api/models"
	pb "eventsie/pb/auth"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Login(svc *client.Services) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		loginData := &models.LoginBody{}

		if err := c.BodyParser(loginData); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid data"})
		}

		request := &pb.LoginRequest{
			Email:    loginData.Email,
			Password: loginData.Password,
		}

		resp, _ := svc.Auth.Login(context.TODO(), request)

		if resp.Error {
			return c.Status(int(resp.Status)).JSON(fiber.Map{"message": resp.Message})
		}

		return c.Status(int(resp.Status)).JSON(fiber.Map{"token": resp.Token})
	}
}
