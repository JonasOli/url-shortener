package main

import (
	"crypto/rsa"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/jonasOli/url-shortener/api/config"
	"github.com/jonasOli/url-shortener/api/internal/handler"
)

var privateKey *rsa.PrivateKey

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := config.InitPostgres()
	redis := config.InitRedis()
	defer db.Close()

	app := fiber.New()

	handler.SetupPublicRoutes(app, db, redis)

	handler.SetupPrivateRoutes(app, db, redis)

	log.Fatal(app.Listen(":8000"))
}
