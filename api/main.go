package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
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
	// redis := config.InitRedis()
	defer db.Close()

	app := fiber.New()
	handler.SetupRoutes(app)

	log.Fatal(app.Listen(":8000"))
}
