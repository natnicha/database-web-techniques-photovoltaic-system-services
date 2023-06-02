package handler

import (
	"photovoltaic-system-services/product/controller"

	"github.com/gin-gonic/gin"
)

type Help interface {
	Create(context *gin.Context)
}

func Create(context *gin.Context) {
	controller.Create(context)
}
