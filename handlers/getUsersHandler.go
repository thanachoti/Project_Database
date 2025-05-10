package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/khemingkapat/been_chillin/queries"
)

func GetUsersHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		users, err := queries.GetUsers(db)
		if err != nil {
			c.Status(http.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(users)
	}
}
