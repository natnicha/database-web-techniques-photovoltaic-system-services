package handler

import (
	controller "photovoltaic-system-services/solar-panel-model/controllers"

	"github.com/gin-gonic/gin"
)

type Help interface {
	Get(context *gin.Context)
}

func Get(context *gin.Context) {
	controller.Get(context)
}
