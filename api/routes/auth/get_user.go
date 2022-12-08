package auth

import (
	"context"
	"eventsie/api/client"
	pb "eventsie/pb/auth"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetUser(svc *client.Services) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		resp, err := svc.Auth.GetUser(context.TODO(), &pb.GetUserRequest{Id: id})
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
