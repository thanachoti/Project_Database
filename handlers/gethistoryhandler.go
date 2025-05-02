package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	object "github.com/khemingkapat/been_chillin/objects"
	"github.com/khemingkapat/been_chillin/queries"
)

// CreateWatchHistoryHandler inserts a new watch history
func CreateWatchHistoryHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		history, err := queries.ParseWatchHistoryData(c)
		if err != nil {
			log.Println("Parse error:", err)
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		query := `
			INSERT INTO WATCH_HISTORY (user_id, content_id, watched_timestamp, progress, language_preference, cc_preference)
			VALUES ($1, $2, CURRENT_TIMESTAMP, $3, $4, $5)
			RETURNING history_id, watched_timestamp
		`

		err = db.QueryRow(query,
			history.UserID,
			history.ContentID,
			history.Progress,
			history.LanguagePreference,
			history.CcPreference,
		).Scan(&history.HistoryID, &history.Progress)

		if err != nil {
			log.Println("DB error:", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "failed to insert watch history"})
		}

		return c.Status(http.StatusCreated).JSON(history)
	}
}

// GetWatchHistoryHandler gets watch history by user
func GetWatchHistoryHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Params("user_id")

		rows, err := db.Query(`SELECT history_id, user_id, content_id, watched_timestamp, progress, language_preference, cc_preference FROM WATCH_HISTORY WHERE user_id = $1`, userID)
		if err != nil {
			log.Println("Query error:", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "failed to fetch watch history"})
		}
		defer rows.Close()

		var histories []object.WatchHistory
		for rows.Next() {
			var h object.WatchHistory
			if err := rows.Scan(&h.HistoryID, &h.UserID, &h.ContentID, &h.WatchedTimestamp, &h.Progress, &h.LanguagePreference, &h.CcPreference); err != nil {
				log.Println("Row scan error:", err)
				continue
			}
			histories = append(histories, h)
		}

		return c.JSON(histories)
	}
}
