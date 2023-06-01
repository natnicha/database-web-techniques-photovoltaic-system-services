package controller

import (
	"net/http"

	"photovoltaic-system-services/user/repositories"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mcuadros/go-defaults"
)

func Update(context *gin.Context) {
	reqBody := new(requestBody)
	defaults.SetDefaults(reqBody)
	err := context.BindJSON(&reqBody)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err = validateStruct(reqBody)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			context.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
			return
		}
	}

	userId, _ := context.Get("user-id")
	existingUser, err := repositories.GetUserByEmail(reqBody.Email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	if existingUser != nil {
		if existingUser.Id != userId {
			context.JSON(http.StatusConflict, gin.H{"error": "The email is already assigned in the system"})
			return
		}
	}

	user, err := prepareUserInfo(*reqBody)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	updatedUser, err := repositories.UpdateUser(userId.(int), user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": updatedUser})
}
