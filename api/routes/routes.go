package routes

import (
	"eventsie/api/client"
	"eventsie/api/middleware"
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
	app.Post("/attendEvent", auth.AttendEvent(svc))
	app.Post("/unattendEvent", auth.UnattendEvent(svc))

	// Events
	app.Get("/events", events.GetEvents(svc))
	app.Get("/events/:id", events.GetEvent(svc))
	app.Get("/categories", events.GetCategories(svc))
	app.Post("/newEvent", middleware.AuthGuard(svc), events.CreateEvent(svc))
	app.Post("/deleteEvent", middleware.AuthGuard(svc), events.DeleteEvent(svc))
	app.Post("/updateEvent", middleware.AuthGuard(svc), events.UpdateEvent(svc))
}
