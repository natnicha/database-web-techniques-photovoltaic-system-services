package handler

import (
	"github.com/gin-gonic/gin"

	"photovoltaic-system-services/user/controller"
)

type Help interface {
	Get(context *gin.Context)
	Create(context *gin.Context)
}

func Get(context *gin.Context) {
	controller.Get(context)
}

func Create(context *gin.Context) {
	controller.Create(context)
}
