package main

import (
	"log"
	"os"

	"github.com/TariqNabeel0/url-shortener/database"
	"github.com/TariqNabeel0/url-shortener/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to laod the env file")

	}

	database.ConnectDB()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()
	r.GET("/pong", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.POST("/shorten", handlers.ShortenURL)
	r.GET("/code", handlers.RedirectOriginal)

	r.Run(":" + port)
}