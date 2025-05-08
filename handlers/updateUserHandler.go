package handlers

import (
	"database/sql"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/khemingkapat/been_chillin/auth"
)

func UpdateUserProfileHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 1. ดึง userID จาก JWT (login)
		jwtUserID, err := auth.ExtractUserID(c)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{"error": "unauthorized"})
		}

		// 2. ดึง user_id จาก URL parameter
		paramUserID := c.Params("user_id")

		// 3. ตรวจสอบว่าเป็นเจ้าของจริงหรือไม่
		if strconv.Itoa(jwtUserID) != paramUserID {
			return c.Status(403).JSON(fiber.Map{"error": "forbidden"})
		}

		// 4. ดึงข้อมูลจากฟอร์ม
		username := c.FormValue("username")
		email := c.FormValue("email")
		subscription := c.FormValue("subscription")

		// 5. อัปเดต DB
		_, err = db.Exec(`
			UPDATE "user" 
			SET username = $1, email = $2, subscription = $3 
			WHERE user_id = $4
		`, username, email, subscription, jwtUserID)

		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "update failed"})
		}

		return c.JSON(fiber.Map{"message": "Profile updated successfully"})
	}
}
