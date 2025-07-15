package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"url-shortener/internal/handlers"
	"url-shortener/pkg/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type ShortenPost struct {
	URL string `json:"url" binding:"required"`
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.InitMongo(os.Getenv("mongoURI"))

	redis := database.InitRedis()

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
			"time":   time.Now(),
		})
	})

	handlers.NewURLHandler(r, redis).RegisterRoutes()

	r.Run()
}
