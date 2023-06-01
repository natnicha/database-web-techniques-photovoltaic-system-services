package controller

import (
	"fmt"
	"net/http"

	"photovoltaic-system-services/user/repositories"

	"github.com/gin-gonic/gin"
)

func Delete(context *gin.Context) {
	userId, _ := context.Get("user-id")
	err := repositories.DeleteUserById(fmt.Sprint(userId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, nil)
}
