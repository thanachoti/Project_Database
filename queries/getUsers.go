package queries

import (
	"database/sql"
	"fmt"

	object "github.com/khemingkapat/been_chillin/objects"
)

func GetUsers(db *sql.DB) ([]object.User, error) {
	rows, err := db.Query(`SELECT user_id, username, email, password, subscription, registration, age, role, profile_pic FROM "user"`)
	if err != nil {
		return nil, fmt.Errorf("querying users: %w", err)
	}
	defer rows.Close()

	var users []object.User

	for rows.Next() {
		var user object.User
		err := rows.Scan(
			&user.UserID,
			&user.UserName,
			&user.Email,
			&user.Password,
			&user.Subscription,
			&user.Registration,
			&user.Age,
			&user.Role,
			&user.ProfilePic,
		)
		if err != nil {
			return nil, fmt.Errorf("scanning user: %w", err)
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("reading users: %w", err)
	}

	return users, nil
}
