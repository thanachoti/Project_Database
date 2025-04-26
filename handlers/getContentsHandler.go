package handlers

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/khemingkapat/been_chillin/queries"
)

func GetContentsHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		contents, err := queries.GetContents(db)
		if err != nil {
			log.Println(err)
			return err
		}
		return c.JSON(contents)
	}
}
