package utils

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func AuthMiddleware(redis_client *redis.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		session_key := c.Cookies("session-id")

		if session_key == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "No session",
			})
		}

		ctx := context.Background()
		user_id, err := redis_client.Get(ctx, session_key).Result()

		if err == redis.Nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or expired session",
			})
		} else if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error on session verification",
			})
		}

		c.Locals("user_id", user_id)

		return c.Next()
	}
}
