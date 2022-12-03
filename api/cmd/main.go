package main

import (
	"eventsie/api/config"
	"eventsie/api/routes"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	cfg := config.GetConfig()

	app := fiber.New()

	// Cors
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	// Routes
	routes.RegisterRoutes(app)

	// Listen
	if err := app.Listen(fmt.Sprintf(":%d", cfg.API_PORT)); err != nil {
		fmt.Println(err)
		panic("Error while starting up app.")
	}

	fmt.Printf("App is listening on port %d\n", cfg.API_PORT)
}
