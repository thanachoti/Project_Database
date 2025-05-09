package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func DeleteReviewHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		reviewID := c.Params("review_id")

		_, err := db.Exec(`DELETE FROM review WHERE review_id = $1`, reviewID)
		if err != nil {
			log.Println("Delete error:", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "failed to delete review"})
		}

		return c.SendStatus(fiber.StatusNoContent) // 204
	}
}
