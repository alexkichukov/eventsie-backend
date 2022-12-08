package events

import (
	"context"
	"eventsie/api/client"
	pb "eventsie/pb/events"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetCategories(svc *client.Services) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		resp, err := svc.Events.GetCategories(context.TODO(), &pb.GetCategoriesRequest{})

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Could not connect to events service"})
		}

		return c.Status(http.StatusOK).JSON(resp.Categories)
	}
}
