package handler

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jonasOli/url-shortener/api/internal/repository"
	"github.com/jonasOli/url-shortener/api/internal/service"
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

		short_url, err := service.ShortenURL(req.Url)

		if err != nil {
			log.Println("My error: %s", err)
			return c.Status(500).JSON(fiber.Map{"error": "Failed to shorten URL"})
		}

		return c.JSON(fiber.Map{"short_url": short_url})
	})

	app.Get("/:short_code", func(c *fiber.Ctx) error {
		short_code := c.Params("short_code")

		original_url, err := service.GetOriginalURL(short_code)

		if err != nil {
			log.Println("Error on get original url: %s", err)

			return c.Status(500).JSON(fiber.Map{"error": "Failed to get original URL"})
		}

		return c.Redirect(original_url, 301)
	})
}
