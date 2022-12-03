package auth

import (
	"context"
	"eventsie/api/client"
	"eventsie/api/models"
	pb "eventsie/pb/auth"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func FavouriteEvent(svc *client.Services) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		body := &models.FavouriteAttendBody{}

		if err := c.BodyParser(body); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid data"})
		}

		request := &pb.FavouriteEventRequest{
			EventId: body.EventID,
			Token:   strings.TrimPrefix(c.Get("Authorization"), "Bearer "),
		}

		resp, _ := svc.Auth.FavouriteEvent(context.TODO(), request)

		if resp.Error {
			return c.Status(int(resp.Status)).JSON(fiber.Map{"message": resp.Message})
		}

		return c.Status(int(resp.Status)).JSON(fiber.Map{"message": resp.Message})
	}
}
