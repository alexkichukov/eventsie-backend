package events

import (
	"context"
	"eventsie/api/client"
	pb "eventsie/pb/events"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetEvents(svc *client.Services) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		category := c.Query("category")
		city := c.Query("city")
		tags := c.Query("tags")
		priceFrom := c.Query("priceFrom")
		priceTo := c.Query("priceTo")

		request := &pb.FindManyRequest{}

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

		return c.Status(int(resp.Status)).JSON(resp.Events)
	}
}
