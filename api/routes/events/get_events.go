package events

import (
	"context"
	"eventsie/api/client"
	authPb "eventsie/pb/auth"
	eventsPb "eventsie/pb/events"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetEvents(svc *client.Services) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id := c.Query("id")
		category := c.Query("category")
		city := c.Query("city")
		tags := c.Query("tags")
		createdBy := c.Query("createdBy")
		priceFrom := c.Query("priceFrom")
		priceTo := c.Query("priceTo")

		request := &eventsPb.FindManyRequest{}

		// Parse the filter parameters
		if len(id) > 0 {
			c := strings.Split(category, ",")
			request.Id = c
		}
		if len(category) > 0 {
			c := strings.Split(category, ",")
			request.Categories = c
		}
		if len(city) > 0 {
			c := strings.Split(city, ",")
			request.Cities = c
		}
		if len(tags) > 0 {
			c := strings.Split(tags, ",")
			request.Tags = c
		}
		if len(createdBy) > 0 {
			c := strings.Split(createdBy, ",")
			request.CreatedBy = c
		}
		if len(priceFrom) > 0 {
			from, err := strconv.ParseFloat(priceFrom, 64)
			if err == nil && from > 0 {
				request.PriceFrom = &from
			}
		}
		if len(priceTo) > 0 {
			to, err := strconv.ParseFloat(priceTo, 64)
			if err == nil && to > 0 {
				request.PriceTo = &to
			}
		}

		resp, _ := svc.Events.FindMany(context.TODO(), request)

		// There is an error
		if resp.Error {
			return c.Status(int(resp.Status)).JSON(fiber.Map{"message": resp.Message})
		}

		// There was no event found
		if resp.Events == nil {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "No events found"})
		}

		// Fiber map to be returned as json
		m := make([]*fiber.Map, len(resp.Events))

		// Get event creators info
		for i, event := range resp.Events {
			resp, _ := svc.Auth.GetUser(context.TODO(), &authPb.GetUserRequest{Id: event.CreatedBy})
			if resp.Error {
				return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Unexpected error"})
			}

			m[i] = &fiber.Map{
				"id":          event.Id,
				"date":        event.Date,
				"title":       event.Title,
				"description": event.Description,
				"category":    event.Category,
				"tags":        event.Tags,
				"location": &fiber.Map{
					"address":  event.Location.Address,
					"city":     event.Location.City,
					"postcode": event.Location.Postcode,
				},
				"price": &fiber.Map{
					"from": event.Price.From,
					"to":   event.Price.To,
				},
				"createdBy": &fiber.Map{
					"id":        event.CreatedBy,
					"firstName": resp.User.FirstName,
					"lastName":  resp.User.LastName,
				},
			}
		}

		return c.Status(int(resp.Status)).JSON(m)
	}
}
