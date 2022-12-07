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
		resp, _ := svc.Events.GetCategories(context.TODO(), &pb.GetCategoriesRequest{})
		return c.Status(http.StatusOK).JSON(resp.Categories)
	}
}
