package auth

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App) {
	svc := NewServiceClient()

	app.Post("/register", svc.Register)
}
