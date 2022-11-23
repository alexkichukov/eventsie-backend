package main

import (
	"events/api/config"
	"events/api/events"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.GetConfig()

	app := fiber.New()

	events.RegisterRoutes(app)

	if err := app.Listen(fmt.Sprintf(":%d", cfg.API_PORT)); err != nil {
		panic("Error while starting up app.")
	}

	fmt.Printf("App is listening on port %d\n", cfg.API_PORT)
}
