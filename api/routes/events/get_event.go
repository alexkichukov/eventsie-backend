package events

import (
	"context"
	"eventsie/api/client"
	authPb "eventsie/pb/auth"
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

		// Find the user who created the event
		authResp, _ := svc.Auth.GetUser(context.TODO(), &authPb.GetUserRequest{Id: resp.Event.CreatedBy})
		if resp.Error {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Unexpected error"})
		}

		return c.Status(int(resp.Status)).JSON(fiber.Map{
			"id":          resp.Event.Id,
			"date":        resp.Event.Date,
			"title":       resp.Event.Title,
			"description": resp.Event.Description,
			"category":    resp.Event.Category,
			"tags":        resp.Event.Tags,
			"location": &fiber.Map{
				"address":  resp.Event.Location.Address,
				"city":     resp.Event.Location.City,
				"postcode": resp.Event.Location.Postcode,
			},
			"price": &fiber.Map{
				"from": resp.Event.Price.From,
				"to":   resp.Event.Price.To,
			},
			"createdBy": &fiber.Map{
				"id":        authResp.User.Id,
				"firstName": authResp.User.FirstName,
				"lastName":  authResp.User.LastName,
			},
		})
	}
}
