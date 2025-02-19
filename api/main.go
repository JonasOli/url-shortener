package main

import (
	"crypto/rand"
	"crypto/rsa"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/jonasOli/url-shortener/api/config"
	"github.com/jonasOli/url-shortener/api/internal/handler"

	jwtware "github.com/gofiber/contrib/jwt"
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

	privateKey, err = rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		log.Fatalf("rsa.GenerateKey: %v", err)
	}

	handler.SetupPublicRoutes(app, db, redis, privateKey)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwtware.RS256,
			Key:    privateKey.Public(),
		},
	}))

	handler.SetupPrivateRoutes(app, db, redis)

	log.Fatal(app.Listen(":8000"))
}
