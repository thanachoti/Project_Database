package queries

import (
	"encoding/json"
	"errors"

	"github.com/gofiber/fiber/v2"
	object "github.com/khemingkapat/been_chillin/objects"
)

func ParseWatchHistoryData(c *fiber.Ctx) (object.WatchHistory, error) {
	var history object.WatchHistory

	if err := json.Unmarshal(c.Body(), &history); err != nil {
		return history, err
	}

	if history.UserID == 0 || history.ContentID == 0 {
		return history, errors.New("missing required fields")
	}

	return history, nil
}
