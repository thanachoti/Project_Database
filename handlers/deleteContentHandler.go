package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/khemingkapat/been_chillin/queries"
)

// DeleteContentHandler deletes a content entry by content_id
func DeleteContentHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		contentIDStr := c.Params("content_id")

		// Convert content_id to int
		contentID, err := strconv.Atoi(contentIDStr)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid content_id",
			})
		}

		// Delete content
		err = queries.DeleteContent(db, contentID)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": "Content deleted successfully",
		})
	}
}
