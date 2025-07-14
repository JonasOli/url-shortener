package handlers

import (
	"net/http"
	"url-shortener/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type URLHandler struct {
	route *gin.Engine
	redis *redis.Client
}

type ShortenPost struct {
	URL string `json:"url" binding:"required"`
}

func NewURLHandler(route *gin.Engine, redis *redis.Client) *URLHandler {
	return &URLHandler{route: route, redis: redis}
}

func (h *URLHandler) RegisterRoutes() {
	service := services.NewURLService(h.redis)

	h.route.POST("/shorten", func(c *gin.Context) {
		var shortenPost ShortenPost

		if err := c.ShouldBindJSON(&shortenPost); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": " Invalid request body"})
			return
		}

		service.CreateShortURL(shortenPost.URL)

		c.JSON(http.StatusCreated, gin.H{})
	})

	h.route.GET("/:shortCode", func(c *gin.Context) {
		shortCode := c.Param("shortCode")

		originalUrl, err := service.FindByShortCode(shortCode)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
			return
		}

		c.Redirect(http.StatusFound, originalUrl)
	})

}
