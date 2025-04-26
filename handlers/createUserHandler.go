package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	object "github.com/khemingkapat/been_chillin/objects"
	"github.com/khemingkapat/been_chillin/queries"
)

// CreateUserHandler inserts a new user
func CreateUserHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user, err := queries.CreateUser(c)
		if err != nil {
			log.Println("Parse error:", err)
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		query := `
			INSERT INTO "user" (username, email, password, subscription, registration, age, profile_pic)
			VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP, $5, $6)
			RETURNING user_id, registration
		`

		err = db.QueryRow(query,
			user.UserName,
			user.Email,
			user.Password,
			user.Subscription,
			user.Age,
			user.ProfilePic,
		).Scan(&user.UserID, &user.Registration)

		if err != nil {
			log.Println("DB error:", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "failed to create user"})
		}

		return c.Status(http.StatusCreated).JSON(user)
	}
}

// GetUserHandler retrieves a user by id
func GetUserHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		var user object.User
		err := db.QueryRow(`SELECT user_id, username, email, subscription, registration, age, profile_pic FROM "user" WHERE user_id = $1`, id).
			Scan(&user.UserID, &user.UserName, &user.Email, &user.Subscription, &user.Registration, &user.Age, &user.ProfilePic)

		if err != nil {
			if err == sql.ErrNoRows {
				return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
			}
			log.Println("Query error:", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "failed to fetch user"})
		}

		return c.JSON(user)
	}
}
