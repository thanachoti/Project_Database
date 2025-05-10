package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/khemingkapat/been_chillin/queries"
)

func UpdateContentHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		contentIDStr := c.Params("content_id")
		contentID, err := strconv.Atoi(contentIDStr)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid content_id",
			})
		}

		updates := c.Queries()
		if len(updates) == 0 {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "No fields provided for update",
			})
		}

		err = queries.UpdateContent(db, contentID, updates)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"message": "Content updated successfully",
		})
	}
}
