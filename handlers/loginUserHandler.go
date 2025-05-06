package handlers

import (
	"database/sql"
	"encoding/base64"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/khemingkapat/been_chillin/auth"
	object "github.com/khemingkapat/been_chillin/objects"
)

func LoginUserHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userLogin := new(object.UserLogin)

		// 1. Parse body
		if err := c.BodyParser(userLogin); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		// 2. ตรวจสอบ email/password และสร้าง token
		token, err := auth.LoginUser(db, userLogin)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
		}

		// 3. ดึง user detail จาก DB
		var user object.User
		err = db.QueryRow(`
			SELECT user_id, username, email, subscription, registration, age, profile_pic
			FROM "user"
			WHERE email = $1
		`, userLogin.Email).Scan(
			&user.UserID,
			&user.UserName,
			&user.Email,
			&user.Subscription,
			&user.Registration,
			&user.Age,
			&user.ProfilePic,
		)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to fetch user info",
			})
		}

		// 4. สร้าง cookie JWT
		c.Cookie(&fiber.Cookie{
			Name:     "jwt",
			Value:    token,
			Expires:  time.Now().Add(time.Hour * 3),
			HTTPOnly: true,
			SameSite: "lax",
			Path:     "/",
		})

		profilePictureBase64 := ""
		if len(user.ProfilePic) > 0 {
			profilePictureBase64 = "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(user.ProfilePic)
			log.Println("📦 Final profilePictureBase64:", profilePictureBase64[:50], "... (truncated)")
		} else {
			log.Println("⚠️ No profile picture found (user.ProfilePic is empty)")
		}

		// 6. ส่งข้อมูลกลับ
		return c.JSON(fiber.Map{
			"message": "Login Succeeded",
			"token":   token,
			"user": fiber.Map{
				"userID":         user.UserID,
				"username":       user.UserName,
				"email":          user.Email,
				"subscription":   user.Subscription,
				"age":            user.Age,
				"registration":   user.Registration.Format(time.RFC3339),
				"profilePicture": profilePictureBase64,
			},
		})
	}
}
