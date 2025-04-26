package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	object "github.com/khemingkapat/been_chillin/objects"
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

// GetAllFavoritesHandler retrieves all favorites
func GetAllFavoritesHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		rows, err := db.Query(`SELECT favorite_id, user_id, content_id FROM FAVORITE`)
		if err != nil {
			log.Println("Query error:", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "failed to fetch favorites"})
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

// DeleteFavoriteHandler deletes a favorite by id
func DeleteFavoriteHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

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
