package auth

import (
	"database/sql"
	// "fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	object "github.com/khemingkapat/been_chillin/objects"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(db *sql.DB, user *object.UserLogin) (string, error) {
	// Find the user by email in the database
	selectedUser := new(object.User)
	err := db.QueryRow("SELECT user_id,username, email,password,role FROM \"user\" WHERE email = $1", user.Email).
		Scan(&selectedUser.UserID, &selectedUser.UserName, &selectedUser.Email, &selectedUser.Password, &selectedUser.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", err
		}
		return "", err
	}

	// Compare the hashed password from the database with the provided password
	err = bcrypt.CompareHashAndPassword([]byte(selectedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err // Return error if password is incorrect
	}

	// Generate JWT token after successful login
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = selectedUser.UserID // Use selectedUser.ID here
	claims["exp"] = time.Now().Add(time.Hour * 3).Unix()
	claims["role"] = selectedUser.Role

	t, err := token.SignedString([]byte(JWTSecretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}
