package queries

import (
	"database/sql"

	"github.com/khemingkapat/been_chillin/auth"
	"github.com/khemingkapat/been_chillin/objects"
)

func CreateUser(db *sql.DB, user *object.User) error {
	err := auth.EncryptUser(user)
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO \"user\" (username, email, password) VALUES ($1, $2, $3)",
		user.UserName, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}
