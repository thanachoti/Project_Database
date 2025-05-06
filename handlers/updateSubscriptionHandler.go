package handlers

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	// "github.com/khemingkapat/been_chillin/queries"
	"github.com/khemingkapat/been_chillin/auth"
	"github.com/khemingkapat/been_chillin/queries"
)

func UpdateSubscriptionHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the subscription_plan query parameter
		subscriptionPlan := c.Query("subscription_plan")
		if subscriptionPlan == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Missing subscription_plan parameter",
			})
		}

		userIDFromToken, err := auth.ExtractUserID(c)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{"error": "unauthorized"})
		}

		err = queries.UpdateSubscription(db, subscriptionPlan, userIDFromToken)

		// Process the subscription plan as needed
		return c.JSON(fiber.Map{
			"message":           "Subscription updated",
			"subscription_plan": subscriptionPlan,
			"user_id":           userIDFromToken,
		})
	}
}
