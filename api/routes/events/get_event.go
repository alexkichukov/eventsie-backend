package events

import (
	"context"
	"eventsie/api/client"
	pb "eventsie/pb/events"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetEvent(svc *client.Services) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		resp, _ := svc.Events.FindOne(context.TODO(), &pb.FindOneRequest{Id: c.Params("id")})

		// There is an error
		if resp.Error {
			return c.Status(int(resp.Status)).JSON(fiber.Map{"message": resp.Message})
		}

		// There was no event found
		if resp.Event == nil {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": fmt.Sprintf("No event with ID %s was found", c.Params("id")),
			})
		}

		return c.Status(int(resp.Status)).JSON(resp.Event)
	}
}
