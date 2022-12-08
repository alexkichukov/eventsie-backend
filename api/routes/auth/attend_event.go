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

func AttendEvent(svc *client.Services) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		body := &models.FavouriteAttendBody{}

		if err := c.BodyParser(body); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid data"})
		}

		// Make sure the event exists
		eventResp, _ := svc.Events.FindOne(context.TODO(), &eventsPb.FindOneRequest{Id: body.EventID})
		if eventResp.Error {
			return c.Status(int(eventResp.Status)).JSON(fiber.Map{"message": "Could not add event to attending"})
		}

		// Add event to favourites
		resp, _ := svc.Auth.AttendEvent(context.TODO(), &authPb.AttendEventRequest{
			EventID: body.EventID,
			Token:   strings.TrimPrefix(c.Get("Authorization"), "Bearer "),
		})

		if resp.Error {
			return c.Status(int(resp.Status)).JSON(fiber.Map{"message": resp.Message})
		}

		if resp.AttendingEvents == nil {
			resp.AttendingEvents = []string{}
		}

		return c.Status(int(resp.Status)).JSON(fiber.Map{"attendingEvents": resp.AttendingEvents})
	}
}

func UnattendEvent(svc *client.Services) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		body := &models.FavouriteAttendBody{}

		if err := c.BodyParser(body); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid data"})
		}

		// Make sure the event exists
		eventResp, err := svc.Events.FindOne(context.TODO(), &eventsPb.FindOneRequest{Id: body.EventID})
		if eventResp.Error {
			return c.Status(int(eventResp.Status)).JSON(fiber.Map{"message": "Could not remove event from attending"})
		}

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Could not connect to events service"})
		}

		// Add event to favourites
		resp, err := svc.Auth.UnattendEvent(context.TODO(), &authPb.AttendEventRequest{
			EventID: body.EventID,
			Token:   strings.TrimPrefix(c.Get("Authorization"), "Bearer "),
		})

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Could not connect to auth service"})
		}

		if resp.Error {
			return c.Status(int(resp.Status)).JSON(fiber.Map{"message": resp.Message})
		}

		if resp.AttendingEvents == nil {
			resp.AttendingEvents = []string{}
		}

		return c.Status(int(resp.Status)).JSON(fiber.Map{"attendingEvents": resp.AttendingEvents})
	}
}
