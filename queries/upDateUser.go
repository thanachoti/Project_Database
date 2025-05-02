package queries

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
)

func UpdateUserFlexible(db *sql.DB, idStr string, fields map[string]interface{}) error {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("invalid user_id: %v", err)
	}

	if len(fields) == 0 {
		return fmt.Errorf("no fields to update")
	}

	setParts := []string{}
	args := []interface{}{}
	argPos := 1

	for col, val := range fields {
		setParts = append(setParts, fmt.Sprintf("%s = $%d", col, argPos))
		args = append(args, val)
		argPos++
	}

	query := fmt.Sprintf(`
		UPDATE "user"
		SET %s
		WHERE user_id = $%d
	`, strings.Join(setParts, ", "), argPos)

	args = append(args, id)

	_, err = db.Exec(query, args...)
	return err
}
