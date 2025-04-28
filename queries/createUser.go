package queries

import (
	"database/sql"
	"time"

	"github.com/khemingkapat/been_chillin/auth"
	object "github.com/khemingkapat/been_chillin/objects"
)

func CreateUser(db *sql.DB, user *object.User) error {
	err := auth.EncryptUser(user)
	if err != nil {
		return err
	}

	user.Registration = time.Now()

	// Insert user ลง database
	_, err = db.Exec(`
		INSERT INTO "user" (username, email, password, subscription, registration, age, profile_pic)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`,
		user.UserName,
		user.Email,
		user.Password,
		user.Subscription,
		user.Registration,
		user.Age,
		user.ProfilePic,
	)
	if err != nil {
		return err
	}
	return nil
}
