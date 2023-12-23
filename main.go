package main

import (
	"os"

	"github.com/gin-gonic/gin"
	routes "github.com/harsh082ip/JWT-AUTHORIZATION-golang/routes"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8001"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-1"})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access grnated for api-2"})
	})

}
