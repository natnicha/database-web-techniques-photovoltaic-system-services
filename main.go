package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	// Creates a gin router with default middleware
	router := gin.Default()

	// A handler for GET request on /example
	router.GET("/example", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "example",
		}) // gin.H is a shortcut for map[string]interface{}
	})
	router.Run(":" + os.Getenv("SERVICE_PORT")) // listen and serve on port 8080
}
