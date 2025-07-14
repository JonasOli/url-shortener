package handlers

import (
	"net/http"
	"url-shortener/internal/repositories"
	"url-shortener/internal/services"

	"github.com/gin-gonic/gin"
)

type URLHandler struct {
	route *gin.Engine
}

type ShortenPost struct {
	URL string `json:"url" binding:"required"`
}

func NewURLHandler(route *gin.Engine) *URLHandler {
	return &URLHandler{route: route}
}

func (h *URLHandler) RegisterRoutes() {
	service := services.NewURLService(repositories.NewURLRepository())

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
