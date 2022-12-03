package auth

import (
	"context"
	"eventsie/api/client"
	pb "eventsie/pb/auth"

	"github.com/gofiber/fiber/v2"
)

func GetUser(svc *client.Services) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		resp, _ := svc.Auth.GetUser(context.TODO(), &pb.GetUserRequest{Id: id})
		if resp.Error {
			return c.Status(int(resp.Status)).JSON(fiber.Map{"message": resp.Message})
		}

		if resp.FavouriteEvents == nil {
			resp.FavouriteEvents = []string{}
		}
		if resp.AttendingEvents == nil {
			resp.AttendingEvents = []string{}
		}

		return c.Status(int(resp.Status)).JSON(fiber.Map{
			"id":              resp.Id,
			"firstName":       resp.FirstName,
			"lastName":        resp.LastName,
			"email":           resp.Email,
			"role":            resp.Role,
			"favouriteEvents": resp.FavouriteEvents,
			"attendingEvents": resp.AttendingEvents,
		})
	}
}
