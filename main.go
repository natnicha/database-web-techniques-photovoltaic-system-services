package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	database "photovoltaic-system-services/db"
	users "photovoltaic-system-services/users/handler"
)

func main() {
	// TODO : loadEnv()
	godotenv.Load(".env")
	database.Connect()
	serveApplication()
}

func serveApplication() {
	router := gin.Default()
	apiV1 := router.Group("/api/v1")

	user := apiV1.Group("/user")
	user.GET("/:id", users.Get)

	router.Run(":" + os.Getenv("SERVICE_PORT")) // listen and serve on port in .env
	fmt.Println("Server running on port " + os.Getenv("SERVICE_PORT"))
}
