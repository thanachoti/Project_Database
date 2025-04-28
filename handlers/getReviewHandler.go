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
		review, err := queries.PrepareReviewData(c)
		if err != nil {
			log.Println("Parse error:", err)
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

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

// GetAllReviewsHandler retrieves all reviews
func GetAllReviewsHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		rows, err := db.Query(`SELECT review_id, user_id, content_id, rating, review_text, review_date FROM REVIEW`)
		if err != nil {
			log.Println("Query error:", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "failed to fetch reviews"})
		}
		defer rows.Close()

		var reviews []object.Review
		for rows.Next() {
			var r object.Review
			if err := rows.Scan(&r.ReviewID, &r.UserID, &r.ContentID, &r.Rating, &r.ReviewText, &r.ReviewDate); err != nil {
				log.Println("Row scan error:", err)
				continue
			}
			reviews = append(reviews, r)
		}

		return c.JSON(reviews)
	}
}

// GetReviewByIDHandler retrieves a specific review by review_id
func GetReviewByIDHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		var r object.Review
		err := db.QueryRow(`SELECT review_id, user_id, content_id, rating, review_text, review_date FROM REVIEW WHERE review_id = $1`, id).
			Scan(&r.ReviewID, &r.UserID, &r.ContentID, &r.Rating, &r.ReviewText, &r.ReviewDate)
		if err != nil {
			if err == sql.ErrNoRows {
				return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "review not found"})
			}
			log.Println("Query error:", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "failed to fetch review"})
		}

		return c.JSON(r)
	}
}
