package auth

import (
	"context"
	"eventsie/api/client"
	"eventsie/api/models"
	authPb "eventsie/pb/auth"
	eventsPb "eventsie/pb/events"
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

		// Make sure the event exists
		eventResp, _ := svc.Events.FindOne(context.TODO(), &eventsPb.FindOneRequest{Id: body.EventID})
		if eventResp.Error {
			return c.Status(int(eventResp.Status)).JSON(fiber.Map{"message": "Could not favourite event"})
		}

		// Add event to favourites
		resp, _ := svc.Auth.FavouriteEvent(context.TODO(), &authPb.FavouriteEventRequest{
			EventID: body.EventID,
			Token:   strings.TrimPrefix(c.Get("Authorization"), "Bearer "),
		})

		if resp.Error {
			return c.Status(int(resp.Status)).JSON(fiber.Map{"message": resp.Message})
		}

		if resp.FavouriteEvents == nil {
			resp.FavouriteEvents = []string{}
		}

		return c.Status(int(resp.Status)).JSON(fiber.Map{"favouriteEvents": resp.FavouriteEvents})
	}
}

func UnfavouriteEvent(svc *client.Services) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		body := &models.FavouriteAttendBody{}

		if err := c.BodyParser(body); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid data"})
		}

		// Make sure the event exists
		eventResp, _ := svc.Events.FindOne(context.TODO(), &eventsPb.FindOneRequest{Id: body.EventID})
		if eventResp.Error {
			return c.Status(int(eventResp.Status)).JSON(fiber.Map{"message": "Could not favourite event"})
		}

		// Add event to favourites
		resp, _ := svc.Auth.UnfavouriteEvent(context.TODO(), &authPb.FavouriteEventRequest{
			EventID: body.EventID,
			Token:   strings.TrimPrefix(c.Get("Authorization"), "Bearer "),
		})

		if resp.Error {
			return c.Status(int(resp.Status)).JSON(fiber.Map{"message": resp.Message})
		}

		if resp.FavouriteEvents == nil {
			resp.FavouriteEvents = []string{}
		}

		return c.Status(int(resp.Status)).JSON(fiber.Map{"favouriteEvents": resp.FavouriteEvents})
	}
}
