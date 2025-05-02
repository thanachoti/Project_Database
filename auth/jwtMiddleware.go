package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var JWTSecretKey string = "mysecretkey"

// JWTMiddleware checks if the JWT token exists in the cookie and parses it
func JWTMiddleware(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	if cookie == "" {
		// If no token is found in the cookie, send an unauthorized response
		return c.Status(fiber.StatusUnauthorized).SendString("Missing or invalid token")
	}

	// Parse the token from the cookie
	token, err := jwt.ParseWithClaims(cookie, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Verify the token's signing method and provide the secret key for validation
		return []byte(JWTSecretKey), nil
	})

	if err != nil || !token.Valid {
		// If there's an error or the token is not valid, return an unauthorized response
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid token")
	}

	// Token is valid, add claims to the context
	_, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid token claims")
	}
	c.Locals("user", token)

	// Continue to the next handler
	return c.Next()
}
