package project

import (
	controller "photovoltaic-system-services/weather/controllers"

	"github.com/gin-gonic/gin"
)

type Help interface {
	Daily(context *gin.Context)
	History(context *gin.Context)
}

func Daily(context *gin.Context) {
	controller.Daily(context)
}

func History(context *gin.Context) {
	controller.History(context)
}
