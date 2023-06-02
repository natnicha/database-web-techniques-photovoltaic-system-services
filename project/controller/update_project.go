package controller

import (
	"net/http"
	"photovoltaic-system-services/project/repositories"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mcuadros/go-defaults"
)

func Update(context *gin.Context) {
	reqBody := new(requestBody)
	defaults.SetDefaults(reqBody)

	err := context.BindJSON(&reqBody)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
	projectId, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	project := repositories.Projects{
		UserId:      userId.(int),
		Name:        reqBody.Name,
		Description: reqBody.Description,
		StartAt:     reqBody.StartAt,
		IsPrinted:   reqBody.IsPrinted,
		UpdateAt:    time.Now(),
	}
	projectNum := repositories.CheckExistProject(projectId, project)
	if projectNum == 0 {
		context.JSON(http.StatusConflict, gin.H{"error": "No project ID with specified user ID"})
		return
	}
	updatedProject, err := repositories.UpdateProject(projectId, project)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": updatedProject})
}
