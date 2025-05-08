package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/khemingkapat/been_chillin/auth"
)

// DeleteFavoriteHandler deletes a favorite by id
func DeleteFavoriteHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		contentID := c.Params("content_id")
		userID, err := auth.ExtractUserID(c)
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "you are not logged in"})
		}

		res, err := db.Exec(`DELETE FROM FAVORITE WHERE user_id = $1 and content_id = $2`, userID, contentID)
		if err != nil {
			log.Println("Delete error:", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "failed to delete favorite"})
		}

		rowsAffected, _ := res.RowsAffected()
		if rowsAffected == 0 {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "favorite not found"})
		}

		return c.SendStatus(http.StatusNoContent)
	}
}
