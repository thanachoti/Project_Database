package handlers

import (
	"database/sql"
	"io"

	"github.com/gofiber/fiber/v2"
	"github.com/khemingkapat/been_chillin/auth"
)

func UploadProfilePictureHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID, err := auth.ExtractUserID(c)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{"error": "unauthorized"})
		}

		file, err := c.FormFile("profile_pic")
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "no file received"})
		}

		src, _ := file.Open()
		defer src.Close()

		imgBytes, _ := io.ReadAll(src)

		_, err = db.Exec(`UPDATE "user" SET profile_pic = $1 WHERE user_id = $2`, imgBytes, userID)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "upload failed"})
		}

		return c.JSON(fiber.Map{"message": "profile picture uploaded"})
	}
}
