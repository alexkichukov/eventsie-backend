package routes

import (
	"eventsie/api/client"
	"eventsie/api/routes/auth"
	"eventsie/api/routes/events"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	// Intialize microservices
	svc := client.NewServicesClient()

	// Auth
	app.Get("/user/:id", auth.GetUser(svc))
	app.Post("/login", auth.Login(svc))
	app.Post("/register", auth.Register(svc))
	app.Post("/favouriteEvent", auth.FavouriteEvent(svc))
	app.Post("/unfavouriteEvent", auth.UnfavouriteEvent(svc))

	// Events
	app.Get("/events", events.GetEvents(svc))
	app.Get("/events/:id", events.GetEvent(svc))
	app.Post("/new-event", events.CreateEvent(svc))
}
