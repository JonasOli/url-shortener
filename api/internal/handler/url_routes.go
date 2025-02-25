package handler

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jonasOli/url-shortener/api/internal/repository"
	"github.com/jonasOli/url-shortener/api/internal/service"
	"github.com/jonasOli/url-shortener/api/internal/utils"
	"github.com/redis/go-redis/v9"
)

func UlrRoutes(app *fiber.App, db *sql.DB, redis *redis.Client) {
	repo := repository.NewURLRepository(db, redis)
	service := service.NewURLService(repo)

	app.Post("/shorten", func(c *fiber.Ctx) error {
		var req struct {
			Url string `json:"url"`
		}

		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}

		user_name := utils.GetUserName(c)

		short_url, err := service.ShortenURL(req.Url, user_name)

		if err != nil {
			log.Error(err)
			return c.Status(500).JSON(fiber.Map{"error": "Failed to shorten URL"})
		}

		return c.JSON(fiber.Map{"short_url": short_url})
	})

	// Make this route public
	app.Get("/:short_code", func(c *fiber.Ctx) error {
		short_code := c.Params("short_code")

		original_url, err := service.GetOriginalURL(short_code)

		if err != nil {
			log.Error("Error on get original url: %s", err)

			return c.Status(500).JSON(fiber.Map{"error": "Failed to get original URL"})
		}

		return c.Redirect(original_url, 301)
	})
}
