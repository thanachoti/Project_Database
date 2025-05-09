package queries

import (
	"database/sql"
	"fmt"
	// "strconv"
	// "strings"
)

func UpdateSubscription(db *sql.DB, subPlan string, userID int) error {
	query := `UPDATE "user" SET subscription = $1 WHERE user_id = $2`

	_, err := db.Exec(query, subPlan, userID)
	if err != nil {
		return fmt.Errorf("failed to update subscription plan: %v", err)
	}

	return nil
}
