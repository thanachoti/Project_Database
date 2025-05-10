package queries

import (
	"database/sql"
	"errors"
	"fmt"
)

func DeleteContent(db *sql.DB, contentID int) error {
	// Check if content exists
	var exists bool
	err := db.QueryRow(`SELECT EXISTS(SELECT 1 FROM CONTENT WHERE content_id = $1)`, contentID).Scan(&exists)
	if err != nil {
		return fmt.Errorf("error checking if content exists: %w", err)
	}

	if !exists {
		return errors.New("content not found")
	}

	// Delete the content
	_, err = db.Exec(`DELETE FROM CONTENT WHERE content_id = $1`, contentID)
	if err != nil {
		return fmt.Errorf("error deleting content: %w", err)
	}

	return nil
}
