package project

import (
	"photovoltaic-system-services/project/controller"

	"github.com/gin-gonic/gin"
)

type Help interface {
	Create(context *gin.Context)
}

func Create(context *gin.Context) {
	controller.Create(context)
}
