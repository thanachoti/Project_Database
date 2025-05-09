package queries

import (
	"database/sql"
	"fmt"
	"log"

	object "github.com/khemingkapat/been_chillin/objects"
)

func GetContents(db *sql.DB) ([]object.Content, error) {
	//  the SQL query
	query := `SELECT content_id, title, description, release_year, duration, content_type, 
              total_seasons, thumbnail_url, video_url, rating 
              FROM CONTENT`

	// Use db.Query to execute the query and get a result set (rows).
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying database: %w", err)
	}
	defer rows.Close()

	//  a slice to store the Content structs.
	var contents []object.Content

	// Iterate over the rows using rows.Next().
	for rows.Next() {
		var content object.Content
		// Use rows.Scan() to read the data from the current row into a Content struct.
		err := rows.Scan(
			&content.ContentID,
			&content.Title,
			&content.Description,
			&content.ReleaseYear,
			&content.Duration,
			&content.ContentType,
			&content.TotalSeasons,
			&content.ThumbnailURL,
			&content.VideoURL,
			&content.Rating,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}

		// Fetch related data (languages, subtitles, categories)
		content.Languages, err = getLanguagesForContent(db, content.ContentID)
		if err != nil {
			log.Printf("Error getting languages for movie %d: %v", content.ContentID, err)
			//  consider whether to continue or return an error
		}
		content.Subtitles, err = getSubtitlesForContent(db, content.ContentID)
		if err != nil {
			log.Printf("Error getting subtitles for movie %d: %v", content.ContentID, err)
			//  consider whether to continue or return an error
		}
		content.Categories, err = getCategoriesForContent(db, content.ContentID)
		if err != nil {
			log.Printf("Error getting categories for movie %d: %v", content.ContentID, err)
			//  consider whether to continue or return an error
		}

		contents = append(contents, content)
	}

	// Check for any errors that occurred during the iteration with rows.Err().
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %w", err)
	}

	return contents, nil
}

// getLanguagesForContent retrieves the languages for a given movie ID.
func getLanguagesForContent(db *sql.DB, contentID int) ([]string, error) {
	query := `SELECT l.language_name 
              FROM CONTENT_LANGUAGE cl
              JOIN LANGUAGE l ON cl.language_id = l.language_id
              WHERE cl.content_id = $1`
	rows, err := db.Query(query, contentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var languages []string
	for rows.Next() {
		var lang string
		err := rows.Scan(&lang)
		if err != nil {
			return nil, err
		}
		languages = append(languages, lang)
	}
	return languages, nil
}

// getSubtitlesForContent retrieves the subtitles for a given movie ID.
func getSubtitlesForContent(db *sql.DB, contentID int) ([]string, error) {
	query := `SELECT l.language_name 
              FROM CONTENT_SUBTITLE cs
              JOIN LANGUAGE l ON cs.language_id = l.language_id
              WHERE cs.content_id = $1`
	rows, err := db.Query(query, contentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subtitles []string
	for rows.Next() {
		var sub string
		err := rows.Scan(&sub)
		if err != nil {
			return nil, err
		}
		subtitles = append(subtitles, sub)
	}
	return subtitles, nil
}

// getCategoriesForContent retrieves the categories for a given movie ID.
func getCategoriesForContent(db *sql.DB, contentID int) ([]string, error) {
	query := `SELECT c.category_name 
              FROM CONTENT_CATEGORY cc
              JOIN CATEGORY c ON cc.category_id = c.category_id
              WHERE cc.content_id = $1`
	rows, err := db.Query(query, contentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []string
	for rows.Next() {
		var category string
		err := rows.Scan(&category)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}
