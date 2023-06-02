package handler

import (
	"photovoltaic-system-services/product/controller"

	"github.com/gin-gonic/gin"
)

type Help interface {
	Create(context *gin.Context)
	Delete(context *gin.Context)
	Get(context *gin.Context)
	Update(context *gin.Context)
}

func Create(context *gin.Context) {
	controller.Create(context)
}

func Delete(context *gin.Context) {
	controller.Delete(context)
}

func Get(context *gin.Context) {
	controller.Get(context)
}

func Update(context *gin.Context) {
	controller.Update(context)
}
