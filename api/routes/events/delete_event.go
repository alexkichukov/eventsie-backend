package events

import (
	"context"
	"eventsie/api/client"
	"eventsie/api/models"
	pb "eventsie/pb/events"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func DeleteEvent(svc *client.Services) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// Parse body
		e := &models.DeleteEventBody{}
		if err := c.BodyParser(e); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid event data"})
		}

		user := c.Locals("user").(fiber.Map)

		resp, err := svc.Events.Delete(context.TODO(), &pb.DeleteRequest{
			EventID:  e.Id,
			UserID:   user["Id"].(string),
			UserRole: user["Role"].(string),
		})

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Unexpected error"})
		}

		return c.Status(int(resp.Status)).JSON(fiber.Map{"message": resp.Message})
	}
}
