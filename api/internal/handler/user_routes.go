package handler

import (
	"database/sql"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jonasOli/url-shortener/api/internal/repository"
	"github.com/jonasOli/url-shortener/api/internal/service"
	"github.com/redis/go-redis/v9"
)

func UserRoutes(app *fiber.App, db *sql.DB, redis *redis.Client) {
	repo := repository.NewUserRepository(db, redis)
	service := service.NewUserService(repo)

	app.Post("/user/signup", func(c *fiber.Ctx) error {
		var req struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}

		session_key, err := service.Signup(req.Name, req.Email, req.Password)

		if err != nil {
			log.Errorf("%s", err)
			return c.Status(500).JSON(fiber.Map{"error": "Failed to create user"})
		}

		c.Cookie(&fiber.Cookie{
			Name:     "session_id",
			Value:    session_key,
			Expires:  time.Now().Add(time.Hour),
			HTTPOnly: true,
			Secure:   true,
			SameSite: "Strict",
		})

		return c.SendStatus(201)
	})

	app.Post("/user/login", func(c *fiber.Ctx) error {
		var req struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}

		session_key, err := service.Signin(req.Email, req.Password)

		if err != nil {
			return err
		}

		c.Cookie(&fiber.Cookie{
			Name:     "session_id",
			Value:    session_key,
			Expires:  time.Now().Add(time.Hour),
			HTTPOnly: true,
			Secure:   true,
			SameSite: "Strict",
		})

		return c.SendStatus(fiber.StatusOK)
	})

	app.Post("/user/signout", func(c *fiber.Ctx) error {
		session_key := c.Cookies("session_id")

		if session_key == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "No session",
			})
		}

		err := service.Signout(session_key)

		if err != nil {
			log.Errorf("%s", err)

			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.SendStatus(fiber.StatusOK)
	})
}
