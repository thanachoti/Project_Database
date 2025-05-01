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

func CreateUserHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := new(object.User)

		// ดึงค่าฟิลด์ปกติจาก Form
		user.UserName = c.FormValue("username")
		user.Email = c.FormValue("email")
		user.Password = c.FormValue("password")
		user.Subscription = c.FormValue("subscription")
		user.Age, _ = strconv.Atoi(c.FormValue("age"))

		// ดึงไฟล์ภาพ
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

		// บันทึก
		err = queries.CreateUser(db, user)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		return c.JSON(fiber.Map{
			"message": "User Created",
		})
	}
}
