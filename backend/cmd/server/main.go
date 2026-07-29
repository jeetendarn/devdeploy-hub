package main

import (
	"log"
	"net/http"


    "github.com/jeetendar/devdeploy-hub/internal/config"
    "github.com/jeetendar/devdeploy-hub/internal/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)
func main() {

	cfg := config.Load()

	if err := database.Connect(cfg); err != nil {
		log.Printf("Database unavailable: %v", err)
		log.Println("Continuing without database...")
	} else {
		log.Println("Database Connected")
	}

	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/api/health", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"service":     cfg.AppName,
			"status":      "healthy",
			"environment": cfg.AppEnv,
			"version":     "1.0.0",
		})

	})

	log.Println("Server running on :" + cfg.AppPort)

	if err := router.Run(":" + cfg.AppPort); err != nil {
		log.Fatal(err)
	}
}