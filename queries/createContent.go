package queries

import (
	"database/sql"
	"fmt"
	"time"

	object "github.com/khemingkapat/been_chillin/objects"
)

func CreateContent(db *sql.DB, content *object.Content) (int, error) {
	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		return 0, fmt.Errorf("error starting transaction: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Insert content and get the new ID
	var contentID int
	err = tx.QueryRow(`
		INSERT INTO CONTENT (
			title, description, release_year, duration, content_type, 
			total_seasons, thumbnail_url, video_url, rating, director
		) VALUES (
			$1, $2, $3, $4, $5, 
			$6, $7, $8, $9, $10
		) RETURNING content_id
	`,
		content.Title,
		content.Description,
		content.ReleaseYear,
		content.Duration,
		content.ContentType,
		content.TotalSeasons,
		content.ThumbnailURL,
		content.VideoURL,
		content.Rating,
		content.Director).Scan(&contentID)
	if err != nil {
		return 0, fmt.Errorf("error inserting content: %w", err)
	}

	// Process categories
	for _, categoryName := range content.Categories {
		// Use PostgreSQL's INSERT ON CONFLICT (upsert) to safely handle race conditions
		var categoryID int
		err = tx.QueryRow(`
			INSERT INTO CATEGORY (category_name) 
			VALUES ($1)
			ON CONFLICT (category_name) 
			DO UPDATE SET category_name = EXCLUDED.category_name
			RETURNING category_id
		`, categoryName).Scan(&categoryID)
		if err != nil {
			return 0, fmt.Errorf("error upserting category: %w", err)
		}

		// Check if this content-category relationship already exists
		var exists bool
		err = tx.QueryRow(`
			SELECT EXISTS(SELECT 1 FROM CONTENT_CATEGORY WHERE content_id = $1 AND category_id = $2)
		`, contentID, categoryID).Scan(&exists)
		if err != nil {
			return 0, fmt.Errorf("error checking category relationship: %w", err)
		}

		// Only insert if relationship doesn't exist
		if !exists {
			// Link content to category
			_, err = tx.Exec(`
				INSERT INTO CONTENT_CATEGORY (content_id, category_id) VALUES ($1, $2)
			`, contentID, categoryID)
			if err != nil {
				return 0, fmt.Errorf("error linking content to category: %w", err)
			}
		}
		time.Sleep(time.Millisecond * 1000)
	}

	// Process languages using the same upsert pattern
	for _, langName := range content.Languages {
		// Use PostgreSQL's INSERT ON CONFLICT (upsert) to safely handle race conditions
		var languageID int
		err = tx.QueryRow(`
			INSERT INTO LANGUAGE (language_name) 
			VALUES ($1)
			ON CONFLICT (language_name) 
			DO UPDATE SET language_name = EXCLUDED.language_name
			RETURNING language_id
		`, langName).Scan(&languageID)
		if err != nil {
			return 0, fmt.Errorf("error upserting language: %w", err)
		}

		// Check if this content-language relationship already exists
		var exists bool
		err = tx.QueryRow(`
			SELECT EXISTS(SELECT 1 FROM CONTENT_LANGUAGE WHERE content_id = $1 AND language_id = $2)
		`, contentID, languageID).Scan(&exists)
		if err != nil {
			return 0, fmt.Errorf("error checking language relationship: %w", err)
		}

		// Only insert if relationship doesn't exist
		if !exists {
			// Link content to language
			_, err = tx.Exec(`
				INSERT INTO CONTENT_LANGUAGE (content_id, language_id) VALUES ($1, $2)
			`, contentID, languageID)
			if err != nil {
				return 0, fmt.Errorf("error linking content to language: %w", err)
			}
		}
		time.Sleep(time.Millisecond * 1000)
	}

	// Process subtitles using the same upsert pattern
	for _, subLang := range content.Subtitles {
		// Use PostgreSQL's INSERT ON CONFLICT (upsert) to safely handle race conditions
		var languageID int
		err = tx.QueryRow(`
			INSERT INTO LANGUAGE (language_name) 
			VALUES ($1)
			ON CONFLICT (language_name) 
			DO UPDATE SET language_name = EXCLUDED.language_name
			RETURNING language_id
		`, subLang).Scan(&languageID)
		if err != nil {
			return 0, fmt.Errorf("error upserting subtitle language: %w", err)
		}

		// Check if this content-subtitle relationship already exists
		var exists bool
		err = tx.QueryRow(`
			SELECT EXISTS(SELECT 1 FROM CONTENT_SUBTITLE WHERE content_id = $1 AND language_id = $2)
		`, contentID, languageID).Scan(&exists)
		if err != nil {
			return 0, fmt.Errorf("error checking subtitle relationship: %w", err)
		}

		// Only insert if relationship doesn't exist
		if !exists {
			// Link content to subtitle language
			_, err = tx.Exec(`
				INSERT INTO CONTENT_SUBTITLE (content_id, language_id) VALUES ($1, $2)
			`, contentID, languageID)
			if err != nil {
				return 0, fmt.Errorf("error linking content to subtitle language: %w", err)
			}
		}
		time.Sleep(time.Millisecond * 1000)
	}

	// Commit the transaction
	if err = tx.Commit(); err != nil {
		return 0, fmt.Errorf("error committing transaction: %w", err)
	}

	return contentID, nil
}
