package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	Database "photovoltaic-system-services/db"
	Middleware "photovoltaic-system-services/middleware"
	Product "photovoltaic-system-services/product/handler"
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
	user.DELETE("/delete", User.Delete)
	user.PUT("/update", User.Update)

	project := apiV1.Group("/project")
	project.GET("/", Project.Get)
	project.POST("/create", Project.Create)
	project.DELETE("/delete/:id", Project.Delete)
	project.PUT("/update/:id", Project.Update)

	solarPanel := apiV1.Group("/solar-panel-model")
	solarPanel.GET("/", SolarPanel.Get)

	product := apiV1.Group("/product")
	product.GET("/", Product.Get)
	product.POST("/create", Product.Create)
	product.DELETE("/delete/:id", Product.Delete)
	product.PUT("/update/:id", Product.Update)

	router.Run(":" + os.Getenv("SERVICE_PORT")) // listen and serve on port in .env
	fmt.Println("Server running on port " + os.Getenv("SERVICE_PORT"))
}
