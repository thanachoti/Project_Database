package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/khemingkapat/been_chillin/objects"
	"github.com/khemingkapat/been_chillin/queries"
)

// CreateFavoriteHandler inserts a favorite into DB
func CreateContentHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		newContent := new(object.Content)

		// 1. Parse body
		if err := c.BodyParser(newContent); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		contentID, err := queries.CreateContent(db, newContent)
		if err != nil {
			return c.Status(http.StatusInternalServerError).SendString(err.Error())
		}

		return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "content created", "content_id": contentID})
	}
}
