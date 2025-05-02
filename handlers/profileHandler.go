package handlers

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/khemingkapat/been_chillin/queries"
)

func GetCurrentUserHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userTokenRaw := c.Locals("user")
		if userTokenRaw == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized or missing token",
			})
		}

		user := userTokenRaw.(*jwt.Token)
		claims := user.Claims.(*jwt.MapClaims)
		userID := int((*claims)["user_id"].(float64)) // ✅ ใช้ชื่อเดียวกัน

		email, name, profilePic, err := queries.GetUserByID(db, userID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "User not found",
			})
		}

		return c.JSON(fiber.Map{
			"email":          email,
			"name":           name,
			"profilePicture": profilePic,
		})
	}
}
