package controller

import (
	"net/http"

	"photovoltaic-system-services/user/repositories"

	"github.com/gin-gonic/gin"
)

func Get(context *gin.Context) {
	id := context.Param("id")
	user, err := repositories.GetUserById(id)
	if err != nil {
		context.JSON(http.StatusCreated, gin.H{"err": err})
		return
	}

	if len(user) == 0 {
		context.JSON(http.StatusOK, gin.H{"user": "user " + id + " does not exists"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"user": user})

}
