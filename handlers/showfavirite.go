package handlers

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	object "github.com/khemingkapat/been_chillin/objects"
)

// GetFavoritesByUserHandler retrieves favorites for a specific user
func GetFavoritesByUserHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Query("user_id")
		if userID == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "missing user_id parameter"})
		}

		query := `SELECT favorite_id, user_id, content_id FROM FAVORITE WHERE user_id = $1`
		rows, err := db.Query(query, userID)
		if err != nil {
			log.Println("Query error:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to fetch favorites"})
		}
		defer rows.Close()

		var favorites []object.Favorite
		for rows.Next() {
			var fav object.Favorite
			if err := rows.Scan(&fav.FavoriteID, &fav.UserID, &fav.ContentID); err != nil {
				log.Println("Row scan error:", err)
				continue
			}
			favorites = append(favorites, fav)
		}

		return c.JSON(favorites)
	}
}
