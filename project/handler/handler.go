package project

import (
	"photovoltaic-system-services/project/controller"

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
