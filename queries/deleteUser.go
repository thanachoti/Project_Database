package queries

import (
	"database/sql"
	"errors"
	"fmt"
)

func DeleteUser(db *sql.DB, userID int) error {
	// Check if content exists
	var exists bool
	err := db.QueryRow(`SELECT EXISTS(SELECT 1 FROM "user" WHERE user_id = $1)`, userID).Scan(&exists)
	if err != nil {
		return fmt.Errorf("error checking if content exists: %w", err)
	}

	if !exists {
		return errors.New("content not found")
	}

	// Delete the content
	_, err = db.Exec(`DELETE FROM "user" WHERE user_id = $1`, userID)
	if err != nil {
		return fmt.Errorf("error deleting content: %w", err)
	}

	return nil
}
