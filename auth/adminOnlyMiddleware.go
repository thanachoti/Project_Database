package auth

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func AdminOnlyMiddleware(c *fiber.Ctx) error {
	// Retrieve the JWT from the cookie
	cookie := c.Cookies("jwt")

	// Parse the token from the cookie
	token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(JWTSecretKey), nil
	})

	if err != nil || !token.Valid {
		fmt.Println("3")
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid token")
	}

	// Extract claims from the token
	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid token claims")
	}

	// Check if the user has the "manager" role
	role, exists := (*claims)["role"].(string)
	if !exists || role != "Admin" {
		return c.Status(fiber.StatusForbidden).SendString("Access denied: Manager role required")
	}

	// Proceed to the next handler or middleware
	return c.Next()
}
