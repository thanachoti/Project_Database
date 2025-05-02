package queries

import (
	"database/sql"
)

func GetUserByID(db *sql.DB, userID int) (email, username string, profilePic string, err error) {
	err = db.QueryRow(`
		SELECT email, username, profile_pic
		FROM "user"
		WHERE user_id = $1
	`, userID).Scan(&email, &username, &profilePic)
	return
}
