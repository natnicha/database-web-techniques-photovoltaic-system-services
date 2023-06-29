package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	Database "photovoltaic-system-services/db"
	Middleware "photovoltaic-system-services/middleware"
	Product "photovoltaic-system-services/product/handlers"
	Project "photovoltaic-system-services/project/handlers"
	SolarPanel "photovoltaic-system-services/solar-panel-model/handlers"
	User "photovoltaic-system-services/user/handlers"
	Weather "photovoltaic-system-services/weather/handlers"
)

func main() {
	godotenv.Load(".env")
	Database.Connect()
	defer Database.Close()
	router := serveApplication()
	doGracefulShutdown(router)
}

func serveApplication() *gin.Engine {
	router := gin.Default()
	router.Use(Middleware.CORSMiddleware())
	auth := router.Group("/auth")
	auth.POST("/register", User.Register)
	auth.POST("/login", User.Login)

	apiV1 := router.Group("/api/v1")
	apiV1.Use(Middleware.JWTAuthMiddleware())
	user := apiV1.Group("/user")
	user.GET("/", User.Get)
	user.DELETE("/delete", User.Delete)
	auth.POST("/logout", User.Logout)
	user.PUT("/update", User.Update)

	project := apiV1.Group("/project")
	project.GET("/", Project.Get)
	project.POST("/create", Project.Create)
	project.DELETE("/delete/:id", Project.Delete)
	project.PUT("/update/:id", Project.Update)
	project.POST("/generate-report/:id", Project.GenerateReport)

	solarPanel := apiV1.Group("/solar-panel-model")
	solarPanel.GET("/", SolarPanel.Get)

	product := apiV1.Group("/product")
	product.GET("/", Product.Get)
	product.POST("/create", Product.Create)
	product.DELETE("/delete/:id", Product.Delete)
	product.PUT("/update/:id", Product.Update)
	product.POST("/generate-report/:id", Product.GenerateReport)

	weather := apiV1.Group("/weather")
	weather.POST("/daily", Weather.Daily)
	weather.POST("/history", Weather.History)

	return router
}

func doGracefulShutdown(router *gin.Engine) {
	srv := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%s", os.Getenv("SERVICE_PORT")),
	}
	fmt.Println("Server running on port " + os.Getenv("SERVICE_PORT"))

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	// Create channel for shutdown signals.
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	signal.Notify(stop, syscall.SIGTERM)

	//Recieve shutdown signals.
	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("error shutting down server %s", err)
	} else {
		log.Println("Server gracefully stopped")
	}
}
