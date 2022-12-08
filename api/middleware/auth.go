package middleware

import (
	"context"
	"eventsie/api/client"
	pb "eventsie/pb/auth"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// Verifies Authentication header for JWT token and returns a 401 if invalid.
// Attaches the information from the token to Locals if valid
func AuthGuard(svc *client.Services) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		token := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")

		resp, _ := svc.Auth.ValidateToken(context.TODO(), &pb.ValidateRequest{Token: token})

		// Add decoded information to locals
		if resp.Valid {
			c.Locals("user", fiber.Map{
				"Id":        resp.Id,
				"FirstName": resp.FirstName,
				"LastName":  resp.LastName,
				"Email":     resp.Email,
				"Role":      resp.Role,
			})
			return c.Next()
		}

		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "You are not authorized"})
	}
}
