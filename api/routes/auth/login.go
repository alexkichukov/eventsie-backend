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

		resp, err := svc.Auth.Login(context.TODO(), request)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Could not connect to auth service"})
		}

		if resp.Error {
			return c.Status(int(resp.Status)).JSON(fiber.Map{"message": resp.Message})
		}

		if resp.User.FavouriteEvents == nil {
			resp.User.FavouriteEvents = []string{}
		}
		if resp.User.AttendingEvents == nil {
			resp.User.AttendingEvents = []string{}
		}

		return c.Status(int(resp.Status)).JSON(fiber.Map{
			"token":           resp.Token,
			"id":              resp.User.Id,
			"firstName":       resp.User.FirstName,
			"lastName":        resp.User.LastName,
			"email":           resp.User.Email,
			"role":            resp.User.Role,
			"favouriteEvents": resp.User.FavouriteEvents,
			"attendingEvents": resp.User.AttendingEvents,
		})
	}
}
