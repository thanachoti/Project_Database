package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/khemingkapat/been_chillin/queries"
)

// CreateFavoriteHandler inserts a favorite into DB
func CreateFavoriteHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		favorite, err := queries.ParseFavoriteData(c)
		if err != nil {
			log.Println("Parse error:", err)
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		query := `INSERT INTO FAVORITE (user_id, content_id) VALUES ($1, $2) RETURNING favorite_id`

		err = db.QueryRow(query, favorite.UserID, favorite.ContentID).Scan(&favorite.FavoriteID)
		if err != nil {
			log.Println("DB error:", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "failed to insert favorite"})
		}

		return c.Status(http.StatusCreated).JSON(favorite)
	}
}
