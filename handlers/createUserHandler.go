package handlers

import (
	"database/sql"
	"log" // ✅ เพิ่ม log

	"github.com/gofiber/fiber/v2"
	// "github.com/khemingkapat/been_chillin/auth" // ✅ ต้องใช้สำหรับ EncryptUser
	object "github.com/khemingkapat/been_chillin/objects"
	"github.com/khemingkapat/been_chillin/queries"
)

func CreateUserHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := new(object.User)

		// 🧪 ตรวจว่า Body ส่งมาถูกหรือไม่
		if err := c.BodyParser(user); err != nil {
			log.Println("❌ BodyParser error:", err)
			return c.Status(fiber.StatusBadRequest).SendString("Invalid input format")
		}

		// 🔐 เข้ารหัสรหัสผ่าน
		// if err := auth.EncryptUser(user); err != nil {
		// 	log.Println("❌ Password hashing error:", err)
		// 	return c.Status(fiber.StatusInternalServerError).SendString("Hashing failed")
		// }

		log.Println("✅ Password hashed:", user.Password)

		// 🚀 บันทึกผู้ใช้ลงฐานข้อมูล
		err := queries.CreateUser(db, user)
		if err != nil {
			log.Println("❌ DB insert error:", err)
			return c.Status(fiber.StatusBadRequest).SendString("Database error")
		}

		log.Printf("✅ User %s (%s) created successfully\n", user.UserName, user.Email)

		return c.JSON(fiber.Map{
			"message": "User Created",
		})
	}
}
