package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	database "photovoltaic-system-services/db"
	"photovoltaic-system-services/users/handler"
)

func main() {
	// loadEnv()
	// loadDatabase()

	godotenv.Load(".env")
	database.Connect()
	serveApplication()

	// Creates a gin router with default middleware
	// router := gin.Default()

	// A handler for GET request on /example
	// router.GET("/example", func(c *gin.Context) {

	// 	c.JSON(200, gin.H{
	// 		"message": "example",
	// 	}) // gin.H is a shortcut for map[string]interface{}
	// })
	// router.Run(":" + os.Getenv("SERVICE_PORT")) // listen and serve on port 8080
}

func serveApplication() {
	router := gin.Default()

	apiV1 := router.Group("/api/v1")

	user := apiV1.Group("/user")
	user.GET("/", handler.Help)
	// publicRoutes.POST("/register", controller.Register)
	// publicRoutes.POST("/login", controller.Login)

	router.Run(":" + os.Getenv("SERVICE_PORT")) // listen and serve on port 8080
	fmt.Println("Server running on port 8000")
}
