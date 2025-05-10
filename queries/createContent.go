package queries

import (
	"context"
	"database/sql"
	"fmt"

	object "github.com/khemingkapat/been_chillin/objects"
)

func CreateContent(db *sql.DB, content *object.Content) (int, error) {
	// Start a transaction
	tx, err := db.BeginTx(context.Background(), &sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
	})
	if err != nil {
		return 0, fmt.Errorf("error starting transaction: %w", err)
	}
	// Defer rollback in case of an error (before committing)
	defer func() {
		if err != nil {
			fmt.Println("Rollback since", err.Error())
			tx.Rollback()
		}
	}()

	// Insert content and get the new ID
	var contentID int
	err = tx.QueryRow(`
		INSERT INTO CONTENT (
			title, description, release_year, duration, content_type, 
			total_seasons, thumbnail_url, video_url, director
		) VALUES (
			$1, $2, $3, $4, $5, 
			$6, $7, $8, $9
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
		content.Director).Scan(&contentID)
	if err != nil {
		return 0, fmt.Errorf("error inserting content: %w", err)
	}

	// Process categories
	for _, categoryName := range content.Categories {
		var categoryID int
		// First, check if the category already exists
		err = tx.QueryRow(`
			SELECT category_id FROM CATEGORY WHERE category_name = $1
		`, categoryName).Scan(&categoryID)
		if err != nil && err != sql.ErrNoRows {
			return 0, fmt.Errorf("error checking category: %w", err)
		}

		// If not found, insert the new category
		if err == sql.ErrNoRows {
			err = tx.QueryRow(`
				INSERT INTO CATEGORY (category_name) 
				VALUES ($1)
				RETURNING category_id
			`, categoryName).Scan(&categoryID)
			if err != nil {
				return 0, fmt.Errorf("error inserting category: %w", err)
			}
		}

		// Check if this content-category relationship already exists
		var exists bool
		err = tx.QueryRow(`
			SELECT EXISTS(SELECT 1 FROM CONTENT_CATEGORY WHERE content_id = $1 AND category_id = $2)
		`, contentID, categoryID).Scan(&exists)
		if err != nil {
			return 0, fmt.Errorf("error checking category relationship: %w", err)
		}

		if !exists {
			_, err = tx.Exec(`
				INSERT INTO CONTENT_CATEGORY (content_id, category_id) VALUES ($1, $2)
			`, contentID, categoryID)
			if err != nil {
				return 0, fmt.Errorf("error linking content to category: %w", err)
			}
		}
	}

	// Process languages using the same approach as categories
	for _, langName := range content.Languages {
		var languageID int
		// First, check if the language already exists
		err = tx.QueryRow(`
			SELECT language_id FROM LANGUAGE WHERE language_name = $1
		`, langName).Scan(&languageID)
		if err != nil && err != sql.ErrNoRows {
			return 0, fmt.Errorf("error checking language: %w", err)
		}

		// If not found, insert the new language
		if err == sql.ErrNoRows {
			err = tx.QueryRow(`
				INSERT INTO LANGUAGE (language_name) 
				VALUES ($1)
				RETURNING language_id
			`, langName).Scan(&languageID)
			if err != nil {
				return 0, fmt.Errorf("error inserting language: %w", err)
			}
		}

		// Check if this content-language relationship already exists
		var exists bool
		err = tx.QueryRow(`
			SELECT EXISTS(SELECT 1 FROM CONTENT_LANGUAGE WHERE content_id = $1 AND language_id = $2)
		`, contentID, languageID).Scan(&exists)
		if err != nil {
			return 0, fmt.Errorf("error checking language relationship: %w", err)
		}

		if !exists {
			_, err = tx.Exec(`
				INSERT INTO CONTENT_LANGUAGE (content_id, language_id) VALUES ($1, $2)
			`, contentID, languageID)
			if err != nil {
				return 0, fmt.Errorf("error linking content to language: %w", err)
			}
		}
	}

	// Process subtitles using the same upsert pattern
	for _, subLang := range content.Subtitles {
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
	}

	// Commit the transaction
	if err = tx.Commit(); err != nil {
		return 0, fmt.Errorf("error committing transaction: %w", err)
	}

	return contentID, nil
}
