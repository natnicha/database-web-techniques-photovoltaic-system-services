package controller

import (
	"fmt"
	"net/http"

	"photovoltaic-system-services/user/repositories"

	"github.com/gin-gonic/gin"
)

func Get(context *gin.Context) {
	userId, _ := context.Get("user-id")
	user, err := repositories.GetUserById(userId.(int))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user.Id == 0 {
		context.JSON(http.StatusNotFound, gin.H{"error": "user " + fmt.Sprint(userId) + " does not exist"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": user})
}
