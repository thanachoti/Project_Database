package handlers

import (
	"database/sql"
	"io"
	"strconv"
	// ✅ เพิ่ม log
	"github.com/gofiber/fiber/v2"
	// "github.com/khemingkapat/been_chillin/auth" // ✅ ต้องใช้สำหรับ EncryptUser
	object "github.com/khemingkapat/been_chillin/objects"
	"github.com/khemingkapat/been_chillin/queries"
)

func CreateUserHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := new(object.User)

		user.UserName = c.FormValue("username")
		user.Email = c.FormValue("email")
		user.Password = c.FormValue("password")
		user.Subscription = c.FormValue("subscription")
		ageStr := c.FormValue("age")
		user.Age, _ = strconv.Atoi(ageStr)

		// อ่านรูปจาก form
		fileHeader, err := c.FormFile("profile_pic")
		if err == nil && fileHeader != nil {
			src, _ := fileHeader.Open()
			defer src.Close()
			imgBytes, _ := io.ReadAll(src)
			user.ProfilePic = imgBytes
		}

		// 🔐 เข้ารหัส (ถ้าจะเปิดใช้) -- ยังเป็น comment
		// if err := auth.EncryptUser(user); err != nil {
		//     return c.Status(500).SendString("Hashing failed")
		// }

		err = queries.CreateUser(db, user)
		if err != nil {
			return c.Status(500).SendString("Database insert failed")
		}

		return c.JSON(fiber.Map{"message": "User created"})
	}
}
