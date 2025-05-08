package handlers

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/khemingkapat/been_chillin/auth"
	object "github.com/khemingkapat/been_chillin/objects"
	"golang.org/x/crypto/bcrypt"
)

func ChangePasswordHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 1. Extract user ID from JWT token
		userID, err := auth.ExtractUserID(c)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
		}

		// 2. Parse request body
		var req object.PasswordChangeRequest
		if err := json.Unmarshal(c.Body(), &req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
		}

		// 3. Get current password hash from DB
		var hashedPassword string
		err = db.QueryRow(`SELECT password FROM "user" WHERE user_id = $1`, userID).Scan(&hashedPassword)
		if err != nil {
			log.Println("🧪 Password SELECT error:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "user not found"})
		}

		// 4. Compare currentPassword with hash
		if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.CurrentPassword)); err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "current password is incorrect"})
		}

		// 5. Hash new password
		newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to hash new password"})
		}

		// 6. Update password in DB
		_, err = db.Exec(`UPDATE "user" SET password = $1 WHERE user_id = $2`, newHashedPassword, userID)
		if err != nil {
			log.Println("🧪 Password UPDATE error:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to update password"})
		}

		return c.JSON(fiber.Map{"message": "password changed successfully"})
	}
}
