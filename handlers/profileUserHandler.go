package handlers

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/khemingkapat/been_chillin/auth"
	"github.com/khemingkapat/been_chillin/queries"
)

func GetCurrentUserHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID, err := auth.ExtractUserID(c)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		email, username, profilePic, err := queries.GetUserByID(db, userID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "User not found",
			})
		}

		return c.JSON(fiber.Map{
			"user_id":        userID,
			"email":          email,
			"username":       username,
			"profilePicture": profilePic,
		})
	}
}
