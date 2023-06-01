package controller

import (
	"net/http"

	"photovoltaic-system-services/user/repositories"

	"github.com/gin-gonic/gin"
)

func Delete(context *gin.Context) {
	userId, _ := context.Get("user-id")
	err := repositories.DeleteUserById(userId.(string))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusOK, nil)
}
