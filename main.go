package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	database "photovoltaic-system-services/db"
	middleware "photovoltaic-system-services/middleware"
	users "photovoltaic-system-services/user/handler"
)

func main() {
	// TODO : loadEnv()
	godotenv.Load(".env")
	database.Connect()
	serveApplication()
}

func serveApplication() {
	router := gin.Default()
	auth := router.Group("/auth")
	auth.POST("/register", users.Register)
	auth.POST("/login", users.Login)

	apiV1 := router.Group("/api/v1")
	apiV1.Use(middleware.JWTAuthMiddleware())
	user := apiV1.Group("/user")
	user.GET("/:id", users.Get)
	user.PUT("/update/:id", users.Update)
	user.DELETE("/delete/:id", users.Delete)

	router.Run(":" + os.Getenv("SERVICE_PORT")) // listen and serve on port in .env
	fmt.Println("Server running on port " + os.Getenv("SERVICE_PORT"))
}
