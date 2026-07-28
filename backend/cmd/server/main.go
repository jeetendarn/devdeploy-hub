package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/api/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"service": "DevDeploy Hub",
			"version": "1.0.0",
		})
	})

	log.Println("Server started at http://localhost:8080")

	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}