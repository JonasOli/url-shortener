package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"url-shortener/internal/handlers"
	"url-shortener/pkg/database"

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

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
			"time":   time.Now(),
		})
	})

	handlers.NewURLHandler(r).RegisterRoutes()

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
