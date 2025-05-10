package queries

import (
	"database/sql"
	"fmt"
	"strings"
)

func UpdateContent(db *sql.DB, contentID int, updates map[string]string) error {
	setClauses := []string{}
	args := []any{}
	argIndex := 1

	for field, value := range updates {
		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", field, argIndex))
		args = append(args, value)
		argIndex++
	}

	// Add content_id to WHERE clause
	args = append(args, contentID)
	query := fmt.Sprintf("UPDATE CONTENT SET %s WHERE content_id = $%d", strings.Join(setClauses, ", "), argIndex)

	_, err := db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("error updating content: %w", err)
	}

	return nil
}
