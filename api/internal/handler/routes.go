package handler

import (
	"crypto/rsa"
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func SetupPublicRoutes(app *fiber.App, db *sql.DB, redis *redis.Client, privateKey *rsa.PrivateKey) {
	UserRoutes(app, db, privateKey)
}

func SetupPrivateRoutes(app *fiber.App, db *sql.DB, redis *redis.Client) {
	UlrRoutes(app, db, redis)
}
