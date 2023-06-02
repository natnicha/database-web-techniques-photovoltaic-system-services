package handler

import (
	"photovoltaic-system-services/product/controller"

	"github.com/gin-gonic/gin"
)

type Help interface {
	Create(context *gin.Context)
	Update(context *gin.Context)
}

func Create(context *gin.Context) {
	controller.Create(context)
}

func Update(context *gin.Context) {
	controller.Update(context)
}
