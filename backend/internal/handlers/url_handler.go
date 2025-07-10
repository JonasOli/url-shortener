package handlers

import (
	"net/http"
	"url-shortener/internal/repositories"
	"url-shortener/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type URLHandler struct {
	route *gin.Engine
	db    *gorm.DB
}

type ShortenPost struct {
	URL string `json:"url" binding:"required"`
}

func NewURLHandler(route *gin.Engine, db *gorm.DB) *URLHandler {
	return &URLHandler{route: route, db: db}
}

func (h *URLHandler) RegisterRoutes() {
	service := services.NewURLService(repositories.NewURLRepository(h.db))

	h.route.POST("/shorten", func(c *gin.Context) {
		var shortenPost ShortenPost

		if err := c.ShouldBindJSON(&shortenPost); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": " Invalid request body"})
			return
		}

		service.CreateShortURL(shortenPost.URL)

		c.JSON(http.StatusCreated, gin.H{})
	})
}
