package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GenerateReport(context *gin.Context) {

	context.JSON(http.StatusAccepted, nil)
	return
}
