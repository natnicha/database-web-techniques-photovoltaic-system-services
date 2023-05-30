package handler

import (
	"github.com/gin-gonic/gin"

	"photovoltaic-system-services/user/controller"
)

type Help interface {
	Delete(context *gin.Context)
	Get(context *gin.Context)
	Login(context *gin.Context)
	Register(context *gin.Context)
	Update(context *gin.Context)
}

func Delete(context *gin.Context) {
	controller.Delete(context)
}

func Get(context *gin.Context) {
	controller.Get(context)
}

func Login(context *gin.Context) {
	controller.Login(context)
}

func Register(context *gin.Context) {
	controller.Create(context)
}

func Update(context *gin.Context) {
	controller.Update(context)
}
