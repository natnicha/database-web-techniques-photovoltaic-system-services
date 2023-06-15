package controller

import (
	"net/http"
	"photovoltaic-system-services/product/repositories"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mcuadros/go-defaults"
)

type requestBody struct {
	ProjectId         int     `json:"project_id" validate:"required"`
	SolarPanelModelId int     `json:"solar_panel_model_id" validate:"required"`
	Orientation       float32 `json:"orientation" validate:"required"`
	Inclination       float32 `json:"inclination" validate:"required"`
	Area              float32 `json:"area" validate:"required"`
	Geolocation       string  `json:"geolocation" validate:"required"`
}

func Create(context *gin.Context) {
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
	productNum := repositories.CheckExistProject(reqBody.ProjectId, userId.(int))
	if productNum == 0 {
		context.JSON(http.StatusConflict, gin.H{"error": "a project ID doesn't belong to a user ID"})
		return
	}
	product, err := repositories.CreateProduct(
		repositories.Product{
			ProjectId:         reqBody.ProjectId,
			SolarPanelModelId: reqBody.SolarPanelModelId,
			Orientation:       reqBody.Orientation,
			Inclination:       reqBody.Inclination,
			Area:              reqBody.Area,
			Geolocation:       reqBody.Geolocation,
		},
	)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"data": product})
}

func validateStruct(obj interface{}) error {
	v := validator.New()
	return v.Struct(obj)
}
