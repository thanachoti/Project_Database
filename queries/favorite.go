package queries

import (
	"encoding/json"
	"errors"

	"github.com/gofiber/fiber/v2"
	object "github.com/khemingkapat/been_chillin/objects"
)

func ParseFavoriteData(c *fiber.Ctx) (object.Favorite, error) {
	var favorite object.Favorite

	if err := json.Unmarshal(c.Body(), &favorite); err != nil {
		return favorite, err
	}

	if favorite.UserID == 0 || favorite.ContentID == 0 {
		return favorite, errors.New("missing required fields")
	}

	return favorite, nil
}
