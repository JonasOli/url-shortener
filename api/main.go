package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/jonasOli/url-shortener/api/config"
	"github.com/jonasOli/url-shortener/api/internal/handler"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := config.InitPostgres()
	redis := config.InitRedis()
	defer db.Close()

	app := fiber.New()

	app.Use(cors.New())

	handler.SetupPublicRoutes(app, db, redis)

	handler.SetupPrivateRoutes(app, db, redis)

	log.Fatal(app.Listen(":8000"))
}
