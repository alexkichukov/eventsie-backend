package events

import (
	"context"
	"eventsie/api/client"
	"eventsie/api/models"
	pb "eventsie/pb/events"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func CreateEvent(svc *client.Services) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		e := &models.CreateEventBody{}

		if err := c.BodyParser(e); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid event data"})
		}

		event := &pb.Event{
			Title:       e.Title,
			Date:        e.Date,
			Description: e.Description,
			Tags:        e.Tags,
			Category:    e.Category,
			CreatedBy:   e.CreatedBy,
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

		resp, _ := svc.Events.Add(context.TODO(), &pb.AddRequest{Event: event})

		return c.Status(int(resp.Status)).JSON(fiber.Map{"message": resp.Message})
	}
}
