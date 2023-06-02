package handler

import (
	"photovoltaic-system-services/solar-panel-model/controller"

	"github.com/gin-gonic/gin"
)

type Help interface {
	Get(context *gin.Context)
}

func Get(context *gin.Context) {
	controller.Get(context)
}
