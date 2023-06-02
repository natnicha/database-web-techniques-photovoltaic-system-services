package controller

import (
	"net/http"
	"photovoltaic-system-services/project/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Delete(context *gin.Context) {
	projectId, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, _ := context.Get("user-id")
	projectNum := repositories.CheckExistProject(projectId, repositories.Projects{UserId: userId.(int)})
	if projectNum == 0 {
		context.JSON(http.StatusConflict, gin.H{"error": "No specified product ID or a project ID doesn't belong to a user ID"})
		return
	}
	err = repositories.DeleteProjectById(projectId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, nil)
}
