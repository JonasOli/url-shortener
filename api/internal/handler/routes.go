package handler

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func SetupRoutes(app *fiber.App, db *sql.DB, redis *redis.Client) {
	UlrRoutes(app, db, redis)
}
