package queries

import (
	"database/sql"
	"fmt"

	object "github.com/khemingkapat/been_chillin/objects"
)

func GetContents(db *sql.DB) ([]object.Content, error) {
	query := `SELECT 
    c.content_id,
    c.title,
    c.description,
    c.release_year,
    c.duration,
    c.content_type,
    c.total_seasons,
    c.thumbnail_url,
    c.video_url,
    c.rating,
    c.director,
    array_remove(array_agg(DISTINCT cl.language_name), NULL) AS languages,
    array_remove(array_agg(DISTINCT sl.language_name), NULL) AS subtitles,
    array_remove(array_agg(DISTINCT cat.category_name), NULL) AS categories
    FROM CONTENT c
    LEFT JOIN CONTENT_LANGUAGE clink ON c.content_id = clink.content_id
    LEFT JOIN LANGUAGE cl ON clink.language_id = cl.language_id

    LEFT JOIN CONTENT_SUBTITLE slink ON c.content_id = slink.content_id
    LEFT JOIN LANGUAGE sl ON slink.language_id = sl.language_id

    LEFT JOIN CONTENT_CATEGORY catlink ON c.content_id = catlink.content_id
    LEFT JOIN CATEGORY cat ON catlink.category_id = cat.category_id

    GROUP BY c.content_id;
    `

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var contents []object.Content
	for rows.Next() {
		var c object.Content
		err := rows.Scan(
			&c.ContentID,
			&c.Title,
			&c.Description,
			&c.ReleaseYear,
			&c.Duration,
			&c.ContentType,
			&c.TotalSeasons,
			&c.ThumbnailURL,
			&c.VideoURL,
			&c.Rating,
			&c.Director,
			&c.Languages,
			&c.Subtitles,
			&c.Categories,
		)
		if err != nil {
			return nil, fmt.Errorf("row scan failed: %w", err)
		}
		contents = append(contents, c)
	}
	return contents, rows.Err()
}
