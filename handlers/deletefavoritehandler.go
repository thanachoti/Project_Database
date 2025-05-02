package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// DeleteFavoriteHandler deletes a favorite by id
func DeleteFavoriteHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("user_id")

		res, err := db.Exec(`DELETE FROM FAVORITE WHERE favorite_id = $1`, id)
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
