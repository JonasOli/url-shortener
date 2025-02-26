package handler

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func SetupPublicRoutes(app *fiber.App, db *sql.DB, redis *redis.Client) {
	UserRoutes(app, db)
	UlrPublicRoutes(app, db, redis)
}

func SetupPrivateRoutes(app *fiber.App, db *sql.DB, redis *redis.Client) {
	UlrPrivateRoutes(app, db, redis)
}
