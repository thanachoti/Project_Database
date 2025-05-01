package queries

import (
	"database/sql"

	object "github.com/khemingkapat/been_chillin/objects"
)

func UpdateUser(db *sql.DB, user *object.User, id string) error {
	_, err := db.Exec(`
		UPDATE "user"
		SET username = $1, email = $2, subscription = $3, age = $4, profile_pic = $5
		WHERE user_id = $6
	`,
		user.UserName,
		user.Email,
		user.Subscription,
		user.Age,
		user.ProfilePic,
		id,
	)
	return err
}
