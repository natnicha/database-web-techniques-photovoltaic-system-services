package controller

import (
	"net/http"

	"photovoltaic-system-services/user/repositories"

	"github.com/gin-gonic/gin"
)

func Delete(context *gin.Context) {
	id := context.Param("id")
	err := repositories.DeleteUserById(id)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusOK, nil)
}
