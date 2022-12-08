package events

import (
	"context"
	"eventsie/api/client"
	"eventsie/api/models"
	pb "eventsie/pb/events"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func UpdateEvent(svc *client.Services) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		e := &models.UpdateEventBody{}

		if err := c.BodyParser(e); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid event data"})
		}

		user := c.Locals("user").(fiber.Map)

		event := &pb.Event{
			Id:          e.Id,
			Title:       e.Title,
			Date:        e.Date,
			Description: e.Description,
			Tags:        e.Tags,
			Category:    e.Category,
			CreatedBy:   user["Id"].(string),
			Location: &pb.Location{
				Address:  e.Location.GetAddress(),
				City:     e.Location.GetCity(),
				Postcode: e.Location.GetPostcode(),
			},
			Price: &pb.Price{
				From: e.Price.GetFrom(),
				To:   e.Price.GetTo(),
			},
		}

		resp, err := svc.Events.Update(context.TODO(), &pb.UpdateRequest{
			Event:    event,
			EventID:  e.Id,
			UserID:   user["Id"].(string),
			UserRole: user["Role"].(string),
		})

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Could not connect to events service"})
		}

		return c.Status(int(resp.Status)).JSON(fiber.Map{"message": resp.Message})
	}
}
