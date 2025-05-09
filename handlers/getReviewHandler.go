package handlers

import (
	"database/sql"
	"encoding/base64"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	object "github.com/khemingkapat/been_chillin/objects"
	"github.com/khemingkapat/been_chillin/queries"
)

// CreateReviewHandler inserts a review into the database
func CreateReviewHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		review, err := queries.PrepareReview(c)
		if err != nil {
			log.Println("Parse error:", err)
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		log.Println("Review input:", review)

		// ✅ 1. Insert review และดึง review_id กลับมา
		insertQuery := `
			INSERT INTO REVIEW (user_id, content_id, rating, review_text, review_date)
			VALUES ($1, $2, $3, $4, $5)
			RETURNING review_id
		`
		err = db.QueryRow(insertQuery, review.UserID, review.ContentID, review.Rating, review.ReviewText, review.ReviewDate).
			Scan(&review.ReviewID)
		if err != nil {
			log.Println("DB insert error:", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "failed to insert review"})
		}

		// ✅ 2. ดึง username จาก user_id
		err = db.QueryRow(`SELECT username FROM "user" WHERE user_id = $1`, review.UserID).
			Scan(&review.Username)
		if err != nil {
			log.Println("DB username fetch error:", err)
			// ถ้าหาไม่เจอให้ fallback เป็น "User"
			review.Username = "User"
		}

		log.Println("✅ Review inserted with username:", review.Username)

		// ✅ 3. ส่งกลับ review (พร้อม username)
		return c.Status(http.StatusCreated).JSON(review)
	}
}

// GetReviewByContentIDHandler retrieves reviews for a specific content_id
// GetReviewByContentIDHandler retrieves reviews for a specific content_id
func GetReviewByContentIDHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		contentID := c.Params("content_id")
		var reviews []object.Review

		rows, err := db.Query(`
			SELECT 
				r.review_id, r.user_id, u.username, u.profile_pic, 
				r.content_id, r.rating, r.review_text, r.review_date
			FROM review r
			JOIN "user" u ON r.user_id = u.user_id
			WHERE r.content_id = $1
			ORDER BY r.review_date DESC
		`, contentID)

		if err != nil {
			log.Println("❌ Query error:", err)
			return c.Status(500).JSON(fiber.Map{"error": "failed to fetch reviews"})
		}
		defer rows.Close()

		for rows.Next() {
			var r object.Review
			var profilePicBytes []byte

			err := rows.Scan(
				&r.ReviewID,
				&r.UserID,
				&r.Username,
				&profilePicBytes,
				&r.ContentID,
				&r.Rating,
				&r.ReviewText,
				&r.ReviewDate,
			)
			if err != nil {
				log.Println("❌ Scan error:", err)
				return c.Status(500).JSON(fiber.Map{"error": "failed to process review"})
			}

			// ✅ แปลง profile_pic เป็น base64 ถ้ามี
			if len(profilePicBytes) > 0 {
				base64Str := base64.StdEncoding.EncodeToString(profilePicBytes)
				r.ProfilePic = "data:image/png;base64," + base64Str
			} else {
				r.ProfilePic = ""
			}

			reviews = append(reviews, r)
		}

		if len(reviews) == 0 {
			return c.Status(404).JSON(fiber.Map{"error": "no reviews found"})
		}

		return c.JSON(reviews)
	}
}
