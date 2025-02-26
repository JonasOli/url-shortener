package handler

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
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

		err := service.Signup(req.Name, req.Email, req.Password)

		if err != nil {
			log.Println(err)
			return c.Status(500).JSON(fiber.Map{"error": "Failed to create user"})
		}

		return c.SendStatus(200)
	})

	app.Post("/user/login", func(c *fiber.Ctx) error {
		var req struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}

		token, err := service.Signin(req.Email, req.Password)

		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{"token": token})
	})
}
