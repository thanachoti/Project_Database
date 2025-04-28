package queries

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	object "github.com/khemingkapat/been_chillin/objects"
)

func PrepareReviewData(c *fiber.Ctx) (object.Review, error) {
	var review object.Review

	if err := json.Unmarshal(c.Body(), &review); err != nil {
		return review, err
	}

	if review.UserID == 0 || review.ContentID == 0 || review.Rating == 0 || review.ReviewText == "" {
		return review, errors.New("missing required fields")
	}

	review.ReviewDate = time.Now()

	return review, nil
}
