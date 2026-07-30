package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/jeetendar/devdeploy-hub/internal/config"
	"github.com/jeetendar/devdeploy-hub/internal/database"

	"github.com/jeetendar/devdeploy-hub/internal/routes"
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

	routes.Register(router)

	log.Println("Server running on :" + cfg.AppPort)

	if err := router.Run(":" + cfg.AppPort); err != nil {
		log.Fatal(err)
	}
}
