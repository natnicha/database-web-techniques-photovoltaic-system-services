package handler

import (
	"github.com/gin-gonic/gin"

	"photovoltaic-system-services/users/controller"
)

type Help interface {
	Get(context *gin.Context)
	CreateUser(context *gin.Context)
}

func Get(context *gin.Context) {
	controller.Get(context)
}
