package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	Database "photovoltaic-system-services/db"
	Middleware "photovoltaic-system-services/middleware"
	Project "photovoltaic-system-services/project/handler"
	SolarPanel "photovoltaic-system-services/solar-panel-model/handler"
	User "photovoltaic-system-services/user/handler"
)

func main() {
	// TODO : loadEnv()
	godotenv.Load(".env")
	Database.Connect()
	serveApplication()
}

func serveApplication() {
	router := gin.Default()
	auth := router.Group("/auth")
	auth.POST("/register", User.Register)
	auth.POST("/login", User.Login)

	apiV1 := router.Group("/api/v1")
	apiV1.Use(Middleware.JWTAuthMiddleware())
	user := apiV1.Group("/user")
	user.GET("/", User.Get)
	user.PUT("/update", User.Update)
	user.DELETE("/delete", User.Delete)

	project := apiV1.Group("/project")
	project.POST("/create", Project.Create)
	project.PUT("/update/:id", Project.Update)
	project.DELETE("/delete/:id", Project.Delete)
	project.GET("/", Project.Get)

	solarPanel := apiV1.Group("/solar-panel-model")
	solarPanel.GET("/", SolarPanel.Get)
	router.Run(":" + os.Getenv("SERVICE_PORT")) // listen and serve on port in .env
	fmt.Println("Server running on port " + os.Getenv("SERVICE_PORT"))
}
