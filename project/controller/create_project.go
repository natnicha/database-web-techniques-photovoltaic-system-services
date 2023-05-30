package controller

import (
	"net/http"
	auth "photovoltaic-system-services/auth/controller"
	"photovoltaic-system-services/project/repositories"
	"time"

	"github.com/gin-gonic/gin"
)

type requestBody struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsPrinted   bool      `json:"is_printed"`
	IsActive    bool      `json:"is_active"`
	StartAt     time.Time `json:"start_at"`
}

func Create(context *gin.Context) {
	var reqBody requestBody
	err := context.BindJSON(&reqBody)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	user, err := auth.GetCurrentUser(context)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	project, err := repositories.CreateProject(
		repositories.Projects{
			Name:        reqBody.Name,
			UserId:      user.Id,
			Description: reqBody.Description,
			IsPrinted:   reqBody.IsPrinted,
			IsActive:    reqBody.IsActive,
			StartAt:     reqBody.StartAt,
		},
	)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"data": project})
}
