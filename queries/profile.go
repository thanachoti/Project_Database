package queries

import (
	"database/sql"
)

func GetUserByID(db *sql.DB, userID int) (email, name, profilePic string, err error) {
	err = db.QueryRow(`
		SELECT email, name, profile_pic
		FROM "user"
		WHERE user_id = $1
	`, userID).Scan(&email, &name, &profilePic)

	return
}
