package queries

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/khemingkapat/been_chillin/objects"
)

// ParseProfileData decodes the Fiber request body into a Profile struct and validates it.
func CreateProfile(c *fiber.Ctx) (object.Profile, error) {
	var profile object.Profile

	if err := c.BodyParser(&profile); err != nil {
		return profile, err
	}

	if profile.UserID == 0 || profile.ProfileName == "" {
		return profile, errors.New("missing required fields")
	}

	return profile, nil
}
