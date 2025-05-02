package handlers

import (
	"bytes"
	"database/sql"
	"io"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/khemingkapat/been_chillin/queries"
)

func UpdateUserHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fields := map[string]interface{}{}

		// ตรวจ username
		if v := c.FormValue("username"); v != "" {
			fields["username"] = v
		}
		if v := c.FormValue("email"); v != "" {
			fields["email"] = v
		}
		if v := c.FormValue("subscription"); v != "" {
			fields["subscription"] = v
		}
		if v := c.FormValue("age"); v != "" {
			if age, err := strconv.Atoi(v); err == nil {
				fields["age"] = age
			} else {
				log.Println("⚠️ age format invalid:", err)
				return c.Status(400).SendString("Invalid age format")
			}
		}

		// แนบ profile_pic ถ้ามี
		file, err := c.FormFile("profile_pic")
		if err == nil && file != nil {
			src, err := file.Open()
			if err != nil {
				return c.Status(500).SendString("Open file error")
			}
			defer src.Close()

			buf := new(bytes.Buffer)
			if _, err := io.Copy(buf, src); err != nil {
				return c.Status(500).SendString("Copy file error")
			}
			fields["profile_pic"] = buf.Bytes()
		}

		err = queries.UpdateUserFlexible(db, c.Params("user_id"), fields)

		if err != nil {
			log.Println("❌ UpdateUserFlexible error:", err)
			return c.Status(500).SendString("Update failed")
		}

		return c.JSON(fiber.Map{
			"message": "User updated successfully",
		})
	}
}
