package auth

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// ExtractUserID extracts the user_id from the JWT token stored in Fiber's context.
// It returns the user_id as an int if successful, or an error if the token or claims are invalid.
func ExtractUserID(c *fiber.Ctx) (int, error) {
	// Retrieve the raw token from the context, set by JWT middleware.
	tokenRaw := c.Locals("user")

	// Assert the token type to *jwt.Token
	token, ok := tokenRaw.(*jwt.Token)
	if !ok {
		return 0, errors.New("invalid token type")
	}

	// Extract claims as a map (standard JWT claims format)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	// Access the "user_id" field from the claims
	// JWT numeric fields are usually decoded as float64, even if originally integers.
	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return 0, errors.New("user_id not found in token")
	}

	// Convert user_id from float64 to int and return it
	return int(userIDFloat), nil
}
