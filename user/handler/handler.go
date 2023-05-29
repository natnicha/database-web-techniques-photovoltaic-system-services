package handler

import (
	"github.com/gin-gonic/gin"

	"photovoltaic-system-services/user/controller"
)

type Help interface {
	Get(context *gin.Context)
	CreateUser(context *gin.Context)
}

func Get(context *gin.Context) {
	controller.Get(context)
}
