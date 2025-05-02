package queries

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	object "github.com/khemingkapat/been_chillin/objects"
)

func PrepareReview(c *fiber.Ctx) (object.Review, error) {
	var review object.Review

	if err := json.Unmarshal(c.Body(), &review); err != nil {
		return review, err
	}
	log.Println("Prepared Review:", review)
	if review.UserID == 0 || review.ContentID == 0 || review.Rating == 0 || review.ReviewText == "" {
		return review, errors.New("missing required fields")
	}

	review.ReviewDate = time.Now()

	return review, nil
}
