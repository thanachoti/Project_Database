package handlers

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/khemingkapat/been_chillin/queries"
)

func CreateProfileHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 🔍 Debug raw request body
		log.Println("Request Body:", string(c.Body()))

		// Parse profile
		profile, err := queries.CreateProfile(c)
		if err != nil {
			log.Println("Parse error:", err) // 🔍 Log error
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		// 🔍 Debug profile after parse
		log.Printf("Parsed Profile: %+v\n", profile)

		// Insert profile into DB
		query := `INSERT INTO PROFILE (user_id, profileName, age_restriction)
                  VALUES ($1, $2, $3) RETURNING profile_id`
		err = db.QueryRow(query, profile.UserID, profile.ProfileName, profile.AgeRestriction).
			Scan(&profile.ProfileID)
		if err != nil {
			log.Println("DB insert error:", err) // 🔍 Log DB error
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create profile: " + err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(profile)
	}
}
