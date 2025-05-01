package handlers

import (
	"bytes"
	"database/sql"
	"io"
	"strconv"

	"github.com/gofiber/fiber/v2"
	object "github.com/khemingkapat/been_chillin/objects"
	"github.com/khemingkapat/been_chillin/queries"
)

func UpdateUserHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		user := new(object.User)

		user.UserName = c.FormValue("username")
		user.Email = c.FormValue("email")
		user.Subscription = c.FormValue("subscription")
		user.Age, _ = strconv.Atoi(c.FormValue("age"))

		file, err := c.FormFile("profile_pic")
		if err == nil {
			src, err := file.Open()
			if err != nil {
				return c.Status(500).SendString("Open file error")
			}
			defer src.Close()

			buf := new(bytes.Buffer)
			if _, err := io.Copy(buf, src); err != nil {
				return c.Status(500).SendString("Copy file error")
			}
			user.ProfilePic = buf.Bytes()
		}

		err = queries.UpdateUser(db, user, id)
		if err != nil {
			return c.Status(500).SendString("Update failed")
		}

		return c.JSON(fiber.Map{
			"message": "User updated successfully",
		})
	}
}
