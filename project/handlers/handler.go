package project

import (
	controller "photovoltaic-system-services/project/controllers"

	"github.com/gin-gonic/gin"
)

type Help interface {
	Create(context *gin.Context)
	Delete(context *gin.Context)
	GenerateReport(context *gin.Context)
	Get(context *gin.Context)
	Update(context *gin.Context)
}

func Create(context *gin.Context) {
	controller.Create(context)
}

func Delete(context *gin.Context) {
	controller.Delete(context)
}

func GenerateReport(context *gin.Context) {
	controller.GenerateReport(context)
}

func Get(context *gin.Context) {
	controller.Get(context)
}

func Update(context *gin.Context) {
	controller.Update(context)
}
