package events

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App) {
	svc := NewServiceClient()

	app.Get("/events", svc.GetAllEvents)
	app.Get("/events/:id", svc.GetEventByID)
	app.Post("/new-event", svc.CreateEvent)
}
