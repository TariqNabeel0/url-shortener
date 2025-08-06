package handlers

import (
	"math/rand"
	"net/http"
	"time"
	"github.com/TariqNabeel0/url-shortener/models"
"github.com/TariqNabeel0/url-shortener/database"

	"github.com/gin-gonic/gin"
)

type ShortenRequest struct {
	URL string `json:"url" binding:"required"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"


var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func generateShortCode(length int) string {
	code := ""

	for i := 0; i < length; i++ {
		randomIndex := seededRand.Intn(len(charset))
		randomChar := string(charset[randomIndex])
		code += randomChar
	}

	return code
}


func ShortenURL(c *gin.Context) {
	var req ShortenRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	shortCode := generateShortCode(6)

	url := models.URL{
		OriginalURL: req.URL,
		ShortCode:   shortCode,
	}

	if err := database.DB.Create(&url).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save URL"})
		return
	}

	host := c.Request.Host
	shortURL := "http://" + host + "/" + shortCode

	c.JSON(http.StatusOK, ShortenResponse{ShortURL: shortURL})
}

func RedirectOriginal(c *gin.Context) {
	shortCode := c.Param("shortcode")

	var url models.URL
	if err := database.DB.Where("short_code = ?", shortCode).First(&url).Error;err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error":"Short URL not found"})
		return
	}

	c.Redirect(http.StatusNotFound, url.OriginalURL)

}