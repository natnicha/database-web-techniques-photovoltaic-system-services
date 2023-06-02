package controller

import (
	"net/http"
	"photovoltaic-system-services/product/repositories"
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

	product := repositories.Product{
		ProjectId:         reqBody.ProjectId,
		SolarPanelModelId: reqBody.SolarPanelModelId,
		Orientation:       reqBody.Orientation,
		Inclination:       reqBody.Inclination,
		Area:              reqBody.Area,
		Geolocation:       reqBody.Geolocation,
		UpdateAt:          time.Now(),
	}
	userId, _ := context.Get("user-id")
	projectNum := repositories.CheckExistProject(reqBody.ProjectId, userId.(int))
	if projectNum == 0 {
		context.JSON(http.StatusConflict, gin.H{"error": "a project ID doesn't belong to a user ID"})
		return
	}
	productId, err := strconv.Atoi(context.Param("id"))
	productNum := repositories.CheckExistProduct(productId)
	if productNum == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "No specified product ID"})
		return
	}
	updatedProject, err := repositories.UpdateProject(productId, product)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": updatedProject})
}
