package handlers

import (
	"database/sql"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/khemingkapat/been_chillin/auth"
	object "github.com/khemingkapat/been_chillin/objects"
)

func GetCurrentUserHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// ❶ ดึง user_id จาก JWT
		userIDFromToken, err := auth.ExtractUserID(c)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{"error": "unauthorized"})
		}

		// ❷ ดึง user_id จาก URL params
		paramID := c.Params("user_id")
		userIDFromParam, err := strconv.Atoi(paramID)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid user_id parameter"})
		}

		// ❸ ตรวจสอบว่า user จาก token เรียกของตัวเองหรือไม่
		if userIDFromToken != userIDFromParam {
			return c.Status(403).JSON(fiber.Map{"error": "forbidden: you can only access your own data"})
		}

		// ❹ query ข้อมูลผู้ใช้จากฐานข้อมูล
		var user object.User
		err = db.QueryRow(`
			SELECT user_id, username, email, subscription, registration, age
			FROM "user"
			WHERE user_id = $1
		`, userIDFromParam).Scan(&user.UserID, &user.UserName, &user.Email, &user.Subscription, &user.Registration, &user.Age)

		if err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "user not found"})
		}

		return c.JSON(user)
	}
}
