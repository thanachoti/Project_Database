package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/khemingkapat/been_chillin/queries"
)

// DeleteContentHandler deletes a content entry by content_id
func DeleteUserHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userIDStr := c.Params("user_id")

		// Convert content_id to int
		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid content_id",
			})
		}

		// Delete content
		err = queries.DeleteUser(db, userID)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": "User Deleted Successfully",
		})
	}
}
