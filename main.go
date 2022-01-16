package main

import (
	"log"
	"os"

	"github.com/Saranya-jena/golang-jwt-project/routes"
	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

func main() {

	err := gotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading.env file")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	// gin.SetMode(gin.ReleaseMode)

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-1"})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-2"})
	})

	router.Run(":" + port)

}
