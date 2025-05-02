package handlers

import (
	"database/sql"
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

		log.Println("Review:", review)
		query := `INSERT INTO REVIEW (user_id, content_id, rating, review_text, review_date)
				  VALUES ($1, $2, $3, $4, $5) RETURNING review_id`

		err = db.QueryRow(query, review.UserID, review.ContentID, review.Rating, review.ReviewText, review.ReviewDate).
			Scan(&review.ReviewID)
		if err != nil {
			log.Println("DB error:", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "failed to insert review"})
		}

		return c.Status(http.StatusCreated).JSON(review)
	}
}

// GetReviewByContentIDHandler retrieves reviews for a specific content_id
func GetReviewByContentIDHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// ดึง content_id จากพารามิเตอร์ใน URL
		contentID := c.Params("content_id")

		// สร้างตัวแปรเพื่อเก็บรีวิวทั้งหมดที่ตรงกับ content_id
		var reviews []object.Review

		// Query ฐานข้อมูลเพื่อค้นหาทุกรีวิวที่ตรงกับ content_id
		rows, err := db.Query(`
			SELECT review_id, user_id, content_id, rating, review_text, review_date 
			FROM REVIEW 
			WHERE content_id = $1`, contentID)

		// ถ้ามีข้อผิดพลาดในการ query
		if err != nil {
			log.Println("Query error:", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "failed to fetch reviews"})
		}
		defer rows.Close() // ปิด rows เมื่อเสร็จการใช้งาน

		// ดึงข้อมูลรีวิวจาก rows
		for rows.Next() {
			var r object.Review
			err := rows.Scan(&r.ReviewID, &r.UserID, &r.ContentID, &r.Rating, &r.ReviewText, &r.ReviewDate)
			if err != nil {
				log.Println("Error scanning row:", err)
				return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "failed to process review"})
			}
			reviews = append(reviews, r)
		}

		// ถ้าไม่พบรีวิวที่ตรงกับ content_id
		if len(reviews) == 0 {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "no reviews found for this content"})
		}

		// ส่งผลลัพธ์รีวิวทั้งหมดที่พบ
		return c.JSON(reviews)
	}
}
