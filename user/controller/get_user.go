package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"photovoltaic-system-services/user/repositories"

	"github.com/gin-gonic/gin"
)

func Get(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	user, err := repositories.GetUserById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	if user.Id == 0 {
		context.JSON(http.StatusNotFound, gin.H{"error": "user " + fmt.Sprint(id) + " does not exist"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": user})
}